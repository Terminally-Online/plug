// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugNounsLib } from "../libraries/Plug.Nouns.Lib.sol";
import { BytesLib } from "../libraries/BytesLib.sol";

/**
 * @title Nouns Id Fuse
 * @notice This Fuse enables the ability to declare a specific Noun tokenId
 *		   that you you would like to bid on, on a regular basis.
 * @author @nftchance <chance@utc24.io>
 */
contract PlugNounsIdFuse is PlugFuseInterface {
    using BytesLib for bytes;

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
        /// @dev Get the current state of the auction.
        (uint256 nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();

        require(nounId == decode($live), "NounsTokenId:invalid-noun-id");

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @notice Decode the live wire into the tokenId.
     * @param $live The live wire to decode.
     * @return The tokenId being requested.
     */
    function decode(bytes calldata $live)
        public
        view
        virtual
        returns (uint256)
    {
        return $live.toUint256(0);
    }

    /**
     * @notice Encode the tokenId into the live wire.
     * @param $value The tokenId to encode.
     * @return The live wire.
     */
    function encode(uint256 $value)
        public
        pure
        virtual
        returns (bytes memory)
    {
        return abi.encode($value);
    }
}
