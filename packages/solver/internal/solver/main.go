package solver

import (
	"encoding/json"
	"fmt"
	"maps"
	"math/big"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3"
	"solver/internal/actions/assert"
	"solver/internal/actions/basepaint"
	"solver/internal/actions/boolean"
	dbactions "solver/internal/actions/database"
	"solver/internal/actions/euler"
	"solver/internal/actions/math"
	"solver/internal/actions/morpho"
	"solver/internal/actions/nouns"
	"solver/internal/actions/plug"
	"solver/internal/actions/yearn_v3"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Solver struct {
	Protocols map[string]actions.Protocol
	IsKilled  bool
}

func New() *Solver {
	return &Solver{
		Protocols: map[string]actions.Protocol{
			actions.AaveV3:    aave_v3.New(),
			actions.Assert:    assert.New(),
			actions.BasePaint: basepaint.New(),
			actions.Boolean:   boolean.New(),
			actions.Euler:     euler.New(),
			actions.Math:      math.New(),
			actions.Morpho:    morpho.New(),
			actions.Nouns:     nouns.New(),
			actions.Plug:      plug.New(),
			actions.Database:  dbactions.New(),
			actions.YearnV3:   yearn_v3.New(),
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
	if !exists || handler.Actions[inputs.Action].Handler == nil {
		return nil, fmt.Errorf("unsupported protocol: %s", inputs.Protocol)
	}

	params := actions.HandlerParams{}
	params, err := params.New(chainId, from)
	if err != nil {
		return nil, err
	}

	transactions, err := handler.Actions[inputs.Action].Handler(rawInputs, params)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		if transactions[i].Value == nil {
			transactions[i].Value = big.NewInt(0)
		}
	}

	return transactions, nil
}

func (s *Solver) GetPlugsArray(head []signature.Plug, inputs []byte, chainId uint64, from common.Address) (plugs []signature.Plug, error error) {
	plugs, err := s.GetTransaction(inputs, chainId, from)
	if err != nil {
		return nil, err
	}

	return append(head, plugs...), nil
}

func (s *Solver) GetPlugs(intent *models.Intent) ([]signature.Plug, error) {
	var plugs []signature.Plug
	for _, input := range intent.Inputs {
		inputsMap := map[string]any{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		maps.Copy(inputsMap, input)
		inputs, err := json.Marshal(inputsMap)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}

		var exclusive bool
		plugs, err = s.GetPlugsArray(plugs, inputs, intent.ChainId, common.HexToAddress(intent.From))
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}
		if exclusive {
			break
		}
	}

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

func (s *Solver) RebuildSolutionFromModels(intent *models.Intent) (*Solution, error) {
	var livePlugs signature.LivePlugs
	if err := database.DB.Where("intent_id = ?", intent.Id).
		Order("created_at DESC").
		First(&livePlugs).Error; err != nil {
		return nil, fmt.Errorf("failed to find live plug: %v", err)
	}

	var run models.Run
	if err := database.DB.Where("intent_id = ? AND live_plug_id = ?", intent.Id, livePlugs.Id).
		Order("created_at DESC").
		First(&run).Error; err != nil {
		return nil, fmt.Errorf("failed to find run: %v", err)
	}

	var plugs []signature.Plug
	if err := database.DB.Where("bundle_id = ?", livePlugs.Id).
		Find(&plugs).Error; err != nil {
		return nil, fmt.Errorf("failed to find plugs: %v", err)
	}
	livePlugs.Plugs.Plugs = plugs

	return &Solution{
		LivePlugs: &livePlugs,
		Run:       &run,
	}, nil
}

func (s *Solver) SolveEOA(intent *models.Intent, simulate bool) (solution *Solution, err error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	identifier := []byte("plug")
	data := append(plugs[0].Data, identifier...)

	var run *models.Run
	if simulate && intent.Options["simulate"].(bool) {
		simTx := &signature.Transaction{
			From:  common.HexToAddress(intent.From),
			To:    plugs[0].To,
			Data:  data,
			Value: plugs[0].Value,
		}

		run, err = simulation.SimulateEOATx(simTx, nil, intent.ChainId)
		if err != nil {
			return nil, err
		}

		if run != nil {
			run.IntentId = intent.Id

			if err := database.DB.Create(run).Error; err != nil {
				return nil, fmt.Errorf("failed to save simulation run: %v", err)
			}
		}
	} else {
		run = &models.Run{
			IntentId: intent.Id,
			Status:   "skipped",
		}

		if err := database.DB.Create(run).Error; err != nil {
			return nil, fmt.Errorf("failed to save pending run: %v", err)
		}
	}

	transactions := make([]*signature.MinimalPlug, len(plugs))
	for i, plug := range plugs {
		transactions[i] = plug.Minify()
	}

	return &Solution{
		Transactions: transactions,
		Run:          run,
	}, nil
}

func (s *Solver) SolveAndSimulateSocket(intent *models.Intent) (solution *Solution, err error) {
	return s.Solve(intent, true, false)
}
func (s *Solver) SolveSocket(intent *models.Intent, simulate bool) (solution *Solution, err error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	livePlugs, err := s.GetLivePlugs(plugs, intent.ChainId, intent.From)
	if err != nil {
		return nil, err
	}

	livePlugs.IntentId = intent.Id
	livePlugs.ChainId = intent.ChainId
	livePlugs.From = intent.From

	if err := database.DB.Create(livePlugs).Error; err != nil {
		return nil, fmt.Errorf("failed to save live plugs: %v", err)
	}

	callData, err := livePlugs.GetCallData()
	if err != nil {
		return nil, err
	}

	run := &models.Run{
		IntentId:    intent.Id,
		LivePlugsId: livePlugs.Id,
		Status:      "pending",
	}
	shouldSimulate := simulate && intent.Options["simulate"] != nil && intent.Options["simulate"].(bool)
	if shouldSimulate {
		simLivePlugs := &signature.LivePlugs{
			Id:        livePlugs.Id,
			ChainId:   livePlugs.ChainId,
			From:      livePlugs.From,
			IntentId:  livePlugs.IntentId,
			Data:      hexutil.Bytes(callData).String(),
			Plugs:     livePlugs.Plugs,
			Signature: livePlugs.Signature,
		}

		run, err = simulation.SimulateLivePlugs(simLivePlugs)
		if err != nil {
			return nil, err
		}
	}

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save pending run: %v", err)
	}

	result := &Solution{
		Run: run,
	}

	if livePlugs != nil {
		routerAddress := livePlugs.GetRouterAddress()
		routerPlug := &signature.MinimalPlug{
			To:    routerAddress,
			Data:  callData,
			Value: big.NewInt(0),
		}
		result.Transactions = []*signature.MinimalPlug{routerPlug}
	}

	return result, nil
}

func (s *Solver) Solve(intent *models.Intent, simulate bool, live bool) (solution *Solution, err error) {
	var result *Solution
	var solveErr error
	var plugs []signature.Plug

	plugs, err = s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	if isEOA, ok := intent.Options["isEOA"].(bool); ok && isEOA {
		result, solveErr = s.SolveEOA(intent, simulate)
		if solveErr != nil {
			return nil, solveErr
		}

		transactions := make([]*signature.MinimalPlug, len(plugs))
		for i, plug := range plugs {
			transactions[i] = plug.Minify()
		}
		result.Transactions = transactions
	} else {
		result, solveErr = s.SolveSocket(intent, simulate)
		if solveErr != nil {
			return nil, solveErr
		}
	}

	if !live {
		result.LivePlugs = nil
	}

	return result, nil
}
