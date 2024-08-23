// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test, PlugEtcherLib, LibClone, PlugFactory, PlugLib
} from "../abstracts/test/Plug.Test.sol";

contract PlugFactoryTest is Test {
    event Transfer(address indexed from, address indexed to, uint256 indexed id);

    function setUp() public virtual {
        setUpPlug();
    }

    function test_salt() public {
        bytes memory salt =
            abi.encode(uint96(1738), signer, oneClicker, address(socketImplementation));

        /// @dev Decode the details used to deploy the Socket and guard the signature.
        (
            uint96 nonce,
            address adminAddress,
            address oneClickerAddress,
            address implementationAddress
        ) = abi.decode(salt, (uint96, address, address, address));

        assertEq(nonce, uint96(1738));
        assertEq(adminAddress, signer);
        assertEq(oneClickerAddress, oneClicker);
        assertEq(implementationAddress, address(socketImplementation));

        /// @dev Decode the single nonce used to guard the signature.
        (uint96 singleNonce) = abi.decode(salt, (uint96));
        assertEq(nonce, singleNonce);
    }

    function test_DeployDeterministic(uint256) public {
        vm.deal(address(this), 1000 ether);
        uint256 initialValue = _random() % 100 ether;
        bytes memory salt =
            abi.encode(uint96(1738), signer, oneClicker, address(socketImplementation));
        (, address vault) = factory.deploy{ value: initialValue }(salt);
        assertEq(address(vault).balance, initialValue);
        (bool alreadyDeployed,) = factory.deploy{ value: initialValue }(salt);
        assertTrue(alreadyDeployed);
    }

    function testRevert_InvalidImplementation(uint256) public {
        bytes memory salt = abi.encode(uint96(1738), signer, oneClicker, address(0));
        vm.expectRevert(abi.encodeWithSelector(PlugLib.SaltInvalid.selector, address(0), signer));
        factory.deploy(salt);
    }

    function testRevert_InvalidAdmin(uint256) public {
        bytes memory salt =
            abi.encode(uint96(1738), address(0), oneClicker, address(socketImplementation));
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.SaltInvalid.selector, address(socketImplementation), address(0)
            )
        );
        factory.deploy(salt);
    }
}
