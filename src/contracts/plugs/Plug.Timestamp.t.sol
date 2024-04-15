// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugTimestamp } from "./Plug.Timestamp.sol";

contract PlugTimestampTest is Test {
    PlugTimestamp internal connector;

    bytes32 plugsHash = bytes32("0");

    uint8 beforeOperator;
    uint8 afterOperator = 1;

    uint256 beforeTimestamp;
    uint256 afterTimestamp;

    function setUp() public virtual {
        connector = new PlugTimestamp();

        skip(200);

        beforeTimestamp = block.timestamp + 100;
        afterTimestamp = block.timestamp - 100;
    }

    function test_enforce_BeforeTimestamp() public {
        bytes memory terms = connector.encode(beforeOperator, beforeTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) = connector.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeTimestamp);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_BeforeTimestamp_Expired() public {
        uint256 expected = beforeTimestamp - 150;
        bytes memory terms = connector.encode(beforeOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, block.timestamp)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AfterTimestamp() public {
        bytes memory terms = connector.encode(afterOperator, afterTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) = connector.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterTimestamp);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AfterTimestamp_Early() public {
        uint256 expected = afterTimestamp + 400;
        bytes memory terms = connector.encode(afterOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdInsufficient.selector, expected, block.timestamp
            )
        );
        connector.enforce(terms, plugsHash);
    }
}
