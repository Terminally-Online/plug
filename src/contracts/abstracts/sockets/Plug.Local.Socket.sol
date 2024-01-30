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
