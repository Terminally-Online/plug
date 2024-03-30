// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import {
    PlugFuseInterface,
    PlugTypesLib
} from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugThresholdFuseEnforce } from
    "../abstracts/fuses/Plug.Threshold.Fuse.Enforce.sol";

import { PlugBalanceInterface } from
    "../interfaces/Plug.Balance.Interface.sol";

/**
 * @title Plug Balance (Semi-Fungible) Fuse
 * @notice A Fuse that provides enforcement of semi-fungible (ERC1155s) balance thresholds.
 * @notice Use cases for enforcing balance thresholds:
 *     - Inherits all the use cases of the fungible and non-fungible balance threshold fuse.
 *     - Tier based access and services resolved through the token id balance held.
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

        /// @dev Ensure the balance of the 1155 token is within the bounds defined.
        _enforceFuse(
            $operator,
            $threshold,
            PlugBalanceInterface($asset).balanceOf($holder, $tokenId)
        );

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
        ($holder, $asset, $tokenId, $operator, $threshold) = abi
            .decode($data, (address, address, uint256, uint8, uint256));
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
        $data = abi.encode(
            $holder, $asset, $tokenId, $operator, $threshold
        );
    }
}
