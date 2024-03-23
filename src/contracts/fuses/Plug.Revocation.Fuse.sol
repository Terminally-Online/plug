// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

/// @dev Plug abstracts.
import { PlugFuseInterface } from
    "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

/// @dev Hash declarations and decoders for the Plug framework.
import { ECDSA } from "solady/src/utils/ECDSA.sol";

/**
 * @title Revocation Enforcer
 * @notice This Fuse Enforcer operates as an independent instance of the
 *         Plug enabling the revocation of previously signed pins.
 *         After revocation, it is not possible for the signer to reuse the
 *         exact same pin therefore it is recommended to set salt as
 *         as the timestamp of generation (in milliseconds) to ensure that
 *         the signer can still reuse the same pin with a new salt.
 * @author @nftchance (chance@onplug.io)
 */
contract PlugRevocationFuse is PlugFuseInterface {
    /// @notice Use the ECDSA library for signature verification.
    using ECDSA for bytes32;

    /// @dev Keep track of which bundles of Plugs have been revoked.
    mapping(address => mapping(bytes32 => bool)) public isRevoked;

    /**
     * See {FuseEnforcer-enforceFuse}.
     *
     * @dev While an individual can delegate permission to a different user,
     *      the controller is is the one who can revoke the bundle of Plugs.
     *      A user cannot "maliciously" revoke a bundle of Plugs with control
     *      outside of the definition of the sender themselves.
     */
    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32 $plugsHash
    )
        public
        view
        virtual
        override
        returns (bytes memory $through)
    {
        /// @dev Decode the (declared) controller of the bundle of Plugs.
        address sender = decode($live);

        /// @dev Ensure the plug has not been revoked.
        require(
            isRevoked[sender][$plugsHash] == false,
            "PlugRevocationFuse:revoked"
        );

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @notice Enables a sender to revoke the intent of a built bundle of Plugs.
     * @dev There is no verification of a signature here because a controller
     *      can only revoke permission to hashes that they control themselves
     *      so the revocation operates with the assumption that a sender is
     *      always revoking the hashes of their own.
     * @param $plugsHash The hash of the bundle of Plugs to revoke.
     * @param $revoked Whether to revoke or un-revoke the bundle of Plugs.
     */
    function revoke(
        bytes32 $plugsHash,
        bool $revoked
    )
        public
        virtual
    {
        /// @dev Mark the bundle of Plugs as revoked.
        isRevoked[msg.sender][$plugsHash] = $revoked;
    }

    /**
     * @dev Decode the clamp data into the two bounds.
     */
    function decode(bytes calldata $data)
        public
        pure
        returns (address $sender)
    {
        $sender = abi.decode($data, (address));
    }

    /**
     * @dev Encode the clamp bounds.
     */
    function encode(address $sender)
        public
        pure
        returns (bytes memory $data)
    {
        $data = abi.encode($sender);
    }
}
