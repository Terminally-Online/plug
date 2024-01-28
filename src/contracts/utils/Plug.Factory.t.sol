// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "./Test.sol";

import { PlugFactory } from "./Plug.Factory.sol";
import { PlugVaultSocket } from "../sockets/Plug.Vault.Socket.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugFactoryTest is Test {
    PlugVaultSocket internal implementation;
    PlugFactory internal factory;

    function setUp() public virtual {
        implementation = new PlugVaultSocket();
        factory = new PlugFactory();
    }

    function test_DeployDeterministic(uint256) public {
        vm.deal(address(this), 100 ether);
        address owner = _randomNonZeroAddress();
        uint256 initialValue = _random() % 100 ether;
        bytes32 salt = _random() % 8 == 0
            ? bytes32(_random())
            : bytes32(uint256(uint96(_random())));

        bool alreadyDeployed;
        address vault;

        if (uint256(salt) >> 96 != uint160(owner) && uint256(salt) >> 96 != 0) {
            vm.expectRevert(LibClone.SaltDoesNotStartWith.selector);
            (alreadyDeployed, vault) = factory.deploy{ value: initialValue }(
                address(implementation), owner, salt
            );
            return;
        } else {
            (alreadyDeployed, vault) = factory.deploy{ value: initialValue }(
                address(implementation), owner, salt
            );
        }

        assertEq(address(vault).balance, initialValue);
        assertEq(PlugVaultSocket(payable(vault)).isSigner(owner), true);
    }

    function test_RepeatedDeployDeterministic() public {
        bytes32 salt = bytes32(_random() & uint256(type(uint96).max));
        address expectedInstance =
            factory.getAddress(address(implementation), salt);
        (, address instance) = factory.deploy{ value: 123 }(
            address(implementation), address(this), salt
        );
        assertEq(instance.balance, 123);
        (bool alreadyDeployed,) = factory.deploy{ value: 456 }(
            address(implementation), address(this), salt
        );
        assertEq(alreadyDeployed, true);
        factory.deploy(address(implementation), address(this), salt);
        assertEq(alreadyDeployed, true);
        assertEq(instance.balance, 123 + 456);
        assertEq(instance, expectedInstance);
    }

    function test_RepeatedDeployDeterministic(uint256) public {
        address owner = _randomNonZeroAddress();
        bytes32 salt = bytes32(
            (_random() & uint256(type(uint96).max))
                | (uint256(uint160(owner)) << 96)
        );
        address expectedInstance =
            factory.getAddress(address(implementation), salt);
        address notOwner = _randomNonZeroAddress();
        while (owner == notOwner) notOwner = _randomNonZeroAddress();
        vm.expectRevert(LibClone.SaltDoesNotStartWith.selector);
        factory.deploy{ value: 123 }(address(implementation), notOwner, salt);
        (, address instance) =
            factory.deploy{ value: 123 }(address(implementation), owner, salt);
        assertEq(instance.balance, 123);
        vm.expectRevert(LibClone.SaltDoesNotStartWith.selector);
        factory.deploy{ value: 123 }(address(implementation), notOwner, salt);
        (, address redeployedInstance) =
            factory.deploy{ value: 456 }(address(implementation), owner, salt);
        assertEq(redeployedInstance, instance);
        assertEq(instance.balance, 123 + 456);
        assertEq(expectedInstance, instance);
    }

    function test_InitCodeHash() public view {
        factory.initCodeHash(address(implementation));
    }
}
