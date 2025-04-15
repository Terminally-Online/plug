// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho_router

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Authorization is an auto generated low-level Go binding around an user-defined struct.
type Authorization struct {
	Authorizer   common.Address
	Authorized   common.Address
	IsAuthorized bool
	Nonce        *big.Int
	Deadline     *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// MorphoRouterMetaData contains all meta data concerning the MorphoRouter contract.
var MorphoRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"}],\"name\":\"EnableIrm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"name\":\"EnableLltv\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usedNonce\",\"type\":\"uint256\"}],\"name\":\"IncrementNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repaidAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repaidShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizedAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badDebtAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badDebtShares\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorized\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newIsAuthorized\",\"type\":\"bool\"}],\"name\":\"SetAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"SupplyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"accrueInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"createMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"}],\"name\":\"enableIrm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"name\":\"enableLltv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"slots\",\"type\":\"bytes32[]\"}],\"name\":\"extSloads\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"res\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"idToMarketParams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isIrmEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isLltvEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizedAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaidShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"market\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"totalSupplyAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastUpdate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"position\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyShares\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"borrowShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"collateral\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorized\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsAuthorized\",\"type\":\"bool\"}],\"name\":\"setAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorized\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structAuthorization\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"setAuthorizationWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"supplyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MorphoRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoRouterMetaData.ABI instead.
var MorphoRouterABI = MorphoRouterMetaData.ABI

// MorphoRouter is an auto generated Go binding around an Ethereum contract.
type MorphoRouter struct {
	MorphoRouterCaller     // Read-only binding to the contract
	MorphoRouterTransactor // Write-only binding to the contract
	MorphoRouterFilterer   // Log filterer for contract events
}

// MorphoRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoRouterSession struct {
	Contract     *MorphoRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoRouterCallerSession struct {
	Contract *MorphoRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MorphoRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoRouterTransactorSession struct {
	Contract     *MorphoRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MorphoRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoRouterRaw struct {
	Contract *MorphoRouter // Generic contract binding to access the raw methods on
}

// MorphoRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoRouterCallerRaw struct {
	Contract *MorphoRouterCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoRouterTransactorRaw struct {
	Contract *MorphoRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphoRouter creates a new instance of MorphoRouter, bound to a specific deployed contract.
func NewMorphoRouter(address common.Address, backend bind.ContractBackend) (*MorphoRouter, error) {
	contract, err := bindMorphoRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphoRouter{MorphoRouterCaller: MorphoRouterCaller{contract: contract}, MorphoRouterTransactor: MorphoRouterTransactor{contract: contract}, MorphoRouterFilterer: MorphoRouterFilterer{contract: contract}}, nil
}

// NewMorphoRouterCaller creates a new read-only instance of MorphoRouter, bound to a specific deployed contract.
func NewMorphoRouterCaller(address common.Address, caller bind.ContractCaller) (*MorphoRouterCaller, error) {
	contract, err := bindMorphoRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterCaller{contract: contract}, nil
}

// NewMorphoRouterTransactor creates a new write-only instance of MorphoRouter, bound to a specific deployed contract.
func NewMorphoRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoRouterTransactor, error) {
	contract, err := bindMorphoRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterTransactor{contract: contract}, nil
}

// NewMorphoRouterFilterer creates a new log filterer instance of MorphoRouter, bound to a specific deployed contract.
func NewMorphoRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoRouterFilterer, error) {
	contract, err := bindMorphoRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterFilterer{contract: contract}, nil
}

// bindMorphoRouter binds a generic wrapper to an already deployed contract.
func bindMorphoRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoRouter *MorphoRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoRouter.Contract.MorphoRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoRouter *MorphoRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoRouter.Contract.MorphoRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoRouter *MorphoRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoRouter.Contract.MorphoRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoRouter *MorphoRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoRouter *MorphoRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoRouter *MorphoRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoRouter.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoRouter *MorphoRouterCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoRouter *MorphoRouterSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphoRouter.Contract.DOMAINSEPARATOR(&_MorphoRouter.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoRouter *MorphoRouterCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphoRouter.Contract.DOMAINSEPARATOR(&_MorphoRouter.CallOpts)
}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_MorphoRouter *MorphoRouterCaller) ExtSloads(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "extSloads", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_MorphoRouter *MorphoRouterSession) ExtSloads(slots [][32]byte) ([][32]byte, error) {
	return _MorphoRouter.Contract.ExtSloads(&_MorphoRouter.CallOpts, slots)
}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_MorphoRouter *MorphoRouterCallerSession) ExtSloads(slots [][32]byte) ([][32]byte, error) {
	return _MorphoRouter.Contract.ExtSloads(&_MorphoRouter.CallOpts, slots)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoRouter *MorphoRouterCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoRouter *MorphoRouterSession) FeeRecipient() (common.Address, error) {
	return _MorphoRouter.Contract.FeeRecipient(&_MorphoRouter.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoRouter *MorphoRouterCallerSession) FeeRecipient() (common.Address, error) {
	return _MorphoRouter.Contract.FeeRecipient(&_MorphoRouter.CallOpts)
}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 ) view returns(address loanToken, address collateralToken, address oracle, address irm, uint256 lltv)
func (_MorphoRouter *MorphoRouterCaller) IdToMarketParams(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "idToMarketParams", arg0)

	outstruct := new(struct {
		LoanToken       common.Address
		CollateralToken common.Address
		Oracle          common.Address
		Irm             common.Address
		Lltv            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CollateralToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Oracle = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Irm = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Lltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 ) view returns(address loanToken, address collateralToken, address oracle, address irm, uint256 lltv)
func (_MorphoRouter *MorphoRouterSession) IdToMarketParams(arg0 [32]byte) (struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}, error) {
	return _MorphoRouter.Contract.IdToMarketParams(&_MorphoRouter.CallOpts, arg0)
}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 ) view returns(address loanToken, address collateralToken, address oracle, address irm, uint256 lltv)
func (_MorphoRouter *MorphoRouterCallerSession) IdToMarketParams(arg0 [32]byte) (struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}, error) {
	return _MorphoRouter.Contract.IdToMarketParams(&_MorphoRouter.CallOpts, arg0)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x65e4ad9e.
//
// Solidity: function isAuthorized(address , address ) view returns(bool)
func (_MorphoRouter *MorphoRouterCaller) IsAuthorized(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "isAuthorized", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x65e4ad9e.
//
// Solidity: function isAuthorized(address , address ) view returns(bool)
func (_MorphoRouter *MorphoRouterSession) IsAuthorized(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _MorphoRouter.Contract.IsAuthorized(&_MorphoRouter.CallOpts, arg0, arg1)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x65e4ad9e.
//
// Solidity: function isAuthorized(address , address ) view returns(bool)
func (_MorphoRouter *MorphoRouterCallerSession) IsAuthorized(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _MorphoRouter.Contract.IsAuthorized(&_MorphoRouter.CallOpts, arg0, arg1)
}

// IsIrmEnabled is a free data retrieval call binding the contract method 0xf2b863ce.
//
// Solidity: function isIrmEnabled(address ) view returns(bool)
func (_MorphoRouter *MorphoRouterCaller) IsIrmEnabled(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "isIrmEnabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIrmEnabled is a free data retrieval call binding the contract method 0xf2b863ce.
//
// Solidity: function isIrmEnabled(address ) view returns(bool)
func (_MorphoRouter *MorphoRouterSession) IsIrmEnabled(arg0 common.Address) (bool, error) {
	return _MorphoRouter.Contract.IsIrmEnabled(&_MorphoRouter.CallOpts, arg0)
}

// IsIrmEnabled is a free data retrieval call binding the contract method 0xf2b863ce.
//
// Solidity: function isIrmEnabled(address ) view returns(bool)
func (_MorphoRouter *MorphoRouterCallerSession) IsIrmEnabled(arg0 common.Address) (bool, error) {
	return _MorphoRouter.Contract.IsIrmEnabled(&_MorphoRouter.CallOpts, arg0)
}

// IsLltvEnabled is a free data retrieval call binding the contract method 0xb485f3b8.
//
// Solidity: function isLltvEnabled(uint256 ) view returns(bool)
func (_MorphoRouter *MorphoRouterCaller) IsLltvEnabled(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "isLltvEnabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLltvEnabled is a free data retrieval call binding the contract method 0xb485f3b8.
//
// Solidity: function isLltvEnabled(uint256 ) view returns(bool)
func (_MorphoRouter *MorphoRouterSession) IsLltvEnabled(arg0 *big.Int) (bool, error) {
	return _MorphoRouter.Contract.IsLltvEnabled(&_MorphoRouter.CallOpts, arg0)
}

// IsLltvEnabled is a free data retrieval call binding the contract method 0xb485f3b8.
//
// Solidity: function isLltvEnabled(uint256 ) view returns(bool)
func (_MorphoRouter *MorphoRouterCallerSession) IsLltvEnabled(arg0 *big.Int) (bool, error) {
	return _MorphoRouter.Contract.IsLltvEnabled(&_MorphoRouter.CallOpts, arg0)
}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 ) view returns(uint128 totalSupplyAssets, uint128 totalSupplyShares, uint128 totalBorrowAssets, uint128 totalBorrowShares, uint128 lastUpdate, uint128 fee)
func (_MorphoRouter *MorphoRouterCaller) Market(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TotalSupplyAssets *big.Int
	TotalSupplyShares *big.Int
	TotalBorrowAssets *big.Int
	TotalBorrowShares *big.Int
	LastUpdate        *big.Int
	Fee               *big.Int
}, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "market", arg0)

	outstruct := new(struct {
		TotalSupplyAssets *big.Int
		TotalSupplyShares *big.Int
		TotalBorrowAssets *big.Int
		TotalBorrowShares *big.Int
		LastUpdate        *big.Int
		Fee               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSupplyAssets = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalSupplyShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowAssets = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowShares = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 ) view returns(uint128 totalSupplyAssets, uint128 totalSupplyShares, uint128 totalBorrowAssets, uint128 totalBorrowShares, uint128 lastUpdate, uint128 fee)
func (_MorphoRouter *MorphoRouterSession) Market(arg0 [32]byte) (struct {
	TotalSupplyAssets *big.Int
	TotalSupplyShares *big.Int
	TotalBorrowAssets *big.Int
	TotalBorrowShares *big.Int
	LastUpdate        *big.Int
	Fee               *big.Int
}, error) {
	return _MorphoRouter.Contract.Market(&_MorphoRouter.CallOpts, arg0)
}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 ) view returns(uint128 totalSupplyAssets, uint128 totalSupplyShares, uint128 totalBorrowAssets, uint128 totalBorrowShares, uint128 lastUpdate, uint128 fee)
func (_MorphoRouter *MorphoRouterCallerSession) Market(arg0 [32]byte) (struct {
	TotalSupplyAssets *big.Int
	TotalSupplyShares *big.Int
	TotalBorrowAssets *big.Int
	TotalBorrowShares *big.Int
	LastUpdate        *big.Int
	Fee               *big.Int
}, error) {
	return _MorphoRouter.Contract.Market(&_MorphoRouter.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_MorphoRouter *MorphoRouterCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_MorphoRouter *MorphoRouterSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _MorphoRouter.Contract.Nonce(&_MorphoRouter.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_MorphoRouter *MorphoRouterCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _MorphoRouter.Contract.Nonce(&_MorphoRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoRouter *MorphoRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoRouter *MorphoRouterSession) Owner() (common.Address, error) {
	return _MorphoRouter.Contract.Owner(&_MorphoRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoRouter *MorphoRouterCallerSession) Owner() (common.Address, error) {
	return _MorphoRouter.Contract.Owner(&_MorphoRouter.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 , address ) view returns(uint256 supplyShares, uint128 borrowShares, uint128 collateral)
func (_MorphoRouter *MorphoRouterCaller) Position(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	SupplyShares *big.Int
	BorrowShares *big.Int
	Collateral   *big.Int
}, error) {
	var out []interface{}
	err := _MorphoRouter.contract.Call(opts, &out, "position", arg0, arg1)

	outstruct := new(struct {
		SupplyShares *big.Int
		BorrowShares *big.Int
		Collateral   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupplyShares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BorrowShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Collateral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 , address ) view returns(uint256 supplyShares, uint128 borrowShares, uint128 collateral)
func (_MorphoRouter *MorphoRouterSession) Position(arg0 [32]byte, arg1 common.Address) (struct {
	SupplyShares *big.Int
	BorrowShares *big.Int
	Collateral   *big.Int
}, error) {
	return _MorphoRouter.Contract.Position(&_MorphoRouter.CallOpts, arg0, arg1)
}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 , address ) view returns(uint256 supplyShares, uint128 borrowShares, uint128 collateral)
func (_MorphoRouter *MorphoRouterCallerSession) Position(arg0 [32]byte, arg1 common.Address) (struct {
	SupplyShares *big.Int
	BorrowShares *big.Int
	Collateral   *big.Int
}, error) {
	return _MorphoRouter.Contract.Position(&_MorphoRouter.CallOpts, arg0, arg1)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x151c1ade.
//
// Solidity: function accrueInterest((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterTransactor) AccrueInterest(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "accrueInterest", marketParams)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x151c1ade.
//
// Solidity: function accrueInterest((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterSession) AccrueInterest(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.Contract.AccrueInterest(&_MorphoRouter.TransactOpts, marketParams)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x151c1ade.
//
// Solidity: function accrueInterest((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) AccrueInterest(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.Contract.AccrueInterest(&_MorphoRouter.TransactOpts, marketParams)
}

// Borrow is a paid mutator transaction binding the contract method 0x50d8cd4b.
//
// Solidity: function borrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactor) Borrow(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "borrow", marketParams, assets, shares, onBehalf, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x50d8cd4b.
//
// Solidity: function borrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterSession) Borrow(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Borrow(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x50d8cd4b.
//
// Solidity: function borrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactorSession) Borrow(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Borrow(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, receiver)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x8c1358a2.
//
// Solidity: function createMarket((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterTransactor) CreateMarket(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "createMarket", marketParams)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x8c1358a2.
//
// Solidity: function createMarket((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterSession) CreateMarket(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.Contract.CreateMarket(&_MorphoRouter.TransactOpts, marketParams)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x8c1358a2.
//
// Solidity: function createMarket((address,address,address,address,uint256) marketParams) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) CreateMarket(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoRouter.Contract.CreateMarket(&_MorphoRouter.TransactOpts, marketParams)
}

// EnableIrm is a paid mutator transaction binding the contract method 0x5a64f51e.
//
// Solidity: function enableIrm(address irm) returns()
func (_MorphoRouter *MorphoRouterTransactor) EnableIrm(opts *bind.TransactOpts, irm common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "enableIrm", irm)
}

// EnableIrm is a paid mutator transaction binding the contract method 0x5a64f51e.
//
// Solidity: function enableIrm(address irm) returns()
func (_MorphoRouter *MorphoRouterSession) EnableIrm(irm common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.EnableIrm(&_MorphoRouter.TransactOpts, irm)
}

// EnableIrm is a paid mutator transaction binding the contract method 0x5a64f51e.
//
// Solidity: function enableIrm(address irm) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) EnableIrm(irm common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.EnableIrm(&_MorphoRouter.TransactOpts, irm)
}

// EnableLltv is a paid mutator transaction binding the contract method 0x4d98a93b.
//
// Solidity: function enableLltv(uint256 lltv) returns()
func (_MorphoRouter *MorphoRouterTransactor) EnableLltv(opts *bind.TransactOpts, lltv *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "enableLltv", lltv)
}

// EnableLltv is a paid mutator transaction binding the contract method 0x4d98a93b.
//
// Solidity: function enableLltv(uint256 lltv) returns()
func (_MorphoRouter *MorphoRouterSession) EnableLltv(lltv *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.Contract.EnableLltv(&_MorphoRouter.TransactOpts, lltv)
}

// EnableLltv is a paid mutator transaction binding the contract method 0x4d98a93b.
//
// Solidity: function enableLltv(uint256 lltv) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) EnableLltv(lltv *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.Contract.EnableLltv(&_MorphoRouter.TransactOpts, lltv)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address token, uint256 assets, bytes data) returns()
func (_MorphoRouter *MorphoRouterTransactor) FlashLoan(opts *bind.TransactOpts, token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "flashLoan", token, assets, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address token, uint256 assets, bytes data) returns()
func (_MorphoRouter *MorphoRouterSession) FlashLoan(token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.FlashLoan(&_MorphoRouter.TransactOpts, token, assets, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address token, uint256 assets, bytes data) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) FlashLoan(token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.FlashLoan(&_MorphoRouter.TransactOpts, token, assets, data)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd8eabcb8.
//
// Solidity: function liquidate((address,address,address,address,uint256) marketParams, address borrower, uint256 seizedAssets, uint256 repaidShares, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterTransactor) Liquidate(opts *bind.TransactOpts, marketParams MarketParams, borrower common.Address, seizedAssets *big.Int, repaidShares *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "liquidate", marketParams, borrower, seizedAssets, repaidShares, data)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd8eabcb8.
//
// Solidity: function liquidate((address,address,address,address,uint256) marketParams, address borrower, uint256 seizedAssets, uint256 repaidShares, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterSession) Liquidate(marketParams MarketParams, borrower common.Address, seizedAssets *big.Int, repaidShares *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Liquidate(&_MorphoRouter.TransactOpts, marketParams, borrower, seizedAssets, repaidShares, data)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd8eabcb8.
//
// Solidity: function liquidate((address,address,address,address,uint256) marketParams, address borrower, uint256 seizedAssets, uint256 repaidShares, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterTransactorSession) Liquidate(marketParams MarketParams, borrower common.Address, seizedAssets *big.Int, repaidShares *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Liquidate(&_MorphoRouter.TransactOpts, marketParams, borrower, seizedAssets, repaidShares, data)
}

// Repay is a paid mutator transaction binding the contract method 0x20b76e81.
//
// Solidity: function repay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactor) Repay(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "repay", marketParams, assets, shares, onBehalf, data)
}

// Repay is a paid mutator transaction binding the contract method 0x20b76e81.
//
// Solidity: function repay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterSession) Repay(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Repay(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, data)
}

// Repay is a paid mutator transaction binding the contract method 0x20b76e81.
//
// Solidity: function repay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactorSession) Repay(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Repay(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, data)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address authorized, bool newIsAuthorized) returns()
func (_MorphoRouter *MorphoRouterTransactor) SetAuthorization(opts *bind.TransactOpts, authorized common.Address, newIsAuthorized bool) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "setAuthorization", authorized, newIsAuthorized)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address authorized, bool newIsAuthorized) returns()
func (_MorphoRouter *MorphoRouterSession) SetAuthorization(authorized common.Address, newIsAuthorized bool) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetAuthorization(&_MorphoRouter.TransactOpts, authorized, newIsAuthorized)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address authorized, bool newIsAuthorized) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SetAuthorization(authorized common.Address, newIsAuthorized bool) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetAuthorization(&_MorphoRouter.TransactOpts, authorized, newIsAuthorized)
}

