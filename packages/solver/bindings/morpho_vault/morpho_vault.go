// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho_vault

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

// MarketAllocation is an auto generated low-level Go binding around an user-defined struct.
type MarketAllocation struct {
	MarketParams MarketParams
	Assets       *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// MorphoVaultMetaData contains all meta data concerning the MorphoVault contract.
var MorphoVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMaxTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllCapsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DuplicateMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InconsistentAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReallocation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalTimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MarketNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxQueueLengthExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllocatorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorNorGuardianRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGuardianRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"PendingCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRemoval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"UnauthorizedMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SetCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"SetCurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"SetGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllocator\",\"type\":\"bool\"}],\"name\":\"SetIsAllocator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"SetSkimRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetSupplyQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SetTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newWithdrawQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetWithdrawQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SubmitCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"SubmitGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SubmitMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SubmitTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedTotalAssets\",\"type\":\"uint256\"}],\"name\":\"UpdateLastTotalAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS_OFFSET\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"contractIMorpho\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"acceptCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint184\",\"name\":\"cap\",\"type\":\"uint184\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"removableAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllocator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTimelock\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAllocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"name\":\"reallocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"setCurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAllocator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsAllocator\",\"type\":\"bool\"}],\"name\":\"setIsAllocator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"setSkimRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"setSupplyQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skimRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newSupplyCap\",\"type\":\"uint256\"}],\"name\":\"submitCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"submitGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"submitMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"submitTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateWithdrawQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MorphoVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoVaultMetaData.ABI instead.
var MorphoVaultABI = MorphoVaultMetaData.ABI

// MorphoVault is an auto generated Go binding around an Ethereum contract.
type MorphoVault struct {
	MorphoVaultCaller     // Read-only binding to the contract
	MorphoVaultTransactor // Write-only binding to the contract
	MorphoVaultFilterer   // Log filterer for contract events
}

// MorphoVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoVaultSession struct {
	Contract     *MorphoVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoVaultCallerSession struct {
	Contract *MorphoVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MorphoVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoVaultTransactorSession struct {
	Contract     *MorphoVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MorphoVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoVaultRaw struct {
	Contract *MorphoVault // Generic contract binding to access the raw methods on
}

// MorphoVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoVaultCallerRaw struct {
	Contract *MorphoVaultCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoVaultTransactorRaw struct {
	Contract *MorphoVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphoVault creates a new instance of MorphoVault, bound to a specific deployed contract.
func NewMorphoVault(address common.Address, backend bind.ContractBackend) (*MorphoVault, error) {
	contract, err := bindMorphoVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphoVault{MorphoVaultCaller: MorphoVaultCaller{contract: contract}, MorphoVaultTransactor: MorphoVaultTransactor{contract: contract}, MorphoVaultFilterer: MorphoVaultFilterer{contract: contract}}, nil
}

// NewMorphoVaultCaller creates a new read-only instance of MorphoVault, bound to a specific deployed contract.
func NewMorphoVaultCaller(address common.Address, caller bind.ContractCaller) (*MorphoVaultCaller, error) {
	contract, err := bindMorphoVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultCaller{contract: contract}, nil
}

// NewMorphoVaultTransactor creates a new write-only instance of MorphoVault, bound to a specific deployed contract.
func NewMorphoVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoVaultTransactor, error) {
	contract, err := bindMorphoVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultTransactor{contract: contract}, nil
}

// NewMorphoVaultFilterer creates a new log filterer instance of MorphoVault, bound to a specific deployed contract.
func NewMorphoVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoVaultFilterer, error) {
	contract, err := bindMorphoVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultFilterer{contract: contract}, nil
}

// bindMorphoVault binds a generic wrapper to an already deployed contract.
func bindMorphoVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoVault *MorphoVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoVault.Contract.MorphoVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoVault *MorphoVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.Contract.MorphoVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoVault *MorphoVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoVault.Contract.MorphoVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoVault *MorphoVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoVault *MorphoVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoVault *MorphoVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoVault.Contract.contract.Transact(opts, method, params...)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MorphoVault *MorphoVaultCaller) DECIMALSOFFSET(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "DECIMALS_OFFSET")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MorphoVault *MorphoVaultSession) DECIMALSOFFSET() (uint8, error) {
	return _MorphoVault.Contract.DECIMALSOFFSET(&_MorphoVault.CallOpts)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MorphoVault *MorphoVaultCallerSession) DECIMALSOFFSET() (uint8, error) {
	return _MorphoVault.Contract.DECIMALSOFFSET(&_MorphoVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoVault *MorphoVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoVault *MorphoVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphoVault.Contract.DOMAINSEPARATOR(&_MorphoVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphoVault *MorphoVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphoVault.Contract.DOMAINSEPARATOR(&_MorphoVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoVault *MorphoVaultCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoVault *MorphoVaultSession) MORPHO() (common.Address, error) {
	return _MorphoVault.Contract.MORPHO(&_MorphoVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) MORPHO() (common.Address, error) {
	return _MorphoVault.Contract.MORPHO(&_MorphoVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.Allowance(&_MorphoVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.Allowance(&_MorphoVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MorphoVault *MorphoVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MorphoVault *MorphoVaultSession) Asset() (common.Address, error) {
	return _MorphoVault.Contract.Asset(&_MorphoVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) Asset() (common.Address, error) {
	return _MorphoVault.Contract.Asset(&_MorphoVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.BalanceOf(&_MorphoVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.BalanceOf(&_MorphoVault.CallOpts, account)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MorphoVault *MorphoVaultCaller) Config(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "config", arg0)

	outstruct := new(struct {
		Cap         *big.Int
		Enabled     bool
		RemovableAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RemovableAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MorphoVault *MorphoVaultSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _MorphoVault.Contract.Config(&_MorphoVault.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MorphoVault *MorphoVaultCallerSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _MorphoVault.Contract.Config(&_MorphoVault.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.ConvertToAssets(&_MorphoVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.ConvertToAssets(&_MorphoVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.ConvertToShares(&_MorphoVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.ConvertToShares(&_MorphoVault.CallOpts, assets)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MorphoVault *MorphoVaultCaller) Curator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "curator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MorphoVault *MorphoVaultSession) Curator() (common.Address, error) {
	return _MorphoVault.Contract.Curator(&_MorphoVault.CallOpts)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) Curator() (common.Address, error) {
	return _MorphoVault.Contract.Curator(&_MorphoVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphoVault *MorphoVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphoVault *MorphoVaultSession) Decimals() (uint8, error) {
	return _MorphoVault.Contract.Decimals(&_MorphoVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphoVault *MorphoVaultCallerSession) Decimals() (uint8, error) {
	return _MorphoVault.Contract.Decimals(&_MorphoVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MorphoVault *MorphoVaultCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MorphoVault *MorphoVaultSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MorphoVault.Contract.Eip712Domain(&_MorphoVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MorphoVault *MorphoVaultCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MorphoVault.Contract.Eip712Domain(&_MorphoVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MorphoVault *MorphoVaultCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MorphoVault *MorphoVaultSession) Fee() (*big.Int, error) {
	return _MorphoVault.Contract.Fee(&_MorphoVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MorphoVault *MorphoVaultCallerSession) Fee() (*big.Int, error) {
	return _MorphoVault.Contract.Fee(&_MorphoVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoVault *MorphoVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoVault *MorphoVaultSession) FeeRecipient() (common.Address, error) {
	return _MorphoVault.Contract.FeeRecipient(&_MorphoVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _MorphoVault.Contract.FeeRecipient(&_MorphoVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MorphoVault *MorphoVaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MorphoVault *MorphoVaultSession) Guardian() (common.Address, error) {
	return _MorphoVault.Contract.Guardian(&_MorphoVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) Guardian() (common.Address, error) {
	return _MorphoVault.Contract.Guardian(&_MorphoVault.CallOpts)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MorphoVault *MorphoVaultCaller) IsAllocator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "isAllocator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MorphoVault *MorphoVaultSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _MorphoVault.Contract.IsAllocator(&_MorphoVault.CallOpts, arg0)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MorphoVault *MorphoVaultCallerSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _MorphoVault.Contract.IsAllocator(&_MorphoVault.CallOpts, arg0)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) LastTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "lastTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MorphoVault *MorphoVaultSession) LastTotalAssets() (*big.Int, error) {
	return _MorphoVault.Contract.LastTotalAssets(&_MorphoVault.CallOpts)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) LastTotalAssets() (*big.Int, error) {
	return _MorphoVault.Contract.LastTotalAssets(&_MorphoVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxDeposit(&_MorphoVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxDeposit(&_MorphoVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxMint(&_MorphoVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxMint(&_MorphoVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxRedeem(&_MorphoVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxRedeem(&_MorphoVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MorphoVault *MorphoVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MorphoVault *MorphoVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxWithdraw(&_MorphoVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MorphoVault *MorphoVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.MaxWithdraw(&_MorphoVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphoVault *MorphoVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphoVault *MorphoVaultSession) Name() (string, error) {
	return _MorphoVault.Contract.Name(&_MorphoVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphoVault *MorphoVaultCallerSession) Name() (string, error) {
	return _MorphoVault.Contract.Name(&_MorphoVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.Nonces(&_MorphoVault.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MorphoVault.Contract.Nonces(&_MorphoVault.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoVault *MorphoVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoVault *MorphoVaultSession) Owner() (common.Address, error) {
	return _MorphoVault.Contract.Owner(&_MorphoVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) Owner() (common.Address, error) {
	return _MorphoVault.Contract.Owner(&_MorphoVault.CallOpts)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultCaller) PendingCap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "pendingCap", arg0)

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingCap(&_MorphoVault.CallOpts, arg0)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultCallerSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingCap(&_MorphoVault.CallOpts, arg0)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MorphoVault *MorphoVaultCaller) PendingGuardian(opts *bind.CallOpts) (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "pendingGuardian")

	outstruct := new(struct {
		Value   common.Address
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MorphoVault *MorphoVaultSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingGuardian(&_MorphoVault.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MorphoVault *MorphoVaultCallerSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingGuardian(&_MorphoVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MorphoVault *MorphoVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MorphoVault *MorphoVaultSession) PendingOwner() (common.Address, error) {
	return _MorphoVault.Contract.PendingOwner(&_MorphoVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) PendingOwner() (common.Address, error) {
	return _MorphoVault.Contract.PendingOwner(&_MorphoVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultCaller) PendingTimelock(opts *bind.CallOpts) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "pendingTimelock")

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingTimelock(&_MorphoVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MorphoVault *MorphoVaultCallerSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MorphoVault.Contract.PendingTimelock(&_MorphoVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewDeposit(&_MorphoVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewDeposit(&_MorphoVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewMint(&_MorphoVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewMint(&_MorphoVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewRedeem(&_MorphoVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewRedeem(&_MorphoVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewWithdraw(&_MorphoVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _MorphoVault.Contract.PreviewWithdraw(&_MorphoVault.CallOpts, assets)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MorphoVault *MorphoVaultCaller) SkimRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "skimRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MorphoVault *MorphoVaultSession) SkimRecipient() (common.Address, error) {
	return _MorphoVault.Contract.SkimRecipient(&_MorphoVault.CallOpts)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MorphoVault *MorphoVaultCallerSession) SkimRecipient() (common.Address, error) {
	return _MorphoVault.Contract.SkimRecipient(&_MorphoVault.CallOpts)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultCaller) SupplyQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "supplyQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _MorphoVault.Contract.SupplyQueue(&_MorphoVault.CallOpts, arg0)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultCallerSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _MorphoVault.Contract.SupplyQueue(&_MorphoVault.CallOpts, arg0)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) SupplyQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "supplyQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultSession) SupplyQueueLength() (*big.Int, error) {
	return _MorphoVault.Contract.SupplyQueueLength(&_MorphoVault.CallOpts)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) SupplyQueueLength() (*big.Int, error) {
	return _MorphoVault.Contract.SupplyQueueLength(&_MorphoVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphoVault *MorphoVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphoVault *MorphoVaultSession) Symbol() (string, error) {
	return _MorphoVault.Contract.Symbol(&_MorphoVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphoVault *MorphoVaultCallerSession) Symbol() (string, error) {
	return _MorphoVault.Contract.Symbol(&_MorphoVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoVault *MorphoVaultSession) Timelock() (*big.Int, error) {
	return _MorphoVault.Contract.Timelock(&_MorphoVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) Timelock() (*big.Int, error) {
	return _MorphoVault.Contract.Timelock(&_MorphoVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MorphoVault *MorphoVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MorphoVault *MorphoVaultSession) TotalAssets() (*big.Int, error) {
	return _MorphoVault.Contract.TotalAssets(&_MorphoVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MorphoVault *MorphoVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _MorphoVault.Contract.TotalAssets(&_MorphoVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphoVault *MorphoVaultSession) TotalSupply() (*big.Int, error) {
	return _MorphoVault.Contract.TotalSupply(&_MorphoVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _MorphoVault.Contract.TotalSupply(&_MorphoVault.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _MorphoVault.Contract.WithdrawQueue(&_MorphoVault.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MorphoVault *MorphoVaultCallerSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _MorphoVault.Contract.WithdrawQueue(&_MorphoVault.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphoVault.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultSession) WithdrawQueueLength() (*big.Int, error) {
	return _MorphoVault.Contract.WithdrawQueueLength(&_MorphoVault.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MorphoVault *MorphoVaultCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _MorphoVault.Contract.WithdrawQueueLength(&_MorphoVault.CallOpts)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultTransactor) AcceptCap(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "acceptCap", marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptCap(&_MorphoVault.TransactOpts, marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultTransactorSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptCap(&_MorphoVault.TransactOpts, marketParams)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MorphoVault *MorphoVaultTransactor) AcceptGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "acceptGuardian")
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MorphoVault *MorphoVaultSession) AcceptGuardian() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptGuardian(&_MorphoVault.TransactOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MorphoVault *MorphoVaultTransactorSession) AcceptGuardian() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptGuardian(&_MorphoVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MorphoVault *MorphoVaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MorphoVault *MorphoVaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptOwnership(&_MorphoVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MorphoVault *MorphoVaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptOwnership(&_MorphoVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MorphoVault *MorphoVaultTransactor) AcceptTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "acceptTimelock")
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MorphoVault *MorphoVaultSession) AcceptTimelock() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptTimelock(&_MorphoVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MorphoVault *MorphoVaultTransactorSession) AcceptTimelock() (*types.Transaction, error) {
	return _MorphoVault.Contract.AcceptTimelock(&_MorphoVault.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.Approve(&_MorphoVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.Approve(&_MorphoVault.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MorphoVault *MorphoVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MorphoVault *MorphoVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Deposit(&_MorphoVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MorphoVault *MorphoVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Deposit(&_MorphoVault.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MorphoVault *MorphoVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MorphoVault *MorphoVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Mint(&_MorphoVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MorphoVault *MorphoVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Mint(&_MorphoVault.TransactOpts, shares, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MorphoVault *MorphoVaultTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MorphoVault *MorphoVaultSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.Multicall(&_MorphoVault.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MorphoVault *MorphoVaultTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.Multicall(&_MorphoVault.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphoVault *MorphoVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphoVault *MorphoVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.Permit(&_MorphoVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphoVault *MorphoVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.Permit(&_MorphoVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MorphoVault *MorphoVaultTransactor) Reallocate(opts *bind.TransactOpts, allocations []MarketAllocation) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "reallocate", allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MorphoVault *MorphoVaultSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _MorphoVault.Contract.Reallocate(&_MorphoVault.TransactOpts, allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MorphoVault *MorphoVaultTransactorSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _MorphoVault.Contract.Reallocate(&_MorphoVault.TransactOpts, allocations)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MorphoVault *MorphoVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MorphoVault *MorphoVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Redeem(&_MorphoVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MorphoVault *MorphoVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Redeem(&_MorphoVault.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphoVault *MorphoVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphoVault *MorphoVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphoVault.Contract.RenounceOwnership(&_MorphoVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphoVault *MorphoVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphoVault.Contract.RenounceOwnership(&_MorphoVault.TransactOpts)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MorphoVault *MorphoVaultTransactor) RevokePendingCap(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "revokePendingCap", id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MorphoVault *MorphoVaultSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingCap(&_MorphoVault.TransactOpts, id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MorphoVault *MorphoVaultTransactorSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingCap(&_MorphoVault.TransactOpts, id)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MorphoVault *MorphoVaultTransactor) RevokePendingGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "revokePendingGuardian")
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MorphoVault *MorphoVaultSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingGuardian(&_MorphoVault.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MorphoVault *MorphoVaultTransactorSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingGuardian(&_MorphoVault.TransactOpts)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MorphoVault *MorphoVaultTransactor) RevokePendingMarketRemoval(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "revokePendingMarketRemoval", id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MorphoVault *MorphoVaultSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingMarketRemoval(&_MorphoVault.TransactOpts, id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MorphoVault *MorphoVaultTransactorSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingMarketRemoval(&_MorphoVault.TransactOpts, id)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MorphoVault *MorphoVaultTransactor) RevokePendingTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "revokePendingTimelock")
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MorphoVault *MorphoVaultSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingTimelock(&_MorphoVault.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MorphoVault *MorphoVaultTransactorSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _MorphoVault.Contract.RevokePendingTimelock(&_MorphoVault.TransactOpts)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MorphoVault *MorphoVaultTransactor) SetCurator(opts *bind.TransactOpts, newCurator common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setCurator", newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MorphoVault *MorphoVaultSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetCurator(&_MorphoVault.TransactOpts, newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetCurator(&_MorphoVault.TransactOpts, newCurator)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MorphoVault *MorphoVaultTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MorphoVault *MorphoVaultSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetFee(&_MorphoVault.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetFee(&_MorphoVault.TransactOpts, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoVault *MorphoVaultTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoVault *MorphoVaultSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetFeeRecipient(&_MorphoVault.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetFeeRecipient(&_MorphoVault.TransactOpts, newFeeRecipient)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MorphoVault *MorphoVaultTransactor) SetIsAllocator(opts *bind.TransactOpts, newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setIsAllocator", newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MorphoVault *MorphoVaultSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetIsAllocator(&_MorphoVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetIsAllocator(&_MorphoVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MorphoVault *MorphoVaultTransactor) SetSkimRecipient(opts *bind.TransactOpts, newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setSkimRecipient", newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MorphoVault *MorphoVaultSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetSkimRecipient(&_MorphoVault.TransactOpts, newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetSkimRecipient(&_MorphoVault.TransactOpts, newSkimRecipient)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MorphoVault *MorphoVaultTransactor) SetSupplyQueue(opts *bind.TransactOpts, newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "setSupplyQueue", newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MorphoVault *MorphoVaultSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetSupplyQueue(&_MorphoVault.TransactOpts, newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MorphoVault.Contract.SetSupplyQueue(&_MorphoVault.TransactOpts, newSupplyQueue)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MorphoVault *MorphoVaultTransactor) Skim(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "skim", token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MorphoVault *MorphoVaultSession) Skim(token common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Skim(&_MorphoVault.TransactOpts, token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MorphoVault *MorphoVaultTransactorSession) Skim(token common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Skim(&_MorphoVault.TransactOpts, token)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MorphoVault *MorphoVaultTransactor) SubmitCap(opts *bind.TransactOpts, marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "submitCap", marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MorphoVault *MorphoVaultSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitCap(&_MorphoVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitCap(&_MorphoVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MorphoVault *MorphoVaultTransactor) SubmitGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "submitGuardian", newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MorphoVault *MorphoVaultSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitGuardian(&_MorphoVault.TransactOpts, newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitGuardian(&_MorphoVault.TransactOpts, newGuardian)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultTransactor) SubmitMarketRemoval(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "submitMarketRemoval", marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitMarketRemoval(&_MorphoVault.TransactOpts, marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitMarketRemoval(&_MorphoVault.TransactOpts, marketParams)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MorphoVault *MorphoVaultTransactor) SubmitTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "submitTimelock", newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MorphoVault *MorphoVaultSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitTimelock(&_MorphoVault.TransactOpts, newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MorphoVault *MorphoVaultTransactorSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.SubmitTimelock(&_MorphoVault.TransactOpts, newTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.Transfer(&_MorphoVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.Transfer(&_MorphoVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.TransferFrom(&_MorphoVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MorphoVault *MorphoVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.TransferFrom(&_MorphoVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphoVault *MorphoVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphoVault *MorphoVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.TransferOwnership(&_MorphoVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphoVault *MorphoVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.TransferOwnership(&_MorphoVault.TransactOpts, newOwner)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MorphoVault *MorphoVaultTransactor) UpdateWithdrawQueue(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "updateWithdrawQueue", indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MorphoVault *MorphoVaultSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.UpdateWithdrawQueue(&_MorphoVault.TransactOpts, indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MorphoVault *MorphoVaultTransactorSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _MorphoVault.Contract.UpdateWithdrawQueue(&_MorphoVault.TransactOpts, indexes)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MorphoVault *MorphoVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MorphoVault *MorphoVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Withdraw(&_MorphoVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MorphoVault *MorphoVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MorphoVault.Contract.Withdraw(&_MorphoVault.TransactOpts, assets, receiver, owner)
}

// MorphoVaultAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the MorphoVault contract.
type MorphoVaultAccrueInterestIterator struct {
	Event *MorphoVaultAccrueInterest // Event containing the contract specifics and raw log

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
func (it *MorphoVaultAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultAccrueInterest)
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
		it.Event = new(MorphoVaultAccrueInterest)
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
func (it *MorphoVaultAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultAccrueInterest represents a AccrueInterest event raised by the MorphoVault contract.
type MorphoVaultAccrueInterest struct {
	NewTotalAssets *big.Int
	FeeShares      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MorphoVault *MorphoVaultFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*MorphoVaultAccrueInterestIterator, error) {

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &MorphoVaultAccrueInterestIterator{contract: _MorphoVault.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MorphoVault *MorphoVaultFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *MorphoVaultAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultAccrueInterest)
				if err := _MorphoVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MorphoVault *MorphoVaultFilterer) ParseAccrueInterest(log types.Log) (*MorphoVaultAccrueInterest, error) {
	event := new(MorphoVaultAccrueInterest)
	if err := _MorphoVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MorphoVault contract.
type MorphoVaultApprovalIterator struct {
	Event *MorphoVaultApproval // Event containing the contract specifics and raw log

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
func (it *MorphoVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultApproval)
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
		it.Event = new(MorphoVaultApproval)
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
func (it *MorphoVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultApproval represents a Approval event raised by the MorphoVault contract.
type MorphoVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MorphoVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultApprovalIterator{contract: _MorphoVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MorphoVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultApproval)
				if err := _MorphoVault.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) ParseApproval(log types.Log) (*MorphoVaultApproval, error) {
	event := new(MorphoVaultApproval)
	if err := _MorphoVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MorphoVault contract.
type MorphoVaultDepositIterator struct {
	Event *MorphoVaultDeposit // Event containing the contract specifics and raw log

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
func (it *MorphoVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultDeposit)
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
		it.Event = new(MorphoVaultDeposit)
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
func (it *MorphoVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultDeposit represents a Deposit event raised by the MorphoVault contract.
type MorphoVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*MorphoVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultDepositIterator{contract: _MorphoVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MorphoVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultDeposit)
				if err := _MorphoVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) ParseDeposit(log types.Log) (*MorphoVaultDeposit, error) {
	event := new(MorphoVaultDeposit)
	if err := _MorphoVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MorphoVault contract.
type MorphoVaultEIP712DomainChangedIterator struct {
	Event *MorphoVaultEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MorphoVaultEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultEIP712DomainChanged)
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
		it.Event = new(MorphoVaultEIP712DomainChanged)
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
func (it *MorphoVaultEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultEIP712DomainChanged represents a EIP712DomainChanged event raised by the MorphoVault contract.
type MorphoVaultEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MorphoVault *MorphoVaultFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MorphoVaultEIP712DomainChangedIterator, error) {

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MorphoVaultEIP712DomainChangedIterator{contract: _MorphoVault.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MorphoVault *MorphoVaultFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MorphoVaultEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultEIP712DomainChanged)
				if err := _MorphoVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MorphoVault *MorphoVaultFilterer) ParseEIP712DomainChanged(log types.Log) (*MorphoVaultEIP712DomainChanged, error) {
	event := new(MorphoVaultEIP712DomainChanged)
	if err := _MorphoVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MorphoVault contract.
type MorphoVaultOwnershipTransferStartedIterator struct {
	Event *MorphoVaultOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MorphoVaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultOwnershipTransferStarted)
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
		it.Event = new(MorphoVaultOwnershipTransferStarted)
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
func (it *MorphoVaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MorphoVault contract.
type MorphoVaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MorphoVault *MorphoVaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphoVaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultOwnershipTransferStartedIterator{contract: _MorphoVault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MorphoVault *MorphoVaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MorphoVaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultOwnershipTransferStarted)
				if err := _MorphoVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MorphoVault *MorphoVaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*MorphoVaultOwnershipTransferStarted, error) {
	event := new(MorphoVaultOwnershipTransferStarted)
	if err := _MorphoVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MorphoVault contract.
type MorphoVaultOwnershipTransferredIterator struct {
	Event *MorphoVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MorphoVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultOwnershipTransferred)
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
		it.Event = new(MorphoVaultOwnershipTransferred)
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
func (it *MorphoVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultOwnershipTransferred represents a OwnershipTransferred event raised by the MorphoVault contract.
type MorphoVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphoVault *MorphoVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphoVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultOwnershipTransferredIterator{contract: _MorphoVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphoVault *MorphoVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MorphoVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultOwnershipTransferred)
				if err := _MorphoVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MorphoVault *MorphoVaultFilterer) ParseOwnershipTransferred(log types.Log) (*MorphoVaultOwnershipTransferred, error) {
	event := new(MorphoVaultOwnershipTransferred)
	if err := _MorphoVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultReallocateSupplyIterator is returned from FilterReallocateSupply and is used to iterate over the raw logs and unpacked data for ReallocateSupply events raised by the MorphoVault contract.
type MorphoVaultReallocateSupplyIterator struct {
	Event *MorphoVaultReallocateSupply // Event containing the contract specifics and raw log

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
func (it *MorphoVaultReallocateSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultReallocateSupply)
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
		it.Event = new(MorphoVaultReallocateSupply)
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
func (it *MorphoVaultReallocateSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultReallocateSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultReallocateSupply represents a ReallocateSupply event raised by the MorphoVault contract.
type MorphoVaultReallocateSupply struct {
	Caller         common.Address
	Id             [32]byte
	SuppliedAssets *big.Int
	SuppliedShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReallocateSupply is a free log retrieval operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MorphoVault *MorphoVaultFilterer) FilterReallocateSupply(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultReallocateSupplyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultReallocateSupplyIterator{contract: _MorphoVault.contract, event: "ReallocateSupply", logs: logs, sub: sub}, nil
}

// WatchReallocateSupply is a free log subscription operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MorphoVault *MorphoVaultFilterer) WatchReallocateSupply(opts *bind.WatchOpts, sink chan<- *MorphoVaultReallocateSupply, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultReallocateSupply)
				if err := _MorphoVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
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

// ParseReallocateSupply is a log parse operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MorphoVault *MorphoVaultFilterer) ParseReallocateSupply(log types.Log) (*MorphoVaultReallocateSupply, error) {
	event := new(MorphoVaultReallocateSupply)
	if err := _MorphoVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultReallocateWithdrawIterator is returned from FilterReallocateWithdraw and is used to iterate over the raw logs and unpacked data for ReallocateWithdraw events raised by the MorphoVault contract.
type MorphoVaultReallocateWithdrawIterator struct {
	Event *MorphoVaultReallocateWithdraw // Event containing the contract specifics and raw log

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
func (it *MorphoVaultReallocateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultReallocateWithdraw)
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
		it.Event = new(MorphoVaultReallocateWithdraw)
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
func (it *MorphoVaultReallocateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultReallocateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultReallocateWithdraw represents a ReallocateWithdraw event raised by the MorphoVault contract.
type MorphoVaultReallocateWithdraw struct {
	Caller          common.Address
	Id              [32]byte
	WithdrawnAssets *big.Int
	WithdrawnShares *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReallocateWithdraw is a free log retrieval operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MorphoVault *MorphoVaultFilterer) FilterReallocateWithdraw(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultReallocateWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultReallocateWithdrawIterator{contract: _MorphoVault.contract, event: "ReallocateWithdraw", logs: logs, sub: sub}, nil
}

// WatchReallocateWithdraw is a free log subscription operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MorphoVault *MorphoVaultFilterer) WatchReallocateWithdraw(opts *bind.WatchOpts, sink chan<- *MorphoVaultReallocateWithdraw, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultReallocateWithdraw)
				if err := _MorphoVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
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

// ParseReallocateWithdraw is a log parse operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MorphoVault *MorphoVaultFilterer) ParseReallocateWithdraw(log types.Log) (*MorphoVaultReallocateWithdraw, error) {
	event := new(MorphoVaultReallocateWithdraw)
	if err := _MorphoVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultRevokePendingCapIterator is returned from FilterRevokePendingCap and is used to iterate over the raw logs and unpacked data for RevokePendingCap events raised by the MorphoVault contract.
type MorphoVaultRevokePendingCapIterator struct {
	Event *MorphoVaultRevokePendingCap // Event containing the contract specifics and raw log

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
func (it *MorphoVaultRevokePendingCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultRevokePendingCap)
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
		it.Event = new(MorphoVaultRevokePendingCap)
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
func (it *MorphoVaultRevokePendingCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultRevokePendingCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultRevokePendingCap represents a RevokePendingCap event raised by the MorphoVault contract.
type MorphoVaultRevokePendingCap struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingCap is a free log retrieval operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) FilterRevokePendingCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultRevokePendingCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultRevokePendingCapIterator{contract: _MorphoVault.contract, event: "RevokePendingCap", logs: logs, sub: sub}, nil
}

// WatchRevokePendingCap is a free log subscription operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) WatchRevokePendingCap(opts *bind.WatchOpts, sink chan<- *MorphoVaultRevokePendingCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultRevokePendingCap)
				if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
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

// ParseRevokePendingCap is a log parse operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) ParseRevokePendingCap(log types.Log) (*MorphoVaultRevokePendingCap, error) {
	event := new(MorphoVaultRevokePendingCap)
	if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultRevokePendingGuardianIterator is returned from FilterRevokePendingGuardian and is used to iterate over the raw logs and unpacked data for RevokePendingGuardian events raised by the MorphoVault contract.
type MorphoVaultRevokePendingGuardianIterator struct {
	Event *MorphoVaultRevokePendingGuardian // Event containing the contract specifics and raw log

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
func (it *MorphoVaultRevokePendingGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultRevokePendingGuardian)
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
		it.Event = new(MorphoVaultRevokePendingGuardian)
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
func (it *MorphoVaultRevokePendingGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultRevokePendingGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultRevokePendingGuardian represents a RevokePendingGuardian event raised by the MorphoVault contract.
type MorphoVaultRevokePendingGuardian struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingGuardian is a free log retrieval operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) FilterRevokePendingGuardian(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultRevokePendingGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultRevokePendingGuardianIterator{contract: _MorphoVault.contract, event: "RevokePendingGuardian", logs: logs, sub: sub}, nil
}

// WatchRevokePendingGuardian is a free log subscription operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) WatchRevokePendingGuardian(opts *bind.WatchOpts, sink chan<- *MorphoVaultRevokePendingGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultRevokePendingGuardian)
				if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
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

// ParseRevokePendingGuardian is a log parse operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) ParseRevokePendingGuardian(log types.Log) (*MorphoVaultRevokePendingGuardian, error) {
	event := new(MorphoVaultRevokePendingGuardian)
	if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultRevokePendingMarketRemovalIterator is returned from FilterRevokePendingMarketRemoval and is used to iterate over the raw logs and unpacked data for RevokePendingMarketRemoval events raised by the MorphoVault contract.
type MorphoVaultRevokePendingMarketRemovalIterator struct {
	Event *MorphoVaultRevokePendingMarketRemoval // Event containing the contract specifics and raw log

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
func (it *MorphoVaultRevokePendingMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultRevokePendingMarketRemoval)
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
		it.Event = new(MorphoVaultRevokePendingMarketRemoval)
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
func (it *MorphoVaultRevokePendingMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultRevokePendingMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultRevokePendingMarketRemoval represents a RevokePendingMarketRemoval event raised by the MorphoVault contract.
type MorphoVaultRevokePendingMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMarketRemoval is a free log retrieval operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) FilterRevokePendingMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultRevokePendingMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultRevokePendingMarketRemovalIterator{contract: _MorphoVault.contract, event: "RevokePendingMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMarketRemoval is a free log subscription operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) WatchRevokePendingMarketRemoval(opts *bind.WatchOpts, sink chan<- *MorphoVaultRevokePendingMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultRevokePendingMarketRemoval)
				if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
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

// ParseRevokePendingMarketRemoval is a log parse operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) ParseRevokePendingMarketRemoval(log types.Log) (*MorphoVaultRevokePendingMarketRemoval, error) {
	event := new(MorphoVaultRevokePendingMarketRemoval)
	if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultRevokePendingTimelockIterator is returned from FilterRevokePendingTimelock and is used to iterate over the raw logs and unpacked data for RevokePendingTimelock events raised by the MorphoVault contract.
type MorphoVaultRevokePendingTimelockIterator struct {
	Event *MorphoVaultRevokePendingTimelock // Event containing the contract specifics and raw log

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
func (it *MorphoVaultRevokePendingTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultRevokePendingTimelock)
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
		it.Event = new(MorphoVaultRevokePendingTimelock)
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
func (it *MorphoVaultRevokePendingTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultRevokePendingTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultRevokePendingTimelock represents a RevokePendingTimelock event raised by the MorphoVault contract.
type MorphoVaultRevokePendingTimelock struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingTimelock is a free log retrieval operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) FilterRevokePendingTimelock(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultRevokePendingTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultRevokePendingTimelockIterator{contract: _MorphoVault.contract, event: "RevokePendingTimelock", logs: logs, sub: sub}, nil
}

// WatchRevokePendingTimelock is a free log subscription operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) WatchRevokePendingTimelock(opts *bind.WatchOpts, sink chan<- *MorphoVaultRevokePendingTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultRevokePendingTimelock)
				if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
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

// ParseRevokePendingTimelock is a log parse operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MorphoVault *MorphoVaultFilterer) ParseRevokePendingTimelock(log types.Log) (*MorphoVaultRevokePendingTimelock, error) {
	event := new(MorphoVaultRevokePendingTimelock)
	if err := _MorphoVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetCapIterator is returned from FilterSetCap and is used to iterate over the raw logs and unpacked data for SetCap events raised by the MorphoVault contract.
type MorphoVaultSetCapIterator struct {
	Event *MorphoVaultSetCap // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetCap)
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
		it.Event = new(MorphoVaultSetCap)
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
func (it *MorphoVaultSetCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetCap represents a SetCap event raised by the MorphoVault contract.
type MorphoVaultSetCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetCap is a free log retrieval operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) FilterSetCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultSetCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetCapIterator{contract: _MorphoVault.contract, event: "SetCap", logs: logs, sub: sub}, nil
}

// WatchSetCap is a free log subscription operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) WatchSetCap(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetCap)
				if err := _MorphoVault.contract.UnpackLog(event, "SetCap", log); err != nil {
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

// ParseSetCap is a log parse operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) ParseSetCap(log types.Log) (*MorphoVaultSetCap, error) {
	event := new(MorphoVaultSetCap)
	if err := _MorphoVault.contract.UnpackLog(event, "SetCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetCuratorIterator is returned from FilterSetCurator and is used to iterate over the raw logs and unpacked data for SetCurator events raised by the MorphoVault contract.
type MorphoVaultSetCuratorIterator struct {
	Event *MorphoVaultSetCurator // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetCuratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetCurator)
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
		it.Event = new(MorphoVaultSetCurator)
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
func (it *MorphoVaultSetCuratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetCuratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetCurator represents a SetCurator event raised by the MorphoVault contract.
type MorphoVaultSetCurator struct {
	NewCurator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCurator is a free log retrieval operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MorphoVault *MorphoVaultFilterer) FilterSetCurator(opts *bind.FilterOpts, newCurator []common.Address) (*MorphoVaultSetCuratorIterator, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetCuratorIterator{contract: _MorphoVault.contract, event: "SetCurator", logs: logs, sub: sub}, nil
}

// WatchSetCurator is a free log subscription operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MorphoVault *MorphoVaultFilterer) WatchSetCurator(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetCurator, newCurator []common.Address) (event.Subscription, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetCurator)
				if err := _MorphoVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
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

// ParseSetCurator is a log parse operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MorphoVault *MorphoVaultFilterer) ParseSetCurator(log types.Log) (*MorphoVaultSetCurator, error) {
	event := new(MorphoVaultSetCurator)
	if err := _MorphoVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the MorphoVault contract.
type MorphoVaultSetFeeIterator struct {
	Event *MorphoVaultSetFee // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetFee)
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
		it.Event = new(MorphoVaultSetFee)
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
func (it *MorphoVaultSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetFee represents a SetFee event raised by the MorphoVault contract.
type MorphoVaultSetFee struct {
	Caller common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MorphoVault *MorphoVaultFilterer) FilterSetFee(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultSetFeeIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetFeeIterator{contract: _MorphoVault.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MorphoVault *MorphoVaultFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetFee, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetFee)
				if err := _MorphoVault.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MorphoVault *MorphoVaultFilterer) ParseSetFee(log types.Log) (*MorphoVaultSetFee, error) {
	event := new(MorphoVaultSetFee)
	if err := _MorphoVault.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the MorphoVault contract.
type MorphoVaultSetFeeRecipientIterator struct {
	Event *MorphoVaultSetFeeRecipient // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetFeeRecipient)
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
		it.Event = new(MorphoVaultSetFeeRecipient)
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
func (it *MorphoVaultSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetFeeRecipient represents a SetFeeRecipient event raised by the MorphoVault contract.
type MorphoVaultSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MorphoVault *MorphoVaultFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts, newFeeRecipient []common.Address) (*MorphoVaultSetFeeRecipientIterator, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetFeeRecipientIterator{contract: _MorphoVault.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MorphoVault *MorphoVaultFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetFeeRecipient, newFeeRecipient []common.Address) (event.Subscription, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetFeeRecipient)
				if err := _MorphoVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
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
func (_MorphoVault *MorphoVaultFilterer) ParseSetFeeRecipient(log types.Log) (*MorphoVaultSetFeeRecipient, error) {
	event := new(MorphoVaultSetFeeRecipient)
	if err := _MorphoVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetGuardianIterator is returned from FilterSetGuardian and is used to iterate over the raw logs and unpacked data for SetGuardian events raised by the MorphoVault contract.
type MorphoVaultSetGuardianIterator struct {
	Event *MorphoVaultSetGuardian // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetGuardian)
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
		it.Event = new(MorphoVaultSetGuardian)
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
func (it *MorphoVaultSetGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetGuardian represents a SetGuardian event raised by the MorphoVault contract.
type MorphoVaultSetGuardian struct {
	Caller   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGuardian is a free log retrieval operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MorphoVault *MorphoVaultFilterer) FilterSetGuardian(opts *bind.FilterOpts, caller []common.Address, guardian []common.Address) (*MorphoVaultSetGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetGuardianIterator{contract: _MorphoVault.contract, event: "SetGuardian", logs: logs, sub: sub}, nil
}

// WatchSetGuardian is a free log subscription operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MorphoVault *MorphoVaultFilterer) WatchSetGuardian(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetGuardian, caller []common.Address, guardian []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetGuardian)
				if err := _MorphoVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
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

// ParseSetGuardian is a log parse operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MorphoVault *MorphoVaultFilterer) ParseSetGuardian(log types.Log) (*MorphoVaultSetGuardian, error) {
	event := new(MorphoVaultSetGuardian)
	if err := _MorphoVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetIsAllocatorIterator is returned from FilterSetIsAllocator and is used to iterate over the raw logs and unpacked data for SetIsAllocator events raised by the MorphoVault contract.
type MorphoVaultSetIsAllocatorIterator struct {
	Event *MorphoVaultSetIsAllocator // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetIsAllocatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetIsAllocator)
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
		it.Event = new(MorphoVaultSetIsAllocator)
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
func (it *MorphoVaultSetIsAllocatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetIsAllocatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetIsAllocator represents a SetIsAllocator event raised by the MorphoVault contract.
type MorphoVaultSetIsAllocator struct {
	Allocator   common.Address
	IsAllocator bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllocator is a free log retrieval operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MorphoVault *MorphoVaultFilterer) FilterSetIsAllocator(opts *bind.FilterOpts, allocator []common.Address) (*MorphoVaultSetIsAllocatorIterator, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetIsAllocatorIterator{contract: _MorphoVault.contract, event: "SetIsAllocator", logs: logs, sub: sub}, nil
}

// WatchSetIsAllocator is a free log subscription operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MorphoVault *MorphoVaultFilterer) WatchSetIsAllocator(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetIsAllocator, allocator []common.Address) (event.Subscription, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetIsAllocator)
				if err := _MorphoVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
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

// ParseSetIsAllocator is a log parse operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MorphoVault *MorphoVaultFilterer) ParseSetIsAllocator(log types.Log) (*MorphoVaultSetIsAllocator, error) {
	event := new(MorphoVaultSetIsAllocator)
	if err := _MorphoVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetSkimRecipientIterator is returned from FilterSetSkimRecipient and is used to iterate over the raw logs and unpacked data for SetSkimRecipient events raised by the MorphoVault contract.
type MorphoVaultSetSkimRecipientIterator struct {
	Event *MorphoVaultSetSkimRecipient // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetSkimRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetSkimRecipient)
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
		it.Event = new(MorphoVaultSetSkimRecipient)
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
func (it *MorphoVaultSetSkimRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetSkimRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetSkimRecipient represents a SetSkimRecipient event raised by the MorphoVault contract.
type MorphoVaultSetSkimRecipient struct {
	NewSkimRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetSkimRecipient is a free log retrieval operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MorphoVault *MorphoVaultFilterer) FilterSetSkimRecipient(opts *bind.FilterOpts, newSkimRecipient []common.Address) (*MorphoVaultSetSkimRecipientIterator, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetSkimRecipientIterator{contract: _MorphoVault.contract, event: "SetSkimRecipient", logs: logs, sub: sub}, nil
}

// WatchSetSkimRecipient is a free log subscription operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MorphoVault *MorphoVaultFilterer) WatchSetSkimRecipient(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetSkimRecipient, newSkimRecipient []common.Address) (event.Subscription, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetSkimRecipient)
				if err := _MorphoVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
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

// ParseSetSkimRecipient is a log parse operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MorphoVault *MorphoVaultFilterer) ParseSetSkimRecipient(log types.Log) (*MorphoVaultSetSkimRecipient, error) {
	event := new(MorphoVaultSetSkimRecipient)
	if err := _MorphoVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetSupplyQueueIterator is returned from FilterSetSupplyQueue and is used to iterate over the raw logs and unpacked data for SetSupplyQueue events raised by the MorphoVault contract.
type MorphoVaultSetSupplyQueueIterator struct {
	Event *MorphoVaultSetSupplyQueue // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetSupplyQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetSupplyQueue)
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
		it.Event = new(MorphoVaultSetSupplyQueue)
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
func (it *MorphoVaultSetSupplyQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetSupplyQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetSupplyQueue represents a SetSupplyQueue event raised by the MorphoVault contract.
type MorphoVaultSetSupplyQueue struct {
	Caller         common.Address
	NewSupplyQueue [][32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyQueue is a free log retrieval operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MorphoVault *MorphoVaultFilterer) FilterSetSupplyQueue(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultSetSupplyQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetSupplyQueueIterator{contract: _MorphoVault.contract, event: "SetSupplyQueue", logs: logs, sub: sub}, nil
}

// WatchSetSupplyQueue is a free log subscription operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MorphoVault *MorphoVaultFilterer) WatchSetSupplyQueue(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetSupplyQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetSupplyQueue)
				if err := _MorphoVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
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

// ParseSetSupplyQueue is a log parse operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MorphoVault *MorphoVaultFilterer) ParseSetSupplyQueue(log types.Log) (*MorphoVaultSetSupplyQueue, error) {
	event := new(MorphoVaultSetSupplyQueue)
	if err := _MorphoVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetTimelockIterator is returned from FilterSetTimelock and is used to iterate over the raw logs and unpacked data for SetTimelock events raised by the MorphoVault contract.
type MorphoVaultSetTimelockIterator struct {
	Event *MorphoVaultSetTimelock // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetTimelock)
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
		it.Event = new(MorphoVaultSetTimelock)
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
func (it *MorphoVaultSetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetTimelock represents a SetTimelock event raised by the MorphoVault contract.
type MorphoVaultSetTimelock struct {
	Caller      common.Address
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTimelock is a free log retrieval operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) FilterSetTimelock(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultSetTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetTimelockIterator{contract: _MorphoVault.contract, event: "SetTimelock", logs: logs, sub: sub}, nil
}

// WatchSetTimelock is a free log subscription operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) WatchSetTimelock(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetTimelock)
				if err := _MorphoVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
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

// ParseSetTimelock is a log parse operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) ParseSetTimelock(log types.Log) (*MorphoVaultSetTimelock, error) {
	event := new(MorphoVaultSetTimelock)
	if err := _MorphoVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSetWithdrawQueueIterator is returned from FilterSetWithdrawQueue and is used to iterate over the raw logs and unpacked data for SetWithdrawQueue events raised by the MorphoVault contract.
type MorphoVaultSetWithdrawQueueIterator struct {
	Event *MorphoVaultSetWithdrawQueue // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSetWithdrawQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSetWithdrawQueue)
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
		it.Event = new(MorphoVaultSetWithdrawQueue)
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
func (it *MorphoVaultSetWithdrawQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSetWithdrawQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSetWithdrawQueue represents a SetWithdrawQueue event raised by the MorphoVault contract.
type MorphoVaultSetWithdrawQueue struct {
	Caller           common.Address
	NewWithdrawQueue [][32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawQueue is a free log retrieval operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MorphoVault *MorphoVaultFilterer) FilterSetWithdrawQueue(opts *bind.FilterOpts, caller []common.Address) (*MorphoVaultSetWithdrawQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSetWithdrawQueueIterator{contract: _MorphoVault.contract, event: "SetWithdrawQueue", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawQueue is a free log subscription operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MorphoVault *MorphoVaultFilterer) WatchSetWithdrawQueue(opts *bind.WatchOpts, sink chan<- *MorphoVaultSetWithdrawQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSetWithdrawQueue)
				if err := _MorphoVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
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

// ParseSetWithdrawQueue is a log parse operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MorphoVault *MorphoVaultFilterer) ParseSetWithdrawQueue(log types.Log) (*MorphoVaultSetWithdrawQueue, error) {
	event := new(MorphoVaultSetWithdrawQueue)
	if err := _MorphoVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the MorphoVault contract.
type MorphoVaultSkimIterator struct {
	Event *MorphoVaultSkim // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSkim)
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
		it.Event = new(MorphoVaultSkim)
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
func (it *MorphoVaultSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSkim represents a Skim event raised by the MorphoVault contract.
type MorphoVaultSkim struct {
	Caller common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MorphoVault *MorphoVaultFilterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*MorphoVaultSkimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSkimIterator{contract: _MorphoVault.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MorphoVault *MorphoVaultFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *MorphoVaultSkim, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSkim)
				if err := _MorphoVault.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MorphoVault *MorphoVaultFilterer) ParseSkim(log types.Log) (*MorphoVaultSkim, error) {
	event := new(MorphoVaultSkim)
	if err := _MorphoVault.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSubmitCapIterator is returned from FilterSubmitCap and is used to iterate over the raw logs and unpacked data for SubmitCap events raised by the MorphoVault contract.
type MorphoVaultSubmitCapIterator struct {
	Event *MorphoVaultSubmitCap // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSubmitCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSubmitCap)
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
		it.Event = new(MorphoVaultSubmitCap)
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
func (it *MorphoVaultSubmitCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSubmitCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSubmitCap represents a SubmitCap event raised by the MorphoVault contract.
type MorphoVaultSubmitCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitCap is a free log retrieval operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) FilterSubmitCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultSubmitCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSubmitCapIterator{contract: _MorphoVault.contract, event: "SubmitCap", logs: logs, sub: sub}, nil
}

// WatchSubmitCap is a free log subscription operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) WatchSubmitCap(opts *bind.WatchOpts, sink chan<- *MorphoVaultSubmitCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSubmitCap)
				if err := _MorphoVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
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

// ParseSubmitCap is a log parse operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MorphoVault *MorphoVaultFilterer) ParseSubmitCap(log types.Log) (*MorphoVaultSubmitCap, error) {
	event := new(MorphoVaultSubmitCap)
	if err := _MorphoVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSubmitGuardianIterator is returned from FilterSubmitGuardian and is used to iterate over the raw logs and unpacked data for SubmitGuardian events raised by the MorphoVault contract.
type MorphoVaultSubmitGuardianIterator struct {
	Event *MorphoVaultSubmitGuardian // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSubmitGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSubmitGuardian)
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
		it.Event = new(MorphoVaultSubmitGuardian)
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
func (it *MorphoVaultSubmitGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSubmitGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSubmitGuardian represents a SubmitGuardian event raised by the MorphoVault contract.
type MorphoVaultSubmitGuardian struct {
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitGuardian is a free log retrieval operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MorphoVault *MorphoVaultFilterer) FilterSubmitGuardian(opts *bind.FilterOpts, newGuardian []common.Address) (*MorphoVaultSubmitGuardianIterator, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSubmitGuardianIterator{contract: _MorphoVault.contract, event: "SubmitGuardian", logs: logs, sub: sub}, nil
}

// WatchSubmitGuardian is a free log subscription operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MorphoVault *MorphoVaultFilterer) WatchSubmitGuardian(opts *bind.WatchOpts, sink chan<- *MorphoVaultSubmitGuardian, newGuardian []common.Address) (event.Subscription, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSubmitGuardian)
				if err := _MorphoVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
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

// ParseSubmitGuardian is a log parse operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MorphoVault *MorphoVaultFilterer) ParseSubmitGuardian(log types.Log) (*MorphoVaultSubmitGuardian, error) {
	event := new(MorphoVaultSubmitGuardian)
	if err := _MorphoVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSubmitMarketRemovalIterator is returned from FilterSubmitMarketRemoval and is used to iterate over the raw logs and unpacked data for SubmitMarketRemoval events raised by the MorphoVault contract.
type MorphoVaultSubmitMarketRemovalIterator struct {
	Event *MorphoVaultSubmitMarketRemoval // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSubmitMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSubmitMarketRemoval)
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
		it.Event = new(MorphoVaultSubmitMarketRemoval)
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
func (it *MorphoVaultSubmitMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSubmitMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSubmitMarketRemoval represents a SubmitMarketRemoval event raised by the MorphoVault contract.
type MorphoVaultSubmitMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitMarketRemoval is a free log retrieval operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) FilterSubmitMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MorphoVaultSubmitMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSubmitMarketRemovalIterator{contract: _MorphoVault.contract, event: "SubmitMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmitMarketRemoval is a free log subscription operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) WatchSubmitMarketRemoval(opts *bind.WatchOpts, sink chan<- *MorphoVaultSubmitMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSubmitMarketRemoval)
				if err := _MorphoVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
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

// ParseSubmitMarketRemoval is a log parse operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MorphoVault *MorphoVaultFilterer) ParseSubmitMarketRemoval(log types.Log) (*MorphoVaultSubmitMarketRemoval, error) {
	event := new(MorphoVaultSubmitMarketRemoval)
	if err := _MorphoVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultSubmitTimelockIterator is returned from FilterSubmitTimelock and is used to iterate over the raw logs and unpacked data for SubmitTimelock events raised by the MorphoVault contract.
type MorphoVaultSubmitTimelockIterator struct {
	Event *MorphoVaultSubmitTimelock // Event containing the contract specifics and raw log

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
func (it *MorphoVaultSubmitTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultSubmitTimelock)
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
		it.Event = new(MorphoVaultSubmitTimelock)
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
func (it *MorphoVaultSubmitTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultSubmitTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultSubmitTimelock represents a SubmitTimelock event raised by the MorphoVault contract.
type MorphoVaultSubmitTimelock struct {
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTimelock is a free log retrieval operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) FilterSubmitTimelock(opts *bind.FilterOpts) (*MorphoVaultSubmitTimelockIterator, error) {

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return &MorphoVaultSubmitTimelockIterator{contract: _MorphoVault.contract, event: "SubmitTimelock", logs: logs, sub: sub}, nil
}

// WatchSubmitTimelock is a free log subscription operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) WatchSubmitTimelock(opts *bind.WatchOpts, sink chan<- *MorphoVaultSubmitTimelock) (event.Subscription, error) {

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultSubmitTimelock)
				if err := _MorphoVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
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

// ParseSubmitTimelock is a log parse operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MorphoVault *MorphoVaultFilterer) ParseSubmitTimelock(log types.Log) (*MorphoVaultSubmitTimelock, error) {
	event := new(MorphoVaultSubmitTimelock)
	if err := _MorphoVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MorphoVault contract.
type MorphoVaultTransferIterator struct {
	Event *MorphoVaultTransfer // Event containing the contract specifics and raw log

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
func (it *MorphoVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultTransfer)
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
		it.Event = new(MorphoVaultTransfer)
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
func (it *MorphoVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultTransfer represents a Transfer event raised by the MorphoVault contract.
type MorphoVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MorphoVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultTransferIterator{contract: _MorphoVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MorphoVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultTransfer)
				if err := _MorphoVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphoVault *MorphoVaultFilterer) ParseTransfer(log types.Log) (*MorphoVaultTransfer, error) {
	event := new(MorphoVaultTransfer)
	if err := _MorphoVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultUpdateLastTotalAssetsIterator is returned from FilterUpdateLastTotalAssets and is used to iterate over the raw logs and unpacked data for UpdateLastTotalAssets events raised by the MorphoVault contract.
type MorphoVaultUpdateLastTotalAssetsIterator struct {
	Event *MorphoVaultUpdateLastTotalAssets // Event containing the contract specifics and raw log

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
func (it *MorphoVaultUpdateLastTotalAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultUpdateLastTotalAssets)
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
		it.Event = new(MorphoVaultUpdateLastTotalAssets)
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
func (it *MorphoVaultUpdateLastTotalAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultUpdateLastTotalAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultUpdateLastTotalAssets represents a UpdateLastTotalAssets event raised by the MorphoVault contract.
type MorphoVaultUpdateLastTotalAssets struct {
	UpdatedTotalAssets *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateLastTotalAssets is a free log retrieval operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MorphoVault *MorphoVaultFilterer) FilterUpdateLastTotalAssets(opts *bind.FilterOpts) (*MorphoVaultUpdateLastTotalAssetsIterator, error) {

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return &MorphoVaultUpdateLastTotalAssetsIterator{contract: _MorphoVault.contract, event: "UpdateLastTotalAssets", logs: logs, sub: sub}, nil
}

// WatchUpdateLastTotalAssets is a free log subscription operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MorphoVault *MorphoVaultFilterer) WatchUpdateLastTotalAssets(opts *bind.WatchOpts, sink chan<- *MorphoVaultUpdateLastTotalAssets) (event.Subscription, error) {

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultUpdateLastTotalAssets)
				if err := _MorphoVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
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

// ParseUpdateLastTotalAssets is a log parse operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MorphoVault *MorphoVaultFilterer) ParseUpdateLastTotalAssets(log types.Log) (*MorphoVaultUpdateLastTotalAssets, error) {
	event := new(MorphoVaultUpdateLastTotalAssets)
	if err := _MorphoVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MorphoVault contract.
type MorphoVaultWithdrawIterator struct {
	Event *MorphoVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *MorphoVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoVaultWithdraw)
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
		it.Event = new(MorphoVaultWithdraw)
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
func (it *MorphoVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoVaultWithdraw represents a Withdraw event raised by the MorphoVault contract.
type MorphoVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*MorphoVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MorphoVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MorphoVaultWithdrawIterator{contract: _MorphoVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MorphoVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MorphoVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoVaultWithdraw)
				if err := _MorphoVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MorphoVault *MorphoVaultFilterer) ParseWithdraw(log types.Log) (*MorphoVaultWithdraw, error) {
	event := new(MorphoVaultWithdraw)
	if err := _MorphoVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
