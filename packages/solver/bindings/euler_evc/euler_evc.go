// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_evc

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

// IEVCBatchItem is an auto generated low-level Go binding around an user-defined struct.
type IEVCBatchItem struct {
	TargetContract    common.Address
	OnBehalfOfAccount common.Address
	Value             *big.Int
	Data              []byte
}

// IEVCBatchItemResult is an auto generated low-level Go binding around an user-defined struct.
type IEVCBatchItemResult struct {
	Success bool
	Result  []byte
}

// IEVCStatusCheckResult is an auto generated low-level Go binding around an user-defined struct.
type IEVCStatusCheckResult struct {
	CheckedAddress common.Address
	IsValid        bool
	Result         []byte
}

// EulerEvcMetaData contains all meta data concerning the EulerEvc contract.
var EulerEvcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EVC_BatchPanic\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_ChecksReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_ControlCollateralReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_ControllerViolation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_EmptyError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidOperatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_LockdownMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_OnBehalfOfAccountNotAuthenticated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_PermitDisabledMode\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.BatchItemResult[]\",\"name\":\"batchItemsResult\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"checkedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.StatusCheckResult[]\",\"name\":\"accountsStatusResult\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"checkedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.StatusCheckResult[]\",\"name\":\"vaultsStatusResult\",\"type\":\"tuple[]\"}],\"name\":\"EVC_RevertedBatchResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_SimulationBatchNested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyElements\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"AccountStatusCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"onBehalfOfAddressPrefix\",\"type\":\"bytes19\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"CallWithContext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"CollateralStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ControllerStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"LockdownModeStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonceNamespace\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonceNamespace\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountOperatorAuthorized\",\"type\":\"uint256\"}],\"name\":\"OperatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PermitDisabledModeStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultStatusCheck\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"areChecksDeferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"areChecksInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.BatchItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"batch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.BatchItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"batchRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.BatchItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"batchSimulation\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.BatchItemResult[]\",\"name\":\"batchItemsResult\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"checkedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.StatusCheckResult[]\",\"name\":\"accountsStatusCheckResult\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"checkedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structIEVC.StatusCheckResult[]\",\"name\":\"vaultsStatusCheckResult\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"controlCollateral\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"disableCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"disableController\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"enableCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"enableController\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"forgiveAccountStatusCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgiveVaultStatusCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAddressPrefix\",\"outputs\":[{\"internalType\":\"bytes19\",\"name\":\"\",\"type\":\"bytes19\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCollaterals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getControllers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controllerToCheck\",\"type\":\"address\"}],\"name\":\"getCurrentOnBehalfOfAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOfAccount\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"controllerEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLastAccountStatusCheckTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"uint256\",\"name\":\"nonceNamespace\",\"type\":\"uint256\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRawExecutionContext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"context\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"otherAccount\",\"type\":\"address\"}],\"name\":\"haveCommonOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isAccountOperatorAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAccountStatusCheckDeferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"isCollateralEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isControlCollateralInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"isControllerEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"}],\"name\":\"isLockdownMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOperatorAuthenticated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"}],\"name\":\"isPermitDisabledMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSimulationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"isVaultStatusCheckDeferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonceNamespace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"index1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"index2\",\"type\":\"uint8\"}],\"name\":\"reorderCollaterals\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"requireAccountAndVaultStatusCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"requireAccountStatusCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireVaultStatusCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"setAccountOperator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setLockdownMode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"uint256\",\"name\":\"nonceNamespace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorBitField\",\"type\":\"uint256\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPermitDisabledMode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EulerEvcABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEvcMetaData.ABI instead.
var EulerEvcABI = EulerEvcMetaData.ABI

// EulerEvc is an auto generated Go binding around an Ethereum contract.
type EulerEvc struct {
	EulerEvcCaller     // Read-only binding to the contract
	EulerEvcTransactor // Write-only binding to the contract
	EulerEvcFilterer   // Log filterer for contract events
}

// EulerEvcCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEvcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEvcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEvcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEvcSession struct {
	Contract     *EulerEvc         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerEvcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEvcCallerSession struct {
	Contract *EulerEvcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EulerEvcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEvcTransactorSession struct {
	Contract     *EulerEvcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EulerEvcRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEvcRaw struct {
	Contract *EulerEvc // Generic contract binding to access the raw methods on
}

// EulerEvcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEvcCallerRaw struct {
	Contract *EulerEvcCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEvcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEvcTransactorRaw struct {
	Contract *EulerEvcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEvc creates a new instance of EulerEvc, bound to a specific deployed contract.
func NewEulerEvc(address common.Address, backend bind.ContractBackend) (*EulerEvc, error) {
	contract, err := bindEulerEvc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEvc{EulerEvcCaller: EulerEvcCaller{contract: contract}, EulerEvcTransactor: EulerEvcTransactor{contract: contract}, EulerEvcFilterer: EulerEvcFilterer{contract: contract}}, nil
}

// NewEulerEvcCaller creates a new read-only instance of EulerEvc, bound to a specific deployed contract.
func NewEulerEvcCaller(address common.Address, caller bind.ContractCaller) (*EulerEvcCaller, error) {
	contract, err := bindEulerEvc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvcCaller{contract: contract}, nil
}

// NewEulerEvcTransactor creates a new write-only instance of EulerEvc, bound to a specific deployed contract.
func NewEulerEvcTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEvcTransactor, error) {
	contract, err := bindEulerEvc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvcTransactor{contract: contract}, nil
}

// NewEulerEvcFilterer creates a new log filterer instance of EulerEvc, bound to a specific deployed contract.
func NewEulerEvcFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEvcFilterer, error) {
	contract, err := bindEulerEvc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEvcFilterer{contract: contract}, nil
}

// bindEulerEvc binds a generic wrapper to an already deployed contract.
func bindEulerEvc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEvcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvc *EulerEvcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvc.Contract.EulerEvcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvc *EulerEvcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvc.Contract.EulerEvcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvc *EulerEvcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvc.Contract.EulerEvcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvc *EulerEvcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvc *EulerEvcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvc *EulerEvcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvc.Contract.contract.Transact(opts, method, params...)
}

// AreChecksDeferred is a free data retrieval call binding the contract method 0x430292b3.
//
// Solidity: function areChecksDeferred() view returns(bool)
func (_EulerEvc *EulerEvcCaller) AreChecksDeferred(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "areChecksDeferred")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreChecksDeferred is a free data retrieval call binding the contract method 0x430292b3.
//
// Solidity: function areChecksDeferred() view returns(bool)
func (_EulerEvc *EulerEvcSession) AreChecksDeferred() (bool, error) {
	return _EulerEvc.Contract.AreChecksDeferred(&_EulerEvc.CallOpts)
}

// AreChecksDeferred is a free data retrieval call binding the contract method 0x430292b3.
//
// Solidity: function areChecksDeferred() view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) AreChecksDeferred() (bool, error) {
	return _EulerEvc.Contract.AreChecksDeferred(&_EulerEvc.CallOpts)
}

// AreChecksInProgress is a free data retrieval call binding the contract method 0xe21e537c.
//
// Solidity: function areChecksInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCaller) AreChecksInProgress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "areChecksInProgress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreChecksInProgress is a free data retrieval call binding the contract method 0xe21e537c.
//
// Solidity: function areChecksInProgress() view returns(bool)
func (_EulerEvc *EulerEvcSession) AreChecksInProgress() (bool, error) {
	return _EulerEvc.Contract.AreChecksInProgress(&_EulerEvc.CallOpts)
}

// AreChecksInProgress is a free data retrieval call binding the contract method 0xe21e537c.
//
// Solidity: function areChecksInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) AreChecksInProgress() (bool, error) {
	return _EulerEvc.Contract.AreChecksInProgress(&_EulerEvc.CallOpts)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0x442b172c.
