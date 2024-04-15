// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugLib, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugBaseFee } from "./Plug.BaseFee.sol";

contract PlugBaseFeeTest is Test {
    PlugBaseFee internal connector;

    bytes32 plugsHash = bytes32(0);
    uint8 belowOperator;
    uint8 aboveOperator = 1;
    uint256 belowBaseFee;
    uint256 aboveBaseFee;

    function setUp() public virtual {
        connector = new PlugBaseFee();

        vm.fee(25 gwei);
        belowBaseFee = block.basefee + 1;
        aboveBaseFee = block.basefee - 1;
    }

    function test_enforce_BelowBaseFee() public {
        bytes memory terms = connector.encode(belowOperator, belowBaseFee);
        (uint256 decodedOperator, uint256 decodedBaseFee) = connector.decode(terms);
        assertEq(decodedOperator, belowOperator);
        assertEq(decodedBaseFee, belowBaseFee);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowBaseFee_Exceeded() public {
        uint256 expected = block.basefee - 1;
        bytes memory terms = connector.encode(belowOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, block.basefee)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AboveBaseFee() public {
        bytes memory terms = connector.encode(aboveOperator, aboveBaseFee);
        (uint256 decodedOperator, uint256 decodedBaseFee) = connector.decode(terms);
        assertEq(decodedOperator, aboveOperator);
        assertEq(decodedBaseFee, aboveBaseFee);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AboveBaseFee_Insufficient() public {
        uint256 expected = block.basefee + 1;
        bytes memory terms = connector.encode(aboveOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, expected, block.basefee)
        );
        connector.enforce(terms, plugsHash);
    }
}
