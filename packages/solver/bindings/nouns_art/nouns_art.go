// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nouns_art

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

// INounsArtNounArtStoragePage is an auto generated low-level Go binding around an user-defined struct.
type INounsArtNounArtStoragePage struct {
	ImageCount         uint16
	DecompressedLength *big.Int
	Pointer            common.Address
}

// INounsArtTrait is an auto generated low-level Go binding around an user-defined struct.
type INounsArtTrait struct {
	StoragePages      []INounsArtNounArtStoragePage
	StoredImagesCount *big.Int
}

// NounsArtMetaData contains all meta data concerning the NounsArt contract.
var NounsArtMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_descriptor\",\"type\":\"address\"},{\"internalType\":\"contractIInflator\",\"name\":\"_inflator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadDecompressedLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadImageCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPaletteLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBytes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPalette\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaletteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotDescriptor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"AccessoriesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"AccessoriesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BackgroundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"BodiesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"BodiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDescriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDescriptor\",\"type\":\"address\"}],\"name\":\"DescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"GlassesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"GlassesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"HeadsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"}],\"name\":\"HeadsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldInflator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInflator\",\"type\":\"address\"}],\"name\":\"InflatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"paletteIndex\",\"type\":\"uint8\"}],\"name\":\"PaletteSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"accessories\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessoriesTrait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessoryCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addAccessoriesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_background\",\"type\":\"string\"}],\"name\":\"addBackground\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addBodies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addBodiesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addGlasses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addGlassesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addHeads\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"addHeadsFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_backgrounds\",\"type\":\"string[]\"}],\"name\":\"addManyBackgrounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backgroundCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"backgrounds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"bodies\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bodiesTrait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bodyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"descriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessoriesTrait\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"}],\"internalType\":\"structINounsArt.NounArtStoragePage[]\",\"name\":\"storagePages\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"internalType\":\"structINounsArt.Trait\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBodiesTrait\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"}],\"internalType\":\"structINounsArt.NounArtStoragePage[]\",\"name\":\"storagePages\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"internalType\":\"structINounsArt.Trait\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlassesTrait\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"}],\"internalType\":\"structINounsArt.NounArtStoragePage[]\",\"name\":\"storagePages\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"internalType\":\"structINounsArt.Trait\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHeadsTrait\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"}],\"internalType\":\"structINounsArt.NounArtStoragePage[]\",\"name\":\"storagePages\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"internalType\":\"structINounsArt.Trait\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"glasses\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glassesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glassesTrait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"heads\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headsTrait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"storedImagesCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflator\",\"outputs\":[{\"internalType\":\"contractIInflator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"paletteIndex\",\"type\":\"uint8\"}],\"name\":\"palettes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"palettesPointers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_descriptor\",\"type\":\"address\"}],\"name\":\"setDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInflator\",\"name\":\"_inflator\",\"type\":\"address\"}],\"name\":\"setInflator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"paletteIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"palette\",\"type\":\"bytes\"}],\"name\":\"setPalette\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"paletteIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"}],\"name\":\"setPalettePointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateAccessoriesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateBodies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateBodiesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateGlasses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateGlassesFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedCompressed\",\"type\":\"bytes\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateHeads\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pointer\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"decompressedLength\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"imageCount\",\"type\":\"uint16\"}],\"name\":\"updateHeadsFromPointer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NounsArtABI is the input ABI used to generate the binding from.
// Deprecated: Use NounsArtMetaData.ABI instead.
var NounsArtABI = NounsArtMetaData.ABI

// NounsArt is an auto generated Go binding around an Ethereum contract.
type NounsArt struct {
	NounsArtCaller     // Read-only binding to the contract
	NounsArtTransactor // Write-only binding to the contract
	NounsArtFilterer   // Log filterer for contract events
}

// NounsArtCaller is an auto generated read-only Go binding around an Ethereum contract.
type NounsArtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsArtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NounsArtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsArtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NounsArtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsArtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NounsArtSession struct {
	Contract     *NounsArt         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NounsArtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NounsArtCallerSession struct {
	Contract *NounsArtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NounsArtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NounsArtTransactorSession struct {
	Contract     *NounsArtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NounsArtRaw is an auto generated low-level Go binding around an Ethereum contract.
type NounsArtRaw struct {
	Contract *NounsArt // Generic contract binding to access the raw methods on
}

// NounsArtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NounsArtCallerRaw struct {
	Contract *NounsArtCaller // Generic read-only contract binding to access the raw methods on
}

// NounsArtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NounsArtTransactorRaw struct {
	Contract *NounsArtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNounsArt creates a new instance of NounsArt, bound to a specific deployed contract.
func NewNounsArt(address common.Address, backend bind.ContractBackend) (*NounsArt, error) {
	contract, err := bindNounsArt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NounsArt{NounsArtCaller: NounsArtCaller{contract: contract}, NounsArtTransactor: NounsArtTransactor{contract: contract}, NounsArtFilterer: NounsArtFilterer{contract: contract}}, nil
}

// NewNounsArtCaller creates a new read-only instance of NounsArt, bound to a specific deployed contract.
func NewNounsArtCaller(address common.Address, caller bind.ContractCaller) (*NounsArtCaller, error) {
	contract, err := bindNounsArt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NounsArtCaller{contract: contract}, nil
}

// NewNounsArtTransactor creates a new write-only instance of NounsArt, bound to a specific deployed contract.
func NewNounsArtTransactor(address common.Address, transactor bind.ContractTransactor) (*NounsArtTransactor, error) {
	contract, err := bindNounsArt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NounsArtTransactor{contract: contract}, nil
}

// NewNounsArtFilterer creates a new log filterer instance of NounsArt, bound to a specific deployed contract.
func NewNounsArtFilterer(address common.Address, filterer bind.ContractFilterer) (*NounsArtFilterer, error) {
	contract, err := bindNounsArt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NounsArtFilterer{contract: contract}, nil
}

// bindNounsArt binds a generic wrapper to an already deployed contract.
func bindNounsArt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NounsArtMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsArt *NounsArtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsArt.Contract.NounsArtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsArt *NounsArtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsArt.Contract.NounsArtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsArt *NounsArtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsArt.Contract.NounsArtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsArt *NounsArtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsArt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsArt *NounsArtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsArt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsArt *NounsArtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsArt.Contract.contract.Transact(opts, method, params...)
}

// Accessories is a free data retrieval call binding the contract method 0x7ca94210.
//
// Solidity: function accessories(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCaller) Accessories(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "accessories", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Accessories is a free data retrieval call binding the contract method 0x7ca94210.
//
// Solidity: function accessories(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtSession) Accessories(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Accessories(&_NounsArt.CallOpts, index)
}

// Accessories is a free data retrieval call binding the contract method 0x7ca94210.
//
// Solidity: function accessories(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCallerSession) Accessories(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Accessories(&_NounsArt.CallOpts, index)
}

// AccessoriesTrait is a free data retrieval call binding the contract method 0xbb5e488c.
//
// Solidity: function accessoriesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCaller) AccessoriesTrait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "accessoriesTrait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccessoriesTrait is a free data retrieval call binding the contract method 0xbb5e488c.
//
// Solidity: function accessoriesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtSession) AccessoriesTrait() (*big.Int, error) {
	return _NounsArt.Contract.AccessoriesTrait(&_NounsArt.CallOpts)
}

// AccessoriesTrait is a free data retrieval call binding the contract method 0xbb5e488c.
//
// Solidity: function accessoriesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCallerSession) AccessoriesTrait() (*big.Int, error) {
	return _NounsArt.Contract.AccessoriesTrait(&_NounsArt.CallOpts)
}