// SetAuthorizationWithSig is a paid mutator transaction binding the contract method 0x8069218f.
//
// Solidity: function setAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature) returns()
func (_MorphoRouter *MorphoRouterTransactor) SetAuthorizationWithSig(opts *bind.TransactOpts, authorization Authorization, signature Signature) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "setAuthorizationWithSig", authorization, signature)
}

// SetAuthorizationWithSig is a paid mutator transaction binding the contract method 0x8069218f.
//
// Solidity: function setAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature) returns()
func (_MorphoRouter *MorphoRouterSession) SetAuthorizationWithSig(authorization Authorization, signature Signature) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetAuthorizationWithSig(&_MorphoRouter.TransactOpts, authorization, signature)
}

// SetAuthorizationWithSig is a paid mutator transaction binding the contract method 0x8069218f.
//
// Solidity: function setAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SetAuthorizationWithSig(authorization Authorization, signature Signature) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetAuthorizationWithSig(&_MorphoRouter.TransactOpts, authorization, signature)
}

// SetFee is a paid mutator transaction binding the contract method 0x2b4f013c.
//
// Solidity: function setFee((address,address,address,address,uint256) marketParams, uint256 newFee) returns()
func (_MorphoRouter *MorphoRouterTransactor) SetFee(opts *bind.TransactOpts, marketParams MarketParams, newFee *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "setFee", marketParams, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x2b4f013c.
//
// Solidity: function setFee((address,address,address,address,uint256) marketParams, uint256 newFee) returns()
func (_MorphoRouter *MorphoRouterSession) SetFee(marketParams MarketParams, newFee *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetFee(&_MorphoRouter.TransactOpts, marketParams, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x2b4f013c.
//
// Solidity: function setFee((address,address,address,address,uint256) marketParams, uint256 newFee) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SetFee(marketParams MarketParams, newFee *big.Int) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetFee(&_MorphoRouter.TransactOpts, marketParams, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoRouter *MorphoRouterTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoRouter *MorphoRouterSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetFeeRecipient(&_MorphoRouter.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetFeeRecipient(&_MorphoRouter.TransactOpts, newFeeRecipient)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoRouter *MorphoRouterTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoRouter *MorphoRouterSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetOwner(&_MorphoRouter.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SetOwner(&_MorphoRouter.TransactOpts, newOwner)
}

// Supply is a paid mutator transaction binding the contract method 0xa99aad89.
//
// Solidity: function supply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterTransactor) Supply(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "supply", marketParams, assets, shares, onBehalf, data)
}

// Supply is a paid mutator transaction binding the contract method 0xa99aad89.
//
// Solidity: function supply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterSession) Supply(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Supply(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, data)
}

