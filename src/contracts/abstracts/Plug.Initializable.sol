// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

import { LibBitmap } from "solady/src/utils/LibBitmap.sol";

/**
 * @title Plug Initializable
 * @notice Initialize a socket with an owner and domain.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugInitializable is
    PlugSocket,
    Ownable,
    Receiver,
    Initializable
{
    /**
     * @notice Initializes a new Plug Vault contract.
     */
    constructor() {
        initialize(msg.sender);
    }

    /**
     * @notice Initialize a new Plug Vault.
     * @param $owner The owner of the vault.
     */
    function initialize(address $owner) public payable virtual initializer {
        /// @dev Initialize the owner.
        _initializeOwner($owner);

        /// @dev Initialize the Plug Socket.
        _initializeSocket(name(), version());
    }

    /**
     * @notice Name used for the domain separator.
     * @dev Must be implemented by the child contract.
     */
    function name() public pure virtual returns (string memory);

    /**
     * @notice Version used for the domain separator.
     * @dev Must be implemented by the child contract.
     */
    function version() public pure virtual returns (string memory) {
        return "0.0.0";
    }
}
