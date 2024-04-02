// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib
} from "../abstracts/test/Plug.Test.sol";
import { PlugRateLimitFuse } from "./Plug.RateLimit.Fuse.sol";

contract PlugCooldownFuseTest is Test {
    PlugRateLimitFuse internal fuse;

    PlugTypesLib.Current current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32(0);

    bool global;
    uint32 replenishRate = 60;
    uint32 max = 3;

    function setUp() public virtual {
        skip(2 days);

        fuse = new PlugRateLimitFuse();
    }

    function test_enforceFuse_Encoding() public {
        bytes memory terms = fuse.encode(global, replenishRate, max);
        (
            bool decodedGlobal,
            uint32 decodedReplenishRate,
            uint32 decodedMax
        ) = fuse.decode(terms);
        assertEq(decodedGlobal, global);
        assertEq(decodedReplenishRate, replenishRate);
        assertEq(decodedMax, max);
    }

    function test_enforceFuse_RateLimit() public {
        bytes memory terms = fuse.encode(global, replenishRate, max);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_RateLimit_Max() public {
        bytes memory terms = fuse.encode(global, replenishRate, max);
        fuse.enforceFuse(terms, current, plugsHash);
        fuse.enforceFuse(terms, current, plugsHash);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_RateLimit_ThresholdExceeded()
        public
    {
        bytes memory terms = fuse.encode(global, replenishRate, 1);
        fuse.enforceFuse(terms, current, plugsHash);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector, 1, 0
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_RateLimit_Replenished() public {
        bytes memory terms = fuse.encode(global, replenishRate, 1);
        fuse.enforceFuse(terms, current, plugsHash);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector, 1, 0
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
        skip(61);
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