// Supply is a paid mutator transaction binding the contract method 0xa99aad89.
//
// Solidity: function supply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, bytes data) returns(uint256, uint256)
func (_MorphoRouter *MorphoRouterTransactorSession) Supply(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Supply(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, data)
}

// SupplyCollateral is a paid mutator transaction binding the contract method 0x238d6579.
//
// Solidity: function supplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) returns()
func (_MorphoRouter *MorphoRouterTransactor) SupplyCollateral(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "supplyCollateral", marketParams, assets, onBehalf, data)
}

// SupplyCollateral is a paid mutator transaction binding the contract method 0x238d6579.
//
// Solidity: function supplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) returns()
func (_MorphoRouter *MorphoRouterSession) SupplyCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SupplyCollateral(&_MorphoRouter.TransactOpts, marketParams, assets, onBehalf, data)
}

// SupplyCollateral is a paid mutator transaction binding the contract method 0x238d6579.
//
// Solidity: function supplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) SupplyCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoRouter.Contract.SupplyCollateral(&_MorphoRouter.TransactOpts, marketParams, assets, onBehalf, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c2bea49.
//
// Solidity: function withdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactor) Withdraw(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "withdraw", marketParams, assets, shares, onBehalf, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c2bea49.
//
// Solidity: function withdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterSession) Withdraw(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Withdraw(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c2bea49.
//
// Solidity: function withdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, address onBehalf, address receiver) returns(uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterTransactorSession) Withdraw(marketParams MarketParams, assets *big.Int, shares *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.Withdraw(&_MorphoRouter.TransactOpts, marketParams, assets, shares, onBehalf, receiver)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x8720316d.
//
// Solidity: function withdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, address receiver) returns()
func (_MorphoRouter *MorphoRouterTransactor) WithdrawCollateral(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.contract.Transact(opts, "withdrawCollateral", marketParams, assets, onBehalf, receiver)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x8720316d.
//
// Solidity: function withdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, address receiver) returns()
func (_MorphoRouter *MorphoRouterSession) WithdrawCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.WithdrawCollateral(&_MorphoRouter.TransactOpts, marketParams, assets, onBehalf, receiver)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x8720316d.
//
// Solidity: function withdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, address receiver) returns()
func (_MorphoRouter *MorphoRouterTransactorSession) WithdrawCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, receiver common.Address) (*types.Transaction, error) {
	return _MorphoRouter.Contract.WithdrawCollateral(&_MorphoRouter.TransactOpts, marketParams, assets, onBehalf, receiver)
}

// MorphoRouterAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the MorphoRouter contract.
type MorphoRouterAccrueInterestIterator struct {
	Event *MorphoRouterAccrueInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterAccrueInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterAccrueInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterAccrueInterest represents a AccrueInterest event raised by the MorphoRouter contract.
type MorphoRouterAccrueInterest struct {
	Id             [32]byte
	PrevBorrowRate *big.Int
	Interest       *big.Int
	FeeShares      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x9d9bd501d0657d7dfe415f779a620a62b78bc508ddc0891fbbd8b7ac0f8fce87.
//
// Solidity: event AccrueInterest(bytes32 indexed id, uint256 prevBorrowRate, uint256 interest, uint256 feeShares)
func (_MorphoRouter *MorphoRouterFilterer) FilterAccrueInterest(opts *bind.FilterOpts, id [][32]byte) (*MorphoRouterAccrueInterestIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "AccrueInterest", idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterAccrueInterestIterator{contract: _MorphoRouter.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x9d9bd501d0657d7dfe415f779a620a62b78bc508ddc0891fbbd8b7ac0f8fce87.
//
// Solidity: event AccrueInterest(bytes32 indexed id, uint256 prevBorrowRate, uint256 interest, uint256 feeShares)
func (_MorphoRouter *MorphoRouterFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *MorphoRouterAccrueInterest, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "AccrueInterest", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterAccrueInterest)
				if err := _MorphoRouter.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrueInterest is a log parse operation binding the contract event 0x9d9bd501d0657d7dfe415f779a620a62b78bc508ddc0891fbbd8b7ac0f8fce87.
//
// Solidity: event AccrueInterest(bytes32 indexed id, uint256 prevBorrowRate, uint256 interest, uint256 feeShares)
func (_MorphoRouter *MorphoRouterFilterer) ParseAccrueInterest(log types.Log) (*MorphoRouterAccrueInterest, error) {
	event := new(MorphoRouterAccrueInterest)
	if err := _MorphoRouter.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the MorphoRouter contract.
type MorphoRouterBorrowIterator struct {
	Event *MorphoRouterBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterBorrow represents a Borrow event raised by the MorphoRouter contract.
type MorphoRouterBorrow struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Receiver common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) FilterBorrow(opts *bind.FilterOpts, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (*MorphoRouterBorrowIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "Borrow", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterBorrowIterator{contract: _MorphoRouter.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *MorphoRouterBorrow, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "Borrow", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterBorrow)
				if err := _MorphoRouter.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) ParseBorrow(log types.Log) (*MorphoRouterBorrow, error) {
	event := new(MorphoRouterBorrow)
	if err := _MorphoRouter.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterCreateMarketIterator is returned from FilterCreateMarket and is used to iterate over the raw logs and unpacked data for CreateMarket events raised by the MorphoRouter contract.
type MorphoRouterCreateMarketIterator struct {
	Event *MorphoRouterCreateMarket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterCreateMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterCreateMarket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterCreateMarket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterCreateMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterCreateMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterCreateMarket represents a CreateMarket event raised by the MorphoRouter contract.
type MorphoRouterCreateMarket struct {
	Id           [32]byte
	MarketParams MarketParams
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateMarket is a free log retrieval operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_MorphoRouter *MorphoRouterFilterer) FilterCreateMarket(opts *bind.FilterOpts, id [][32]byte) (*MorphoRouterCreateMarketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "CreateMarket", idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterCreateMarketIterator{contract: _MorphoRouter.contract, event: "CreateMarket", logs: logs, sub: sub}, nil
}

// WatchCreateMarket is a free log subscription operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_MorphoRouter *MorphoRouterFilterer) WatchCreateMarket(opts *bind.WatchOpts, sink chan<- *MorphoRouterCreateMarket, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "CreateMarket", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterCreateMarket)
				if err := _MorphoRouter.contract.UnpackLog(event, "CreateMarket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateMarket is a log parse operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_MorphoRouter *MorphoRouterFilterer) ParseCreateMarket(log types.Log) (*MorphoRouterCreateMarket, error) {
	event := new(MorphoRouterCreateMarket)
	if err := _MorphoRouter.contract.UnpackLog(event, "CreateMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterEnableIrmIterator is returned from FilterEnableIrm and is used to iterate over the raw logs and unpacked data for EnableIrm events raised by the MorphoRouter contract.
type MorphoRouterEnableIrmIterator struct {
	Event *MorphoRouterEnableIrm // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterEnableIrmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterEnableIrm)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterEnableIrm)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterEnableIrmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterEnableIrmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterEnableIrm represents a EnableIrm event raised by the MorphoRouter contract.
type MorphoRouterEnableIrm struct {
	Irm common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableIrm is a free log retrieval operation binding the contract event 0x590e04cdebeccba40f566186b9746ad295a4cd358ea4fefaaea6ce79630d96c0.
//
// Solidity: event EnableIrm(address indexed irm)
func (_MorphoRouter *MorphoRouterFilterer) FilterEnableIrm(opts *bind.FilterOpts, irm []common.Address) (*MorphoRouterEnableIrmIterator, error) {

	var irmRule []interface{}
	for _, irmItem := range irm {
		irmRule = append(irmRule, irmItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "EnableIrm", irmRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterEnableIrmIterator{contract: _MorphoRouter.contract, event: "EnableIrm", logs: logs, sub: sub}, nil
}

// WatchEnableIrm is a free log subscription operation binding the contract event 0x590e04cdebeccba40f566186b9746ad295a4cd358ea4fefaaea6ce79630d96c0.
//
// Solidity: event EnableIrm(address indexed irm)
func (_MorphoRouter *MorphoRouterFilterer) WatchEnableIrm(opts *bind.WatchOpts, sink chan<- *MorphoRouterEnableIrm, irm []common.Address) (event.Subscription, error) {

	var irmRule []interface{}
	for _, irmItem := range irm {
		irmRule = append(irmRule, irmItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "EnableIrm", irmRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterEnableIrm)
				if err := _MorphoRouter.contract.UnpackLog(event, "EnableIrm", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnableIrm is a log parse operation binding the contract event 0x590e04cdebeccba40f566186b9746ad295a4cd358ea4fefaaea6ce79630d96c0.
//
// Solidity: event EnableIrm(address indexed irm)
func (_MorphoRouter *MorphoRouterFilterer) ParseEnableIrm(log types.Log) (*MorphoRouterEnableIrm, error) {
	event := new(MorphoRouterEnableIrm)
	if err := _MorphoRouter.contract.UnpackLog(event, "EnableIrm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterEnableLltvIterator is returned from FilterEnableLltv and is used to iterate over the raw logs and unpacked data for EnableLltv events raised by the MorphoRouter contract.
type MorphoRouterEnableLltvIterator struct {
	Event *MorphoRouterEnableLltv // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterEnableLltvIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterEnableLltv)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterEnableLltv)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterEnableLltvIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterEnableLltvIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterEnableLltv represents a EnableLltv event raised by the MorphoRouter contract.
type MorphoRouterEnableLltv struct {
	Lltv *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEnableLltv is a free log retrieval operation binding the contract event 0x297b80e7a896fad470c630f6575072d609bde997260ff3db851939405ec29139.
//
// Solidity: event EnableLltv(uint256 lltv)
func (_MorphoRouter *MorphoRouterFilterer) FilterEnableLltv(opts *bind.FilterOpts) (*MorphoRouterEnableLltvIterator, error) {

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "EnableLltv")
	if err != nil {
		return nil, err
	}
	return &MorphoRouterEnableLltvIterator{contract: _MorphoRouter.contract, event: "EnableLltv", logs: logs, sub: sub}, nil
}

// WatchEnableLltv is a free log subscription operation binding the contract event 0x297b80e7a896fad470c630f6575072d609bde997260ff3db851939405ec29139.
//
// Solidity: event EnableLltv(uint256 lltv)
func (_MorphoRouter *MorphoRouterFilterer) WatchEnableLltv(opts *bind.WatchOpts, sink chan<- *MorphoRouterEnableLltv) (event.Subscription, error) {

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "EnableLltv")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterEnableLltv)
				if err := _MorphoRouter.contract.UnpackLog(event, "EnableLltv", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnableLltv is a log parse operation binding the contract event 0x297b80e7a896fad470c630f6575072d609bde997260ff3db851939405ec29139.
//
// Solidity: event EnableLltv(uint256 lltv)
func (_MorphoRouter *MorphoRouterFilterer) ParseEnableLltv(log types.Log) (*MorphoRouterEnableLltv, error) {
	event := new(MorphoRouterEnableLltv)
	if err := _MorphoRouter.contract.UnpackLog(event, "EnableLltv", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the MorphoRouter contract.
type MorphoRouterFlashLoanIterator struct {
	Event *MorphoRouterFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterFlashLoan represents a FlashLoan event raised by the MorphoRouter contract.
type MorphoRouterFlashLoan struct {
	Caller common.Address
	Token  common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xc76f1b4fe4396ac07a9fa55a415d4ca430e72651d37d3401f3bed7cb13fc4f12.
//
// Solidity: event FlashLoan(address indexed caller, address indexed token, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) FilterFlashLoan(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*MorphoRouterFlashLoanIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "FlashLoan", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterFlashLoanIterator{contract: _MorphoRouter.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xc76f1b4fe4396ac07a9fa55a415d4ca430e72651d37d3401f3bed7cb13fc4f12.
//
// Solidity: event FlashLoan(address indexed caller, address indexed token, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *MorphoRouterFlashLoan, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "FlashLoan", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterFlashLoan)
				if err := _MorphoRouter.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0xc76f1b4fe4396ac07a9fa55a415d4ca430e72651d37d3401f3bed7cb13fc4f12.
//
// Solidity: event FlashLoan(address indexed caller, address indexed token, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) ParseFlashLoan(log types.Log) (*MorphoRouterFlashLoan, error) {
	event := new(MorphoRouterFlashLoan)
	if err := _MorphoRouter.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterIncrementNonceIterator is returned from FilterIncrementNonce and is used to iterate over the raw logs and unpacked data for IncrementNonce events raised by the MorphoRouter contract.
type MorphoRouterIncrementNonceIterator struct {
	Event *MorphoRouterIncrementNonce // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterIncrementNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterIncrementNonce)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterIncrementNonce)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterIncrementNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterIncrementNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterIncrementNonce represents a IncrementNonce event raised by the MorphoRouter contract.
type MorphoRouterIncrementNonce struct {
	Caller     common.Address
	Authorizer common.Address
	UsedNonce  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIncrementNonce is a free log retrieval operation binding the contract event 0xa58af1a0c70dba0c7aa60d1a1a147ebd61000d1690a968828ac718bca927f2c7.
//
// Solidity: event IncrementNonce(address indexed caller, address indexed authorizer, uint256 usedNonce)
func (_MorphoRouter *MorphoRouterFilterer) FilterIncrementNonce(opts *bind.FilterOpts, caller []common.Address, authorizer []common.Address) (*MorphoRouterIncrementNonceIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "IncrementNonce", callerRule, authorizerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterIncrementNonceIterator{contract: _MorphoRouter.contract, event: "IncrementNonce", logs: logs, sub: sub}, nil
}

// WatchIncrementNonce is a free log subscription operation binding the contract event 0xa58af1a0c70dba0c7aa60d1a1a147ebd61000d1690a968828ac718bca927f2c7.
//
// Solidity: event IncrementNonce(address indexed caller, address indexed authorizer, uint256 usedNonce)
func (_MorphoRouter *MorphoRouterFilterer) WatchIncrementNonce(opts *bind.WatchOpts, sink chan<- *MorphoRouterIncrementNonce, caller []common.Address, authorizer []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "IncrementNonce", callerRule, authorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterIncrementNonce)
				if err := _MorphoRouter.contract.UnpackLog(event, "IncrementNonce", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncrementNonce is a log parse operation binding the contract event 0xa58af1a0c70dba0c7aa60d1a1a147ebd61000d1690a968828ac718bca927f2c7.
//
// Solidity: event IncrementNonce(address indexed caller, address indexed authorizer, uint256 usedNonce)
func (_MorphoRouter *MorphoRouterFilterer) ParseIncrementNonce(log types.Log) (*MorphoRouterIncrementNonce, error) {
	event := new(MorphoRouterIncrementNonce)
	if err := _MorphoRouter.contract.UnpackLog(event, "IncrementNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the MorphoRouter contract.
type MorphoRouterLiquidateIterator struct {
	Event *MorphoRouterLiquidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterLiquidate represents a Liquidate event raised by the MorphoRouter contract.
type MorphoRouterLiquidate struct {
	Id            [32]byte
	Caller        common.Address
	Borrower      common.Address
	RepaidAssets  *big.Int
	RepaidShares  *big.Int
	SeizedAssets  *big.Int
	BadDebtAssets *big.Int
	BadDebtShares *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xa4946ede45d0c6f06a0f5ce92c9ad3b4751452d2fe0e25010783bcab57a67e41.
//
// Solidity: event Liquidate(bytes32 indexed id, address indexed caller, address indexed borrower, uint256 repaidAssets, uint256 repaidShares, uint256 seizedAssets, uint256 badDebtAssets, uint256 badDebtShares)
func (_MorphoRouter *MorphoRouterFilterer) FilterLiquidate(opts *bind.FilterOpts, id [][32]byte, caller []common.Address, borrower []common.Address) (*MorphoRouterLiquidateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "Liquidate", idRule, callerRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterLiquidateIterator{contract: _MorphoRouter.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xa4946ede45d0c6f06a0f5ce92c9ad3b4751452d2fe0e25010783bcab57a67e41.
//
// Solidity: event Liquidate(bytes32 indexed id, address indexed caller, address indexed borrower, uint256 repaidAssets, uint256 repaidShares, uint256 seizedAssets, uint256 badDebtAssets, uint256 badDebtShares)
func (_MorphoRouter *MorphoRouterFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *MorphoRouterLiquidate, id [][32]byte, caller []common.Address, borrower []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "Liquidate", idRule, callerRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterLiquidate)
				if err := _MorphoRouter.contract.UnpackLog(event, "Liquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidate is a log parse operation binding the contract event 0xa4946ede45d0c6f06a0f5ce92c9ad3b4751452d2fe0e25010783bcab57a67e41.
//
// Solidity: event Liquidate(bytes32 indexed id, address indexed caller, address indexed borrower, uint256 repaidAssets, uint256 repaidShares, uint256 seizedAssets, uint256 badDebtAssets, uint256 badDebtShares)
func (_MorphoRouter *MorphoRouterFilterer) ParseLiquidate(log types.Log) (*MorphoRouterLiquidate, error) {
	event := new(MorphoRouterLiquidate)
	if err := _MorphoRouter.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the MorphoRouter contract.
type MorphoRouterRepayIterator struct {
	Event *MorphoRouterRepay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterRepay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterRepay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterRepay represents a Repay event raised by the MorphoRouter contract.
type MorphoRouterRepay struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) FilterRepay(opts *bind.FilterOpts, id [][32]byte, caller []common.Address, onBehalf []common.Address) (*MorphoRouterRepayIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "Repay", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterRepayIterator{contract: _MorphoRouter.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *MorphoRouterRepay, id [][32]byte, caller []common.Address, onBehalf []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "Repay", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterRepay)
				if err := _MorphoRouter.contract.UnpackLog(event, "Repay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepay is a log parse operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) ParseRepay(log types.Log) (*MorphoRouterRepay, error) {
	event := new(MorphoRouterRepay)
	if err := _MorphoRouter.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSetAuthorizationIterator is returned from FilterSetAuthorization and is used to iterate over the raw logs and unpacked data for SetAuthorization events raised by the MorphoRouter contract.
type MorphoRouterSetAuthorizationIterator struct {
	Event *MorphoRouterSetAuthorization // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSetAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSetAuthorization)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSetAuthorization)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSetAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSetAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSetAuthorization represents a SetAuthorization event raised by the MorphoRouter contract.
type MorphoRouterSetAuthorization struct {
	Caller          common.Address
	Authorizer      common.Address
	Authorized      common.Address
	NewIsAuthorized bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetAuthorization is a free log retrieval operation binding the contract event 0xd5e969f01efe921d3f766bdebad25f0a05e3f237311f56482bf132d0326309c0.
//
// Solidity: event SetAuthorization(address indexed caller, address indexed authorizer, address indexed authorized, bool newIsAuthorized)
func (_MorphoRouter *MorphoRouterFilterer) FilterSetAuthorization(opts *bind.FilterOpts, caller []common.Address, authorizer []common.Address, authorized []common.Address) (*MorphoRouterSetAuthorizationIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var authorizedRule []interface{}
	for _, authorizedItem := range authorized {
		authorizedRule = append(authorizedRule, authorizedItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "SetAuthorization", callerRule, authorizerRule, authorizedRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSetAuthorizationIterator{contract: _MorphoRouter.contract, event: "SetAuthorization", logs: logs, sub: sub}, nil
}

// WatchSetAuthorization is a free log subscription operation binding the contract event 0xd5e969f01efe921d3f766bdebad25f0a05e3f237311f56482bf132d0326309c0.
//
// Solidity: event SetAuthorization(address indexed caller, address indexed authorizer, address indexed authorized, bool newIsAuthorized)
func (_MorphoRouter *MorphoRouterFilterer) WatchSetAuthorization(opts *bind.WatchOpts, sink chan<- *MorphoRouterSetAuthorization, caller []common.Address, authorizer []common.Address, authorized []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var authorizedRule []interface{}
	for _, authorizedItem := range authorized {
		authorizedRule = append(authorizedRule, authorizedItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "SetAuthorization", callerRule, authorizerRule, authorizedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSetAuthorization)
				if err := _MorphoRouter.contract.UnpackLog(event, "SetAuthorization", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAuthorization is a log parse operation binding the contract event 0xd5e969f01efe921d3f766bdebad25f0a05e3f237311f56482bf132d0326309c0.
//
// Solidity: event SetAuthorization(address indexed caller, address indexed authorizer, address indexed authorized, bool newIsAuthorized)
func (_MorphoRouter *MorphoRouterFilterer) ParseSetAuthorization(log types.Log) (*MorphoRouterSetAuthorization, error) {
	event := new(MorphoRouterSetAuthorization)
	if err := _MorphoRouter.contract.UnpackLog(event, "SetAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the MorphoRouter contract.
type MorphoRouterSetFeeIterator struct {
	Event *MorphoRouterSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSetFee represents a SetFee event raised by the MorphoRouter contract.
type MorphoRouterSetFee struct {
	Id     [32]byte
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x139d6f58e9a127229667c8e3b36e88890a66cfc8ab1024ddc513e189e125b75b.
//
// Solidity: event SetFee(bytes32 indexed id, uint256 newFee)
func (_MorphoRouter *MorphoRouterFilterer) FilterSetFee(opts *bind.FilterOpts, id [][32]byte) (*MorphoRouterSetFeeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "SetFee", idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSetFeeIterator{contract: _MorphoRouter.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x139d6f58e9a127229667c8e3b36e88890a66cfc8ab1024ddc513e189e125b75b.
//
// Solidity: event SetFee(bytes32 indexed id, uint256 newFee)
func (_MorphoRouter *MorphoRouterFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *MorphoRouterSetFee, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "SetFee", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSetFee)
				if err := _MorphoRouter.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x139d6f58e9a127229667c8e3b36e88890a66cfc8ab1024ddc513e189e125b75b.
//
// Solidity: event SetFee(bytes32 indexed id, uint256 newFee)
func (_MorphoRouter *MorphoRouterFilterer) ParseSetFee(log types.Log) (*MorphoRouterSetFee, error) {
	event := new(MorphoRouterSetFee)
	if err := _MorphoRouter.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the MorphoRouter contract.
type MorphoRouterSetFeeRecipientIterator struct {
	Event *MorphoRouterSetFeeRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSetFeeRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSetFeeRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSetFeeRecipient represents a SetFeeRecipient event raised by the MorphoRouter contract.
type MorphoRouterSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MorphoRouter *MorphoRouterFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts, newFeeRecipient []common.Address) (*MorphoRouterSetFeeRecipientIterator, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSetFeeRecipientIterator{contract: _MorphoRouter.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MorphoRouter *MorphoRouterFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *MorphoRouterSetFeeRecipient, newFeeRecipient []common.Address) (event.Subscription, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSetFeeRecipient)
				if err := _MorphoRouter.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeRecipient is a log parse operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MorphoRouter *MorphoRouterFilterer) ParseSetFeeRecipient(log types.Log) (*MorphoRouterSetFeeRecipient, error) {
	event := new(MorphoRouterSetFeeRecipient)
	if err := _MorphoRouter.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the MorphoRouter contract.
type MorphoRouterSetOwnerIterator struct {
	Event *MorphoRouterSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSetOwner represents a SetOwner event raised by the MorphoRouter contract.
type MorphoRouterSetOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed newOwner)
func (_MorphoRouter *MorphoRouterFilterer) FilterSetOwner(opts *bind.FilterOpts, newOwner []common.Address) (*MorphoRouterSetOwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "SetOwner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSetOwnerIterator{contract: _MorphoRouter.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed newOwner)
func (_MorphoRouter *MorphoRouterFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *MorphoRouterSetOwner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "SetOwner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSetOwner)
				if err := _MorphoRouter.contract.UnpackLog(event, "SetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetOwner is a log parse operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed newOwner)
func (_MorphoRouter *MorphoRouterFilterer) ParseSetOwner(log types.Log) (*MorphoRouterSetOwner, error) {
	event := new(MorphoRouterSetOwner)
	if err := _MorphoRouter.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the MorphoRouter contract.
type MorphoRouterSupplyIterator struct {
	Event *MorphoRouterSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSupply represents a Supply event raised by the MorphoRouter contract.
type MorphoRouterSupply struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0xedf8870433c83823eb071d3df1caa8d008f12f6440918c20d75a3602cda30fe0.
//
// Solidity: event Supply(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) FilterSupply(opts *bind.FilterOpts, id [][32]byte, caller []common.Address, onBehalf []common.Address) (*MorphoRouterSupplyIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "Supply", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSupplyIterator{contract: _MorphoRouter.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0xedf8870433c83823eb071d3df1caa8d008f12f6440918c20d75a3602cda30fe0.
//
// Solidity: event Supply(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *MorphoRouterSupply, id [][32]byte, caller []common.Address, onBehalf []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "Supply", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSupply)
				if err := _MorphoRouter.contract.UnpackLog(event, "Supply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupply is a log parse operation binding the contract event 0xedf8870433c83823eb071d3df1caa8d008f12f6440918c20d75a3602cda30fe0.
//
// Solidity: event Supply(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) ParseSupply(log types.Log) (*MorphoRouterSupply, error) {
	event := new(MorphoRouterSupply)
	if err := _MorphoRouter.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterSupplyCollateralIterator is returned from FilterSupplyCollateral and is used to iterate over the raw logs and unpacked data for SupplyCollateral events raised by the MorphoRouter contract.
type MorphoRouterSupplyCollateralIterator struct {
	Event *MorphoRouterSupplyCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterSupplyCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterSupplyCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterSupplyCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterSupplyCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterSupplyCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterSupplyCollateral represents a SupplyCollateral event raised by the MorphoRouter contract.
type MorphoRouterSupplyCollateral struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Assets   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSupplyCollateral is a free log retrieval operation binding the contract event 0xa3b9472a1399e17e123f3c2e6586c23e504184d504de59cdaa2b375e880c6184.
//
// Solidity: event SupplyCollateral(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) FilterSupplyCollateral(opts *bind.FilterOpts, id [][32]byte, caller []common.Address, onBehalf []common.Address) (*MorphoRouterSupplyCollateralIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "SupplyCollateral", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterSupplyCollateralIterator{contract: _MorphoRouter.contract, event: "SupplyCollateral", logs: logs, sub: sub}, nil
}

// WatchSupplyCollateral is a free log subscription operation binding the contract event 0xa3b9472a1399e17e123f3c2e6586c23e504184d504de59cdaa2b375e880c6184.
//
// Solidity: event SupplyCollateral(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) WatchSupplyCollateral(opts *bind.WatchOpts, sink chan<- *MorphoRouterSupplyCollateral, id [][32]byte, caller []common.Address, onBehalf []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "SupplyCollateral", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterSupplyCollateral)
				if err := _MorphoRouter.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupplyCollateral is a log parse operation binding the contract event 0xa3b9472a1399e17e123f3c2e6586c23e504184d504de59cdaa2b375e880c6184.
//
// Solidity: event SupplyCollateral(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) ParseSupplyCollateral(log types.Log) (*MorphoRouterSupplyCollateral, error) {
	event := new(MorphoRouterSupplyCollateral)
	if err := _MorphoRouter.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MorphoRouter contract.
type MorphoRouterWithdrawIterator struct {
	Event *MorphoRouterWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterWithdraw represents a Withdraw event raised by the MorphoRouter contract.
type MorphoRouterWithdraw struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Receiver common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xa56fc0ad5702ec05ce63666221f796fb62437c32db1aa1aa075fc6484cf58fbf.
//
// Solidity: event Withdraw(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) FilterWithdraw(opts *bind.FilterOpts, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (*MorphoRouterWithdrawIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "Withdraw", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterWithdrawIterator{contract: _MorphoRouter.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xa56fc0ad5702ec05ce63666221f796fb62437c32db1aa1aa075fc6484cf58fbf.
//
// Solidity: event Withdraw(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MorphoRouterWithdraw, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "Withdraw", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterWithdraw)
				if err := _MorphoRouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xa56fc0ad5702ec05ce63666221f796fb62437c32db1aa1aa075fc6484cf58fbf.
//
// Solidity: event Withdraw(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_MorphoRouter *MorphoRouterFilterer) ParseWithdraw(log types.Log) (*MorphoRouterWithdraw, error) {
	event := new(MorphoRouterWithdraw)
	if err := _MorphoRouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRouterWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the MorphoRouter contract.
type MorphoRouterWithdrawCollateralIterator struct {
	Event *MorphoRouterWithdrawCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MorphoRouterWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRouterWithdrawCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MorphoRouterWithdrawCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MorphoRouterWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRouterWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRouterWithdrawCollateral represents a WithdrawCollateral event raised by the MorphoRouter contract.
type MorphoRouterWithdrawCollateral struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Receiver common.Address
	Assets   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0xe80ebd7cc9223d7382aab2e0d1d6155c65651f83d53c8b9b06901d167e321142.
//
// Solidity: event WithdrawCollateral(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (*MorphoRouterWithdrawCollateralIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.FilterLogs(opts, "WithdrawCollateral", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRouterWithdrawCollateralIterator{contract: _MorphoRouter.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0xe80ebd7cc9223d7382aab2e0d1d6155c65651f83d53c8b9b06901d167e321142.
//
// Solidity: event WithdrawCollateral(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *MorphoRouterWithdrawCollateral, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MorphoRouter.contract.WatchLogs(opts, "WithdrawCollateral", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRouterWithdrawCollateral)
				if err := _MorphoRouter.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawCollateral is a log parse operation binding the contract event 0xe80ebd7cc9223d7382aab2e0d1d6155c65651f83d53c8b9b06901d167e321142.
//
// Solidity: event WithdrawCollateral(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets)
func (_MorphoRouter *MorphoRouterFilterer) ParseWithdrawCollateral(log types.Log) (*MorphoRouterWithdrawCollateral, error) {
	event := new(MorphoRouterWithdrawCollateral)
	if err := _MorphoRouter.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
