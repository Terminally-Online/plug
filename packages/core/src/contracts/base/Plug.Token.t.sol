// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { SuperchainERC20 } from "op/SuperchainERC20.sol";
import { Test } from "../abstracts/test/Plug.Test.sol";
import { PlugToken } from "./Plug.Token.sol";
import { PredeployAddresses } from "op/libraries/PredeployAddresses.sol";
import { Ownable } from "solady/auth/Ownable.sol";

contract PlugTokenTest is Test {
    PlugToken internal token;
    address internal owner;
    uint32 internal unlockTime;
    uint256 internal constant TOTAL_SUPPLY = 1_000_000 ether;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event CrosschainMint(
        address indexed to, uint256 amount, address indexed sender
    );

    function setUp() public virtual {
        owner = _randomNonZeroAddress();
        unlockTime = uint32(block.timestamp + 1 days);
        token = new PlugToken();
        token.initialize(unlockTime, owner, TOTAL_SUPPLY);
    }

    function test_name() public {
        assertEq(token.name(), "Plug");
    }

    function test_symbol() public {
        assertEq(token.symbol(), "PLUG");
    }

    function test_transferUnlock() public {
        assertEq(token.transferUnlock(), unlockTime);
    }

    function test_bridgeUnlock() public {
        assertEq(token.bridgeUnlock(), unlockTime);
    }

    function test_setTransferUnlock() public {
        uint32 newUnlockTime = uint32(block.timestamp + 2 days);
        vm.prank(owner);
        token.setTransferUnlock(newUnlockTime);
        assertEq(token.transferUnlock(), newUnlockTime);
    }

    function test_setBridgeUnlock() public {
        uint32 newUnlockTime = uint32(block.timestamp + 2 days);
        vm.prank(owner);
        token.setBridgeUnlock(newUnlockTime);
        assertEq(token.bridgeUnlock(), newUnlockTime);
    }

    function testRevert_setTransferUnlock_NotOwner() public {
        uint32 newUnlockTime = uint32(block.timestamp + 2 days);
        vm.expectRevert(Ownable.Unauthorized.selector);
        token.setTransferUnlock(newUnlockTime);
    }

    function testRevert_setBridgeUnlock_NotOwner() public {
        uint32 newUnlockTime = uint32(block.timestamp + 2 days);
        vm.expectRevert(Ownable.Unauthorized.selector);
        token.setBridgeUnlock(newUnlockTime);
    }

    function test_transfer_AfterUnlock() public {
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.transfer(address(this), amount);
        vm.warp(unlockTime + 1);
        vm.expectEmit(true, true, false, true);
        emit Transfer(address(this), recipient, amount);
        token.transfer(recipient, amount);
        assertEq(token.balanceOf(recipient), amount);
    }

    function testRevert_transfer_BeforeUnlock() public {
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.transfer(address(this), amount);
        vm.expectRevert("PlugToken:transfer-locked");
        token.transfer(recipient, amount);
    }

    function test_initialSupply() public {
        assertEq(token.totalSupply(), TOTAL_SUPPLY);
        assertEq(token.balanceOf(owner), TOTAL_SUPPLY);
    }

    function testRevert_initialize_Twice() public {
        vm.expectRevert();
        token.initialize(unlockTime, owner, TOTAL_SUPPLY);
    }

    function test_constructor() public {
        PlugToken newToken = new PlugToken();
        assertEq(newToken.owner(), address(1));
    }

    function test_allowedSender_BeforeUnlock() public {
        address sender = _randomNonZeroAddress();
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.setSenderAllowed(sender, true);
        vm.prank(owner);
        token.transfer(sender, amount);
        vm.prank(sender);
        token.transfer(recipient, amount);
        assertEq(token.balanceOf(recipient), amount);
    }

    function test_allowedSender_Toggle() public {
        address sender = _randomNonZeroAddress();
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.setSenderAllowed(sender, true);
        vm.prank(owner);
        token.transfer(sender, amount);
        vm.prank(sender);
        token.transfer(recipient, amount / 2);
        assertEq(token.balanceOf(recipient), amount / 2);
        vm.prank(owner);
        token.setSenderAllowed(sender, false);
        vm.expectRevert("PlugToken:transfer-locked");
        vm.prank(sender);
        token.transfer(recipient, amount / 2);
    }

    function test_superchainERC20_Compatibility() public {
        SuperchainERC20 superchainToken = SuperchainERC20(address(token));
        assertEq(address(superchainToken), address(token));
    }

    function test_superchainERC20_TransferAfterUnlock() public {
        SuperchainERC20 superchainToken = SuperchainERC20(address(token));
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.transfer(address(this), amount);
        vm.warp(unlockTime + 1);
        vm.expectEmit(true, true, false, true);
        emit Transfer(address(this), recipient, amount);
        superchainToken.transfer(recipient, amount);
        assertEq(superchainToken.balanceOf(recipient), amount);
    }

    function testRevert_superchainERC20_TransferBeforeUnlock() public {
        SuperchainERC20 superchainToken = SuperchainERC20(address(token));
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;
        vm.prank(owner);
        token.transfer(address(this), amount);
        vm.expectRevert("PlugToken:transfer-locked");
        superchainToken.transfer(recipient, amount);
    }

    function test_crosschainMint_BeforeUnlock() public {
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        vm.prank(PredeployAddresses.SUPERCHAIN_TOKEN_BRIDGE);
        vm.expectRevert("PlugToken:bridge-locked");
        token.crosschainMint(recipient, amount);
    }

    function testRevert_crosschainMint_NotBridge() public {
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        vm.expectRevert("Unauthorized");
        token.crosschainMint(recipient, amount);
    }

    function test_crosschainBurn_BeforeUnlock() public {
        address sender = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        // Try to burn tokens before unlock
        vm.prank(PredeployAddresses.SUPERCHAIN_TOKEN_BRIDGE);
        vm.expectRevert("PlugToken:bridge-locked");
        token.crosschainBurn(sender, amount);
    }

    function testRevert_crosschainBurn_NotBridge() public {
        address sender = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        vm.expectRevert("Unauthorized");
        token.crosschainBurn(sender, amount);
    }

    function testRevert_crosschainMint_BridgeLocked() public {
        address recipient = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        vm.prank(owner);
        token.setBridgeUnlock(uint32(block.timestamp + 1 days));

        vm.prank(PredeployAddresses.SUPERCHAIN_TOKEN_BRIDGE);
        vm.expectRevert("PlugToken:bridge-locked");
        token.crosschainMint(recipient, amount);
    }

    function testRevert_crosschainBurn_BridgeLocked() public {
        address sender = _randomNonZeroAddress();
        uint256 amount = 100 ether;

        // First set bridge unlock to future time
        vm.prank(owner);
        token.setBridgeUnlock(uint32(block.timestamp + 1 days));

        // Try to burn tokens
        vm.prank(PredeployAddresses.SUPERCHAIN_TOKEN_BRIDGE);
        vm.expectRevert("PlugToken:bridge-locked");
        token.crosschainBurn(sender, amount);
    }
}
