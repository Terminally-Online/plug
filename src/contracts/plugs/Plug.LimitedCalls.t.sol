// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib, LibClone } from "../abstracts/test/Plug.Test.sol";
import { PlugLimitedCalls } from "./Plug.LimitedCalls.sol";

contract PlugLimitedCallsTest is Test {
    PlugLimitedCalls internal connector;

    bytes32 plugsHash = bytes32(0);

    function setUp() public virtual {
        connector = new PlugLimitedCalls();
    }

    function test_enforce() public {
        uint256 calls = 1;
        bytes memory terms = connector.encode(1);
        uint256 decodedCalls = connector.decode(terms);
        assertEq(decodedCalls, calls);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_Exceeded() public {
        uint256 calls = 1;
        bytes memory terms = connector.encode(calls);
        connector.enforce(terms, plugsHash);
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, calls, 2));
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_ZeroCalls() public {
        uint256 calls = 0;
        bytes memory terms = connector.encode(calls);
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, calls, 1));
        connector.enforce(terms, plugsHash);
    }
}
