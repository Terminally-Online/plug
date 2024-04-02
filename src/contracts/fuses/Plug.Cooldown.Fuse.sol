// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    PlugFuseInterface,
    PlugTypesLib
} from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Cooldown Fuse
 * @notice This Fuse enables the ability to limit how often a bundle of Plugs can
 *         be used onchain with the most basic form of a cooldown.
 * @notice Use cases for having a cooldown between calls:
 *     - Prevent a user from spamming a function that should only be called
 *       once every so often.
 *     - Enable recurring transactions to be executed on a regular basis that
 *       do not have an always-enforced start and end time.
 * @author nftchance (chance@onplug.io)
 */
contract PlugCooldownFuse is PlugFuseInterface {
    /// @dev Keep track of the last time a bundle of Plugs was used.
    mapping(address => mapping(bytes32 => uint256))
        socketToPlugsToLastUsed;

    /**
     * See {PlugFuseInterface-enforceFuse}.
     */
    function enforceFuse(
        bytes calldata $terms,
        PlugTypesLib.Current calldata $current,
        bytes32 $plugsHash
    )
        public
        virtual
        returns (bytes memory $through)
    {
        /// @dev Snapshot the current state of the cooldown and use.
        uint256 lastUsed =
            socketToPlugsToLastUsed[msg.sender][$plugsHash];

        /// @dev Confirm one use has already happened.
        if (lastUsed != 0) {
            /// @dev Determine how long it has been since the last use.
            uint256 timeSince = block.timestamp - lastUsed;
            /// @dev Decode the cooldown from the terms.
            uint256 cooldown = decode($terms);

            /// @dev Confirm the cooldown has not been exceeded.
            if (timeSince < cooldown) {
                revert PlugLib.ThresholdExceeded(cooldown, timeSince);
            }
        }

        /// @dev Update the last used timestamp for the next cooldown check.
        socketToPlugsToLastUsed[msg.sender][$plugsHash] =
            block.timestamp;

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * See {PlugFuseInterface-decode}.
     */
    function decode(bytes calldata $terms)
        public
        pure
        returns (uint256 $cooldown)
    {
        $cooldown = abi.decode($terms, (uint256));
    }

    /**
     * See {PlugFuseInterface-encode}.
     */
    function encode(uint256 $cooldown)
        public
        pure
        returns (bytes memory $terms)
    {
        $terms = abi.encode($cooldown);
    }
}