// AccessoryCount is a free data retrieval call binding the contract method 0x4daebac2.
//
// Solidity: function accessoryCount() view returns(uint256)
func (_NounsArt *NounsArtCaller) AccessoryCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "accessoryCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccessoryCount is a free data retrieval call binding the contract method 0x4daebac2.
//
// Solidity: function accessoryCount() view returns(uint256)
func (_NounsArt *NounsArtSession) AccessoryCount() (*big.Int, error) {
	return _NounsArt.Contract.AccessoryCount(&_NounsArt.CallOpts)
}

// AccessoryCount is a free data retrieval call binding the contract method 0x4daebac2.
//
// Solidity: function accessoryCount() view returns(uint256)
func (_NounsArt *NounsArtCallerSession) AccessoryCount() (*big.Int, error) {
	return _NounsArt.Contract.AccessoryCount(&_NounsArt.CallOpts)
}

// BackgroundCount is a free data retrieval call binding the contract method 0x4531c0a8.
//
// Solidity: function backgroundCount() view returns(uint256)
func (_NounsArt *NounsArtCaller) BackgroundCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "backgroundCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BackgroundCount is a free data retrieval call binding the contract method 0x4531c0a8.
//
// Solidity: function backgroundCount() view returns(uint256)
func (_NounsArt *NounsArtSession) BackgroundCount() (*big.Int, error) {
	return _NounsArt.Contract.BackgroundCount(&_NounsArt.CallOpts)
}

// BackgroundCount is a free data retrieval call binding the contract method 0x4531c0a8.
//
// Solidity: function backgroundCount() view returns(uint256)
func (_NounsArt *NounsArtCallerSession) BackgroundCount() (*big.Int, error) {
	return _NounsArt.Contract.BackgroundCount(&_NounsArt.CallOpts)
}

// Backgrounds is a free data retrieval call binding the contract method 0x04bde4dd.
//
// Solidity: function backgrounds(uint256 ) view returns(string)
func (_NounsArt *NounsArtCaller) Backgrounds(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "backgrounds", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Backgrounds is a free data retrieval call binding the contract method 0x04bde4dd.
//
// Solidity: function backgrounds(uint256 ) view returns(string)
func (_NounsArt *NounsArtSession) Backgrounds(arg0 *big.Int) (string, error) {
	return _NounsArt.Contract.Backgrounds(&_NounsArt.CallOpts, arg0)
}

// Backgrounds is a free data retrieval call binding the contract method 0x04bde4dd.
//
// Solidity: function backgrounds(uint256 ) view returns(string)
func (_NounsArt *NounsArtCallerSession) Backgrounds(arg0 *big.Int) (string, error) {
	return _NounsArt.Contract.Backgrounds(&_NounsArt.CallOpts, arg0)
}

// Bodies is a free data retrieval call binding the contract method 0x44cee73c.
//
// Solidity: function bodies(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCaller) Bodies(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "bodies", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bodies is a free data retrieval call binding the contract method 0x44cee73c.
//
// Solidity: function bodies(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtSession) Bodies(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Bodies(&_NounsArt.CallOpts, index)
}

// Bodies is a free data retrieval call binding the contract method 0x44cee73c.
//
// Solidity: function bodies(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCallerSession) Bodies(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Bodies(&_NounsArt.CallOpts, index)
}

// BodiesTrait is a free data retrieval call binding the contract method 0x5c0910be.
//
// Solidity: function bodiesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCaller) BodiesTrait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "bodiesTrait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BodiesTrait is a free data retrieval call binding the contract method 0x5c0910be.
//
// Solidity: function bodiesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtSession) BodiesTrait() (*big.Int, error) {
	return _NounsArt.Contract.BodiesTrait(&_NounsArt.CallOpts)
}

// BodiesTrait is a free data retrieval call binding the contract method 0x5c0910be.
//
// Solidity: function bodiesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCallerSession) BodiesTrait() (*big.Int, error) {
	return _NounsArt.Contract.BodiesTrait(&_NounsArt.CallOpts)
}

// BodyCount is a free data retrieval call binding the contract method 0xeba81806.
//
// Solidity: function bodyCount() view returns(uint256)
func (_NounsArt *NounsArtCaller) BodyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "bodyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BodyCount is a free data retrieval call binding the contract method 0xeba81806.
//
// Solidity: function bodyCount() view returns(uint256)
func (_NounsArt *NounsArtSession) BodyCount() (*big.Int, error) {
	return _NounsArt.Contract.BodyCount(&_NounsArt.CallOpts)
}

// BodyCount is a free data retrieval call binding the contract method 0xeba81806.
//
// Solidity: function bodyCount() view returns(uint256)
func (_NounsArt *NounsArtCallerSession) BodyCount() (*big.Int, error) {
	return _NounsArt.Contract.BodyCount(&_NounsArt.CallOpts)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsArt *NounsArtCaller) Descriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "descriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsArt *NounsArtSession) Descriptor() (common.Address, error) {
	return _NounsArt.Contract.Descriptor(&_NounsArt.CallOpts)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsArt *NounsArtCallerSession) Descriptor() (common.Address, error) {
	return _NounsArt.Contract.Descriptor(&_NounsArt.CallOpts)
}

// GetAccessoriesTrait is a free data retrieval call binding the contract method 0xc64b2f5d.
//
// Solidity: function getAccessoriesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCaller) GetAccessoriesTrait(opts *bind.CallOpts) (INounsArtTrait, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "getAccessoriesTrait")

	if err != nil {
		return *new(INounsArtTrait), err
	}

	out0 := *abi.ConvertType(out[0], new(INounsArtTrait)).(*INounsArtTrait)

	return out0, err

}

// GetAccessoriesTrait is a free data retrieval call binding the contract method 0xc64b2f5d.
//
// Solidity: function getAccessoriesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtSession) GetAccessoriesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetAccessoriesTrait(&_NounsArt.CallOpts)
}

// GetAccessoriesTrait is a free data retrieval call binding the contract method 0xc64b2f5d.
//
// Solidity: function getAccessoriesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCallerSession) GetAccessoriesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetAccessoriesTrait(&_NounsArt.CallOpts)
}

// GetBodiesTrait is a free data retrieval call binding the contract method 0x222a36d0.
//
// Solidity: function getBodiesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCaller) GetBodiesTrait(opts *bind.CallOpts) (INounsArtTrait, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "getBodiesTrait")

	if err != nil {
		return *new(INounsArtTrait), err
	}

	out0 := *abi.ConvertType(out[0], new(INounsArtTrait)).(*INounsArtTrait)

	return out0, err

}

// GetBodiesTrait is a free data retrieval call binding the contract method 0x222a36d0.
//
// Solidity: function getBodiesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtSession) GetBodiesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetBodiesTrait(&_NounsArt.CallOpts)
}

// GetBodiesTrait is a free data retrieval call binding the contract method 0x222a36d0.
//
// Solidity: function getBodiesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCallerSession) GetBodiesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetBodiesTrait(&_NounsArt.CallOpts)
}

// GetGlassesTrait is a free data retrieval call binding the contract method 0xe73dd383.
//
// Solidity: function getGlassesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCaller) GetGlassesTrait(opts *bind.CallOpts) (INounsArtTrait, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "getGlassesTrait")

	if err != nil {
		return *new(INounsArtTrait), err
	}

	out0 := *abi.ConvertType(out[0], new(INounsArtTrait)).(*INounsArtTrait)

	return out0, err

}

// GetGlassesTrait is a free data retrieval call binding the contract method 0xe73dd383.
//
// Solidity: function getGlassesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtSession) GetGlassesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetGlassesTrait(&_NounsArt.CallOpts)
}

// GetGlassesTrait is a free data retrieval call binding the contract method 0xe73dd383.
//
// Solidity: function getGlassesTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCallerSession) GetGlassesTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetGlassesTrait(&_NounsArt.CallOpts)
}

