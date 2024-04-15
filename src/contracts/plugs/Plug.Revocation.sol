// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../interfaces/Plug.Connector.Interface.sol";

/**
 * @title Plug Revocation
 * @notice This Plug operates enables the ability to revoke execution rights
 *         of any previously signed bundle of Plugs.
 * @notice Use cases for bidding on Nouns:
 *     - Enable safe intent reuse by enabling the ability to declare a bundle
 *       of Plugs that can be revoked at any time.
 *     - Operate with programmable revocation fulfilled by conditions outside
 *       of manual and/or human desire.
 * @dev After revocation, it is not possible for the signer to reuse the
 *      exact same pin therefore it is recommended to set salt as
 *      as the timestamp of generation (in milliseconds) to ensure that
 *      the signer can still reuse the same pin with a new salt.
 * @author @nftchance (chance@onplug.io)
 */
contract PlugRevocation is PlugConnectorInterface {
    /// @dev Keep track of which bundles of Plugs have been revoked.
    mapping(address => mapping(bytes32 => bool)) public isRevoked;

    /**
     * See {PlugConnectorInterface-enforce}.
     *
     * @dev While an individual can delegate permission to a different user,
     *      the controller is is the one who can revoke the bundle of Plugs.
     *      A user cannot "maliciously" revoke a bundle of Plugs with control
     *      outside of the definition of the sender themselves.
     */
    function enforce(bytes calldata $terms, bytes32 $plugsHash) public view virtual {
        /// @dev Decode the (declared) controller of the bundle of Plugs.
        address sender = decode($terms);

        /// @dev Ensure the plug has not been revoked.
        require(isRevoked[sender][$plugsHash] == false, "PlugRevocation:revoked");
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
    function revoke(bytes32 $plugsHash, bool $revoked) public virtual {
        /// @dev Mark the bundle of Plugs as revoked.
        isRevoked[msg.sender][$plugsHash] = $revoked;
    }

    /**
     * @dev Decode the clamp data into the two bounds.
     */
    function decode(bytes calldata $data) public pure returns (address $sender) {
        $sender = abi.decode($data, (address));
    }

    /**
     * @dev Encode the clamp bounds.
     */
    function encode(address $sender) public pure returns (bytes memory $data) {
        $data = abi.encode($sender);
    }
}
