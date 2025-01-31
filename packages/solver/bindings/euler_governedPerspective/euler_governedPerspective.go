// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_governedPerspective

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

// EulerGovernedPerspectiveMetaData contains all meta data concerning the EulerGovernedPerspective contract.
var EulerGovernedPerspectiveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"perspective\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"codes\",\"type\":\"uint256\"}],\"name\":\"PerspectiveError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerspectivePanic\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"PerspectiveUnverified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"PerspectiveVerified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"isVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"perspectiveUnverify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"failEarly\",\"type\":\"bool\"}],\"name\":\"perspectiveVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultFactory\",\"outputs\":[{\"internalType\":\"contractGenericFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerGovernedPerspectiveABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerGovernedPerspectiveMetaData.ABI instead.
var EulerGovernedPerspectiveABI = EulerGovernedPerspectiveMetaData.ABI

// EulerGovernedPerspective is an auto generated Go binding around an Ethereum contract.
type EulerGovernedPerspective struct {
	EulerGovernedPerspectiveCaller     // Read-only binding to the contract
	EulerGovernedPerspectiveTransactor // Write-only binding to the contract
	EulerGovernedPerspectiveFilterer   // Log filterer for contract events
}

// EulerGovernedPerspectiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerGovernedPerspectiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerGovernedPerspectiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerGovernedPerspectiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerGovernedPerspectiveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerGovernedPerspectiveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerGovernedPerspectiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerGovernedPerspectiveSession struct {
	Contract     *EulerGovernedPerspective // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EulerGovernedPerspectiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerGovernedPerspectiveCallerSession struct {
	Contract *EulerGovernedPerspectiveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// EulerGovernedPerspectiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerGovernedPerspectiveTransactorSession struct {
	Contract     *EulerGovernedPerspectiveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EulerGovernedPerspectiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerGovernedPerspectiveRaw struct {
	Contract *EulerGovernedPerspective // Generic contract binding to access the raw methods on
}

// EulerGovernedPerspectiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerGovernedPerspectiveCallerRaw struct {
	Contract *EulerGovernedPerspectiveCaller // Generic read-only contract binding to access the raw methods on
}

// EulerGovernedPerspectiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerGovernedPerspectiveTransactorRaw struct {
	Contract *EulerGovernedPerspectiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerGovernedPerspective creates a new instance of EulerGovernedPerspective, bound to a specific deployed contract.
func NewEulerGovernedPerspective(address common.Address, backend bind.ContractBackend) (*EulerGovernedPerspective, error) {
	contract, err := bindEulerGovernedPerspective(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspective{EulerGovernedPerspectiveCaller: EulerGovernedPerspectiveCaller{contract: contract}, EulerGovernedPerspectiveTransactor: EulerGovernedPerspectiveTransactor{contract: contract}, EulerGovernedPerspectiveFilterer: EulerGovernedPerspectiveFilterer{contract: contract}}, nil
}

// NewEulerGovernedPerspectiveCaller creates a new read-only instance of EulerGovernedPerspective, bound to a specific deployed contract.
func NewEulerGovernedPerspectiveCaller(address common.Address, caller bind.ContractCaller) (*EulerGovernedPerspectiveCaller, error) {
	contract, err := bindEulerGovernedPerspective(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectiveCaller{contract: contract}, nil
}

// NewEulerGovernedPerspectiveTransactor creates a new write-only instance of EulerGovernedPerspective, bound to a specific deployed contract.
func NewEulerGovernedPerspectiveTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerGovernedPerspectiveTransactor, error) {
	contract, err := bindEulerGovernedPerspective(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectiveTransactor{contract: contract}, nil
}

// NewEulerGovernedPerspectiveFilterer creates a new log filterer instance of EulerGovernedPerspective, bound to a specific deployed contract.
func NewEulerGovernedPerspectiveFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerGovernedPerspectiveFilterer, error) {
	contract, err := bindEulerGovernedPerspective(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectiveFilterer{contract: contract}, nil
}

// bindEulerGovernedPerspective binds a generic wrapper to an already deployed contract.
func bindEulerGovernedPerspective(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerGovernedPerspectiveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerGovernedPerspective.Contract.EulerGovernedPerspectiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.EulerGovernedPerspectiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.EulerGovernedPerspectiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerGovernedPerspective.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) EVC() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.EVC(&_EulerGovernedPerspective.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) EVC() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.EVC(&_EulerGovernedPerspective.CallOpts)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address vault) view returns(bool)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) IsVerified(opts *bind.CallOpts, vault common.Address) (bool, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "isVerified", vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address vault) view returns(bool)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) IsVerified(vault common.Address) (bool, error) {
	return _EulerGovernedPerspective.Contract.IsVerified(&_EulerGovernedPerspective.CallOpts, vault)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address vault) view returns(bool)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) IsVerified(vault common.Address) (bool, error) {
	return _EulerGovernedPerspective.Contract.IsVerified(&_EulerGovernedPerspective.CallOpts, vault)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) Name() (string, error) {
	return _EulerGovernedPerspective.Contract.Name(&_EulerGovernedPerspective.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) Name() (string, error) {
	return _EulerGovernedPerspective.Contract.Name(&_EulerGovernedPerspective.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) Owner() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.Owner(&_EulerGovernedPerspective.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) Owner() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.Owner(&_EulerGovernedPerspective.CallOpts)
}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) VaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "vaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) VaultFactory() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.VaultFactory(&_EulerGovernedPerspective.CallOpts)
}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() view returns(address)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) VaultFactory() (common.Address, error) {
	return _EulerGovernedPerspective.Contract.VaultFactory(&_EulerGovernedPerspective.CallOpts)
}

