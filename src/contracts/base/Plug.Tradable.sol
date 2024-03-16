// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {PlugLib} from '../libraries/Plug.Lib.sol';

import {ERC721} from 'solady/src/tokens/ERC721.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';

import {LibString} from 'solady/src/utils/LibString.sol';

/**
 * @title Plug Tradable
 * @notice This enables the single-housing of Plug Vaults that have been deployed
 *         irrespective of the version they are actively using. Notably, this
 *         contract is responsible for managing the ownership of the vault through
 *         a mirror-like function which means when the owner of this token changes,
 *         the owner of the vault does as well.
 * @dev The usage of this model increase the underlying gas costs of execution by a
 *      small margin though the vault-internal implementation already circuits out of
 *      `if` loops that are to be checking this resulting in this only ever being
 *      when an owner actually calls the state of something or someone with a bad
 *      bundler packet attempts running execution through a vault that is not their own.
 * @author nftchance (chance@onplug.io)
 */
contract PlugTradable is ERC721, Ownable {
	using LibString for uint256;

	/// @dev The base endpoint for the metadata.
	string private baseURI;

	/// @dev The address of the Plug router that is allowed to
	///      call the mint function.
	address private plug;

	/// @dev Make sure only the router can call this function.
	modifier onlyPlug() {
		require(msg.sender == plug, 'PlugTradable:forbidden-sender');
		_;
	}

	constructor(address $owner, string memory $baseURI, address $plug) {
		/// @dev Initialize the metadata controller.
		_initializeOwner($owner);

		/// @dev Set the base URI for the token.
		baseURI = $baseURI;

		/// @dev Set the address of the Plug router.
		plug = $plug;
	}

	/**
	 * @notice Set the base URI for the token.
	 * @param $baseURI The base URI to be set for the token.
	 */
	function setBaseURI(string memory $baseURI) external onlyOwner {
		baseURI = $baseURI;
	}

	/**
	 * @notice Set the address of the Plug router.
	 * @param $plug The address of the Plug router.
	 */
	function setPlug(address $plug) external onlyOwner {
		plug = $plug;
	}

	/**
	 * @notice Mint a new vault token to the owner of the vault.
	 * @param $admin The owner of the vault.
	 * @param $version The version of the vault.
	 * @param $vault The address of the vault.
	 */
	function mint(
		address $admin,
		uint96 $version,
		address $vault
	) public virtual onlyPlug returns (uint256 $tokenId) {
		/// @dev Prefix the token id of the vault with the version number followed
		///      by the vault address that is being deployed.
		$tokenId = uint256(uint160($vault)) | ($version << 160);

		/// @dev Mint the token to the new owner of the vault.
		_mint($admin, $tokenId);
	}

	/**
	 * @notice Metadata response for the name of the collection.
	 * @return $name The name of the collection.
	 */
	function name() public pure override returns (string memory $name) {
		$name = 'Plug Vaults';
	}

	/**
	 * @notice Metadata response for the symbol of the collection.
	 * @return $symbol The symbol of the collection.
	 */
	function symbol() public pure override returns (string memory $symbol) {
		$symbol = 'PLUG';
	}

	/**
	 * @notice Returns the compiled URI for a given token ID.
	 * @param $tokenId The token ID to query the URI for.
	 * @return $uri The URI for the given token ID.
	 */
	function tokenURI(
		uint256 $tokenId
	) public view override returns (string memory $uri) {
		$uri = string(abi.encodePacked(baseURI, $tokenId.toString()));
	}
}
