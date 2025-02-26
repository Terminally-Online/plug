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

func (s *Solver) GetPlugs(intent models.Intent) ([]signature.Plug, error) {
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

func (s *Solver) GetLivePlugs(intent models.Intent) (*signature.LivePlugs, error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}
	solver, err := signature.GetSolverHash()
	if err != nil {
		return nil, err
	}
	from := common.HexToAddress(intent.From)
	salt, err := signature.GetSaltHash(from)
	if err != nil {
		return nil, err
	}

	plugsSigned, plugsSignature, err := signature.GetSignature(
		big.NewInt(int64(intent.ChainId)),
		from,
		signature.Plugs{
			Socket: from,
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

func (s *Solver) BuildPlugTransaction(intent models.Intent, livePlugs signature.LivePlugs) (transaction *simulation.Transaction, err error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("PlugRouter")
	}

	plugCalldata, err := routerAbi.Pack("plug", livePlugs)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	transaction = &simulation.Transaction{
		From:    intent.From,
		To:      references.Networks[intent.ChainId].References["plug"]["router"],
		ChainId: intent.ChainId,
		Value:   hexutil.EncodeBig(intent.Value),
		Data:    hexutil.Bytes.String(plugCalldata),
	}

	if intent.GasLimit != nil {
		gasLimitStr := hexutil.EncodeUint64(*intent.GasLimit)
		transaction.Gas = &gasLimitStr
	}

	return transaction, nil
}

func (s *Solver) SolveEOA(intent models.Intent) (solution *Solution, err error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	// TODO: How do we handle multiple transactions for an EOA?
	if len(plugs) > 1 {
		return nil, utils.ErrField("plugs", "eoa can only run one transaction at a time")
	}

	transaction := simulation.Transaction{
		From:    intent.From,
		ChainId: intent.ChainId,
		To:      plugs[0].To.Hex(),
		Value:   plugs[0].Value.String(),
		Data:    hexutil.Bytes.String(plugs[0].Data),
	}

	if intent.GasLimit != nil {
		gasLimitStr := hexutil.EncodeUint64(*intent.GasLimit)
		transaction.Gas = &gasLimitStr
	}

	var run *models.Run
	if simulate, ok := intent.Options["simulate"].(bool); ok && simulate {
		run, err = simulation.SimulateRaw(transaction, nil)
		if err != nil {
			return nil, err
		}
	}
	run.IntentId = intent.Id

	fmt.Printf("run: %v\n", run)
	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save simulation run: %v", err)
	}

	return &Solution{
		Transactions: plugs,
		Intent:       &intent,
		Run:          run,
		Transaction:  &transaction,
	}, nil
}

func (s *Solver) Solve(intent models.Intent) (solution *Solution, err error) {
	if isEOA, ok := intent.Options["isEOA"].(bool); ok && isEOA {
		return s.SolveEOA(intent)
	}

	livePlugs, err := s.GetLivePlugs(intent)
	if err != nil {
		return nil, err
	}

	transaction, err := s.BuildPlugTransaction(intent, *livePlugs)
	if err != nil {
		return nil, err
	}

	var run *models.Run
	if simulate, ok := intent.Options["simulate"].(bool); ok && simulate {
		run, err = simulation.Simulate(*transaction)
		if err != nil {
			return nil, err
		}
	}
	run.IntentId = intent.Id

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save simulation run: %v", err)
	}

	return &Solution{
		Transactions: livePlugs.Plugs.Plugs, // Transactions in the `livePlug`.
		LivePlugs:    livePlugs,             // The `livePlug` included in the bundle.
		Intent:       &intent,               // Intent the solver built from.
		Run:          run,                   // Simulation results of solver run.
		Transaction:  transaction,           // Transaction the solver runs.
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

		solution, err := s.Solve(intent)
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
