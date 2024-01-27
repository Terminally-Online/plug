// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";

// TODO: Implement storage of the signer and sender so that contracts can choose
//       to associate the execution of a contract to the intended party rather
//       than always having to reference things to the royalty address.

/**
 * @title Plug Router Socket
 * @notice This contract represents a general purpose relay socket that can be
 * used
 *         to route transactions to other contracts.
 * @notice Do not approve assets to this contract as anyone can sign and/or
 * execute
 *         transactions which means they can steal your tokens if you do.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugRouterSocket is PlugSocket, Ownable, Receiver {
    using LibBitmap for LibBitmap.Bitmap;

    /// @dev Whether or not the contract has been initialized.
    bool private initialized;

    /**
     * @notice Initializes a new Plug Vault contract.
     */
    constructor() {
        initialize(msg.sender);
    }

    /**
     * @notice Modifier to ensure that the contract has not been initialized.
     */
    modifier initializer() {
        require(!initialized, "PlugVaultSocket:already-initialized");

        initialized = true;
        _;
    }

    /**
     * @notice Initialize a new Plug Vault.
     * @param $owner The owner of the vault.
     */
    function initialize(address $owner) public payable virtual initializer {
        /// @dev Initialize the owner.
        _initializeOwner($owner);

        /// @dev Initialize the Plug Socket.
        _initializeSocket("PlugVaultSocket", "0.0.0");
    }
}
