// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Test, PlugMockEcho } from "../abstracts/test/Plug.Test.sol";

contract PlugMockSocketTest is Test {
    event EchoInvoked(address $sender, string $message);

    function setUp() public virtual {
        setUpPlug();
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
