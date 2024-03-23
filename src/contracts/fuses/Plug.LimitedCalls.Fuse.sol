// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugFuseInterface } from
    "../interfaces/Plug.Fuse.Interface.sol";
import { PlugLib, PlugTypesLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Limited Calls Fuse
 * @notice This Fuse enables the ability to limit the number of times an
 *         intent can be used onchain.
 * @notice Use cases for limiting calls:
 *     - Allow an intent declaration to be reused multiple times such as DCA
 *       models or expiring subscription definitions.
 *     - Remove the need for single nonce-use with a multi-or-infinite nonce pattern.
 *     - Lack of use results in infinite signatures that are dependent on conditions
 *       outside of a nonce architecture.
 * @author nftchance (chance@onplug.io)
 */
contract PlugLimitedCallsFuse is PlugFuseInterface {
    /// @dev Keep track of how many times a pin has been used.
    mapping(address => mapping(bytes32 => uint256)) callCounts;

    /**
     * See {PlugFuseInterface-enforceFuse}.
     */
    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32 $plugsHash
    )
        public
        override
        returns (bytes memory $through)
    {
        /// @dev Snapshot the current state of the call allowance and consumption.
        uint256 allowedCalls = decode($live);
        uint256 calls = ++callCounts[msg.sender][$plugsHash];

        /// @dev Confirm the declared limit has not been exceeded.
        if (allowedCalls < calls) {
            revert PlugLib.ThresholdExceeded(allowedCalls, calls);
        }

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * See {PlugFuseInterface-decode}.
     */
    function decode(bytes calldata $terms)
        public
        pure
        returns (uint256 $callCount)
    {
        $callCount = abi.decode($terms, (uint256));
    }

    /**
     * See {PlugFuseInterface-encode}.
     */
    function encode(uint256 $callCount)
        public
        pure
        returns (bytes memory $terms)
    {
        $terms = abi.encode($callCount);
    }
}
