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
     * @notice Automatically initialize the contract that is used for the
     *         the implementation to prevent nefarious interaction with
     *         this contract.
     * @dev We initialize to address(1) instead of address(0) because
     *      it is more difficult to check if a contract has been
     *      initialized if we use address(0).
     */
    constructor() {
        initialize(address(1));
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
     * @dev May be implemented by the child contract.
     */
    function version() public pure virtual returns (string memory) {
        return "0.0.0";
    }

    /**
     * @notice Prevent the owner from being double initialized.
     * @return $guard True if the owner has been initialized.
     */
    function _guardInitializeOwner()
        internal
        pure
        virtual
        override
        returns (bool $guard)
    {
        $guard = true;
    }
}
