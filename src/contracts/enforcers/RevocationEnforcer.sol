//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Framework abstracts.
import {CaveatEnforcer} from '../abstracts/CaveatEnforcer.sol';
import {FrameworkCore} from '../abstracts/FrameworkCore.sol';

/// @dev Hash declarations and decoders for the Emporium framework.
import {ECDSA} from 'solady/src/utils/ECDSA.sol';

/**
 * @title Revocation Enforcer
 * @notice This Caveat Enforcer operates as an independent instance of the
 *         Framework enabling the revocation of previously signed delegations.
 *         After revocation, it is not possible for the signer to reuse the
 *         exact same delegation therefore it is recommended to set salt as
 *         as the timestamp of generation (in milliseconds) to ensure that
 *         the signer can still reuse the same delegation with a new salt.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
contract RevocationEnforcer is CaveatEnforcer, FrameworkCore {
	/// @notice Use the ECDSA library for signature verification.
	using ECDSA for bytes32;

	/// @dev Mapping of revoked delegations.
	mapping(bytes32 => bool) isRevoked;

	constructor() FrameworkCore('RevocationEnforcer', '1') {}

	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata,
		Transaction calldata,
		bytes32 $delegationHash
	) public view override returns (bool $success) {
		/// @dev Ensure the delegation has not been revoked.
		require(!isRevoked[$delegationHash], 'RevocationEnforcer:revoked');

		/// @dev Otherwise, clear for takeoff.
		$success = true;
	}

	/**
	 * @notice Enables a Delegator to revoke the permissions of a previously
	 *         signed signature.
	 * @param $signedDelegation The signed delegation to revoke.
	 * @param $domainHash The domain hash of the delegation.
	 */
	function revoke(
		SignedDelegation calldata $signedDelegation,
		bytes32 $domainHash
	) public {
		/// @dev Only allow signers of delegations to revoke a signature.
		///      Of course, revocation itself could be delegated.
		require(
			getSigner($signedDelegation, $domainHash) == _msgSender(),
			'RevocationEnforcer:invalid-revoker'
		);

		/// @dev Determine the hash of the delegation.
		bytes32 delegationHash = getSignedDelegationPacketHash(
			$signedDelegation
		);

		/// @dev Ensure the delegation has not already been revoked.
		require(
			!isRevoked[delegationHash],
			'RevocationEnforcer:already-revoked'
		);

		/// @dev Mark the delegation as revoked.
		isRevoked[delegationHash] = true;
	}

	/**
	 * @notice Determine the signer of a signed delegation.
	 * @param $signedDelegation The signed delegation to determine the signer of.
	 * @param $domainHash The domain hash of the delegation.
	 * @return $signer The address of the signer.
	 */
	function getSigner(
		SignedDelegation memory $signedDelegation,
		bytes32 $domainHash
	) public view returns (address $signer) {
		/// @dev Determine the digest of the delegation and recover the signer.
		$signer = getDigest($signedDelegation.delegation, $domainHash).recover(
			$signedDelegation.signature
		);
	}

	/**
	 * @notice Determine the digest of a delegation.
	 * @param $delegation The delegation to determine the digest of.
	 * @param $domainHash The domain hash of the delegation.
	 * @return $digest The digest of the delegation.
	 */
	function getDigest(
		Delegation memory $delegation,
		bytes32 $domainHash
	) public pure returns (bytes32 $digest) {
		/// @dev Encode the delegation and domain hash and hash them.
		$digest = keccak256(
			abi.encodePacked(
				'\x19\x01',
				$domainHash,
				getDelegationPacketHash($delegation)
			)
		);
	}
}
