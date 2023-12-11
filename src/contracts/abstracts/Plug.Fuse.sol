//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {IFuse} from '../interfaces/IFuse.sol';

abstract contract PlugFuse is IFuse {
	/**
	 * See {IFuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata $live,
		PlugTypesLib.Current calldata $current,
		bytes32 $pinHash
	) public virtual returns (bytes memory $callback);
}
