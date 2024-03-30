// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import {
    PlugFuseInterface,
    PlugTypesLib
} from "../../interfaces/Plug.Fuse.Interface.sol";
import { PlugNounsLib } from
    "../../libraries/protocols/Plug.Nouns.Lib.sol";

/**
 * @title Plug Nouns Id Fuse
 * @notice This Fuse enables the ability to declare a specific Noun tokenId
 *		   that you you would like to bid on, on a regular basis.
 * @notice Use cases for bidding on Nouns:
 *     - Schedule bids for specific Nouns.
 * @author nftchance (chance@onplug.io)
 */
contract PlugNounsIdFuse is PlugFuseInterface {
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
        /// @dev Get the current state of the auction.
        (uint256 nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();

        require(
            nounId == decode($live), "NounsTokenId:invalid-noun-id"
        );

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @notice Decode the live wire into the tokenId.
     * @param $live The live wire to decode.
     * @return $tokenId The id of the token being requested.
     */
    function decode(bytes calldata $live)
        public
        view
        virtual
        returns (uint256 $tokenId)
    {
        return abi.decode($live, (uint256));
    }

    /**
     * @notice Encode the tokenId into the live wire.
     * @param $value The tokenId to encode.
     * @return $data The encoded token id of the noun being bid on.
     */
    function encode(uint256 $value)
        public
        pure
        virtual
        returns (bytes memory $data)
    {
        $data = abi.encode($value);
    }
}
