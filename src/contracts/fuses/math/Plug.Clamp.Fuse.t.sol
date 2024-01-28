// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../../utils/Test.sol";
import { PlugTypesLib } from "../../abstracts/Plug.Types.sol";

import { PlugClampFuse } from "./Plug.Clamp.Fuse.sol";

contract PlugClampFuseTest is Test {
    PlugClampFuse internal fuse;

    function setUp() public {
        fuse = new PlugClampFuse();
    }

    function test_EnforceFuse() public {
        bytes memory live = fuse.encode(10, 50);
        bytes memory pass = fuse.enforceFuse(
            live,
            PlugTypesLib.Current({
                ground: address(0),
                voltage: 0,
                data: abi.encodePacked(uint256(51))
            }),
            bytes32(0)
        );
        assertEq(abi.decode(pass, (uint256)), 50);
    }
}
