// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { Test } from "forge-std/Test.sol";
import { PlugRewards } from "./Plug.Rewards.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";
import { PlugToken } from "./Plug.Token.sol";
import { MerkleProofLib } from "solady/utils/MerkleProofLib.sol";
import { Ownable } from "solady/auth/Ownable.sol";
import { PredeployAddresses } from "../mocks/Plug.Rewards.Test.Mocks.sol";

contract PlugRewardsTest is Test {
    PlugRewards internal rewards;
    PlugToken internal rewardToken;
    address internal owner;
    address internal user1;
    address internal user2;

    // Data for merkle tree
    bytes32 internal merkleRoot;
    bytes32[] internal proofUser1;
    bytes32[] internal proofUser2;
    uint256 internal amountUser1 = 100 ether;
    uint256 internal amountUser2 = 50 ether;
    uint256 internal totalAmount = 150 ether;

    function setUp() public {
        owner = makeAddr("owner");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        
        // Deploy the actual PlugToken with transfer unlocked (timestamp = 0)
        rewardToken = new PlugToken();
        rewardToken.initialize(0, owner, 1000 ether);
        
        rewards = new PlugRewards(owner, address(rewardToken));
        
        // Create merkle tree with two users
        // These values would typically be generated off-chain
        bytes32 leafUser1 = keccak256(abi.encodePacked(user1, amountUser1));
        bytes32 leafUser2 = keccak256(abi.encodePacked(user2, amountUser2));
        
        // Simple merkle tree with just two leaves for testing
        bytes32 node = keccak256(abi.encodePacked(leafUser1, leafUser2));
        merkleRoot = node;
        
        // Create proofs
        proofUser1 = new bytes32[](1);
        proofUser1[0] = leafUser2;
        
        proofUser2 = new bytes32[](1);
        proofUser2[0] = leafUser1;
        
        // Approve rewards contract to spend owner's tokens
        vm.prank(owner);
        rewardToken.approve(address(rewards), type(uint256).max);
    }
    
    function test_constructor() public view {
        assertEq(rewards.rewardToken(), address(rewardToken));
        assertEq(rewards.currentPeriod(), 0);
        assertEq(rewards.owner(), owner);
    }
    
    function test_createRewardPeriod() public {
        vm.prank(owner);
        vm.expectEmit(true, false, false, true);
        emit PlugLib.NewRewardPeriod(1, merkleRoot, totalAmount);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        assertEq(rewards.currentPeriod(), 1);
        assertEq(rewards.periodMerkleRoots(1), merkleRoot);
        assertEq(rewards.periodTotalAmounts(1), totalAmount);
    }
    
    function testRevert_createRewardPeriod_NotOwner() public {
        vm.prank(user1);
        vm.expectRevert(Ownable.Unauthorized.selector);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
    }
    
    function testRevert_createRewardPeriod_ZeroAmount() public {
        vm.prank(owner);
        vm.expectRevert(PlugLib.ZeroAmount.selector);
        rewards.createRewardPeriod(merkleRoot, 0);
    }
    
    function test_fundRewards() public {
        uint256 fundAmount = 200 ether;
        
        vm.prank(owner);
        rewards.fundRewards(fundAmount);
        
        assertEq(rewardToken.balanceOf(address(rewards)), fundAmount);
        assertEq(rewards.getRewardBalance(), fundAmount);
    }
    
    function testRevert_fundRewards_NotOwner() public {
        uint256 fundAmount = 200 ether;
        
        vm.prank(user1);
        vm.expectRevert(Ownable.Unauthorized.selector);
        rewards.fundRewards(fundAmount);
    }
    
    function testRevert_fundRewards_ZeroAmount() public {
        vm.prank(owner);
        vm.expectRevert(PlugLib.ZeroAmount.selector);
        rewards.fundRewards(0);
    }
    
    function test_claimReward() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Initial balances
        assertEq(rewardToken.balanceOf(user1), 0);
        
        // Claim rewards for user1
        vm.prank(user1);
        vm.expectEmit(true, true, false, true);
        emit PlugLib.RewardClaimed(1, user1, amountUser1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // Verify claim was processed
        assertEq(rewardToken.balanceOf(user1), amountUser1);
        assertTrue(rewards.rewardClaimed(1, user1));
        
        // Check remaining balance
        assertEq(rewards.getRewardBalance(), totalAmount - amountUser1);
    }
    
    function testRevert_claimReward_AlreadyClaimed() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // First claim
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // Try to claim again
        vm.prank(user1);
        vm.expectRevert(PlugLib.RewardsAlreadyClaimed.selector);
        rewards.claimReward(1, amountUser1, proofUser1);
    }
    
    function testRevert_claimReward_PeriodNotInitialized() public {
        vm.prank(user1);
        vm.expectRevert(PlugLib.PeriodNotInitialized.selector);
        rewards.claimReward(1, amountUser1, proofUser1);
    }
    
    function testRevert_claimReward_InvalidProof() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Try to claim with wrong amount
        vm.prank(user1);
        vm.expectRevert(PlugLib.InvalidMerkleProof.selector);
        rewards.claimReward(1, amountUser2, proofUser1);
        
        // Try to claim with wrong proof
        vm.prank(user1);
        vm.expectRevert(PlugLib.InvalidMerkleProof.selector);
        rewards.claimReward(1, amountUser1, proofUser2);
    }
    
    function testRevert_claimReward_InsufficientBalance() public {
        // Setup period but with insufficient funds
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        // Fund with less than required
        vm.prank(owner);
        rewards.fundRewards(amountUser1 / 2);
        
        // Try to claim
        vm.prank(user1);
        vm.expectRevert(PlugLib.InsufficientRewardBalance.selector);
        rewards.claimReward(1, amountUser1, proofUser1);
    }
    
    function test_hasValidClaim() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        // Check valid claim
        assertTrue(rewards.hasValidClaim(1, user1, amountUser1, proofUser1));
        assertTrue(rewards.hasValidClaim(1, user2, amountUser2, proofUser2));
        
        // Check invalid amount
        assertFalse(rewards.hasValidClaim(1, user1, amountUser2, proofUser1));
        assertFalse(rewards.hasValidClaim(1, user2, amountUser1, proofUser2));
        
        // Check invalid proof
        assertFalse(rewards.hasValidClaim(1, user1, amountUser1, proofUser2));
        assertFalse(rewards.hasValidClaim(1, user2, amountUser2, proofUser1));
        
        // Claim rewards and verify claim is no longer valid after claiming
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        assertFalse(rewards.hasValidClaim(1, user1, amountUser1, proofUser1));
        assertTrue(rewards.hasValidClaim(1, user2, amountUser2, proofUser2));
    }
    
    function test_hasValidClaim_PeriodNotInitialized() public view {
        // Non-existent period should return false
        assertFalse(rewards.hasValidClaim(1, user1, amountUser1, proofUser1));
    }
    
    function test_getRewardBalance() public {
        assertEq(rewards.getRewardBalance(), 0);
        
        uint256 fundAmount = 200 ether;
        vm.prank(owner);
        rewards.fundRewards(fundAmount);
        
        assertEq(rewards.getRewardBalance(), fundAmount);
        
        // Create period and claim some rewards
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        assertEq(rewards.getRewardBalance(), fundAmount - amountUser1);
    }
    
    function test_multipleRewardPeriods() public {
        // Create first period
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Claim from first period
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // Create second period with same merkle root
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // User1 already claimed from period 1, but should be able to claim from period 2
        vm.prank(user1);
        rewards.claimReward(2, amountUser1, proofUser1);
        
        // Check balances
        assertEq(rewardToken.balanceOf(user1), amountUser1 * 2);
        
        // Check claim status
        assertTrue(rewards.rewardClaimed(1, user1));
        assertTrue(rewards.rewardClaimed(2, user1));
        assertFalse(rewards.rewardClaimed(1, user2));
        assertFalse(rewards.rewardClaimed(2, user2));
    }
    
    // Additional tests for token locking mechanism
    
    function test_tokenTransferWhenLocked() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Set token to locked state
        uint32 futureTime = uint32(block.timestamp + 1 days);
        vm.prank(owner);
        rewardToken.setTransferUnlock(futureTime);
        
        // Claim rewards works even with locked token because rewards contract is allowed
        vm.prank(owner);
        rewardToken.setSenderAllowed(address(rewards), true);
        
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // User can't transfer their tokens yet because they're locked
        vm.prank(user1);
        vm.expectRevert("PlugToken:transfer-locked");
        rewardToken.transfer(user2, amountUser1);
        
        // Forward time to unlock
        vm.warp(futureTime + 1);
        
        // Now user can transfer
        vm.prank(user1);
        rewardToken.transfer(user2, amountUser1);
        
        assertEq(rewardToken.balanceOf(user1), 0);
        assertEq(rewardToken.balanceOf(user2), amountUser1);
    }
    
    function test_rewardsContractAllowedSender() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Lock token transfers
        uint32 futureTime = uint32(block.timestamp + 1 days);
        vm.prank(owner);
        rewardToken.setTransferUnlock(futureTime);
        
        // Don't allow the rewards contract
        vm.prank(owner);
        rewardToken.setSenderAllowed(address(rewards), false);
        
        // Claim would fail because rewards contract can't transfer tokens
        vm.prank(user1);
        vm.expectRevert();
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // Now allow the rewards contract
        vm.prank(owner);
        rewardToken.setSenderAllowed(address(rewards), true);
        
        // Claim should work now
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        assertEq(rewardToken.balanceOf(user1), amountUser1);
    }
    
    function test_ownerCanBypassLock() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Lock token transfers
        uint32 futureTime = uint32(block.timestamp + 1 days);
        vm.prank(owner);
        rewardToken.setTransferUnlock(futureTime);
        
        // Disallow rewards contract
        vm.prank(owner);
        rewardToken.setSenderAllowed(address(rewards), false);
        
        // Owner should still be able to transfer tokens even when locked
        vm.prank(owner);
        rewardToken.transfer(user2, 100 ether);
        
        assertEq(rewardToken.balanceOf(user2), 100 ether);
    }
    
    function test_unlockChangesGlobalLockState() public {
        // Setup with lock
        uint32 futureTime = uint32(block.timestamp + 1 days);
        vm.prank(owner);
        rewardToken.setTransferUnlock(futureTime);
        
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Allow rewards contract
        vm.prank(owner);
        rewardToken.setSenderAllowed(address(rewards), true);
        
        // User can claim
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // But can't transfer
        vm.prank(user1);
        vm.expectRevert("PlugToken:transfer-locked");
        rewardToken.transfer(user2, amountUser1);
        
        // Change global lock to current time (unlocked)
        vm.prank(owner);
        rewardToken.setTransferUnlock(uint32(block.timestamp));
        
        // Now user can transfer
        vm.prank(user1);
        rewardToken.transfer(user2, amountUser1);
        
        assertEq(rewardToken.balanceOf(user2), amountUser1);
    }
    
    function test_separateBridgeAndTransferLocks() public {
        // Setup
        vm.prank(owner);
        rewards.createRewardPeriod(merkleRoot, totalAmount);
        
        vm.prank(owner);
        rewards.fundRewards(totalAmount);
        
        // Set different unlock times for transfer and bridge
        uint32 transferUnlock = uint32(block.timestamp + 1 days);
        uint32 bridgeUnlock = uint32(block.timestamp + 2 days);
        
        vm.startPrank(owner);
        rewardToken.setTransferUnlock(transferUnlock);
        rewardToken.setBridgeUnlock(bridgeUnlock);
        rewardToken.setSenderAllowed(address(rewards), true);
        vm.stopPrank();
        
        // User can still claim rewards because rewards contract is allowed
        vm.prank(user1);
        rewards.claimReward(1, amountUser1, proofUser1);
        
        // User now has tokens but can't transfer them yet
        assertEq(rewardToken.balanceOf(user1), amountUser1);
        vm.prank(user1);
        vm.expectRevert("PlugToken:transfer-locked");
        rewardToken.transfer(user2, amountUser1);
        
        // Advance time to after transfer unlock but before bridge unlock
        vm.warp(transferUnlock + 1);
        
        // Now user can transfer tokens
        vm.prank(user1);
        rewardToken.transfer(user2, amountUser1);
        assertEq(rewardToken.balanceOf(user2), amountUser1);
        
        // Verify that bridge and transfer locks can be independently controlled
        assertEq(rewardToken.transferUnlock(), transferUnlock);
        assertEq(rewardToken.bridgeUnlock(), bridgeUnlock);
    }
}