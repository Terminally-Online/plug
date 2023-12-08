// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';
import {ERC20} from 'solady/src/tokens/ERC20.sol';

/**
 * @title Plug ERC20 Socket.
 * @dev This ERC20 is designed to be a fair-release token where all addresses
 *      of Ethereum can claim a portion of the supply.
 * @notice A mock contract for testing the Plug framework.
 * @dev This contract is for testing purposes only.
 */
contract PlugERC20Socket is PlugSocket, ERC20 {
	/// @dev Store the metadata.
	string public tokenName;
	string public tokenSymbol;

	/// @dev Mapping of claimed addresses.
	mapping(address => bool) public addressToClaimed;

	/**
	 * @notice Initializes a new Plug contract.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) {
		/// @dev Initialize the Plug Socket.
		_initializeSocket($name, $version);

		/// @dev Initialize the ERC20.
		tokenName = $name;
		tokenSymbol = $version;
	}

	function name() public view override returns (string memory) {
		return tokenName;
	}

	function symbol() public view override returns (string memory) {
		return tokenSymbol;
	}

	function claim() public virtual { 
		/// @dev Ensure the address has not already claimed.
		require(!addressToClaimed[_msgSender()], 'PlugERC20Socket:claimed');

		/// @dev Mark the address as claimed.
		addressToClaimed[_msgSender()] = true;

		/// @dev Mint the token to the address.
		_mint(_msgSender(), 1);
	}
}
