//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

/// @dev Plug abstracts.
import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugLocalSocket } from "../abstracts/sockets/Plug.Local.Socket.sol";

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
 * @author @nftchance (chance@utc24.io)
 */
contract PlugRevocationFuse is PlugFuseInterface, PlugLocalSocket {
    /// @notice Use the ECDSA library for signature verification.
    using ECDSA for bytes32;

    /// @dev Keep track of which bundles of Plugs have been revoked.
    mapping(address => mapping(bytes32 => bool)) public isRevoked;

    /// @dev Initialize the local socket.
    constructor() PlugLocalSocket() { }

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
            isRevoked[sender][$plugsHash] == false, "PlugRevocationFuse:revoked"
        );

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @notice Enables a controller to revoke the pins of a
     *         previously signed signature.
     * @dev There is no verification of a signature here because a controller
     *      can only revoke permission to hashes that they control themselves
     *      so the revocation operates with the assumption that only the
     * @param $livePlugs The signed pin to revoke.
     * @param $domainHash The domain hash of the pin.
     */
    function revoke(
        PlugTypesLib.LivePlugs memory $livePlugs,
        bytes32 $domainHash,
        bool $revoked
    )
        public
        virtual
    {
        /// @dev Retrieve the sender.
        address sender = _msgSender();

        /// @dev Determine the hash of the pin.
        bytes32 plugsHash = getPlugsHash($livePlugs.plugs);

        /// @dev Determine the digest of the pin and recover the signer.
        // address signer = keccak256(
        //     abi.encodePacked("\x19\x01", $domainHash, plugsHash)
        // ).recover($livePlugs.signature);

        /// @dev Only allow the signer of a Plugs bundle to revoke a
        ///      signature. Revocation itself could be plugged.
        require(
            getSigner($livePlugs, $domainHash) == _msgSender(),
            "PlugRevocationFuse:invalid-revoker"
        );

        /// @dev Ensure the bundle of Plugs does not already have the same
        ///      revocation state that is being set.
        require(
            isRevoked[sender][plugsHash] != $revoked,
            "PlugRevocationFuse:same-state"
        );

        /// @dev Mark the bundle of Plugs as revoked.
        isRevoked[sender][plugsHash] = $revoked;
    }

    /**
     * @notice Determine the signer of a signed pin.
     * @dev We use custom functions here because the domain separator is
     *      different for each LivePin.
     * @param $signedPin The signed pin to determine the signer of.
     * @param $domainHash The domain hash of the pin.
     * @return $signer The address of the signer.
     */
    function getSigner(
        PlugTypesLib.LivePlugs memory $signedPin,
        bytes32 $domainHash
    )
        public
        view
        returns (address $signer)
    {
        /// @dev Determine the digest of the pin and recover the signer.
        $signer = getDigest($signedPin.plugs, $domainHash).recover(
            $signedPin.signature
        );
    }

    /**
     * @notice Determine the digest of a pin.
     * @param $pin The pin to determine the digest of.
     * @param $domainHash The domain hash of the pin.
     * @return $digest The digest of the pin.
     */
    function getDigest(
        PlugTypesLib.Plugs memory $pin,
        bytes32 $domainHash
    )
        public
        pure
        returns (bytes32 $digest)
    {
        /// @dev Encode the pin and domain hash and hash them.
        $digest = keccak256(
            abi.encodePacked("\x19\x01", $domainHash, getPlugsHash($pin))
        );
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
    function encode(address $sender) public pure returns (bytes memory $data) {
        $data = abi.encode($sender);
    }

    /**
     * See {PlugInitializable-name}.
     */
    function name() public pure virtual override returns (string memory) {
        return "PlugRevocationFuse";
    }
}
