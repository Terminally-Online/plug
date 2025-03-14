package morpho

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/morpho_distributor"
	"solver/bindings/morpho_router"
	"solver/bindings/morpho_vault"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func HandleEarn(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal earn inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert earn amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(inputs.Vault),
		amount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack approve calldata: %w", err)
	}

	vaultAbi, err := morpho_vault.MorphoVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho vault abi: %w", err)
	}
	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		params.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack deposit calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(inputs.Vault),
		Data: depositCalldata,
	}}, nil
}

func HandleSupplyCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal supply collateral inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert supply collateral amount to uint: %w", err)
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	supplyCollateralCalldata, err := morphoAbi.Pack(
		"supplyCollateral",
		market.Params,
		amount,
		params.From,
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack supply calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: supplyCollateralCalldata,
	}}, nil
}

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert withdraw collateral amount to uint: %w", err)
	}

	if len(inputs.Target) == 42 {
		_, err := GetVault(inputs.Target, params.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vault: %w", err)
		}

		vaultAbi, err := morpho_vault.MorphoVaultMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get vault abi: %w", err)
		}
		withdrawCalldata, err := vaultAbi.Pack(
			"withdraw",
			amount,
			params.From,
			params.From,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
		}

		return []signature.Plug{{
			To:   common.HexToAddress(inputs.Target),
			Data: withdrawCalldata,
		}}, nil
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
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
		amount,
		big.NewInt(0),
		params.From,
		params.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: withdrawCalldata,
	}}, nil
}

func HandleWithdrawAll(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	if len(inputs.Target) == 42 {
		_, err := GetVault(inputs.Target, params.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vault: %w", err)
		}

		erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get erc20 abi: %w", err)
		}
		balance, err := erc20Abi.Pack("balanceOf", params.From)
		if err != nil {
			return nil, fmt.Errorf("failed to pack balance of calldata: %w", err)
		}

		vaultAbi, err := morpho_vault.MorphoVaultMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get vault abi: %w", err)
		}
		withdrawCalldata, err := vaultAbi.Pack(
			"redeem",
			balance,
			params.From,
			params.From,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
		}

		return []signature.Plug{{
			To:   common.HexToAddress(inputs.Target),
			Data: withdrawCalldata,
		}}, nil
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		params.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho router: %w", err)
	}
	position, err := morpho.Position(
		params.Client.ReadOptions(params.From),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		params.From,
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
		params.From,
		params.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: withdrawCalldata,
	}}, nil
}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw collateral inputs: %w", err)
	}

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
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
		amount,
		big.NewInt(0),
		params.From,
		params.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack borrow calldata: %w", err)
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: borrowCalldata,
	}}, nil
}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repay amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(
		references.Networks[params.ChainId].References["morpho"]["router"]),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
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
		amount,
		big.NewInt(0),
		params.From,
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack repay calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: repayCalldata,
	}}, nil
}

func HandleRepayAll(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay max inputs: %w", err)
	}

	token, _, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		params.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho router: %w", err)
	}
	position, err := morpho.Position(
		params.Client.ReadOptions(params.From),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		params.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get position: %w", err)
	}

	borrowAssets := mulDivUp(
		position.BorrowShares,
		new(big.Int).Add(market.State.BorrowAssets, VirtualAssets),
		new(big.Int).Add(market.State.BorrowShares, VirtualShares),
	)

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(
		references.Networks[params.ChainId].References["morpho"]["router"]),
		borrowAssets,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
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
		params.From,
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack repay calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		Data: repayCalldata,
	}}, nil
}

func HandleClaimRewards(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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

	claims := []signature.Plug{}
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
			params.From,
			common.HexToAddress(distribution.Asset.Address),
			claimable,
			proof,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to pack claim calldata: %w", err)
		}
		claims = append(claims, signature.Plug{
			To:   common.HexToAddress(distribution.Distributor.Address),
			Data: claimCalldata,
		})
	}

	return claims, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Target    string `json:"target"`
		Operator  int8   `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health factor constraint inputs: %w", err)
	}

	// Morpho has health factor formatted as a WAD (18 decimals)
	threshold, err := utils.StringToUint(inputs.Threshold, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert health factor threshold to uint: %w", err)
	}

	market, err := GetMarket(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(references.Networks[params.ChainId].References["morpho"]["router"]),
		params.Client,
	)
	if err != nil {
		return nil, err
	}
	position, err := morpho.Position(
		params.Client.ReadOptions(params.From),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		params.From,
	)
	if err != nil {
		return nil, err
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
	healthFactor := wDivDown(maxBorrow, borrowAssets)

	switch inputs.Operator {
	case -1:
		if healthFactor.Cmp(threshold) >= 0 {
			return nil, fmt.Errorf("health factor %.2f is not less than threshold %.2f", healthFactor, threshold)
		}
	case 1:
		if healthFactor.Cmp(threshold) <= 0 {
			return nil, fmt.Errorf("health factor %.2f is not greater than threshold %.2f", healthFactor, threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Direction int    `json:"direction"` // -1 for borrow, 1 for deposit
		Target    string `json:"target"`    // Underlying market or vault
		Operator  int    `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold string `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold float: %w", err)
	}

	var currentRate float64
	if len(inputs.Target) == 42 {
		vault, err := GetVault(inputs.Target, params.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch vault: %w", err)
		}

		if inputs.Direction != 1 {
			return nil, fmt.Errorf("vaults only support deposit direction (1), got %d", inputs.Direction)
		}

		currentRate = vault.DailyApys.NetApy * 100
	} else {
		market, err := GetMarket(inputs.Target, params.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch market: %w", err)
		}

		switch inputs.Direction {
		case -1:
			currentRate = market.DailyApys.BorrowApy * 100
		case 1:
			currentRate = market.DailyApys.SupplyApy * 100
		default:
			return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", inputs.Direction)
		}
	}

	switch inputs.Operator {
	case -1:
		if currentRate >= thresholdFloat {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", currentRate, thresholdFloat)
		}
	case 1:
		if currentRate <= thresholdFloat {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", currentRate, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
