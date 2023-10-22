//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {CaveatEnforcer} from '../CaveatEnforcer.sol';
import {BytesLib} from 'solidity-bytes-utils/contracts/BytesLib.sol';

abstract contract ThresholdEnforcer is CaveatEnforcer {
	using BytesLib for bytes;

	function _threshold() internal view virtual returns (uint256);

	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata,
		bytes32
	) public view override returns (bool $success) {
		/// @dev Retrieve the logic operator set by the Delegator in the terms.
		uint256 logicOperator = $terms.toUint128(0);
		/// @dev Move 16 bytes to the right to get the threshold.
		uint256 blockThreshold = $terms.toUint128(16);

		/// @dev Make sure the block number is before the threshold.
		if (logicOperator == 0) {
			if (blockThreshold <= _threshold())
				revert('BlockNumberBeforeEnforcer:expired-delegation');
		}
		/// @dev Make sure the block number is after the threshold.
		else if (blockThreshold >= _threshold())
			revert('BlockNumberAfterEnforcer:early-delegation');

		$success = true;
	}
}
