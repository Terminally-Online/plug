//SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {CaveatEnforcer} from '../abstracts/CaveatEnforcer.sol';
import {BytesLib} from 'solidity-bytes-utils/contracts/BytesLib.sol';

contract EIP1271Enforcer is CaveatEnforcer {
	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata terms,
		Transaction calldata transaction,
		bytes32 delegationHash
	) public override returns (bool) {
		return true;
	}

	function _callERC1271isValidSignature(
		address _addr,
		bytes32 _hash,
		bytes calldata _signature
	) internal view {
		bytes4 result = IERC1271Wallet(_addr).isValidSignature(
			_hash,
			_signature
		);

		require(result == 0x1626ba7e, 'INVALID_SIGNATURE');
	}
}
