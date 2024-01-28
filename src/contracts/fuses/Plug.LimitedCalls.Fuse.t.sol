// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugLimitedCallsFuse } from "./Plug.LimitedCalls.Fuse.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugLimitedCallsFuseTest is Test {
    PlugLimitedCallsFuse internal fuse;

    PlugTypesLib.Current current =
        PlugTypesLib.Current({ ground: address(fuse), voltage: 0, data: "0x" });
    bytes32 pinHash = bytes32("0");

    function setUp() public virtual {
        fuse = new PlugLimitedCallsFuse();
    }

    function test_enforceFuse() public {
        uint256 calls = 1;
        bytes memory terms = fuse.encode(calls);
        uint256 decodedCalls = fuse.decode(terms);
        assertEq(decodedCalls, calls);
        fuse.enforceFuse(terms, current, pinHash);
        vm.expectRevert(bytes("PlugLimitedCallsFuse:limit-exceeded"));
        fuse.enforceFuse(terms, current, pinHash);
    }

    function test_enforceFuseWithZeroCalls() public {
        uint256 calls = 0;
        bytes memory terms = fuse.encode(calls);
        vm.expectRevert(bytes("PlugLimitedCallsFuse:limit-exceeded"));
        fuse.enforceFuse(terms, current, pinHash);
    }
}
