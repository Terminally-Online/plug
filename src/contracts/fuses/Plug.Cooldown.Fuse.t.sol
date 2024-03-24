// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import {
    Test,
    PlugLib,
    PlugTypesLib
} from "../abstracts/test/Plug.Test.sol";
import { PlugCooldownFuse } from "./Plug.Cooldown.Fuse.sol";

contract PlugCooldownFuseTest is Test {
    PlugCooldownFuse internal fuse;

    PlugTypesLib.Current current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32(0);

    uint256 cooldown = 1 hours;

    function setUp() public virtual {
        fuse = new PlugCooldownFuse();
    }

    function test_enforceFuse_Encoding() public {
        bytes memory terms = fuse.encode(cooldown);
        uint256 decodedCooldoown = fuse.decode(terms);
        assertEq(decodedCooldoown, cooldown);
    }

    function test_enforceFuse_FirstInteraction() public {
        bytes memory terms = fuse.encode(cooldown);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_RepeatedInteraction() public {
        bytes memory terms = fuse.encode(cooldown);
        fuse.enforceFuse(terms, current, plugsHash);
        skip(cooldown + 1);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_ThresholdExceeded() public {
        bytes memory terms = fuse.encode(cooldown);
        fuse.enforceFuse(terms, current, plugsHash);
        skip(120);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector, cooldown, 120
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
