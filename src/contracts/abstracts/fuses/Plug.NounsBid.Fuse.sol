// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugFuseInterface } from "../../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../../abstracts/Plug.Types.sol";
import { PlugNounsLib } from "../../libraries/Plug.Nouns.Lib.sol";

abstract contract PlugNounsBidFuse is PlugFuseInterface {
    /// @dev Keep track of the balances of each user.
    mapping(address => uint256) public balances;

    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32
    )
        public
        view
        override
        returns (bytes memory $through)
    {
        (bool $settleUnsettled, address $bidder, uint256 $bid) = decode($live);

        /// @dev Get the current state of the auction.
        (,,, uint256 $endTime, address $winner, bool $settled) =
            PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev Prevent the user from bidding on an auction that they
        ///      have already won / are winning.
        if ($winner == $bidder) revert PlugNounsLib.InsufficientReason();

        /// @dev Prevent the user from bidding on an auction that has
        ///      not yet been settled.
        if (!$settled && $endTime <= block.timestamp && !$settleUnsettled) {
            revert PlugNounsLib.InsufficientSettlement();
        }

        /// @dev Make sure the user has enough money to bid.
        if (balances[$bidder] < $bid) revert PlugNounsLib.InsufficientBalance();

        /// @dev Make sure the bid - fees is large enough to cover
        ///		 the minimum bid.

        /// @dev Callback to transfer the fee to the protocol.
        $through = $current.data;
    }

    function decode(bytes calldata $live)
        public
        pure
        returns (bool $settleUnsettled, address $bidder, uint256 $bid)
    {
        /// @dev Decode the live data.
        ($settleUnsettled, $bidder, $bid) =
            abi.decode($live, (bool, address, uint256));
    }

    function encode(
        bool $settleUnsettled,
        address $bidder,
        uint256 $bid
    )
        public
        pure
        returns (bytes memory $live)
    {
        /// @dev Encode the live data.
        $live = abi.encode($settleUnsettled, $bidder, $bid);
    }
}
