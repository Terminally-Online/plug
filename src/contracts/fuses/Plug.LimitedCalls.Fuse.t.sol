// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Test, PlugTypesLib, LibClone } from "../abstracts/test/Plug.Test.sol";
import { PlugLimitedCallsFuse } from "./Plug.LimitedCalls.Fuse.sol";

contract PlugLimitedCallsFuseTest is Test {
    PlugLimitedCallsFuse internal fuse;

    PlugTypesLib.Current current =
        PlugTypesLib.Current({ target: address(fuse), value: 0, data: "0x" });
    bytes32 plugsHash = bytes32(0);

    function setUp() public virtual {
        fuse = new PlugLimitedCallsFuse();
    }

    function test_enforceFuse() public {
        uint256 calls = 1;
        bytes memory terms = fuse.encode(1);
        uint256 decodedCalls = fuse.decode(terms);
        assertEq(decodedCalls, calls);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_Exceeded() public {
        uint256 calls = 1;
        bytes memory terms = fuse.encode(calls);
        fuse.enforceFuse(terms, current, plugsHash);
        vm.expectRevert(bytes("PlugLimitedCallsFuse:limit-exceeded"));
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_ZeroCalls() public {
        uint256 calls = 0;
        bytes memory terms = fuse.encode(calls);
        vm.expectRevert(bytes("PlugLimitedCallsFuse:limit-exceeded"));
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
