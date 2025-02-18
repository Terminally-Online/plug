package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"solver/bindings/multicall_primary"
	"solver/bindings/plug_router"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	chainId          uint64
	solverAddress    common.Address
	solverPrivateKey *ecdsa.PrivateKey
	*ethclient.Client
}

func New(chainId uint64) (*Client, error) {
	solverAddress := common.HexToAddress(os.Getenv("SOLVER_ADDRESS"))
	solverPrivateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return nil, utils.ErrBuild(err.Error())
	}

	rpcUrl, err := GetQuicknodeUrl(chainId)
	if err != nil {
		return nil, err
	}
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, utils.ErrEthClient(err.Error())
	}

	return &Client{
		chainId:          chainId,
		solverAddress:    solverAddress,
		solverPrivateKey: solverPrivateKey,
		Client:           ethClient,
	}, nil
}

func (c *Client) ReadOptions(address string) *bind.CallOpts {
	return &bind.CallOpts{
		From:    common.HexToAddress(address),
		Pending: true,
		Context: context.Background(),
	}
}
func (c *Client) SolverReadOptions() *bind.CallOpts {
	return c.ReadOptions(os.Getenv("SOLVER_ADDRESS"))
}

func (c *Client) WriteOptions(address string, value *big.Int) *bind.TransactOpts {
	transactionForwarder := func(_ common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return transaction, nil
	}

	return &bind.TransactOpts{
		From:   common.HexToAddress(address),
		Signer: transactionForwarder,
		NoSend: true,
		Value:  value,
	}
}
func (c *Client) SolverWriteOptions() *bind.TransactOpts {
	return c.WriteOptions(os.Getenv("SOLVER_ADDRESS"), big.NewInt(0))
}

func (c *Client) Plug(livePlugs []signature.LivePlugs) (*ethtypes.Transaction, error) {
	router, err := plug_router.NewPlugRouter(
		common.HexToAddress(references.Networks[c.chainId].References["plug"]["router"]),
		c,
	)
	if err != nil {
		return nil, err
	}

	// NOTE: We are doing manual transformation here instead of doing something that is type
	//       based we want the ability to include extra things like `meta` in the response
	//       body shape which would not be straightforward to do if we use the protocol
	//       exposed types that are generated with the abi.
	// TODO: If there is a better way to do this please implement it. - CHANCE
	lps := make([]plug_router.PlugTypesLibLivePlugs, len(livePlugs))
	for _, livePlug := range livePlugs {
		lps = append(lps, livePlug.Wrap())
	}
	plugged, err := router.Plug0(c.SolverWriteOptions(), lps)
	if err != nil {
		return nil, err
	}

	return plugged, nil
}

func (c *Client) Multicall(calls []MulticallCalldata) ([]interface{}, error) {
	multicallAddress := common.HexToAddress(references.Networks[c.chainId].References["multicall"]["primary"])
	if multicallAddress == (common.Address{}) {
		return nil, fmt.Errorf("multicall not found for chain id: %d", c.chainId)
	}

	multicallCalls := make([]multicall_primary.Multicall3Call, len(calls))
	for i, call := range calls {
		callData, err := call.ABI.Pack(call.Method, call.Args...)
		if err != nil {
			return nil, fmt.Errorf("failed to pack %s call: %w", call.Method, err)
		}

		multicallCalls[i] = multicall_primary.Multicall3Call{
			Target:   call.Target,
			CallData: callData,
		}
	}

	multicallAbi, err := multicall_primary.MulticallPrimaryMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Multicall")
	}

	input, err := multicallAbi.Pack("aggregate", multicallCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall aggregate: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   &multicallAddress,
		Data: input,
	}

	output, err := c.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make multicall: %w", err)
	}

	unpacked, err := multicallAbi.Unpack("aggregate", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack multicall aggregate: %w", err)
	}

	_, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("multicall result 0 was not a big.Int")
	}

	returnData, ok := unpacked[1].([][]byte)
	if !ok {
		return nil, fmt.Errorf("multicall result 1 was not a [][]byte")
	}

	results := make([]interface{}, len(returnData))
	for i, data := range returnData {
		unpacked, err := calls[i].ABI.Unpack(calls[i].Method, data)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack data for call %d: %w", i, err)
		}

		if len(unpacked) == 0 {
			return nil, fmt.Errorf("empty result for call %d", i)
		}

		jsonData, err := json.Marshal(unpacked[0])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data for call %d: %w", i, err)
		}

		result := calls[i].OutputType
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data for call %d: %w", i, err)
		}

		results[i] = result
	}

	return results, nil
}
