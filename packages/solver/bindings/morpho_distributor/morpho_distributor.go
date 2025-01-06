// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho_distributor

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

// MorphoDistributorMetaData contains all meta data concerning the MorphoDistributor contract.
var MorphoDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"initialRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initialIpfsHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"acceptRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUpdater\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ipfsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"validAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newIpfsHash\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"setTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newIpfsHash\",\"type\":\"bytes32\"}],\"name\":\"submitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MorphoDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoDistributorMetaData.ABI instead.
var MorphoDistributorABI = MorphoDistributorMetaData.ABI

// MorphoDistributor is an auto generated Go binding around an Ethereum contract.
type MorphoDistributor struct {
	MorphoDistributorCaller     // Read-only binding to the contract
	MorphoDistributorTransactor // Write-only binding to the contract
	MorphoDistributorFilterer   // Log filterer for contract events
}

// MorphoDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoDistributorSession struct {
	Contract     *MorphoDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MorphoDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoDistributorCallerSession struct {
	Contract *MorphoDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MorphoDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoDistributorTransactorSession struct {
	Contract     *MorphoDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MorphoDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoDistributorRaw struct {
	Contract *MorphoDistributor // Generic contract binding to access the raw methods on
}

// MorphoDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoDistributorCallerRaw struct {
	Contract *MorphoDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoDistributorTransactorRaw struct {
	Contract *MorphoDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphoDistributor creates a new instance of MorphoDistributor, bound to a specific deployed contract.
func NewMorphoDistributor(address common.Address, backend bind.ContractBackend) (*MorphoDistributor, error) {
	contract, err := bindMorphoDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphoDistributor{MorphoDistributorCaller: MorphoDistributorCaller{contract: contract}, MorphoDistributorTransactor: MorphoDistributorTransactor{contract: contract}, MorphoDistributorFilterer: MorphoDistributorFilterer{contract: contract}}, nil
}

// NewMorphoDistributorCaller creates a new read-only instance of MorphoDistributor, bound to a specific deployed contract.
func NewMorphoDistributorCaller(address common.Address, caller bind.ContractCaller) (*MorphoDistributorCaller, error) {
	contract, err := bindMorphoDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoDistributorCaller{contract: contract}, nil
}

// NewMorphoDistributorTransactor creates a new write-only instance of MorphoDistributor, bound to a specific deployed contract.
func NewMorphoDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoDistributorTransactor, error) {
	contract, err := bindMorphoDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoDistributorTransactor{contract: contract}, nil
}

// NewMorphoDistributorFilterer creates a new log filterer instance of MorphoDistributor, bound to a specific deployed contract.
func NewMorphoDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoDistributorFilterer, error) {
	contract, err := bindMorphoDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoDistributorFilterer{contract: contract}, nil
}

// bindMorphoDistributor binds a generic wrapper to an already deployed contract.
func bindMorphoDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoDistributor *MorphoDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoDistributor.Contract.MorphoDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoDistributor *MorphoDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.MorphoDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoDistributor *MorphoDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.MorphoDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoDistributor *MorphoDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoDistributor *MorphoDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoDistributor *MorphoDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0x0c9cbf0e.
//
// Solidity: function claimed(address account, address reward) view returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorCaller) Claimed(opts *bind.CallOpts, account common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "claimed", account, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x0c9cbf0e.
//
// Solidity: function claimed(address account, address reward) view returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorSession) Claimed(account common.Address, reward common.Address) (*big.Int, error) {
	return _MorphoDistributor.Contract.Claimed(&_MorphoDistributor.CallOpts, account, reward)
}

