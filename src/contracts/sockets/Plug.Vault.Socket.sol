// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';

contract PlugVaultSocket is PlugSocket, Ownable {
	/**
	 * @notice Initialize a new Plug Vault.
	 * @param $owner The owner of the vault.
	 * @param $name The name of the vault used in Domain.
	 * @param $version The version of the vault used in Domain.
	 */
	function _initializeVault(
		address $owner,
		string memory $name,
		string memory $version
	) internal virtual {
		/// @dev Initialize the owner.
		_initializeOwner($owner);

		/// @dev Initialize the Plug Socket.
		_initializeSocket($name, $version);
	}

	/**
	 * @notice Execute through the transaction on behalf of the owner.
	 * @param $to The address to execute the transaction on.
	 * @param $data The data to execute the transaction with.
	 * @param $value The value to execute the transaction with.
	 * @return $success The success of the transaction.
	 * @return $returnData The data returned from the transaction.
	 */
	function execute(
		address $to,
		bytes calldata $data,
		uint256 $value
	)
		public
		payable
		virtual
		onlyOwner
		returns (bool $success, bytes memory $returnData)
	{
		/// @dev Execute the transaction and bubble up the response.
		($success, $returnData) = $to.call{value: $value}($data);
	}
}
