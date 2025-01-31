// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_earn_implementation

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

// CheckpointsCheckpoint208 is an auto generated low-level Go binding around an user-defined struct.
type CheckpointsCheckpoint208 struct {
	Key   *big.Int
	Value *big.Int
}

// IEulerEarnDeploymentParams is an auto generated low-level Go binding around an user-defined struct.
type IEulerEarnDeploymentParams struct {
	EulerEarnVaultModule  common.Address
	RewardsModule         common.Address
	HooksModule           common.Address
	FeeModule             common.Address
	StrategyModule        common.Address
	WithdrawalQueueModule common.Address
}

// IEulerEarnInitParams is an auto generated low-level Go binding around an user-defined struct.
type IEulerEarnInitParams struct {
	EulerEarnVaultOwner         common.Address
	Asset                       common.Address
	Name                        string
	Symbol                      string
	InitialCashAllocationPoints *big.Int
	SmearingPeriod              *big.Int
}

// IEulerEarnStrategy is an auto generated low-level Go binding around an user-defined struct.
type IEulerEarnStrategy struct {
	Allocated        *big.Int
	AllocationPoints *big.Int
	Cap              uint16
	Status           uint8
}

// SharedIntegrationsParams is an auto generated low-level Go binding around an user-defined struct.
type SharedIntegrationsParams struct {
	Evc                      common.Address
	BalanceTracker           common.Address
	Permit2                  common.Address
	IsHarvestCoolDownCheckOn bool
}

// EulerEarnImplementationMetaData contains all meta data concerning the EulerEarnImplementation contract.
var EulerEarnImplementationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isHarvestCoolDownCheckOn\",\"type\":\"bool\"}],\"internalType\":\"structShared.IntegrationsParams\",\"name\":\"_integrationsParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"eulerEarnVaultModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hooksModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategyModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawalQueueModule\",\"type\":\"address\"}],\"internalType\":\"structIEulerEarn.DeploymentParams\",\"name\":\"_deploymentParams\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededSafeSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededSafeSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"clock\",\"type\":\"uint48\"}],\"name\":\"ERC5805FutureLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC6372InconsistentClock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialAllocationPointsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAssetAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmearingPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewReentrancy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"VotesExpiredSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotes\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocationPoints\",\"type\":\"uint256\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPoints\",\"type\":\"uint256\"}],\"name\":\"adjustAllocationPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"_key\",\"type\":\"uint48\"},{\"internalType\":\"uint208\",\"name\":\"_value\",\"type\":\"uint208\"}],\"internalType\":\"structCheckpoints.Checkpoint208\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_forfeitRecentReward\",\"type\":\"bool\"}],\"name\":\"claimStrategyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_forfeitRecentReward\",\"type\":\"bool\"}],\"name\":\"disableRewardForStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"}],\"name\":\"enableRewardForStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eulerEarnVaultModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEulerEarnSavingRate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"uint168\",\"name\":\"\",\"type\":\"uint168\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHooksConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timepoint\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timepoint\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"getStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint120\",\"name\":\"allocated\",\"type\":\"uint120\"},{\"internalType\":\"uint96\",\"name\":\"allocationPoints\",\"type\":\"uint96\"},{\"internalType\":\"AmountCap\",\"name\":\"cap\",\"type\":\"uint16\"},{\"internalType\":\"enumIEulerEarn.StrategyStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIEulerEarn.Strategy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hooksModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"eulerEarnVaultOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialCashAllocationPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"smearingPeriod\",\"type\":\"uint24\"}],\"internalType\":\"structIEulerEarn.InitParams\",\"name\":\"_initParams\",\"type\":\"tuple\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestSmearingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCheckingHarvestCoolDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHarvestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"optInStrategyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"optOutStrategyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_strategies\",\"type\":\"address[]\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"removeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_index1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_index2\",\"type\":\"uint8\"}],\"name\":\"reorderWithdrawalQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hooksTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_hookedFns\",\"type\":\"uint32\"}],\"name\":\"setHooksConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_newFee\",\"type\":\"uint96\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_cap\",\"type\":\"uint16\"}],\"name\":\"setStrategyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"toggleStrategyEmergencyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocationPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssetsAllocatable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssetsDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterestAccrued\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalQueueModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerEarnImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEarnImplementationMetaData.ABI instead.
var EulerEarnImplementationABI = EulerEarnImplementationMetaData.ABI

// EulerEarnImplementation is an auto generated Go binding around an Ethereum contract.
type EulerEarnImplementation struct {
	EulerEarnImplementationCaller     // Read-only binding to the contract
	EulerEarnImplementationTransactor // Write-only binding to the contract
	EulerEarnImplementationFilterer   // Log filterer for contract events
}

// EulerEarnImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEarnImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEarnImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEarnImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEarnImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEarnImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEarnImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEarnImplementationSession struct {
	Contract     *EulerEarnImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EulerEarnImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEarnImplementationCallerSession struct {
	Contract *EulerEarnImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// EulerEarnImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEarnImplementationTransactorSession struct {
	Contract     *EulerEarnImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// EulerEarnImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEarnImplementationRaw struct {
	Contract *EulerEarnImplementation // Generic contract binding to access the raw methods on
}

// EulerEarnImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEarnImplementationCallerRaw struct {
	Contract *EulerEarnImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEarnImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEarnImplementationTransactorRaw struct {
	Contract *EulerEarnImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEarnImplementation creates a new instance of EulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEarnImplementation(address common.Address, backend bind.ContractBackend) (*EulerEarnImplementation, error) {
	contract, err := bindEulerEarnImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementation{EulerEarnImplementationCaller: EulerEarnImplementationCaller{contract: contract}, EulerEarnImplementationTransactor: EulerEarnImplementationTransactor{contract: contract}, EulerEarnImplementationFilterer: EulerEarnImplementationFilterer{contract: contract}}, nil
}

// NewEulerEarnImplementationCaller creates a new read-only instance of EulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEarnImplementationCaller(address common.Address, caller bind.ContractCaller) (*EulerEarnImplementationCaller, error) {
	contract, err := bindEulerEarnImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationCaller{contract: contract}, nil
}

// NewEulerEarnImplementationTransactor creates a new write-only instance of EulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEarnImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEarnImplementationTransactor, error) {
	contract, err := bindEulerEarnImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationTransactor{contract: contract}, nil
}

// NewEulerEarnImplementationFilterer creates a new log filterer instance of EulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEarnImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEarnImplementationFilterer, error) {
	contract, err := bindEulerEarnImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationFilterer{contract: contract}, nil
}

