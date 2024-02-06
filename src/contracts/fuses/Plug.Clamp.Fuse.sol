// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { BytesLib } from "../libraries/BytesLib.sol";

/**
 * @title Plug Clamp Fuse
 * @dev A fuse that clamps the current value between two bounds.
 * @author nftchance (chance@utc24.io)
 *
 */
contract PlugClampFuse is PlugFuseInterface {
    using BytesLib for bytes;

    /**
     * See {FuseEnforcer-enforceFuse}.
     */
    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32
    )
        public
        pure
        override
        returns (bytes memory $through)
    {
        /// @dev Determine the bounds of the clamp.
        (uint256 $min, uint256 $max) = decode($live);

        /// @dev Make sure the current value is within the bounds.
        uint256 value = $current.data.toUint256(0);

        /// @dev If it is above the max, return the max.
        if (value > $max) return abi.encodePacked($max);
        /// @dev If it is below the min, return the min.
        if (value < $min) return abi.encodePacked($min);
        /// @dev Otherwise, return the current value.
        return $current.data;
    }

    /**
     * @dev Decode the clamp data into the two bounds.
     */
    function decode(bytes calldata $data)
        public
        pure
        returns (uint256 $min, uint256 $max)
    {
        ($min, $max) = abi.decode($data, (uint256, uint256));
    }

    /**
     * @dev Encode the clamp bounds.
     */
    function encode(
        uint256 $min,
        uint256 $max
    )
        public
        pure
        returns (bytes memory $data)
    {
        $data = abi.encode($min, $max);
    }
}
