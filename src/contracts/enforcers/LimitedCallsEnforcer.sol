//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {CaveatEnforcer} from '../abstracts/CaveatEnforcer.sol';
import {BytesLib} from '../libraries/BytesLib.sol';

/**
 * @title LimitedCallsEnforcer
 * @notice This Caveat Enforcer powers the ability to limit the number of times
 *         a delegate can call a function with the same permission hash.
 */
contract LimitedCallsEnforcer is CaveatEnforcer {
	/// @dev Use the BytesLib library for bytes manipulation.
	using BytesLib for bytes;

	/// @dev Keep track of how many times a permission has been used.
	mapping(address => mapping(bytes32 => uint256)) callCounts;

	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata,
		bytes32 $permissionHash
	) public override returns (bool $success) {
		/// @dev Confirm the allowed limit has not yet been reached by the sender
		///      of the declared permission.
		require(
			decode($terms) >= callCounts[msg.sender][$permissionHash]++,
			'LimitedCallsEnforcer:limit-exceeded'
		);

		$success = true;
	}

	/**
	 * @dev Decode the callCount defined by the terms at a given bytes index.
	 */
	function decode(
		bytes calldata $terms
	) public pure returns (uint256 $callCount) {
		$callCount = $terms.toUint256(0);
	}

	/**
	 * @dev  Encode the limit into the terms of the Caveat.
	 */
	function encode(
		uint256 $callCount
	) public pure returns (bytes memory $terms) {
		$terms = abi.encode($callCount);
	}
}