//
// Solidity: function getAccountOwner(address account) view returns(address)
func (_EulerEvc *EulerEvcCaller) GetAccountOwner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getAccountOwner", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0x442b172c.
//
// Solidity: function getAccountOwner(address account) view returns(address)
func (_EulerEvc *EulerEvcSession) GetAccountOwner(account common.Address) (common.Address, error) {
	return _EulerEvc.Contract.GetAccountOwner(&_EulerEvc.CallOpts, account)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0x442b172c.
//
// Solidity: function getAccountOwner(address account) view returns(address)
func (_EulerEvc *EulerEvcCallerSession) GetAccountOwner(account common.Address) (common.Address, error) {
	return _EulerEvc.Contract.GetAccountOwner(&_EulerEvc.CallOpts, account)
}

// GetAddressPrefix is a free data retrieval call binding the contract method 0x506d8c92.
//
// Solidity: function getAddressPrefix(address account) pure returns(bytes19)
func (_EulerEvc *EulerEvcCaller) GetAddressPrefix(opts *bind.CallOpts, account common.Address) ([19]byte, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getAddressPrefix", account)

	if err != nil {
		return *new([19]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([19]byte)).(*[19]byte)

	return out0, err

}

// GetAddressPrefix is a free data retrieval call binding the contract method 0x506d8c92.
//
// Solidity: function getAddressPrefix(address account) pure returns(bytes19)
func (_EulerEvc *EulerEvcSession) GetAddressPrefix(account common.Address) ([19]byte, error) {
	return _EulerEvc.Contract.GetAddressPrefix(&_EulerEvc.CallOpts, account)
}

// GetAddressPrefix is a free data retrieval call binding the contract method 0x506d8c92.
//
// Solidity: function getAddressPrefix(address account) pure returns(bytes19)
func (_EulerEvc *EulerEvcCallerSession) GetAddressPrefix(account common.Address) ([19]byte, error) {
	return _EulerEvc.Contract.GetAddressPrefix(&_EulerEvc.CallOpts, account)
}

// GetCollaterals is a free data retrieval call binding the contract method 0xa4d25d1e.
//
// Solidity: function getCollaterals(address account) view returns(address[])
func (_EulerEvc *EulerEvcCaller) GetCollaterals(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getCollaterals", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollaterals is a free data retrieval call binding the contract method 0xa4d25d1e.
//
// Solidity: function getCollaterals(address account) view returns(address[])
func (_EulerEvc *EulerEvcSession) GetCollaterals(account common.Address) ([]common.Address, error) {
	return _EulerEvc.Contract.GetCollaterals(&_EulerEvc.CallOpts, account)
}

// GetCollaterals is a free data retrieval call binding the contract method 0xa4d25d1e.
//
// Solidity: function getCollaterals(address account) view returns(address[])
func (_EulerEvc *EulerEvcCallerSession) GetCollaterals(account common.Address) ([]common.Address, error) {
	return _EulerEvc.Contract.GetCollaterals(&_EulerEvc.CallOpts, account)
}

// GetControllers is a free data retrieval call binding the contract method 0xfd6046d7.
//
// Solidity: function getControllers(address account) view returns(address[])
func (_EulerEvc *EulerEvcCaller) GetControllers(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getControllers", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetControllers is a free data retrieval call binding the contract method 0xfd6046d7.
//
// Solidity: function getControllers(address account) view returns(address[])
func (_EulerEvc *EulerEvcSession) GetControllers(account common.Address) ([]common.Address, error) {
	return _EulerEvc.Contract.GetControllers(&_EulerEvc.CallOpts, account)
}

// GetControllers is a free data retrieval call binding the contract method 0xfd6046d7.
//
// Solidity: function getControllers(address account) view returns(address[])
func (_EulerEvc *EulerEvcCallerSession) GetControllers(account common.Address) ([]common.Address, error) {
	return _EulerEvc.Contract.GetControllers(&_EulerEvc.CallOpts, account)
}

// GetCurrentOnBehalfOfAccount is a free data retrieval call binding the contract method 0x18503a1e.
//
// Solidity: function getCurrentOnBehalfOfAccount(address controllerToCheck) view returns(address onBehalfOfAccount, bool controllerEnabled)
func (_EulerEvc *EulerEvcCaller) GetCurrentOnBehalfOfAccount(opts *bind.CallOpts, controllerToCheck common.Address) (struct {
	OnBehalfOfAccount common.Address
	ControllerEnabled bool
}, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getCurrentOnBehalfOfAccount", controllerToCheck)

	outstruct := new(struct {
		OnBehalfOfAccount common.Address
		ControllerEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OnBehalfOfAccount = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ControllerEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetCurrentOnBehalfOfAccount is a free data retrieval call binding the contract method 0x18503a1e.
//
// Solidity: function getCurrentOnBehalfOfAccount(address controllerToCheck) view returns(address onBehalfOfAccount, bool controllerEnabled)
func (_EulerEvc *EulerEvcSession) GetCurrentOnBehalfOfAccount(controllerToCheck common.Address) (struct {
	OnBehalfOfAccount common.Address
	ControllerEnabled bool
}, error) {
	return _EulerEvc.Contract.GetCurrentOnBehalfOfAccount(&_EulerEvc.CallOpts, controllerToCheck)
}

// GetCurrentOnBehalfOfAccount is a free data retrieval call binding the contract method 0x18503a1e.
//
// Solidity: function getCurrentOnBehalfOfAccount(address controllerToCheck) view returns(address onBehalfOfAccount, bool controllerEnabled)
func (_EulerEvc *EulerEvcCallerSession) GetCurrentOnBehalfOfAccount(controllerToCheck common.Address) (struct {
	OnBehalfOfAccount common.Address
	ControllerEnabled bool
}, error) {
	return _EulerEvc.Contract.GetCurrentOnBehalfOfAccount(&_EulerEvc.CallOpts, controllerToCheck)
}

// GetLastAccountStatusCheckTimestamp is a free data retrieval call binding the contract method 0xdf7c1384.
//
// Solidity: function getLastAccountStatusCheckTimestamp(address account) view returns(uint256)
func (_EulerEvc *EulerEvcCaller) GetLastAccountStatusCheckTimestamp(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getLastAccountStatusCheckTimestamp", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastAccountStatusCheckTimestamp is a free data retrieval call binding the contract method 0xdf7c1384.
//
// Solidity: function getLastAccountStatusCheckTimestamp(address account) view returns(uint256)
func (_EulerEvc *EulerEvcSession) GetLastAccountStatusCheckTimestamp(account common.Address) (*big.Int, error) {
	return _EulerEvc.Contract.GetLastAccountStatusCheckTimestamp(&_EulerEvc.CallOpts, account)
}

// GetLastAccountStatusCheckTimestamp is a free data retrieval call binding the contract method 0xdf7c1384.
//
// Solidity: function getLastAccountStatusCheckTimestamp(address account) view returns(uint256)
func (_EulerEvc *EulerEvcCallerSession) GetLastAccountStatusCheckTimestamp(account common.Address) (*big.Int, error) {
	return _EulerEvc.Contract.GetLastAccountStatusCheckTimestamp(&_EulerEvc.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x12d6c936.
//
// Solidity: function getNonce(bytes19 addressPrefix, uint256 nonceNamespace) view returns(uint256)
func (_EulerEvc *EulerEvcCaller) GetNonce(opts *bind.CallOpts, addressPrefix [19]byte, nonceNamespace *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getNonce", addressPrefix, nonceNamespace)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x12d6c936.
//
// Solidity: function getNonce(bytes19 addressPrefix, uint256 nonceNamespace) view returns(uint256)
func (_EulerEvc *EulerEvcSession) GetNonce(addressPrefix [19]byte, nonceNamespace *big.Int) (*big.Int, error) {
	return _EulerEvc.Contract.GetNonce(&_EulerEvc.CallOpts, addressPrefix, nonceNamespace)
}

// GetNonce is a free data retrieval call binding the contract method 0x12d6c936.
//
// Solidity: function getNonce(bytes19 addressPrefix, uint256 nonceNamespace) view returns(uint256)
func (_EulerEvc *EulerEvcCallerSession) GetNonce(addressPrefix [19]byte, nonceNamespace *big.Int) (*big.Int, error) {
	return _EulerEvc.Contract.GetNonce(&_EulerEvc.CallOpts, addressPrefix, nonceNamespace)
}

// GetOperator is a free data retrieval call binding the contract method 0xb03c130d.
//
// Solidity: function getOperator(bytes19 addressPrefix, address operator) view returns(uint256)
func (_EulerEvc *EulerEvcCaller) GetOperator(opts *bind.CallOpts, addressPrefix [19]byte, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getOperator", addressPrefix, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xb03c130d.
//
// Solidity: function getOperator(bytes19 addressPrefix, address operator) view returns(uint256)
func (_EulerEvc *EulerEvcSession) GetOperator(addressPrefix [19]byte, operator common.Address) (*big.Int, error) {
	return _EulerEvc.Contract.GetOperator(&_EulerEvc.CallOpts, addressPrefix, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0xb03c130d.
//
// Solidity: function getOperator(bytes19 addressPrefix, address operator) view returns(uint256)
func (_EulerEvc *EulerEvcCallerSession) GetOperator(addressPrefix [19]byte, operator common.Address) (*big.Int, error) {
	return _EulerEvc.Contract.GetOperator(&_EulerEvc.CallOpts, addressPrefix, operator)
}

// GetRawExecutionContext is a free data retrieval call binding the contract method 0x3a1a3a1d.
//
// Solidity: function getRawExecutionContext() view returns(uint256 context)
func (_EulerEvc *EulerEvcCaller) GetRawExecutionContext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "getRawExecutionContext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRawExecutionContext is a free data retrieval call binding the contract method 0x3a1a3a1d.
//
// Solidity: function getRawExecutionContext() view returns(uint256 context)
func (_EulerEvc *EulerEvcSession) GetRawExecutionContext() (*big.Int, error) {
	return _EulerEvc.Contract.GetRawExecutionContext(&_EulerEvc.CallOpts)
}

// GetRawExecutionContext is a free data retrieval call binding the contract method 0x3a1a3a1d.
//
// Solidity: function getRawExecutionContext() view returns(uint256 context)
func (_EulerEvc *EulerEvcCallerSession) GetRawExecutionContext() (*big.Int, error) {
	return _EulerEvc.Contract.GetRawExecutionContext(&_EulerEvc.CallOpts)
}

// HaveCommonOwner is a free data retrieval call binding the contract method 0xc760d921.
//
// Solidity: function haveCommonOwner(address account, address otherAccount) pure returns(bool)
func (_EulerEvc *EulerEvcCaller) HaveCommonOwner(opts *bind.CallOpts, account common.Address, otherAccount common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "haveCommonOwner", account, otherAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HaveCommonOwner is a free data retrieval call binding the contract method 0xc760d921.
//
// Solidity: function haveCommonOwner(address account, address otherAccount) pure returns(bool)
func (_EulerEvc *EulerEvcSession) HaveCommonOwner(account common.Address, otherAccount common.Address) (bool, error) {
	return _EulerEvc.Contract.HaveCommonOwner(&_EulerEvc.CallOpts, account, otherAccount)
}

// HaveCommonOwner is a free data retrieval call binding the contract method 0xc760d921.
//
// Solidity: function haveCommonOwner(address account, address otherAccount) pure returns(bool)
func (_EulerEvc *EulerEvcCallerSession) HaveCommonOwner(account common.Address, otherAccount common.Address) (bool, error) {
	return _EulerEvc.Contract.HaveCommonOwner(&_EulerEvc.CallOpts, account, otherAccount)
}

// IsAccountOperatorAuthorized is a free data retrieval call binding the contract method 0x1647292a.
//
// Solidity: function isAccountOperatorAuthorized(address account, address operator) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsAccountOperatorAuthorized(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isAccountOperatorAuthorized", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountOperatorAuthorized is a free data retrieval call binding the contract method 0x1647292a.
//
// Solidity: function isAccountOperatorAuthorized(address account, address operator) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsAccountOperatorAuthorized(account common.Address, operator common.Address) (bool, error) {
	return _EulerEvc.Contract.IsAccountOperatorAuthorized(&_EulerEvc.CallOpts, account, operator)
}

// IsAccountOperatorAuthorized is a free data retrieval call binding the contract method 0x1647292a.
//
// Solidity: function isAccountOperatorAuthorized(address account, address operator) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsAccountOperatorAuthorized(account common.Address, operator common.Address) (bool, error) {
	return _EulerEvc.Contract.IsAccountOperatorAuthorized(&_EulerEvc.CallOpts, account, operator)
}

// IsAccountStatusCheckDeferred is a free data retrieval call binding the contract method 0x42e53499.
//
// Solidity: function isAccountStatusCheckDeferred(address account) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsAccountStatusCheckDeferred(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isAccountStatusCheckDeferred", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountStatusCheckDeferred is a free data retrieval call binding the contract method 0x42e53499.
//
// Solidity: function isAccountStatusCheckDeferred(address account) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsAccountStatusCheckDeferred(account common.Address) (bool, error) {
	return _EulerEvc.Contract.IsAccountStatusCheckDeferred(&_EulerEvc.CallOpts, account)
}

// IsAccountStatusCheckDeferred is a free data retrieval call binding the contract method 0x42e53499.
//
// Solidity: function isAccountStatusCheckDeferred(address account) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsAccountStatusCheckDeferred(account common.Address) (bool, error) {
	return _EulerEvc.Contract.IsAccountStatusCheckDeferred(&_EulerEvc.CallOpts, account)
}

// IsCollateralEnabled is a free data retrieval call binding the contract method 0x9e716d58.
//
// Solidity: function isCollateralEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsCollateralEnabled(opts *bind.CallOpts, account common.Address, vault common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isCollateralEnabled", account, vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralEnabled is a free data retrieval call binding the contract method 0x9e716d58.
//
// Solidity: function isCollateralEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsCollateralEnabled(account common.Address, vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsCollateralEnabled(&_EulerEvc.CallOpts, account, vault)
}

// IsCollateralEnabled is a free data retrieval call binding the contract method 0x9e716d58.
//
// Solidity: function isCollateralEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsCollateralEnabled(account common.Address, vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsCollateralEnabled(&_EulerEvc.CallOpts, account, vault)
}

// IsControlCollateralInProgress is a free data retrieval call binding the contract method 0x863789d7.
//
// Solidity: function isControlCollateralInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsControlCollateralInProgress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isControlCollateralInProgress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsControlCollateralInProgress is a free data retrieval call binding the contract method 0x863789d7.
//
// Solidity: function isControlCollateralInProgress() view returns(bool)
func (_EulerEvc *EulerEvcSession) IsControlCollateralInProgress() (bool, error) {
	return _EulerEvc.Contract.IsControlCollateralInProgress(&_EulerEvc.CallOpts)
}

// IsControlCollateralInProgress is a free data retrieval call binding the contract method 0x863789d7.
//
// Solidity: function isControlCollateralInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsControlCollateralInProgress() (bool, error) {
	return _EulerEvc.Contract.IsControlCollateralInProgress(&_EulerEvc.CallOpts)
}

// IsControllerEnabled is a free data retrieval call binding the contract method 0x47cfdac4.
//
// Solidity: function isControllerEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsControllerEnabled(opts *bind.CallOpts, account common.Address, vault common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isControllerEnabled", account, vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsControllerEnabled is a free data retrieval call binding the contract method 0x47cfdac4.
//
// Solidity: function isControllerEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsControllerEnabled(account common.Address, vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsControllerEnabled(&_EulerEvc.CallOpts, account, vault)
}

// IsControllerEnabled is a free data retrieval call binding the contract method 0x47cfdac4.
//
// Solidity: function isControllerEnabled(address account, address vault) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsControllerEnabled(account common.Address, vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsControllerEnabled(&_EulerEvc.CallOpts, account, vault)
}

// IsLockdownMode is a free data retrieval call binding the contract method 0x3b10f3ef.
//
// Solidity: function isLockdownMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsLockdownMode(opts *bind.CallOpts, addressPrefix [19]byte) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isLockdownMode", addressPrefix)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockdownMode is a free data retrieval call binding the contract method 0x3b10f3ef.
//
// Solidity: function isLockdownMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsLockdownMode(addressPrefix [19]byte) (bool, error) {
	return _EulerEvc.Contract.IsLockdownMode(&_EulerEvc.CallOpts, addressPrefix)
}

// IsLockdownMode is a free data retrieval call binding the contract method 0x3b10f3ef.
//
// Solidity: function isLockdownMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsLockdownMode(addressPrefix [19]byte) (bool, error) {
	return _EulerEvc.Contract.IsLockdownMode(&_EulerEvc.CallOpts, addressPrefix)
}

// IsOperatorAuthenticated is a free data retrieval call binding the contract method 0x3b2416be.
//
// Solidity: function isOperatorAuthenticated() view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsOperatorAuthenticated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isOperatorAuthenticated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAuthenticated is a free data retrieval call binding the contract method 0x3b2416be.
//
// Solidity: function isOperatorAuthenticated() view returns(bool)
func (_EulerEvc *EulerEvcSession) IsOperatorAuthenticated() (bool, error) {
	return _EulerEvc.Contract.IsOperatorAuthenticated(&_EulerEvc.CallOpts)
}

// IsOperatorAuthenticated is a free data retrieval call binding the contract method 0x3b2416be.
//
// Solidity: function isOperatorAuthenticated() view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsOperatorAuthenticated() (bool, error) {
	return _EulerEvc.Contract.IsOperatorAuthenticated(&_EulerEvc.CallOpts)
}

// IsPermitDisabledMode is a free data retrieval call binding the contract method 0xcb29955a.
//
// Solidity: function isPermitDisabledMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsPermitDisabledMode(opts *bind.CallOpts, addressPrefix [19]byte) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isPermitDisabledMode", addressPrefix)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPermitDisabledMode is a free data retrieval call binding the contract method 0xcb29955a.
//
// Solidity: function isPermitDisabledMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsPermitDisabledMode(addressPrefix [19]byte) (bool, error) {
	return _EulerEvc.Contract.IsPermitDisabledMode(&_EulerEvc.CallOpts, addressPrefix)
}

// IsPermitDisabledMode is a free data retrieval call binding the contract method 0xcb29955a.
//
// Solidity: function isPermitDisabledMode(bytes19 addressPrefix) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsPermitDisabledMode(addressPrefix [19]byte) (bool, error) {
	return _EulerEvc.Contract.IsPermitDisabledMode(&_EulerEvc.CallOpts, addressPrefix)
}

// IsSimulationInProgress is a free data retrieval call binding the contract method 0x92d2fc01.
//
// Solidity: function isSimulationInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsSimulationInProgress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isSimulationInProgress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSimulationInProgress is a free data retrieval call binding the contract method 0x92d2fc01.
//
// Solidity: function isSimulationInProgress() view returns(bool)
func (_EulerEvc *EulerEvcSession) IsSimulationInProgress() (bool, error) {
	return _EulerEvc.Contract.IsSimulationInProgress(&_EulerEvc.CallOpts)
}

// IsSimulationInProgress is a free data retrieval call binding the contract method 0x92d2fc01.
//
// Solidity: function isSimulationInProgress() view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsSimulationInProgress() (bool, error) {
	return _EulerEvc.Contract.IsSimulationInProgress(&_EulerEvc.CallOpts)
}

