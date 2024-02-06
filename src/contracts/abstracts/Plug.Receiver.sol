// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { LibBitmap } from "solady/src/utils/LibBitmap.sol";

/**
 * @title PlugSender
 * @notice A BitMap is used to cast the address to a uint160. It is honestly rather
 *         unlikely though that we will ever have multiple addresses that are
 *         in the same slot. Still, it is a possibility and thus gas would be saved.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugReceiver {
    using LibBitmap for LibBitmap.Bitmap;

    /// @dev The address of the Plug Router.
    address internal constant ROUTER_SOCKET_ADDRESS =
        0x00b09C89Ace100AB7A4Dc47ebfBd1E7997920062;

    /// @dev The bitmap that stores the trusted forwarders.
    LibBitmap.Bitmap internal trustedForwarders;

    /**
     * @notice Toggle the state of a forwarder on or off.
     * @param $forwarder The address of the forwarder.
     * @return $isTrusted true if the address is a trusted forwarder, false otherwise.
     */
    function toggleTrustedForwarder(address $forwarder)
        public
        virtual
        returns (bool $isTrusted)
    {
        $isTrusted = trustedForwarders.toggle(uint160($forwarder));
    }

    /**
     * @notice Determine and return whether or not an address is a trusted forwarder
     *         that will enable the ability to safely recover the signer (or intended
     *         sender) from the end of the calldata.
     * @param $forwarder The address of the forwarder.
     * @return $trusted true if the address is a trusted forwarder, false otherwise.
     */
    function isTrustedForwarder(address $forwarder)
        public
        view
        virtual
        returns (bool $trusted)
    {
        $trusted = msg.sender == ROUTER_SOCKET_ADDRESS
            || trustedForwarders.get(uint160($forwarder));
    }

    /**
     * @notice Returns the intended sender of a Plug Current.
     * @dev The intended sender is the address that sent the transaction to the Plug Router.
     * @return $sender The intended sender of the Plug Current.
     */
    function _msgSender() internal view returns (address $sender) {
        /// @dev The call is coming from an external source such as an EOA or a
        ///      Vault contract that is not the Plug Router.
        $sender = msg.sender;

        /// @dev The call came from a Trusted Forwarder.
        if (isTrustedForwarder(msg.sender)) {
            assembly {
                /// @dev Check if the call data size is at least 20 bytes.
                if gt(calldatasize(), 19) {
                    /// @dev Extract the sender address from the last 20 bytes of call data.
                    $sender := shr(96, calldataload(sub(calldatasize(), 20)))
                    /// @dev Prevent address(0) from being a valid sender by resetting it to
                    ///      the caller if it is the case.
                    /// @dev Given the proper setup this should never really be possible,
                    ///      if only properly functioning trusted forwarders are enabled,
                    ///      but it is better to be safe than sorry because significant damage
                    ///      could be done if the address(0) is a valid sender to effectively
                    ///      every smart contract that exists.
                    if iszero($sender) { $sender := caller() }
                }
            }
        }
    }
}
