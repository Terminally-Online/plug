// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugFuseInterface } from
    "../interfaces/Plug.Fuse.Interface.sol";
import { PlugLib, PlugTypesLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Rate Limit Fuse
 * @notice This Fuse enables the ability to rate limit the throughput of a Plug
 *         a variable period between uses rather than an explicit fixed period.
 *         The rate limit can be applied to a single bundle of Plugs or to the
 *         global state of a Socket.
 * @notice Use cases for having a cooldown between calls:
 *     - Prevent a user from spamming a function that should only be called
 *       once every so often.
 *     - Enable recurring transactions to be executed on a regular basis that
 *       do not have an always-enforced start and end time.
 * @author nftchance (chance@onplug.io)
 */
contract PlugRateLimitFuse is PlugFuseInterface {
    /// @dev The structure of a rate limit bucket.
    struct Bucket {
        uint32 lastUpdatedAt;
        uint224 availableTokens;
    }

    /// @dev Keep track of the last time a bundle of Plugs was used.
    mapping(address => mapping(bytes32 => Bucket)) consumption;

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
        (bool global, uint32 replenishRate, uint32 max) =
            decode($terms);

        /// @dev Retrieve the current state of the timestamp and use from storage.
        Bucket storage bucket =
            consumption[msg.sender][global ? bytes32(0) : $plugsHash];

        /// @dev Determine how many tokens have been replenished since last use.
        uint256 tokensToAdd =
            (block.timestamp - bucket.lastUpdatedAt) / replenishRate;

        /// @dev Account for partial tokens by adjusting the refresh window.
        bucket.lastUpdatedAt = uint32(
            block.timestamp
                - (
                    (block.timestamp - bucket.lastUpdatedAt)
                        % replenishRate
                )
        );

        /// @dev Determine how many tokens are available while acounting for
        ///      the amount of time that has passed since last use.
        bucket.availableTokens = uint224(
            (bucket.availableTokens + tokensToAdd > max)
                ? max
                : bucket.availableTokens + tokensToAdd
        );

        /// @dev Make sure that there are tokens that can be used.
        if (bucket.availableTokens == 0) {
            revert PlugLib.ThresholdExceeded(max, 0);
        }

        /// @dev Decrement the available tokens.
        bucket.availableTokens -= 1;

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * See {PlugFuseInterface-decode}.
     */
    function decode(bytes calldata $terms)
        public
        pure
        returns (bool $global, uint32 $replenishRate, uint32 $max)
    {
        ($global, $replenishRate, $max) =
            abi.decode($terms, (bool, uint32, uint32));
    }

    /**
     * See {PlugFuseInterface-encode}.
     */
    function encode(
        bool $global,
        uint32 $replenishRate,
        uint128 $max
    )
        public
        pure
        returns (bytes memory $terms)
    {
        $terms = abi.encode($global, $replenishRate, $max);
    }
}
