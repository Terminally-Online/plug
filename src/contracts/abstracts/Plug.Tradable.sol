// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import {ERC721} from 'solady/src/tokens/ERC721.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';

import {LibString} from 'solady/src/utils/LibString.sol';

import {PlugTradingInterface} from '../interfaces/Plug.Trading.Interface.sol';

/**
 * @title Plug Tradable
 * @notice This enables the single-housing of Plug Vaults that have been deployed
 *         irrespective of the version they are actively using. Notably, this
 *         contract is responsible for managing the ownership of the vault through
 *         a mirror-like function which means when the owner of this token changes,
 *         the owner of the vault does as well.
 * @author nftchance (chance@onplug.io)
 */
abstract contract PlugTradable is ERC721, Ownable {
	using LibString for uint256;

	/// @dev The base endpoint for the metadata.
	string private baseURI;

	constructor(address $owner, string memory $baseURI) {
		/// @dev Initialize the metadata controller.
		_initializeOwner($owner);

		/// @dev Set the base URI for the token.
		baseURI = $baseURI;
	}

	/**
	 * @notice Set the base URI for the token.
	 * @param $baseURI The base URI to be set for the token.
	 */
	function setBaseURI(string memory $baseURI) external onlyOwner {
		baseURI = $baseURI;
	}

	/**
	 * @notice Mint a new vault token to the owner of the vault.
	 * @param $admin The owner of the vault.
	 * @param $vault The address of the vault.
	 */
	function mint(
		address $admin,
		address $vault
	) public virtual {
		/// @dev Mint the token to the new owner of the vault.
		_mint($admin, uint256(uint160($vault)));
	}

	function _afterTokenTransfer(
		address from,
		address to,
		uint256 tokenId
	) internal override {
		super._afterTokenTransfer(from, to, tokenId);

        /// @dev Call into the vault that is representing this asset and update
        ///      the owner in storage. We do not rely on a runtime read for ownership
        ///      references because that would end up costing more than just storing
        ///      the data over there as well.
		PlugTradingInterface(address(uint160(tokenId))).transferOwnership(to);
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