// IsVaultStatusCheckDeferred is a free data retrieval call binding the contract method 0xcdd8ea78.
//
// Solidity: function isVaultStatusCheckDeferred(address vault) view returns(bool)
func (_EulerEvc *EulerEvcCaller) IsVaultStatusCheckDeferred(opts *bind.CallOpts, vault common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "isVaultStatusCheckDeferred", vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultStatusCheckDeferred is a free data retrieval call binding the contract method 0xcdd8ea78.
//
// Solidity: function isVaultStatusCheckDeferred(address vault) view returns(bool)
func (_EulerEvc *EulerEvcSession) IsVaultStatusCheckDeferred(vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsVaultStatusCheckDeferred(&_EulerEvc.CallOpts, vault)
}

// IsVaultStatusCheckDeferred is a free data retrieval call binding the contract method 0xcdd8ea78.
//
// Solidity: function isVaultStatusCheckDeferred(address vault) view returns(bool)
func (_EulerEvc *EulerEvcCallerSession) IsVaultStatusCheckDeferred(vault common.Address) (bool, error) {
	return _EulerEvc.Contract.IsVaultStatusCheckDeferred(&_EulerEvc.CallOpts, vault)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvc *EulerEvcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEvc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvc *EulerEvcSession) Name() (string, error) {
	return _EulerEvc.Contract.Name(&_EulerEvc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvc *EulerEvcCallerSession) Name() (string, error) {
	return _EulerEvc.Contract.Name(&_EulerEvc.CallOpts)
}

// Batch is a paid mutator transaction binding the contract method 0xc16ae7a4.
//
// Solidity: function batch((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcTransactor) Batch(opts *bind.TransactOpts, items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "batch", items)
}

// Batch is a paid mutator transaction binding the contract method 0xc16ae7a4.
//
// Solidity: function batch((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcSession) Batch(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.Batch(&_EulerEvc.TransactOpts, items)
}

