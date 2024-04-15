// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../../interfaces/Plug.Connector.Interface.sol";
import { PlugNounsLib } from "../../libraries/protocols/Plug.Nouns.Lib.sol";

/**
 * @title Plug Nouns Id Fuse
 * @notice This Fuse enables the ability to declare a specific Noun tokenId
 *		   that you you would like to bid on, on a regular basis.
 * @notice Use cases for bidding on Nouns:
 *     - Schedule bids for specific Nouns.
 * @author nftchance (chance@onplug.io)
 */
contract PlugNounsId is PlugConnectorInterface {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view {
        /// @dev Get the current state of the auction.
        (uint256 nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();

        require(nounId == decode($terms), "NounsTokenId:invalid-noun-id");
    }

    /**
     * @notice Decode the live wire into the tokenId.
     * @param $live The live wire to decode.
     * @return $tokenId The id of the token being requested.
     */
    function decode(bytes calldata $live) public view virtual returns (uint256 $tokenId) {
        return abi.decode($live, (uint256));
    }

    /**
     * @notice Encode the tokenId into the live wire.
     * @param $value The tokenId to encode.
     * @return $data The encoded token id of the noun being bid on.
     */
    function encode(uint256 $value) public pure virtual returns (bytes memory $data) {
        $data = abi.encode($value);
    }
}