// GetHeadsTrait is a free data retrieval call binding the contract method 0x368013dc.
//
// Solidity: function getHeadsTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCaller) GetHeadsTrait(opts *bind.CallOpts) (INounsArtTrait, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "getHeadsTrait")

	if err != nil {
		return *new(INounsArtTrait), err
	}

	out0 := *abi.ConvertType(out[0], new(INounsArtTrait)).(*INounsArtTrait)

	return out0, err

}

// GetHeadsTrait is a free data retrieval call binding the contract method 0x368013dc.
//
// Solidity: function getHeadsTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtSession) GetHeadsTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetHeadsTrait(&_NounsArt.CallOpts)
}

// GetHeadsTrait is a free data retrieval call binding the contract method 0x368013dc.
//
// Solidity: function getHeadsTrait() view returns(((uint16,uint80,address)[],uint256))
func (_NounsArt *NounsArtCallerSession) GetHeadsTrait() (INounsArtTrait, error) {
	return _NounsArt.Contract.GetHeadsTrait(&_NounsArt.CallOpts)
}

// Glasses is a free data retrieval call binding the contract method 0xb982d1b9.
//
// Solidity: function glasses(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCaller) Glasses(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "glasses", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Glasses is a free data retrieval call binding the contract method 0xb982d1b9.
//
// Solidity: function glasses(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtSession) Glasses(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Glasses(&_NounsArt.CallOpts, index)
}

// Glasses is a free data retrieval call binding the contract method 0xb982d1b9.
//
// Solidity: function glasses(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCallerSession) Glasses(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Glasses(&_NounsArt.CallOpts, index)
}

// GlassesCount is a free data retrieval call binding the contract method 0x4479cef2.
//
// Solidity: function glassesCount() view returns(uint256)
func (_NounsArt *NounsArtCaller) GlassesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "glassesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlassesCount is a free data retrieval call binding the contract method 0x4479cef2.
//
// Solidity: function glassesCount() view returns(uint256)
func (_NounsArt *NounsArtSession) GlassesCount() (*big.Int, error) {
	return _NounsArt.Contract.GlassesCount(&_NounsArt.CallOpts)
}

// GlassesCount is a free data retrieval call binding the contract method 0x4479cef2.
//
// Solidity: function glassesCount() view returns(uint256)
func (_NounsArt *NounsArtCallerSession) GlassesCount() (*big.Int, error) {
	return _NounsArt.Contract.GlassesCount(&_NounsArt.CallOpts)
}

// GlassesTrait is a free data retrieval call binding the contract method 0xfc362a70.
//
// Solidity: function glassesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCaller) GlassesTrait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "glassesTrait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlassesTrait is a free data retrieval call binding the contract method 0xfc362a70.
//
// Solidity: function glassesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtSession) GlassesTrait() (*big.Int, error) {
	return _NounsArt.Contract.GlassesTrait(&_NounsArt.CallOpts)
}

// GlassesTrait is a free data retrieval call binding the contract method 0xfc362a70.
//
// Solidity: function glassesTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCallerSession) GlassesTrait() (*big.Int, error) {
	return _NounsArt.Contract.GlassesTrait(&_NounsArt.CallOpts)
}

// HeadCount is a free data retrieval call binding the contract method 0xcc2aa091.
//
// Solidity: function headCount() view returns(uint256)
func (_NounsArt *NounsArtCaller) HeadCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "headCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeadCount is a free data retrieval call binding the contract method 0xcc2aa091.
//
// Solidity: function headCount() view returns(uint256)
func (_NounsArt *NounsArtSession) HeadCount() (*big.Int, error) {
	return _NounsArt.Contract.HeadCount(&_NounsArt.CallOpts)
}

// HeadCount is a free data retrieval call binding the contract method 0xcc2aa091.
//
// Solidity: function headCount() view returns(uint256)
func (_NounsArt *NounsArtCallerSession) HeadCount() (*big.Int, error) {
	return _NounsArt.Contract.HeadCount(&_NounsArt.CallOpts)
}

// Heads is a free data retrieval call binding the contract method 0x5a503f13.
//
// Solidity: function heads(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCaller) Heads(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "heads", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Heads is a free data retrieval call binding the contract method 0x5a503f13.
//
// Solidity: function heads(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtSession) Heads(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Heads(&_NounsArt.CallOpts, index)
}

// Heads is a free data retrieval call binding the contract method 0x5a503f13.
//
// Solidity: function heads(uint256 index) view returns(bytes)
func (_NounsArt *NounsArtCallerSession) Heads(index *big.Int) ([]byte, error) {
	return _NounsArt.Contract.Heads(&_NounsArt.CallOpts, index)
}

// HeadsTrait is a free data retrieval call binding the contract method 0x970b2271.
//
// Solidity: function headsTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCaller) HeadsTrait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "headsTrait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeadsTrait is a free data retrieval call binding the contract method 0x970b2271.
//
// Solidity: function headsTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtSession) HeadsTrait() (*big.Int, error) {
	return _NounsArt.Contract.HeadsTrait(&_NounsArt.CallOpts)
}

// HeadsTrait is a free data retrieval call binding the contract method 0x970b2271.
//
// Solidity: function headsTrait() view returns(uint256 storedImagesCount)
func (_NounsArt *NounsArtCallerSession) HeadsTrait() (*big.Int, error) {
	return _NounsArt.Contract.HeadsTrait(&_NounsArt.CallOpts)
}

// Inflator is a free data retrieval call binding the contract method 0xe1d46ae6.
//
// Solidity: function inflator() view returns(address)
func (_NounsArt *NounsArtCaller) Inflator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "inflator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inflator is a free data retrieval call binding the contract method 0xe1d46ae6.
//
// Solidity: function inflator() view returns(address)
func (_NounsArt *NounsArtSession) Inflator() (common.Address, error) {
	return _NounsArt.Contract.Inflator(&_NounsArt.CallOpts)
}

// Inflator is a free data retrieval call binding the contract method 0xe1d46ae6.
//
// Solidity: function inflator() view returns(address)
func (_NounsArt *NounsArtCallerSession) Inflator() (common.Address, error) {
	return _NounsArt.Contract.Inflator(&_NounsArt.CallOpts)
}

// Palettes is a free data retrieval call binding the contract method 0xbc2d45fe.
//
// Solidity: function palettes(uint8 paletteIndex) view returns(bytes)
func (_NounsArt *NounsArtCaller) Palettes(opts *bind.CallOpts, paletteIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "palettes", paletteIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Palettes is a free data retrieval call binding the contract method 0xbc2d45fe.
//
// Solidity: function palettes(uint8 paletteIndex) view returns(bytes)
func (_NounsArt *NounsArtSession) Palettes(paletteIndex uint8) ([]byte, error) {
	return _NounsArt.Contract.Palettes(&_NounsArt.CallOpts, paletteIndex)
}

// Palettes is a free data retrieval call binding the contract method 0xbc2d45fe.
//
// Solidity: function palettes(uint8 paletteIndex) view returns(bytes)
func (_NounsArt *NounsArtCallerSession) Palettes(paletteIndex uint8) ([]byte, error) {
	return _NounsArt.Contract.Palettes(&_NounsArt.CallOpts, paletteIndex)
}

// PalettesPointers is a free data retrieval call binding the contract method 0x72c84d3f.
//
// Solidity: function palettesPointers(uint8 ) view returns(address)
func (_NounsArt *NounsArtCaller) PalettesPointers(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _NounsArt.contract.Call(opts, &out, "palettesPointers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PalettesPointers is a free data retrieval call binding the contract method 0x72c84d3f.
//
// Solidity: function palettesPointers(uint8 ) view returns(address)
func (_NounsArt *NounsArtSession) PalettesPointers(arg0 uint8) (common.Address, error) {
	return _NounsArt.Contract.PalettesPointers(&_NounsArt.CallOpts, arg0)
}

// PalettesPointers is a free data retrieval call binding the contract method 0x72c84d3f.
//
// Solidity: function palettesPointers(uint8 ) view returns(address)
func (_NounsArt *NounsArtCallerSession) PalettesPointers(arg0 uint8) (common.Address, error) {
	return _NounsArt.Contract.PalettesPointers(&_NounsArt.CallOpts, arg0)
}

// AddAccessories is a paid mutator transaction binding the contract method 0x0ba3db1a.
//
// Solidity: function addAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddAccessories(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addAccessories", encodedCompressed, decompressedLength, imageCount)
}

// AddAccessories is a paid mutator transaction binding the contract method 0x0ba3db1a.
//
// Solidity: function addAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddAccessories(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddAccessories(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddAccessories is a paid mutator transaction binding the contract method 0x0ba3db1a.
//
// Solidity: function addAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddAccessories(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddAccessories(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x6e856531.
//
// Solidity: function addAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddAccessoriesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addAccessoriesFromPointer", pointer, decompressedLength, imageCount)
}

// AddAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x6e856531.
//
// Solidity: function addAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddAccessoriesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddAccessoriesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x6e856531.
//
// Solidity: function addAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddAccessoriesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddAccessoriesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddBackground is a paid mutator transaction binding the contract method 0x5e70664c.
//
// Solidity: function addBackground(string _background) returns()
func (_NounsArt *NounsArtTransactor) AddBackground(opts *bind.TransactOpts, _background string) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addBackground", _background)
}