// Claimed is a free data retrieval call binding the contract method 0x0c9cbf0e.
//
// Solidity: function claimed(address account, address reward) view returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorCallerSession) Claimed(account common.Address, reward common.Address) (*big.Int, error) {
	return _MorphoDistributor.Contract.Claimed(&_MorphoDistributor.CallOpts, account, reward)
}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorCaller) IpfsHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "ipfsHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorSession) IpfsHash() ([32]byte, error) {
	return _MorphoDistributor.Contract.IpfsHash(&_MorphoDistributor.CallOpts)
}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorCallerSession) IpfsHash() ([32]byte, error) {
	return _MorphoDistributor.Contract.IpfsHash(&_MorphoDistributor.CallOpts)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_MorphoDistributor *MorphoDistributorCaller) IsUpdater(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "isUpdater", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_MorphoDistributor *MorphoDistributorSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _MorphoDistributor.Contract.IsUpdater(&_MorphoDistributor.CallOpts, arg0)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_MorphoDistributor *MorphoDistributorCallerSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _MorphoDistributor.Contract.IsUpdater(&_MorphoDistributor.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoDistributor *MorphoDistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoDistributor *MorphoDistributorSession) Owner() (common.Address, error) {
	return _MorphoDistributor.Contract.Owner(&_MorphoDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoDistributor *MorphoDistributorCallerSession) Owner() (common.Address, error) {
	return _MorphoDistributor.Contract.Owner(&_MorphoDistributor.CallOpts)
}

// PendingRoot is a free data retrieval call binding the contract method 0x750588cf.
//
// Solidity: function pendingRoot() view returns(bytes32 root, bytes32 ipfsHash, uint256 validAt)
func (_MorphoDistributor *MorphoDistributorCaller) PendingRoot(opts *bind.CallOpts) (struct {
	Root     [32]byte
	IpfsHash [32]byte
	ValidAt  *big.Int
}, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "pendingRoot")

	outstruct := new(struct {
		Root     [32]byte
		IpfsHash [32]byte
		ValidAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.IpfsHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ValidAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingRoot is a free data retrieval call binding the contract method 0x750588cf.
//
// Solidity: function pendingRoot() view returns(bytes32 root, bytes32 ipfsHash, uint256 validAt)
func (_MorphoDistributor *MorphoDistributorSession) PendingRoot() (struct {
	Root     [32]byte
	IpfsHash [32]byte
	ValidAt  *big.Int
}, error) {
	return _MorphoDistributor.Contract.PendingRoot(&_MorphoDistributor.CallOpts)
}

// PendingRoot is a free data retrieval call binding the contract method 0x750588cf.
//
// Solidity: function pendingRoot() view returns(bytes32 root, bytes32 ipfsHash, uint256 validAt)
func (_MorphoDistributor *MorphoDistributorCallerSession) PendingRoot() (struct {
	Root     [32]byte
	IpfsHash [32]byte
	ValidAt  *big.Int
}, error) {
	return _MorphoDistributor.Contract.PendingRoot(&_MorphoDistributor.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorSession) Root() ([32]byte, error) {
	return _MorphoDistributor.Contract.Root(&_MorphoDistributor.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MorphoDistributor *MorphoDistributorCallerSession) Root() ([32]byte, error) {
	return _MorphoDistributor.Contract.Root(&_MorphoDistributor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoDistributor *MorphoDistributorCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoDistributor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoDistributor *MorphoDistributorSession) Timelock() (*big.Int, error) {
	return _MorphoDistributor.Contract.Timelock(&_MorphoDistributor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoDistributor *MorphoDistributorCallerSession) Timelock() (*big.Int, error) {
	return _MorphoDistributor.Contract.Timelock(&_MorphoDistributor.CallOpts)
}

// AcceptRoot is a paid mutator transaction binding the contract method 0xed075ec5.
//
// Solidity: function acceptRoot() returns()
func (_MorphoDistributor *MorphoDistributorTransactor) AcceptRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "acceptRoot")
}

// AcceptRoot is a paid mutator transaction binding the contract method 0xed075ec5.
//
// Solidity: function acceptRoot() returns()
func (_MorphoDistributor *MorphoDistributorSession) AcceptRoot() (*types.Transaction, error) {
	return _MorphoDistributor.Contract.AcceptRoot(&_MorphoDistributor.TransactOpts)
}

// AcceptRoot is a paid mutator transaction binding the contract method 0xed075ec5.
//
// Solidity: function acceptRoot() returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) AcceptRoot() (*types.Transaction, error) {
	return _MorphoDistributor.Contract.AcceptRoot(&_MorphoDistributor.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xfabed412.
//
// Solidity: function claim(address account, address reward, uint256 claimable, bytes32[] proof) returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorTransactor) Claim(opts *bind.TransactOpts, account common.Address, reward common.Address, claimable *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "claim", account, reward, claimable, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xfabed412.
//
// Solidity: function claim(address account, address reward, uint256 claimable, bytes32[] proof) returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorSession) Claim(account common.Address, reward common.Address, claimable *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.Claim(&_MorphoDistributor.TransactOpts, account, reward, claimable, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xfabed412.
//
// Solidity: function claim(address account, address reward, uint256 claimable, bytes32[] proof) returns(uint256 amount)
func (_MorphoDistributor *MorphoDistributorTransactorSession) Claim(account common.Address, reward common.Address, claimable *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.Claim(&_MorphoDistributor.TransactOpts, account, reward, claimable, proof)
}

// RevokePendingRoot is a paid mutator transaction binding the contract method 0x4b387053.
//
// Solidity: function revokePendingRoot() returns()
func (_MorphoDistributor *MorphoDistributorTransactor) RevokePendingRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "revokePendingRoot")
}

// RevokePendingRoot is a paid mutator transaction binding the contract method 0x4b387053.
//
// Solidity: function revokePendingRoot() returns()
func (_MorphoDistributor *MorphoDistributorSession) RevokePendingRoot() (*types.Transaction, error) {
	return _MorphoDistributor.Contract.RevokePendingRoot(&_MorphoDistributor.TransactOpts)
}

// RevokePendingRoot is a paid mutator transaction binding the contract method 0x4b387053.
//
// Solidity: function revokePendingRoot() returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) RevokePendingRoot() (*types.Transaction, error) {
	return _MorphoDistributor.Contract.RevokePendingRoot(&_MorphoDistributor.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoDistributor *MorphoDistributorTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoDistributor *MorphoDistributorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetOwner(&_MorphoDistributor.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetOwner(&_MorphoDistributor.TransactOpts, newOwner)
}

// SetRoot is a paid mutator transaction binding the contract method 0x42af83fb.
//
// Solidity: function setRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorTransactor) SetRoot(opts *bind.TransactOpts, newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "setRoot", newRoot, newIpfsHash)
}

// SetRoot is a paid mutator transaction binding the contract method 0x42af83fb.
//
// Solidity: function setRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorSession) SetRoot(newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetRoot(&_MorphoDistributor.TransactOpts, newRoot, newIpfsHash)
}

// SetRoot is a paid mutator transaction binding the contract method 0x42af83fb.
//
// Solidity: function setRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) SetRoot(newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetRoot(&_MorphoDistributor.TransactOpts, newRoot, newIpfsHash)
}

// SetRootUpdater is a paid mutator transaction binding the contract method 0x53739410.
//
// Solidity: function setRootUpdater(address updater, bool active) returns()
func (_MorphoDistributor *MorphoDistributorTransactor) SetRootUpdater(opts *bind.TransactOpts, updater common.Address, active bool) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "setRootUpdater", updater, active)
}

// SetRootUpdater is a paid mutator transaction binding the contract method 0x53739410.
//
// Solidity: function setRootUpdater(address updater, bool active) returns()
func (_MorphoDistributor *MorphoDistributorSession) SetRootUpdater(updater common.Address, active bool) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetRootUpdater(&_MorphoDistributor.TransactOpts, updater, active)
}

