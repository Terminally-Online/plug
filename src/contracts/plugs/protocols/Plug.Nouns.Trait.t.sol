// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugMockERC20,
    PlugMockERC721
} from "../../abstracts/test/Plug.Test.sol";
import {
    PlugNounsLib,
    NounsArtInterface,
    NounsTokenInterface
} from "../../libraries/protocols/nouns/Plug.Nouns.Lib.sol";
import { NounsArtMock } from "../../mocks/protocols/nouns/Nouns.Art.Mock.sol";
import { NounsTokenMock } from "../../mocks/protocols/nouns/Nouns.Token.Mock.sol";
import { NounsAuctionHouseMock } from "../../mocks/protocols/nouns/Nouns.AuctionHouse.Mock.sol";
import { PlugNounsTrait } from "./Plug.Nouns.Trait.sol";
import { Receiver } from "solady/accounts/Receiver.sol";

contract PlugNounsTraitTest is Test, Receiver {
    PlugNounsTrait internal connector;

    NounsArtInterface internal art;
    NounsTokenInterface internal token;

    function setUp() public virtual {
        art = setUpArt();
        token = setUpToken();
        setUpActionHouse();

        setUpPlug();

        connector = new PlugNounsTrait();
    }

    function setUpArt() internal virtual returns (NounsArtInterface $nounsArt) {
        vm.etch(PlugNounsLib.NOUNS_ART_ADDRESS, address(new NounsArtMock()).code);
        $nounsArt = NounsArtMock(PlugNounsLib.NOUNS_ART_ADDRESS);
    }

    function setUpToken() internal virtual returns (NounsTokenInterface $nounsToken) {
        vm.etch(PlugNounsLib.NOUNS_TOKEN_ADDRESS, address(new NounsTokenMock()).code);
        $nounsToken = NounsTokenMock(PlugNounsLib.NOUNS_TOKEN_ADDRESS);
    }

    function setUpActionHouse()
        internal
        virtual
        returns (NounsAuctionHouseMock $nounsAuctionHouse)
    {
        vm.etch(PlugNounsLib.NOUNS_AUCTION_HOUSE_ADDRESS, address(new NounsAuctionHouseMock()).code);
        $nounsAuctionHouse = NounsAuctionHouseMock(PlugNounsLib.NOUNS_AUCTION_HOUSE_ADDRESS);
    }

    function test_mock() public {
        assertEq(art.backgrounds(0), string(bytes(abi.encode(keccak256("background")))));
        assertEq(art.heads(0), bytes(abi.encode(keccak256("head"))));
        assertEq(art.bodies(0), bytes(abi.encode(keccak256("body"))));
        assertEq(art.accessories(0), bytes(abi.encode(keccak256("accessory"))));
        assertEq(art.glasses(0), bytes(abi.encode(keccak256("glasses"))));

        (uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses) =
            token.seeds(0);
        assertEq(background, 0);
        assertEq(body, 0);
        assertEq(accessory, 0);
        assertEq(head, 0);
        assertEq(glasses, 0);
    }

    function test_encode() public {
        bytes memory terms =
            connector.encode(PlugNounsLib.BACKGROUND_SELECTOR, keccak256("backgrounds"));
        (bytes32 decodedSelector, bytes32 decodedTrait) = connector.decode(terms);

        assertEq(decodedSelector, PlugNounsLib.BACKGROUND_SELECTOR);
        assertEq(decodedTrait, keccak256("backgrounds"));
    }

    function testRevert_encode_InvalidSelector() public {
        vm.expectRevert(
            abi.encodeWithSelector(PlugNounsLib.InvalidSelector.selector, keccak256("invalid"))
        );
        connector.encode(keccak256("invalid"), keccak256("backgrounds"));
    }

    function testRevert_enforce_InvalidSelector() public {
        bytes memory terms = abi.encode(keccak256("invalid"), keccak256("head"));
        vm.expectRevert(
            abi.encodeWithSelector(PlugNounsLib.InvalidSelector.selector, keccak256("invalid"))
        );
        connector.enforce(terms, bytes32(0));
    }
}