// AddBackground is a paid mutator transaction binding the contract method 0x5e70664c.
//
// Solidity: function addBackground(string _background) returns()
func (_NounsArt *NounsArtSession) AddBackground(_background string) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBackground(&_NounsArt.TransactOpts, _background)
}

// AddBackground is a paid mutator transaction binding the contract method 0x5e70664c.
//
// Solidity: function addBackground(string _background) returns()
func (_NounsArt *NounsArtTransactorSession) AddBackground(_background string) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBackground(&_NounsArt.TransactOpts, _background)
}

// AddBodies is a paid mutator transaction binding the contract method 0xaa5bf7d8.
//
// Solidity: function addBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddBodies(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addBodies", encodedCompressed, decompressedLength, imageCount)
}

// AddBodies is a paid mutator transaction binding the contract method 0xaa5bf7d8.
//
// Solidity: function addBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddBodies(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBodies(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddBodies is a paid mutator transaction binding the contract method 0xaa5bf7d8.
//
// Solidity: function addBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddBodies(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBodies(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddBodiesFromPointer is a paid mutator transaction binding the contract method 0xcd2b8250.
//
// Solidity: function addBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddBodiesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addBodiesFromPointer", pointer, decompressedLength, imageCount)
}

// AddBodiesFromPointer is a paid mutator transaction binding the contract method 0xcd2b8250.
//
// Solidity: function addBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddBodiesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBodiesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddBodiesFromPointer is a paid mutator transaction binding the contract method 0xcd2b8250.
//
// Solidity: function addBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddBodiesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddBodiesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddGlasses is a paid mutator transaction binding the contract method 0x353c36a0.
//
// Solidity: function addGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddGlasses(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addGlasses", encodedCompressed, decompressedLength, imageCount)
}

// AddGlasses is a paid mutator transaction binding the contract method 0x353c36a0.
//
// Solidity: function addGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddGlasses(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddGlasses(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddGlasses is a paid mutator transaction binding the contract method 0x353c36a0.
//
// Solidity: function addGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddGlasses(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddGlasses(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddGlassesFromPointer is a paid mutator transaction binding the contract method 0x73ac736b.
//
// Solidity: function addGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddGlassesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addGlassesFromPointer", pointer, decompressedLength, imageCount)
}

// AddGlassesFromPointer is a paid mutator transaction binding the contract method 0x73ac736b.
//
// Solidity: function addGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddGlassesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddGlassesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddGlassesFromPointer is a paid mutator transaction binding the contract method 0x73ac736b.
//
// Solidity: function addGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddGlassesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddGlassesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddHeads is a paid mutator transaction binding the contract method 0x94f3df61.
//
// Solidity: function addHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddHeads(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addHeads", encodedCompressed, decompressedLength, imageCount)
}

// AddHeads is a paid mutator transaction binding the contract method 0x94f3df61.
//
// Solidity: function addHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddHeads(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddHeads(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddHeads is a paid mutator transaction binding the contract method 0x94f3df61.
//
// Solidity: function addHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddHeads(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddHeads(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// AddHeadsFromPointer is a paid mutator transaction binding the contract method 0x461fc5af.
//
// Solidity: function addHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) AddHeadsFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addHeadsFromPointer", pointer, decompressedLength, imageCount)
}

// AddHeadsFromPointer is a paid mutator transaction binding the contract method 0x461fc5af.
//
// Solidity: function addHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) AddHeadsFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddHeadsFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddHeadsFromPointer is a paid mutator transaction binding the contract method 0x461fc5af.
//
// Solidity: function addHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) AddHeadsFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.AddHeadsFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// AddManyBackgrounds is a paid mutator transaction binding the contract method 0x91b7916a.
//
// Solidity: function addManyBackgrounds(string[] _backgrounds) returns()
func (_NounsArt *NounsArtTransactor) AddManyBackgrounds(opts *bind.TransactOpts, _backgrounds []string) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "addManyBackgrounds", _backgrounds)
}

// AddManyBackgrounds is a paid mutator transaction binding the contract method 0x91b7916a.
//
// Solidity: function addManyBackgrounds(string[] _backgrounds) returns()
func (_NounsArt *NounsArtSession) AddManyBackgrounds(_backgrounds []string) (*types.Transaction, error) {
	return _NounsArt.Contract.AddManyBackgrounds(&_NounsArt.TransactOpts, _backgrounds)
}

// AddManyBackgrounds is a paid mutator transaction binding the contract method 0x91b7916a.
//
// Solidity: function addManyBackgrounds(string[] _backgrounds) returns()
func (_NounsArt *NounsArtTransactorSession) AddManyBackgrounds(_backgrounds []string) (*types.Transaction, error) {
	return _NounsArt.Contract.AddManyBackgrounds(&_NounsArt.TransactOpts, _backgrounds)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsArt *NounsArtTransactor) SetDescriptor(opts *bind.TransactOpts, _descriptor common.Address) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "setDescriptor", _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsArt *NounsArtSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetDescriptor(&_NounsArt.TransactOpts, _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsArt *NounsArtTransactorSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetDescriptor(&_NounsArt.TransactOpts, _descriptor)
}

// SetInflator is a paid mutator transaction binding the contract method 0x72aa4a96.
//
// Solidity: function setInflator(address _inflator) returns()
func (_NounsArt *NounsArtTransactor) SetInflator(opts *bind.TransactOpts, _inflator common.Address) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "setInflator", _inflator)
}

// SetInflator is a paid mutator transaction binding the contract method 0x72aa4a96.
//
// Solidity: function setInflator(address _inflator) returns()
func (_NounsArt *NounsArtSession) SetInflator(_inflator common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetInflator(&_NounsArt.TransactOpts, _inflator)
}

// SetInflator is a paid mutator transaction binding the contract method 0x72aa4a96.
//
// Solidity: function setInflator(address _inflator) returns()
func (_NounsArt *NounsArtTransactorSession) SetInflator(_inflator common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetInflator(&_NounsArt.TransactOpts, _inflator)
}

// SetPalette is a paid mutator transaction binding the contract method 0xe79c9ea6.
//
// Solidity: function setPalette(uint8 paletteIndex, bytes palette) returns()
func (_NounsArt *NounsArtTransactor) SetPalette(opts *bind.TransactOpts, paletteIndex uint8, palette []byte) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "setPalette", paletteIndex, palette)
}

