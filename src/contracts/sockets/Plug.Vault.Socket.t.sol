// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PRBTest } from "@prb/test/PRBTest.sol";
import { console2 } from "forge-std/console2.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { TestPlus } from "../tests/TestPlus.sol";

import { PlugFactorySocket } from "./Plug.Factory.Socket.sol";
import { PlugVaultSocket } from "./Plug.Vault.Socket.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugFactorySocketTest is PRBTest, StdCheats, TestPlus {
    PlugVaultSocket internal implementation;
    PlugFactorySocket internal factory;

    PlugVaultSocket internal vault;

    function setUp() public virtual {
        implementation = new PlugVaultSocket();
        factory = new PlugFactorySocket("PlugMockSocket", "0.0.0");

        (, address vaultAddress) = factory.deploy(address(implementation), address(this), bytes32(0));
        vault = PlugVaultSocket(payable(vaultAddress));
    }

    function test_SingletonUse(uint256) public {
        vm.deal(address(vault), 100 ether);
        vm.expectRevert("PlugVaultSocket:already-initialized");
        vault.initialize(address(this));
    }

    function test_ToggleSigner(uint256) public {
        assertEq(vault.isSigner(address(this)), true);
        address nonZeroAddress = _randomNonZeroAddress();
        assertEq(vault.isSigner(nonZeroAddress), false);
        vault.toggleSigner(nonZeroAddress);
        assertEq(vault.isSigner(nonZeroAddress), true);
    }
}
