//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {ERC20} from 'solady/src/tokens/ERC20.sol';

import {PlugFuse} from '../abstracts/Plug.Fuse.sol';
import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {BytesLib} from '../libraries/BytesLib.sol';

contract PlugERC20AllowanceFuse is PlugFuse {
	using BytesLib for bytes;

	/// @dev Function signature of the ERC20 `transfer` method.
	///      | transfer | a9059cbb | transfer(address,uint256) |
	bytes4 public constant ERC20_TRANSFER_FROM =
		bytes4(keccak256('transfer(address,uint256)'));

	/// @dev Balance of amount spend per pin hash.
	mapping(address => mapping(bytes32 => uint256)) spentMap;

	/**
	 * See {FuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata $live,
		PlugTypesLib.Current calldata $current,
		bytes32 $pinHash
	) public override returns (bytes memory $callback) {
		/// @dev Determine the function being called by the transaction.
		bytes4 targetSig = bytes4($current.data[0:4]);

		/// @dev Ensure the method being called is `transferFrom`.
		require(
			targetSig == ERC20_TRANSFER_FROM,
			'ERC20AllowanceEnforcer:invalid-method'
		);

		/// @dev Retrieve the limit set by the Delegator.
		uint256 limit = decode($live);
		/// @dev Retrieve the amount of tokens being transferred and starts
		///      at the 36th byte of the transaction data because the first
		///      4 bytes are the function signature and the next 32 bytes are
		///      the address of the token being transferred.
		uint256 sending = BytesLib.toUint256($current.data, 36);

		/// @dev Adjust the spent amount for the pin hash.
		spentMap[msg.sender][$pinHash] += sending;

		/// @dev Retrieve the balance of the Delegator.
		uint256 spent = spentMap[msg.sender][$pinHash];

		/// @dev Make sure amount spent will not exceed the limit.
		require(spent <= limit, 'ERC20AllowanceEnforcer:allowance-exceeded');

		$callback = bytes('');
	}

	/**
	 * @dev Decode the limit defined by the terms at a given bytes index.
	 */
	function decode(
		bytes calldata $terms
	) public pure returns (uint256 $limit) {
		/// @dev Decode the limit.
		$limit = BytesLib.toUint256($terms, 0);
	}

	/**
	 * @dev  Encode the limit into the terms of the Fuse.
	 */
	function encode(uint256 $limit) public pure returns (bytes memory $terms) {
		/// @dev Encode the limit.
		$terms = abi.encodePacked($limit);
	}
}
