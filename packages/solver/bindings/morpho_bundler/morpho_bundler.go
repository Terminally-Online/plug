// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho_bundler

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

// IAllowanceTransferPermitDetails is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferPermitDetails struct {
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

// IAllowanceTransferPermitSingle is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferPermitSingle struct {
	Details     IAllowanceTransferPermitDetails
	Spender     common.Address
	SigDeadline *big.Int
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

// Withdrawal is an auto generated low-level Go binding around an user-defined struct.
type Withdrawal struct {
	MarketParams MarketParams
	Amount       *big.Int
}

// MorphoBundlerMetaData contains all meta data concerning the MorphoBundler contract.
var MorphoBundlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"contractIMorpho\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ST_ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WST_ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint48\",\"name\":\"expiration\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"nonce\",\"type\":\"uint48\"}],\"internalType\":\"structIAllowanceTransfer.PermitDetails\",\"name\":\"details\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structIAllowanceTransfer.PermitSingle\",\"name\":\"permitSingle\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"}],\"name\":\"approve2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"erc20Transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"erc20TransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"erc20WrapperDepositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"erc20WrapperWithdrawTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"erc4626Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAssets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"erc4626Mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAssets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"erc4626Redeem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"erc4626Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"morphoBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"morphoFlashLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"morphoRepay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorized\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structAuthorization\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"}],\"name\":\"morphoSetAuthorizationWithSig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"morphoSupply\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"morphoSupplyCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"morphoWithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"morphoWithdrawCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"nativeTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoFlashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoRepay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoSupplyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"}],\"name\":\"permitDai\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicAllocator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"internalType\":\"structWithdrawal[]\",\"name\":\"withdrawals\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"supplyMarketParams\",\"type\":\"tuple\"}],\"name\":\"reallocateTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"stakeEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrapNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrapStEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"}],\"name\":\"urdClaim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"wrapNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"wrapStEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MorphoBundlerABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoBundlerMetaData.ABI instead.
var MorphoBundlerABI = MorphoBundlerMetaData.ABI

// MorphoBundler is an auto generated Go binding around an Ethereum contract.
type MorphoBundler struct {
	MorphoBundlerCaller     // Read-only binding to the contract
	MorphoBundlerTransactor // Write-only binding to the contract
	MorphoBundlerFilterer   // Log filterer for contract events
}

// MorphoBundlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoBundlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoBundlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoBundlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoBundlerSession struct {
	Contract     *MorphoBundler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoBundlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoBundlerCallerSession struct {
	Contract *MorphoBundlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MorphoBundlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoBundlerTransactorSession struct {
	Contract     *MorphoBundlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MorphoBundlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoBundlerRaw struct {
	Contract *MorphoBundler // Generic contract binding to access the raw methods on
}

// MorphoBundlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoBundlerCallerRaw struct {
	Contract *MorphoBundlerCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoBundlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoBundlerTransactorRaw struct {
	Contract *MorphoBundlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphoBundler creates a new instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundler(address common.Address, backend bind.ContractBackend) (*MorphoBundler, error) {
	contract, err := bindMorphoBundler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphoBundler{MorphoBundlerCaller: MorphoBundlerCaller{contract: contract}, MorphoBundlerTransactor: MorphoBundlerTransactor{contract: contract}, MorphoBundlerFilterer: MorphoBundlerFilterer{contract: contract}}, nil
}

// NewMorphoBundlerCaller creates a new read-only instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerCaller(address common.Address, caller bind.ContractCaller) (*MorphoBundlerCaller, error) {
	contract, err := bindMorphoBundler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerCaller{contract: contract}, nil
}

// NewMorphoBundlerTransactor creates a new write-only instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoBundlerTransactor, error) {
	contract, err := bindMorphoBundler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerTransactor{contract: contract}, nil
}

// NewMorphoBundlerFilterer creates a new log filterer instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoBundlerFilterer, error) {
	contract, err := bindMorphoBundler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerFilterer{contract: contract}, nil
}

