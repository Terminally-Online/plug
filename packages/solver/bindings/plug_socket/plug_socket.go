// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_socket

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

// PlugTypesLibEIP712Domain is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibEIP712Domain struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}

// PlugTypesLibLivePlugs is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibLivePlugs struct {
	Plugs     PlugTypesLibPlugs
	Signature []byte
}

// PlugTypesLibPlug is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibPlug struct {
	Selector uint8
	To       common.Address
	Data     []byte
	Value    *big.Int
	Updates  []PlugTypesLibUpdate
}

// PlugTypesLibPlugs is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibPlugs struct {
	Socket common.Address
	Plugs  []PlugTypesLibPlug
	Solver []byte
	Salt   []byte
}

// PlugTypesLibResult is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibResult struct {
	Index uint8
	Error string
}

// PlugTypesLibSlice is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibSlice struct {
	Index  uint8
	Start  *big.Int
	Length *big.Int
}

// PlugTypesLibUpdate is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibUpdate struct {
	Start *big.Int
	Slice PlugTypesLibSlice
}

// PlugSocketMetaData contains all meta data concerning the PlugSocket contract.
var PlugSocketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"domain\",\"inputs\":[],\"outputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.EIP712Domain\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEIP712DomainHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.EIP712Domain\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLivePlugsHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.LivePlugs\",\"components\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLivePlugsSigner\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.LivePlugs\",\"components\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlugArrayHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPlugHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plug\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPlugsDigest\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlugsHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSliceHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getUpdateArrayHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getUpdateHash\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Update\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"typeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oneClicker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"oneClick\",\"inputs\":[{\"name\":\"oneClickers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowance\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oneClickersToAllowed\",\"inputs\":[{\"name\":\"oneClicker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"plug\",\"inputs\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Result\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"plug\",\"inputs\":[{\"name\":\"livePlugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.LivePlugs\",\"components\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"selector\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Update[]\",\"components\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slice\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Slice\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Result\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FnSelectorNotRecognized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PlugFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpgradeFailed\",\"inputs\":[]}]",
}

// PlugSocketABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugSocketMetaData.ABI instead.
var PlugSocketABI = PlugSocketMetaData.ABI

// PlugSocket is an auto generated Go binding around an Ethereum contract.
type PlugSocket struct {
	PlugSocketCaller     // Read-only binding to the contract
	PlugSocketTransactor // Write-only binding to the contract
	PlugSocketFilterer   // Log filterer for contract events
}

// PlugSocketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugSocketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugSocketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugSocketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugSocketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugSocketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugSocketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugSocketSession struct {
	Contract     *PlugSocket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugSocketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugSocketCallerSession struct {
	Contract *PlugSocketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlugSocketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugSocketTransactorSession struct {
	Contract     *PlugSocketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlugSocketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugSocketRaw struct {
	Contract *PlugSocket // Generic contract binding to access the raw methods on
}

// PlugSocketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugSocketCallerRaw struct {
	Contract *PlugSocketCaller // Generic read-only contract binding to access the raw methods on
}

// PlugSocketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugSocketTransactorRaw struct {
	Contract *PlugSocketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugSocket creates a new instance of PlugSocket, bound to a specific deployed contract.
func NewPlugSocket(address common.Address, backend bind.ContractBackend) (*PlugSocket, error) {
	contract, err := bindPlugSocket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugSocket{PlugSocketCaller: PlugSocketCaller{contract: contract}, PlugSocketTransactor: PlugSocketTransactor{contract: contract}, PlugSocketFilterer: PlugSocketFilterer{contract: contract}}, nil
}

// NewPlugSocketCaller creates a new read-only instance of PlugSocket, bound to a specific deployed contract.
func NewPlugSocketCaller(address common.Address, caller bind.ContractCaller) (*PlugSocketCaller, error) {
	contract, err := bindPlugSocket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugSocketCaller{contract: contract}, nil
}

// NewPlugSocketTransactor creates a new write-only instance of PlugSocket, bound to a specific deployed contract.
func NewPlugSocketTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugSocketTransactor, error) {
	contract, err := bindPlugSocket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugSocketTransactor{contract: contract}, nil
}

