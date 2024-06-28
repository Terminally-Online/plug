// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../../interfaces/Plug.Connector.Interface.sol";
import { PlugNounsLib, NounsArtInterface } from "../../libraries/protocols/nouns/Plug.Nouns.Lib.sol";

/**
 * @title Plug Nouns Trait
 * @notice This Plug enables the ability to declare a specific Noun tokenId
 *		   that you you would like to bid on, on a regular basis.
 * @notice Use cases for bidding on Nouns:
 *     - Schedule bids for specific Nouns traits.
 * @author nftchance (chance@onplug.io)
 * @author @masonchain
 */
contract PlugNounsTrait is PlugConnectorInterface {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view {
        /// @dev Recover the encoded trait data.
        (bytes32 selector, bytes32 trait) = decode($terms);

        /// @dev Ensure that the trait is the same as the one we are looking for.
        if (nounTrait(selector) != trait) {
            revert PlugNounsLib.InsufficientReason();
        }
    }

    /**
     * @notice Decode the live wire into the selector and trait hashes.
     * @param $live The live wire to decode.
     * @return $selector The selector of the trait to retrieve.
     * @return $trait The trait to retrieve.
     */
    function decode(bytes calldata $live)
        public
        view
        virtual
        returns (bytes32 $selector, bytes32 $trait)
    {
        ($selector, $trait) = abi.decode($live, (bytes32, bytes32));
    }

    /**
     * @notice Encode the selector and trait hashes into a live wire.
     * @param $selector The selector of the trait to retrieve.
     * @param $trait The trait to retrieve.
     * @return $live The live wire to decode.
     */
    function encode(
        bytes32 $selector,
        bytes32 $trait
    )
        public
        pure
        virtual
        returns (bytes memory $live)
    {
        if (
            $selector == PlugNounsLib.HEAD_SELECTOR || $selector == PlugNounsLib.GLASSES_SELECTOR
                || $selector == PlugNounsLib.BODY_SELECTOR
                || $selector == PlugNounsLib.ACCESSORY_SELECTOR
                || $selector == PlugNounsLib.BACKGROUND_SELECTOR
        ) { } else {
            revert PlugNounsLib.InvalidSelector($selector);
        }

        $live = abi.encode($selector, $trait);
    }

    /**
     * @notice Retrieve the trait bytes for a given trait function selector and seed.
     * @param $selector The function selector of the trait being checked.
     * @return $traitHash The hash of the trait.
     */
    function nounTrait(bytes32 $selector) public view virtual returns (bytes32 $traitHash) {
        /// @dev Get the current state of the auction.
        (uint256 nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev Retrieve the metadata seeds of the current Noun.
        (uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses) =
            PlugNounsLib.TOKEN.seeds(nounId);

        /// @dev Get the seed for the specified trait from the Seed struct.
        bytes memory trait;
        if ($selector == PlugNounsLib.HEAD_SELECTOR) {
            trait = PlugNounsLib.ART.heads(head);
        } else if ($selector == PlugNounsLib.GLASSES_SELECTOR) {
            trait = PlugNounsLib.ART.glasses(glasses);
        } else if ($selector == PlugNounsLib.BODY_SELECTOR) {
            trait = PlugNounsLib.ART.bodies(body);
        } else if ($selector == PlugNounsLib.ACCESSORY_SELECTOR) {
            trait = PlugNounsLib.ART.accessories(accessory);
        } else if ($selector == PlugNounsLib.BACKGROUND_SELECTOR) {
            /// @dev For some reason `backgrounds` returns a string
            ///      instead of bytes, so we need to convert it.
            trait = bytes(PlugNounsLib.ART.backgrounds(background));
        } else {
            revert PlugNounsLib.InvalidSelector($selector);
        }

        $traitHash = keccak256(trait);
    }
}