// bindMorphoBundler binds a generic wrapper to an already deployed contract.
func bindMorphoBundler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoBundlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoBundler *MorphoBundlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoBundler.Contract.MorphoBundlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoBundler *MorphoBundlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBundlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoBundler *MorphoBundlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBundlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoBundler *MorphoBundlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoBundler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoBundler *MorphoBundlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoBundler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoBundler *MorphoBundlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoBundler.Contract.contract.Transact(opts, method, params...)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) MORPHO() (common.Address, error) {
	return _MorphoBundler.Contract.MORPHO(&_MorphoBundler.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) MORPHO() (common.Address, error) {
	return _MorphoBundler.Contract.MORPHO(&_MorphoBundler.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0x338346d2.
//
// Solidity: function ST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "ST_ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0x338346d2.
//
// Solidity: function ST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) STETH() (common.Address, error) {
	return _MorphoBundler.Contract.STETH(&_MorphoBundler.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0x338346d2.
//
// Solidity: function ST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) STETH() (common.Address, error) {
	return _MorphoBundler.Contract.STETH(&_MorphoBundler.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _MorphoBundler.Contract.WRAPPEDNATIVE(&_MorphoBundler.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _MorphoBundler.Contract.WRAPPEDNATIVE(&_MorphoBundler.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0x5664cb48.
//
// Solidity: function WST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) WSTETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "WST_ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WSTETH is a free data retrieval call binding the contract method 0x5664cb48.
//
// Solidity: function WST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) WSTETH() (common.Address, error) {
	return _MorphoBundler.Contract.WSTETH(&_MorphoBundler.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0x5664cb48.
//
// Solidity: function WST_ETH() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) WSTETH() (common.Address, error) {
	return _MorphoBundler.Contract.WSTETH(&_MorphoBundler.CallOpts)
}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) Initiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "initiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) Initiator() (common.Address, error) {
	return _MorphoBundler.Contract.Initiator(&_MorphoBundler.CallOpts)
}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) Initiator() (common.Address, error) {
	return _MorphoBundler.Contract.Initiator(&_MorphoBundler.CallOpts)
}

// Approve2 is a paid mutator transaction binding the contract method 0xaf504202.
//
// Solidity: function approve2(((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Approve2(opts *bind.TransactOpts, permitSingle IAllowanceTransferPermitSingle, signature []byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "approve2", permitSingle, signature, skipRevert)
}

// Approve2 is a paid mutator transaction binding the contract method 0xaf504202.
//
// Solidity: function approve2(((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Approve2(permitSingle IAllowanceTransferPermitSingle, signature []byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Approve2(&_MorphoBundler.TransactOpts, permitSingle, signature, skipRevert)
}

// Approve2 is a paid mutator transaction binding the contract method 0xaf504202.
//
// Solidity: function approve2(((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Approve2(permitSingle IAllowanceTransferPermitSingle, signature []byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Approve2(&_MorphoBundler.TransactOpts, permitSingle, signature, skipRevert)
}

// Erc20Transfer is a paid mutator transaction binding the contract method 0x3790767d.
//
// Solidity: function erc20Transfer(address asset, address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc20Transfer(opts *bind.TransactOpts, asset common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc20Transfer", asset, recipient, amount)
}

// Erc20Transfer is a paid mutator transaction binding the contract method 0x3790767d.
//
// Solidity: function erc20Transfer(address asset, address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc20Transfer(asset common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20Transfer(&_MorphoBundler.TransactOpts, asset, recipient, amount)
}

// Erc20Transfer is a paid mutator transaction binding the contract method 0x3790767d.
//
// Solidity: function erc20Transfer(address asset, address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc20Transfer(asset common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20Transfer(&_MorphoBundler.TransactOpts, asset, recipient, amount)
}

// Erc20TransferFrom is a paid mutator transaction binding the contract method 0x70dc41fe.
//
// Solidity: function erc20TransferFrom(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc20TransferFrom(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc20TransferFrom", asset, amount)
}

// Erc20TransferFrom is a paid mutator transaction binding the contract method 0x70dc41fe.
//
// Solidity: function erc20TransferFrom(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc20TransferFrom(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20TransferFrom(&_MorphoBundler.TransactOpts, asset, amount)
}