// Batch is a paid mutator transaction binding the contract method 0xc16ae7a4.
//
// Solidity: function batch((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) Batch(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.Batch(&_EulerEvc.TransactOpts, items)
}

// BatchRevert is a paid mutator transaction binding the contract method 0x7f5c92f3.
//
// Solidity: function batchRevert((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcTransactor) BatchRevert(opts *bind.TransactOpts, items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "batchRevert", items)
}

// BatchRevert is a paid mutator transaction binding the contract method 0x7f5c92f3.
//
// Solidity: function batchRevert((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcSession) BatchRevert(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.BatchRevert(&_EulerEvc.TransactOpts, items)
}

// BatchRevert is a paid mutator transaction binding the contract method 0x7f5c92f3.
//
// Solidity: function batchRevert((address,address,uint256,bytes)[] items) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) BatchRevert(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.BatchRevert(&_EulerEvc.TransactOpts, items)
}

// BatchSimulation is a paid mutator transaction binding the contract method 0x7f17c377.
//
// Solidity: function batchSimulation((address,address,uint256,bytes)[] items) payable returns((bool,bytes)[] batchItemsResult, (address,bool,bytes)[] accountsStatusCheckResult, (address,bool,bytes)[] vaultsStatusCheckResult)
func (_EulerEvc *EulerEvcTransactor) BatchSimulation(opts *bind.TransactOpts, items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "batchSimulation", items)
}

// BatchSimulation is a paid mutator transaction binding the contract method 0x7f17c377.
//
// Solidity: function batchSimulation((address,address,uint256,bytes)[] items) payable returns((bool,bytes)[] batchItemsResult, (address,bool,bytes)[] accountsStatusCheckResult, (address,bool,bytes)[] vaultsStatusCheckResult)
func (_EulerEvc *EulerEvcSession) BatchSimulation(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.BatchSimulation(&_EulerEvc.TransactOpts, items)
}

// BatchSimulation is a paid mutator transaction binding the contract method 0x7f17c377.
//
// Solidity: function batchSimulation((address,address,uint256,bytes)[] items) payable returns((bool,bytes)[] batchItemsResult, (address,bool,bytes)[] accountsStatusCheckResult, (address,bool,bytes)[] vaultsStatusCheckResult)
func (_EulerEvc *EulerEvcTransactorSession) BatchSimulation(items []IEVCBatchItem) (*types.Transaction, error) {
	return _EulerEvc.Contract.BatchSimulation(&_EulerEvc.TransactOpts, items)
}

// Call is a paid mutator transaction binding the contract method 0x1f8b5215.
//
// Solidity: function call(address targetContract, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcTransactor) Call(opts *bind.TransactOpts, targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "call", targetContract, onBehalfOfAccount, value, data)
}

// Call is a paid mutator transaction binding the contract method 0x1f8b5215.
//
// Solidity: function call(address targetContract, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcSession) Call(targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.Call(&_EulerEvc.TransactOpts, targetContract, onBehalfOfAccount, value, data)
}

// Call is a paid mutator transaction binding the contract method 0x1f8b5215.
//
// Solidity: function call(address targetContract, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcTransactorSession) Call(targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.Call(&_EulerEvc.TransactOpts, targetContract, onBehalfOfAccount, value, data)
}

// ControlCollateral is a paid mutator transaction binding the contract method 0xb9b70ff5.
//
// Solidity: function controlCollateral(address targetCollateral, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcTransactor) ControlCollateral(opts *bind.TransactOpts, targetCollateral common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "controlCollateral", targetCollateral, onBehalfOfAccount, value, data)
}

// ControlCollateral is a paid mutator transaction binding the contract method 0xb9b70ff5.
//
// Solidity: function controlCollateral(address targetCollateral, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcSession) ControlCollateral(targetCollateral common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.ControlCollateral(&_EulerEvc.TransactOpts, targetCollateral, onBehalfOfAccount, value, data)
}

// ControlCollateral is a paid mutator transaction binding the contract method 0xb9b70ff5.
//
// Solidity: function controlCollateral(address targetCollateral, address onBehalfOfAccount, uint256 value, bytes data) payable returns(bytes result)
func (_EulerEvc *EulerEvcTransactorSession) ControlCollateral(targetCollateral common.Address, onBehalfOfAccount common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.ControlCollateral(&_EulerEvc.TransactOpts, targetCollateral, onBehalfOfAccount, value, data)
}

// DisableCollateral is a paid mutator transaction binding the contract method 0xe920e8e0.
//
// Solidity: function disableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactor) DisableCollateral(opts *bind.TransactOpts, account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "disableCollateral", account, vault)
}

// DisableCollateral is a paid mutator transaction binding the contract method 0xe920e8e0.
//
// Solidity: function disableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcSession) DisableCollateral(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.DisableCollateral(&_EulerEvc.TransactOpts, account, vault)
}

// DisableCollateral is a paid mutator transaction binding the contract method 0xe920e8e0.
//
// Solidity: function disableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) DisableCollateral(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.DisableCollateral(&_EulerEvc.TransactOpts, account, vault)
}

// DisableController is a paid mutator transaction binding the contract method 0xf4fc3570.
//
// Solidity: function disableController(address account) payable returns()
func (_EulerEvc *EulerEvcTransactor) DisableController(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "disableController", account)
}

// DisableController is a paid mutator transaction binding the contract method 0xf4fc3570.
//
// Solidity: function disableController(address account) payable returns()
func (_EulerEvc *EulerEvcSession) DisableController(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.DisableController(&_EulerEvc.TransactOpts, account)
}

// DisableController is a paid mutator transaction binding the contract method 0xf4fc3570.
//
// Solidity: function disableController(address account) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) DisableController(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.DisableController(&_EulerEvc.TransactOpts, account)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0xd44fee5a.
//
// Solidity: function enableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactor) EnableCollateral(opts *bind.TransactOpts, account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "enableCollateral", account, vault)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0xd44fee5a.
//
// Solidity: function enableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcSession) EnableCollateral(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.EnableCollateral(&_EulerEvc.TransactOpts, account, vault)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0xd44fee5a.
//
// Solidity: function enableCollateral(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) EnableCollateral(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.EnableCollateral(&_EulerEvc.TransactOpts, account, vault)
}

// EnableController is a paid mutator transaction binding the contract method 0xc368516c.
//
// Solidity: function enableController(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactor) EnableController(opts *bind.TransactOpts, account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "enableController", account, vault)
}

// EnableController is a paid mutator transaction binding the contract method 0xc368516c.
//
// Solidity: function enableController(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcSession) EnableController(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.EnableController(&_EulerEvc.TransactOpts, account, vault)
}

// EnableController is a paid mutator transaction binding the contract method 0xc368516c.
//
// Solidity: function enableController(address account, address vault) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) EnableController(account common.Address, vault common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.EnableController(&_EulerEvc.TransactOpts, account, vault)
}

// ForgiveAccountStatusCheck is a paid mutator transaction binding the contract method 0x10a75198.
//
// Solidity: function forgiveAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactor) ForgiveAccountStatusCheck(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "forgiveAccountStatusCheck", account)
}

// ForgiveAccountStatusCheck is a paid mutator transaction binding the contract method 0x10a75198.
//
// Solidity: function forgiveAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcSession) ForgiveAccountStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.ForgiveAccountStatusCheck(&_EulerEvc.TransactOpts, account)
}

// ForgiveAccountStatusCheck is a paid mutator transaction binding the contract method 0x10a75198.
//
// Solidity: function forgiveAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) ForgiveAccountStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.ForgiveAccountStatusCheck(&_EulerEvc.TransactOpts, account)
}

// ForgiveVaultStatusCheck is a paid mutator transaction binding the contract method 0xebf1ea86.
//
// Solidity: function forgiveVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcTransactor) ForgiveVaultStatusCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "forgiveVaultStatusCheck")
}

// ForgiveVaultStatusCheck is a paid mutator transaction binding the contract method 0xebf1ea86.
//
// Solidity: function forgiveVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcSession) ForgiveVaultStatusCheck() (*types.Transaction, error) {
	return _EulerEvc.Contract.ForgiveVaultStatusCheck(&_EulerEvc.TransactOpts)
}

