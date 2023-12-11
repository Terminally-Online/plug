//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugFuse} from '../abstracts/Plug.Fuse.sol';
import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {BytesLib} from '../libraries/BytesLib.sol';

/**
 * @title LimitedCallsEnforcer
 * @notice This Fuse Enforcer powers the ability to limit the number of times
 *         a delegate can call a function with the same pin hash.
 */
contract PlugLimitedCallsFuse is PlugFuse {
	/// @dev Use the BytesLib library for bytes manipulation.
	using BytesLib for bytes;

	/// @dev Keep track of how many times a pin has been used.
	mapping(address => mapping(bytes32 => uint256)) callCounts;

	/**
	 * See {FuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata $live,
		PlugTypesLib.Current calldata,
		bytes32 $pinHash
	) public override returns (bytes memory $callback) {
		/// @dev Confirm the allowed limit has not yet been reached by the sender
		///      of the declared pin.
		require(
			decode($live) >= callCounts[msg.sender][$pinHash]++,
			'LimitedCallsEnforcer:limit-exceeded'
		);

		$callback = bytes('');
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
	 * @dev  Encode the limit into the terms of the Fuse.
	 */
	function encode(
		uint256 $callCount
	) public pure returns (bytes memory $terms) {
		$terms = abi.encode($callCount);
	}
}
