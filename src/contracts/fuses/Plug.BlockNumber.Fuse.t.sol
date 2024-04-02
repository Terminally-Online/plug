// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib
} from "../abstracts/test/Plug.Test.sol";
import { PlugBlockNumberFuse } from "./Plug.BlockNumber.Fuse.sol";

contract PlugBlockNumberFuseTest is Test {
    PlugBlockNumberFuse internal fuse;

    PlugTypesLib.Current current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32(0);

    uint8 beforeOperator;
    uint8 afterOperator = 1;
    uint256 beforeBlock;
    uint256 afterBlock;

    function setUp() public virtual {
        fuse = new PlugBlockNumberFuse();

        skip(12 * 80);

        beforeBlock = block.number + 1;
        afterBlock = block.number - 1;
    }

    function test_enforceFuse_BeforeBlock() public {
        bytes memory terms = fuse.encode(beforeOperator, beforeBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) =
            fuse.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeBlock);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_BeforeBlock_Expired() public {
        uint256 expected = block.number - 1;
        bytes memory terms = fuse.encode(beforeOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector,
                expected,
                block.number
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_AfterBlock() public {
        bytes memory terms = fuse.encode(afterOperator, afterBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) =
            fuse.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterBlock);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_AfterBlock_Early() public {
        uint256 expected = block.number + 1;
        bytes memory terms = fuse.encode(afterOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdInsufficient.selector,
                expected,
                block.number
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
