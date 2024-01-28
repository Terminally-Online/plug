// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";

// TODO: Implement storage of the signer and sender so that contracts can choose
//       to associate the execution of a contract to the intended party rather
//       than always having to reference things to the royalty address.

/**
 * @title Plug Router Socket
 * @notice This contract represents a general purpose relay socket that can be
 *         used to route transactions to other contracts.
 * @notice Do not approve assets to this contract as anyone can sign and/or
 *         execute transactions which means they can use your approvals.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugRouterSocket is PlugSocket {
    /// @dev Whether or not the contract has been initialized.
    bool private initialized;

    /**
     * @notice Initializes a new Plug Vault contract.
     */
    constructor() {
        /// @dev Initialize the Plug Socket.
        _initializeSocket("PlugVaultSocket", "0.0.0");
    }

    /**
     * @notice Modifier to ensure that the contract has not been initialized.
     */
    modifier initializer() {
        require(!initialized, "PlugVaultSocket:already-initialized");

        initialized = true;
        _;
    }
}
