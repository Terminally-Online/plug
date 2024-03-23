// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugThresholdFuseEnforce } from
    "../abstracts/fuses/Plug.Threshold.Fuse.Enforce.sol";
import { PlugLib, PlugTypesLib } from "../libraries/Plug.Lib.sol";
import { ERC1155 } from "solady/src/tokens/ERC1155.sol";

/**
 * @title Plug Balance Semi Fungible Fuse
 * @dev A fuse that provides enforcement for semi-fungible balance thresholds.
 * @author nftchance (chance@onplug.io)
 */
contract PlugBalanceSemiFungibleFuse is
    PlugFuseInterface,
    PlugThresholdFuseEnforce
{
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
        returns (bytes memory $through)
    {
        /// @dev Determine the balance lookup definition.
        (
            address $holder,
            address $asset,
            uint256 $tokenId,
            uint8 $operator,
            uint256 $threshold
        ) = decode($live);

        /// @dev Memory reference for the balance of the holder.
        uint256 balance = ERC1155($asset).balanceOf($holder, $tokenId);

        _enforceFuse($operator, $threshold, balance);

        /// @dev Otherwise, return the current value.
        $through = $current.data;
    }

    /**
     * See { PlugFuseInterface-decode }.
     */
    function decode(bytes calldata $data)
        public
        pure
        returns (
            address $holder,
            address $asset,
            uint256 $tokenId,
            uint8 $operator,
            uint256 $threshold
        )
    {
        ($holder, $asset, $tokenId, $operator, $threshold) =
            abi.decode($data, (address, address, uint256, uint8, uint256));
    }

    /**
     * See { PlugFuseInterface-encode }.
     */
    function encode(
        address $holder,
        address $asset,
        uint256 $tokenId,
        uint8 $operator,
        uint256 $threshold
    )
        public
        pure
        returns (bytes memory $data)
    {
        $data = abi.encode($holder, $asset, $tokenId, $operator, $threshold);
    }
}
