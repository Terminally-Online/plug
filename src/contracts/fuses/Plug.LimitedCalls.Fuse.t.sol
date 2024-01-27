// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PRBTest } from "@prb/test/PRBTest.sol";
import { console2 } from "forge-std/console2.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { TestPlus } from "../tests/TestPlus.sol";

import { PlugTypes, PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { PlugLimitedCallsFuse } from "./Plug.LimitedCalls.Fuse.sol";
import { PlugMockERC1155 } from "../mocks/Plug.Mock.ERC1155.sol";

import { LibClone } from "solady/src/utils/LibClone.sol";

contract PlugLimitedCallsFuseTest is PRBTest, StdCheats, TestPlus {
    PlugLimitedCallsFuse internal fuse;
    PlugMockERC1155 internal mock;

    bytes internal encodedTransaction =
        abi.encodeWithSelector(mock.mint.selector);
    PlugTypesLib.Current current = PlugTypesLib.Current({
        ground: address(mock),
        voltage: 0,
        data: encodedTransaction
    });
    bytes32 pinHash = "0x01";

    function setUp() public virtual {
        fuse = new PlugLimitedCallsFuse();
    }

    function test_enforceFuse() public {
        uint256 calls = 1;
        bytes memory terms = fuse.encode(calls);
        uint256 decodedCalls = fuse.decode(terms);
        assertEq(decodedCalls, calls);
        fuse.enforceFuse(terms, current, pinHash);
    }

    function testRevert_enforceFuse() public {
        uint256 calls = 1;
        bytes memory terms = fuse.encode(calls);
        fuse.enforceFuse(terms, current, pinHash);
        vm.expectRevert(bytes("LimitedCallsEnforcer:limit-exceeded"));
        fuse.enforceFuse(terms, current, pinHash);
    }

    function testRevert_enforceFuseWithZeroCalls() public {
        uint256 calls = 0;
        bytes memory terms = fuse.encode(calls);
        vm.expectRevert(bytes("LimitedCallsEnforcer:limit-exceeded"));
        fuse.enforceFuse(terms, current, pinHash);
    }
}