// VerifiedArray is a free data retrieval call binding the contract method 0x8d5e21d3.
//
// Solidity: function verifiedArray() view returns(address[])
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) VerifiedArray(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "verifiedArray")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// VerifiedArray is a free data retrieval call binding the contract method 0x8d5e21d3.
//
// Solidity: function verifiedArray() view returns(address[])
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) VerifiedArray() ([]common.Address, error) {
	return _EulerGovernedPerspective.Contract.VerifiedArray(&_EulerGovernedPerspective.CallOpts)
}

// VerifiedArray is a free data retrieval call binding the contract method 0x8d5e21d3.
//
// Solidity: function verifiedArray() view returns(address[])
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) VerifiedArray() ([]common.Address, error) {
	return _EulerGovernedPerspective.Contract.VerifiedArray(&_EulerGovernedPerspective.CallOpts)
}

// VerifiedLength is a free data retrieval call binding the contract method 0x138721d9.
//
// Solidity: function verifiedLength() view returns(uint256)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCaller) VerifiedLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerGovernedPerspective.contract.Call(opts, &out, "verifiedLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifiedLength is a free data retrieval call binding the contract method 0x138721d9.
//
// Solidity: function verifiedLength() view returns(uint256)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) VerifiedLength() (*big.Int, error) {
	return _EulerGovernedPerspective.Contract.VerifiedLength(&_EulerGovernedPerspective.CallOpts)
}

// VerifiedLength is a free data retrieval call binding the contract method 0x138721d9.
//
// Solidity: function verifiedLength() view returns(uint256)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveCallerSession) VerifiedLength() (*big.Int, error) {
	return _EulerGovernedPerspective.Contract.VerifiedLength(&_EulerGovernedPerspective.CallOpts)
}

// PerspectiveUnverify is a paid mutator transaction binding the contract method 0x6270684f.
//
// Solidity: function perspectiveUnverify(address vault) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactor) PerspectiveUnverify(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.contract.Transact(opts, "perspectiveUnverify", vault)
}

// PerspectiveUnverify is a paid mutator transaction binding the contract method 0x6270684f.
//
// Solidity: function perspectiveUnverify(address vault) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) PerspectiveUnverify(vault common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.PerspectiveUnverify(&_EulerGovernedPerspective.TransactOpts, vault)
}

// PerspectiveUnverify is a paid mutator transaction binding the contract method 0x6270684f.
//
// Solidity: function perspectiveUnverify(address vault) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorSession) PerspectiveUnverify(vault common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.PerspectiveUnverify(&_EulerGovernedPerspective.TransactOpts, vault)
}

