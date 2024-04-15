// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../interfaces/Plug.Connector.Interface.sol";
import { PlugThresholdEnforce } from "../abstracts/plugs/Plug.Threshold.Enforce.sol";

import { PlugBalanceInterface } from "../interfaces/Plug.Balance.Interface.sol";

/**
 * @title Plug Balance
 * @notice A Plug that provides enforcement of balance thresholds for Native, ERC20,
 *         and ERC721 tokens, but not ERC1155s.
 * @notice Use cases for enforcing balance thresholds:
 *     - Limit the balance of a given asset that an account should hold.
 *     - Limit the amount of tokens that a given recipient should be sent.
 *     - Swap assets when a holder's balance follows below a threshold.
 *     - Swap assets when a holder's balance exceeds a threshold.
 *     - Allow a account to manage/use the state of an external contract and/or
 *       application by giving permission of interaction when holdings are present
 *       with a read-only interaction pattern (subscription services).
 *     - Systematically give and revoke permissions to other accounts with an
 *       onchain write interaction pattern.
 *     - Automatically adjust the collateralization of a loan or position based on
 *       the balance of collateral assets, reducing the risk of liquidiation.
 *     - Determine and set the fee structure for a given account based on the
 *       asset holdings of the targeted holder.
 *     - Airdrop eligibility based on the current balance of a given asset.
 *     - Automatically trigger donations when balance exceeds the declared threshold.
 * @author nftchance (chance@onplug.io)
 */
contract PlugBalance is PlugConnectorInterface, PlugThresholdEnforce {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view {
        /// @dev Determine the balance lookup definition.
        (address $holder, address $asset, uint8 $type, uint8 $operator, uint256 $threshold) =
            decode($terms);

        /// @dev If it is a native asset, ensure the balance is within bounds defined.
        if ($type == 0) {
            _enforce($operator, $threshold, $holder.balance);
        }
        /// @dev Otherwise, ensure the balance of an ERC20 or ERC721 token is within
        ///      specification of the intent.
        else {
            _enforce($operator, $threshold, PlugBalanceInterface($asset).balanceOf($holder));
        }
    }

    /**
     * See { PlugConnectorInterface-decode }.
     */
    function decode(bytes calldata $data)
        public
        pure
        returns (address $holder, address $asset, uint8 $type, uint8 $operator, uint256 $threshold)
    {
        ($holder, $asset, $type, $operator, $threshold) =
            abi.decode($data, (address, address, uint8, uint8, uint256));
    }

    /**
     * See { PlugConnectorInterface-encode }.
     */
    function encode(
        address $holder,
        address $asset,
        uint8 $type,
        uint8 $operator,
        uint256 $threshold
    )
        public
        pure
        returns (bytes memory $data)
    {
        /// @dev Encode the holder, asset, operator, and threshold.
        $data = abi.encode($holder, $asset, $type, $operator, $threshold);
    }
}