// bindEulerEarnImplementation binds a generic wrapper to an already deployed contract.
func bindEulerEarnImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEarnImplementationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEarnImplementation *EulerEarnImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEarnImplementation.Contract.EulerEarnImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEarnImplementation *EulerEarnImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EulerEarnImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEarnImplementation *EulerEarnImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EulerEarnImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEarnImplementation *EulerEarnImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEarnImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEarnImplementation *EulerEarnImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEarnImplementation *EulerEarnImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationSession) CLOCKMODE() (string, error) {
	return _EulerEarnImplementation.Contract.CLOCKMODE(&_EulerEarnImplementation.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) CLOCKMODE() (string, error) {
	return _EulerEarnImplementation.Contract.CLOCKMODE(&_EulerEarnImplementation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EulerEarnImplementation.Contract.DEFAULTADMINROLE(&_EulerEarnImplementation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EulerEarnImplementation.Contract.DEFAULTADMINROLE(&_EulerEarnImplementation.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) EVC() (common.Address, error) {
	return _EulerEarnImplementation.Contract.EVC(&_EulerEarnImplementation.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) EVC() (common.Address, error) {
	return _EulerEarnImplementation.Contract.EVC(&_EulerEarnImplementation.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Allowance(&_EulerEarnImplementation.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Allowance(&_EulerEarnImplementation.CallOpts, _owner, _spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Asset() (common.Address, error) {
	return _EulerEarnImplementation.Contract.Asset(&_EulerEarnImplementation.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Asset() (common.Address, error) {
	return _EulerEarnImplementation.Contract.Asset(&_EulerEarnImplementation.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) BalanceForwarderEnabled(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "balanceForwarderEnabled", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) BalanceForwarderEnabled(_account common.Address) (bool, error) {
	return _EulerEarnImplementation.Contract.BalanceForwarderEnabled(&_EulerEarnImplementation.CallOpts, _account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) BalanceForwarderEnabled(_account common.Address) (bool, error) {
	return _EulerEarnImplementation.Contract.BalanceForwarderEnabled(&_EulerEarnImplementation.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.BalanceOf(&_EulerEarnImplementation.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.BalanceOf(&_EulerEarnImplementation.CallOpts, _account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEarnImplementation.Contract.BalanceTrackerAddress(&_EulerEarnImplementation.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEarnImplementation.Contract.BalanceTrackerAddress(&_EulerEarnImplementation.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Checkpoints(opts *bind.CallOpts, _account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "checkpoints", _account, _pos)

	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}

	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEarnImplementation *EulerEarnImplementationSession) Checkpoints(_account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	return _EulerEarnImplementation.Contract.Checkpoints(&_EulerEarnImplementation.CallOpts, _account, _pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Checkpoints(_account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	return _EulerEarnImplementation.Contract.Checkpoints(&_EulerEarnImplementation.CallOpts, _account, _pos)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Clock() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Clock(&_EulerEarnImplementation.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Clock() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Clock(&_EulerEarnImplementation.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.ConvertToAssets(&_EulerEarnImplementation.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.ConvertToAssets(&_EulerEarnImplementation.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.ConvertToShares(&_EulerEarnImplementation.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.ConvertToShares(&_EulerEarnImplementation.CallOpts, _assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Decimals() (uint8, error) {
	return _EulerEarnImplementation.Contract.Decimals(&_EulerEarnImplementation.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Decimals() (uint8, error) {
	return _EulerEarnImplementation.Contract.Decimals(&_EulerEarnImplementation.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Delegates(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "delegates", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Delegates(_account common.Address) (common.Address, error) {
	return _EulerEarnImplementation.Contract.Delegates(&_EulerEarnImplementation.CallOpts, _account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Delegates(_account common.Address) (common.Address, error) {
	return _EulerEarnImplementation.Contract.Delegates(&_EulerEarnImplementation.CallOpts, _account)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "eip712Domain")

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
func (_EulerEarnImplementation *EulerEarnImplementationSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EulerEarnImplementation.Contract.Eip712Domain(&_EulerEarnImplementation.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EulerEarnImplementation.Contract.Eip712Domain(&_EulerEarnImplementation.CallOpts)
}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) EulerEarnVaultModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "eulerEarnVaultModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) EulerEarnVaultModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.EulerEarnVaultModule(&_EulerEarnImplementation.CallOpts)
}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) EulerEarnVaultModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.EulerEarnVaultModule(&_EulerEarnImplementation.CallOpts)
}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) FeeModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "feeModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) FeeModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.FeeModule(&_EulerEarnImplementation.CallOpts)
}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) FeeModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.FeeModule(&_EulerEarnImplementation.CallOpts)
}

// GetEulerEarnSavingRate is a free data retrieval call binding the contract method 0xbeb9c69b.
//
// Solidity: function getEulerEarnSavingRate() view returns(uint40, uint40, uint168)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetEulerEarnSavingRate(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getEulerEarnSavingRate")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetEulerEarnSavingRate is a free data retrieval call binding the contract method 0xbeb9c69b.
//
// Solidity: function getEulerEarnSavingRate() view returns(uint40, uint40, uint168)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetEulerEarnSavingRate() (*big.Int, *big.Int, *big.Int, error) {
	return _EulerEarnImplementation.Contract.GetEulerEarnSavingRate(&_EulerEarnImplementation.CallOpts)
}

// GetEulerEarnSavingRate is a free data retrieval call binding the contract method 0xbeb9c69b.
//
// Solidity: function getEulerEarnSavingRate() view returns(uint40, uint40, uint168)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetEulerEarnSavingRate() (*big.Int, *big.Int, *big.Int, error) {
	return _EulerEarnImplementation.Contract.GetEulerEarnSavingRate(&_EulerEarnImplementation.CallOpts)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xe6f2aae5.
//
// Solidity: function getHooksConfig() view returns(address, uint32)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetHooksConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getHooksConfig")

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetHooksConfig is a free data retrieval call binding the contract method 0xe6f2aae5.
//
// Solidity: function getHooksConfig() view returns(address, uint32)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetHooksConfig() (common.Address, uint32, error) {
	return _EulerEarnImplementation.Contract.GetHooksConfig(&_EulerEarnImplementation.CallOpts)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xe6f2aae5.
//
// Solidity: function getHooksConfig() view returns(address, uint32)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetHooksConfig() (common.Address, uint32, error) {
	return _EulerEarnImplementation.Contract.GetHooksConfig(&_EulerEarnImplementation.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetPastTotalSupply(opts *bind.CallOpts, _timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getPastTotalSupply", _timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetPastTotalSupply(_timepoint *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetPastTotalSupply(&_EulerEarnImplementation.CallOpts, _timepoint)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetPastTotalSupply(_timepoint *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetPastTotalSupply(&_EulerEarnImplementation.CallOpts, _timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetPastVotes(opts *bind.CallOpts, _account common.Address, _timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getPastVotes", _account, _timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetPastVotes(_account common.Address, _timepoint *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetPastVotes(&_EulerEarnImplementation.CallOpts, _account, _timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetPastVotes(_account common.Address, _timepoint *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetPastVotes(&_EulerEarnImplementation.CallOpts, _account, _timepoint)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EulerEarnImplementation.Contract.GetRoleAdmin(&_EulerEarnImplementation.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EulerEarnImplementation.Contract.GetRoleAdmin(&_EulerEarnImplementation.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EulerEarnImplementation.Contract.GetRoleMember(&_EulerEarnImplementation.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EulerEarnImplementation.Contract.GetRoleMember(&_EulerEarnImplementation.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetRoleMemberCount(&_EulerEarnImplementation.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetRoleMemberCount(&_EulerEarnImplementation.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EulerEarnImplementation.Contract.GetRoleMembers(&_EulerEarnImplementation.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EulerEarnImplementation.Contract.GetRoleMembers(&_EulerEarnImplementation.CallOpts, role)
}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetStrategy(opts *bind.CallOpts, _strategy common.Address) (IEulerEarnStrategy, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getStrategy", _strategy)

	if err != nil {
		return *new(IEulerEarnStrategy), err
	}

	out0 := *abi.ConvertType(out[0], new(IEulerEarnStrategy)).(*IEulerEarnStrategy)

	return out0, err

}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetStrategy(_strategy common.Address) (IEulerEarnStrategy, error) {
	return _EulerEarnImplementation.Contract.GetStrategy(&_EulerEarnImplementation.CallOpts, _strategy)
}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetStrategy(_strategy common.Address) (IEulerEarnStrategy, error) {
	return _EulerEarnImplementation.Contract.GetStrategy(&_EulerEarnImplementation.CallOpts, _strategy)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) GetVotes(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "getVotes", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) GetVotes(_account common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetVotes(&_EulerEarnImplementation.CallOpts, _account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) GetVotes(_account common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.GetVotes(&_EulerEarnImplementation.CallOpts, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EulerEarnImplementation.Contract.HasRole(&_EulerEarnImplementation.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EulerEarnImplementation.Contract.HasRole(&_EulerEarnImplementation.CallOpts, role, account)
}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) HooksModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "hooksModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) HooksModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.HooksModule(&_EulerEarnImplementation.CallOpts)
}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) HooksModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.HooksModule(&_EulerEarnImplementation.CallOpts)
}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) InterestAccrued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "interestAccrued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) InterestAccrued() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.InterestAccrued(&_EulerEarnImplementation.CallOpts)
}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) InterestAccrued() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.InterestAccrued(&_EulerEarnImplementation.CallOpts)
}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) InterestSmearingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "interestSmearingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) InterestSmearingPeriod() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.InterestSmearingPeriod(&_EulerEarnImplementation.CallOpts)
}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) InterestSmearingPeriod() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.InterestSmearingPeriod(&_EulerEarnImplementation.CallOpts)
}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) IsCheckingHarvestCoolDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "isCheckingHarvestCoolDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) IsCheckingHarvestCoolDown() (bool, error) {
	return _EulerEarnImplementation.Contract.IsCheckingHarvestCoolDown(&_EulerEarnImplementation.CallOpts)
}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) IsCheckingHarvestCoolDown() (bool, error) {
	return _EulerEarnImplementation.Contract.IsCheckingHarvestCoolDown(&_EulerEarnImplementation.CallOpts)
}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) LastHarvestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "lastHarvestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) LastHarvestTimestamp() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.LastHarvestTimestamp(&_EulerEarnImplementation.CallOpts)
}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) LastHarvestTimestamp() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.LastHarvestTimestamp(&_EulerEarnImplementation.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) MaxDeposit(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "maxDeposit", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxDeposit(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxDeposit(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) MaxMint(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "maxMint", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxMint(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxMint(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxRedeem(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxRedeem(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxWithdraw(&_EulerEarnImplementation.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.MaxWithdraw(&_EulerEarnImplementation.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Name() (string, error) {
	return _EulerEarnImplementation.Contract.Name(&_EulerEarnImplementation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Name() (string, error) {
	return _EulerEarnImplementation.Contract.Name(&_EulerEarnImplementation.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Nonces(owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Nonces(&_EulerEarnImplementation.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.Nonces(&_EulerEarnImplementation.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) NumCheckpoints(opts *bind.CallOpts, _account common.Address) (uint32, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "numCheckpoints", _account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEarnImplementation *EulerEarnImplementationSession) NumCheckpoints(_account common.Address) (uint32, error) {
	return _EulerEarnImplementation.Contract.NumCheckpoints(&_EulerEarnImplementation.CallOpts, _account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) NumCheckpoints(_account common.Address) (uint32, error) {
	return _EulerEarnImplementation.Contract.NumCheckpoints(&_EulerEarnImplementation.CallOpts, _account)
}

// PerformanceFeeConfig is a free data retrieval call binding the contract method 0x3eda8287.
//
// Solidity: function performanceFeeConfig() view returns(address, uint96)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) PerformanceFeeConfig(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "performanceFeeConfig")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PerformanceFeeConfig is a free data retrieval call binding the contract method 0x3eda8287.
//
// Solidity: function performanceFeeConfig() view returns(address, uint96)
func (_EulerEarnImplementation *EulerEarnImplementationSession) PerformanceFeeConfig() (common.Address, *big.Int, error) {
	return _EulerEarnImplementation.Contract.PerformanceFeeConfig(&_EulerEarnImplementation.CallOpts)
}

// PerformanceFeeConfig is a free data retrieval call binding the contract method 0x3eda8287.
//
// Solidity: function performanceFeeConfig() view returns(address, uint96)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) PerformanceFeeConfig() (common.Address, *big.Int, error) {
	return _EulerEarnImplementation.Contract.PerformanceFeeConfig(&_EulerEarnImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Permit2Address() (common.Address, error) {
	return _EulerEarnImplementation.Contract.Permit2Address(&_EulerEarnImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Permit2Address() (common.Address, error) {
	return _EulerEarnImplementation.Contract.Permit2Address(&_EulerEarnImplementation.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewDeposit(&_EulerEarnImplementation.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewDeposit(&_EulerEarnImplementation.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewMint(&_EulerEarnImplementation.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewMint(&_EulerEarnImplementation.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) PreviewRedeem(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "previewRedeem", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewRedeem(&_EulerEarnImplementation.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewRedeem(&_EulerEarnImplementation.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewWithdraw(&_EulerEarnImplementation.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _EulerEarnImplementation.Contract.PreviewWithdraw(&_EulerEarnImplementation.CallOpts, _assets)
}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) RewardsModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "rewardsModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) RewardsModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.RewardsModule(&_EulerEarnImplementation.CallOpts)
}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) RewardsModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.RewardsModule(&_EulerEarnImplementation.CallOpts)
}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) StrategyModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "strategyModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) StrategyModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.StrategyModule(&_EulerEarnImplementation.CallOpts)
}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) StrategyModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.StrategyModule(&_EulerEarnImplementation.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EulerEarnImplementation.Contract.SupportsInterface(&_EulerEarnImplementation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EulerEarnImplementation.Contract.SupportsInterface(&_EulerEarnImplementation.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Symbol() (string, error) {
	return _EulerEarnImplementation.Contract.Symbol(&_EulerEarnImplementation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) Symbol() (string, error) {
	return _EulerEarnImplementation.Contract.Symbol(&_EulerEarnImplementation.CallOpts)
}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalAllocated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalAllocated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalAllocated() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAllocated(&_EulerEarnImplementation.CallOpts)
}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalAllocated() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAllocated(&_EulerEarnImplementation.CallOpts)
}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalAllocationPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalAllocationPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalAllocationPoints() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAllocationPoints(&_EulerEarnImplementation.CallOpts)
}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalAllocationPoints() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAllocationPoints(&_EulerEarnImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalAssets() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssets(&_EulerEarnImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssets(&_EulerEarnImplementation.CallOpts)
}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalAssetsAllocatable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalAssetsAllocatable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalAssetsAllocatable() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssetsAllocatable(&_EulerEarnImplementation.CallOpts)
}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalAssetsAllocatable() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssetsAllocatable(&_EulerEarnImplementation.CallOpts)
}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalAssetsDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalAssetsDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalAssetsDeposited() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssetsDeposited(&_EulerEarnImplementation.CallOpts)
}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalAssetsDeposited() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalAssetsDeposited(&_EulerEarnImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TotalSupply() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalSupply(&_EulerEarnImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerEarnImplementation.Contract.TotalSupply(&_EulerEarnImplementation.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationCaller) WithdrawalQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "withdrawalQueue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationSession) WithdrawalQueue() ([]common.Address, error) {
	return _EulerEarnImplementation.Contract.WithdrawalQueue(&_EulerEarnImplementation.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) WithdrawalQueue() ([]common.Address, error) {
	return _EulerEarnImplementation.Contract.WithdrawalQueue(&_EulerEarnImplementation.CallOpts)
}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCaller) WithdrawalQueueModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEarnImplementation.contract.Call(opts, &out, "withdrawalQueueModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationSession) WithdrawalQueueModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.WithdrawalQueueModule(&_EulerEarnImplementation.CallOpts)
}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEarnImplementation *EulerEarnImplementationCallerSession) WithdrawalQueueModule() (common.Address, error) {
	return _EulerEarnImplementation.Contract.WithdrawalQueueModule(&_EulerEarnImplementation.CallOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) AddStrategy(opts *bind.TransactOpts, _strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "addStrategy", _strategy, _allocationPoints)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) AddStrategy(_strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.AddStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _allocationPoints)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) AddStrategy(_strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.AddStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _allocationPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) AdjustAllocationPoints(opts *bind.TransactOpts, _strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "adjustAllocationPoints", _strategy, _newPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) AdjustAllocationPoints(_strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.AdjustAllocationPoints(&_EulerEarnImplementation.TransactOpts, _strategy, _newPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) AdjustAllocationPoints(_strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.AdjustAllocationPoints(&_EulerEarnImplementation.TransactOpts, _strategy, _newPoints)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Approve(&_EulerEarnImplementation.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Approve(&_EulerEarnImplementation.TransactOpts, _spender, _value)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) ClaimStrategyReward(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "claimStrategyReward", _strategy, _reward, _recipient, _forfeitRecentReward)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) ClaimStrategyReward(_strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ClaimStrategyReward(&_EulerEarnImplementation.TransactOpts, _strategy, _reward, _recipient, _forfeitRecentReward)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) ClaimStrategyReward(_strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ClaimStrategyReward(&_EulerEarnImplementation.TransactOpts, _strategy, _reward, _recipient, _forfeitRecentReward)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Delegate(opts *bind.TransactOpts, _delegatee common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "delegate", _delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Delegate(_delegatee common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Delegate(&_EulerEarnImplementation.TransactOpts, _delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Delegate(_delegatee common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Delegate(&_EulerEarnImplementation.TransactOpts, _delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) DelegateBySig(opts *bind.TransactOpts, _delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "delegateBySig", _delegatee, _nonce, _expiry, _v, _r, _s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) DelegateBySig(_delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DelegateBySig(&_EulerEarnImplementation.TransactOpts, _delegatee, _nonce, _expiry, _v, _r, _s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) DelegateBySig(_delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DelegateBySig(&_EulerEarnImplementation.TransactOpts, _delegatee, _nonce, _expiry, _v, _r, _s)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Deposit(&_EulerEarnImplementation.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Deposit(&_EulerEarnImplementation.TransactOpts, _assets, _receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DisableBalanceForwarder(&_EulerEarnImplementation.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DisableBalanceForwarder(&_EulerEarnImplementation.TransactOpts)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) DisableRewardForStrategy(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "disableRewardForStrategy", _strategy, _reward, _forfeitRecentReward)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) DisableRewardForStrategy(_strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DisableRewardForStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _reward, _forfeitRecentReward)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) DisableRewardForStrategy(_strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.DisableRewardForStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _reward, _forfeitRecentReward)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EnableBalanceForwarder(&_EulerEarnImplementation.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EnableBalanceForwarder(&_EulerEarnImplementation.TransactOpts)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) EnableRewardForStrategy(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "enableRewardForStrategy", _strategy, _reward)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) EnableRewardForStrategy(_strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EnableRewardForStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _reward)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) EnableRewardForStrategy(_strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.EnableRewardForStrategy(&_EulerEarnImplementation.TransactOpts, _strategy, _reward)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) GrantRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "grantRole", _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.GrantRole(&_EulerEarnImplementation.TransactOpts, _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.GrantRole(&_EulerEarnImplementation.TransactOpts, _role, _account)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Gulp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "gulp")
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Gulp() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Gulp(&_EulerEarnImplementation.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Gulp() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Gulp(&_EulerEarnImplementation.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Harvest() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Harvest(&_EulerEarnImplementation.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Harvest() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Harvest(&_EulerEarnImplementation.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Init(opts *bind.TransactOpts, _initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "init", _initParams)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Init(_initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Init(&_EulerEarnImplementation.TransactOpts, _initParams)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Init(_initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Init(&_EulerEarnImplementation.TransactOpts, _initParams)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Mint(&_EulerEarnImplementation.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Mint(&_EulerEarnImplementation.TransactOpts, _shares, _receiver)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) OptInStrategyRewards(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "optInStrategyRewards", _strategy)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) OptInStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.OptInStrategyRewards(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) OptInStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.OptInStrategyRewards(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) OptOutStrategyRewards(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "optOutStrategyRewards", _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) OptOutStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.OptOutStrategyRewards(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) OptOutStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.OptOutStrategyRewards(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Rebalance(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "rebalance", _strategies)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Rebalance(_strategies []common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Rebalance(&_EulerEarnImplementation.TransactOpts, _strategies)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Rebalance(_strategies []common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Rebalance(&_EulerEarnImplementation.TransactOpts, _strategies)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Redeem(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "redeem", _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Redeem(&_EulerEarnImplementation.TransactOpts, _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Redeem(&_EulerEarnImplementation.TransactOpts, _shares, _receiver, _owner)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) RemoveStrategy(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "removeStrategy", _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RemoveStrategy(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RemoveStrategy(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) RenounceRole(opts *bind.TransactOpts, _role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "renounceRole", _role, _callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) RenounceRole(_role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RenounceRole(&_EulerEarnImplementation.TransactOpts, _role, _callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) RenounceRole(_role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RenounceRole(&_EulerEarnImplementation.TransactOpts, _role, _callerConfirmation)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) ReorderWithdrawalQueue(opts *bind.TransactOpts, _index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "reorderWithdrawalQueue", _index1, _index2)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) ReorderWithdrawalQueue(_index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ReorderWithdrawalQueue(&_EulerEarnImplementation.TransactOpts, _index1, _index2)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) ReorderWithdrawalQueue(_index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ReorderWithdrawalQueue(&_EulerEarnImplementation.TransactOpts, _index1, _index2)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) RevokeRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "revokeRole", _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RevokeRole(&_EulerEarnImplementation.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.RevokeRole(&_EulerEarnImplementation.TransactOpts, _role, _account)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) SetFeeRecipient(opts *bind.TransactOpts, _newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "setFeeRecipient", _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetFeeRecipient(&_EulerEarnImplementation.TransactOpts, _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetFeeRecipient(&_EulerEarnImplementation.TransactOpts, _newFeeRecipient)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) SetHooksConfig(opts *bind.TransactOpts, _hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "setHooksConfig", _hooksTarget, _hookedFns)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) SetHooksConfig(_hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetHooksConfig(&_EulerEarnImplementation.TransactOpts, _hooksTarget, _hookedFns)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) SetHooksConfig(_hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetHooksConfig(&_EulerEarnImplementation.TransactOpts, _hooksTarget, _hookedFns)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) SetPerformanceFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "setPerformanceFee", _newFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) SetPerformanceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetPerformanceFee(&_EulerEarnImplementation.TransactOpts, _newFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) SetPerformanceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetPerformanceFee(&_EulerEarnImplementation.TransactOpts, _newFee)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) SetStrategyCap(opts *bind.TransactOpts, _strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "setStrategyCap", _strategy, _cap)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) SetStrategyCap(_strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetStrategyCap(&_EulerEarnImplementation.TransactOpts, _strategy, _cap)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) SetStrategyCap(_strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.SetStrategyCap(&_EulerEarnImplementation.TransactOpts, _strategy, _cap)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Skim(opts *bind.TransactOpts, _token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "skim", _token, _recipient)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) Skim(_token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Skim(&_EulerEarnImplementation.TransactOpts, _token, _recipient)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Skim(_token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Skim(&_EulerEarnImplementation.TransactOpts, _token, _recipient)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) ToggleStrategyEmergencyStatus(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "toggleStrategyEmergencyStatus", _strategy)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) ToggleStrategyEmergencyStatus(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ToggleStrategyEmergencyStatus(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) ToggleStrategyEmergencyStatus(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.ToggleStrategyEmergencyStatus(&_EulerEarnImplementation.TransactOpts, _strategy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Transfer(&_EulerEarnImplementation.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Transfer(&_EulerEarnImplementation.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.TransferFrom(&_EulerEarnImplementation.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.TransferFrom(&_EulerEarnImplementation.TransactOpts, _from, _to, _value)
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) UpdateInterestAccrued(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "updateInterestAccrued")
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEarnImplementation *EulerEarnImplementationSession) UpdateInterestAccrued() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.UpdateInterestAccrued(&_EulerEarnImplementation.TransactOpts)
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) UpdateInterestAccrued() (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.UpdateInterestAccrued(&_EulerEarnImplementation.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationTransactor) Withdraw(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.contract.Transact(opts, "withdraw", _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Withdraw(&_EulerEarnImplementation.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationTransactorSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEarnImplementation.Contract.Withdraw(&_EulerEarnImplementation.TransactOpts, _assets, _receiver, _owner)
}

// EulerEarnImplementationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationApprovalIterator struct {
	Event *EulerEarnImplementationApproval // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationApproval)
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
		it.Event = new(EulerEarnImplementationApproval)
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
func (it *EulerEarnImplementationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationApproval represents a Approval event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerEarnImplementationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationApprovalIterator{contract: _EulerEarnImplementation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationApproval)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseApproval(log types.Log) (*EulerEarnImplementationApproval, error) {
	event := new(EulerEarnImplementationApproval)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDelegateChangedIterator struct {
	Event *EulerEarnImplementationDelegateChanged // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationDelegateChanged)
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
		it.Event = new(EulerEarnImplementationDelegateChanged)
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
func (it *EulerEarnImplementationDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationDelegateChanged represents a DelegateChanged event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*EulerEarnImplementationDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationDelegateChangedIterator{contract: _EulerEarnImplementation.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationDelegateChanged)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseDelegateChanged(log types.Log) (*EulerEarnImplementationDelegateChanged, error) {
	event := new(EulerEarnImplementationDelegateChanged)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDelegateVotesChangedIterator struct {
	Event *EulerEarnImplementationDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationDelegateVotesChanged)
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
		it.Event = new(EulerEarnImplementationDelegateVotesChanged)
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
func (it *EulerEarnImplementationDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationDelegateVotesChanged represents a DelegateVotesChanged event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*EulerEarnImplementationDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationDelegateVotesChangedIterator{contract: _EulerEarnImplementation.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationDelegateVotesChanged)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseDelegateVotesChanged(log types.Log) (*EulerEarnImplementationDelegateVotesChanged, error) {
	event := new(EulerEarnImplementationDelegateVotesChanged)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDepositIterator struct {
	Event *EulerEarnImplementationDeposit // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationDeposit)
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
		it.Event = new(EulerEarnImplementationDeposit)
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
func (it *EulerEarnImplementationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationDeposit represents a Deposit event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerEarnImplementationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationDepositIterator{contract: _EulerEarnImplementation.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationDeposit)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseDeposit(log types.Log) (*EulerEarnImplementationDeposit, error) {
	event := new(EulerEarnImplementationDeposit)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationEIP712DomainChangedIterator struct {
	Event *EulerEarnImplementationEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationEIP712DomainChanged)
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
		it.Event = new(EulerEarnImplementationEIP712DomainChanged)
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
func (it *EulerEarnImplementationEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationEIP712DomainChanged represents a EIP712DomainChanged event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EulerEarnImplementationEIP712DomainChangedIterator, error) {

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationEIP712DomainChangedIterator{contract: _EulerEarnImplementation.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationEIP712DomainChanged)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseEIP712DomainChanged(log types.Log) (*EulerEarnImplementationEIP712DomainChanged, error) {
	event := new(EulerEarnImplementationEIP712DomainChanged)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationInitializedIterator struct {
	Event *EulerEarnImplementationInitialized // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationInitialized)
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
		it.Event = new(EulerEarnImplementationInitialized)
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
func (it *EulerEarnImplementationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationInitialized represents a Initialized event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterInitialized(opts *bind.FilterOpts) (*EulerEarnImplementationInitializedIterator, error) {

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationInitializedIterator{contract: _EulerEarnImplementation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationInitialized) (event.Subscription, error) {

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationInitialized)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseInitialized(log types.Log) (*EulerEarnImplementationInitialized, error) {
	event := new(EulerEarnImplementationInitialized)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleAdminChangedIterator struct {
	Event *EulerEarnImplementationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationRoleAdminChanged)
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
		it.Event = new(EulerEarnImplementationRoleAdminChanged)
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
func (it *EulerEarnImplementationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationRoleAdminChanged represents a RoleAdminChanged event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EulerEarnImplementationRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationRoleAdminChangedIterator{contract: _EulerEarnImplementation.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationRoleAdminChanged)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseRoleAdminChanged(log types.Log) (*EulerEarnImplementationRoleAdminChanged, error) {
	event := new(EulerEarnImplementationRoleAdminChanged)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleGrantedIterator struct {
	Event *EulerEarnImplementationRoleGranted // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationRoleGranted)
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
		it.Event = new(EulerEarnImplementationRoleGranted)
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
func (it *EulerEarnImplementationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationRoleGranted represents a RoleGranted event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EulerEarnImplementationRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationRoleGrantedIterator{contract: _EulerEarnImplementation.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationRoleGranted)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseRoleGranted(log types.Log) (*EulerEarnImplementationRoleGranted, error) {
	event := new(EulerEarnImplementationRoleGranted)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleRevokedIterator struct {
	Event *EulerEarnImplementationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationRoleRevoked)
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
		it.Event = new(EulerEarnImplementationRoleRevoked)
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
func (it *EulerEarnImplementationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationRoleRevoked represents a RoleRevoked event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EulerEarnImplementationRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationRoleRevokedIterator{contract: _EulerEarnImplementation.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationRoleRevoked)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseRoleRevoked(log types.Log) (*EulerEarnImplementationRoleRevoked, error) {
	event := new(EulerEarnImplementationRoleRevoked)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationTransferIterator struct {
	Event *EulerEarnImplementationTransfer // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationTransfer)
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
		it.Event = new(EulerEarnImplementationTransfer)
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
func (it *EulerEarnImplementationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationTransfer represents a Transfer event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEarnImplementationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationTransferIterator{contract: _EulerEarnImplementation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationTransfer)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseTransfer(log types.Log) (*EulerEarnImplementationTransfer, error) {
	event := new(EulerEarnImplementationTransfer)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEarnImplementationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerEarnImplementation contract.
type EulerEarnImplementationWithdrawIterator struct {
	Event *EulerEarnImplementationWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerEarnImplementationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEarnImplementationWithdraw)
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
		it.Event = new(EulerEarnImplementationWithdraw)
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
func (it *EulerEarnImplementationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEarnImplementationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEarnImplementationWithdraw represents a Withdraw event raised by the EulerEarnImplementation contract.
type EulerEarnImplementationWithdraw struct {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerEarnImplementationWithdrawIterator, error) {

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

	logs, sub, err := _EulerEarnImplementation.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEarnImplementationWithdrawIterator{contract: _EulerEarnImplementation.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerEarnImplementationWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEarnImplementation.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEarnImplementationWithdraw)
				if err := _EulerEarnImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EulerEarnImplementation *EulerEarnImplementationFilterer) ParseWithdraw(log types.Log) (*EulerEarnImplementationWithdraw, error) {
	event := new(EulerEarnImplementationWithdraw)
	if err := _EulerEarnImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
