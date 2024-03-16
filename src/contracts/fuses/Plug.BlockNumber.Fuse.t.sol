// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Test } from "../utils/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugBlockNumberFuse } from "./Plug.BlockNumber.Fuse.sol";

contract PlugBlockNumberFuseTest is Test {
    PlugBlockNumberFuse internal fuse;

    PlugTypesLib.Current current =
        PlugTypesLib.Current({ target: address(fuse), value: 0, data: "0x" });
    bytes32 plugsHash = bytes32(0);

    uint128 beforeOperator;
    uint128 beforeBlock;
    uint128 afterOperator;
    uint128 afterBlock;

    function setUp() public virtual {
        fuse = new PlugBlockNumberFuse();

        skip(12 * 80);

        beforeOperator = 0;
        beforeBlock = uint128(block.number + 1);
        afterOperator = 1;
        afterBlock = uint128(block.number - 1);
    }

    function test_enforceFuse_BeforeBlock() public {
        bytes memory terms = fuse.encode(beforeOperator, beforeBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) = fuse.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeBlock);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_BeforeBlock_Expired() public {
        bytes memory terms =
            fuse.encode(beforeOperator, uint128(block.number - 1));
        vm.expectRevert(bytes("PlugBlockNumberFuse:expired"));
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_AfterBlock() public {
        bytes memory terms = fuse.encode(afterOperator, afterBlock);
        (uint256 decodedOperator, uint256 decodedTimestamp) = fuse.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterBlock);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_AfterBlock_Early() public {
        bytes memory terms =
            fuse.encode(afterOperator, uint128(block.number + 1));
        vm.expectRevert(bytes("PlugBlockNumberFuse:early"));
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