// NewPlugSocketFilterer creates a new log filterer instance of PlugSocket, bound to a specific deployed contract.
func NewPlugSocketFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugSocketFilterer, error) {
	contract, err := bindPlugSocket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugSocketFilterer{contract: contract}, nil
}

// bindPlugSocket binds a generic wrapper to an already deployed contract.
func bindPlugSocket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugSocketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugSocket *PlugSocketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugSocket.Contract.PlugSocketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugSocket *PlugSocketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.Contract.PlugSocketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugSocket *PlugSocketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugSocket.Contract.PlugSocketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugSocket *PlugSocketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugSocket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugSocket *PlugSocketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugSocket *PlugSocketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugSocket.Contract.contract.Transact(opts, method, params...)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns((string,string,uint256,address) domain)
func (_PlugSocket *PlugSocketCaller) Domain(opts *bind.CallOpts) (PlugTypesLibEIP712Domain, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "domain")

	if err != nil {
		return *new(PlugTypesLibEIP712Domain), err
	}

	out0 := *abi.ConvertType(out[0], new(PlugTypesLibEIP712Domain)).(*PlugTypesLibEIP712Domain)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns((string,string,uint256,address) domain)
func (_PlugSocket *PlugSocketSession) Domain() (PlugTypesLibEIP712Domain, error) {
	return _PlugSocket.Contract.Domain(&_PlugSocket.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns((string,string,uint256,address) domain)
func (_PlugSocket *PlugSocketCallerSession) Domain() (PlugTypesLibEIP712Domain, error) {
	return _PlugSocket.Contract.Domain(&_PlugSocket.CallOpts)
}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_PlugSocket *PlugSocketCaller) DomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "domainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_PlugSocket *PlugSocketSession) DomainHash() ([32]byte, error) {
	return _PlugSocket.Contract.DomainHash(&_PlugSocket.CallOpts)
}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_PlugSocket *PlugSocketCallerSession) DomainHash() ([32]byte, error) {
	return _PlugSocket.Contract.DomainHash(&_PlugSocket.CallOpts)
}

// GetEIP712DomainHash is a free data retrieval call binding the contract method 0x34bb3700.
//
// Solidity: function getEIP712DomainHash((string,string,uint256,address) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetEIP712DomainHash(opts *bind.CallOpts, input PlugTypesLibEIP712Domain) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getEIP712DomainHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEIP712DomainHash is a free data retrieval call binding the contract method 0x34bb3700.
//
// Solidity: function getEIP712DomainHash((string,string,uint256,address) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetEIP712DomainHash(input PlugTypesLibEIP712Domain) ([32]byte, error) {
	return _PlugSocket.Contract.GetEIP712DomainHash(&_PlugSocket.CallOpts, input)
}

// GetEIP712DomainHash is a free data retrieval call binding the contract method 0x34bb3700.
//
// Solidity: function getEIP712DomainHash((string,string,uint256,address) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetEIP712DomainHash(input PlugTypesLibEIP712Domain) ([32]byte, error) {
	return _PlugSocket.Contract.GetEIP712DomainHash(&_PlugSocket.CallOpts, input)
}

// GetLivePlugsHash is a free data retrieval call binding the contract method 0xb51f81f6.
//
// Solidity: function getLivePlugsHash(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetLivePlugsHash(opts *bind.CallOpts, input PlugTypesLibLivePlugs) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getLivePlugsHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLivePlugsHash is a free data retrieval call binding the contract method 0xb51f81f6.
//
// Solidity: function getLivePlugsHash(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetLivePlugsHash(input PlugTypesLibLivePlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetLivePlugsHash(&_PlugSocket.CallOpts, input)
}

// GetLivePlugsHash is a free data retrieval call binding the contract method 0xb51f81f6.
//
// Solidity: function getLivePlugsHash(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetLivePlugsHash(input PlugTypesLibLivePlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetLivePlugsHash(&_PlugSocket.CallOpts, input)
}

