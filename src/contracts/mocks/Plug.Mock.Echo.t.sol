// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Test } from "../utils/Test.sol";

import { PlugTypes, PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugFactory } from "../base/Plug.Factory.sol";
import { PlugMockEcho } from "./Plug.Mock.Echo.sol";

import "forge-std/console.sol";

contract PlugMockSocketTest is Test {
    PlugMockEcho internal mock;

    event EchoInvoked(address $sender, string $message);

    function setUp() public virtual {
        mock = new PlugMockEcho();
    }

    function test_Echo() public {
        string memory expected = "Hello World";
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(this), expected);
        mock.echo(expected);
    }

    function test_EmptyEcho() public {
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(this), "Hello World");
        mock.emptyEcho();
    }

    function test_MutedEcho(uint256 $echo) public view {
        mock.mutedEcho($echo);
    }
}
