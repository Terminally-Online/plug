// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugBlockNumber } from "./Plug.BlockNumber.sol";

contract PlugBlockNumberTest is Test {
    PlugBlockNumber internal connector;

    bytes32 plugsHash = bytes32(0);

    uint8 beforeOperator;
    uint8 afterOperator = 1;
    uint256 beforeBlock;
    uint256 afterBlock;

    function setUp() public virtual {
        connector = new PlugBlockNumber();

        skip(12 * 80);

        beforeBlock = block.number + 1;
        afterBlock = block.number - 1;
    }

    function test_enforce_BeforeBlock() public {
        bytes memory terms = connector.encode(beforeOperator, beforeBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) = connector.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeBlock);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BeforeBlock_Expired() public {
        uint256 expected = block.number - 1;
        bytes memory terms = connector.encode(beforeOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, block.number)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AfterBlock() public {
        bytes memory terms = connector.encode(afterOperator, afterBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) = connector.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterBlock);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AfterBlock_Early() public {
        uint256 expected = block.number + 1;
        bytes memory terms = connector.encode(afterOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, expected, block.number)
        );
        connector.enforce(terms, plugsHash);
    }
}
