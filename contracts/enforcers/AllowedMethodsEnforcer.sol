// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {CaveatEnforcer} from '../abstracts/CaveatEnforcer.sol';

contract AllowedMethodsEnforcer is CaveatEnforcer {
	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata $transaction,
		bytes32
	) public pure override returns (bool) {
		/// @dev The signature of the function that is being called.
		bytes4 targetSig = bytes4($transaction.data[0:4]);

		/// @dev Load the stack.
		uint256 i;

		for (i; i < $terms.length; ) {
			/// @dev Slice the next 4 bytes from the terms array.
			bytes4 allowedSig = bytes4($terms[i:i + 4]);

			/// @dev If we have a match, return true.
			if (allowedSig == targetSig) return true;

			/// @dev Go to the next 4 bytes.
			unchecked {
				i += 4;
			}
		}

		revert('AllowedMethodsEnforcer:method-not-allowed');
	}
}
