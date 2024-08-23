// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

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
        assertEq(socket.symbol(), "PS");
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
}
