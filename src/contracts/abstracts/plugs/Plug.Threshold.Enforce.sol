// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugLib } from "../../libraries/Plug.Lib.sol";

abstract contract PlugThresholdEnforce {
    function _enforce(uint8 $operator, uint256 $threshold, uint256 $denominator) internal pure {
        /// @dev Make sure the base denominator is below (or before) the threshold.
        if ($operator == 0) {
            if ($threshold < $denominator) {
                revert PlugLib.ThresholdExceeded($threshold, $denominator);
            }
        }
        /// @dev Make sure the base denominator is above (or after) after the threshold.
        else if ($threshold > $denominator) {
            revert PlugLib.ThresholdInsufficient($threshold, $denominator);
        }
    }
}