// PerspectiveVerify is a paid mutator transaction binding the contract method 0x2e5896e5.
//
// Solidity: function perspectiveVerify(address vault, bool failEarly) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactor) PerspectiveVerify(opts *bind.TransactOpts, vault common.Address, failEarly bool) (*types.Transaction, error) {
	return _EulerGovernedPerspective.contract.Transact(opts, "perspectiveVerify", vault, failEarly)
}

// PerspectiveVerify is a paid mutator transaction binding the contract method 0x2e5896e5.
//
// Solidity: function perspectiveVerify(address vault, bool failEarly) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) PerspectiveVerify(vault common.Address, failEarly bool) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.PerspectiveVerify(&_EulerGovernedPerspective.TransactOpts, vault, failEarly)
}

// PerspectiveVerify is a paid mutator transaction binding the contract method 0x2e5896e5.
//
// Solidity: function perspectiveVerify(address vault, bool failEarly) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorSession) PerspectiveVerify(vault common.Address, failEarly bool) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.PerspectiveVerify(&_EulerGovernedPerspective.TransactOpts, vault, failEarly)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerGovernedPerspective.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) RenounceOwnership() (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.RenounceOwnership(&_EulerGovernedPerspective.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.RenounceOwnership(&_EulerGovernedPerspective.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.TransferOwnership(&_EulerGovernedPerspective.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EulerGovernedPerspective *EulerGovernedPerspectiveTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EulerGovernedPerspective.Contract.TransferOwnership(&_EulerGovernedPerspective.TransactOpts, newOwner)
}

// EulerGovernedPerspectiveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectiveOwnershipTransferredIterator struct {
	Event *EulerGovernedPerspectiveOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EulerGovernedPerspectiveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerGovernedPerspectiveOwnershipTransferred)
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
		it.Event = new(EulerGovernedPerspectiveOwnershipTransferred)
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
func (it *EulerGovernedPerspectiveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerGovernedPerspectiveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerGovernedPerspectiveOwnershipTransferred represents a OwnershipTransferred event raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectiveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EulerGovernedPerspectiveOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectiveOwnershipTransferredIterator{contract: _EulerGovernedPerspective.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EulerGovernedPerspectiveOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerGovernedPerspectiveOwnershipTransferred)
				if err := _EulerGovernedPerspective.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) ParseOwnershipTransferred(log types.Log) (*EulerGovernedPerspectiveOwnershipTransferred, error) {
	event := new(EulerGovernedPerspectiveOwnershipTransferred)
	if err := _EulerGovernedPerspective.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerGovernedPerspectivePerspectiveUnverifiedIterator is returned from FilterPerspectiveUnverified and is used to iterate over the raw logs and unpacked data for PerspectiveUnverified events raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectivePerspectiveUnverifiedIterator struct {
	Event *EulerGovernedPerspectivePerspectiveUnverified // Event containing the contract specifics and raw log

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
func (it *EulerGovernedPerspectivePerspectiveUnverifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerGovernedPerspectivePerspectiveUnverified)
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
		it.Event = new(EulerGovernedPerspectivePerspectiveUnverified)
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
func (it *EulerGovernedPerspectivePerspectiveUnverifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerGovernedPerspectivePerspectiveUnverifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerGovernedPerspectivePerspectiveUnverified represents a PerspectiveUnverified event raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectivePerspectiveUnverified struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPerspectiveUnverified is a free log retrieval operation binding the contract event 0xc4f218915c2884639e62cc3bea5ef5c711dc26f0e3d6b23e2de05cc12c041a66.
//
// Solidity: event PerspectiveUnverified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) FilterPerspectiveUnverified(opts *bind.FilterOpts, vault []common.Address) (*EulerGovernedPerspectivePerspectiveUnverifiedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.FilterLogs(opts, "PerspectiveUnverified", vaultRule)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectivePerspectiveUnverifiedIterator{contract: _EulerGovernedPerspective.contract, event: "PerspectiveUnverified", logs: logs, sub: sub}, nil
}

// WatchPerspectiveUnverified is a free log subscription operation binding the contract event 0xc4f218915c2884639e62cc3bea5ef5c711dc26f0e3d6b23e2de05cc12c041a66.
//
// Solidity: event PerspectiveUnverified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) WatchPerspectiveUnverified(opts *bind.WatchOpts, sink chan<- *EulerGovernedPerspectivePerspectiveUnverified, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.WatchLogs(opts, "PerspectiveUnverified", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerGovernedPerspectivePerspectiveUnverified)
				if err := _EulerGovernedPerspective.contract.UnpackLog(event, "PerspectiveUnverified", log); err != nil {
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

// ParsePerspectiveUnverified is a log parse operation binding the contract event 0xc4f218915c2884639e62cc3bea5ef5c711dc26f0e3d6b23e2de05cc12c041a66.
//
// Solidity: event PerspectiveUnverified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) ParsePerspectiveUnverified(log types.Log) (*EulerGovernedPerspectivePerspectiveUnverified, error) {
	event := new(EulerGovernedPerspectivePerspectiveUnverified)
	if err := _EulerGovernedPerspective.contract.UnpackLog(event, "PerspectiveUnverified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerGovernedPerspectivePerspectiveVerifiedIterator is returned from FilterPerspectiveVerified and is used to iterate over the raw logs and unpacked data for PerspectiveVerified events raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectivePerspectiveVerifiedIterator struct {
	Event *EulerGovernedPerspectivePerspectiveVerified // Event containing the contract specifics and raw log

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
func (it *EulerGovernedPerspectivePerspectiveVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerGovernedPerspectivePerspectiveVerified)
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
		it.Event = new(EulerGovernedPerspectivePerspectiveVerified)
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
func (it *EulerGovernedPerspectivePerspectiveVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerGovernedPerspectivePerspectiveVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerGovernedPerspectivePerspectiveVerified represents a PerspectiveVerified event raised by the EulerGovernedPerspective contract.
type EulerGovernedPerspectivePerspectiveVerified struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPerspectiveVerified is a free log retrieval operation binding the contract event 0x570e1c1f1f2e6e95bfd6d0cae607f36c3cd5ebb7bc35c2f87299924b1bcd3920.
//
// Solidity: event PerspectiveVerified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) FilterPerspectiveVerified(opts *bind.FilterOpts, vault []common.Address) (*EulerGovernedPerspectivePerspectiveVerifiedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.FilterLogs(opts, "PerspectiveVerified", vaultRule)
	if err != nil {
		return nil, err
	}
	return &EulerGovernedPerspectivePerspectiveVerifiedIterator{contract: _EulerGovernedPerspective.contract, event: "PerspectiveVerified", logs: logs, sub: sub}, nil
}

// WatchPerspectiveVerified is a free log subscription operation binding the contract event 0x570e1c1f1f2e6e95bfd6d0cae607f36c3cd5ebb7bc35c2f87299924b1bcd3920.
//
// Solidity: event PerspectiveVerified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) WatchPerspectiveVerified(opts *bind.WatchOpts, sink chan<- *EulerGovernedPerspectivePerspectiveVerified, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _EulerGovernedPerspective.contract.WatchLogs(opts, "PerspectiveVerified", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerGovernedPerspectivePerspectiveVerified)
				if err := _EulerGovernedPerspective.contract.UnpackLog(event, "PerspectiveVerified", log); err != nil {
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

// ParsePerspectiveVerified is a log parse operation binding the contract event 0x570e1c1f1f2e6e95bfd6d0cae607f36c3cd5ebb7bc35c2f87299924b1bcd3920.
//
// Solidity: event PerspectiveVerified(address indexed vault)
func (_EulerGovernedPerspective *EulerGovernedPerspectiveFilterer) ParsePerspectiveVerified(log types.Log) (*EulerGovernedPerspectivePerspectiveVerified, error) {
	event := new(EulerGovernedPerspectivePerspectiveVerified)
	if err := _EulerGovernedPerspective.contract.UnpackLog(event, "PerspectiveVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