// Erc20TransferFrom is a paid mutator transaction binding the contract method 0x70dc41fe.
//
// Solidity: function erc20TransferFrom(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc20TransferFrom(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20TransferFrom(&_MorphoBundler.TransactOpts, asset, amount)
}

// Erc20WrapperDepositFor is a paid mutator transaction binding the contract method 0x12c0d70b.
//
// Solidity: function erc20WrapperDepositFor(address wrapper, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc20WrapperDepositFor(opts *bind.TransactOpts, wrapper common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc20WrapperDepositFor", wrapper, amount)
}

// Erc20WrapperDepositFor is a paid mutator transaction binding the contract method 0x12c0d70b.
//
// Solidity: function erc20WrapperDepositFor(address wrapper, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc20WrapperDepositFor(wrapper common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20WrapperDepositFor(&_MorphoBundler.TransactOpts, wrapper, amount)
}

// Erc20WrapperDepositFor is a paid mutator transaction binding the contract method 0x12c0d70b.
//
// Solidity: function erc20WrapperDepositFor(address wrapper, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc20WrapperDepositFor(wrapper common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20WrapperDepositFor(&_MorphoBundler.TransactOpts, wrapper, amount)
}

// Erc20WrapperWithdrawTo is a paid mutator transaction binding the contract method 0x60244408.
//
// Solidity: function erc20WrapperWithdrawTo(address wrapper, address account, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc20WrapperWithdrawTo(opts *bind.TransactOpts, wrapper common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc20WrapperWithdrawTo", wrapper, account, amount)
}

// Erc20WrapperWithdrawTo is a paid mutator transaction binding the contract method 0x60244408.
//
// Solidity: function erc20WrapperWithdrawTo(address wrapper, address account, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc20WrapperWithdrawTo(wrapper common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20WrapperWithdrawTo(&_MorphoBundler.TransactOpts, wrapper, account, amount)
}

// Erc20WrapperWithdrawTo is a paid mutator transaction binding the contract method 0x60244408.
//
// Solidity: function erc20WrapperWithdrawTo(address wrapper, address account, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc20WrapperWithdrawTo(wrapper common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc20WrapperWithdrawTo(&_MorphoBundler.TransactOpts, wrapper, account, amount)
}

// Erc4626Deposit is a paid mutator transaction binding the contract method 0x6ef5eeae.
//
// Solidity: function erc4626Deposit(address vault, uint256 assets, uint256 minShares, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc4626Deposit(opts *bind.TransactOpts, vault common.Address, assets *big.Int, minShares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc4626Deposit", vault, assets, minShares, receiver)
}

// Erc4626Deposit is a paid mutator transaction binding the contract method 0x6ef5eeae.
//
// Solidity: function erc4626Deposit(address vault, uint256 assets, uint256 minShares, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc4626Deposit(vault common.Address, assets *big.Int, minShares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Deposit(&_MorphoBundler.TransactOpts, vault, assets, minShares, receiver)
}

// Erc4626Deposit is a paid mutator transaction binding the contract method 0x6ef5eeae.
//
// Solidity: function erc4626Deposit(address vault, uint256 assets, uint256 minShares, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc4626Deposit(vault common.Address, assets *big.Int, minShares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Deposit(&_MorphoBundler.TransactOpts, vault, assets, minShares, receiver)
}

// Erc4626Mint is a paid mutator transaction binding the contract method 0x39029ab6.
//
// Solidity: function erc4626Mint(address vault, uint256 shares, uint256 maxAssets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc4626Mint(opts *bind.TransactOpts, vault common.Address, shares *big.Int, maxAssets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc4626Mint", vault, shares, maxAssets, receiver)
}

// Erc4626Mint is a paid mutator transaction binding the contract method 0x39029ab6.
//
// Solidity: function erc4626Mint(address vault, uint256 shares, uint256 maxAssets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc4626Mint(vault common.Address, shares *big.Int, maxAssets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Mint(&_MorphoBundler.TransactOpts, vault, shares, maxAssets, receiver)
}

