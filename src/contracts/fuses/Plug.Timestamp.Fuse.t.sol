// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../tests/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugTimestampFuse } from "./Plug.Timestamp.Fuse.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugTimestampFuseTest is Test {
    PlugTimestampFuse internal fuse;

    PlugTypesLib.Current current =
        PlugTypesLib.Current({ ground: address(fuse), voltage: 0, data: "0x" });
    bytes32 pinHash = bytes32("0");

    uint128 beforeOperator;
    uint128 beforeTimestamp;
    uint128 afterOperator;
    uint128 afterTimestamp;

    function setUp() public virtual {
        fuse = new PlugTimestampFuse();

        skip(200);

        beforeOperator = 0;
        beforeTimestamp = uint128(block.timestamp + 100);
        afterOperator = 1;
        afterTimestamp = uint128(block.timestamp - 100);
    }

    function test_enforceFuse_BeforeTimestamp() public {
        bytes memory terms = fuse.encode(beforeOperator, beforeTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) = fuse.decode(terms);
        assertEq(decodedOperator, beforeOperator);
        assertEq(decodedTimestamp, beforeTimestamp);
        fuse.enforceFuse(terms, current, pinHash);

        bytes memory revertingTerms =
            fuse.encode(beforeOperator, uint128(beforeTimestamp - 150));
        vm.expectRevert(bytes("PlugTimestampFuse:expired-pin"));
        fuse.enforceFuse(revertingTerms, current, pinHash);
    }

    function test_enforceFuse_AfterTimestamp() public {
        bytes memory terms = fuse.encode(afterOperator, afterTimestamp);
        (uint256 decodedOperator, uint256 decodedTimestamp) = fuse.decode(terms);
        assertEq(decodedOperator, afterOperator);
        assertEq(decodedTimestamp, afterTimestamp);
        fuse.enforceFuse(terms, current, pinHash);
        bytes memory revertingTerms =
            fuse.encode(afterOperator, uint128(afterTimestamp + 400));
        vm.expectRevert(bytes("PlugTimestampFuse:early-pin"));
        fuse.enforceFuse(revertingTerms, current, pinHash);
    }
}
