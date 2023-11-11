//SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

/// @dev Plug abstracts.
import {CaveatEnforcer} from '../abstracts/CaveatEnforcer.sol';
import {PlugCore} from '../abstracts/PlugCore.sol';

/// @dev Hash declarations and decoders for the Plug framework.
import {ECDSA} from 'solady/src/utils/ECDSA.sol';

/**
 * @title Revocation Enforcer
 * @notice This Caveat Enforcer operates as an independent instance of the
 *         Plug enabling the revocation of previously signed permissions.
 *         After revocation, it is not possible for the signer to reuse the
 *         exact same permission therefore it is recommended to set salt as
 *         as the timestamp of generation (in milliseconds) to ensure that
 *         the signer can still reuse the same permission with a new salt.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
contract RevocationEnforcer is CaveatEnforcer, PlugCore {
	/// @notice Use the ECDSA library for signature verification.
	using ECDSA for bytes32;

	/// @dev Mapping of revoked permissions.
	mapping(bytes32 => bool) isRevoked;

	constructor() PlugCore('RevocationEnforcer', '1') {}

	/**
	 * See {CaveatEnforcer-enforceCaveat}.
	 */
	function enforceCaveat(
		bytes calldata,
		Transaction calldata,
		bytes32 $permissionHash
	) public view override returns (bool $success) {
		/// @dev Ensure the permission has not been revoked.
		require(!isRevoked[$permissionHash], 'RevocationEnforcer:revoked');

		/// @dev Otherwise, clear for takeoff.
		$success = true;
	}

	/**
	 * @notice Enables a Delegator to revoke the permissions of a previously
	 *         signed signature.
	 * @param $signedPermission The signed permission to revoke.
	 * @param $domainHash The domain hash of the permission.
	 */
	function revoke(
		SignedPermission calldata $signedPermission,
		bytes32 $domainHash
	) public {
		/// @dev Only allow signers of permissions to revoke a signature.
		///      Of course, revocation itself could be delegated.
		require(
			getSigner($signedPermission, $domainHash) == _msgSender(),
			'RevocationEnforcer:invalid-revoker'
		);

		/// @dev Determine the hash of the permission.
		bytes32 permissionHash = getSignedPermissionHash($signedPermission);

		/// @dev Ensure the permission has not already been revoked.
		require(
			!isRevoked[permissionHash],
			'RevocationEnforcer:already-revoked'
		);

		/// @dev Mark the permission as revoked.
		isRevoked[permissionHash] = true;
	}

	/**
	 * @notice Determine the signer of a signed permission.
	 * @dev We use custom functions here because the domain separator is
	 *      different for each SignedPermission.
	 * @param $signedPermission The signed permission to determine the signer of.
	 * @param $domainHash The domain hash of the permission.
	 * @return $signer The address of the signer.
	 */
	function getSigner(
		SignedPermission memory $signedPermission,
		bytes32 $domainHash
	) public view returns (address $signer) {
		/// @dev Determine the digest of the permission and recover the signer.
		$signer = getDigest($signedPermission.permission, $domainHash).recover(
			$signedPermission.signature
		);
	}

	/**
	 * @notice Determine the digest of a permission.
	 * @param $permission The permission to determine the digest of.
	 * @param $domainHash The domain hash of the permission.
	 * @return $digest The digest of the permission.
	 */
	function getDigest(
		Permission memory $permission,
		bytes32 $domainHash
	) public pure returns (bytes32 $digest) {
		/// @dev Encode the permission and domain hash and hash them.
		$digest = keccak256(
			abi.encodePacked(
				'\x19\x01',
				$domainHash,
				getPermissionHash($permission)
			)
		);
	}
}
