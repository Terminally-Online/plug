// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { NounsTokenInterface } from
    "../interfaces/Nouns.Token.Interface.sol";
import { NounsAuctionHouseInterface } from
    "../interfaces/Nouns.AuctionHouse.Interface.sol";

library PlugNounsLib {
    NounsTokenInterface internal constant TOKEN = NounsTokenInterface(
        0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03
    );
    NounsAuctionHouseInterface internal constant AUCTION_HOUSE =
    NounsAuctionHouseInterface(
        0x830BD73E4184ceF73443C15111a1DF14e495C706
    );

    error InsufficientBalance();
    error InsufficientBid();
    error InsufficientReason();
    error InsufficientSettlement();
    error InsufficientOwnership();
    error InvalidSender();

    event Give(
        address indexed sender, address indexed asset, uint256 value
    );
    event Bid(
        address indexed sender,
        address indexed onBehalf,
        uint256 value,
        uint256 nounId
    );
    event Take(
        address indexed sender,
        address indexed onBehalf,
        address indexed asset,
        uint256 nounId
    );
}
