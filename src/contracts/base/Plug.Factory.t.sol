// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Test } from "../utils/Test.sol";

import { PlugFactory } from "./Plug.Factory.sol";
import { PlugVaultSocket } from "../sockets/Plug.Vault.Socket.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugFactoryTest is Test {
    PlugVaultSocket internal implementation;
    PlugFactory internal factory;

    address factoryOwner;
    string baseURI = "https://onplug.io/metadata/";

    function setUp() public virtual {
        factoryOwner = _randomNonZeroAddress();

        implementation = new PlugVaultSocket();
        factory = new PlugFactory(factoryOwner, baseURI);
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
        (, bool isSigner) = PlugVaultSocket(payable(vault)).getAccess(owner);
        assertEq(isSigner, true);
    }

    function test_InitCodeHash() public view {
        factory.initCodeHash(address(implementation));
    }
}
