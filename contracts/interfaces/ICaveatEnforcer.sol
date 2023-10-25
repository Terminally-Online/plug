//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Shape declarations in the Emporium framework.
import {ITypes} from '../abstracts/Types.sol';

interface ICaveatEnforcer is ITypes {
	/**
	 * @notice Enforces a caveat on a transaction.
	 * @param $terms The terms of the caveat.
	 * @param $transaction The transaction to enforce the caveat on.
	 * @param $delegationHash The hash of the delegation.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata $transaction,
		bytes32 $delegationHash
	) external returns (bool);
}
