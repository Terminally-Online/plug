// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugInitializable } from "../abstracts/Plug.Initializable.sol";

import { LibBitmap } from "solady/src/utils/LibBitmap.sol";

/**
 * @title Plug Vault Socket
 * @notice This contract represents an personal relay for a single owner, and
 *         declared set of signers.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugVaultSocket is PlugInitializable {
    using LibBitmap for LibBitmap.Bitmap;

    /// @dev The signers of the contract.
    LibBitmap.Bitmap internal signers;

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
     * @notice Name used for the domain separator.
     */
    function name() public pure override returns (string memory) {
        return "PlugVaultSocket";
    }

    /**
     * @notice Version used for the domain separator.
     */
    function version() public pure override returns (string memory) {
        return "0.0.0";
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
