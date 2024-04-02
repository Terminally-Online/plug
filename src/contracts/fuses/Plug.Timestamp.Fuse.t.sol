// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib
} from "../abstracts/test/Plug.Test.sol";
import { PlugTimestampFuse } from "./Plug.Timestamp.Fuse.sol";

contract PlugTimestampFuseTest is Test {
    PlugTimestampFuse internal fuse;
    PlugTypesLib.Current internal current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32("0");

    uint8 beforeOperator;
    uint8 afterOperator = 1;

    uint256 beforeTimestamp;
    uint256 afterTimestamp;

    function setUp() public virtual {
        fuse = new PlugTimestampFuse();

        skip(200);

        beforeTimestamp = block.timestamp + 100;
        afterTimestamp = block.timestamp - 100;
    }

    function test_enforceFuse_BeforeTimestamp() public {
        bytes memory terms =
            fuse.encode(beforeOperator, beforeTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) =
            fuse.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeTimestamp);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_BeforeTimestamp_Expired()
        public
    {
        uint256 expected = beforeTimestamp - 150;
        bytes memory terms = fuse.encode(beforeOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector,
                expected,
                block.timestamp
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_AfterTimestamp() public {
        bytes memory terms =
            fuse.encode(afterOperator, afterTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) =
            fuse.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterTimestamp);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_AfterTimestamp_Early() public {
        uint256 expected = afterTimestamp + 400;
        bytes memory terms = fuse.encode(afterOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdInsufficient.selector,
                expected,
                block.timestamp
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
