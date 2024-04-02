// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { ThresholdFuse } from
    "../abstracts/fuses/Plug.Threshold.Fuse.sol";

/**
 * @title Plug Timestamp Fuse
 * @notice A Fuse that enables thresholds on the current block number.
 * @notice Use cases for enforcing timestamp thresholds:
 *     - Automatic transaction execution at a future timestamp such as subscription
 *       transactions or other forms such as option expirations, bond maturity dates, etc.
 *     - Prevent the execution of transactions if they take too long to land onchain or
 *       see an attempt of usage after the intent has expired.
 *     - Transparently update the state of a sale or auction based on the current block number
 *       with definition and open access before taking funds.
 *     - Re-up the emissions of a protocol on a regular and clear basis.
 * @author nftchance (chance@onplug.io)
 */
contract PlugTimestampFuse is ThresholdFuse {
    /// @dev Returns the current timestamp.
    function _threshold() internal view override returns (uint256) {
        return block.timestamp;
    }
}
