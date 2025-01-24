package solver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3"
	"solver/internal/actions/ens"
	"solver/internal/actions/morpho"
	"solver/internal/actions/nouns"
	"solver/internal/actions/plug"
	"solver/internal/actions/yearn_v3"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Solver struct {
	IsKilled  bool
	protocols map[string]actions.BaseProtocolHandler
}

func New() *Solver {
	return &Solver{
		IsKilled: false,
		protocols: map[string]actions.BaseProtocolHandler{
			actions.ProtocolPlug:    plug.New(),
			actions.ProtocolAaveV3:  aave_v3.New(),
			actions.ProtocolYearnV3: yearn_v3.New(),
			actions.ProtocolENS:     ens.New(),
			actions.ProtocolNouns:   nouns.New(),
			actions.ProtocolMorpho:  morpho.New(),
		},
	}
}

func (s *Solver) GetProtocols() map[string]actions.BaseProtocolHandler {
	return s.protocols
}

func (s *Solver) GetProtocolHandler(protocol string) (actions.BaseProtocolHandler, bool) {
	handler, exists := s.protocols[protocol]
	return handler, exists
}

func (s *Solver) GetSupportedProtocols(action string) []string {
	supported := make([]string, 0)
	for protocol, handler := range s.protocols {
		for _, supportedAction := range handler.GetActions() {
			if supportedAction == action {
				supported = append(supported, protocol)
				break
			}
		}
	}
	return supported
}

func (s *Solver) GetExecutions() (ExecutionsRequest, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulateNext")
	response, err := utils.MakeHTTPRequest(
		url,
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		nil,
		ExecutionsRequest{},
	)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (s *Solver) GetTransaction(rawInputs json.RawMessage, chainId int, from string) ([]signature.Plug, error) {
	var inputs struct {
		Protocol string `json:"protocol"`
		Action   string `json:"action"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	handler, exists := s.protocols[inputs.Protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", inputs.Protocol)
	}

	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}

	params := actions.HandlerParams{
		Provider: provider,
		ChainId:  chainId,
		From:     from,
	}

	return handler.GetTransaction(inputs.Action, rawInputs, params)
}

func (s *Solver) GetTransactions(execution ExecutionRequest) ([]signature.Plug, error) {
	var breakOuter bool
	transactionsBatch := make([]signature.Plug, 0)
	errors := make([]error, len(execution.Inputs))
	for i, input := range execution.Inputs {
		inputMap := map[string]interface{}{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		for k, v := range input {
			inputMap[k] = v
		}

		inputsJson, err := json.Marshal(inputMap)
		if err != nil {
			errors[i] = err
			continue
		}

		transactions, err := s.GetTransaction(inputsJson, execution.ChainId, execution.From)
		if err != nil {
			errors[i] = err
			continue
		}

		// NOTE: Some plug actions have exclusive transactions that need to be run alone
		//       before the rest of the Plug can run. For this, we will just break out
		//       of the loop and execute any solo transactions that are needed for
		//       the rest of the batch to run in sequence.
		for _, transaction := range transactions {
			if transaction.Exclusive {
				// NOTE: Set the field to false to avoid tarnishing the response shape.
				transaction.Exclusive = false
				transactionsBatch = []signature.Plug{transaction}
				breakOuter = true
				break
			}
		}

		if breakOuter {
			break
		}

		transactionsBatch = append(transactionsBatch, transactions...)
	}

	// If there were any errors we will return a failure.
	for _, err := range errors {
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}
	}

	// If there was no transaction to execute we will return a warning because
	// we will be halting the simulation of this workflow.
	if len(transactionsBatch) == 0 {
		return nil, utils.ErrBuild("no transactions to execute")
	}

	return transactionsBatch, nil
}

func (s *Solver) GetPlugs(chainId int, from string, transactions []signature.Plug) (*signature.LivePlugs, error) {
	// NOTE: This sets the expiration of a Solver provided order to five minutes from now so that our Solver
	//       cannot sign a message, someone else get a hold if it and execute way in the future or us
	//       end up having the case where things are Plugs are not properly executed because they are being
	//       executed 10k blocks late after it was held from execution.
	expiration := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(300))
	solver, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(expiration, common.HexToAddress(os.Getenv("SOLVER_ADDRESS")))
	if err != nil {
		return nil, utils.ErrBuild("failed to pack solver: " + err.Error())
	}

	salt, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		big.NewInt(time.Now().Unix()),
		common.HexToAddress(from),
		common.HexToAddress(os.Getenv("ONE_CLICKER_ADDRESS")),
		common.HexToAddress(os.Getenv("IMPLEMENTATION_ADDRESS")),
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack salt: " + err.Error())
	}
	plugs := signature.Plugs{
		Socket: common.HexToAddress(from),
		Plugs:  transactions,
		Solver: solver,
		Salt:   salt,
	}
	plugsSignature, err := signature.GetSignature(
		big.NewInt(int64(chainId)),
		common.HexToAddress(from),
		plugs,
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to sign: " + err.Error())
	}

	return &signature.LivePlugs{
		Plugs:     plugs,
		Signature: plugsSignature,
	}, nil
}

func (s *Solver) GetSimulation(id string, plugs *signature.LivePlugs) (SimulationRequest, error) {
	// NOTE: This is where we will pick back up for simulation integration.
	return SimulationRequest{
		Id:          id,
		Status:      "success",
		GasEstimate: 100000,
	}, nil
}

func (s *Solver) GetRun(transactions []signature.Plug) error {
	// TODO: Run the transactions through the entrypoint with our executor account.
	return nil
}

func (s *Solver) PostSimulations(simulations []SimulationRequest) error {
	response := SimulationsRequest{
		Json: simulations,
	}
	body, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = utils.MakeHTTPRequest(
		fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulated"),
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		bytes.NewReader(body),
		SimulationsResponse{},
	)
	if err != nil {
		return err
	}

	return nil
}