// GetLivePlugsSigner is a free data retrieval call binding the contract method 0x190dc9ac.
//
// Solidity: function getLivePlugsSigner(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) view returns(address signer)
func (_PlugSocket *PlugSocketCaller) GetLivePlugsSigner(opts *bind.CallOpts, input PlugTypesLibLivePlugs) (common.Address, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getLivePlugsSigner", input)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLivePlugsSigner is a free data retrieval call binding the contract method 0x190dc9ac.
//
// Solidity: function getLivePlugsSigner(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) view returns(address signer)
func (_PlugSocket *PlugSocketSession) GetLivePlugsSigner(input PlugTypesLibLivePlugs) (common.Address, error) {
	return _PlugSocket.Contract.GetLivePlugsSigner(&_PlugSocket.CallOpts, input)
}

// GetLivePlugsSigner is a free data retrieval call binding the contract method 0x190dc9ac.
//
// Solidity: function getLivePlugsSigner(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) input) view returns(address signer)
func (_PlugSocket *PlugSocketCallerSession) GetLivePlugsSigner(input PlugTypesLibLivePlugs) (common.Address, error) {
	return _PlugSocket.Contract.GetLivePlugsSigner(&_PlugSocket.CallOpts, input)
}

// GetPlugArrayHash is a free data retrieval call binding the contract method 0xe450bc51.
//
// Solidity: function getPlugArrayHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetPlugArrayHash(opts *bind.CallOpts, input []PlugTypesLibPlug) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getPlugArrayHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPlugArrayHash is a free data retrieval call binding the contract method 0xe450bc51.
//
// Solidity: function getPlugArrayHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetPlugArrayHash(input []PlugTypesLibPlug) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugArrayHash(&_PlugSocket.CallOpts, input)
}

// GetPlugArrayHash is a free data retrieval call binding the contract method 0xe450bc51.
//
// Solidity: function getPlugArrayHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetPlugArrayHash(input []PlugTypesLibPlug) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugArrayHash(&_PlugSocket.CallOpts, input)
}

// GetPlugHash is a free data retrieval call binding the contract method 0x06d2d460.
//
// Solidity: function getPlugHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[]) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetPlugHash(opts *bind.CallOpts, input PlugTypesLibPlug) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getPlugHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPlugHash is a free data retrieval call binding the contract method 0x06d2d460.
//
// Solidity: function getPlugHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[]) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetPlugHash(input PlugTypesLibPlug) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugHash(&_PlugSocket.CallOpts, input)
}

// GetPlugHash is a free data retrieval call binding the contract method 0x06d2d460.
//
// Solidity: function getPlugHash((uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[]) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetPlugHash(input PlugTypesLibPlug) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugHash(&_PlugSocket.CallOpts, input)
}

// GetPlugsDigest is a free data retrieval call binding the contract method 0x17943105.
//
// Solidity: function getPlugsDigest((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) view returns(bytes32 digest)
func (_PlugSocket *PlugSocketCaller) GetPlugsDigest(opts *bind.CallOpts, input PlugTypesLibPlugs) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getPlugsDigest", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPlugsDigest is a free data retrieval call binding the contract method 0x17943105.
//
// Solidity: function getPlugsDigest((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) view returns(bytes32 digest)
func (_PlugSocket *PlugSocketSession) GetPlugsDigest(input PlugTypesLibPlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugsDigest(&_PlugSocket.CallOpts, input)
}

// GetPlugsDigest is a free data retrieval call binding the contract method 0x17943105.
//
// Solidity: function getPlugsDigest((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) view returns(bytes32 digest)
func (_PlugSocket *PlugSocketCallerSession) GetPlugsDigest(input PlugTypesLibPlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugsDigest(&_PlugSocket.CallOpts, input)
}

// GetPlugsHash is a free data retrieval call binding the contract method 0xe75e0ad6.
//
// Solidity: function getPlugsHash((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetPlugsHash(opts *bind.CallOpts, input PlugTypesLibPlugs) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getPlugsHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPlugsHash is a free data retrieval call binding the contract method 0xe75e0ad6.
//
// Solidity: function getPlugsHash((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetPlugsHash(input PlugTypesLibPlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugsHash(&_PlugSocket.CallOpts, input)
}

