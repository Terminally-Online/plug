// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {PlugLib} from '../libraries/Plug.Lib.sol';
import {ERC721Interface} from '../interfaces/ERC.721.Interface.sol';

abstract contract PlugTrading {
	address public tokenOwner;

	/**
	 * @notice Make a single call to the ownership reference to save on gas
	 *         and then clear the owner when the function is done being consumed.
	 */
	modifier withOwner() {
		/// @dev If the tokenOwner has already been set there is no need to call
		///      the ownership proxy again.
		if (tokenOwner == address(0))
			/// @dev Retrieve the owner of the token from the ownership proxy.
			tokenOwner = owner();
		_;
		/// @dev Reset the owner to the zero address to prevent any potential
		///      misuse of the owner reference while reclaiming gas spent.
		delete tokenOwner;
	}

    /**
     * @notice Only the owner of the token can call functions that have this
     *         modifier applied onto it.
     */
	modifier onlyOwner() {
		require(msg.sender == tokenOwner, 'PlugVaultSocket:forbidden-caller');
		_;
	}

	function owner() public view returns (address $owner) {
		$owner =
			ERC721Interface(PlugLib.PLUG_TRADABLE_ADDRESS).ownerOf(
				uint256(uint160(address(this)))
			);
	}
}
