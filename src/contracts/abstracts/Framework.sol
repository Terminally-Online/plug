// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {IPlug} from '../interfaces/IPlug.sol';
import {PlugCore} from './PlugCore.sol';

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable permission of extremely
 *         granular permission and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract Plug is IPlug, PlugCore {
	/**
	 * @notice Instantiates a new Plug contract.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) PlugCore($name, $version) {}

	/**
	 * See {IPlug-contractInvoke}.
	 */
	function contractInvoke(
		Intent[] calldata $intent
	) external returns (bool $success) {
		$success = _invoke($intent, msg.sender);
	}

	/**
	 * See {IPlug-invoke}.
	 */
	function invoke(
		SignedIntents[] calldata $signedIntents
	) external returns (bool $success) {
		/// @dev Load the stack.
		uint256 i;

		/// @dev Loop through the signed intents.
		for (i; i < $signedIntents.length; ) {
			/// @dev Load the signed intent as a hot reference.
			SignedIntents calldata signedIntent = $signedIntents[i];

			/// @dev Determine who signed the intent.
			address intentSigner = getSignedIntentsSigner(signedIntent);

			/// @dev Load the intents as a hot reference.
			Intents calldata intents = signedIntent.intents;

			/// @dev Prevent replay attacks by enforcing replay protection.
			_enforceReplayProtection(intentSigner, intents.replayProtection);

			/// @dev Invoke the intents.
			$success = _invoke(intents.batch, intentSigner);

			unchecked {
				++i;
			}
		}
	}
}
