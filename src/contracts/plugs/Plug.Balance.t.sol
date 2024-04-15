// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugMockERC20,
    PlugMockERC721
} from "../abstracts/test/Plug.Test.sol";
import { PlugBalance } from "./Plug.Balance.sol";

contract PlugBalanceTest is Test {
    PlugBalance internal connector;

    bytes32 plugsHash = bytes32(0);
    uint8 belowOperator;
    uint8 aboveOperator = 1;
    uint256 belowBalance = 100 + 1;
    uint256 aboveBalance = 100 - 1;
    uint8 nativeType;
    uint8 erc20Type = 1;
    uint8 erc721Type = 2;
    uint256 balance = 100;

    function setUp() public virtual {
        setUpPlug();

        connector = new PlugBalance();

        vm.deal(address(this), balance);
        mockERC20.mint(address(this), balance);
        mockERC721.mint(address(this), balance);
    }

    function test_enforce_Encoding() public {
        bytes memory terms =
            connector.encode(address(this), address(0), nativeType, belowOperator, belowBalance);
        (
            address decodedHolder,
            address decodedAsset,
            uint8 decodedType,
            uint8 decodedOperator,
            uint256 decodedBalance
        ) = connector.decode(terms);

        assertEq(decodedHolder, address(this));
        assertEq(decodedAsset, address(0));
        assertEq(decodedType, nativeType);
        assertEq(decodedOperator, belowOperator);
        assertEq(decodedBalance, belowBalance);
    }

    function test_enforce_BelowNativeBalance() public view {
        bytes memory terms =
            connector.encode(address(this), address(0), nativeType, belowOperator, belowBalance);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowNativeBalance_Exceeded() public {
        uint256 expected = balance - 1;
        bytes memory terms =
            connector.encode(address(this), address(0), nativeType, belowOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, balance)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AboveNativeBalance() public view {
        bytes memory terms =
            connector.encode(address(this), address(0), nativeType, aboveOperator, aboveBalance);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AboveNativeBalance_Insufficient() public {
        uint256 expected = balance + 1;
        bytes memory terms =
            connector.encode(address(this), address(0), nativeType, aboveOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, expected, balance)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowERC20Balance() public view {
        bytes memory terms = connector.encode(
            address(this), address(mockERC20), erc20Type, belowOperator, belowBalance
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowERC20Balance_Exceeded() public {
        uint256 expected = balance - 1;
        bytes memory terms =
            connector.encode(address(this), address(mockERC20), erc20Type, belowOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, balance)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AboveERC20Balance() public view {
        bytes memory terms = connector.encode(
            address(this), address(mockERC20), erc20Type, aboveOperator, aboveBalance
        );
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AboveERC20Balance_Insufficient() public {
        uint256 expected = balance + 1;
        bytes memory terms =
            connector.encode(address(this), address(mockERC20), erc20Type, aboveOperator, expected);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, expected, balance)
        );
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowERC721Balance() public view {
        bytes memory terms =
            connector.encode(address(this), address(mockERC721), erc721Type, belowOperator, 2);
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_BelowERC721Balance_Exceeded() public {
        uint256 expected = 0;
        bytes memory terms = connector.encode(
            address(this), address(mockERC721), erc721Type, belowOperator, expected
        );
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdExceeded.selector, expected, 1));
        connector.enforce(terms, plugsHash);
    }

    function test_enforce_AboveERC721Balance() public view {
        bytes memory terms =
            connector.encode(address(this), address(mockERC721), erc721Type, aboveOperator, 0);
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_AboveERC721Balance_Insufficient() public {
        uint256 expected = 1 + 1;
        bytes memory terms = connector.encode(
            address(this), address(mockERC721), erc721Type, aboveOperator, expected
        );
        vm.expectRevert(abi.encodeWithSelector(PlugLib.ThresholdInsufficient.selector, expected, 1));
        connector.enforce(terms, plugsHash);
    }
}
