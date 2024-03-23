// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import {
    Test,
    PlugEtcherLib,
    LibClone,
    PlugFactory,
    PlugVaultSocket
} from "../abstracts/test/Plug.Test.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract PlugTreasuryTest is Test {
    function setUp() public virtual {
        setUpPlug();
    }

    function test_Execute_WithdrawETH() public {
        vm.deal(address(treasury), 1 ether);
        assertEq(address(treasury).balance, 1 ether);
        address to = treasury.owner();
        uint256 value = 1 ether;
        bytes memory data = "";
        vm.prank(factoryOwner);
        treasury.execute(to, value, data);
        assertEq(address(treasury).balance, 0);
    }

    function testRevert_Execute_Unauthorized() public {
        vm.deal(address(treasury), 1 ether);
        address to = treasury.owner();
        uint256 value = 1 ether;
        bytes memory data = "";
        vm.expectRevert(Ownable.Unauthorized.selector);
        treasury.execute(to, value, data);
    }
}
