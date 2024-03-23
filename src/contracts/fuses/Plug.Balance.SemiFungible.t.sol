// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugMockERC20,
    PlugMockERC721
} from "../abstracts/test/Plug.Test.sol";
import { PlugBalanceSemiFungibleFuse } from
    "./Plug.Balance.SemiFungible.Fuse.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";

contract PlugBalanceSemiFungibleFuseTest is Test, Receiver {
    PlugBalanceSemiFungibleFuse internal fuse;

    PlugTypesLib.Current current = PlugTypesLib.Current({
        target: address(fuse),
        value: 0,
        data: "0x"
    });
    bytes32 plugsHash = bytes32(0);

    uint8 belowOperator;
    uint8 aboveOperator = 1;

    uint256 tokenId = 0;
    uint256 balance = 100;
    uint256 belowBalance = balance + 1;
    uint256 aboveBalance = balance - 1;

    function setUp() public virtual {
        setUpPlug();

        fuse = new PlugBalanceSemiFungibleFuse();

        mockERC1155.mint(address(this), tokenId, balance, "");
    }

    function test_enforceFuse_Encoding() public {
        bytes memory terms = fuse.encode(
            address(this),
            address(mockERC1155),
            tokenId,
            belowOperator,
            belowBalance
        );
        (
            address decodedHolder,
            address decodedAsset,
            uint256 decodedTokenId,
            uint8 decodedOperator,
            uint256 decodedBalance
        ) = fuse.decode(terms);

        assertEq(decodedHolder, address(this));
        assertEq(decodedAsset, address(mockERC1155));
        assertEq(decodedTokenId, tokenId);
        assertEq(decodedOperator, belowOperator);
        assertEq(decodedBalance, belowBalance);
    }

    function test_enforceFuse_BelowERC1155Balance() public view {
        bytes memory terms = fuse.encode(
            address(this),
            address(mockERC1155),
            tokenId,
            belowOperator,
            belowBalance
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_BelowERC1155Balance_Exceeded() public {
        bytes memory terms = fuse.encode(
            address(this),
            address(mockERC1155),
            tokenId,
            belowOperator,
            aboveBalance
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdExceeded.selector,
                aboveBalance,
                balance
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function test_enforceFuse_AboveERC1155Balance() public view {
        bytes memory terms = fuse.encode(
            address(this),
            address(mockERC1155),
            tokenId,
            aboveOperator,
            aboveBalance
        );

        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_AboveERC1155Balance_Insufficient()
        public
    {
        bytes memory terms = fuse.encode(
            address(this),
            address(mockERC1155),
            tokenId,
            aboveOperator,
            belowBalance
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ThresholdInsufficient.selector,
                belowBalance,
                balance
            )
        );
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
