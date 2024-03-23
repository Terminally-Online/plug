// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Test, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";

import { PlugClampFuse } from "./Plug.Clamp.Fuse.sol";

contract PlugClampFuseTest is Test {
    PlugClampFuse internal fuse;

    PlugTypesLib.Current internal current = PlugTypesLib.Current({
        target: address(0),
        value: 0,
        data: abi.encode(uint256(51))
    });
    bytes32 plugsHash = bytes32(0);

    function setUp() public {
        fuse = new PlugClampFuse();
    }

    function test_EnforceFuse() public {
        bytes memory terms = fuse.encode(10, 50);
        bytes memory pass = fuse.enforceFuse(terms, current, plugsHash);
        assertEq(abi.decode(pass, (uint256)), 50);
    }
}
