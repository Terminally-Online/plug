// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package othentic_attestation

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Operator    common.Address
	OperatorId  *big.Int
	VotingPower *big.Int
}

// OthenticAttestationMetaData contains all meta data concerning the OthenticAttestation contract.
var OthenticAttestationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getActiveOperatorsDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"struct\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OthenticAttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use OthenticAttestationMetaData.ABI instead.
var OthenticAttestationABI = OthenticAttestationMetaData.ABI

// OthenticAttestation is an auto generated Go binding around an Ethereum contract.
type OthenticAttestation struct {
	OthenticAttestationCaller     // Read-only binding to the contract
	OthenticAttestationTransactor // Write-only binding to the contract
	OthenticAttestationFilterer   // Log filterer for contract events
}

// OthenticAttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type OthenticAttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticAttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OthenticAttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticAttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OthenticAttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticAttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OthenticAttestationSession struct {
	Contract     *OthenticAttestation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OthenticAttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OthenticAttestationCallerSession struct {
	Contract *OthenticAttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OthenticAttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OthenticAttestationTransactorSession struct {
	Contract     *OthenticAttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OthenticAttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type OthenticAttestationRaw struct {
	Contract *OthenticAttestation // Generic contract binding to access the raw methods on
}

// OthenticAttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OthenticAttestationCallerRaw struct {
	Contract *OthenticAttestationCaller // Generic read-only contract binding to access the raw methods on
}

// OthenticAttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OthenticAttestationTransactorRaw struct {
	Contract *OthenticAttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOthenticAttestation creates a new instance of OthenticAttestation, bound to a specific deployed contract.
func NewOthenticAttestation(address common.Address, backend bind.ContractBackend) (*OthenticAttestation, error) {
	contract, err := bindOthenticAttestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OthenticAttestation{OthenticAttestationCaller: OthenticAttestationCaller{contract: contract}, OthenticAttestationTransactor: OthenticAttestationTransactor{contract: contract}, OthenticAttestationFilterer: OthenticAttestationFilterer{contract: contract}}, nil
}

// NewOthenticAttestationCaller creates a new read-only instance of OthenticAttestation, bound to a specific deployed contract.
func NewOthenticAttestationCaller(address common.Address, caller bind.ContractCaller) (*OthenticAttestationCaller, error) {
	contract, err := bindOthenticAttestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OthenticAttestationCaller{contract: contract}, nil
}

// NewOthenticAttestationTransactor creates a new write-only instance of OthenticAttestation, bound to a specific deployed contract.
func NewOthenticAttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*OthenticAttestationTransactor, error) {
	contract, err := bindOthenticAttestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OthenticAttestationTransactor{contract: contract}, nil
}

// NewOthenticAttestationFilterer creates a new log filterer instance of OthenticAttestation, bound to a specific deployed contract.
func NewOthenticAttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*OthenticAttestationFilterer, error) {
	contract, err := bindOthenticAttestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OthenticAttestationFilterer{contract: contract}, nil
}

// bindOthenticAttestation binds a generic wrapper to an already deployed contract.
func bindOthenticAttestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OthenticAttestationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OthenticAttestation *OthenticAttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OthenticAttestation.Contract.OthenticAttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OthenticAttestation *OthenticAttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OthenticAttestation.Contract.OthenticAttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OthenticAttestation *OthenticAttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OthenticAttestation.Contract.OthenticAttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OthenticAttestation *OthenticAttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OthenticAttestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OthenticAttestation *OthenticAttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OthenticAttestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OthenticAttestation *OthenticAttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OthenticAttestation.Contract.contract.Transact(opts, method, params...)
}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[])
func (_OthenticAttestation *OthenticAttestationCaller) GetActiveOperatorsDetails(opts *bind.CallOpts) ([]Struct0, error) {
	var out []interface{}
	err := _OthenticAttestation.contract.Call(opts, &out, "getActiveOperatorsDetails")

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[])
func (_OthenticAttestation *OthenticAttestationSession) GetActiveOperatorsDetails() ([]Struct0, error) {
	return _OthenticAttestation.Contract.GetActiveOperatorsDetails(&_OthenticAttestation.CallOpts)
}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[])
func (_OthenticAttestation *OthenticAttestationCallerSession) GetActiveOperatorsDetails() ([]Struct0, error) {
	return _OthenticAttestation.Contract.GetActiveOperatorsDetails(&_OthenticAttestation.CallOpts)
}