// SetRootUpdater is a paid mutator transaction binding the contract method 0x53739410.
//
// Solidity: function setRootUpdater(address updater, bool active) returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) SetRootUpdater(updater common.Address, active bool) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetRootUpdater(&_MorphoDistributor.TransactOpts, updater, active)
}

// SetTimelock is a paid mutator transaction binding the contract method 0x1e891c0a.
//
// Solidity: function setTimelock(uint256 newTimelock) returns()
func (_MorphoDistributor *MorphoDistributorTransactor) SetTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "setTimelock", newTimelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0x1e891c0a.
//
// Solidity: function setTimelock(uint256 newTimelock) returns()
func (_MorphoDistributor *MorphoDistributorSession) SetTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetTimelock(&_MorphoDistributor.TransactOpts, newTimelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0x1e891c0a.
//
// Solidity: function setTimelock(uint256 newTimelock) returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) SetTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SetTimelock(&_MorphoDistributor.TransactOpts, newTimelock)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xd61825ef.
//
// Solidity: function submitRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorTransactor) SubmitRoot(opts *bind.TransactOpts, newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.contract.Transact(opts, "submitRoot", newRoot, newIpfsHash)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xd61825ef.
//
// Solidity: function submitRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorSession) SubmitRoot(newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SubmitRoot(&_MorphoDistributor.TransactOpts, newRoot, newIpfsHash)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xd61825ef.
//
// Solidity: function submitRoot(bytes32 newRoot, bytes32 newIpfsHash) returns()
func (_MorphoDistributor *MorphoDistributorTransactorSession) SubmitRoot(newRoot [32]byte, newIpfsHash [32]byte) (*types.Transaction, error) {
	return _MorphoDistributor.Contract.SubmitRoot(&_MorphoDistributor.TransactOpts, newRoot, newIpfsHash)
}
