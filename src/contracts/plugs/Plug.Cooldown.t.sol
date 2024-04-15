// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugCooldown } from "./Plug.Cooldown.sol";

contract PlugCooldownTest is Test {
    PlugCooldown internal connector;

    bytes32 plugsHash = bytes32(0);
    uint256 cooldown = 1 hours;

    function setUp() public virtual {
        connector = new PlugCooldown();
    }

    function test_enforce_Encoding() public {
        bytes memory terms = connector.encode(cooldown);
        uint256 decodedCooldoown = connector.decode(terms);
        assertEq(decodedCooldoown, cooldown);
    }

    function test_enforce_FirstInteraction() public {
        bytes memory terms = connector.encode(cooldown);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_RepeatedInteraction() public {
        bytes memory terms = connector.encode(cooldown);
        connector.enforce(terms, plugsHash);
        skip(cooldown + 1);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_ThresholdExceeded() public {
        bytes memory terms = connector.encode(cooldown);
        connector.enforce(terms, plugsHash);
        skip(120);
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, cooldown, 120));
        connector.enforce(terms, plugsHash);
    }
}
