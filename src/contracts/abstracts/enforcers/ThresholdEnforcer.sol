//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {CaveatEnforcer} from '../CaveatEnforcer.sol';
import {BytesLib} from '../../libraries/BytesLib.sol';

abstract contract ThresholdEnforcer is CaveatEnforcer {
	using BytesLib for bytes;

	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata,
		bytes32
	) public view override returns (bool $success) {
		/// @dev Decode the terms to get the logic operator and threshold.
		(uint256 $operator, uint256 $threshold) = decode($terms);

		/// @dev Make sure the block number is before the threshold.
		if ($operator == 0) {
			if ($threshold <= _threshold())
				revert('BlockNumberBeforeEnforcer:expired-permission');
		}
		/// @dev Make sure the block number is after the threshold.
		else if ($threshold >= _threshold())
			revert('BlockNumberAfterEnforcer:early-permission');

		$success = true;
	}

	/**
	 * @dev Decode the terms to get the logic operator and threshold.
	 */
	function decode(
		bytes calldata $data
	) public pure returns (uint128 $operator, uint128 $threshold) {
		/// @dev Retrieve the logic operator set in the terms.
		$operator = $data.toUint128(0);
		/// @dev Move 16 bytes to the right to get the threshold.
		$threshold = $data.toUint128(16);
	}

	/**
	 * @dev Encode the logic operator and threshold.
	 */
	function encode(
		uint128 $operator,
		uint128 $threshold
	) public pure returns (bytes memory $data) {
		/// @dev Encode the logic operator and threshold.
		$data = abi.encodePacked($operator, $threshold);
	}

	/**
	 * @dev Unit denomination of the threshold.
	 */
	function _threshold() internal view virtual returns (uint256);
}
