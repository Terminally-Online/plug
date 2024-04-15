// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugMockERC20,
    PlugMockERC721
} from "../abstracts/test/Plug.Test.sol";
import { PlugBalanceSemiFungible } from "./Plug.Balance.SemiFungible.sol";
import { Receiver } from "solady/accounts/Receiver.sol";

contract PlugBalanceSemiFungibleTest is Test, Receiver {
    PlugBalanceSemiFungible internal connector;

    bytes32 plugsHash = bytes32(0);
    uint8 belowOperator;
    uint8 aboveOperator = 1;
    uint256 tokenId = 0;
    uint256 balance = 100;
    uint256 belowBalance = balance + 1;
    uint256 aboveBalance = balance - 1;

    function setUp() public virtual {
        setUpPlug();

        connector = new PlugBalanceSemiFungible();

        mockERC1155.mint(address(this), tokenId, balance, "");
    }

    function test_enforce_Encoding() public {
        bytes memory terms = connector.encode(
            address(this), address(mockERC1155), tokenId, belowOperator, belowBalance
        );
        (
            address decodedHolder,
            address decodedAsset,
            uint256 decodedTokenId,
            uint8 decodedOperator,
            uint256 decodedBalance
        ) = connector.decode(terms);

        assertEq(decodedHolder, address(this));
        assertEq(decodedAsset, address(mockERC1155));
        assertEq(decodedTokenId, tokenId);
        assertEq(decodedOperator, belowOperator);
        assertEq(decodedBalance, belowBalance);
    }

    function test_enforce_BelowERC1155Balance() public view {
        bytes memory terms = connector.encode(
            address(this), address(mockERC1155), tokenId, belowOperator, belowBalance
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowERC1155Balance_Exceeded() public {
        bytes memory terms = connector.encode(
            address(this), address(mockERC1155), tokenId, belowOperator, aboveBalance
        );
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, aboveBalance, balance)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AboveERC1155Balance() public view {
        bytes memory terms = connector.encode(
            address(this), address(mockERC1155), tokenId, aboveOperator, aboveBalance
        );

        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AboveERC1155Balance_Insufficient() public {
        bytes memory terms = connector.encode(
            address(this), address(mockERC1155), tokenId, aboveOperator, belowBalance
        );
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, belowBalance, balance)
        );
        connector.enforce(terms, plugsHash);
    }
}
