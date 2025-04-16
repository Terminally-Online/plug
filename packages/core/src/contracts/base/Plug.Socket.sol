// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {Ownable} from 'solady/auth/Ownable.sol';
import {Receiver} from 'solady/accounts/Receiver.sol';
import {UUPSUpgradeable} from 'solady/utils/UUPSUpgradeable.sol';
import {ReentrancyGuard} from 'solady/utils/ReentrancyGuard.sol';
import {LibBitmap} from 'solady/utils/LibBitmap.sol';
import {ECDSA} from 'solady/utils/ECDSA.sol';
import {LibBytes} from 'solady/utils/LibBytes.sol';

import {PlugSocketInterface} from '../interfaces/Plug.Socket.Interface.sol';
import {PlugTypes} from '../abstracts/Plug.Types.sol';
import {PlugLib, PlugTypesLib} from '../libraries/Plug.Lib.sol';
import {PlugCoilLib} from '../libraries/Plug.Coil.Lib.sol';

/**
 * @title PlugSocket
 * @notice The "account" contract for a user of Plug.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugSocket is
	PlugSocketInterface,
	PlugTypes,
	Ownable,
	Receiver,
	UUPSUpgradeable,
	ReentrancyGuard
{
	using ECDSA for bytes32;
	using LibBitmap for LibBitmap.Bitmap;
	using PlugCoilLib for bytes;

	uint256 private constant TYPE_CALL = 0x00;
	uint256 private constant TYPE_DELEGATECALL = 0x01;
	uint256 private constant TYPE_CALL_WITH_VALUE = 0x02;
	uint256 private constant TYPE_STATICCALL = 0x03;

	LibBitmap.Bitmap private nonces;

	mapping(address oneClicker => bool allowed) public oneClickersToAllowed;

	constructor() {
		initialize(address(1), address(1));
	}

	/**
	 * @notice Modifier to enforce the signer of the transaction.
	 * @param $input The LivePlugs the definition of execution as well as the
	 *               signature used to verify the execution permission.
	 */
	modifier enforceSignature(PlugTypesLib.LivePlugs calldata $input) {
		if (_enforceSignature($input) == false) {
			revert PlugLib.PlugFailed(
				type(uint8).max,
				PlugLib.PlugCoreSignatureInvalid
			);
		}
		_;
	}

	/**
	 * @notice Modifier to enforce the sender of the transaction.
	 */
	modifier enforceSender() {
		if (_enforceSender(msg.sender) == false) {
			revert PlugLib.PlugFailed(
				type(uint8).max,
				PlugLib.PlugCoreSenderInvalid
			);
		}
		_;
	}

	/**
	 * @notice Modifier to enforce the validity of the solver proof provided.
	 * @param $proof The encoded data that defines the solver.
	 * @param $solver The address of the alleged solver running the transaction.
	 */
	modifier enforceSolver(bytes calldata $proof, address $solver) {
		if ($proof.length != 0) {
			if ($proof.length < 0x40) {
				revert PlugLib.PlugFailed(
					type(uint8).max,
					PlugLib.PlugCoreSolverMalformed
				);
			}
			if (
				uint256(LibBytes.loadCalldata($proof, 0x00)) < block.timestamp
			) {
				revert PlugLib.PlugFailed(
					type(uint8).max,
					PlugLib.PlugCoreSolverExpired
				);
			}
			if (
				address(
					uint160(uint256(LibBytes.loadCalldata($proof, 0x20)))
				) != $solver
			) {
				revert PlugLib.PlugFailed(
					type(uint8).max,
					PlugLib.PlugCoreSolverInvalid
				);
			}
		}
		_;
	}

	/**
	 * @notice See {PlugSocketInterface-initialize}.
	 */
	function initialize(address $owner, address $oneClicker) public {
		_initializeOwner($owner);

		if ($oneClicker != address(0)) {
			oneClickersToAllowed[$oneClicker] = true;
		}
	}

	/**
	 * @notice See {PlugSocketInterface-plug}.
	 */
	function plug(
		PlugTypesLib.LivePlugs calldata $livePlugs,
		address $solver
	) external payable virtual nonReentrant enforceSignature($livePlugs) {
		_plug($livePlugs.plugs, $solver);
	}

	/**
	 * @notice See {PlugSocketInterface-plug}.
	 */
	function plug(
		PlugTypesLib.Plugs calldata $plugs
	) external payable virtual nonReentrant enforceSender {
		_plug($plugs, address(0));
	}

	/**
	 * @notice Enable specific addresses to build the final route of the Plug.
	 * @param $oneClickers The address of the one clicker.
	 * @param $allowance The allowance of the one clicker.
	 */
	function oneClick(
		address[] calldata $oneClickers,
		bool[] calldata $allowance
	) public virtual onlyOwner {
		for (uint256 i; i < $oneClickers.length; i++) {
			oneClickersToAllowed[$oneClickers[i]] = $allowance[i];
		}
	}

	/**
	 * @notice See { PlugSocket-name }
	 */
	function name() public pure override returns (string memory $name) {
		$name = 'Plug Socket';
	}

	/**
	 * @notice See { PlugSocket-symbol }
	 */
	function symbol() public pure override returns (string memory $symbol) {
		$symbol = 'PLUGS';
	}

	/**
	 * @notice See { PlugSocket-version }
	 */
	function version() public pure override returns (string memory $version) {
		$version = '0.0.1';
	}

	/**
	 * @notice See { PlugSocket-hash }
	 */
	function hash(
		PlugTypesLib.LivePlugs calldata $livePlugs
	) public pure override returns (bytes32 $livePlugsHash) {
		return getLivePlugsHash($livePlugs);
	}

	/**
	 * @notice Confirm that signer has permission to declare execution of a
	 *         Plug bundle on the parent-socket that inherits this contract.
	 * @dev Inheriting contracts must implement the logic of this function to make
	 *      sure that only signatures intended for this scope are allowed.
	 * @param $input The LivePlugs object that contains the Plugs object as well as
	 *               the signature defining the permission to execute the bundle.
	 * @return $allowed True if the signature is valid, false otherwise.
	 */
	function _enforceSignature(
		PlugTypesLib.LivePlugs calldata $input
	) internal virtual returns (bool $allowed) {
		address signer = getPlugsDigest($input.plugs).recover($input.signature);
		uint256 nonce = uint256(uint96(bytes12($input.plugs.salt)));
		if (nonces.get(nonce) == true) {
			revert PlugLib.PlugFailed(
				type(uint8).max,
				PlugLib.PlugCoreNonceInvalid
			);
		}
		nonces.set(nonce);
		$allowed = oneClickersToAllowed[signer] || owner() == signer;
	}

	/**
	 * @notice Confirm that the sender of the transaction is allowed as the Socket
	 *         was directly interacted with.
	 * @dev Inheriting contracts must implement the logic of this function to make
	 *      sure that only senders intended for this scope are allowed.
	 * @param $sender The sender of the transaction.
	 * @return $allowed True if the sender is allowed, false otherwise.
	 */
	function _enforceSender(
		address $sender
	) internal view virtual returns (bool $allowed) {
		$allowed = $sender == owner() || $sender == address(this);
	}

	/**
	 * @notice Submit the next transaction with the appropriate call-type that
	 *         will result in the proper side effects and responses.
	 * @param $plug The transaction data that has been post update application.
	 * @return $success The results of the execution.
	 * @return $result The results of the execution.
	 */
	function _call(
		bytes memory $plug
	) internal returns (bool $success, bytes memory $result) {
		(uint8 selector, address to, uint256 value, bytes memory data) = $plug
			.decode();

		if (selector == TYPE_DELEGATECALL) {
			($success, $result) = to.delegatecall(data);
		} else if (selector == TYPE_CALL) {
			($success, $result) = to.call(data);
		} else if (selector == TYPE_CALL_WITH_VALUE) {
			($success, $result) = to.call{value: value}(data);
		} else if (selector == TYPE_STATICCALL) {
			($success, $result) = to.staticcall(data);
		}
	}

	/**
	 * @notice Execute a set of Plugs.
	 * @param $plugs The Plugs to execute containing the bundle and side effects.
	 * @param $solver Encoded data defining the Solver and compensation.
	 */
	function _plug(
		PlugTypesLib.Plugs calldata $plugs,
		address $solver
	) internal enforceSolver($plugs.solver, $solver) {
		bool success;
		bytes[] memory state = new bytes[]($plugs.plugs.length * 2);
		for (uint256 i; i < $plugs.plugs.length; i++) {
			uint256 input = i * 2;
			state[input] = $plugs.plugs[i].data;
			for (uint256 ii; ii < $plugs.plugs[i].updates.length; ii++) {
				(state[input], ) = state[
					$plugs.plugs[i].updates[ii].slice.index
				].transform(i, $plugs.plugs[i].updates[ii], state[input]);
			}
			(success, state[input + 1]) = _call(state[input]);
			if (!success) {
				revert PlugLib.PlugFailed(i, PlugLib.PlugCorePlugFailed);
			}
		}
	}

	/**
	 * @notice See { Ownable._guardInitializeOwner }
	 */
	function _guardInitializeOwnership() internal pure returns (bool $guard) {
		$guard = true;
	}

	/**
	 * @notice See { UUPSUpgradeable._authorizeUpgrade }
	 */
	function _authorizeUpgrade(address) internal virtual override onlyOwner {}
}
