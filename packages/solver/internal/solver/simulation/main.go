package simulation

import (
	"context"
	"fmt"
	"math/big"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Simulator struct {
	provider *ethclient.Client
}

func NewSimulator(chainId int) (*Simulator, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	return &Simulator{
		provider: provider,
	}, nil
}

func (s *Simulator) SimulateTransaction(ctx context.Context, plugs *signature.LivePlugs) (*SimulationResult, error) {
	// Get latest block for simulation context
	block, err := s.provider.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	solverAddr := common.HexToAddress(hexutil.Encode(plugs.Plugs.Solver))

	// Create call message
	msg := ethereum.CallMsg{
		From:      plugs.Plugs.Socket,
		To:        &solverAddr,
		Gas:       block.GasLimit(),
		GasPrice:  block.BaseFee(),
		GasFeeCap: block.BaseFee(),
		GasTipCap: big.NewInt(0),
		Value:     big.NewInt(0),
		Data:      plugs.Signature,
	}

	// Simulate transaction
	result, err := s.provider.CallContract(ctx, msg, block.Number())
	if err != nil {
		// Parse revert reason if available
		revertErr, ok := err.(*revertError)
		if !ok {
			return &SimulationResult{
				Success: false,
				Error: &SimulationError{
					Message: err.Error(),
				},
			}, nil
		}

		return &SimulationResult{
			Success: false,
			Error: &SimulationError{
				Message:      "Transaction reverted",
				RevertData:   revertErr.Data(),
				RevertReason: hexutil.Encode(revertErr.Data()),
			},
		}, nil
	}

	// Estimate gas (we do this after successful simulation)
	gasEstimate, err := s.provider.EstimateGas(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	return &SimulationResult{
		Success:    true,
		GasUsed:    gasEstimate,
		ReturnData: result,
		// Note: State changes tracking would require additional RPC calls to track storage changes
		StateChanges: []StateChange{},
	}, nil
}

// Helper type for parsing revert errors
type revertError struct {
	error
	data []byte
}

func (e *revertError) Data() []byte {
	return e.data
}

func GetSimulation(chainId int, executionId string, plugs *signature.LivePlugs) (solver.SimulationRequest, error) {
	simulator, err := NewSimulator(chainId)
	if err != nil {
		return solver.SimulationRequest{
			Id:     executionId,
			Status: "error",
			Error:  fmt.Sprintf("Failed to create simulator: %v", err),
		}, nil
	}

	result, err := simulator.SimulateTransaction(context.Background(), plugs)
	if err != nil {
		return solver.SimulationRequest{
			Id:     executionId,
			Status: "error",
			Error:  fmt.Sprintf("Simulation failed: %v", err),
		}, nil
	}

	if !result.Success {
		return solver.SimulationRequest{
			Id:     executionId,
			Status: "error",
			Error:  result.Error.Message,
		}, nil
	}

	return solver.SimulationRequest{
		Id:          executionId,
		Status:      "success",
		GasEstimate: int(result.GasUsed),
	}, nil
}
