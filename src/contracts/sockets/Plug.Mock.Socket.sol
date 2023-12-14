// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';

/**
 * @title Plug Mock
 * @notice A mock contract for testing the Plug framework.
 * @dev This contract is for testing purposes only.
 */
contract PlugMockSocket is PlugSocket {
	/// @dev Active revert when echo is muted.
	error EchoMuted();

	/// @dev Event emitted when an Echo is invoked.
	event EchoInvoked(address $reality, address $perception, string $message);

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
	}

	/**
	 * @notice A mock function for testing the framework.
	 * @param $message The message to echo.
	 */
	function echo(string memory $message) external {
		emit EchoInvoked(msg.sender, _msgSender(), $message);
	}

	function emptyEcho() external {
		emit EchoInvoked(msg.sender, _msgSender(), 'Hello World');
	}

	/**
	 * @notice A mock function for testing the framework.
	 */
	function mutedEcho(uint256 $echo) external pure returns (uint256 $slot) {
        if($echo % 8 == 0) {
            $slot = 1;
        }
        $slot = 2;
	}
}
