//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {IFuse} from '../interfaces/IFuse.sol';

abstract contract Fuse is IFuse {
	/**
	 * See {IFuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata $live,
		Current calldata $current,
		bytes32 $pinHash
	) public virtual returns (bool);
}
