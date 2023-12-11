//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {ThresholdFuse} from '../abstracts/fuses/Plug.Threshold.Fuse.sol';

contract PlugTimestampFuse is ThresholdFuse {
	/// @dev Returns the current timestamp.
	function _threshold() internal view override returns (uint256) {
		return block.timestamp;
	}
}
