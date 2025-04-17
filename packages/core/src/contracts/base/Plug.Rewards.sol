// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.26;

import { Ownable } from "solady/auth/Ownable.sol";
import { MerkleProofLib } from "solady/utils/MerkleProofLib.sol";
import { SafeTransferLib } from "solady/utils/SafeTransferLib.sol";
import { FixedPointMathLib } from "solady/utils/FixedPointMathLib.sol";

import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Rewards
 * @notice This contract implements a reward distribution system using merkle trees
 *         to efficiently distribute rewards to users on a rolling period basis.
 * @dev Each reward period has its own merkle root, allowing for on-chain verification
 *      of reward claims. Rewards are distributed in the project's ERC20 token.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugRewards is Ownable {
    using FixedPointMathLib for uint256;

    /// @dev The ERC20 token used for rewards
    address public immutable rewardToken;

    /// @dev Tracks the merkle root for each reward period
    mapping(uint256 period => bytes32 merkleRoot) public periodMerkleRoots;

    /// @dev Tracks the total amount allocated for each reward period
    mapping(uint256 period => uint256 amount) public periodTotalAmounts;

    /// @dev Tracks whether an address has claimed rewards for a specific period
    mapping(uint256 period => mapping(address user => bool claimed)) public
        rewardClaimed;

    uint256 public currentPeriod;

    /**
     * @notice Initializes the rewards contract with an owner and reward token
     * @param $owner The owner of the contract
     * @param $token The ERC20 token used for rewards
     */
    constructor(address $owner, address $token) {
        _initializeOwner($owner);

        rewardToken = $token;
    }

    /**
     * @notice Creates a new reward period with a specified merkle root
     * @dev The merkle tree should be constructed off-chain with leaf nodes as
     *      keccak256(abi.encodePacked(address user, uint256 amount))
     * @param $merkleRoot The merkle root of the reward distribution tree
     * @param $totalAmount Total amount of rewards allocated for this period
     */
    function createRewardPeriod(
        bytes32 $merkleRoot,
        uint256 $totalAmount
    )
        external
        onlyOwner
    {
        if ($totalAmount == 0) revert PlugLib.ZeroAmount();

        uint256 period = ++currentPeriod;

        periodMerkleRoots[period] = $merkleRoot;
        periodTotalAmounts[period] = $totalAmount;

        emit PlugLib.NewRewardPeriod(period, $merkleRoot, $totalAmount);
    }

    /**
     * @notice Allows contract owner to fund the contract with reward tokens
     * @dev Transfers tokens from the owner to this contract
     * @param $amount Amount of tokens to transfer
     */
    function fundRewards(uint256 $amount) external onlyOwner {
        if ($amount == 0) revert PlugLib.ZeroAmount();
        SafeTransferLib.safeTransferFrom(
            rewardToken, msg.sender, address(this), $amount
        );
    }

    /**
     * @notice Allows users to claim rewards for a specific period
     * @dev Verifies the merkle proof and transfers the rewards to the user
     * @param $period The reward period to claim for
     * @param $amount The amount to be claimed
     * @param $merkleProof The merkle proof to verify the claim
     */
    function claimReward(
        uint256 $period,
        uint256 $amount,
        bytes32[] calldata $merkleProof
    )
        external
    {
        if (rewardClaimed[$period][msg.sender]) {
            revert PlugLib.RewardsAlreadyClaimed();
        }

        bytes32 merkleRoot = periodMerkleRoots[$period];
        if (merkleRoot == bytes32(0)) revert PlugLib.PeriodNotInitialized();

        bytes32 leaf = keccak256(abi.encodePacked(msg.sender, $amount));
        if (!MerkleProofLib.verify($merkleProof, merkleRoot, leaf)) {
            revert PlugLib.InvalidMerkleProof();
        }

        rewardClaimed[$period][msg.sender] = true;

        if (SafeTransferLib.balanceOf(rewardToken, address(this)) < $amount) {
            revert PlugLib.InsufficientRewardBalance();
        }
        SafeTransferLib.safeTransfer(rewardToken, msg.sender, $amount);

        emit PlugLib.RewardClaimed($period, msg.sender, $amount);
    }

    /**
     * @notice Checks if a user has a valid claim for a given period
     * @dev This is a view function to verify off-chain if a claim is valid
     * @param $period The reward period to check
     * @param $user The user address to check
     * @param $amount The amount to verify
     * @param $merkleProof The merkle proof to verify
     * @return True if the claim is valid, false otherwise
     */
    function hasValidClaim(
        uint256 $period,
        address $user,
        uint256 $amount,
        bytes32[] calldata $merkleProof
    )
        external
        view
        returns (bool)
    {
        bytes32 merkleRoot = periodMerkleRoots[$period];
        if (merkleRoot == bytes32(0)) return false;

        if (rewardClaimed[$period][$user]) return false;

        bytes32 leaf = keccak256(abi.encodePacked($user, $amount));
        return MerkleProofLib.verify($merkleProof, merkleRoot, leaf);
    }

    /**
     * @notice Get the current balance of reward tokens available in the contract
     * @return The balance of the reward token
     */
    function getRewardBalance() external view returns (uint256) {
        return SafeTransferLib.balanceOf(rewardToken, address(this));
    }
}
