// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugFactory } from "../utils/Plug.Factory.sol";
import { PlugVaultSocket } from "./Plug.Vault.Socket.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

contract PlugVaultSocketTest is Test {
    PlugVaultSocket internal implementation;
    PlugVaultSocket internal vault;
    PlugFactory internal factory;

    function setUp() public virtual {
        implementation = new PlugVaultSocket();
        factory = new PlugFactory();

        (, address vaultAddress) =
            factory.deploy(address(implementation), address(this), bytes32(0));
        vault = PlugVaultSocket(payable(vaultAddress));
    }

    function test_SingletonUse(uint256) public {
        vm.deal(address(vault), 100 ether);
        vm.expectRevert(Initializable.InvalidInitialization.selector);
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
