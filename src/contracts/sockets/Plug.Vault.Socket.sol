// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';
import {PlugTrading} from '../abstracts/Plug.Trading.sol';
import {Receiver} from 'solady/src/accounts/Receiver.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';
import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {PlugLib} from '../libraries/Plug.Lib.sol';
import {ERC721Interface} from '../interfaces/ERC.721.Interface.sol';

/**
 * @title Plug Vault Socket
 * @author @nftchance (chance@utc24.io)
 */
contract PlugVaultSocket is PlugSocket, PlugTrading, Receiver {
	/// @dev Bit related shifts and masks for access management.
	uint8 internal constant DEFAULT_ACCESS = 0x10;
	uint8 internal constant ACCESS = 0x1;
	uint8 internal constant ACCESS_DENIED = 0x0;
	uint8 internal constant SIGNER_SHIFT = 0x4;

	/// @notice Mapping that holds the access definition of addresses.
	/// @dev This stores two "boolean" values in a single uint9 slot
	///      to save on storage costs. The first 4 bits are for the
	///      router, and the last 4 bits are for the signer.
	mapping(uint160 => mapping(uint160 => uint8)) public access;

	/*
	 * @notice The constructor for the Plug Vault Socket will
	 *         initialize to address(1) when not deployed through
	 *         a Socket factory.
	 */
	constructor() {
		initialize(address(1));
	}

	/**
	 * @notice Initializes a new Plug Vault Socket contract.
	 */
	function initialize(address $ownership) public {
        /// @dev Associate the ownership proxy as the factory that deployed
        ///      the contract at the same time as deployment.
        _initializeOwnership($ownership);
		/// @dev Initialize the Plug Socket.
		_initializePlug();
	}

	/**
	 * @notice Set the access of a router or signer.
	 * @param $address The address to set the access for.
	 * @param $allowance The bitpacked allowance to set for the address.
	 */
	function setAccess(
		address $address,
		uint8 $allowance
	) public virtual onlyOwner {
		_setAccess($address, $allowance);
	}

	/**
	 * @notice Helper view function that can be used to build the packed
	 *         access value for a router or signer.
	 * @param $isRouter If the address is a router.
	 * @param $isSigner If the address is a signer.
	 * @return $access The packed access value.
	 */
	function getAccess(
		bool $isRouter,
		bool $isSigner
	) public pure returns (uint8 $access) {
		/// @dev Set the router value.
		$access = $isRouter ? ACCESS : ACCESS_DENIED;
		/// @dev Set the signer value preserving the value of the
		///      router flag previously set.
		$access |= $isSigner ? ACCESS << SIGNER_SHIFT : ACCESS_DENIED;
	}

	/**
	 * @notice Helper view function that can be used to unpack the access
	 *         value for a router or signer.
	 * @param $address The address to get the access for.
	 * @return $isRouter If the address is a router.
	 * @return $isSigner If the address is a signer.
	 */
	function getAccess(
		address $address
	) public view returns (bool $isRouter, bool $isSigner) {
		/// @dev Retrieve the state from storage.
		uint8 $access = access[uint160(owner())][uint160($address)];
		/// @dev Unpack the router and signer flags.
		$isRouter = _enforceAccess($access);
		$isSigner = $address == owner() || _enforceAccess($access >> SIGNER_SHIFT);
	}

	/**
	 * @notice Internal management of the access a router or signer has.
	 * @dev Note that you cannot toggle off the canonical router. All participants
	 *      of the Plug ecosystem are expected to follow pie-growing principles.
	 *      Even by explicitly setting the router to false, the router will still
	 *      be considered a router. This is to prevent any malicious actors from
	 *      attempting to disable the router and disrupt the ecosystem.
	 * @dev If you would like to append additional logic to this function, you can
	 *      override it in a derived contract and call super._setAccess() to ensure
	 *      that the access state is properly managed.
	 * @param $address The address to manage access for.
	 * @param $allowance The bitpacked allowance to set for the address.
	 */
	function _setAccess(address $address, uint8 $allowance) internal virtual {
		/// @dev Set the packed access state for the address.
		access[uint160(owner())][uint160($address)] = $allowance;
	}

	/**
	 * See { PlugEnforce._enforceRouter }
	 */
	function _enforceRouter(
		address $router
	) internal view override returns (bool $allowed) {
		/// @dev Confirm the router is allowed by recovering the packed access
		///      state as well as checking if the router is the canonical router.
		$allowed =
			_enforceAccess(access[uint160(owner())][uint160($router)]) ||
			super._enforceRouter($router);
	}

	/**
	 * See { PlugEnforce._enforceSigner }
	 */
	function _enforceSigner(
		address $signer
	) internal view override returns (bool $allowed) {
		/// @dev Confirm the signer is allowed by recovering the packed access state.
		$allowed = $signer == owner() || _enforceAccess(
			access[uint160(owner())][uint160($signer)] >> SIGNER_SHIFT
		);
	}

	/**
	 * @notice Enforce the access of a router or signer.
	 * @param $state The state of the access.
	 * @return $allowed If the access is allowed.
	 */
	function _enforceAccess(
		uint8 $state
	) internal pure returns (bool $allowed) {
		/// @dev Confirm the masked state is equal to the flag state.
		$allowed = $state & ACCESS == ACCESS;
	}

	/**
	 * See { PlugSocket-name }
	 */
	function name() public pure override returns (string memory $name) {
		$name = 'PlugVaultSocket';
	}

	/**
	 * See { PlugSocket-version }
	 */
	function version() public pure override returns (string memory $version) {
		$version = '0.0.1';
	}
}
