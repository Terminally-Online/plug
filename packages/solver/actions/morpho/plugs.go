package morpho

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/erc_20"
	"solver/bindings/morpho_distributor"
	"solver/bindings/morpho_router"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleSupplyCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Target string   `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal supply collateral inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(utils.Mainnet.References["morpho"]["router"]),
		inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	supplyCalldata, err := morphoAbi.Pack(
		"supply",
		market.Params,
		inputs.Amount,
		common.HexToAddress(params.From),
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack supply calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   inputs.Token,
			Data: "0x" + common.Bytes2Hex(approveCalldata),
		},
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(supplyCalldata),
		},
	}, nil
}

func HandleWithdrawCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Target string   `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	withdrawCalldata, err := morphoAbi.Pack(
		"withdraw",
		market.Params,
		inputs.Amount,
		big.NewInt(0),
		common.HexToAddress(params.From),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(withdrawCalldata),
		},
	}, nil
}

func HandleWithdrawMaxCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Target string   `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(utils.Mainnet.References["morpho"]["router"]),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho router: %w", err)
	}
	position, err := morpho.Position(
		utils.BuildCallOpts(params.From, big.NewInt(0)),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get position: %w", err)
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	withdrawCalldata, err := morphoAbi.Pack(
		"withdraw",
		market.Params,
		big.NewInt(0),
		position.SupplyShares,
		common.HexToAddress(params.From),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(withdrawCalldata),
		},
	}, nil
}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Target string   `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	borrowCalldata, err := morphoAbi.Pack(
		"borrow",
		market.Params,
		inputs.Amount,
		big.NewInt(0),
		common.HexToAddress(params.From),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack borrow calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(borrowCalldata),
		},
	}, nil
}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Target string   `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(
		utils.Mainnet.References["morpho"]["router"]),
		inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	repayCalldata, err := morphoAbi.Pack(
		"repay",
		market.Params,
		inputs.Amount,
		big.NewInt(0),
		common.HexToAddress(params.From),
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack repay calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   inputs.Token,
			Data: "0x" + common.Bytes2Hex(approveCalldata),
		},
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(repayCalldata),
		},
	}, nil
}

func HandleRepayMax(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay max inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(utils.Mainnet.References["morpho"]["router"]),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho router: %w", err)
	}
	position, err := morpho.Position(
		utils.BuildCallOpts(params.From, big.NewInt(0)),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get position: %w", err)
	}

	borrowAssets := mulDivUp(
		position.BorrowShares,
		new(big.Int).Add(market.State.BorrowAssets, VirtualAssets),
		new(big.Int).Add(market.State.BorrowShares, VirtualShares),
	)
	erc20, err := erc_20.NewErc20(common.HexToAddress(market.LoanAsset.Address), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get erc20: %w", err)
	}
	balance, err := erc20.BalanceOf(nil, common.HexToAddress(params.From))
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	// NOTE: Not sure if we want to revert or use the entire balance they have in this case.
	if balance.Cmp(borrowAssets) < 0 {
		return nil, fmt.Errorf("balance %s is less than borrow assets %s", balance, borrowAssets)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(
		utils.Mainnet.References["morpho"]["router"]),
		borrowAssets,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	// NOTE: Shares is set to zero because assets informs the amount of shares to burn while we
	//       maintain the precise control over the amount of assets to approve.
	repayCalldata, err := morphoAbi.Pack(
		"repay",
		market.Params,
		borrowAssets,
		big.NewInt(0),
		common.HexToAddress(params.From),
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack repay calldata: %w", err)
	}

	return []*types.Transaction{
		{
			To:   inputs.Token,
			Data: "0x" + common.Bytes2Hex(approveCalldata),
		},
		{
			To:   utils.Mainnet.References["morpho"]["router"],
			Data: "0x" + common.Bytes2Hex(repayCalldata),
		},
	}, nil
}

func HandleClaimRewards(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	distributions, err := GetDistributions(params.From, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get distributions: %w", err)
	}
	if len(distributions) == 0 {
		return nil, fmt.Errorf("no active distributions found for address: %s", params.From)
	}

	distributorAbi, err := morpho_distributor.MorphoDistributorMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho distributor abi: %w", err)
	}

	claims := []*types.Transaction{}
	for _, distribution := range distributions {
		claimable, ok := new(big.Int).SetString(distribution.Claimable, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse claimable amount: %s", distribution.Claimable)
		}
		proof := make([][32]byte, len(distribution.Proof))
		for i, proofItem := range distribution.Proof {
			proof[i] = common.HexToHash(proofItem)
		}
		claimCalldata, err := distributorAbi.Pack(
			"claim",
			common.HexToAddress(params.From),
			common.HexToAddress(distribution.Asset.Address),
			claimable,
			proof,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to pack claim calldata: %w", err)
		}
		claims = append(claims, &types.Transaction{
			To:   distribution.Distributor.Address,
			Data: "0x" + common.Bytes2Hex(claimCalldata),
		})
	}

	return claims, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Target    string     `json:"target"`
		Operator  int8       `json:"operator"`
		Threshold *big.Float `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health factor constraint inputs: %w", err)
	}

	market, err := GetMarket(inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(utils.Mainnet.References["morpho"]["router"]),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho router: %w", err)
	}
	position, err := morpho.Position(
		utils.BuildCallOpts(params.From, big.NewInt(0)),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get position: %w", err)
	}

	borrowAssets := mulDivUp(
		position.BorrowShares,
		new(big.Int).Add(market.State.BorrowAssets, VirtualAssets),
		new(big.Int).Add(market.State.BorrowShares, VirtualShares),
	)
	maxBorrow := wMulDown(
		mulDivDown(position.Collateral, market.State.Price, OraclePriceScale),
		market.LLTV,
	)
	healthFactor := new(big.Float).SetInt(wDivDown(maxBorrow, borrowAssets))

	switch inputs.Operator {
	case -1:
		if healthFactor.Cmp(inputs.Threshold) >= 0 {
			return nil, fmt.Errorf("health factor %.2f is not less than threshold %.2f", healthFactor, inputs.Threshold)
		}
	case 1:
		if healthFactor.Cmp(inputs.Threshold) <= 0 {
			return nil, fmt.Errorf("health factor %.2f is not greater than threshold %.2f", healthFactor, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Direction int     `json:"direction"` // -1 for borrow, 1 for deposit
		Address   string  `json:"address"`   // Underlying market or vault
		Operator  int     `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold float64 `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs: %w", err)
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch markets: %w", err)
	}

	var market *Market
	for _, m := range markets {
		if m.UniqueKey == inputs.Address {
			market = &m
			break
		}
	}

	if market == nil {
		return nil, fmt.Errorf("market not found for address: %s", inputs.Address)
	}

	var currentRate float64
	switch inputs.Direction {
	case -1:
		currentRate = market.State.BorrowApy * 100 // Convert to percentage
	case 1:
		currentRate = market.State.SupplyApy * 100 // Convert to percentage
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", inputs.Direction)
	}

	switch inputs.Operator {
	case -1:
		if currentRate >= inputs.Threshold {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", currentRate, inputs.Threshold)
		}
	case 1:
		if currentRate <= inputs.Threshold {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", currentRate, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
