// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.19;

import {Framework} from '../abstracts/Framework.sol';

/**
 * @title Framework Mock
 * @notice A mock contract for testing the Emporium framework.
 * @dev This contract is for testing purposes only.
 */
contract FrameworkMock is Framework {
	/// @dev Active revert when echo is muted.
	error EchoMuted();

	/// @dev Event emitted when an Echo is invoked.
	event EchoInvoked(address $reality, address $perception, string $message);

	/**
	 * @notice Instantiates a new Framework contract.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) Framework($name, $version) {}

	/**
	 * @notice A mock function for testing the framework.
	 * @param $message The message to echo.
	 */
	function echo(string memory $message) external {
		emit EchoInvoked(msg.sender, _msgSender(), $message);
	}

	function pureEcho() external pure returns (string memory $message) {
		$message = 'Hello World';
	}

	/**
	 * @notice A mock function for testing the framework.
	 */
	function echoMuted() external pure {
		revert EchoMuted();
	}
}
