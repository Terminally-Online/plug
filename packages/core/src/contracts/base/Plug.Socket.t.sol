// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {
    Test,
    PlugLib,
    PlugEtcherLib,
    PlugTypesLib,
    Plug,
    PlugFactory,
    PlugSocket,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";
import { Ownable } from "solady/auth/Ownable.sol";

contract PlugSocketTest is Test {
    function setUp() public virtual {
        setUpPlug();
    }

    function test_name() public {
        assertEq(socket.name(), "Plug Socket");
    }

    function test_symbol() public {
        assertEq(socket.symbol(), "PLUGS");
    }

    function testRevert_Initialize_Again() public {
        vm.deal(address(socket), 100 ether);
        socket.initialize(signer, oneClicker);
    }

    function test_owner_Implementation() public {
        assertEq(socketImplementation.owner(), address(1));
    }

    function test_owner() public {
        assertEq(socket.owner(), signer);
    }

    function test_transferOwnership() public {
        vm.prank(signer);
        socket.transferOwnership(_randomNonZeroAddress());
    }

    function testRevert_transferOwnership() public {
        vm.expectRevert(Ownable.Unauthorized.selector);
        socket.transferOwnership(_randomNonZeroAddress());
    }

    function test_domainHashSet() public {
        // Deploy a new vault
        socket = deployVault();
        
        // Get the expected domain hash by constructing it manually
        bytes32 expectedDomainHash = socket.getEIP712DomainHash(
            PlugTypesLib.EIP712Domain({
                name: socket.name(),
                version: socket.version(),
                chainId: block.chainid,
                verifyingContract: address(socket)
            })
        );

        // Compare the actual domain hash with the expected one
        assertEq(
            socket.domainHash(),
            expectedDomainHash,
            "Domain hash should match expected value"
        );
        
        // Verify it's not zero
        assertTrue(socket.domainHash() != 0x0, "Domain hash should not be zero");
    }
}
