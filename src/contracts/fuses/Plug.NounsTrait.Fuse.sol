// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Ownable } from "solady/src/auth/Ownable.sol";

import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugNounsLib } from "../libraries/Plug.Nouns.Lib.sol";

/**
 * @title Nouns Trait Fuse
 * @notice This Fuse enables the ability to declare a specific trait that you
 *         you would like to bid on, on a regular basis. To be very explicit with
 *         this model, a hash of the returned trait metadata is used to verify
 *         to minimize the amount of data that needs to be relayed.
 * @dev The verification of this takes place through onchain means
 *      therefore it can be rather expensive.
 * @author @nftchance <chance@utc24.io>
 * @author @masonchain
 */
contract PlugNounsTraitFuse is PlugFuseInterface, Ownable {
    /// @dev Function hashes of the trait getters.
    bytes32 public constant BACKGROUND_SELECTOR =
        keccak256(abi.encode("background(uint256 index)"));
    bytes32 public constant HEAD_SELECTOR =
        keccak256(abi.encode("head(uint256 index)"));
    bytes32 public constant GLASSES_SELECTOR =
        keccak256(abi.encode("glasses(uint256 index)"));
    bytes32 public constant BODY_SELECTOR =
        keccak256(abi.encode("body(uint256 index)"));
    bytes32 public constant ACCESSORY_SELECTOR =
        keccak256(abi.encode("accessory(uint256 index)"));

    /// @dev Metadata storage contract for Nouns.
    /// @notice We use a raw address instead of interface here because we are dynamically building
    ///         the function to be called by decoding the selector from the live wire.
    address art;

    constructor(address $art) {
        /// @dev Prepare the inferaces.
        // NOUN = $noun;
        // AUCTION_HOUSE = $auctionHouse;

        /// @dev Set the scope of the Fuse.
        art = $art;
    }

    /**
     * See {Fuse-enforceFuse}.
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
        (bytes32 selector, bytes32 trait) = decode($live);

        require(nounTrait(selector) == trait, "NounsTraitFuse:invalid-trait");

        /// @dev Continue the pass through.
        $through = $current.data;
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
        returns (bytes memory)
    {
        bool isValid;
        if ($selector == HEAD_SELECTOR) {
            isValid = true;
        } else if ($selector == GLASSES_SELECTOR) {
            isValid = true;
        } else if ($selector == BODY_SELECTOR) {
            isValid = true;
        } else if ($selector == ACCESSORY_SELECTOR) {
            isValid = true;
        } else if ($selector == BACKGROUND_SELECTOR) {
            isValid = true;
        } else {
            revert("NounsTraitFuse:invalid-selector");
        }

        return abi.encode($selector, $trait);
    }

    /**
     * @notice Retrieve the trait bytes for a given trait function selector and seed.
     * @param $selector The function selector of the trait being checked.
     * @return $traitHash The hash of the trait.
     */
    function nounTrait(bytes32 $selector)
        public
        view
        virtual
        returns (bytes32 $traitHash)
    {
        /// @dev Get the current state of the auction.
        (uint256 nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev Retrieve the metadata seeds of the current Noun.
        (
            uint48 background,
            uint48 body,
            uint48 accessory,
            uint48 head,
            uint48 glasses
        ) = PlugNounsLib.TOKEN.seeds(nounId);

        /// @dev Get the seed for the specified trait from the Seed struct.
        uint256 traitSeed;
        if ($selector == HEAD_SELECTOR) {
            traitSeed = head;
        } else if ($selector == GLASSES_SELECTOR) {
            traitSeed = glasses;
        } else if ($selector == BODY_SELECTOR) {
            traitSeed = body;
        } else if ($selector == ACCESSORY_SELECTOR) {
            traitSeed = accessory;
        } else if ($selector == BACKGROUND_SELECTOR) {
            traitSeed = background;
        } else {
            revert("NounsTraitFuse:invalid-selector");
        }

        /// @dev Build the transaction data to call the Noun contract.
        bytes memory data = abi.encodeWithSelector(bytes4($selector), traitSeed);

        /// @dev Retrieve the trait.
        (bool success, bytes memory returnData) = art.staticcall(data);

        /// @dev Ensure the call succeeded -- out of scope Nouns are not alive.
        require(success, "NounsTraitFuse:trait-call-failed");

        $traitHash = keccak256(returnData);
    }

    /**
     * @notice Set the address of the contract that supplies the art to Nouns.
     * @dev This is here just in case Nouns updates their metadata infrastructure
     *      as it has happened before (v2).
     * @param $art The address of the art contract.
     */
    function setArt(address $art) public onlyOwner {
        art = $art;
    }
}