// GetPlugsHash is a free data retrieval call binding the contract method 0xe75e0ad6.
//
// Solidity: function getPlugsHash((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetPlugsHash(input PlugTypesLibPlugs) ([32]byte, error) {
	return _PlugSocket.Contract.GetPlugsHash(&_PlugSocket.CallOpts, input)
}

// GetSliceHash is a free data retrieval call binding the contract method 0x161937ea.
//
// Solidity: function getSliceHash((uint8,uint256,uint256) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetSliceHash(opts *bind.CallOpts, input PlugTypesLibSlice) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getSliceHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSliceHash is a free data retrieval call binding the contract method 0x161937ea.
//
// Solidity: function getSliceHash((uint8,uint256,uint256) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetSliceHash(input PlugTypesLibSlice) ([32]byte, error) {
	return _PlugSocket.Contract.GetSliceHash(&_PlugSocket.CallOpts, input)
}

// GetSliceHash is a free data retrieval call binding the contract method 0x161937ea.
//
// Solidity: function getSliceHash((uint8,uint256,uint256) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetSliceHash(input PlugTypesLibSlice) ([32]byte, error) {
	return _PlugSocket.Contract.GetSliceHash(&_PlugSocket.CallOpts, input)
}

// GetUpdateArrayHash is a free data retrieval call binding the contract method 0x0b21e86f.
//
// Solidity: function getUpdateArrayHash((uint256,(uint8,uint256,uint256))[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetUpdateArrayHash(opts *bind.CallOpts, input []PlugTypesLibUpdate) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getUpdateArrayHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUpdateArrayHash is a free data retrieval call binding the contract method 0x0b21e86f.
//
// Solidity: function getUpdateArrayHash((uint256,(uint8,uint256,uint256))[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetUpdateArrayHash(input []PlugTypesLibUpdate) ([32]byte, error) {
	return _PlugSocket.Contract.GetUpdateArrayHash(&_PlugSocket.CallOpts, input)
}

// GetUpdateArrayHash is a free data retrieval call binding the contract method 0x0b21e86f.
//
// Solidity: function getUpdateArrayHash((uint256,(uint8,uint256,uint256))[] input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetUpdateArrayHash(input []PlugTypesLibUpdate) ([32]byte, error) {
	return _PlugSocket.Contract.GetUpdateArrayHash(&_PlugSocket.CallOpts, input)
}

// GetUpdateHash is a free data retrieval call binding the contract method 0xb2f65adb.
//
// Solidity: function getUpdateHash((uint256,(uint8,uint256,uint256)) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCaller) GetUpdateHash(opts *bind.CallOpts, input PlugTypesLibUpdate) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "getUpdateHash", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUpdateHash is a free data retrieval call binding the contract method 0xb2f65adb.
//
// Solidity: function getUpdateHash((uint256,(uint8,uint256,uint256)) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketSession) GetUpdateHash(input PlugTypesLibUpdate) ([32]byte, error) {
	return _PlugSocket.Contract.GetUpdateHash(&_PlugSocket.CallOpts, input)
}

