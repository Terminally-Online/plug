// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_euler_earn_implementation

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

// EulerEulerEarnImplementationMetaData contains all meta data concerning the EulerEulerEarnImplementation contract.
var EulerEulerEarnImplementationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isHarvestCoolDownCheckOn\",\"type\":\"bool\"}],\"internalType\":\"structShared.IntegrationsParams\",\"name\":\"_integrationsParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"eulerEarnVaultModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hooksModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategyModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawalQueueModule\",\"type\":\"address\"}],\"internalType\":\"structIEulerEarn.DeploymentParams\",\"name\":\"_deploymentParams\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededSafeSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededSafeSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"clock\",\"type\":\"uint48\"}],\"name\":\"ERC5805FutureLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC6372InconsistentClock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EVC_InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialAllocationPointsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAssetAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmearingPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewReentrancy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"VotesExpiredSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotes\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocationPoints\",\"type\":\"uint256\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPoints\",\"type\":\"uint256\"}],\"name\":\"adjustAllocationPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"_key\",\"type\":\"uint48\"},{\"internalType\":\"uint208\",\"name\":\"_value\",\"type\":\"uint208\"}],\"internalType\":\"structCheckpoints.Checkpoint208\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_forfeitRecentReward\",\"type\":\"bool\"}],\"name\":\"claimStrategyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_forfeitRecentReward\",\"type\":\"bool\"}],\"name\":\"disableRewardForStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"}],\"name\":\"enableRewardForStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eulerEarnVaultModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEulerEarnSavingRate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"uint168\",\"name\":\"\",\"type\":\"uint168\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHooksConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timepoint\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timepoint\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"getStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint120\",\"name\":\"allocated\",\"type\":\"uint120\"},{\"internalType\":\"uint96\",\"name\":\"allocationPoints\",\"type\":\"uint96\"},{\"internalType\":\"AmountCap\",\"name\":\"cap\",\"type\":\"uint16\"},{\"internalType\":\"enumIEulerEarn.StrategyStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIEulerEarn.Strategy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hooksModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"eulerEarnVaultOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialCashAllocationPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"smearingPeriod\",\"type\":\"uint24\"}],\"internalType\":\"structIEulerEarn.InitParams\",\"name\":\"_initParams\",\"type\":\"tuple\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestSmearingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCheckingHarvestCoolDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHarvestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"optInStrategyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"optOutStrategyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_strategies\",\"type\":\"address[]\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"removeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_index1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_index2\",\"type\":\"uint8\"}],\"name\":\"reorderWithdrawalQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hooksTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_hookedFns\",\"type\":\"uint32\"}],\"name\":\"setHooksConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_newFee\",\"type\":\"uint96\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_cap\",\"type\":\"uint16\"}],\"name\":\"setStrategyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"toggleStrategyEmergencyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocationPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssetsAllocatable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssetsDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterestAccrued\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalQueueModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerEulerEarnImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEulerEarnImplementationMetaData.ABI instead.
var EulerEulerEarnImplementationABI = EulerEulerEarnImplementationMetaData.ABI

// EulerEulerEarnImplementation is an auto generated Go binding around an Ethereum contract.
type EulerEulerEarnImplementation struct {
	EulerEulerEarnImplementationCaller     // Read-only binding to the contract
	EulerEulerEarnImplementationTransactor // Write-only binding to the contract
	EulerEulerEarnImplementationFilterer   // Log filterer for contract events
}

// EulerEulerEarnImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEulerEarnImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEulerEarnImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEulerEarnImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEulerEarnImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEulerEarnImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEulerEarnImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEulerEarnImplementationSession struct {
	Contract     *EulerEulerEarnImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EulerEulerEarnImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEulerEarnImplementationCallerSession struct {
	Contract *EulerEulerEarnImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// EulerEulerEarnImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEulerEarnImplementationTransactorSession struct {
	Contract     *EulerEulerEarnImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// EulerEulerEarnImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEulerEarnImplementationRaw struct {
	Contract *EulerEulerEarnImplementation // Generic contract binding to access the raw methods on
}

// EulerEulerEarnImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEulerEarnImplementationCallerRaw struct {
	Contract *EulerEulerEarnImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEulerEarnImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEulerEarnImplementationTransactorRaw struct {
	Contract *EulerEulerEarnImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEulerEarnImplementation creates a new instance of EulerEulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEulerEarnImplementation(address common.Address, backend bind.ContractBackend) (*EulerEulerEarnImplementation, error) {
	contract, err := bindEulerEulerEarnImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementation{EulerEulerEarnImplementationCaller: EulerEulerEarnImplementationCaller{contract: contract}, EulerEulerEarnImplementationTransactor: EulerEulerEarnImplementationTransactor{contract: contract}, EulerEulerEarnImplementationFilterer: EulerEulerEarnImplementationFilterer{contract: contract}}, nil
}

// NewEulerEulerEarnImplementationCaller creates a new read-only instance of EulerEulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEulerEarnImplementationCaller(address common.Address, caller bind.ContractCaller) (*EulerEulerEarnImplementationCaller, error) {
	contract, err := bindEulerEulerEarnImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationCaller{contract: contract}, nil
}

