// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugFuse} from '../abstracts/Plug.Fuse.sol';
import {PlugTypesLib} from '../abstracts/Plug.Types.sol';

contract PlugAllowedMethodsFuse is PlugFuse {
	/**
	 * See {FuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata $live,
		PlugTypesLib.Current calldata $current,
		bytes32
	) public pure override returns (bytes memory $callback) {
		/// @dev The signature of the function that is being called.
		bytes4 targetSig = bytes4($current.data[0:4]);

		/// @dev Load the stack.
		uint256 i;

		for (i; i < $live.length; ) {
			/// @dev Slice the next 4 bytes from the terms array.
			bytes4 allowedSig = bytes4($live[i:i + 4]);
			/// @dev If we have a match, return true.
			if (allowedSig == targetSig) return bytes('');

			/// @dev Go to the next 4 bytes.
			unchecked {
				i += 4;
			}
		}

		revert('AllowedMethodsEnforcer:method-not-allowed');
	}

	/**
	 * @dev Decode the terms to get a specific signature.
	 */
	function decode(
		bytes calldata $data
	) public pure returns (bytes4[] memory) {
		/// @dev Load the stack.
		uint256 i;
		uint256 length = $data.length;

		/// @dev Create a new array to store the signatures.
		bytes4[] memory $signatures = new bytes4[](length / 4);

		for (i; i < length; ) {
			/// @dev Slice the next 4 bytes from the terms array.
			bytes4 signature = bytes4($data[i:i + 4]);
			/// @dev Push the signature to the array.
			$signatures[i / 4] = signature;

			/// @dev Go to the next 4 bytes.
			unchecked {
				i += 4;
			}
		}

		return $signatures;
	}

	/**
	 * @dev Encode all of the signatures that are allowed.
	 */
	function encode(
		bytes4[] memory $signatures
	) public pure returns (bytes memory $signature) {
		/// @dev Load the stack.
		uint256 i;
		uint256 length = $signatures.length;

		for (i; i < length; ) {
			$signature = abi.encodePacked($signature, $signatures[i]);

			unchecked {
				i++;
			}
		}
	}
}
