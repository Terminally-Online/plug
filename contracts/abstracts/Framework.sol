// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.19;

import {IFramework} from '../interfaces/IFramework.sol';
import {FrameworkCore} from './FrameworkCore.sol';

/**
 * @title Framework
 * @notice The core contract for the Emporium framework that enables
 *         counterfactual revokable delegation of extremely
 *         granular permission and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract Framework is IFramework, FrameworkCore {
	/**
	 * @notice Instantiates a new Framework contract.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) FrameworkCore($name, $version) {}

	/**
	 * See {IFramework-contractInvoke}.
	 */
	function contractInvoke(
		Invocation[] calldata $invocation
	) external returns (bool $success) {
		$success = _invoke($invocation, msg.sender);
	}

	/**
	 * See {IFramework-invoke}.
	 */
	function invoke(
		SignedInvocations[] calldata $signedInvocations
	) external returns (bool $success) {
		/// @dev Load the stack.
		uint256 i;

		/// @dev Loop through the signed invocations.
		for (i; i < $signedInvocations.length; ) {
			/// @dev Load the signed invocation as a hot reference.
			SignedInvocations calldata signedInvocation = $signedInvocations[i];

			/// @dev Determine who signed the invocation.
			address invocationSigner = getSigner(signedInvocation);

			/// @dev Load the invocations as a hot reference.
			Invocations calldata invocations = signedInvocation.invocations;

			/// @dev Prevent replay attacks by enforcing replay protection.
			_enforceReplayProtection(
				invocationSigner,
				invocations.replayProtection
			);

			/// @dev Invoke the invocations.
			$success = _invoke(invocations.batch, invocationSigner);

			unchecked {
				++i;
			}
		}
	}
}