// SetPalette is a paid mutator transaction binding the contract method 0xe79c9ea6.
//
// Solidity: function setPalette(uint8 paletteIndex, bytes palette) returns()
func (_NounsArt *NounsArtSession) SetPalette(paletteIndex uint8, palette []byte) (*types.Transaction, error) {
	return _NounsArt.Contract.SetPalette(&_NounsArt.TransactOpts, paletteIndex, palette)
}

// SetPalette is a paid mutator transaction binding the contract method 0xe79c9ea6.
//
// Solidity: function setPalette(uint8 paletteIndex, bytes palette) returns()
func (_NounsArt *NounsArtTransactorSession) SetPalette(paletteIndex uint8, palette []byte) (*types.Transaction, error) {
	return _NounsArt.Contract.SetPalette(&_NounsArt.TransactOpts, paletteIndex, palette)
}

// SetPalettePointer is a paid mutator transaction binding the contract method 0x8bd54c06.
//
// Solidity: function setPalettePointer(uint8 paletteIndex, address pointer) returns()
func (_NounsArt *NounsArtTransactor) SetPalettePointer(opts *bind.TransactOpts, paletteIndex uint8, pointer common.Address) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "setPalettePointer", paletteIndex, pointer)
}

// SetPalettePointer is a paid mutator transaction binding the contract method 0x8bd54c06.
//
// Solidity: function setPalettePointer(uint8 paletteIndex, address pointer) returns()
func (_NounsArt *NounsArtSession) SetPalettePointer(paletteIndex uint8, pointer common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetPalettePointer(&_NounsArt.TransactOpts, paletteIndex, pointer)
}

// SetPalettePointer is a paid mutator transaction binding the contract method 0x8bd54c06.
//
// Solidity: function setPalettePointer(uint8 paletteIndex, address pointer) returns()
func (_NounsArt *NounsArtTransactorSession) SetPalettePointer(paletteIndex uint8, pointer common.Address) (*types.Transaction, error) {
	return _NounsArt.Contract.SetPalettePointer(&_NounsArt.TransactOpts, paletteIndex, pointer)
}

// UpdateAccessories is a paid mutator transaction binding the contract method 0xd6d7f8b6.
//
// Solidity: function updateAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateAccessories(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateAccessories", encodedCompressed, decompressedLength, imageCount)
}

// UpdateAccessories is a paid mutator transaction binding the contract method 0xd6d7f8b6.
//
// Solidity: function updateAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateAccessories(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateAccessories(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateAccessories is a paid mutator transaction binding the contract method 0xd6d7f8b6.
//
// Solidity: function updateAccessories(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateAccessories(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateAccessories(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x83f74116.
//
// Solidity: function updateAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateAccessoriesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateAccessoriesFromPointer", pointer, decompressedLength, imageCount)
}

// UpdateAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x83f74116.
//
// Solidity: function updateAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateAccessoriesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateAccessoriesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateAccessoriesFromPointer is a paid mutator transaction binding the contract method 0x83f74116.
//
// Solidity: function updateAccessoriesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateAccessoriesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateAccessoriesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateBodies is a paid mutator transaction binding the contract method 0xc0721cd6.
//
// Solidity: function updateBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateBodies(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateBodies", encodedCompressed, decompressedLength, imageCount)
}

// UpdateBodies is a paid mutator transaction binding the contract method 0xc0721cd6.
//
// Solidity: function updateBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateBodies(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateBodies(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateBodies is a paid mutator transaction binding the contract method 0xc0721cd6.
//
// Solidity: function updateBodies(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateBodies(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateBodies(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateBodiesFromPointer is a paid mutator transaction binding the contract method 0x7b425651.
//
// Solidity: function updateBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateBodiesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateBodiesFromPointer", pointer, decompressedLength, imageCount)
}

// UpdateBodiesFromPointer is a paid mutator transaction binding the contract method 0x7b425651.
//
// Solidity: function updateBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateBodiesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateBodiesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateBodiesFromPointer is a paid mutator transaction binding the contract method 0x7b425651.
//
// Solidity: function updateBodiesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateBodiesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateBodiesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateGlasses is a paid mutator transaction binding the contract method 0xcd5b95c8.
//
// Solidity: function updateGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateGlasses(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateGlasses", encodedCompressed, decompressedLength, imageCount)
}

// UpdateGlasses is a paid mutator transaction binding the contract method 0xcd5b95c8.
//
// Solidity: function updateGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateGlasses(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateGlasses(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateGlasses is a paid mutator transaction binding the contract method 0xcd5b95c8.
//
// Solidity: function updateGlasses(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateGlasses(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateGlasses(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateGlassesFromPointer is a paid mutator transaction binding the contract method 0xdbb55d2d.
//
// Solidity: function updateGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateGlassesFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateGlassesFromPointer", pointer, decompressedLength, imageCount)
}

// UpdateGlassesFromPointer is a paid mutator transaction binding the contract method 0xdbb55d2d.
//
// Solidity: function updateGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateGlassesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateGlassesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateGlassesFromPointer is a paid mutator transaction binding the contract method 0xdbb55d2d.
//
// Solidity: function updateGlassesFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateGlassesFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateGlassesFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateHeads is a paid mutator transaction binding the contract method 0xe4973050.
//
// Solidity: function updateHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateHeads(opts *bind.TransactOpts, encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateHeads", encodedCompressed, decompressedLength, imageCount)
}

// UpdateHeads is a paid mutator transaction binding the contract method 0xe4973050.
//
// Solidity: function updateHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateHeads(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateHeads(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateHeads is a paid mutator transaction binding the contract method 0xe4973050.
//
// Solidity: function updateHeads(bytes encodedCompressed, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateHeads(encodedCompressed []byte, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateHeads(&_NounsArt.TransactOpts, encodedCompressed, decompressedLength, imageCount)
}

// UpdateHeadsFromPointer is a paid mutator transaction binding the contract method 0x925b63a7.
//
// Solidity: function updateHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactor) UpdateHeadsFromPointer(opts *bind.TransactOpts, pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.contract.Transact(opts, "updateHeadsFromPointer", pointer, decompressedLength, imageCount)
}

// UpdateHeadsFromPointer is a paid mutator transaction binding the contract method 0x925b63a7.
//
// Solidity: function updateHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtSession) UpdateHeadsFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateHeadsFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// UpdateHeadsFromPointer is a paid mutator transaction binding the contract method 0x925b63a7.
//
// Solidity: function updateHeadsFromPointer(address pointer, uint80 decompressedLength, uint16 imageCount) returns()
func (_NounsArt *NounsArtTransactorSession) UpdateHeadsFromPointer(pointer common.Address, decompressedLength *big.Int, imageCount uint16) (*types.Transaction, error) {
	return _NounsArt.Contract.UpdateHeadsFromPointer(&_NounsArt.TransactOpts, pointer, decompressedLength, imageCount)
}

