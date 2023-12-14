// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";

/**
 * @title Plug Vault Socket
 * @notice This contract represents an personal relay for a single owner, and
 *         declared set of signers.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugVaultSocket is PlugSocket, Ownable, Receiver {
    using LibBitmap for LibBitmap.Bitmap;

    /// @dev Whether or not the contract has been initialized.
    bool private initialized;

    /// @dev The signers of the contract.
    LibBitmap.Bitmap internal signers;

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

    /**
     * @notice Toggle a signer on or off.
     * @param $signer The address of the signer.
     */
    function toggleSigner(address $signer) public onlyOwner {
        signers.toggle(uint160($signer));
    }

    /**
     * @notice Determine whether or not an address is a declared signer
     *         or the implicit owner of the vault.
     * @param $isSigner true if the address is a signer, false otherwise.
     */
    function isSigner(address $signer) public view returns (bool $isSigner) {
        $isSigner = $signer == owner() || signers.get(uint160($signer));
    }

    /**
     * @notice Prevent the contract from executing the transaction
     *         if the sender is not an approved signer.
     * @param $signer The address of the signer.
     */
    function _enforceSigner(address $signer) internal view override {
        require(isSigner($signer), "PlugSigners:signer-invalid");
    }
}
