// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugInitializable } from "../Plug.Initializable.sol";
import { PlugReceiver } from "../Plug.Receiver.sol";

import { PlugTypesLib } from "../Plug.Types.sol";

/**
 * @title Plug Router Socket
 * @notice This contract represents a general purpose relay socket that can be
 *         used to route transactions to other contracts.
 * @notice Consumers of this abstract MUST implement `.name()` and MAY choose
 *         to implement `.version()` in the case it diverges from the core
 *         or is not a native primitive of Plug.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugLocalSocket is PlugInitializable, PlugReceiver {
    /**
     * @notice Initializes a new Plug Vault contract.
     */
    constructor() {
        /// @dev Initialize the contract when deployed through a factory.
        initialize(msg.sender);
    }

    /**
     * See {PlugReceiver-isTrustedForwarder}.
     *
     * @notice This method should not be relied upon if the contract has
     *         the capability to make external calls, or call itself through
     *         a means that is not the mechanisms provided by Local Socket.
     * @notice It is of critical importance that a trusted forwarder is not simultaneously
     *         a LocalSocket and Multicallable (or another variant) as one could
     *         create malicious calldata that allows them to impersonate an account
     *         they should not have permission to. If the only means of self-calling
     *         is through a Plug Local Socket that is okay because the decoded
     *         sender is always appended to the end of the calldata.
     */
    function isTrustedForwarder(address $sender)
        public
        view
        virtual
        override
        returns (bool $trusted)
    {
        $trusted =
            msg.sender == address(this) || super.isTrustedForwarder($sender);
    }

    /**
     * @notice Ensure that the only valid ground for the current
     *         is the current contract.
     * @param $current The current to enforce.
     */
    function _enforceCurrent(PlugTypesLib.Current memory $current)
        internal
        view
        override
    {
        require($current.ground == address(this), "PlugCore:invalid-ground");
    }
}
