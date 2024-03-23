// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { ThresholdFuse } from
    "../abstracts/fuses/Plug.Threshold.Fuse.sol";

contract PlugBlockNumberFuse is ThresholdFuse {
    /// @dev Returns the current block number.
    function _threshold() internal view override returns (uint256) {
        return block.number;
    }
}
