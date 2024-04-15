// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../interfaces/Plug.Connector.Interface.sol";
import { PlugThresholdEnforce } from "../abstracts/plugs/Plug.Threshold.Enforce.sol";

import { PlugBalanceInterface } from "../interfaces/Plug.Balance.Interface.sol";

/**
 * @title Plug Balance (Semi-Fungible)
 * @notice A Plug that provides enforcement of semi-fungible (ERC1155s) balance thresholds.
 * @notice Use cases for enforcing balance thresholds:
 *     - Inherits all the use cases of the fungible and non-fungible balance threshold Plug.
 *     - Tier based access and services resolved through the token id balance held.
 * @author nftchance (chance@onplug.io)
 */
contract PlugBalanceSemiFungible is PlugConnectorInterface, PlugThresholdEnforce {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view {
        /// @dev Determine the balance lookup definition.
        (address $holder, address $asset, uint256 $tokenId, uint8 $operator, uint256 $threshold) =
            decode($terms);

        /// @dev Ensure the balance of the 1155 token is within the bounds defined.
        _enforce($operator, $threshold, PlugBalanceInterface($asset).balanceOf($holder, $tokenId));
    }

    /**
     * See { PlugConnectorInterface-decode }.
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
     * See { PlugConnectorInterface-encode }.
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
