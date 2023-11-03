//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {ICaveatEnforcer} from '../interfaces/ICaveatEnforcer.sol';

abstract contract CaveatEnforcer is ICaveatEnforcer {
	/**
	 * See {ICaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata $terms,
		Transaction calldata $transaction,
		bytes32 $permissionHash
	) public virtual returns (bool);
}