// Erc4626Mint is a paid mutator transaction binding the contract method 0x39029ab6.
//
// Solidity: function erc4626Mint(address vault, uint256 shares, uint256 maxAssets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc4626Mint(vault common.Address, shares *big.Int, maxAssets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Mint(&_MorphoBundler.TransactOpts, vault, shares, maxAssets, receiver)
}

// Erc4626Redeem is a paid mutator transaction binding the contract method 0xa7f6e606.
//
// Solidity: function erc4626Redeem(address vault, uint256 shares, uint256 minAssets, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc4626Redeem(opts *bind.TransactOpts, vault common.Address, shares *big.Int, minAssets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc4626Redeem", vault, shares, minAssets, receiver, owner)
}

// Erc4626Redeem is a paid mutator transaction binding the contract method 0xa7f6e606.
//
// Solidity: function erc4626Redeem(address vault, uint256 shares, uint256 minAssets, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc4626Redeem(vault common.Address, shares *big.Int, minAssets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Redeem(&_MorphoBundler.TransactOpts, vault, shares, minAssets, receiver, owner)
}

// Erc4626Redeem is a paid mutator transaction binding the contract method 0xa7f6e606.
//
// Solidity: function erc4626Redeem(address vault, uint256 shares, uint256 minAssets, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc4626Redeem(vault common.Address, shares *big.Int, minAssets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Redeem(&_MorphoBundler.TransactOpts, vault, shares, minAssets, receiver, owner)
}

// Erc4626Withdraw is a paid mutator transaction binding the contract method 0xc9565706.
//
// Solidity: function erc4626Withdraw(address vault, uint256 assets, uint256 maxShares, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Erc4626Withdraw(opts *bind.TransactOpts, vault common.Address, assets *big.Int, maxShares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "erc4626Withdraw", vault, assets, maxShares, receiver, owner)
}

// Erc4626Withdraw is a paid mutator transaction binding the contract method 0xc9565706.
//
// Solidity: function erc4626Withdraw(address vault, uint256 assets, uint256 maxShares, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Erc4626Withdraw(vault common.Address, assets *big.Int, maxShares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Withdraw(&_MorphoBundler.TransactOpts, vault, assets, maxShares, receiver, owner)
}

// Erc4626Withdraw is a paid mutator transaction binding the contract method 0xc9565706.
//
// Solidity: function erc4626Withdraw(address vault, uint256 assets, uint256 maxShares, address receiver, address owner) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Erc4626Withdraw(vault common.Address, assets *big.Int, maxShares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Erc4626Withdraw(&_MorphoBundler.TransactOpts, vault, assets, maxShares, receiver, owner)
}

// MorphoBorrow is a paid mutator transaction binding the contract method 0x62577ad0.
//
// Solidity: function morphoBorrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoBorrow(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoBorrow", marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoBorrow is a paid mutator transaction binding the contract method 0x62577ad0.
//
// Solidity: function morphoBorrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoBorrow(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBorrow(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoBorrow is a paid mutator transaction binding the contract method 0x62577ad0.
//
// Solidity: function morphoBorrow((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoBorrow(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBorrow(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoFlashLoan is a paid mutator transaction binding the contract method 0xe2975912.
//
// Solidity: function morphoFlashLoan(address token, uint256 assets, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoFlashLoan(opts *bind.TransactOpts, token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoFlashLoan", token, assets, data)
}

// MorphoFlashLoan is a paid mutator transaction binding the contract method 0xe2975912.
//
// Solidity: function morphoFlashLoan(address token, uint256 assets, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoFlashLoan(token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoFlashLoan(&_MorphoBundler.TransactOpts, token, assets, data)
}

// MorphoFlashLoan is a paid mutator transaction binding the contract method 0xe2975912.
//
// Solidity: function morphoFlashLoan(address token, uint256 assets, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoFlashLoan(token common.Address, assets *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoFlashLoan(&_MorphoBundler.TransactOpts, token, assets, data)
}

