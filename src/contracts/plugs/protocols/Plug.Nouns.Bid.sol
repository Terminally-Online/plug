// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../../interfaces/Plug.Connector.Interface.sol";
import { PlugNounsLib } from "../../libraries/protocols/nouns/Plug.Nouns.Lib.sol";

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
contract PlugNounsBid is PlugConnectorInterface {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public override {
        uint256 $bid = decode($terms);

        /// @dev Get the current state of the auction.
        (,,, uint256 $endTime, address $winner, bool $settled) =
            PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev If the previous auction has not been settled and needs to be, it must be
        ///      settled before a new auction can be started and bid placed.
        if ($settled == false && block.timestamp >= $endTime) {
            /// @dev Place the bid on the upcoming noun.
            PlugNounsLib.AUCTION_HOUSE.settleCurrentAndCreateNewAuction();
        }

        /// @dev Prevent the user from bidding on an auction that they
        ///      have already won / are winning.
        if ($winner == msg.sender) {
            revert PlugNounsLib.InsufficientReason();
        }

        /// @dev Make sure the Socket has enough money to bid.
        if (msg.sender.balance < $bid) {
            revert PlugNounsLib.InsufficientBalance();
        }
    }

    function decode(bytes calldata $live) public pure returns (uint256 $bid) {
        /// @dev Decode the live data.
        ($bid) = abi.decode($live, (uint256));
    }

    function encode(uint256 $bid) public pure returns (bytes memory $live) {
        /// @dev Encode the live data.
        $live = abi.encode($bid);
    }
}