// ForgiveVaultStatusCheck is a paid mutator transaction binding the contract method 0xebf1ea86.
//
// Solidity: function forgiveVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcTransactorSession) ForgiveVaultStatusCheck() (*types.Transaction, error) {
	return _EulerEvc.Contract.ForgiveVaultStatusCheck(&_EulerEvc.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0x5bedd1cd.
//
// Solidity: function permit(address signer, address sender, uint256 nonceNamespace, uint256 nonce, uint256 deadline, uint256 value, bytes data, bytes signature) payable returns()
func (_EulerEvc *EulerEvcTransactor) Permit(opts *bind.TransactOpts, signer common.Address, sender common.Address, nonceNamespace *big.Int, nonce *big.Int, deadline *big.Int, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "permit", signer, sender, nonceNamespace, nonce, deadline, value, data, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x5bedd1cd.
//
// Solidity: function permit(address signer, address sender, uint256 nonceNamespace, uint256 nonce, uint256 deadline, uint256 value, bytes data, bytes signature) payable returns()
func (_EulerEvc *EulerEvcSession) Permit(signer common.Address, sender common.Address, nonceNamespace *big.Int, nonce *big.Int, deadline *big.Int, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.Permit(&_EulerEvc.TransactOpts, signer, sender, nonceNamespace, nonce, deadline, value, data, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x5bedd1cd.
//
// Solidity: function permit(address signer, address sender, uint256 nonceNamespace, uint256 nonce, uint256 deadline, uint256 value, bytes data, bytes signature) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) Permit(signer common.Address, sender common.Address, nonceNamespace *big.Int, nonce *big.Int, deadline *big.Int, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _EulerEvc.Contract.Permit(&_EulerEvc.TransactOpts, signer, sender, nonceNamespace, nonce, deadline, value, data, signature)
}

// ReorderCollaterals is a paid mutator transaction binding the contract method 0x642ea23f.
//
// Solidity: function reorderCollaterals(address account, uint8 index1, uint8 index2) payable returns()
func (_EulerEvc *EulerEvcTransactor) ReorderCollaterals(opts *bind.TransactOpts, account common.Address, index1 uint8, index2 uint8) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "reorderCollaterals", account, index1, index2)
}

// ReorderCollaterals is a paid mutator transaction binding the contract method 0x642ea23f.
//
// Solidity: function reorderCollaterals(address account, uint8 index1, uint8 index2) payable returns()
func (_EulerEvc *EulerEvcSession) ReorderCollaterals(account common.Address, index1 uint8, index2 uint8) (*types.Transaction, error) {
	return _EulerEvc.Contract.ReorderCollaterals(&_EulerEvc.TransactOpts, account, index1, index2)
}

// ReorderCollaterals is a paid mutator transaction binding the contract method 0x642ea23f.
//
// Solidity: function reorderCollaterals(address account, uint8 index1, uint8 index2) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) ReorderCollaterals(account common.Address, index1 uint8, index2 uint8) (*types.Transaction, error) {
	return _EulerEvc.Contract.ReorderCollaterals(&_EulerEvc.TransactOpts, account, index1, index2)
}

// RequireAccountAndVaultStatusCheck is a paid mutator transaction binding the contract method 0x30f31667.
//
// Solidity: function requireAccountAndVaultStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactor) RequireAccountAndVaultStatusCheck(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "requireAccountAndVaultStatusCheck", account)
}

// RequireAccountAndVaultStatusCheck is a paid mutator transaction binding the contract method 0x30f31667.
//
// Solidity: function requireAccountAndVaultStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcSession) RequireAccountAndVaultStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireAccountAndVaultStatusCheck(&_EulerEvc.TransactOpts, account)
}

// RequireAccountAndVaultStatusCheck is a paid mutator transaction binding the contract method 0x30f31667.
//
// Solidity: function requireAccountAndVaultStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) RequireAccountAndVaultStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireAccountAndVaultStatusCheck(&_EulerEvc.TransactOpts, account)
}

// RequireAccountStatusCheck is a paid mutator transaction binding the contract method 0x46591032.
//
// Solidity: function requireAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactor) RequireAccountStatusCheck(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "requireAccountStatusCheck", account)
}

// RequireAccountStatusCheck is a paid mutator transaction binding the contract method 0x46591032.
//
// Solidity: function requireAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcSession) RequireAccountStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireAccountStatusCheck(&_EulerEvc.TransactOpts, account)
}

// RequireAccountStatusCheck is a paid mutator transaction binding the contract method 0x46591032.
//
// Solidity: function requireAccountStatusCheck(address account) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) RequireAccountStatusCheck(account common.Address) (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireAccountStatusCheck(&_EulerEvc.TransactOpts, account)
}

// RequireVaultStatusCheck is a paid mutator transaction binding the contract method 0xa37d54af.
//
// Solidity: function requireVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcTransactor) RequireVaultStatusCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "requireVaultStatusCheck")
}

// RequireVaultStatusCheck is a paid mutator transaction binding the contract method 0xa37d54af.
//
// Solidity: function requireVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcSession) RequireVaultStatusCheck() (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireVaultStatusCheck(&_EulerEvc.TransactOpts)
}

// RequireVaultStatusCheck is a paid mutator transaction binding the contract method 0xa37d54af.
//
// Solidity: function requireVaultStatusCheck() payable returns()
func (_EulerEvc *EulerEvcTransactorSession) RequireVaultStatusCheck() (*types.Transaction, error) {
	return _EulerEvc.Contract.RequireVaultStatusCheck(&_EulerEvc.TransactOpts)
}

// SetAccountOperator is a paid mutator transaction binding the contract method 0x9f5c462a.
//
// Solidity: function setAccountOperator(address account, address operator, bool authorized) payable returns()
func (_EulerEvc *EulerEvcTransactor) SetAccountOperator(opts *bind.TransactOpts, account common.Address, operator common.Address, authorized bool) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "setAccountOperator", account, operator, authorized)
}

// SetAccountOperator is a paid mutator transaction binding the contract method 0x9f5c462a.
//
// Solidity: function setAccountOperator(address account, address operator, bool authorized) payable returns()
func (_EulerEvc *EulerEvcSession) SetAccountOperator(account common.Address, operator common.Address, authorized bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetAccountOperator(&_EulerEvc.TransactOpts, account, operator, authorized)
}

// SetAccountOperator is a paid mutator transaction binding the contract method 0x9f5c462a.
//
// Solidity: function setAccountOperator(address account, address operator, bool authorized) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) SetAccountOperator(account common.Address, operator common.Address, authorized bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetAccountOperator(&_EulerEvc.TransactOpts, account, operator, authorized)
}

// SetLockdownMode is a paid mutator transaction binding the contract method 0x129d21a0.
//
// Solidity: function setLockdownMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcTransactor) SetLockdownMode(opts *bind.TransactOpts, addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "setLockdownMode", addressPrefix, enabled)
}

// SetLockdownMode is a paid mutator transaction binding the contract method 0x129d21a0.
//
// Solidity: function setLockdownMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcSession) SetLockdownMode(addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetLockdownMode(&_EulerEvc.TransactOpts, addressPrefix, enabled)
}

// SetLockdownMode is a paid mutator transaction binding the contract method 0x129d21a0.
//
// Solidity: function setLockdownMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) SetLockdownMode(addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetLockdownMode(&_EulerEvc.TransactOpts, addressPrefix, enabled)
}

// SetNonce is a paid mutator transaction binding the contract method 0xa829aaf5.
//
// Solidity: function setNonce(bytes19 addressPrefix, uint256 nonceNamespace, uint256 nonce) payable returns()
func (_EulerEvc *EulerEvcTransactor) SetNonce(opts *bind.TransactOpts, addressPrefix [19]byte, nonceNamespace *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "setNonce", addressPrefix, nonceNamespace, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xa829aaf5.
//
// Solidity: function setNonce(bytes19 addressPrefix, uint256 nonceNamespace, uint256 nonce) payable returns()
func (_EulerEvc *EulerEvcSession) SetNonce(addressPrefix [19]byte, nonceNamespace *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetNonce(&_EulerEvc.TransactOpts, addressPrefix, nonceNamespace, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xa829aaf5.
//
// Solidity: function setNonce(bytes19 addressPrefix, uint256 nonceNamespace, uint256 nonce) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) SetNonce(addressPrefix [19]byte, nonceNamespace *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetNonce(&_EulerEvc.TransactOpts, addressPrefix, nonceNamespace, nonce)
}

// SetOperator is a paid mutator transaction binding the contract method 0xc14c11bf.
//
// Solidity: function setOperator(bytes19 addressPrefix, address operator, uint256 operatorBitField) payable returns()
func (_EulerEvc *EulerEvcTransactor) SetOperator(opts *bind.TransactOpts, addressPrefix [19]byte, operator common.Address, operatorBitField *big.Int) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "setOperator", addressPrefix, operator, operatorBitField)
}

// SetOperator is a paid mutator transaction binding the contract method 0xc14c11bf.
//
// Solidity: function setOperator(bytes19 addressPrefix, address operator, uint256 operatorBitField) payable returns()
func (_EulerEvc *EulerEvcSession) SetOperator(addressPrefix [19]byte, operator common.Address, operatorBitField *big.Int) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetOperator(&_EulerEvc.TransactOpts, addressPrefix, operator, operatorBitField)
}

// SetOperator is a paid mutator transaction binding the contract method 0xc14c11bf.
//
// Solidity: function setOperator(bytes19 addressPrefix, address operator, uint256 operatorBitField) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) SetOperator(addressPrefix [19]byte, operator common.Address, operatorBitField *big.Int) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetOperator(&_EulerEvc.TransactOpts, addressPrefix, operator, operatorBitField)
}

// SetPermitDisabledMode is a paid mutator transaction binding the contract method 0x116d0e93.
//
// Solidity: function setPermitDisabledMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcTransactor) SetPermitDisabledMode(opts *bind.TransactOpts, addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.contract.Transact(opts, "setPermitDisabledMode", addressPrefix, enabled)
}

// SetPermitDisabledMode is a paid mutator transaction binding the contract method 0x116d0e93.
//
// Solidity: function setPermitDisabledMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcSession) SetPermitDisabledMode(addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetPermitDisabledMode(&_EulerEvc.TransactOpts, addressPrefix, enabled)
}

