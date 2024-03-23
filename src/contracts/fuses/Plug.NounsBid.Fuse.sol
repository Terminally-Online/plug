// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugFuseInterface } from
    "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../libraries/Plug.Lib.sol";
import { PlugNounsLib } from "../libraries/fuses/Plug.Nouns.Lib.sol";

/**
 * @title Plug Nouns Bid Fuse
 * @notice This Fuse enables the confirmation that a Nouns bid is valid.
 * @notice Use cases for bidding on Nouns:
 *     - Actively participate in the Nouns auction every day without needing
 *       to manually bid or even remember to bid.
 *     - Ensure that the treasury value is protected from undervalued closing
 *       bids that damage the underlying value of Nouns.
 * @author nftchance (chance@onplug.io)
 */
contract PlugNounsBidFuse is PlugFuseInterface {
    /**
     * See {PlugFuseInterface-enforceFuse}.
     */
    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32
    )
        public
        view
        override
        returns (bytes memory $through)
    {
        (address $bidder, uint256 $bid) = decode($live);

        /// @dev Get the current state of the auction.
        (,,,, address $winner,) = PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev Prevent the user from bidding on an auction that they
        ///      have already won / are winning.
        if ($winner == $bidder) {
            revert PlugNounsLib.InsufficientReason();
        }

        /// @dev Make sure the Socket has enough money to bid.
        if (msg.sender.balance < $bid) {
            revert PlugNounsLib.InsufficientBalance();
        }

        /// @dev Make sure the bid - fees is large enough to cover
        ///		 the minimum bid.

        /// @dev Callback to transfer the fee to the protocol.
        $through = $current.data;
    }

    function decode(bytes calldata $live)
        public
        pure
        returns (address $bidder, uint256 $bid)
    {
        /// @dev Decode the live data.
        ($bidder, $bid) = abi.decode($live, (address, uint256));
    }

    function encode(
        address $bidder,
        uint256 $bid
    )
        public
        pure
        returns (bytes memory $live)
    {
        /// @dev Encode the live data.
        $live = abi.encode($bidder, $bid);
    }
}
