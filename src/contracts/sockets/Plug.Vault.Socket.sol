// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';
import {Receiver} from 'solady/src/accounts/Receiver.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';
import {LibZip} from 'solady/src/utils/LibZip.sol';

contract PlugVaultSocket is PlugSocket, Receiver, Ownable {
	bool private initialized;

	constructor() {
		/// @dev Initialize the owner.
		initialize(msg.sender);
	}

	modifier initializer() {
		require(!initialized, 'PlugVaultSocket: already initialized');
		_;
		initialized = true;
	}

	/**
	 * @notice Initialize a new Plug Vault.
	 * @param $owner The owner of the vault.
	 */
	function initialize(address $owner) public payable virtual initializer {
		/// @dev Initialize the owner.
		_initializeOwner($owner);

		/// @dev Initialize the Plug Socket.
		_initializeSocket('PlugVaultSocket', '0.0.0');
	}
}
