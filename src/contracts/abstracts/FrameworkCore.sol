// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Hash declarations and decoders for the Emporium framework.
import {Types} from './Types.sol';
/// @dev Error utilities for the Emporium framework.
import {FrameworkErrors} from './FrameworkErrors.sol';

/// @dev Core Framework dependencies.
import {CaveatEnforcer} from './CaveatEnforcer.sol';

/**
 * @title Framework Core
 * @notice The core contract for the Emporium framework that enables
 *         counterfactual revokable delegation of extremely
 *         granular permission and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract FrameworkCore is Types {
	using FrameworkErrors for bytes;

	/// @notice Multi-dimensional account delegation nonce management.
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
	 *         current context. This is important for delegations, as the
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
			'FrameworkCore:nonce2-out-of-order'
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
	 * @notice Execute a batch of invocations
	 * @param $batch The batch of invocations to execute.
	 * @param $sender The address of the sender.
	 * @return $success Whether the transaction was successfully executed.
	 */
	function _invoke(
		Invocation[] calldata $batch,
		address $sender
	) internal returns (bool $success) {
		/// @dev Load the stack.
		uint256 i;
		uint256 j;
		uint256 k;
		address canGrant;
		address intendedSender;
		address delegationSigner;
		bytes32 authHash;
		bytes32 delegationHash;

		/// @dev Load the structs into a hot reference.
		Invocation memory invocation;
		SignedDelegation memory signedDelegation;
		Delegation memory delegation;
		Transaction memory transaction;

		/// @dev Iterate over the batch of invocations.
		for (i; i < $batch.length; ) {
			/// @dev Load the invocation from the batch.
			invocation = $batch[i];

			/// @dev If there are no delegations, this invocation comes from the signer
			if (invocation.authority.length == 0) {
				canGrant = intendedSender = $sender;
			}

			/// @dev Reset the hot reference to the authHash.
			authHash = 0x0;
			j = 0;
			k = 0;

			/// @dev Load the transaction from the invocation.
			transaction = invocation.transaction;

			require(
				transaction.to == address(this),
				'FrameworkCore:invalid-invocation-target'
			);

			/// @dev Iterate over the authority delegations.
			for (j; j < invocation.authority.length; j++) {
				/// @dev Load the delegation from the invocation.
				signedDelegation = invocation.authority[j];

				/// @dev Determine the signer of the delegation.
				delegationSigner = getSignedDelegationSigner(signedDelegation);

				/// @dev Implied sending account is the signer of the first delegation.
				if (j == 0) canGrant = intendedSender = delegationSigner;

				/// @dev Ensure the delegation signer has authority to grant
				///      the claimed permission.
				require(
					delegationSigner == canGrant,
					'FrameworkCore:invalid-delegation-signer'
				);

				/// @dev Warm up the delegation reference.
				delegation = signedDelegation.delegation;

				/// @dev Ensure the delegation is valid.
				require(
					delegation.authority == authHash,
					'FrameworkCore:invalid-authority-delegation-link'
				);

				/// @dev Retrieve the packet hash for the delegation.
				delegationHash = getSignedDelegationPacketHash(
					signedDelegation
				);

				/// @dev Loop through all the execution caveats declared in the delegation
				///      and ensure they are all valid.
				for (k; k < delegation.caveats.length; ) {
					/// @dev Call the enforcer to determine if the caveat is valid.
					require(
						CaveatEnforcer(delegation.caveats[k].enforcer)
							.enforceCaveat(
								delegation.caveats[k].terms,
								invocation.transaction,
								delegationHash
							),
						'FrameworkCore:caveat-rejected'
					);

					unchecked {
						++k;
					}
				}

				/// @dev Store the hash of this delegation in `authHash` to verify the
				///      next delegation can be verified against it.
				authHash = delegationHash;

				/// @dev Set the next delegation signer as the current delegation signer.
				canGrant = delegation.delegate;
			}

			/// @dev Verify the delegate at the end of the delegation chain is the signer.
			require(canGrant == $sender, 'DelegatableCore:invalid-delegate');

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