// NounsArtAccessoriesAddedIterator is returned from FilterAccessoriesAdded and is used to iterate over the raw logs and unpacked data for AccessoriesAdded events raised by the NounsArt contract.
type NounsArtAccessoriesAddedIterator struct {
	Event *NounsArtAccessoriesAdded // Event containing the contract specifics and raw log

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
func (it *NounsArtAccessoriesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtAccessoriesAdded)
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
		it.Event = new(NounsArtAccessoriesAdded)
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
func (it *NounsArtAccessoriesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtAccessoriesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtAccessoriesAdded represents a AccessoriesAdded event raised by the NounsArt contract.
type NounsArtAccessoriesAdded struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccessoriesAdded is a free log retrieval operation binding the contract event 0x2d0f2274b911553cd85de198b828dfe7ef8309c67e0b7674c045bcb0e5b5ba2e.
//
// Solidity: event AccessoriesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterAccessoriesAdded(opts *bind.FilterOpts) (*NounsArtAccessoriesAddedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "AccessoriesAdded")
	if err != nil {
		return nil, err
	}
	return &NounsArtAccessoriesAddedIterator{contract: _NounsArt.contract, event: "AccessoriesAdded", logs: logs, sub: sub}, nil
}

// WatchAccessoriesAdded is a free log subscription operation binding the contract event 0x2d0f2274b911553cd85de198b828dfe7ef8309c67e0b7674c045bcb0e5b5ba2e.
//
// Solidity: event AccessoriesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchAccessoriesAdded(opts *bind.WatchOpts, sink chan<- *NounsArtAccessoriesAdded) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "AccessoriesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtAccessoriesAdded)
				if err := _NounsArt.contract.UnpackLog(event, "AccessoriesAdded", log); err != nil {
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

// ParseAccessoriesAdded is a log parse operation binding the contract event 0x2d0f2274b911553cd85de198b828dfe7ef8309c67e0b7674c045bcb0e5b5ba2e.
//
// Solidity: event AccessoriesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseAccessoriesAdded(log types.Log) (*NounsArtAccessoriesAdded, error) {
	event := new(NounsArtAccessoriesAdded)
	if err := _NounsArt.contract.UnpackLog(event, "AccessoriesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtAccessoriesUpdatedIterator is returned from FilterAccessoriesUpdated and is used to iterate over the raw logs and unpacked data for AccessoriesUpdated events raised by the NounsArt contract.
type NounsArtAccessoriesUpdatedIterator struct {
	Event *NounsArtAccessoriesUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtAccessoriesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtAccessoriesUpdated)
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
		it.Event = new(NounsArtAccessoriesUpdated)
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
func (it *NounsArtAccessoriesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtAccessoriesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtAccessoriesUpdated represents a AccessoriesUpdated event raised by the NounsArt contract.
type NounsArtAccessoriesUpdated struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccessoriesUpdated is a free log retrieval operation binding the contract event 0x43a2ee835d7ab3047a2c726114a6e738684a8c2635356a461eb18ba63711ad51.
//
// Solidity: event AccessoriesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterAccessoriesUpdated(opts *bind.FilterOpts) (*NounsArtAccessoriesUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "AccessoriesUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtAccessoriesUpdatedIterator{contract: _NounsArt.contract, event: "AccessoriesUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessoriesUpdated is a free log subscription operation binding the contract event 0x43a2ee835d7ab3047a2c726114a6e738684a8c2635356a461eb18ba63711ad51.
//
// Solidity: event AccessoriesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchAccessoriesUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtAccessoriesUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "AccessoriesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtAccessoriesUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "AccessoriesUpdated", log); err != nil {
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

// ParseAccessoriesUpdated is a log parse operation binding the contract event 0x43a2ee835d7ab3047a2c726114a6e738684a8c2635356a461eb18ba63711ad51.
//
// Solidity: event AccessoriesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseAccessoriesUpdated(log types.Log) (*NounsArtAccessoriesUpdated, error) {
	event := new(NounsArtAccessoriesUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "AccessoriesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtBackgroundsAddedIterator is returned from FilterBackgroundsAdded and is used to iterate over the raw logs and unpacked data for BackgroundsAdded events raised by the NounsArt contract.
type NounsArtBackgroundsAddedIterator struct {
	Event *NounsArtBackgroundsAdded // Event containing the contract specifics and raw log

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
func (it *NounsArtBackgroundsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtBackgroundsAdded)
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
		it.Event = new(NounsArtBackgroundsAdded)
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
func (it *NounsArtBackgroundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtBackgroundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtBackgroundsAdded represents a BackgroundsAdded event raised by the NounsArt contract.
type NounsArtBackgroundsAdded struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBackgroundsAdded is a free log retrieval operation binding the contract event 0x379976e1287af3c12aafa34c6a1a61b0cbcb9dce67b3b220ece3b474a4a74276.
//
// Solidity: event BackgroundsAdded(uint256 count)
func (_NounsArt *NounsArtFilterer) FilterBackgroundsAdded(opts *bind.FilterOpts) (*NounsArtBackgroundsAddedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "BackgroundsAdded")
	if err != nil {
		return nil, err
	}
	return &NounsArtBackgroundsAddedIterator{contract: _NounsArt.contract, event: "BackgroundsAdded", logs: logs, sub: sub}, nil
}

// WatchBackgroundsAdded is a free log subscription operation binding the contract event 0x379976e1287af3c12aafa34c6a1a61b0cbcb9dce67b3b220ece3b474a4a74276.
//
// Solidity: event BackgroundsAdded(uint256 count)
func (_NounsArt *NounsArtFilterer) WatchBackgroundsAdded(opts *bind.WatchOpts, sink chan<- *NounsArtBackgroundsAdded) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "BackgroundsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtBackgroundsAdded)
				if err := _NounsArt.contract.UnpackLog(event, "BackgroundsAdded", log); err != nil {
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

// ParseBackgroundsAdded is a log parse operation binding the contract event 0x379976e1287af3c12aafa34c6a1a61b0cbcb9dce67b3b220ece3b474a4a74276.
//
// Solidity: event BackgroundsAdded(uint256 count)
func (_NounsArt *NounsArtFilterer) ParseBackgroundsAdded(log types.Log) (*NounsArtBackgroundsAdded, error) {
	event := new(NounsArtBackgroundsAdded)
	if err := _NounsArt.contract.UnpackLog(event, "BackgroundsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtBodiesAddedIterator is returned from FilterBodiesAdded and is used to iterate over the raw logs and unpacked data for BodiesAdded events raised by the NounsArt contract.
type NounsArtBodiesAddedIterator struct {
	Event *NounsArtBodiesAdded // Event containing the contract specifics and raw log

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
func (it *NounsArtBodiesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtBodiesAdded)
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
		it.Event = new(NounsArtBodiesAdded)
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
func (it *NounsArtBodiesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtBodiesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtBodiesAdded represents a BodiesAdded event raised by the NounsArt contract.
type NounsArtBodiesAdded struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBodiesAdded is a free log retrieval operation binding the contract event 0xeb09489df35ba64745f59c5a7efc6df50d432df8cfc3708deb7075e3c8a4f76a.
//
// Solidity: event BodiesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterBodiesAdded(opts *bind.FilterOpts) (*NounsArtBodiesAddedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "BodiesAdded")
	if err != nil {
		return nil, err
	}
	return &NounsArtBodiesAddedIterator{contract: _NounsArt.contract, event: "BodiesAdded", logs: logs, sub: sub}, nil
}

// WatchBodiesAdded is a free log subscription operation binding the contract event 0xeb09489df35ba64745f59c5a7efc6df50d432df8cfc3708deb7075e3c8a4f76a.
//
// Solidity: event BodiesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchBodiesAdded(opts *bind.WatchOpts, sink chan<- *NounsArtBodiesAdded) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "BodiesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtBodiesAdded)
				if err := _NounsArt.contract.UnpackLog(event, "BodiesAdded", log); err != nil {
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

// ParseBodiesAdded is a log parse operation binding the contract event 0xeb09489df35ba64745f59c5a7efc6df50d432df8cfc3708deb7075e3c8a4f76a.
//
// Solidity: event BodiesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseBodiesAdded(log types.Log) (*NounsArtBodiesAdded, error) {
	event := new(NounsArtBodiesAdded)
	if err := _NounsArt.contract.UnpackLog(event, "BodiesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtBodiesUpdatedIterator is returned from FilterBodiesUpdated and is used to iterate over the raw logs and unpacked data for BodiesUpdated events raised by the NounsArt contract.
type NounsArtBodiesUpdatedIterator struct {
	Event *NounsArtBodiesUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtBodiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtBodiesUpdated)
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
		it.Event = new(NounsArtBodiesUpdated)
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
func (it *NounsArtBodiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtBodiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtBodiesUpdated represents a BodiesUpdated event raised by the NounsArt contract.
type NounsArtBodiesUpdated struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBodiesUpdated is a free log retrieval operation binding the contract event 0x50e61f6fb788a0f606b23048f8a08e40e6c3eed4c2f871c149f3036fc42d95cd.
//
// Solidity: event BodiesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterBodiesUpdated(opts *bind.FilterOpts) (*NounsArtBodiesUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "BodiesUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtBodiesUpdatedIterator{contract: _NounsArt.contract, event: "BodiesUpdated", logs: logs, sub: sub}, nil
}

// WatchBodiesUpdated is a free log subscription operation binding the contract event 0x50e61f6fb788a0f606b23048f8a08e40e6c3eed4c2f871c149f3036fc42d95cd.
//
// Solidity: event BodiesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchBodiesUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtBodiesUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "BodiesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtBodiesUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "BodiesUpdated", log); err != nil {
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

// ParseBodiesUpdated is a log parse operation binding the contract event 0x50e61f6fb788a0f606b23048f8a08e40e6c3eed4c2f871c149f3036fc42d95cd.
//
// Solidity: event BodiesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseBodiesUpdated(log types.Log) (*NounsArtBodiesUpdated, error) {
	event := new(NounsArtBodiesUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "BodiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtDescriptorUpdatedIterator is returned from FilterDescriptorUpdated and is used to iterate over the raw logs and unpacked data for DescriptorUpdated events raised by the NounsArt contract.
type NounsArtDescriptorUpdatedIterator struct {
	Event *NounsArtDescriptorUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtDescriptorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtDescriptorUpdated)
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
		it.Event = new(NounsArtDescriptorUpdated)
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
func (it *NounsArtDescriptorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtDescriptorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtDescriptorUpdated represents a DescriptorUpdated event raised by the NounsArt contract.
type NounsArtDescriptorUpdated struct {
	OldDescriptor common.Address
	NewDescriptor common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDescriptorUpdated is a free log retrieval operation binding the contract event 0x6a470e5dd4b354979dc3b984575294975f737cb9ee3ae3cca949e998dbc7cee9.
//
// Solidity: event DescriptorUpdated(address oldDescriptor, address newDescriptor)
func (_NounsArt *NounsArtFilterer) FilterDescriptorUpdated(opts *bind.FilterOpts) (*NounsArtDescriptorUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtDescriptorUpdatedIterator{contract: _NounsArt.contract, event: "DescriptorUpdated", logs: logs, sub: sub}, nil
}

// WatchDescriptorUpdated is a free log subscription operation binding the contract event 0x6a470e5dd4b354979dc3b984575294975f737cb9ee3ae3cca949e998dbc7cee9.
//
// Solidity: event DescriptorUpdated(address oldDescriptor, address newDescriptor)
func (_NounsArt *NounsArtFilterer) WatchDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtDescriptorUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtDescriptorUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
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

// ParseDescriptorUpdated is a log parse operation binding the contract event 0x6a470e5dd4b354979dc3b984575294975f737cb9ee3ae3cca949e998dbc7cee9.
//
// Solidity: event DescriptorUpdated(address oldDescriptor, address newDescriptor)
func (_NounsArt *NounsArtFilterer) ParseDescriptorUpdated(log types.Log) (*NounsArtDescriptorUpdated, error) {
	event := new(NounsArtDescriptorUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtGlassesAddedIterator is returned from FilterGlassesAdded and is used to iterate over the raw logs and unpacked data for GlassesAdded events raised by the NounsArt contract.
type NounsArtGlassesAddedIterator struct {
	Event *NounsArtGlassesAdded // Event containing the contract specifics and raw log

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
func (it *NounsArtGlassesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtGlassesAdded)
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
		it.Event = new(NounsArtGlassesAdded)
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
func (it *NounsArtGlassesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtGlassesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtGlassesAdded represents a GlassesAdded event raised by the NounsArt contract.
type NounsArtGlassesAdded struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlassesAdded is a free log retrieval operation binding the contract event 0xfbb56d0e73d76edc5867b20b68684b671a625696e50d8c985c2830fd1566aaec.
//
// Solidity: event GlassesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterGlassesAdded(opts *bind.FilterOpts) (*NounsArtGlassesAddedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "GlassesAdded")
	if err != nil {
		return nil, err
	}
	return &NounsArtGlassesAddedIterator{contract: _NounsArt.contract, event: "GlassesAdded", logs: logs, sub: sub}, nil
}

// WatchGlassesAdded is a free log subscription operation binding the contract event 0xfbb56d0e73d76edc5867b20b68684b671a625696e50d8c985c2830fd1566aaec.
//
// Solidity: event GlassesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchGlassesAdded(opts *bind.WatchOpts, sink chan<- *NounsArtGlassesAdded) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "GlassesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtGlassesAdded)
				if err := _NounsArt.contract.UnpackLog(event, "GlassesAdded", log); err != nil {
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

// ParseGlassesAdded is a log parse operation binding the contract event 0xfbb56d0e73d76edc5867b20b68684b671a625696e50d8c985c2830fd1566aaec.
//
// Solidity: event GlassesAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseGlassesAdded(log types.Log) (*NounsArtGlassesAdded, error) {
	event := new(NounsArtGlassesAdded)
	if err := _NounsArt.contract.UnpackLog(event, "GlassesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtGlassesUpdatedIterator is returned from FilterGlassesUpdated and is used to iterate over the raw logs and unpacked data for GlassesUpdated events raised by the NounsArt contract.
type NounsArtGlassesUpdatedIterator struct {
	Event *NounsArtGlassesUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtGlassesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtGlassesUpdated)
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
		it.Event = new(NounsArtGlassesUpdated)
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
func (it *NounsArtGlassesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtGlassesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtGlassesUpdated represents a GlassesUpdated event raised by the NounsArt contract.
type NounsArtGlassesUpdated struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlassesUpdated is a free log retrieval operation binding the contract event 0x838d24ac2709c5e0923bc034d6e342aaa59b0bbb8138a133996e3cd6cd24f80c.
//
// Solidity: event GlassesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterGlassesUpdated(opts *bind.FilterOpts) (*NounsArtGlassesUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "GlassesUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtGlassesUpdatedIterator{contract: _NounsArt.contract, event: "GlassesUpdated", logs: logs, sub: sub}, nil
}

// WatchGlassesUpdated is a free log subscription operation binding the contract event 0x838d24ac2709c5e0923bc034d6e342aaa59b0bbb8138a133996e3cd6cd24f80c.
//
// Solidity: event GlassesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchGlassesUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtGlassesUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "GlassesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtGlassesUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "GlassesUpdated", log); err != nil {
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

// ParseGlassesUpdated is a log parse operation binding the contract event 0x838d24ac2709c5e0923bc034d6e342aaa59b0bbb8138a133996e3cd6cd24f80c.
//
// Solidity: event GlassesUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseGlassesUpdated(log types.Log) (*NounsArtGlassesUpdated, error) {
	event := new(NounsArtGlassesUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "GlassesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtHeadsAddedIterator is returned from FilterHeadsAdded and is used to iterate over the raw logs and unpacked data for HeadsAdded events raised by the NounsArt contract.
type NounsArtHeadsAddedIterator struct {
	Event *NounsArtHeadsAdded // Event containing the contract specifics and raw log

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
func (it *NounsArtHeadsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtHeadsAdded)
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
		it.Event = new(NounsArtHeadsAdded)
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
func (it *NounsArtHeadsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtHeadsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtHeadsAdded represents a HeadsAdded event raised by the NounsArt contract.
type NounsArtHeadsAdded struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHeadsAdded is a free log retrieval operation binding the contract event 0xe74953497d5d03198c809f0f4a324019e503e87fef8e2081636487743ae29d62.
//
// Solidity: event HeadsAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterHeadsAdded(opts *bind.FilterOpts) (*NounsArtHeadsAddedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "HeadsAdded")
	if err != nil {
		return nil, err
	}
	return &NounsArtHeadsAddedIterator{contract: _NounsArt.contract, event: "HeadsAdded", logs: logs, sub: sub}, nil
}

// WatchHeadsAdded is a free log subscription operation binding the contract event 0xe74953497d5d03198c809f0f4a324019e503e87fef8e2081636487743ae29d62.
//
// Solidity: event HeadsAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchHeadsAdded(opts *bind.WatchOpts, sink chan<- *NounsArtHeadsAdded) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "HeadsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtHeadsAdded)
				if err := _NounsArt.contract.UnpackLog(event, "HeadsAdded", log); err != nil {
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

// ParseHeadsAdded is a log parse operation binding the contract event 0xe74953497d5d03198c809f0f4a324019e503e87fef8e2081636487743ae29d62.
//
// Solidity: event HeadsAdded(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseHeadsAdded(log types.Log) (*NounsArtHeadsAdded, error) {
	event := new(NounsArtHeadsAdded)
	if err := _NounsArt.contract.UnpackLog(event, "HeadsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtHeadsUpdatedIterator is returned from FilterHeadsUpdated and is used to iterate over the raw logs and unpacked data for HeadsUpdated events raised by the NounsArt contract.
type NounsArtHeadsUpdatedIterator struct {
	Event *NounsArtHeadsUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtHeadsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtHeadsUpdated)
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
		it.Event = new(NounsArtHeadsUpdated)
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
func (it *NounsArtHeadsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtHeadsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtHeadsUpdated represents a HeadsUpdated event raised by the NounsArt contract.
type NounsArtHeadsUpdated struct {
	Count uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHeadsUpdated is a free log retrieval operation binding the contract event 0xb16d3cdff51a0e572b8c53ed281e6ff13c20c7851ee81c05df66be13ae786301.
//
// Solidity: event HeadsUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) FilterHeadsUpdated(opts *bind.FilterOpts) (*NounsArtHeadsUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "HeadsUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtHeadsUpdatedIterator{contract: _NounsArt.contract, event: "HeadsUpdated", logs: logs, sub: sub}, nil
}

// WatchHeadsUpdated is a free log subscription operation binding the contract event 0xb16d3cdff51a0e572b8c53ed281e6ff13c20c7851ee81c05df66be13ae786301.
//
// Solidity: event HeadsUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) WatchHeadsUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtHeadsUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "HeadsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtHeadsUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "HeadsUpdated", log); err != nil {
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

// ParseHeadsUpdated is a log parse operation binding the contract event 0xb16d3cdff51a0e572b8c53ed281e6ff13c20c7851ee81c05df66be13ae786301.
//
// Solidity: event HeadsUpdated(uint16 count)
func (_NounsArt *NounsArtFilterer) ParseHeadsUpdated(log types.Log) (*NounsArtHeadsUpdated, error) {
	event := new(NounsArtHeadsUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "HeadsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtInflatorUpdatedIterator is returned from FilterInflatorUpdated and is used to iterate over the raw logs and unpacked data for InflatorUpdated events raised by the NounsArt contract.
type NounsArtInflatorUpdatedIterator struct {
	Event *NounsArtInflatorUpdated // Event containing the contract specifics and raw log

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
func (it *NounsArtInflatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtInflatorUpdated)
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
		it.Event = new(NounsArtInflatorUpdated)
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
func (it *NounsArtInflatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtInflatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtInflatorUpdated represents a InflatorUpdated event raised by the NounsArt contract.
type NounsArtInflatorUpdated struct {
	OldInflator common.Address
	NewInflator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInflatorUpdated is a free log retrieval operation binding the contract event 0xad22bb31e9e983d055eeb60a03a1e572d4905254640c9ee3cd36c8d643124830.
//
// Solidity: event InflatorUpdated(address oldInflator, address newInflator)
func (_NounsArt *NounsArtFilterer) FilterInflatorUpdated(opts *bind.FilterOpts) (*NounsArtInflatorUpdatedIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "InflatorUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsArtInflatorUpdatedIterator{contract: _NounsArt.contract, event: "InflatorUpdated", logs: logs, sub: sub}, nil
}

// WatchInflatorUpdated is a free log subscription operation binding the contract event 0xad22bb31e9e983d055eeb60a03a1e572d4905254640c9ee3cd36c8d643124830.
//
// Solidity: event InflatorUpdated(address oldInflator, address newInflator)
func (_NounsArt *NounsArtFilterer) WatchInflatorUpdated(opts *bind.WatchOpts, sink chan<- *NounsArtInflatorUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "InflatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtInflatorUpdated)
				if err := _NounsArt.contract.UnpackLog(event, "InflatorUpdated", log); err != nil {
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

// ParseInflatorUpdated is a log parse operation binding the contract event 0xad22bb31e9e983d055eeb60a03a1e572d4905254640c9ee3cd36c8d643124830.
//
// Solidity: event InflatorUpdated(address oldInflator, address newInflator)
func (_NounsArt *NounsArtFilterer) ParseInflatorUpdated(log types.Log) (*NounsArtInflatorUpdated, error) {
	event := new(NounsArtInflatorUpdated)
	if err := _NounsArt.contract.UnpackLog(event, "InflatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsArtPaletteSetIterator is returned from FilterPaletteSet and is used to iterate over the raw logs and unpacked data for PaletteSet events raised by the NounsArt contract.
type NounsArtPaletteSetIterator struct {
	Event *NounsArtPaletteSet // Event containing the contract specifics and raw log

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
func (it *NounsArtPaletteSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsArtPaletteSet)
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
		it.Event = new(NounsArtPaletteSet)
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
func (it *NounsArtPaletteSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsArtPaletteSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsArtPaletteSet represents a PaletteSet event raised by the NounsArt contract.
type NounsArtPaletteSet struct {
	PaletteIndex uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaletteSet is a free log retrieval operation binding the contract event 0x3469f6a12aa5e5edc4ea6e284300f2621e073fce4374d4673ded8f2ea7c18b4f.
//
// Solidity: event PaletteSet(uint8 paletteIndex)
func (_NounsArt *NounsArtFilterer) FilterPaletteSet(opts *bind.FilterOpts) (*NounsArtPaletteSetIterator, error) {

	logs, sub, err := _NounsArt.contract.FilterLogs(opts, "PaletteSet")
	if err != nil {
		return nil, err
	}
	return &NounsArtPaletteSetIterator{contract: _NounsArt.contract, event: "PaletteSet", logs: logs, sub: sub}, nil
}

// WatchPaletteSet is a free log subscription operation binding the contract event 0x3469f6a12aa5e5edc4ea6e284300f2621e073fce4374d4673ded8f2ea7c18b4f.
//
// Solidity: event PaletteSet(uint8 paletteIndex)
func (_NounsArt *NounsArtFilterer) WatchPaletteSet(opts *bind.WatchOpts, sink chan<- *NounsArtPaletteSet) (event.Subscription, error) {

	logs, sub, err := _NounsArt.contract.WatchLogs(opts, "PaletteSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsArtPaletteSet)
				if err := _NounsArt.contract.UnpackLog(event, "PaletteSet", log); err != nil {
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

// ParsePaletteSet is a log parse operation binding the contract event 0x3469f6a12aa5e5edc4ea6e284300f2621e073fce4374d4673ded8f2ea7c18b4f.
//
// Solidity: event PaletteSet(uint8 paletteIndex)
func (_NounsArt *NounsArtFilterer) ParsePaletteSet(log types.Log) (*NounsArtPaletteSet, error) {
	event := new(NounsArtPaletteSet)
	if err := _NounsArt.contract.UnpackLog(event, "PaletteSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
