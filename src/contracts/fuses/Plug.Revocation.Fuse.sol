//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

/// @dev Plug abstracts.
import {PlugFuse} from '../abstracts/Plug.Fuse.sol';
import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {PlugSocket} from '../abstracts/Plug.Socket.sol';

/// @dev Hash declarations and decoders for the Plug framework.
import {ECDSA} from 'solady/src/utils/ECDSA.sol';

/**
 * @title Revocation Enforcer
 * @notice This Fuse Enforcer operates as an independent instance of the
 *         Plug enabling the revocation of previously signed pins.
 *         After revocation, it is not possible for the signer to reuse the
 *         exact same pin therefore it is recommended to set salt as
 *         as the timestamp of generation (in milliseconds) to ensure that
 *         the signer can still reuse the same pin with a new salt.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
contract PlugRevocationFuse is PlugFuse, PlugSocket {
	/// @notice Use the ECDSA library for signature verification.
	using ECDSA for bytes32;

	/// @dev Mapping of revoked pins.
	mapping(bytes32 => bool) isRevoked;

	constructor() {
		_initializeSocket('RevocationEnforcer', '0.0.1');
	}

	/**
	 * See {FuseEnforcer-enforceFuse}.
	 */
	function enforceFuse(
		bytes calldata,
		PlugTypesLib.Current calldata,
		bytes32 $pinHash
	) public view override returns (bytes memory $callback) {
		/// @dev Ensure the pin has not been revoked.
		require(!isRevoked[$pinHash], 'RevocationEnforcer:revoked');

		/// @dev Otherwise, clear for takeoff.
		$callback = bytes('');
	}

	/**
	 * @notice Enables a Delegator to revoke the pins of a previously
	 *         signed signature.
	 * @param $signedPin The signed pin to revoke.
	 * @param $domainHash The domain hash of the pin.
	 */
	function revoke(
		PlugTypesLib.LivePin calldata $signedPin,
		bytes32 $domainHash
	) public {
		/// @dev Only allow signers of pins to revoke a signature.
		///      Of course, revocation itself could be delegated.
		require(
			getSigner($signedPin, $domainHash) == _msgSender(),
			'RevocationEnforcer:invalid-revoker'
		);

		/// @dev Determine the hash of the pin.
		bytes32 pinHash = getLivePinHash($signedPin);

		/// @dev Ensure the pin has not already been revoked.
		require(!isRevoked[pinHash], 'RevocationEnforcer:already-revoked');

		/// @dev Mark the pin as revoked.
		isRevoked[pinHash] = true;
	}

	/**
	 * @notice Determine the signer of a signed pin.
	 * @dev We use custom functions here because the domain separator is
	 *      different for each LivePin.
	 * @param $signedPin The signed pin to determine the signer of.
	 * @param $domainHash The domain hash of the pin.
	 * @return $signer The address of the signer.
	 */
	function getSigner(
		PlugTypesLib.LivePin memory $signedPin,
		bytes32 $domainHash
	) public view returns (address $signer) {
		/// @dev Determine the digest of the pin and recover the signer.
		$signer = getDigest($signedPin.pin, $domainHash).recover(
			$signedPin.signature
		);
	}

	/**
	 * @notice Determine the digest of a pin.
	 * @param $pin The pin to determine the digest of.
	 * @param $domainHash The domain hash of the pin.
	 * @return $digest The digest of the pin.
	 */
	function getDigest(
		PlugTypesLib.Pin memory $pin,
		bytes32 $domainHash
	) public pure returns (bytes32 $digest) {
		/// @dev Encode the pin and domain hash and hash them.
		$digest = keccak256(
			abi.encodePacked('\x19\x01', $domainHash, getPinHash($pin))
		);
	}
}