// GetUpdateHash is a free data retrieval call binding the contract method 0xb2f65adb.
//
// Solidity: function getUpdateHash((uint256,(uint8,uint256,uint256)) input) pure returns(bytes32 typeHash)
func (_PlugSocket *PlugSocketCallerSession) GetUpdateHash(input PlugTypesLibUpdate) ([32]byte, error) {
	return _PlugSocket.Contract.GetUpdateHash(&_PlugSocket.CallOpts, input)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugSocket *PlugSocketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugSocket *PlugSocketSession) Name() (string, error) {
	return _PlugSocket.Contract.Name(&_PlugSocket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugSocket *PlugSocketCallerSession) Name() (string, error) {
	return _PlugSocket.Contract.Name(&_PlugSocket.CallOpts)
}

// OneClickersToAllowed is a free data retrieval call binding the contract method 0x60cd27ad.
//
// Solidity: function oneClickersToAllowed(address oneClicker) view returns(bool allowed)
func (_PlugSocket *PlugSocketCaller) OneClickersToAllowed(opts *bind.CallOpts, oneClicker common.Address) (bool, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "oneClickersToAllowed", oneClicker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OneClickersToAllowed is a free data retrieval call binding the contract method 0x60cd27ad.
//
// Solidity: function oneClickersToAllowed(address oneClicker) view returns(bool allowed)
func (_PlugSocket *PlugSocketSession) OneClickersToAllowed(oneClicker common.Address) (bool, error) {
	return _PlugSocket.Contract.OneClickersToAllowed(&_PlugSocket.CallOpts, oneClicker)
}

// OneClickersToAllowed is a free data retrieval call binding the contract method 0x60cd27ad.
//
// Solidity: function oneClickersToAllowed(address oneClicker) view returns(bool allowed)
func (_PlugSocket *PlugSocketCallerSession) OneClickersToAllowed(oneClicker common.Address) (bool, error) {
	return _PlugSocket.Contract.OneClickersToAllowed(&_PlugSocket.CallOpts, oneClicker)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PlugSocket *PlugSocketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PlugSocket *PlugSocketSession) Owner() (common.Address, error) {
	return _PlugSocket.Contract.Owner(&_PlugSocket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PlugSocket *PlugSocketCallerSession) Owner() (common.Address, error) {
	return _PlugSocket.Contract.Owner(&_PlugSocket.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PlugSocket *PlugSocketCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PlugSocket *PlugSocketSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PlugSocket.Contract.OwnershipHandoverExpiresAt(&_PlugSocket.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PlugSocket *PlugSocketCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PlugSocket.Contract.OwnershipHandoverExpiresAt(&_PlugSocket.CallOpts, pendingOwner)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlugSocket *PlugSocketCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlugSocket *PlugSocketSession) ProxiableUUID() ([32]byte, error) {
	return _PlugSocket.Contract.ProxiableUUID(&_PlugSocket.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlugSocket *PlugSocketCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PlugSocket.Contract.ProxiableUUID(&_PlugSocket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol)
func (_PlugSocket *PlugSocketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol)
func (_PlugSocket *PlugSocketSession) Symbol() (string, error) {
	return _PlugSocket.Contract.Symbol(&_PlugSocket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol)
func (_PlugSocket *PlugSocketCallerSession) Symbol() (string, error) {
	return _PlugSocket.Contract.Symbol(&_PlugSocket.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string version)
func (_PlugSocket *PlugSocketCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlugSocket.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string version)
func (_PlugSocket *PlugSocketSession) Version() (string, error) {
	return _PlugSocket.Contract.Version(&_PlugSocket.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string version)
func (_PlugSocket *PlugSocketCallerSession) Version() (string, error) {
	return _PlugSocket.Contract.Version(&_PlugSocket.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PlugSocket.Contract.CancelOwnershipHandover(&_PlugSocket.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PlugSocket.Contract.CancelOwnershipHandover(&_PlugSocket.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PlugSocket *PlugSocketTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PlugSocket *PlugSocketSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.CompleteOwnershipHandover(&_PlugSocket.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PlugSocket *PlugSocketTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.CompleteOwnershipHandover(&_PlugSocket.TransactOpts, pendingOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address oneClicker) returns()
func (_PlugSocket *PlugSocketTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, oneClicker common.Address) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "initialize", owner, oneClicker)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address oneClicker) returns()
func (_PlugSocket *PlugSocketSession) Initialize(owner common.Address, oneClicker common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.Initialize(&_PlugSocket.TransactOpts, owner, oneClicker)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address oneClicker) returns()
func (_PlugSocket *PlugSocketTransactorSession) Initialize(owner common.Address, oneClicker common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.Initialize(&_PlugSocket.TransactOpts, owner, oneClicker)
}

// OneClick is a paid mutator transaction binding the contract method 0x6f69f3d6.
//
// Solidity: function oneClick(address[] oneClickers, bool[] allowance) returns()
func (_PlugSocket *PlugSocketTransactor) OneClick(opts *bind.TransactOpts, oneClickers []common.Address, allowance []bool) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "oneClick", oneClickers, allowance)
}

// OneClick is a paid mutator transaction binding the contract method 0x6f69f3d6.
//
// Solidity: function oneClick(address[] oneClickers, bool[] allowance) returns()
func (_PlugSocket *PlugSocketSession) OneClick(oneClickers []common.Address, allowance []bool) (*types.Transaction, error) {
	return _PlugSocket.Contract.OneClick(&_PlugSocket.TransactOpts, oneClickers, allowance)
}

// OneClick is a paid mutator transaction binding the contract method 0x6f69f3d6.
//
// Solidity: function oneClick(address[] oneClickers, bool[] allowance) returns()
func (_PlugSocket *PlugSocketTransactorSession) OneClick(oneClickers []common.Address, allowance []bool) (*types.Transaction, error) {
	return _PlugSocket.Contract.OneClick(&_PlugSocket.TransactOpts, oneClickers, allowance)
}

// Plug is a paid mutator transaction binding the contract method 0x53f022b4.
//
// Solidity: function plug((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) plugs) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketTransactor) Plug(opts *bind.TransactOpts, plugs PlugTypesLibPlugs) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "plug", plugs)
}

// Plug is a paid mutator transaction binding the contract method 0x53f022b4.
//
// Solidity: function plug((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) plugs) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketSession) Plug(plugs PlugTypesLibPlugs) (*types.Transaction, error) {
	return _PlugSocket.Contract.Plug(&_PlugSocket.TransactOpts, plugs)
}

// Plug is a paid mutator transaction binding the contract method 0x53f022b4.
//
// Solidity: function plug((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes) plugs) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketTransactorSession) Plug(plugs PlugTypesLibPlugs) (*types.Transaction, error) {
	return _PlugSocket.Contract.Plug(&_PlugSocket.TransactOpts, plugs)
}

// Plug0 is a paid mutator transaction binding the contract method 0xaf433c7c.
//
// Solidity: function plug(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) livePlugs, address solver) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketTransactor) Plug0(opts *bind.TransactOpts, livePlugs PlugTypesLibLivePlugs, solver common.Address) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "plug0", livePlugs, solver)
}

// Plug0 is a paid mutator transaction binding the contract method 0xaf433c7c.
//
// Solidity: function plug(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) livePlugs, address solver) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketSession) Plug0(livePlugs PlugTypesLibLivePlugs, solver common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.Plug0(&_PlugSocket.TransactOpts, livePlugs, solver)
}

// Plug0 is a paid mutator transaction binding the contract method 0xaf433c7c.
//
// Solidity: function plug(((address,(uint8,address,bytes,uint256,(uint256,(uint8,uint256,uint256))[])[],bytes,bytes),bytes) livePlugs, address solver) payable returns((uint8,string) results)
func (_PlugSocket *PlugSocketTransactorSession) Plug0(livePlugs PlugTypesLibLivePlugs, solver common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.Plug0(&_PlugSocket.TransactOpts, livePlugs, solver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PlugSocket *PlugSocketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PlugSocket *PlugSocketSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlugSocket.Contract.RenounceOwnership(&_PlugSocket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PlugSocket *PlugSocketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlugSocket.Contract.RenounceOwnership(&_PlugSocket.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PlugSocket.Contract.RequestOwnershipHandover(&_PlugSocket.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PlugSocket *PlugSocketTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PlugSocket.Contract.RequestOwnershipHandover(&_PlugSocket.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PlugSocket *PlugSocketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PlugSocket *PlugSocketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.TransferOwnership(&_PlugSocket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PlugSocket *PlugSocketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlugSocket.Contract.TransferOwnership(&_PlugSocket.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlugSocket *PlugSocketTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlugSocket.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlugSocket *PlugSocketSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlugSocket.Contract.UpgradeToAndCall(&_PlugSocket.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlugSocket *PlugSocketTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlugSocket.Contract.UpgradeToAndCall(&_PlugSocket.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PlugSocket *PlugSocketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PlugSocket.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PlugSocket *PlugSocketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PlugSocket.Contract.Fallback(&_PlugSocket.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PlugSocket *PlugSocketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PlugSocket.Contract.Fallback(&_PlugSocket.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugSocket *PlugSocketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugSocket.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugSocket *PlugSocketSession) Receive() (*types.Transaction, error) {
	return _PlugSocket.Contract.Receive(&_PlugSocket.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugSocket *PlugSocketTransactorSession) Receive() (*types.Transaction, error) {
	return _PlugSocket.Contract.Receive(&_PlugSocket.TransactOpts)
}

// PlugSocketOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the PlugSocket contract.
type PlugSocketOwnershipHandoverCanceledIterator struct {
	Event *PlugSocketOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *PlugSocketOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugSocketOwnershipHandoverCanceled)
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
		it.Event = new(PlugSocketOwnershipHandoverCanceled)
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
func (it *PlugSocketOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugSocketOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugSocketOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the PlugSocket contract.
type PlugSocketOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*PlugSocketOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugSocketOwnershipHandoverCanceledIterator{contract: _PlugSocket.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *PlugSocketOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugSocketOwnershipHandoverCanceled)
				if err := _PlugSocket.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*PlugSocketOwnershipHandoverCanceled, error) {
	event := new(PlugSocketOwnershipHandoverCanceled)
	if err := _PlugSocket.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugSocketOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the PlugSocket contract.
type PlugSocketOwnershipHandoverRequestedIterator struct {
	Event *PlugSocketOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *PlugSocketOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugSocketOwnershipHandoverRequested)
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
		it.Event = new(PlugSocketOwnershipHandoverRequested)
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
func (it *PlugSocketOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugSocketOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugSocketOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the PlugSocket contract.
type PlugSocketOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*PlugSocketOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugSocketOwnershipHandoverRequestedIterator{contract: _PlugSocket.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *PlugSocketOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugSocketOwnershipHandoverRequested)
				if err := _PlugSocket.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PlugSocket *PlugSocketFilterer) ParseOwnershipHandoverRequested(log types.Log) (*PlugSocketOwnershipHandoverRequested, error) {
	event := new(PlugSocketOwnershipHandoverRequested)
	if err := _PlugSocket.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugSocketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PlugSocket contract.
type PlugSocketOwnershipTransferredIterator struct {
	Event *PlugSocketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PlugSocketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugSocketOwnershipTransferred)
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
		it.Event = new(PlugSocketOwnershipTransferred)
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
func (it *PlugSocketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugSocketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugSocketOwnershipTransferred represents a OwnershipTransferred event raised by the PlugSocket contract.
type PlugSocketOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PlugSocket *PlugSocketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PlugSocketOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugSocketOwnershipTransferredIterator{contract: _PlugSocket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PlugSocket *PlugSocketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlugSocketOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PlugSocket.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugSocketOwnershipTransferred)
				if err := _PlugSocket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PlugSocket *PlugSocketFilterer) ParseOwnershipTransferred(log types.Log) (*PlugSocketOwnershipTransferred, error) {
	event := new(PlugSocketOwnershipTransferred)
	if err := _PlugSocket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugSocketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PlugSocket contract.
type PlugSocketUpgradedIterator struct {
	Event *PlugSocketUpgraded // Event containing the contract specifics and raw log

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
func (it *PlugSocketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugSocketUpgraded)
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
		it.Event = new(PlugSocketUpgraded)
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
func (it *PlugSocketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugSocketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugSocketUpgraded represents a Upgraded event raised by the PlugSocket contract.
type PlugSocketUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PlugSocket *PlugSocketFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PlugSocketUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PlugSocket.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PlugSocketUpgradedIterator{contract: _PlugSocket.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PlugSocket *PlugSocketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PlugSocketUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PlugSocket.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugSocketUpgraded)
				if err := _PlugSocket.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PlugSocket *PlugSocketFilterer) ParseUpgraded(log types.Log) (*PlugSocketUpgraded, error) {
	event := new(PlugSocketUpgraded)
	if err := _PlugSocket.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
