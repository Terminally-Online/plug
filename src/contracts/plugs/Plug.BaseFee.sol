// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugThreshold } from "../abstracts/plugs/Plug.Threshold.sol";

/**
 * @title Plug Base Fee
 * @notice A Plug that provides enforcement of thresholds on the gas base fee.
 * @notice Use cases for enforcing balance thresholds:
 *     - Limit the economic value that can be used to prioritize transactions.
 *     - Prevent transactions from being executed when the volatility of the base fee
 *       results in exceeding the declared threshold.
 *     - Automatically run a transaction when the base fee is low and a typically
 *       expensive operation can be executed without such a large fee.
 *     - Provide predictability (and bounds) to Solvers and Paymasters that route
 *       transactions to onchain on behalf of a user.
 * @author nftchance (chance@onplug.io)
 */
contract PlugBaseFee is PlugThreshold {
    /// @dev Returns the current block basefee.
    function _threshold() internal view override returns (uint256) {
        return block.basefee;
    }
}
