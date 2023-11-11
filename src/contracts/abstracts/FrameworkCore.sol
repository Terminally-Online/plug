// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Hash declarations and decoders for the Plug framework.
import {Types} from './Types.sol';
/// @dev Error utilities for the Plug framework.
import {PlugErrors} from '../libraries/PlugErrors.sol';

/// @dev Core Plug dependencies.
import {CaveatEnforcer} from './CaveatEnforcer.sol';

/**
 * @title Plug Core
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable permission of extremely
 *         granular permission and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract PlugCore is Types {
	using PlugErrors for bytes;

	/// @notice Multi-dimensional account permission nonce management.
	mapping(address => mapping(uint256 => uint256)) public nonce;

	/**
	 * @notice Load the Core alongside all the Types driving
	 *         the parent consumer.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) Types($name, $version) {}

	/**
	 * @notice Determine the address representing the message sender in the
	 *         current context. This is important for permissions, as the
	 *         message sender may be the framework itself, in which case
	 *         the sender must be extracted from the data.
	 * @dev Because of this function, when building on top of the framework,
	 *      you should never use `msg.sender` directly, but instead use
	 *      `_msgSender()`, which will correctly identify the sender whether
	 *      it's the framework or an external account.
	 * @return $sender The address of the message sender.
	 */
	function _msgSender() internal view virtual returns (address $sender) {
		/// @dev If the message sender is the framework, we need to extract the
		///      sender from the data.
		if (msg.sender == address(this)) {
			/// @dev Load the data as a hot reference.
			bytes memory array = msg.data;

			/// @dev Load the length of the data as a hot reference.
			uint256 index = array.length;

			assembly {
				/// @dev Load the sender from the data by applying a
				///      bitwise AND operation to the data and the
				///      maximum uint256 value, keeping only the last
				///      20 bytes (160 bits) (address size).
				$sender := and(
					/// @dev Load the bytes at the computer pointer.
					mload(
						/// @dev Computes the sum of the starting address of array and index.
						///      This effectively points to the end of array.
						add(array, index)
					),
					0xffffffffffffffffffffffffffffffffffffffff
				)
			}
		}
		/// @dev Otherwise, the sender is the message sender.
		else $sender = msg.sender;
	}

	/**
	 * @notice Update the nonce for a given account and queue.
	 * @param $intendedSender The address of the intended sender.
	 * @param $protection The replay protection struct.
	 */
	function _enforceReplayProtection(
		address $intendedSender,
		ReplayProtection memory $protection
	) internal {
		/// @dev Ensure the nonce is in order.
		require(
			$protection.nonce == ++nonce[$intendedSender][$protection.queue],
			'PlugCore:nonce2-out-of-order'
		);
	}

	/**
	 * @notice Execution a built transaction.
	 * @param $to The address of the contract to execute.
	 * @param $data The data to execute on the contract.
	 * @param $gasLimit The gas limit for the transaction.
	 * @param $sender The address of the sender.
	 * @return $success Whether the transaction was successfully executed.
	 */
	function _execute(
		address $to,
		bytes memory $data,
		uint256 $gasLimit,
		address $sender
	) internal returns (bool $success) {
		/// @dev Build the final call data.
		bytes memory full = abi.encodePacked($data, $sender);

		/// @dev Warm up the slot for the return data.
		bytes memory errorMessage;

		/// @dev Make the external call that was delegated.
		($success, errorMessage) = address($to).call{gas: $gasLimit}(full);

		/// @dev If the call failed, bubble up the revert reason if possible.
		if ($success == false) errorMessage.bubbleRevert();
	}

	/**
	 * @notice Execute a batch of intents
	 * @param $batch The batch of intents to execute.
	 * @param $sender The address of the sender.
	 * @return $success Whether the transaction was successfully executed.
	 */
	function _invoke(
		Intent[] calldata $batch,
		address $sender
	) internal returns (bool $success) {
		/// @dev Load the stack.
		uint256 i;
		uint256 j;
		uint256 k;
		address canGrant;
		address intendedSender;
		address permissionSigner;
		bytes32 authHash;
		bytes32 permissionHash;

		/// @dev Load the structs into a hot reference.
		Intent memory intent;
		SignedPermission memory signedPermission;
		Permission memory permission;
		Transaction memory transaction;

		/// @dev Iterate over the batch of intents.
		for (i; i < $batch.length; ) {
			/// @dev Load the intent from the batch.
			intent = $batch[i];

			/// @dev If there are no permissions, this intent comes from the signer
			if (intent.authority.length == 0) {
				canGrant = intendedSender = $sender;
			}

			/// @dev Reset the hot reference to the authHash.
			authHash = 0x0;
			j = 0;
			k = 0;

			/// @dev Load the transaction from the intent.
			transaction = intent.transaction;

			require(
				transaction.to == address(this),
				'PlugCore:invalid-intent-target'
			);

			/// @dev Iterate over the authority permissions.
			for (j; j < intent.authority.length; j++) {
				/// @dev Load the permission from the intent.
				signedPermission = intent.authority[j];

				/// @dev Determine the signer of the permission.
				permissionSigner = getSignedPermissionSigner(signedPermission);

				/// @dev Implied sending account is the signer of the first permission.
				if (j == 0) canGrant = intendedSender = permissionSigner;

				/// @dev Ensure the permission signer has authority to grant
				///      the claimed permission.
				require(
					permissionSigner == canGrant,
					'PlugCore:invalid-permission-signer'
				);

				/// @dev Warm up the permission reference.
				permission = signedPermission.permission;

				/// @dev Ensure the permission is valid.
				require(
					permission.authority == authHash,
					'PlugCore:invalid-authority-permission-link'
				);

				/// @dev Retrieve the packet hash for the permission.
				permissionHash = getSignedPermissionHash(signedPermission);

				/// @dev Loop through all the execution caveats declared in the permission
				///      and ensure they are all valid.
				for (k; k < permission.caveats.length; ) {
					/// @dev Call the enforcer to determine if the caveat is valid.
					require(
						CaveatEnforcer(permission.caveats[k].enforcer)
							.enforceCaveat(
								permission.caveats[k].terms,
								intent.transaction,
								permissionHash
							),
						'PlugCore:caveat-rejected'
					);

					unchecked {
						++k;
					}
				}

				/// @dev Store the hash of this permission in `authHash` to verify the
				///      next permission can be verified against it.
				authHash = permissionHash;

				/// @dev Set the next permission signer as the current permission signer.
				canGrant = permission.delegate;
			}

			/// @dev Verify the delegate at the end of the permission chain is the signer.
			require(canGrant == $sender, 'PlugCore:invalid-signer');

			/// @dev Execute the transaction.
			$success = _execute(
				transaction.to,
				transaction.data,
				transaction.gasLimit,
				intendedSender
			);

			unchecked {
				++i;
			}
		}
	}
}