// MorphoRepay is a paid mutator transaction binding the contract method 0x4d5fcf68.
//
// Solidity: function morphoRepay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoRepay(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoRepay", marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoRepay is a paid mutator transaction binding the contract method 0x4d5fcf68.
//
// Solidity: function morphoRepay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoRepay(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoRepay(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoRepay is a paid mutator transaction binding the contract method 0x4d5fcf68.
//
// Solidity: function morphoRepay((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoRepay(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoRepay(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoSetAuthorizationWithSig is a paid mutator transaction binding the contract method 0xbea88fda.
//
// Solidity: function morphoSetAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoSetAuthorizationWithSig(opts *bind.TransactOpts, authorization Authorization, signature Signature, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoSetAuthorizationWithSig", authorization, signature, skipRevert)
}

// MorphoSetAuthorizationWithSig is a paid mutator transaction binding the contract method 0xbea88fda.
//
// Solidity: function morphoSetAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoSetAuthorizationWithSig(authorization Authorization, signature Signature, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSetAuthorizationWithSig(&_MorphoBundler.TransactOpts, authorization, signature, skipRevert)
}

// MorphoSetAuthorizationWithSig is a paid mutator transaction binding the contract method 0xbea88fda.
//
// Solidity: function morphoSetAuthorizationWithSig((address,address,bool,uint256,uint256) authorization, (uint8,bytes32,bytes32) signature, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoSetAuthorizationWithSig(authorization Authorization, signature Signature, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSetAuthorizationWithSig(&_MorphoBundler.TransactOpts, authorization, signature, skipRevert)
}

// MorphoSupply is a paid mutator transaction binding the contract method 0x5b866db6.
//
// Solidity: function morphoSupply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoSupply(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoSupply", marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoSupply is a paid mutator transaction binding the contract method 0x5b866db6.
//
// Solidity: function morphoSupply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoSupply(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSupply(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoSupply is a paid mutator transaction binding the contract method 0x5b866db6.
//
// Solidity: function morphoSupply((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoSupply(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSupply(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, onBehalf, data)
}

// MorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xca463673.
//
// Solidity: function morphoSupplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoSupplyCollateral(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoSupplyCollateral", marketParams, assets, onBehalf, data)
}

// MorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xca463673.
//
// Solidity: function morphoSupplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoSupplyCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSupplyCollateral(&_MorphoBundler.TransactOpts, marketParams, assets, onBehalf, data)
}

// MorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xca463673.
//
// Solidity: function morphoSupplyCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address onBehalf, bytes data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoSupplyCollateral(marketParams MarketParams, assets *big.Int, onBehalf common.Address, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoSupplyCollateral(&_MorphoBundler.TransactOpts, marketParams, assets, onBehalf, data)
}

// MorphoWithdraw is a paid mutator transaction binding the contract method 0x84d287ef.
//
// Solidity: function morphoWithdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoWithdraw(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoWithdraw", marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoWithdraw is a paid mutator transaction binding the contract method 0x84d287ef.
//
// Solidity: function morphoWithdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoWithdraw(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoWithdraw(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoWithdraw is a paid mutator transaction binding the contract method 0x84d287ef.
//
// Solidity: function morphoWithdraw((address,address,address,address,uint256) marketParams, uint256 assets, uint256 shares, uint256 slippageAmount, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoWithdraw(marketParams MarketParams, assets *big.Int, shares *big.Int, slippageAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoWithdraw(&_MorphoBundler.TransactOpts, marketParams, assets, shares, slippageAmount, receiver)
}

// MorphoWithdrawCollateral is a paid mutator transaction binding the contract method 0x1af3bbc6.
//
// Solidity: function morphoWithdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) MorphoWithdrawCollateral(opts *bind.TransactOpts, marketParams MarketParams, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "morphoWithdrawCollateral", marketParams, assets, receiver)
}

// MorphoWithdrawCollateral is a paid mutator transaction binding the contract method 0x1af3bbc6.
//
// Solidity: function morphoWithdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerSession) MorphoWithdrawCollateral(marketParams MarketParams, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoWithdrawCollateral(&_MorphoBundler.TransactOpts, marketParams, assets, receiver)
}

