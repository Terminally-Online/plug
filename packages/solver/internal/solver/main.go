package solver

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/plug_router"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3"
	"solver/internal/actions/ens"
	"solver/internal/actions/euler"
	"solver/internal/actions/morpho"
	"solver/internal/actions/nouns"
	"solver/internal/actions/plug"
	"solver/internal/actions/yearn_v3"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/database/types"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Solver struct {
	Protocols map[string]actions.BaseProtocolHandler
	IsKilled  bool
}

func New() Solver {
	return Solver{
		Protocols: map[string]actions.BaseProtocolHandler{
			actions.ProtocolPlug:    plug.New(),
			actions.ProtocolAaveV3:  aave_v3.New(),
			actions.ProtocolYearnV3: yearn_v3.New(),
			actions.ProtocolENS:     ens.New(),
			actions.ProtocolNouns:   nouns.New(),
			actions.ProtocolMorpho:  morpho.New(),
			actions.ProtocolEuler:   euler.New(),
		},
		IsKilled: false,
	}
}

func (s *Solver) GetTransaction(rawInputs json.RawMessage, chainId uint64, from common.Address) ([]signature.Plug, error) {
	var inputs struct {
		Protocol string `json:"protocol"`
		Action   string `json:"action"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	handler, exists := s.Protocols[inputs.Protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", inputs.Protocol)
	}

	client, err := client.New(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err)
	}

	params := actions.HandlerParams{
		Client:  client,
		ChainId: chainId,
		From:    from,
	}

	transactions, err := handler.GetTransaction(inputs.Action, rawInputs, params)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		if transactions[i].Value == nil {
			transactions[i].Value = big.NewInt(0)
		}
		// TODO: Only include the gas amount when we can properly estimate it with the traces
		//       that are generated from the simulation.
		transactions[i].Gas = big.NewInt(600000)
	}

	return transactions, nil
}

func (s *Solver) GetPlugsArray(
	head []signature.Plug,
	inputs []byte,
	chainId uint64,
	from common.Address,
) (plugs []signature.Plug, exclusive bool, error error) {
	plugs, err := s.GetTransaction(inputs, chainId, from)
	if err != nil {
		return nil, false, err
	}

	// NOTE: Some plug actions have exclusive transactions that need to be run alone
	//       before the rest of the Plug can run. For this, we will just break out
	//       of the loop and execute any solo transactions that are needed for
	//       the rest of the batch to run in sequence.
	for _, plug := range plugs {
		if plug.Exclusive {
			// NOTE: Set the field to false to avoid tarnishing the response shape.
			plug.Exclusive = false
			return []signature.Plug{plug}, true, nil
		}
	}

	return append(head, plugs...), false, nil
}

func (s *Solver) GetPlugs(intent *models.Intent) ([]signature.Plug, error) {
	var plugs []signature.Plug
	for _, input := range intent.Inputs {
		inputsMap := map[string]interface{}{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		for k, v := range input {
			inputsMap[k] = v
		}
		inputs, err := json.Marshal(inputsMap)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}

		var exclusive bool
		plugs, exclusive, err = s.GetPlugsArray(plugs, inputs, intent.ChainId, common.HexToAddress(intent.From))
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}
		if exclusive {
			break
		}
	}

	// NOTE: If there was no transaction to execute we will return a warning because
	//		 we will be halting the simulation of this workflow.
	if len(plugs) == 0 {
		return nil, utils.ErrBuild("no transactions to execute")
	}

	return plugs, nil
}

func (s *Solver) GetLivePlugs(plugs []signature.Plug, chainId uint64, from string) (*signature.LivePlugs, error) {
	solver, err := signature.GetSolverHash()
	if err != nil {
		return nil, err
	}
	fromAddress := common.HexToAddress(from)
	salt, err := signature.GetSaltHash(fromAddress)
	if err != nil {
		return nil, err
	}

	plugsSigned, plugsSignature, err := signature.GetSignature(
		big.NewInt(int64(chainId)),
		fromAddress,
		signature.Plugs{
			Socket: fromAddress,
			Plugs:  plugs,
			Solver: solver,
			Salt:   salt,
		},
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to sign: " + err.Error())
	}

	return &signature.LivePlugs{
		Plugs:     plugsSigned,
		Signature: plugsSignature,
	}, nil
}

func (s *Solver) BuildLivePlugModels(intent *models.Intent, livePlugs signature.LivePlugs) (transactionBundle *models.LivePlug, err error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("PlugRouter")
	}

	// TODO MASON: we should be able to use the pack method with translated models.Transactions instead of signature.Plugs to be sure it's always backwards compatible.
	plugCalldata, err := routerAbi.Pack("plug", livePlugs)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	identifier := []byte("plug")
	plugs := make([]models.Plug, len(livePlugs.Plugs.Plugs))
	for idx, plug := range livePlugs.Plugs.Plugs {
		data := append(plug.Data, identifier...)
		plugModel := models.Plug{
			From:      intent.From,
			To:        plug.To.Hex(),
			Data:      hexutil.Bytes(data).String(),
			Value:     &types.BigInt{Int: plug.Value},
			Gas:       &types.BigInt{Int: plug.Gas},
			Exclusive: plug.Exclusive,
		}
		plugs[idx] = plugModel
	}

	data := append(plugCalldata, identifier...)
	signature := hexutil.Bytes(livePlugs.Signature).String()
	bundle := models.LivePlug{
		IntentId:  intent.Id,
		Signature: &signature,
		Plugs:     plugs,
		Data:      hexutil.Bytes(data).String(),
		ChainId:   intent.ChainId,
		From:      intent.From,
		To:        references.Networks[intent.ChainId].References["plug"]["router"],
	}

	return &bundle, nil
}

func (s *Solver) RebuildSolutionFromModels(intent *models.Intent) (*Solution, error) {
	// Get the latest LivePlug for this intent
	var livePlug models.LivePlug
	if err := database.DB.Where("intent_id = ?", intent.Id).
		Order("created_at DESC").
		First(&livePlug).Error; err != nil {
		return nil, fmt.Errorf("failed to find live plug: %v", err)
	}

	// Get the latest Run for this intent
	var run models.Run
	if err := database.DB.Where("intent_id = ? AND live_plug_id = ?", intent.Id, livePlug.Id).
		Order("created_at DESC").
		First(&run).Error; err != nil {
		return nil, fmt.Errorf("failed to find run: %v", err)
	}

	// Get the associated plugs
	var plugs []models.Plug
	if err := database.DB.Where("bundle_id = ?", livePlug.Id).
		Find(&plugs).Error; err != nil {
		return nil, fmt.Errorf("failed to find plugs: %v", err)
	}
	livePlug.Plugs = plugs

	// Reconstruct LivePlugs from the stored signature and plugs
	var livePlugs *signature.LivePlugs
	if livePlug.Signature != nil {
		// Only reconstruct LivePlugs if this wasn't an EOA transaction
		signatureBytes := []byte(*livePlug.Signature)

		// Convert models.Plug to signature.Plug
		signaturePugs := make([]signature.Plug, len(plugs))
		for i, plug := range plugs {
			signaturePugs[i] = signature.Plug{
				To:        common.HexToAddress(plug.To),
				Value:     plug.Value.Int,
				Gas:       plug.Gas.Int,
				Data:      []byte(plug.Data),
				Exclusive: plug.Exclusive,
			}
		}

		livePlugs = &signature.LivePlugs{
			Signature: signatureBytes,
			Plugs: signature.Plugs{
				Socket: common.HexToAddress(intent.From),
				Plugs:  signaturePugs,
			},
		}
	}

	return &Solution{
		Status:       SolutionStatus{Success: run.Status == "success"},
		Transactions: &livePlug.Plugs,
		LivePlugs:    livePlugs,
		Intent:       intent,
		Run:          &run,
		Transaction:  &livePlug,
	}, nil
}

func (s *Solver) SolveEOA(intent *models.Intent) (solution *Solution, err error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	// TODO: How do we handle multiple transactions for an EOA?
	if len(plugs) > 1 {
		return nil, utils.ErrField("plugs", "eoa can only run one transaction at a time")
	}

	identifier := []byte("plug")
	data := append(plugs[0].Data, identifier...)

	plugModels := make([]models.Plug, len(plugs))
	plugModels[0] = models.Plug{
		From:      intent.From,
		To:        plugs[0].To.Hex(),
		Data:      hexutil.Bytes(data).String(),
		Value:     &types.BigInt{Int: plugs[0].Value},
		Gas:       &types.BigInt{Int: plugs[0].Gas},
		Exclusive: plugs[0].Exclusive,
	}

	livePlug := models.LivePlug{
		IntentId: intent.Id,
		ChainId:  intent.ChainId,
		From:     intent.From,
		To:       plugs[0].To.Hex(),
		Value:    &types.BigInt{Int: plugs[0].Value},
		Gas:      &types.BigInt{Int: plugs[0].Gas},
		Data:     hexutil.Bytes(data).String(),
		Plugs:    plugModels,
	}

	if err := database.DB.Create(&livePlug).Error; err != nil {
		return nil, fmt.Errorf("failed to save transaction bundle: %v", err)
	}

	var run *models.Run
	if simulate, ok := intent.Options["simulate"].(bool); ok && simulate {
		run, err = simulation.SimulateRaw(&livePlug, nil)
		if err != nil {
			return nil, err
		}
	}
	run.IntentId = intent.Id

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save simulation run: %v", err)
	}

	return &Solution{
		Status:       SolutionStatus{Success: true},
		Transactions: &plugModels,
		Intent:       intent,
		Run:          run,
		Transaction:  &livePlug,
	}, nil
}

func (s *Solver) Solve(intent *models.Intent) (solution *Solution, err error) {
	if isEOA, ok := intent.Options["isEOA"].(bool); ok && isEOA {
		return s.SolveEOA(intent)
	}

	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	livePlugs, err := s.GetLivePlugs(plugs, intent.ChainId, intent.From)
	if err != nil {
		return nil, err
	}

	livePlugModel, err := s.BuildLivePlugModels(intent, *livePlugs)
	if err != nil {
		return nil, err
	}

	if err := database.DB.Create(livePlugModel).Error; err != nil {
		return nil, fmt.Errorf("failed to save transaction bundle: %v", err)
	}

	var run *models.Run
	if simulate, ok := intent.Options["simulate"].(bool); ok && simulate {
		run, err = simulation.Simulate(livePlugModel)
		if err != nil {
			return nil, err
		}
	}

	run.IntentId = intent.Id
	intent.PeriodEndAt, intent.NextSimulationAt = intent.GetNextSimulationAt()

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save simulation run: %v", err)
	}

	if err := database.DB.Model(&intent).Updates(map[string]interface{}{
		"period_end_at":      intent.PeriodEndAt,
		"next_simulation_at": intent.NextSimulationAt,
	}).Error; err != nil {
		return nil, fmt.Errorf("failed to update intent: %v", err)
	}

	return &Solution{
		Status:       SolutionStatus{Success: true},
		Transactions: &livePlugModel.Plugs, // Transactions in the `livePlug`.
		LivePlugs:    livePlugs,            // The `livePlug` included in the bundle.
		Intent:       intent,               // Intent the solver built from.
		Run:          run,                  // Simulation results of solver run.
		Transaction:  livePlugModel,        // Transaction the solver runs.
	}, nil
}

func (s *Solver) Submit(intents []models.Intent) ([]signature.Result, error) {
	if len(intents) == 0 {
		return nil, utils.ErrBuild("no plugs generated to execute")
	}

	chainId := intents[0].ChainId
	errors := make([]error, len(intents))

	var livePlugs []*signature.LivePlugs
	for i, intent := range intents {
		if intent.ChainId != chainId {
			errors[i] = utils.ErrChainId("chainId", intent.ChainId)
			continue
		}

		solution, err := s.Solve(&intent)
		if err != nil {
			errors[i] = err
			continue
		}

		if submit, ok := intent.Options["submit"].(bool); ok && submit && solution.Run.Status == "success" {
			livePlugs = append(livePlugs, solution.LivePlugs)
		}
	}

	provider, err := client.New(chainId)
	if err != nil {
		return nil, err
	}
	results, err := provider.Plug(livePlugs)
	if err != nil {
		return nil, err
	}

	return results, nil
}
