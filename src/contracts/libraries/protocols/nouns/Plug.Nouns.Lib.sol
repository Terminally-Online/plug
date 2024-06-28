// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { NounsArtInterface } from "../../../interfaces/protocols/nouns/Nouns.Art.Interface.sol";
import { NounsTokenInterface } from "../../../interfaces/protocols/nouns/Nouns.Token.Interface.sol";
import { NounsAuctionHouseInterface } from
    "../../../interfaces/protocols/nouns/Nouns.AuctionHouse.Interface.sol";

library PlugNounsLib {
    address internal constant NOUNS_ART_ADDRESS = 0x921687c7A40a7F209100Db16AA95f787f6B4E677;
    address internal constant NOUNS_TOKEN_ADDRESS = 0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03;
    address internal constant NOUNS_AUCTION_HOUSE_ADDRESS =
        0x830BD73E4184ceF73443C15111a1DF14e495C706;

    NounsArtInterface internal constant ART = NounsArtInterface(NOUNS_ART_ADDRESS);
    NounsTokenInterface internal constant TOKEN = NounsTokenInterface(NOUNS_TOKEN_ADDRESS);
    NounsAuctionHouseInterface internal constant AUCTION_HOUSE =
        NounsAuctionHouseInterface(NOUNS_AUCTION_HOUSE_ADDRESS);

    /// @dev Selector hashes of the trait getters.
    bytes32 public constant BACKGROUND_SELECTOR = keccak256("background");
    bytes32 public constant HEAD_SELECTOR = keccak256("head");
    bytes32 public constant GLASSES_SELECTOR = keccak256("glasses");
    bytes32 public constant BODY_SELECTOR = keccak256("body");
    bytes32 public constant ACCESSORY_SELECTOR = keccak256("accessory");

    error InvalidSelector(bytes32 $selector);
    error InvalidTrait(bytes32 $expected, bytes32 $reality);

    error InsufficientBalance();
    error InsufficientBid();
    error InsufficientReason();
    error InsufficientSettlement();
    error InsufficientOwnership();
    error InvalidSender();

    event Give(address indexed sender, address indexed asset, uint256 value);
    event Bid(address indexed sender, address indexed onBehalf, uint256 value, uint256 nounId);
    event Take(
        address indexed sender, address indexed onBehalf, address indexed asset, uint256 nounId
    );
}
