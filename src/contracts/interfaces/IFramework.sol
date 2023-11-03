//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Shape declarations in the Emporium framework.
import {ITypes} from '../abstracts/Types.sol';

interface IFramework is ITypes {
	/**
	 * @notice Allows a smart contract to submit a batch of intents for processing,
	 *         allowing itself to be the delegate.
	 * @param $intent The batch of intents to execute.
	 * @return success Whether the batch of intent was successfully processed.
	 */
	function contractInvoke(Intent[] calldata $intent) external returns (bool);

	/**
	 * @notice Allows anyone to submit a batch of signed intents for processing.
	 * @param $signedIntents The batch of signed intents to process.
	 * @return success Whether the batch of intents was successfully processed.
	 */
	function invoke(
		SignedIntents[] calldata $signedIntents
	) external returns (bool success);
}
