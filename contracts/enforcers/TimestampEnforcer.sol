//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {ThresholdEnforcer} from '../abstracts/enforcers/ThresholdEnforcer.sol';

contract TimestampEnforcer is ThresholdEnforcer {
	/// @dev Returns the current timestamp.
	function _threshold() internal view override returns (uint256) {
		return block.timestamp;
	}
}
