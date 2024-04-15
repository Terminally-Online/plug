// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../../interfaces/Plug.Connector.Interface.sol";
import { PlugThresholdEnforce } from "../../abstracts/plugs/Plug.Threshold.Enforce.sol";

import { FraxlendVaultInterface } from "../../interfaces/protocols/Fraxlend.Vault.Interface.sol";

/**
 * @title Plug Fraxlend APY
 * @notice A Plug that provides enforcement of APY thresholds Flexlend Vaults.
 * @notice Use cases for enforcing APY thresholds:
 *     - Borrow significant amounts of assets from a Fraxlend Vault when
 *       the APY is below a certain threshold.
 *     - Lend significant amounts of assets to a Fraxlend Vault when the
 *       APY is above a certain threshold.
 *     - Withdraw assets from Fraxlend Vault when the APY falls below a
 *       declared threshold
 * @author nftchance (chance@onplug.io)
 */
contract PlugFraxlendAPY is PlugConnectorInterface, PlugThresholdEnforce {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view {
        /// @dev Determine the balance lookup definition.
        (address $vault, uint8 $vaultOperator, uint8 $operator, uint256 $threshold) = decode($terms);

        /// @dev Connect to the vault and retrieve the current rate information.
        FraxlendVaultInterface vault = FraxlendVaultInterface($vault);

        /// @dev Retrieve the current rate information from the Vault being
        ///      entered from within the Frax Vault.
        (,,, uint256 ratePerSec) = vault.currentRateInfo();

        /// @dev Retrieve the amount of assets that are active in the Vault.
        (uint128 totalAssetAmount,) = vault.totalAsset();
        (uint128 totalBorrowAmount,) = vault.totalBorrow();

        /// @dev Solve for the base borrow.
        uint256 apyBaseBorrow = ((ratePerSec * 365 days) / 10 ** 18) / 100;

        /// @dev Determine the amount of apy boosting that is taking place.
        uint256 apyRewardBorrow = (apyBaseBorrow * totalBorrowAmount) / totalAssetAmount;

        /// @dev Account for the APY if it is a lend or borrow action.
        uint256 apy = $vaultOperator == 0 ? apyRewardBorrow : apyBaseBorrow + apyRewardBorrow;

        /// @dev Enforce that the APY is within bounds.
        _enforce($operator, $threshold, apy);
    }

    /**
     * See { PlugConnectorInterface-decode }.
     */
    function decode(bytes calldata $data)
        public
        pure
        returns (address $vault, uint8 $vaultOperator, uint8 $operator, uint256 $threshold)
    {
        ($vault, $vaultOperator, $operator, $threshold) =
            abi.decode($data, (address, uint8, uint8, uint256));
    }

    /**
     * See { PlugConnectorInterface-encode }.
     */
    function encode(
        address $vault,
        uint8 $vaultOperator,
        uint8 $operator,
        uint256 $threshold
    )
        public
        pure
        returns (bytes memory $data)
    {
        /// @dev Encode the holder, asset, operator, and threshold.
        $data = abi.encode($vault, $vaultOperator, $operator, $threshold);
    }
}