// SetPermitDisabledMode is a paid mutator transaction binding the contract method 0x116d0e93.
//
// Solidity: function setPermitDisabledMode(bytes19 addressPrefix, bool enabled) payable returns()
func (_EulerEvc *EulerEvcTransactorSession) SetPermitDisabledMode(addressPrefix [19]byte, enabled bool) (*types.Transaction, error) {
	return _EulerEvc.Contract.SetPermitDisabledMode(&_EulerEvc.TransactOpts, addressPrefix, enabled)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EulerEvc *EulerEvcTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EulerEvc *EulerEvcSession) Receive() (*types.Transaction, error) {
	return _EulerEvc.Contract.Receive(&_EulerEvc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EulerEvc *EulerEvcTransactorSession) Receive() (*types.Transaction, error) {
	return _EulerEvc.Contract.Receive(&_EulerEvc.TransactOpts)
}

// EulerEvcAccountStatusCheckIterator is returned from FilterAccountStatusCheck and is used to iterate over the raw logs and unpacked data for AccountStatusCheck events raised by the EulerEvc contract.
type EulerEvcAccountStatusCheckIterator struct {
	Event *EulerEvcAccountStatusCheck // Event containing the contract specifics and raw log

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
func (it *EulerEvcAccountStatusCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcAccountStatusCheck)
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
		it.Event = new(EulerEvcAccountStatusCheck)
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
func (it *EulerEvcAccountStatusCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcAccountStatusCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcAccountStatusCheck represents a AccountStatusCheck event raised by the EulerEvc contract.
type EulerEvcAccountStatusCheck struct {
	Account    common.Address
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountStatusCheck is a free log retrieval operation binding the contract event 0x889a4d4628b31342e420737e2aeb45387087570710d26239aa8a5f13d3e829d4.
//
// Solidity: event AccountStatusCheck(address indexed account, address indexed controller)
func (_EulerEvc *EulerEvcFilterer) FilterAccountStatusCheck(opts *bind.FilterOpts, account []common.Address, controller []common.Address) (*EulerEvcAccountStatusCheckIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "AccountStatusCheck", accountRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcAccountStatusCheckIterator{contract: _EulerEvc.contract, event: "AccountStatusCheck", logs: logs, sub: sub}, nil
}

// WatchAccountStatusCheck is a free log subscription operation binding the contract event 0x889a4d4628b31342e420737e2aeb45387087570710d26239aa8a5f13d3e829d4.
//
// Solidity: event AccountStatusCheck(address indexed account, address indexed controller)
func (_EulerEvc *EulerEvcFilterer) WatchAccountStatusCheck(opts *bind.WatchOpts, sink chan<- *EulerEvcAccountStatusCheck, account []common.Address, controller []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "AccountStatusCheck", accountRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcAccountStatusCheck)
				if err := _EulerEvc.contract.UnpackLog(event, "AccountStatusCheck", log); err != nil {
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

// ParseAccountStatusCheck is a log parse operation binding the contract event 0x889a4d4628b31342e420737e2aeb45387087570710d26239aa8a5f13d3e829d4.
//
// Solidity: event AccountStatusCheck(address indexed account, address indexed controller)
func (_EulerEvc *EulerEvcFilterer) ParseAccountStatusCheck(log types.Log) (*EulerEvcAccountStatusCheck, error) {
	event := new(EulerEvcAccountStatusCheck)
	if err := _EulerEvc.contract.UnpackLog(event, "AccountStatusCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcCallWithContextIterator is returned from FilterCallWithContext and is used to iterate over the raw logs and unpacked data for CallWithContext events raised by the EulerEvc contract.
type EulerEvcCallWithContextIterator struct {
	Event *EulerEvcCallWithContext // Event containing the contract specifics and raw log

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
func (it *EulerEvcCallWithContextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcCallWithContext)
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
		it.Event = new(EulerEvcCallWithContext)
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
func (it *EulerEvcCallWithContextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcCallWithContextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcCallWithContext represents a CallWithContext event raised by the EulerEvc contract.
type EulerEvcCallWithContext struct {
	Caller                  common.Address
	OnBehalfOfAddressPrefix [19]byte
	OnBehalfOfAccount       common.Address
	TargetContract          common.Address
	Selector                [4]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCallWithContext is a free log retrieval operation binding the contract event 0x6e9738e5aa38fe1517adbb480351ec386ece82947737b18badbcad1e911133ec.
//
// Solidity: event CallWithContext(address indexed caller, bytes19 indexed onBehalfOfAddressPrefix, address onBehalfOfAccount, address indexed targetContract, bytes4 selector)
func (_EulerEvc *EulerEvcFilterer) FilterCallWithContext(opts *bind.FilterOpts, caller []common.Address, onBehalfOfAddressPrefix [][19]byte, targetContract []common.Address) (*EulerEvcCallWithContextIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfAddressPrefixRule []interface{}
	for _, onBehalfOfAddressPrefixItem := range onBehalfOfAddressPrefix {
		onBehalfOfAddressPrefixRule = append(onBehalfOfAddressPrefixRule, onBehalfOfAddressPrefixItem)
	}

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "CallWithContext", callerRule, onBehalfOfAddressPrefixRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcCallWithContextIterator{contract: _EulerEvc.contract, event: "CallWithContext", logs: logs, sub: sub}, nil
}

// WatchCallWithContext is a free log subscription operation binding the contract event 0x6e9738e5aa38fe1517adbb480351ec386ece82947737b18badbcad1e911133ec.
//
// Solidity: event CallWithContext(address indexed caller, bytes19 indexed onBehalfOfAddressPrefix, address onBehalfOfAccount, address indexed targetContract, bytes4 selector)
func (_EulerEvc *EulerEvcFilterer) WatchCallWithContext(opts *bind.WatchOpts, sink chan<- *EulerEvcCallWithContext, caller []common.Address, onBehalfOfAddressPrefix [][19]byte, targetContract []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfAddressPrefixRule []interface{}
	for _, onBehalfOfAddressPrefixItem := range onBehalfOfAddressPrefix {
		onBehalfOfAddressPrefixRule = append(onBehalfOfAddressPrefixRule, onBehalfOfAddressPrefixItem)
	}

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "CallWithContext", callerRule, onBehalfOfAddressPrefixRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcCallWithContext)
				if err := _EulerEvc.contract.UnpackLog(event, "CallWithContext", log); err != nil {
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

// ParseCallWithContext is a log parse operation binding the contract event 0x6e9738e5aa38fe1517adbb480351ec386ece82947737b18badbcad1e911133ec.
//
// Solidity: event CallWithContext(address indexed caller, bytes19 indexed onBehalfOfAddressPrefix, address onBehalfOfAccount, address indexed targetContract, bytes4 selector)
func (_EulerEvc *EulerEvcFilterer) ParseCallWithContext(log types.Log) (*EulerEvcCallWithContext, error) {
	event := new(EulerEvcCallWithContext)
	if err := _EulerEvc.contract.UnpackLog(event, "CallWithContext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcCollateralStatusIterator is returned from FilterCollateralStatus and is used to iterate over the raw logs and unpacked data for CollateralStatus events raised by the EulerEvc contract.
type EulerEvcCollateralStatusIterator struct {
	Event *EulerEvcCollateralStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcCollateralStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcCollateralStatus)
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
		it.Event = new(EulerEvcCollateralStatus)
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
func (it *EulerEvcCollateralStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcCollateralStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcCollateralStatus represents a CollateralStatus event raised by the EulerEvc contract.
type EulerEvcCollateralStatus struct {
	Account    common.Address
	Collateral common.Address
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralStatus is a free log retrieval operation binding the contract event 0xf022705c827017c972043d1984cfddc7958c9f4685b4d9ce8bd68696f4381cd2.
//
// Solidity: event CollateralStatus(address indexed account, address indexed collateral, bool enabled)
func (_EulerEvc *EulerEvcFilterer) FilterCollateralStatus(opts *bind.FilterOpts, account []common.Address, collateral []common.Address) (*EulerEvcCollateralStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "CollateralStatus", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcCollateralStatusIterator{contract: _EulerEvc.contract, event: "CollateralStatus", logs: logs, sub: sub}, nil
}

// WatchCollateralStatus is a free log subscription operation binding the contract event 0xf022705c827017c972043d1984cfddc7958c9f4685b4d9ce8bd68696f4381cd2.
//
// Solidity: event CollateralStatus(address indexed account, address indexed collateral, bool enabled)
func (_EulerEvc *EulerEvcFilterer) WatchCollateralStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcCollateralStatus, account []common.Address, collateral []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "CollateralStatus", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcCollateralStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "CollateralStatus", log); err != nil {
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

// ParseCollateralStatus is a log parse operation binding the contract event 0xf022705c827017c972043d1984cfddc7958c9f4685b4d9ce8bd68696f4381cd2.
//
// Solidity: event CollateralStatus(address indexed account, address indexed collateral, bool enabled)
func (_EulerEvc *EulerEvcFilterer) ParseCollateralStatus(log types.Log) (*EulerEvcCollateralStatus, error) {
	event := new(EulerEvcCollateralStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "CollateralStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcControllerStatusIterator is returned from FilterControllerStatus and is used to iterate over the raw logs and unpacked data for ControllerStatus events raised by the EulerEvc contract.
type EulerEvcControllerStatusIterator struct {
	Event *EulerEvcControllerStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcControllerStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcControllerStatus)
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
		it.Event = new(EulerEvcControllerStatus)
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
func (it *EulerEvcControllerStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcControllerStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcControllerStatus represents a ControllerStatus event raised by the EulerEvc contract.
type EulerEvcControllerStatus struct {
	Account    common.Address
	Controller common.Address
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerStatus is a free log retrieval operation binding the contract event 0x9919d437ee612d4ec7bba88a7d9bc4fc36a0a23608ad6259252711a46b708af9.
//
// Solidity: event ControllerStatus(address indexed account, address indexed controller, bool enabled)
func (_EulerEvc *EulerEvcFilterer) FilterControllerStatus(opts *bind.FilterOpts, account []common.Address, controller []common.Address) (*EulerEvcControllerStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "ControllerStatus", accountRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcControllerStatusIterator{contract: _EulerEvc.contract, event: "ControllerStatus", logs: logs, sub: sub}, nil
}

// WatchControllerStatus is a free log subscription operation binding the contract event 0x9919d437ee612d4ec7bba88a7d9bc4fc36a0a23608ad6259252711a46b708af9.
//
// Solidity: event ControllerStatus(address indexed account, address indexed controller, bool enabled)
func (_EulerEvc *EulerEvcFilterer) WatchControllerStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcControllerStatus, account []common.Address, controller []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "ControllerStatus", accountRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcControllerStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "ControllerStatus", log); err != nil {
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

// ParseControllerStatus is a log parse operation binding the contract event 0x9919d437ee612d4ec7bba88a7d9bc4fc36a0a23608ad6259252711a46b708af9.
//
// Solidity: event ControllerStatus(address indexed account, address indexed controller, bool enabled)
func (_EulerEvc *EulerEvcFilterer) ParseControllerStatus(log types.Log) (*EulerEvcControllerStatus, error) {
	event := new(EulerEvcControllerStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "ControllerStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcLockdownModeStatusIterator is returned from FilterLockdownModeStatus and is used to iterate over the raw logs and unpacked data for LockdownModeStatus events raised by the EulerEvc contract.
type EulerEvcLockdownModeStatusIterator struct {
	Event *EulerEvcLockdownModeStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcLockdownModeStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcLockdownModeStatus)
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
		it.Event = new(EulerEvcLockdownModeStatus)
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
func (it *EulerEvcLockdownModeStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcLockdownModeStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcLockdownModeStatus represents a LockdownModeStatus event raised by the EulerEvc contract.
type EulerEvcLockdownModeStatus struct {
	AddressPrefix [19]byte
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockdownModeStatus is a free log retrieval operation binding the contract event 0xaf5120bc58372f0063d8362c9bba9070c462c07ae24c24082d080a426432798b.
//
// Solidity: event LockdownModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) FilterLockdownModeStatus(opts *bind.FilterOpts, addressPrefix [][19]byte) (*EulerEvcLockdownModeStatusIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "LockdownModeStatus", addressPrefixRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcLockdownModeStatusIterator{contract: _EulerEvc.contract, event: "LockdownModeStatus", logs: logs, sub: sub}, nil
}

// WatchLockdownModeStatus is a free log subscription operation binding the contract event 0xaf5120bc58372f0063d8362c9bba9070c462c07ae24c24082d080a426432798b.
//
// Solidity: event LockdownModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) WatchLockdownModeStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcLockdownModeStatus, addressPrefix [][19]byte) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "LockdownModeStatus", addressPrefixRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcLockdownModeStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "LockdownModeStatus", log); err != nil {
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

// ParseLockdownModeStatus is a log parse operation binding the contract event 0xaf5120bc58372f0063d8362c9bba9070c462c07ae24c24082d080a426432798b.
//
// Solidity: event LockdownModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) ParseLockdownModeStatus(log types.Log) (*EulerEvcLockdownModeStatus, error) {
	event := new(EulerEvcLockdownModeStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "LockdownModeStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcNonceStatusIterator is returned from FilterNonceStatus and is used to iterate over the raw logs and unpacked data for NonceStatus events raised by the EulerEvc contract.
type EulerEvcNonceStatusIterator struct {
	Event *EulerEvcNonceStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcNonceStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcNonceStatus)
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
		it.Event = new(EulerEvcNonceStatus)
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
func (it *EulerEvcNonceStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcNonceStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcNonceStatus represents a NonceStatus event raised by the EulerEvc contract.
type EulerEvcNonceStatus struct {
	AddressPrefix  [19]byte
	NonceNamespace *big.Int
	OldNonce       *big.Int
	NewNonce       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonceStatus is a free log retrieval operation binding the contract event 0x3b8510174a91acb36200f7427c1889f934941fd89ed86faf390749b4c2b46337.
//
// Solidity: event NonceStatus(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 oldNonce, uint256 newNonce)
func (_EulerEvc *EulerEvcFilterer) FilterNonceStatus(opts *bind.FilterOpts, addressPrefix [][19]byte, nonceNamespace []*big.Int) (*EulerEvcNonceStatusIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var nonceNamespaceRule []interface{}
	for _, nonceNamespaceItem := range nonceNamespace {
		nonceNamespaceRule = append(nonceNamespaceRule, nonceNamespaceItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "NonceStatus", addressPrefixRule, nonceNamespaceRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcNonceStatusIterator{contract: _EulerEvc.contract, event: "NonceStatus", logs: logs, sub: sub}, nil
}

// WatchNonceStatus is a free log subscription operation binding the contract event 0x3b8510174a91acb36200f7427c1889f934941fd89ed86faf390749b4c2b46337.
//
// Solidity: event NonceStatus(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 oldNonce, uint256 newNonce)
func (_EulerEvc *EulerEvcFilterer) WatchNonceStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcNonceStatus, addressPrefix [][19]byte, nonceNamespace []*big.Int) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var nonceNamespaceRule []interface{}
	for _, nonceNamespaceItem := range nonceNamespace {
		nonceNamespaceRule = append(nonceNamespaceRule, nonceNamespaceItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "NonceStatus", addressPrefixRule, nonceNamespaceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcNonceStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "NonceStatus", log); err != nil {
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

// ParseNonceStatus is a log parse operation binding the contract event 0x3b8510174a91acb36200f7427c1889f934941fd89ed86faf390749b4c2b46337.
//
// Solidity: event NonceStatus(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 oldNonce, uint256 newNonce)
func (_EulerEvc *EulerEvcFilterer) ParseNonceStatus(log types.Log) (*EulerEvcNonceStatus, error) {
	event := new(EulerEvcNonceStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "NonceStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcNonceUsedIterator is returned from FilterNonceUsed and is used to iterate over the raw logs and unpacked data for NonceUsed events raised by the EulerEvc contract.
type EulerEvcNonceUsedIterator struct {
	Event *EulerEvcNonceUsed // Event containing the contract specifics and raw log

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
func (it *EulerEvcNonceUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcNonceUsed)
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
		it.Event = new(EulerEvcNonceUsed)
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
func (it *EulerEvcNonceUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcNonceUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcNonceUsed represents a NonceUsed event raised by the EulerEvc contract.
type EulerEvcNonceUsed struct {
	AddressPrefix  [19]byte
	NonceNamespace *big.Int
	Nonce          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonceUsed is a free log retrieval operation binding the contract event 0xb0dcec731e48090736be6db10ad9f9581d0ec5fc0f1925a8e267b64b614b08d6.
//
// Solidity: event NonceUsed(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 nonce)
func (_EulerEvc *EulerEvcFilterer) FilterNonceUsed(opts *bind.FilterOpts, addressPrefix [][19]byte, nonceNamespace []*big.Int) (*EulerEvcNonceUsedIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var nonceNamespaceRule []interface{}
	for _, nonceNamespaceItem := range nonceNamespace {
		nonceNamespaceRule = append(nonceNamespaceRule, nonceNamespaceItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "NonceUsed", addressPrefixRule, nonceNamespaceRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcNonceUsedIterator{contract: _EulerEvc.contract, event: "NonceUsed", logs: logs, sub: sub}, nil
}

// WatchNonceUsed is a free log subscription operation binding the contract event 0xb0dcec731e48090736be6db10ad9f9581d0ec5fc0f1925a8e267b64b614b08d6.
//
// Solidity: event NonceUsed(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 nonce)
func (_EulerEvc *EulerEvcFilterer) WatchNonceUsed(opts *bind.WatchOpts, sink chan<- *EulerEvcNonceUsed, addressPrefix [][19]byte, nonceNamespace []*big.Int) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var nonceNamespaceRule []interface{}
	for _, nonceNamespaceItem := range nonceNamespace {
		nonceNamespaceRule = append(nonceNamespaceRule, nonceNamespaceItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "NonceUsed", addressPrefixRule, nonceNamespaceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcNonceUsed)
				if err := _EulerEvc.contract.UnpackLog(event, "NonceUsed", log); err != nil {
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

// ParseNonceUsed is a log parse operation binding the contract event 0xb0dcec731e48090736be6db10ad9f9581d0ec5fc0f1925a8e267b64b614b08d6.
//
// Solidity: event NonceUsed(bytes19 indexed addressPrefix, uint256 indexed nonceNamespace, uint256 nonce)
func (_EulerEvc *EulerEvcFilterer) ParseNonceUsed(log types.Log) (*EulerEvcNonceUsed, error) {
	event := new(EulerEvcNonceUsed)
	if err := _EulerEvc.contract.UnpackLog(event, "NonceUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcOperatorStatusIterator is returned from FilterOperatorStatus and is used to iterate over the raw logs and unpacked data for OperatorStatus events raised by the EulerEvc contract.
type EulerEvcOperatorStatusIterator struct {
	Event *EulerEvcOperatorStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcOperatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcOperatorStatus)
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
		it.Event = new(EulerEvcOperatorStatus)
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
func (it *EulerEvcOperatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcOperatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcOperatorStatus represents a OperatorStatus event raised by the EulerEvc contract.
type EulerEvcOperatorStatus struct {
	AddressPrefix             [19]byte
	Operator                  common.Address
	AccountOperatorAuthorized *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterOperatorStatus is a free log retrieval operation binding the contract event 0x7ba31654d8467e98b6bd4fc56ddde246de9ade831cf860c7ac695579aecb9564.
//
// Solidity: event OperatorStatus(bytes19 indexed addressPrefix, address indexed operator, uint256 accountOperatorAuthorized)
func (_EulerEvc *EulerEvcFilterer) FilterOperatorStatus(opts *bind.FilterOpts, addressPrefix [][19]byte, operator []common.Address) (*EulerEvcOperatorStatusIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "OperatorStatus", addressPrefixRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcOperatorStatusIterator{contract: _EulerEvc.contract, event: "OperatorStatus", logs: logs, sub: sub}, nil
}

// WatchOperatorStatus is a free log subscription operation binding the contract event 0x7ba31654d8467e98b6bd4fc56ddde246de9ade831cf860c7ac695579aecb9564.
//
// Solidity: event OperatorStatus(bytes19 indexed addressPrefix, address indexed operator, uint256 accountOperatorAuthorized)
func (_EulerEvc *EulerEvcFilterer) WatchOperatorStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcOperatorStatus, addressPrefix [][19]byte, operator []common.Address) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "OperatorStatus", addressPrefixRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcOperatorStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "OperatorStatus", log); err != nil {
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

// ParseOperatorStatus is a log parse operation binding the contract event 0x7ba31654d8467e98b6bd4fc56ddde246de9ade831cf860c7ac695579aecb9564.
//
// Solidity: event OperatorStatus(bytes19 indexed addressPrefix, address indexed operator, uint256 accountOperatorAuthorized)
func (_EulerEvc *EulerEvcFilterer) ParseOperatorStatus(log types.Log) (*EulerEvcOperatorStatus, error) {
	event := new(EulerEvcOperatorStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "OperatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcOwnerRegisteredIterator is returned from FilterOwnerRegistered and is used to iterate over the raw logs and unpacked data for OwnerRegistered events raised by the EulerEvc contract.
type EulerEvcOwnerRegisteredIterator struct {
	Event *EulerEvcOwnerRegistered // Event containing the contract specifics and raw log

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
func (it *EulerEvcOwnerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcOwnerRegistered)
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
		it.Event = new(EulerEvcOwnerRegistered)
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
func (it *EulerEvcOwnerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcOwnerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcOwnerRegistered represents a OwnerRegistered event raised by the EulerEvc contract.
type EulerEvcOwnerRegistered struct {
	AddressPrefix [19]byte
	Owner         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnerRegistered is a free log retrieval operation binding the contract event 0x67cb2734834e775d6db886bf16ac03d7273b290223ee5363354b385ec5b643b0.
//
// Solidity: event OwnerRegistered(bytes19 indexed addressPrefix, address indexed owner)
func (_EulerEvc *EulerEvcFilterer) FilterOwnerRegistered(opts *bind.FilterOpts, addressPrefix [][19]byte, owner []common.Address) (*EulerEvcOwnerRegisteredIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "OwnerRegistered", addressPrefixRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcOwnerRegisteredIterator{contract: _EulerEvc.contract, event: "OwnerRegistered", logs: logs, sub: sub}, nil
}

// WatchOwnerRegistered is a free log subscription operation binding the contract event 0x67cb2734834e775d6db886bf16ac03d7273b290223ee5363354b385ec5b643b0.
//
// Solidity: event OwnerRegistered(bytes19 indexed addressPrefix, address indexed owner)
func (_EulerEvc *EulerEvcFilterer) WatchOwnerRegistered(opts *bind.WatchOpts, sink chan<- *EulerEvcOwnerRegistered, addressPrefix [][19]byte, owner []common.Address) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "OwnerRegistered", addressPrefixRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcOwnerRegistered)
				if err := _EulerEvc.contract.UnpackLog(event, "OwnerRegistered", log); err != nil {
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

// ParseOwnerRegistered is a log parse operation binding the contract event 0x67cb2734834e775d6db886bf16ac03d7273b290223ee5363354b385ec5b643b0.
//
// Solidity: event OwnerRegistered(bytes19 indexed addressPrefix, address indexed owner)
func (_EulerEvc *EulerEvcFilterer) ParseOwnerRegistered(log types.Log) (*EulerEvcOwnerRegistered, error) {
	event := new(EulerEvcOwnerRegistered)
	if err := _EulerEvc.contract.UnpackLog(event, "OwnerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcPermitDisabledModeStatusIterator is returned from FilterPermitDisabledModeStatus and is used to iterate over the raw logs and unpacked data for PermitDisabledModeStatus events raised by the EulerEvc contract.
type EulerEvcPermitDisabledModeStatusIterator struct {
	Event *EulerEvcPermitDisabledModeStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvcPermitDisabledModeStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcPermitDisabledModeStatus)
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
		it.Event = new(EulerEvcPermitDisabledModeStatus)
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
func (it *EulerEvcPermitDisabledModeStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcPermitDisabledModeStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcPermitDisabledModeStatus represents a PermitDisabledModeStatus event raised by the EulerEvc contract.
type EulerEvcPermitDisabledModeStatus struct {
	AddressPrefix [19]byte
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPermitDisabledModeStatus is a free log retrieval operation binding the contract event 0x6321df4e44267d425279195e7488fadba1a42d5cce9e84f596d5cf696f4449cd.
//
// Solidity: event PermitDisabledModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) FilterPermitDisabledModeStatus(opts *bind.FilterOpts, addressPrefix [][19]byte) (*EulerEvcPermitDisabledModeStatusIterator, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "PermitDisabledModeStatus", addressPrefixRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcPermitDisabledModeStatusIterator{contract: _EulerEvc.contract, event: "PermitDisabledModeStatus", logs: logs, sub: sub}, nil
}

// WatchPermitDisabledModeStatus is a free log subscription operation binding the contract event 0x6321df4e44267d425279195e7488fadba1a42d5cce9e84f596d5cf696f4449cd.
//
// Solidity: event PermitDisabledModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) WatchPermitDisabledModeStatus(opts *bind.WatchOpts, sink chan<- *EulerEvcPermitDisabledModeStatus, addressPrefix [][19]byte) (event.Subscription, error) {

	var addressPrefixRule []interface{}
	for _, addressPrefixItem := range addressPrefix {
		addressPrefixRule = append(addressPrefixRule, addressPrefixItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "PermitDisabledModeStatus", addressPrefixRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcPermitDisabledModeStatus)
				if err := _EulerEvc.contract.UnpackLog(event, "PermitDisabledModeStatus", log); err != nil {
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

// ParsePermitDisabledModeStatus is a log parse operation binding the contract event 0x6321df4e44267d425279195e7488fadba1a42d5cce9e84f596d5cf696f4449cd.
//
// Solidity: event PermitDisabledModeStatus(bytes19 indexed addressPrefix, bool enabled)
func (_EulerEvc *EulerEvcFilterer) ParsePermitDisabledModeStatus(log types.Log) (*EulerEvcPermitDisabledModeStatus, error) {
	event := new(EulerEvcPermitDisabledModeStatus)
	if err := _EulerEvc.contract.UnpackLog(event, "PermitDisabledModeStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvcVaultStatusCheckIterator is returned from FilterVaultStatusCheck and is used to iterate over the raw logs and unpacked data for VaultStatusCheck events raised by the EulerEvc contract.
type EulerEvcVaultStatusCheckIterator struct {
	Event *EulerEvcVaultStatusCheck // Event containing the contract specifics and raw log

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
func (it *EulerEvcVaultStatusCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvcVaultStatusCheck)
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
		it.Event = new(EulerEvcVaultStatusCheck)
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
func (it *EulerEvcVaultStatusCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvcVaultStatusCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvcVaultStatusCheck represents a VaultStatusCheck event raised by the EulerEvc contract.
type EulerEvcVaultStatusCheck struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultStatusCheck is a free log retrieval operation binding the contract event 0xaea973cfb51ea8ca328767d72f105b5b9d2360c65f5ac4110a2c4470434471c9.
//
// Solidity: event VaultStatusCheck(address indexed vault)
func (_EulerEvc *EulerEvcFilterer) FilterVaultStatusCheck(opts *bind.FilterOpts, vault []common.Address) (*EulerEvcVaultStatusCheckIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerEvc.contract.FilterLogs(opts, "VaultStatusCheck", vaultRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvcVaultStatusCheckIterator{contract: _EulerEvc.contract, event: "VaultStatusCheck", logs: logs, sub: sub}, nil
}

// WatchVaultStatusCheck is a free log subscription operation binding the contract event 0xaea973cfb51ea8ca328767d72f105b5b9d2360c65f5ac4110a2c4470434471c9.
//
// Solidity: event VaultStatusCheck(address indexed vault)
func (_EulerEvc *EulerEvcFilterer) WatchVaultStatusCheck(opts *bind.WatchOpts, sink chan<- *EulerEvcVaultStatusCheck, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerEvc.contract.WatchLogs(opts, "VaultStatusCheck", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvcVaultStatusCheck)
				if err := _EulerEvc.contract.UnpackLog(event, "VaultStatusCheck", log); err != nil {
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

// ParseVaultStatusCheck is a log parse operation binding the contract event 0xaea973cfb51ea8ca328767d72f105b5b9d2360c65f5ac4110a2c4470434471c9.
//
// Solidity: event VaultStatusCheck(address indexed vault)
func (_EulerEvc *EulerEvcFilterer) ParseVaultStatusCheck(log types.Log) (*EulerEvcVaultStatusCheck, error) {
	event := new(EulerEvcVaultStatusCheck)
	if err := _EulerEvc.contract.UnpackLog(event, "VaultStatusCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
