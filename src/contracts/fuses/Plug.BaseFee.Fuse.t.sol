// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib
} from "../abstracts/test/Plug.Test.sol";
import { PlugBaseFeeFuse } from "./Plug.BaseFee.Fuse.sol";

contract PlugBaseFeeFuseTest is Test {
    PlugBaseFeeFuse internal fuse;

    PlugTypesLib.Current current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32(0);

    uint8 belowOperator;
    uint8 aboveOperator = 1;

    uint256 belowBaseFee;
    uint256 aboveBaseFee;

    function setUp() public virtual {
        fuse = new PlugBaseFeeFuse();

        vm.fee(25 gwei);
        belowBaseFee = block.basefee + 1;
        aboveBaseFee = block.basefee - 1;
    }

    function test_enforceFuse_BelowBaseFee() public {
        bytes memory terms = fuse.encode(belowOperator, belowBaseFee);
        (uint256 decodedOperator, uint256 decodedBaseFee) =
            fuse.decode(terms);
        assertEq(decodedOperator, belowOperator);
        assertEq(decodedBaseFee, belowBaseFee);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_BelowBaseFee_Exceeded() public {
        uint256 expected = block.basefee - 1;
        bytes memory terms = fuse.encode(belowOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector,
                expected,
                block.basefee
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_AboveBaseFee() public {
        bytes memory terms = fuse.encode(aboveOperator, aboveBaseFee);
        (uint256 decodedOperator, uint256 decodedBaseFee) =
            fuse.decode(terms);
        assertEq(decodedOperator, aboveOperator);
        assertEq(decodedBaseFee, aboveBaseFee);
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_AboveBaseFee_Insufficient()
        public
    {
        uint256 expected = block.basefee + 1;
        bytes memory terms = fuse.encode(aboveOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdInsufficient.selector,
                expected,
                block.basefee
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