// MorphoWithdrawCollateral is a paid mutator transaction binding the contract method 0x1af3bbc6.
//
// Solidity: function morphoWithdrawCollateral((address,address,address,address,uint256) marketParams, uint256 assets, address receiver) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) MorphoWithdrawCollateral(marketParams MarketParams, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoWithdrawCollateral(&_MorphoBundler.TransactOpts, marketParams, assets, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Multicall(&_MorphoBundler.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Multicall(&_MorphoBundler.TransactOpts, data)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) NativeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "nativeTransfer", recipient, amount)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) NativeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.NativeTransfer(&_MorphoBundler.TransactOpts, recipient, amount)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) NativeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.NativeTransfer(&_MorphoBundler.TransactOpts, recipient, amount)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactor) OnMorphoFlashLoan(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "onMorphoFlashLoan", arg0, data)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerSession) OnMorphoFlashLoan(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoFlashLoan(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) OnMorphoFlashLoan(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoFlashLoan(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoRepay is a paid mutator transaction binding the contract method 0x05b4591c.
//
// Solidity: function onMorphoRepay(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactor) OnMorphoRepay(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "onMorphoRepay", arg0, data)
}

// OnMorphoRepay is a paid mutator transaction binding the contract method 0x05b4591c.
//
// Solidity: function onMorphoRepay(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerSession) OnMorphoRepay(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoRepay(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoRepay is a paid mutator transaction binding the contract method 0x05b4591c.
//
// Solidity: function onMorphoRepay(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) OnMorphoRepay(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoRepay(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoSupply is a paid mutator transaction binding the contract method 0x2075be03.
//
// Solidity: function onMorphoSupply(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactor) OnMorphoSupply(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "onMorphoSupply", arg0, data)
}

// OnMorphoSupply is a paid mutator transaction binding the contract method 0x2075be03.
//
// Solidity: function onMorphoSupply(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerSession) OnMorphoSupply(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoSupply(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoSupply is a paid mutator transaction binding the contract method 0x2075be03.
//
// Solidity: function onMorphoSupply(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) OnMorphoSupply(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoSupply(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xb1022fdf.
//
// Solidity: function onMorphoSupplyCollateral(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactor) OnMorphoSupplyCollateral(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "onMorphoSupplyCollateral", arg0, data)
}

// OnMorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xb1022fdf.
//
// Solidity: function onMorphoSupplyCollateral(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerSession) OnMorphoSupplyCollateral(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoSupplyCollateral(&_MorphoBundler.TransactOpts, arg0, data)
}

// OnMorphoSupplyCollateral is a paid mutator transaction binding the contract method 0xb1022fdf.
//
// Solidity: function onMorphoSupplyCollateral(uint256 , bytes data) returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) OnMorphoSupplyCollateral(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphoBundler.Contract.OnMorphoSupplyCollateral(&_MorphoBundler.TransactOpts, arg0, data)
}

// Permit is a paid mutator transaction binding the contract method 0xa184a5a3.
//
// Solidity: function permit(address asset, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Permit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "permit", asset, amount, deadline, v, r, s, skipRevert)
}

// Permit is a paid mutator transaction binding the contract method 0xa184a5a3.
//
// Solidity: function permit(address asset, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Permit(asset common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Permit(&_MorphoBundler.TransactOpts, asset, amount, deadline, v, r, s, skipRevert)
}

// Permit is a paid mutator transaction binding the contract method 0xa184a5a3.
//
// Solidity: function permit(address asset, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Permit(asset common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Permit(&_MorphoBundler.TransactOpts, asset, amount, deadline, v, r, s, skipRevert)
}

// PermitDai is a paid mutator transaction binding the contract method 0x7336bc6a.
//
// Solidity: function permitDai(uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) PermitDai(opts *bind.TransactOpts, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "permitDai", nonce, expiry, allowed, v, r, s, skipRevert)
}

// PermitDai is a paid mutator transaction binding the contract method 0x7336bc6a.
//
// Solidity: function permitDai(uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerSession) PermitDai(nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.PermitDai(&_MorphoBundler.TransactOpts, nonce, expiry, allowed, v, r, s, skipRevert)
}

