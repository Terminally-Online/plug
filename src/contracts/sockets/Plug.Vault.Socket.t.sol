// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugEtcherLib,
    PlugTypesLib,
    Plug,
    PlugFactory,
    PlugVaultSocket,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";

contract PlugVaultSocketTest is Test {
    function setUp() public virtual {
        setUpPlug();
    }

    function deployVault()
        internal
        override
        returns (PlugVaultSocket $vault)
    {
        (, address vaultAddress) = factory.deploy(
            bytes32(abi.encodePacked(address(this), uint96(0))),
            address(plug)
        );
        $vault = PlugVaultSocket(payable(vaultAddress));
    }

    function test_name() public {
        assertEq(vault.name(), "Plug Vault Socket");
    }

    function test_symbol() public {
        assertEq(vault.symbol(), "PVS");
    }

    function testRevert_Initialize_Again() public {
        vm.deal(address(vault), 100 ether);
        vm.expectRevert(PlugLib.TradingAlreadyInitialized.selector);
        vault.initialize(address(this), address(plug));
    }

    function testRevert_ReinitializeImplementation() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ImplementationAlreadyInitialized.selector,
                uint16(0)
            )
        );
        vm.prank(factoryOwner);
        factory.setImplementation(0, address(vaultImplementation));
    }

    function testRevert_UninitializedImplementation() public {
        bytes32 salt = bytes32(
            abi.encodePacked(address(this), uint80(0), uint16(2))
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ImplementationInvalid.selector, uint16(2)
            )
        );
        factory.deploy(salt, address(plug));
    }

    function test_ownership_Implementation() public {
        assertEq(vaultImplementation.ownership(), address(1));
    }

    function testRevert_owner_Implementation() public {
        vm.expectRevert();
        vaultImplementation.owner();
    }

    function test_owner() public {
        assertEq(vault.owner(), address(this));
    }

    function test_transferOwnership_Token() public {
        factory.transferFrom(
            address(this),
            _randomNonZeroAddress(),
            uint256(uint160(address(vault)))
        );
    }

    function testRevert_transferOwnership() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.CallerInvalid.selector,
                address(factory),
                address(this)
            )
        );
        vault.transferOwnership(_randomNonZeroAddress());
    }
}
