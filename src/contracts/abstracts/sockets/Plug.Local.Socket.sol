// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypesLib } from "../Plug.Types.sol";
import { PlugInitializable } from "../Plug.Initializable.sol";

/**
 * @title Plug Router Socket
 * @notice This contract represents a general purpose relay socket that can be
 *         used to route transactions to other contracts.
 * @notice Consumers of this abstract must implemented `.name` and `.version`.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugLocalSocket is PlugInitializable {
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