// PermitDai is a paid mutator transaction binding the contract method 0x7336bc6a.
//
// Solidity: function permitDai(uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) PermitDai(nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.PermitDai(&_MorphoBundler.TransactOpts, nonce, expiry, allowed, v, r, s, skipRevert)
}

// ReallocateTo is a paid mutator transaction binding the contract method 0xef653419.
//
// Solidity: function reallocateTo(address publicAllocator, address vault, uint256 value, ((address,address,address,address,uint256),uint128)[] withdrawals, (address,address,address,address,uint256) supplyMarketParams) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) ReallocateTo(opts *bind.TransactOpts, publicAllocator common.Address, vault common.Address, value *big.Int, withdrawals []Withdrawal, supplyMarketParams MarketParams) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "reallocateTo", publicAllocator, vault, value, withdrawals, supplyMarketParams)
}

// ReallocateTo is a paid mutator transaction binding the contract method 0xef653419.
//
// Solidity: function reallocateTo(address publicAllocator, address vault, uint256 value, ((address,address,address,address,uint256),uint128)[] withdrawals, (address,address,address,address,uint256) supplyMarketParams) payable returns()
func (_MorphoBundler *MorphoBundlerSession) ReallocateTo(publicAllocator common.Address, vault common.Address, value *big.Int, withdrawals []Withdrawal, supplyMarketParams MarketParams) (*types.Transaction, error) {
	return _MorphoBundler.Contract.ReallocateTo(&_MorphoBundler.TransactOpts, publicAllocator, vault, value, withdrawals, supplyMarketParams)
}

// ReallocateTo is a paid mutator transaction binding the contract method 0xef653419.
//
// Solidity: function reallocateTo(address publicAllocator, address vault, uint256 value, ((address,address,address,address,uint256),uint128)[] withdrawals, (address,address,address,address,uint256) supplyMarketParams) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) ReallocateTo(publicAllocator common.Address, vault common.Address, value *big.Int, withdrawals []Withdrawal, supplyMarketParams MarketParams) (*types.Transaction, error) {
	return _MorphoBundler.Contract.ReallocateTo(&_MorphoBundler.TransactOpts, publicAllocator, vault, value, withdrawals, supplyMarketParams)
}

// StakeEth is a paid mutator transaction binding the contract method 0xd27dd11e.
//
// Solidity: function stakeEth(uint256 amount, uint256 minShares, address referral) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) StakeEth(opts *bind.TransactOpts, amount *big.Int, minShares *big.Int, referral common.Address) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "stakeEth", amount, minShares, referral)
}

// StakeEth is a paid mutator transaction binding the contract method 0xd27dd11e.
//
// Solidity: function stakeEth(uint256 amount, uint256 minShares, address referral) payable returns()
func (_MorphoBundler *MorphoBundlerSession) StakeEth(amount *big.Int, minShares *big.Int, referral common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.StakeEth(&_MorphoBundler.TransactOpts, amount, minShares, referral)
}

// StakeEth is a paid mutator transaction binding the contract method 0xd27dd11e.
//
// Solidity: function stakeEth(uint256 amount, uint256 minShares, address referral) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) StakeEth(amount *big.Int, minShares *big.Int, referral common.Address) (*types.Transaction, error) {
	return _MorphoBundler.Contract.StakeEth(&_MorphoBundler.TransactOpts, amount, minShares, referral)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x54c53ef0.
//
// Solidity: function transferFrom2(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) TransferFrom2(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "transferFrom2", asset, amount)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x54c53ef0.
//
// Solidity: function transferFrom2(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) TransferFrom2(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.TransferFrom2(&_MorphoBundler.TransactOpts, asset, amount)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x54c53ef0.
//
// Solidity: function transferFrom2(address asset, uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) TransferFrom2(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.TransferFrom2(&_MorphoBundler.TransactOpts, asset, amount)
}

// UnwrapNative is a paid mutator transaction binding the contract method 0x34b10a6d.
//
// Solidity: function unwrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) UnwrapNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "unwrapNative", amount)
}

