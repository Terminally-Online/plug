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

    function test_setTargetsAllowed() public {
        address[] memory targets = new address[](1);
        targets[0] = address(0xbeef);

        vm.prank(factoryOwner);
        treasury.setTargetsAllowed(targets, true);
        assertTrue(treasury.targetToAllowed(address(0xbeef)));
    }

    function testRevert_setTargetsAllowed() public {
        address[] memory targets = new address[](1);
        targets[0] = address(0xbeef);

        vm.expectRevert(Ownable.Unauthorized.selector);
        treasury.setTargetsAllowed(targets, true);
    }

    function test_Multicall_WithdrawETH() public {
        vm.deal(address(treasury), 1 ether);
        assertEq(address(treasury).balance, 1 ether);

        address[] memory targets = new address[](1);
        targets[0] = treasury.owner();

        uint256[] memory values = new uint256[](1);
        values[0] = 1 ether;

        bytes[] memory datas = new bytes[](1);
        datas[0] = "";

        vm.prank(factoryOwner);
        treasury.execute(targets, values, datas);
        assertEq(address(treasury).balance, 0);
    }

    function testRevert_Multicall_WithdrawETH() public {
        vm.deal(address(treasury), 1 ether);
        assertEq(address(treasury).balance, 1 ether);

        address[] memory targets = new address[](1);
        targets[0] = treasury.owner();

        uint256[] memory values = new uint256[](1);
        values[0] = 1 ether;

        bytes[] memory datas = new bytes[](1);
        datas[0] = "";

        vm.expectRevert(Ownable.Unauthorized.selector);
        treasury.execute(targets, values, datas);
        assertEq(address(treasury).balance, 1 ether);
    }
}
