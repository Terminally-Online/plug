package solver

import (
	"encoding/json"
	"fmt"
	"maps"
	"math/big"
	"os"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3"
	"solver/internal/actions/assert"
	"solver/internal/actions/basepaint"
	"solver/internal/actions/boolean"
	"solver/internal/actions/euler"
	"solver/internal/actions/math"
	"solver/internal/actions/morpho"
	"solver/internal/actions/nouns"
	"solver/internal/actions/plug"
	"solver/internal/actions/yearn_v3"
	"solver/internal/client"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type Solver struct {
	Protocols map[string]actions.Protocol
	IsKilled  bool
}

var (
	SolverInstance = New()
)

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
			actions.YearnV3:   yearn_v3.New(),
		},
		IsKilled: false,
	}
}

func (s *Solver) GetTransaction(plugs []signature.Plug, raw json.RawMessage, chainId uint64, from common.Address, prevAction *actions.ActionDefinitionInterface) ([]signature.Plug, error) {
	var inputs struct {
		Protocol string `json:"protocol"`
		Action   string `json:"action"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	protocol, protocolExists := s.Protocols[inputs.Protocol]
	action, actionExists := protocol.Actions[inputs.Action]
	if !protocolExists || !actionExists {
		return nil, fmt.Errorf("unsupported schema lookup: %s-%s", inputs.Protocol, inputs.Action)
	}

	var lookupPrevAction actions.ActionDefinitionInterface
	if prevAction != nil {
		lookupPrevAction = *prevAction
	}

	lookup, err := actions.NewSchemaLookup[any](chainId, from, nil, &raw, lookupPrevAction)
	if err != nil {
		return nil, err
	}

	handler := action.GetHandler()
	transactions, err := handler(lookup)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		transactions[i].Data = hexutil.Bytes(transactions[i].Data)

		if transactions[i].Value == nil {
			transactions[i].Value = big.NewInt(0)
		}

		if transactions[i].Value.Cmp(big.NewInt(0)) != 0 {
			transactions[i].Selector = signature.CallWithValue
		}
	}

	return transactions, nil
}

func (s *Solver) GetPlugsArray(head []signature.Plug, inputs []byte, chainId uint64, from common.Address, prevAction *actions.ActionDefinitionInterface) (plugs []signature.Plug, error error) {
	plugs, err := s.GetTransaction(plugs, inputs, chainId, from, prevAction)
	if err != nil {
		return nil, err
	}

	return append(head, plugs...), nil
}

func (s *Solver) GetPlugs(intent *models.Intent) ([]signature.Plug, error) {
	var plugs []signature.Plug
	var prevAction *actions.ActionDefinitionInterface

	for _, input := range intent.Inputs {
		protocolKey, protocolKeyOk := input["protocol"].(string)
		actionKey, actionKeyOk := input["action"].(string)
		if !protocolKeyOk || !actionKeyOk {
			return nil, utils.ErrBuild("protocol and action must be defined in input")
		}

		protocol, protocolExists := s.Protocols[protocolKey]
		action, actionExists := protocol.Actions[actionKey]
		if !protocolExists || !actionExists {
			return nil, fmt.Errorf("unsupported schema lookup: %s-%s", protocolKey, actionKey)
		}

		inputsMap := map[string]any{
			"protocol": protocol,
			"action":   action,
		}
		maps.Copy(inputsMap, input)
		inputs, err := json.Marshal(inputsMap)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}

		plugs, err = s.GetPlugsArray(plugs, inputs, intent.ChainId, common.HexToAddress(intent.From), prevAction)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}

		// Store the current action as prevAction for the next iteration
		prevAction = &action
	}

	if len(plugs) == 0 {
		return nil, utils.ErrBuild("no transactions to execute")
	}

	return plugs, nil
}

func GetSignature(chainId *big.Int, socket common.Address, plugs signature.Plugs) (signature.Plugs, []byte, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return signature.Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	plugs.Salt, _ = signature.GetSaltHash(socket)
	plugs.Solver, _ = signature.GetSolverHash()

	wrapped, _ := plugs.Wrap()
	client, _ := client.New(chainId.Uint64())
	digest, err := client.Digest(chainId, socket, wrapped)
	if err != nil {
		return signature.Plugs{}, nil, utils.ErrBuild(err.Error())
	}
	sig, err := crypto.Sign(digest[:], privateKey)
	if err != nil {
		return signature.Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	sig[crypto.RecoveryIDOffset] += 27

	return plugs, sig, nil
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

	plugsSigned, plugsSignature, err := GetSignature(
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

	livePlugs := signature.LivePlugs{
		Plugs:     plugsSigned,
		Signature: plugsSignature,
	}

	// client, _ := client.New(chainId)
	// socketAbi, _ := plug_socket.PlugSocketMetaData.GetAbi()
	// lp, _ := livePlugs.Wrap()

	// TODO: This validation is left here at the time of writing this comment because we MAY have
	//       fixed signatures but it is tempermental a bit. Delete this upon 06/01/2025.

	// { // [Passing] Domain hash validation
	// 	socketContract, _ := plug_socket.NewPlugSocket(fromAddress, client)
	// 	domain, _ := socketContract.Domain(client.ReadOptions(fromAddress))
	//
	// 	domainHash, _ := socketContract.GetEIP712DomainHash(client.ReadOptions(fromAddress), domain)
	// 	log.Printf("Domain hash %s", hexutil.Encode(domainHash[:]))
	// }
	//
	//
	// { // [Failing] Plugs digest validation
	// 	wrappedPlugs, _ := plugsSigned.Wrap()
	// 	digestCalldata, err := socketAbi.Pack("getPlugsDigest", wrappedPlugs)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to build getLivePlugsSigner calldata: %w", err)
	// 	}
	//
	// 	output, err := client.CallContract(context.Background(), ethereum.CallMsg{
	// 		To:   &fromAddress,
	// 		Data: digestCalldata,
	// 	}, nil)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to read digest: %w", err)
	// 	}
	// 	var digest [32]byte
	// 	err = socketAbi.UnpackIntoInterface(&digest, "getPlugsDigest", output)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to unpack digest: %w", err)
	// 	}
	// 	log.Printf("Digest %s", hexutil.Encode(digest[:]))
	// }
	//
	// { // [Passing] Signer validation
	// 	signerCalldata, err := socketAbi.Pack("getLivePlugsSigner", lp)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to build getLivePlugsSigner calldata: %w", err)
	// 	}
	//
	// 	output, err := client.CallContract(context.Background(), ethereum.CallMsg{
	// 		To:   &fromAddress,
	// 		Data: signerCalldata,
	// 	}, nil)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to read signer: %w", err)
	// 	}
	//
	// 	var signer common.Address
	// 	err = socketAbi.UnpackIntoInterface(&signer, "getLivePlugsSigner", output)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to unpack signer: %w", err)
	// 	}
	//
	// 	solverAddress := common.HexToAddress(os.Getenv("SOLVER_ADDRESS"))
	// 	if signer != solverAddress {
	// 		return nil, fmt.Errorf("failed to generate valid signature: %s != %s", solverAddress, signer)
	// 	} else {
	// 		log.Println("SIGNERS ARE MATCHING AND VALID")
	// 	}
	// }

	return &livePlugs, nil
}

func (s *Solver) RebuildSolutionFromModels(intent *models.Intent) (*Solution, error) {
	if err := database.DB.Where("id = ?", intent.Id).
		Order("created_at DESC").
		Preload("Runs", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1)
		}).
		Preload("LivePlugs", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1)
		}).
		First(&intent).Error; err != nil {
		return nil, fmt.Errorf("failed to find live plug: %v", err)
	}

	livePlug := intent.LivePlugs[0]
	run := intent.Runs[0]

	return &Solution{
		LivePlugs:    &livePlug,
		Transactions: livePlug.Plugs.Plugs,
		Run:          &run,
	}, nil
}

func (s *Solver) SolveEOA(intent *models.Intent, simulate bool) (solution *Solution, err error) {
	plugs, err := s.GetPlugs(intent)
	if err != nil {
		return nil, err
	}

	identifier := []byte("plug")
	combinedData := append([]byte(plugs[0].Data), identifier...)

	var run *models.Run

	// Safely handle the simulate option to prevent crashes from type assertions
	shouldSimulate := false
	if simulate {
		if simulateVal, ok := intent.Options["simulate"]; ok && simulateVal != nil {
			if boolVal, ok := simulateVal.(bool); ok {
				shouldSimulate = boolVal
			}
		}
	}

	if shouldSimulate {
		simTx := &signature.Transaction{
			From:  common.HexToAddress(intent.From),
			To:    plugs[0].To,
			Data:  hexutil.Bytes(combinedData),
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

	return &Solution{
		Transactions: plugs,
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

	shouldSimulate := false
	if simulate {
		if simulateVal, ok := intent.Options["simulate"]; ok && simulateVal != nil {
			if boolVal, ok := simulateVal.(bool); ok {
				shouldSimulate = boolVal
			}
		}
	}

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

		if run != nil {
			run.IntentId = intent.Id
			run.LivePlugsId = livePlugs.Id

			if err := database.DB.Create(run).Error; err != nil {
				return nil, fmt.Errorf("failed to save simulation run: %v", err)
			}
		}
	}

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save pending run: %v", err)
	}

	result := &Solution{
		Run:       run,
		LivePlugs: livePlugs,
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

		result.Transactions = plugs
	} else {
		result, solveErr = s.SolveSocket(intent, simulate)
		if solveErr != nil {
			return nil, solveErr
		}
	}

	result.IntentId = intent.Id

	if !live {
		result.LivePlugs = nil
	}

	return result, nil
}