// UnwrapNative is a paid mutator transaction binding the contract method 0x34b10a6d.
//
// Solidity: function unwrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) UnwrapNative(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UnwrapNative(&_MorphoBundler.TransactOpts, amount)
}

// UnwrapNative is a paid mutator transaction binding the contract method 0x34b10a6d.
//
// Solidity: function unwrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) UnwrapNative(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UnwrapNative(&_MorphoBundler.TransactOpts, amount)
}

// UnwrapStEth is a paid mutator transaction binding the contract method 0x240f78a3.
//
// Solidity: function unwrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) UnwrapStEth(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "unwrapStEth", amount)
}

// UnwrapStEth is a paid mutator transaction binding the contract method 0x240f78a3.
//
// Solidity: function unwrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) UnwrapStEth(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UnwrapStEth(&_MorphoBundler.TransactOpts, amount)
}

// UnwrapStEth is a paid mutator transaction binding the contract method 0x240f78a3.
//
// Solidity: function unwrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) UnwrapStEth(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UnwrapStEth(&_MorphoBundler.TransactOpts, amount)
}

// UrdClaim is a paid mutator transaction binding the contract method 0x6b89026a.
//
// Solidity: function urdClaim(address distributor, address account, address reward, uint256 amount, bytes32[] proof, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) UrdClaim(opts *bind.TransactOpts, distributor common.Address, account common.Address, reward common.Address, amount *big.Int, proof [][32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "urdClaim", distributor, account, reward, amount, proof, skipRevert)
}

// UrdClaim is a paid mutator transaction binding the contract method 0x6b89026a.
//
// Solidity: function urdClaim(address distributor, address account, address reward, uint256 amount, bytes32[] proof, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerSession) UrdClaim(distributor common.Address, account common.Address, reward common.Address, amount *big.Int, proof [][32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UrdClaim(&_MorphoBundler.TransactOpts, distributor, account, reward, amount, proof, skipRevert)
}

// UrdClaim is a paid mutator transaction binding the contract method 0x6b89026a.
//
// Solidity: function urdClaim(address distributor, address account, address reward, uint256 amount, bytes32[] proof, bool skipRevert) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) UrdClaim(distributor common.Address, account common.Address, reward common.Address, amount *big.Int, proof [][32]byte, skipRevert bool) (*types.Transaction, error) {
	return _MorphoBundler.Contract.UrdClaim(&_MorphoBundler.TransactOpts, distributor, account, reward, amount, proof, skipRevert)
}

// WrapNative is a paid mutator transaction binding the contract method 0x9169d833.
//
// Solidity: function wrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) WrapNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "wrapNative", amount)
}

// WrapNative is a paid mutator transaction binding the contract method 0x9169d833.
//
// Solidity: function wrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) WrapNative(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.WrapNative(&_MorphoBundler.TransactOpts, amount)
}

// WrapNative is a paid mutator transaction binding the contract method 0x9169d833.
//
// Solidity: function wrapNative(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) WrapNative(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.WrapNative(&_MorphoBundler.TransactOpts, amount)
}

// WrapStEth is a paid mutator transaction binding the contract method 0xbc2a294c.
//
// Solidity: function wrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) WrapStEth(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "wrapStEth", amount)
}

// WrapStEth is a paid mutator transaction binding the contract method 0xbc2a294c.
//
// Solidity: function wrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerSession) WrapStEth(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.WrapStEth(&_MorphoBundler.TransactOpts, amount)
}

// WrapStEth is a paid mutator transaction binding the contract method 0xbc2a294c.
//
// Solidity: function wrapStEth(uint256 amount) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) WrapStEth(amount *big.Int) (*types.Transaction, error) {
	return _MorphoBundler.Contract.WrapStEth(&_MorphoBundler.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoBundler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MorphoBundler *MorphoBundlerSession) Receive() (*types.Transaction, error) {
	return _MorphoBundler.Contract.Receive(&_MorphoBundler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Receive() (*types.Transaction, error) {
	return _MorphoBundler.Contract.Receive(&_MorphoBundler.TransactOpts)
}
