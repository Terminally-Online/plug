// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugRateLimit } from "./Plug.RateLimit.sol";

contract PlugRateLimitTest is Test {
    PlugRateLimit internal connector;

    bytes32 plugsHash = bytes32(0);

    bool global;
    uint32 replenishRate = 60;
    uint32 max = 3;

    function setUp() public virtual {
        skip(2 days);

        connector = new PlugRateLimit();
    }

    function test_enforce_Encoding() public {
        bytes memory terms = connector.encode(global, replenishRate, max);
        (bool decodedGlobal, uint32 decodedReplenishRate, uint32 decodedMax) =
            connector.decode(terms);
        assertEq(decodedGlobal, global);
        assertEq(decodedReplenishRate, replenishRate);
        assertEq(decodedMax, max);
    }

    function test_enforce_RateLimit() public {
        bytes memory terms = connector.encode(global, replenishRate, max);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_RateLimit_Max() public {
        bytes memory terms = connector.encode(global, replenishRate, max);
        connector.enforce(terms, plugsHash);
        connector.enforce(terms, plugsHash);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_RateLimit_ThresholdExceeded() public {
        bytes memory terms = connector.encode(global, replenishRate, 1);
        connector.enforce(terms, plugsHash);
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, 1, 0));
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_RateLimit_Replenished() public {
        bytes memory terms = connector.encode(global, replenishRate, 1);
        connector.enforce(terms, plugsHash);
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, 1, 0));
        connector.enforce(terms, plugsHash);
        skip(61);
        connector.enforce(terms, plugsHash);
    }
}
