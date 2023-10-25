//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Shape declarations in the Emporium framework.
import {ITypes} from '../abstracts/Types.sol';

interface IFramework is ITypes {
	/**
	 * @notice Allows a smart contract to submit a batch of invocations for processing,
	 *         allowing itself to be the delegate.
	 * @param $invocation The batch of invocations to process.
	 * @return success Whether the batch of invocations was successfully processed.
	 */
	function contractInvoke(
		Invocation[] calldata $invocation
	) external returns (bool);

	/**
	 * @notice Allows anyone to submit a batch of signed invocations for processing.
	 * @param $signedInvocations The batch of signed invocations to process.
	 * @return success Whether the batch of invocations was successfully processed.
	 */
	function invoke(
		SignedInvocations[] calldata $signedInvocations
	) external returns (bool success);
}