// NewEulerEulerEarnImplementationTransactor creates a new write-only instance of EulerEulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEulerEarnImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEulerEarnImplementationTransactor, error) {
	contract, err := bindEulerEulerEarnImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationTransactor{contract: contract}, nil
}

// NewEulerEulerEarnImplementationFilterer creates a new log filterer instance of EulerEulerEarnImplementation, bound to a specific deployed contract.
func NewEulerEulerEarnImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEulerEarnImplementationFilterer, error) {
	contract, err := bindEulerEulerEarnImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationFilterer{contract: contract}, nil
}

// bindEulerEulerEarnImplementation binds a generic wrapper to an already deployed contract.
func bindEulerEulerEarnImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEulerEarnImplementationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEulerEarnImplementation.Contract.EulerEulerEarnImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EulerEulerEarnImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EulerEulerEarnImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEulerEarnImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) CLOCKMODE() (string, error) {
	return _EulerEulerEarnImplementation.Contract.CLOCKMODE(&_EulerEulerEarnImplementation.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) CLOCKMODE() (string, error) {
	return _EulerEulerEarnImplementation.Contract.CLOCKMODE(&_EulerEulerEarnImplementation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EulerEulerEarnImplementation.Contract.DEFAULTADMINROLE(&_EulerEulerEarnImplementation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EulerEulerEarnImplementation.Contract.DEFAULTADMINROLE(&_EulerEulerEarnImplementation.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) EVC() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.EVC(&_EulerEulerEarnImplementation.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) EVC() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.EVC(&_EulerEulerEarnImplementation.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Allowance(&_EulerEulerEarnImplementation.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Allowance(&_EulerEulerEarnImplementation.CallOpts, _owner, _spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Asset() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Asset(&_EulerEulerEarnImplementation.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Asset() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Asset(&_EulerEulerEarnImplementation.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) BalanceForwarderEnabled(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "balanceForwarderEnabled", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) BalanceForwarderEnabled(_account common.Address) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceForwarderEnabled(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address _account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) BalanceForwarderEnabled(_account common.Address) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceForwarderEnabled(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceOf(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceOf(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceTrackerAddress(&_EulerEulerEarnImplementation.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.BalanceTrackerAddress(&_EulerEulerEarnImplementation.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Checkpoints(opts *bind.CallOpts, _account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "checkpoints", _account, _pos)

	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}

	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Checkpoints(_account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	return _EulerEulerEarnImplementation.Contract.Checkpoints(&_EulerEulerEarnImplementation.CallOpts, _account, _pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address _account, uint32 _pos) view returns((uint48,uint208))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Checkpoints(_account common.Address, _pos uint32) (CheckpointsCheckpoint208, error) {
	return _EulerEulerEarnImplementation.Contract.Checkpoints(&_EulerEulerEarnImplementation.CallOpts, _account, _pos)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Clock() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Clock(&_EulerEulerEarnImplementation.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Clock() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Clock(&_EulerEulerEarnImplementation.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.ConvertToAssets(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.ConvertToAssets(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.ConvertToShares(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.ConvertToShares(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Decimals() (uint8, error) {
	return _EulerEulerEarnImplementation.Contract.Decimals(&_EulerEulerEarnImplementation.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Decimals() (uint8, error) {
	return _EulerEulerEarnImplementation.Contract.Decimals(&_EulerEulerEarnImplementation.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Delegates(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "delegates", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Delegates(_account common.Address) (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Delegates(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address _account) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Delegates(_account common.Address) (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Delegates(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "eip712Domain")

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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EulerEulerEarnImplementation.Contract.Eip712Domain(&_EulerEulerEarnImplementation.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EulerEulerEarnImplementation.Contract.Eip712Domain(&_EulerEulerEarnImplementation.CallOpts)
}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) EulerEarnVaultModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "eulerEarnVaultModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) EulerEarnVaultModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.EulerEarnVaultModule(&_EulerEulerEarnImplementation.CallOpts)
}

// EulerEarnVaultModule is a free data retrieval call binding the contract method 0x6bd757c4.
//
// Solidity: function eulerEarnVaultModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) EulerEarnVaultModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.EulerEarnVaultModule(&_EulerEulerEarnImplementation.CallOpts)
}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) FeeModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "feeModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) FeeModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.FeeModule(&_EulerEulerEarnImplementation.CallOpts)
}

// FeeModule is a free data retrieval call binding the contract method 0xbae41cbf.
//
// Solidity: function feeModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) FeeModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.FeeModule(&_EulerEulerEarnImplementation.CallOpts)
}

// GetEulerEarnSavingRate is a free data retrieval call binding the contract method 0xbeb9c69b.
//
// Solidity: function getEulerEarnSavingRate() view returns(uint40, uint40, uint168)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetEulerEarnSavingRate(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getEulerEarnSavingRate")

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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetEulerEarnSavingRate() (*big.Int, *big.Int, *big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetEulerEarnSavingRate(&_EulerEulerEarnImplementation.CallOpts)
}

// GetEulerEarnSavingRate is a free data retrieval call binding the contract method 0xbeb9c69b.
//
// Solidity: function getEulerEarnSavingRate() view returns(uint40, uint40, uint168)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetEulerEarnSavingRate() (*big.Int, *big.Int, *big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetEulerEarnSavingRate(&_EulerEulerEarnImplementation.CallOpts)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xe6f2aae5.
//
// Solidity: function getHooksConfig() view returns(address, uint32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetHooksConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getHooksConfig")

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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetHooksConfig() (common.Address, uint32, error) {
	return _EulerEulerEarnImplementation.Contract.GetHooksConfig(&_EulerEulerEarnImplementation.CallOpts)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xe6f2aae5.
//
// Solidity: function getHooksConfig() view returns(address, uint32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetHooksConfig() (common.Address, uint32, error) {
	return _EulerEulerEarnImplementation.Contract.GetHooksConfig(&_EulerEulerEarnImplementation.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetPastTotalSupply(opts *bind.CallOpts, _timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getPastTotalSupply", _timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetPastTotalSupply(_timepoint *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetPastTotalSupply(&_EulerEulerEarnImplementation.CallOpts, _timepoint)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetPastTotalSupply(_timepoint *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetPastTotalSupply(&_EulerEulerEarnImplementation.CallOpts, _timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetPastVotes(opts *bind.CallOpts, _account common.Address, _timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getPastVotes", _account, _timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetPastVotes(_account common.Address, _timepoint *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetPastVotes(&_EulerEulerEarnImplementation.CallOpts, _account, _timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address _account, uint256 _timepoint) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetPastVotes(_account common.Address, _timepoint *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetPastVotes(&_EulerEulerEarnImplementation.CallOpts, _account, _timepoint)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleAdmin(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleAdmin(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMember(&_EulerEulerEarnImplementation.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMember(&_EulerEulerEarnImplementation.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMemberCount(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMemberCount(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMembers(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.GetRoleMembers(&_EulerEulerEarnImplementation.CallOpts, role)
}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetStrategy(opts *bind.CallOpts, _strategy common.Address) (IEulerEarnStrategy, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getStrategy", _strategy)

	if err != nil {
		return *new(IEulerEarnStrategy), err
	}

	out0 := *abi.ConvertType(out[0], new(IEulerEarnStrategy)).(*IEulerEarnStrategy)

	return out0, err

}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetStrategy(_strategy common.Address) (IEulerEarnStrategy, error) {
	return _EulerEulerEarnImplementation.Contract.GetStrategy(&_EulerEulerEarnImplementation.CallOpts, _strategy)
}

// GetStrategy is a free data retrieval call binding the contract method 0xf8806a13.
//
// Solidity: function getStrategy(address _strategy) view returns((uint120,uint96,uint16,uint8))
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetStrategy(_strategy common.Address) (IEulerEarnStrategy, error) {
	return _EulerEulerEarnImplementation.Contract.GetStrategy(&_EulerEulerEarnImplementation.CallOpts, _strategy)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) GetVotes(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "getVotes", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GetVotes(_account common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetVotes(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address _account) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) GetVotes(_account common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.GetVotes(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.HasRole(&_EulerEulerEarnImplementation.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.HasRole(&_EulerEulerEarnImplementation.CallOpts, role, account)
}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) HooksModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "hooksModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) HooksModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.HooksModule(&_EulerEulerEarnImplementation.CallOpts)
}

// HooksModule is a free data retrieval call binding the contract method 0x1ee91276.
//
// Solidity: function hooksModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) HooksModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.HooksModule(&_EulerEulerEarnImplementation.CallOpts)
}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) InterestAccrued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "interestAccrued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) InterestAccrued() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.InterestAccrued(&_EulerEulerEarnImplementation.CallOpts)
}

// InterestAccrued is a free data retrieval call binding the contract method 0x20dcc342.
//
// Solidity: function interestAccrued() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) InterestAccrued() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.InterestAccrued(&_EulerEulerEarnImplementation.CallOpts)
}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) InterestSmearingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "interestSmearingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) InterestSmearingPeriod() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.InterestSmearingPeriod(&_EulerEulerEarnImplementation.CallOpts)
}

// InterestSmearingPeriod is a free data retrieval call binding the contract method 0x210da9cd.
//
// Solidity: function interestSmearingPeriod() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) InterestSmearingPeriod() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.InterestSmearingPeriod(&_EulerEulerEarnImplementation.CallOpts)
}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) IsCheckingHarvestCoolDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "isCheckingHarvestCoolDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) IsCheckingHarvestCoolDown() (bool, error) {
	return _EulerEulerEarnImplementation.Contract.IsCheckingHarvestCoolDown(&_EulerEulerEarnImplementation.CallOpts)
}

// IsCheckingHarvestCoolDown is a free data retrieval call binding the contract method 0x8db7e68f.
//
// Solidity: function isCheckingHarvestCoolDown() view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) IsCheckingHarvestCoolDown() (bool, error) {
	return _EulerEulerEarnImplementation.Contract.IsCheckingHarvestCoolDown(&_EulerEulerEarnImplementation.CallOpts)
}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) LastHarvestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "lastHarvestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) LastHarvestTimestamp() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.LastHarvestTimestamp(&_EulerEulerEarnImplementation.CallOpts)
}

// LastHarvestTimestamp is a free data retrieval call binding the contract method 0x2257a738.
//
// Solidity: function lastHarvestTimestamp() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) LastHarvestTimestamp() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.LastHarvestTimestamp(&_EulerEulerEarnImplementation.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) MaxDeposit(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "maxDeposit", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxDeposit(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxDeposit(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) MaxMint(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "maxMint", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxMint(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxMint(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxRedeem(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxRedeem(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxWithdraw(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.MaxWithdraw(&_EulerEulerEarnImplementation.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Name() (string, error) {
	return _EulerEulerEarnImplementation.Contract.Name(&_EulerEulerEarnImplementation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Name() (string, error) {
	return _EulerEulerEarnImplementation.Contract.Name(&_EulerEulerEarnImplementation.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Nonces(owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Nonces(&_EulerEulerEarnImplementation.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.Nonces(&_EulerEulerEarnImplementation.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) NumCheckpoints(opts *bind.CallOpts, _account common.Address) (uint32, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "numCheckpoints", _account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) NumCheckpoints(_account common.Address) (uint32, error) {
	return _EulerEulerEarnImplementation.Contract.NumCheckpoints(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address _account) view returns(uint32)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) NumCheckpoints(_account common.Address) (uint32, error) {
	return _EulerEulerEarnImplementation.Contract.NumCheckpoints(&_EulerEulerEarnImplementation.CallOpts, _account)
}

// PerformanceFeeConfig is a free data retrieval call binding the contract method 0x3eda8287.
//
// Solidity: function performanceFeeConfig() view returns(address, uint96)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) PerformanceFeeConfig(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "performanceFeeConfig")

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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) PerformanceFeeConfig() (common.Address, *big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PerformanceFeeConfig(&_EulerEulerEarnImplementation.CallOpts)
}

// PerformanceFeeConfig is a free data retrieval call binding the contract method 0x3eda8287.
//
// Solidity: function performanceFeeConfig() view returns(address, uint96)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) PerformanceFeeConfig() (common.Address, *big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PerformanceFeeConfig(&_EulerEulerEarnImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Permit2Address() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Permit2Address(&_EulerEulerEarnImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Permit2Address() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.Permit2Address(&_EulerEulerEarnImplementation.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewDeposit(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewDeposit(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewMint(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewMint(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) PreviewRedeem(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "previewRedeem", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewRedeem(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewRedeem(&_EulerEulerEarnImplementation.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewWithdraw(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.PreviewWithdraw(&_EulerEulerEarnImplementation.CallOpts, _assets)
}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) RewardsModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "rewardsModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) RewardsModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.RewardsModule(&_EulerEulerEarnImplementation.CallOpts)
}

// RewardsModule is a free data retrieval call binding the contract method 0x23c8738a.
//
// Solidity: function rewardsModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) RewardsModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.RewardsModule(&_EulerEulerEarnImplementation.CallOpts)
}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) StrategyModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "strategyModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) StrategyModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.StrategyModule(&_EulerEulerEarnImplementation.CallOpts)
}

// StrategyModule is a free data retrieval call binding the contract method 0x8aec2834.
//
// Solidity: function strategyModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) StrategyModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.StrategyModule(&_EulerEulerEarnImplementation.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.SupportsInterface(&_EulerEulerEarnImplementation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EulerEulerEarnImplementation.Contract.SupportsInterface(&_EulerEulerEarnImplementation.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Symbol() (string, error) {
	return _EulerEulerEarnImplementation.Contract.Symbol(&_EulerEulerEarnImplementation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) Symbol() (string, error) {
	return _EulerEulerEarnImplementation.Contract.Symbol(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalAllocated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalAllocated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalAllocated() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAllocated(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAllocated is a free data retrieval call binding the contract method 0x45f7f249.
//
// Solidity: function totalAllocated() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalAllocated() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAllocated(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalAllocationPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalAllocationPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalAllocationPoints() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAllocationPoints(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAllocationPoints is a free data retrieval call binding the contract method 0xe7f3fbde.
//
// Solidity: function totalAllocationPoints() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalAllocationPoints() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAllocationPoints(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalAssets() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssets(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssets(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalAssetsAllocatable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalAssetsAllocatable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalAssetsAllocatable() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssetsAllocatable(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssetsAllocatable is a free data retrieval call binding the contract method 0x23e55160.
//
// Solidity: function totalAssetsAllocatable() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalAssetsAllocatable() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssetsAllocatable(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalAssetsDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalAssetsDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalAssetsDeposited() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssetsDeposited(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalAssetsDeposited is a free data retrieval call binding the contract method 0x6c63c2da.
//
// Solidity: function totalAssetsDeposited() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalAssetsDeposited() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalAssetsDeposited(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TotalSupply() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalSupply(&_EulerEulerEarnImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerEulerEarnImplementation.Contract.TotalSupply(&_EulerEulerEarnImplementation.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) WithdrawalQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "withdrawalQueue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) WithdrawalQueue() ([]common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.WithdrawalQueue(&_EulerEulerEarnImplementation.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address[])
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) WithdrawalQueue() ([]common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.WithdrawalQueue(&_EulerEulerEarnImplementation.CallOpts)
}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCaller) WithdrawalQueueModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEulerEarnImplementation.contract.Call(opts, &out, "withdrawalQueueModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) WithdrawalQueueModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.WithdrawalQueueModule(&_EulerEulerEarnImplementation.CallOpts)
}

// WithdrawalQueueModule is a free data retrieval call binding the contract method 0x39f7444e.
//
// Solidity: function withdrawalQueueModule() view returns(address)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationCallerSession) WithdrawalQueueModule() (common.Address, error) {
	return _EulerEulerEarnImplementation.Contract.WithdrawalQueueModule(&_EulerEulerEarnImplementation.CallOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) AddStrategy(opts *bind.TransactOpts, _strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "addStrategy", _strategy, _allocationPoints)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) AddStrategy(_strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.AddStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _allocationPoints)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address _strategy, uint256 _allocationPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) AddStrategy(_strategy common.Address, _allocationPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.AddStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _allocationPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) AdjustAllocationPoints(opts *bind.TransactOpts, _strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "adjustAllocationPoints", _strategy, _newPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) AdjustAllocationPoints(_strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.AdjustAllocationPoints(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _newPoints)
}

// AdjustAllocationPoints is a paid mutator transaction binding the contract method 0x1671fcad.
//
// Solidity: function adjustAllocationPoints(address _strategy, uint256 _newPoints) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) AdjustAllocationPoints(_strategy common.Address, _newPoints *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.AdjustAllocationPoints(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _newPoints)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Approve(&_EulerEulerEarnImplementation.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Approve(&_EulerEulerEarnImplementation.TransactOpts, _spender, _value)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) ClaimStrategyReward(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "claimStrategyReward", _strategy, _reward, _recipient, _forfeitRecentReward)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) ClaimStrategyReward(_strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ClaimStrategyReward(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward, _recipient, _forfeitRecentReward)
}

// ClaimStrategyReward is a paid mutator transaction binding the contract method 0xed0e5187.
//
// Solidity: function claimStrategyReward(address _strategy, address _reward, address _recipient, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) ClaimStrategyReward(_strategy common.Address, _reward common.Address, _recipient common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ClaimStrategyReward(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward, _recipient, _forfeitRecentReward)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Delegate(opts *bind.TransactOpts, _delegatee common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "delegate", _delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Delegate(_delegatee common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Delegate(&_EulerEulerEarnImplementation.TransactOpts, _delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _delegatee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Delegate(_delegatee common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Delegate(&_EulerEulerEarnImplementation.TransactOpts, _delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) DelegateBySig(opts *bind.TransactOpts, _delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "delegateBySig", _delegatee, _nonce, _expiry, _v, _r, _s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) DelegateBySig(_delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DelegateBySig(&_EulerEulerEarnImplementation.TransactOpts, _delegatee, _nonce, _expiry, _v, _r, _s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address _delegatee, uint256 _nonce, uint256 _expiry, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) DelegateBySig(_delegatee common.Address, _nonce *big.Int, _expiry *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DelegateBySig(&_EulerEulerEarnImplementation.TransactOpts, _delegatee, _nonce, _expiry, _v, _r, _s)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Deposit(&_EulerEulerEarnImplementation.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Deposit(&_EulerEulerEarnImplementation.TransactOpts, _assets, _receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DisableBalanceForwarder(&_EulerEulerEarnImplementation.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DisableBalanceForwarder(&_EulerEulerEarnImplementation.TransactOpts)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) DisableRewardForStrategy(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "disableRewardForStrategy", _strategy, _reward, _forfeitRecentReward)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) DisableRewardForStrategy(_strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DisableRewardForStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward, _forfeitRecentReward)
}

// DisableRewardForStrategy is a paid mutator transaction binding the contract method 0xc3129850.
//
// Solidity: function disableRewardForStrategy(address _strategy, address _reward, bool _forfeitRecentReward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) DisableRewardForStrategy(_strategy common.Address, _reward common.Address, _forfeitRecentReward bool) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.DisableRewardForStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward, _forfeitRecentReward)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EnableBalanceForwarder(&_EulerEulerEarnImplementation.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EnableBalanceForwarder(&_EulerEulerEarnImplementation.TransactOpts)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) EnableRewardForStrategy(opts *bind.TransactOpts, _strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "enableRewardForStrategy", _strategy, _reward)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) EnableRewardForStrategy(_strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EnableRewardForStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward)
}

// EnableRewardForStrategy is a paid mutator transaction binding the contract method 0x2419b45b.
//
// Solidity: function enableRewardForStrategy(address _strategy, address _reward) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) EnableRewardForStrategy(_strategy common.Address, _reward common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.EnableRewardForStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _reward)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) GrantRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "grantRole", _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.GrantRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.GrantRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _account)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Gulp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "gulp")
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Gulp() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Gulp(&_EulerEulerEarnImplementation.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Gulp() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Gulp(&_EulerEulerEarnImplementation.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Harvest() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Harvest(&_EulerEulerEarnImplementation.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Harvest() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Harvest(&_EulerEulerEarnImplementation.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Init(opts *bind.TransactOpts, _initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "init", _initParams)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Init(_initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Init(&_EulerEulerEarnImplementation.TransactOpts, _initParams)
}

// Init is a paid mutator transaction binding the contract method 0xa0515064.
//
// Solidity: function init((address,address,string,string,uint256,uint24) _initParams) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Init(_initParams IEulerEarnInitParams) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Init(&_EulerEulerEarnImplementation.TransactOpts, _initParams)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Mint(&_EulerEulerEarnImplementation.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Mint(&_EulerEulerEarnImplementation.TransactOpts, _shares, _receiver)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) OptInStrategyRewards(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "optInStrategyRewards", _strategy)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) OptInStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.OptInStrategyRewards(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// OptInStrategyRewards is a paid mutator transaction binding the contract method 0x612bddc9.
//
// Solidity: function optInStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) OptInStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.OptInStrategyRewards(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) OptOutStrategyRewards(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "optOutStrategyRewards", _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) OptOutStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.OptOutStrategyRewards(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// OptOutStrategyRewards is a paid mutator transaction binding the contract method 0xefb913eb.
//
// Solidity: function optOutStrategyRewards(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) OptOutStrategyRewards(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.OptOutStrategyRewards(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Rebalance(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "rebalance", _strategies)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Rebalance(_strategies []common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Rebalance(&_EulerEulerEarnImplementation.TransactOpts, _strategies)
}

// Rebalance is a paid mutator transaction binding the contract method 0xbea9db6d.
//
// Solidity: function rebalance(address[] _strategies) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Rebalance(_strategies []common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Rebalance(&_EulerEulerEarnImplementation.TransactOpts, _strategies)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Redeem(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "redeem", _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Redeem(&_EulerEulerEarnImplementation.TransactOpts, _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256 assets)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Redeem(&_EulerEulerEarnImplementation.TransactOpts, _shares, _receiver, _owner)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) RemoveStrategy(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "removeStrategy", _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RemoveStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RemoveStrategy(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) RenounceRole(opts *bind.TransactOpts, _role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "renounceRole", _role, _callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) RenounceRole(_role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RenounceRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _callerConfirmation) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) RenounceRole(_role [32]byte, _callerConfirmation common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RenounceRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _callerConfirmation)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) ReorderWithdrawalQueue(opts *bind.TransactOpts, _index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "reorderWithdrawalQueue", _index1, _index2)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) ReorderWithdrawalQueue(_index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ReorderWithdrawalQueue(&_EulerEulerEarnImplementation.TransactOpts, _index1, _index2)
}

// ReorderWithdrawalQueue is a paid mutator transaction binding the contract method 0x7a7e401b.
//
// Solidity: function reorderWithdrawalQueue(uint8 _index1, uint8 _index2) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) ReorderWithdrawalQueue(_index1 uint8, _index2 uint8) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ReorderWithdrawalQueue(&_EulerEulerEarnImplementation.TransactOpts, _index1, _index2)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) RevokeRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "revokeRole", _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RevokeRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.RevokeRole(&_EulerEulerEarnImplementation.TransactOpts, _role, _account)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) SetFeeRecipient(opts *bind.TransactOpts, _newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "setFeeRecipient", _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetFeeRecipient(&_EulerEulerEarnImplementation.TransactOpts, _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetFeeRecipient(&_EulerEulerEarnImplementation.TransactOpts, _newFeeRecipient)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) SetHooksConfig(opts *bind.TransactOpts, _hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "setHooksConfig", _hooksTarget, _hookedFns)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) SetHooksConfig(_hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetHooksConfig(&_EulerEulerEarnImplementation.TransactOpts, _hooksTarget, _hookedFns)
}

// SetHooksConfig is a paid mutator transaction binding the contract method 0x24506add.
//
// Solidity: function setHooksConfig(address _hooksTarget, uint32 _hookedFns) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) SetHooksConfig(_hooksTarget common.Address, _hookedFns uint32) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetHooksConfig(&_EulerEulerEarnImplementation.TransactOpts, _hooksTarget, _hookedFns)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) SetPerformanceFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "setPerformanceFee", _newFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) SetPerformanceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetPerformanceFee(&_EulerEulerEarnImplementation.TransactOpts, _newFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x87451de6.
//
// Solidity: function setPerformanceFee(uint96 _newFee) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) SetPerformanceFee(_newFee *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetPerformanceFee(&_EulerEulerEarnImplementation.TransactOpts, _newFee)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) SetStrategyCap(opts *bind.TransactOpts, _strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "setStrategyCap", _strategy, _cap)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) SetStrategyCap(_strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetStrategyCap(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _cap)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0xd51988d0.
//
// Solidity: function setStrategyCap(address _strategy, uint16 _cap) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) SetStrategyCap(_strategy common.Address, _cap uint16) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.SetStrategyCap(&_EulerEulerEarnImplementation.TransactOpts, _strategy, _cap)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Skim(opts *bind.TransactOpts, _token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "skim", _token, _recipient)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Skim(_token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Skim(&_EulerEulerEarnImplementation.TransactOpts, _token, _recipient)
}

// Skim is a paid mutator transaction binding the contract method 0x712b772f.
//
// Solidity: function skim(address _token, address _recipient) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Skim(_token common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Skim(&_EulerEulerEarnImplementation.TransactOpts, _token, _recipient)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) ToggleStrategyEmergencyStatus(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "toggleStrategyEmergencyStatus", _strategy)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) ToggleStrategyEmergencyStatus(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ToggleStrategyEmergencyStatus(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// ToggleStrategyEmergencyStatus is a paid mutator transaction binding the contract method 0x55c00f24.
//
// Solidity: function toggleStrategyEmergencyStatus(address _strategy) returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) ToggleStrategyEmergencyStatus(_strategy common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.ToggleStrategyEmergencyStatus(&_EulerEulerEarnImplementation.TransactOpts, _strategy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Transfer(&_EulerEulerEarnImplementation.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Transfer(&_EulerEulerEarnImplementation.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.TransferFrom(&_EulerEulerEarnImplementation.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.TransferFrom(&_EulerEulerEarnImplementation.TransactOpts, _from, _to, _value)
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) UpdateInterestAccrued(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "updateInterestAccrued")
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) UpdateInterestAccrued() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.UpdateInterestAccrued(&_EulerEulerEarnImplementation.TransactOpts)
}

// UpdateInterestAccrued is a paid mutator transaction binding the contract method 0xf0618791.
//
// Solidity: function updateInterestAccrued() returns()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) UpdateInterestAccrued() (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.UpdateInterestAccrued(&_EulerEulerEarnImplementation.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactor) Withdraw(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.contract.Transact(opts, "withdraw", _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Withdraw(&_EulerEulerEarnImplementation.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationTransactorSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _EulerEulerEarnImplementation.Contract.Withdraw(&_EulerEulerEarnImplementation.TransactOpts, _assets, _receiver, _owner)
}

// EulerEulerEarnImplementationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationApprovalIterator struct {
	Event *EulerEulerEarnImplementationApproval // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationApproval)
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
		it.Event = new(EulerEulerEarnImplementationApproval)
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
func (it *EulerEulerEarnImplementationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationApproval represents a Approval event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerEulerEarnImplementationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationApprovalIterator{contract: _EulerEulerEarnImplementation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationApproval)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseApproval(log types.Log) (*EulerEulerEarnImplementationApproval, error) {
	event := new(EulerEulerEarnImplementationApproval)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDelegateChangedIterator struct {
	Event *EulerEulerEarnImplementationDelegateChanged // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationDelegateChanged)
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
		it.Event = new(EulerEulerEarnImplementationDelegateChanged)
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
func (it *EulerEulerEarnImplementationDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationDelegateChanged represents a DelegateChanged event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*EulerEulerEarnImplementationDelegateChangedIterator, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationDelegateChangedIterator{contract: _EulerEulerEarnImplementation.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationDelegateChanged)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseDelegateChanged(log types.Log) (*EulerEulerEarnImplementationDelegateChanged, error) {
	event := new(EulerEulerEarnImplementationDelegateChanged)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDelegateVotesChangedIterator struct {
	Event *EulerEulerEarnImplementationDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationDelegateVotesChanged)
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
		it.Event = new(EulerEulerEarnImplementationDelegateVotesChanged)
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
func (it *EulerEulerEarnImplementationDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationDelegateVotesChanged represents a DelegateVotesChanged event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*EulerEulerEarnImplementationDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationDelegateVotesChangedIterator{contract: _EulerEulerEarnImplementation.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationDelegateVotesChanged)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseDelegateVotesChanged(log types.Log) (*EulerEulerEarnImplementationDelegateVotesChanged, error) {
	event := new(EulerEulerEarnImplementationDelegateVotesChanged)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDepositIterator struct {
	Event *EulerEulerEarnImplementationDeposit // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationDeposit)
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
		it.Event = new(EulerEulerEarnImplementationDeposit)
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
func (it *EulerEulerEarnImplementationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationDeposit represents a Deposit event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerEulerEarnImplementationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationDepositIterator{contract: _EulerEulerEarnImplementation.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationDeposit)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseDeposit(log types.Log) (*EulerEulerEarnImplementationDeposit, error) {
	event := new(EulerEulerEarnImplementationDeposit)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationEIP712DomainChangedIterator struct {
	Event *EulerEulerEarnImplementationEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationEIP712DomainChanged)
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
		it.Event = new(EulerEulerEarnImplementationEIP712DomainChanged)
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
func (it *EulerEulerEarnImplementationEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationEIP712DomainChanged represents a EIP712DomainChanged event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EulerEulerEarnImplementationEIP712DomainChangedIterator, error) {

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationEIP712DomainChangedIterator{contract: _EulerEulerEarnImplementation.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationEIP712DomainChanged)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseEIP712DomainChanged(log types.Log) (*EulerEulerEarnImplementationEIP712DomainChanged, error) {
	event := new(EulerEulerEarnImplementationEIP712DomainChanged)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationInitializedIterator struct {
	Event *EulerEulerEarnImplementationInitialized // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationInitialized)
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
		it.Event = new(EulerEulerEarnImplementationInitialized)
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
func (it *EulerEulerEarnImplementationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationInitialized represents a Initialized event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterInitialized(opts *bind.FilterOpts) (*EulerEulerEarnImplementationInitializedIterator, error) {

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationInitializedIterator{contract: _EulerEulerEarnImplementation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationInitialized) (event.Subscription, error) {

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationInitialized)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseInitialized(log types.Log) (*EulerEulerEarnImplementationInitialized, error) {
	event := new(EulerEulerEarnImplementationInitialized)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleAdminChangedIterator struct {
	Event *EulerEulerEarnImplementationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationRoleAdminChanged)
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
		it.Event = new(EulerEulerEarnImplementationRoleAdminChanged)
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
func (it *EulerEulerEarnImplementationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationRoleAdminChanged represents a RoleAdminChanged event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EulerEulerEarnImplementationRoleAdminChangedIterator, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationRoleAdminChangedIterator{contract: _EulerEulerEarnImplementation.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationRoleAdminChanged)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseRoleAdminChanged(log types.Log) (*EulerEulerEarnImplementationRoleAdminChanged, error) {
	event := new(EulerEulerEarnImplementationRoleAdminChanged)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleGrantedIterator struct {
	Event *EulerEulerEarnImplementationRoleGranted // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationRoleGranted)
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
		it.Event = new(EulerEulerEarnImplementationRoleGranted)
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
func (it *EulerEulerEarnImplementationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationRoleGranted represents a RoleGranted event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EulerEulerEarnImplementationRoleGrantedIterator, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationRoleGrantedIterator{contract: _EulerEulerEarnImplementation.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationRoleGranted)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseRoleGranted(log types.Log) (*EulerEulerEarnImplementationRoleGranted, error) {
	event := new(EulerEulerEarnImplementationRoleGranted)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleRevokedIterator struct {
	Event *EulerEulerEarnImplementationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationRoleRevoked)
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
		it.Event = new(EulerEulerEarnImplementationRoleRevoked)
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
func (it *EulerEulerEarnImplementationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationRoleRevoked represents a RoleRevoked event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EulerEulerEarnImplementationRoleRevokedIterator, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationRoleRevokedIterator{contract: _EulerEulerEarnImplementation.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationRoleRevoked)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseRoleRevoked(log types.Log) (*EulerEulerEarnImplementationRoleRevoked, error) {
	event := new(EulerEulerEarnImplementationRoleRevoked)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationTransferIterator struct {
	Event *EulerEulerEarnImplementationTransfer // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationTransfer)
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
		it.Event = new(EulerEulerEarnImplementationTransfer)
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
func (it *EulerEulerEarnImplementationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationTransfer represents a Transfer event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEulerEarnImplementationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationTransferIterator{contract: _EulerEulerEarnImplementation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationTransfer)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseTransfer(log types.Log) (*EulerEulerEarnImplementationTransfer, error) {
	event := new(EulerEulerEarnImplementationTransfer)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEulerEarnImplementationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationWithdrawIterator struct {
	Event *EulerEulerEarnImplementationWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerEulerEarnImplementationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEulerEarnImplementationWithdraw)
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
		it.Event = new(EulerEulerEarnImplementationWithdraw)
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
func (it *EulerEulerEarnImplementationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEulerEarnImplementationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEulerEarnImplementationWithdraw represents a Withdraw event raised by the EulerEulerEarnImplementation contract.
type EulerEulerEarnImplementationWithdraw struct {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerEulerEarnImplementationWithdrawIterator, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEulerEarnImplementationWithdrawIterator{contract: _EulerEulerEarnImplementation.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerEulerEarnImplementationWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEulerEarnImplementation.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEulerEarnImplementationWithdraw)
				if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EulerEulerEarnImplementation *EulerEulerEarnImplementationFilterer) ParseWithdraw(log types.Log) (*EulerEulerEarnImplementationWithdraw, error) {
	event := new(EulerEulerEarnImplementationWithdraw)
	if err := _EulerEulerEarnImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
