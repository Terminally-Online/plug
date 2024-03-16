// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import {PlugTradingInterface} from '../interfaces/Plug.Trading.Interface.sol';

/**
 * @title Plug Trading
 * @notice Enables the ability to represent Vault ownership through the current
 *         state of an ERC721 that is managed inside the factory that deployed
 *         the Vault. This way, Vaults can be traded on any major marketplace
 *         enabling the ability to spread workflows and earnings of
 *         aforementioned workflows such as points and yield.
 * @author nftchance (chance@onplug.io)
 */
abstract contract PlugTrading is PlugTradingInterface {
    /// @dev The address that houses the ownership information.
	address public ownership;

    /// @dev Track the active owner of the Vault.
	address private _owner;

    /**
     * @notice Modifier enforcing the caller to be the ownership proxy.
     */
	modifier onlyTradable() {
		require(msg.sender == ownership, 'PlugTrading:forbidden-caller');
		_;
	}

	/**
	 * @notice Only the owner of the token can call functions that have
     *         this modifier applied onto it.
	 */
	modifier onlyOwner() {
		require(msg.sender == owner(), 'PlugTrading:forbidden-caller');
		_;
	}

    /**
     * @notice Set the address of the ownership proxy which is a ERC721
     *         compliant contract that lives inside of the factory.
     */
	function _initializeOwnership(address $ownership) internal {
		ownership = $ownership;
	}

    /**
     * @notice Transfer the ownership of a Vault to a new address when the
     *         NFT is transferred.
     */
	function transferOwnership(address $newOwner) public virtual onlyTradable {
		_owner = $newOwner;
	}

    /**
     * @notice Get the owner of the Vault.
     */
	function owner() public view virtual returns (address) {
		return _owner;
	}
}
