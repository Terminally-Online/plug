// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugFuse} from '../../abstracts/Plug.Fuse.sol';
import {PlugTypesLib} from '../../abstracts/Plug.Types.sol';
import {INounsAuctionHouse} from '../../interfaces/nouns/INounsAuctionHouse.sol';
import {NounsBidLib} from '../../libraries/nouns/Nouns.Bid.Lib.sol';

contract NounsBidFuse is PlugFuse {
	INounsAuctionHouse public immutable AUCTION_HOUSE;

	/// @dev Keep track of the balances of each user.
	mapping(address => uint256) public balances;

	constructor(address $auctionHouse) {
		/// @dev Initialize the Auction House and Nouns interfaces.
		AUCTION_HOUSE = INounsAuctionHouse($auctionHouse);
	}

	function enforceFuse(
		bytes calldata $live,
		PlugTypesLib.Current calldata,
		bytes32
	) public view override returns (bytes memory $callback) {
		(bool $settleUnsettled, address $bidder, uint256 $bid) = decode($live);

		/// @dev Get the current state of the auction.
		(, , , uint256 $endTime, address $winner, bool $settled) = AUCTION_HOUSE
			.auction();

		/// @dev Prevent the user from bidding on an auction that they
		///      have already won / are winning.
		if ($winner == $bidder) revert NounsBidLib.InsufficientReason();

		/// @dev Prevent the user from bidding on an auction that has
		///      not yet been settled.
		if (!$settled && $endTime <= block.timestamp && !$settleUnsettled)
			revert NounsBidLib.InsufficientSettlement();

		/// @dev Make sure the user has enough money to bid.
		if (balances[$bidder] < $bid) revert NounsBidLib.InsufficientBalance();

		/// @dev Make sure the bid - fees is large enough to cover
		///		 the minimum bid.

		/// @dev Callback to transfer the fee to the protocol.
		$callback = bytes('');
	}

	function decode(
		bytes calldata $live
	)
		public
		pure
		returns (bool $settleUnsettled, address $bidder, uint256 $bid)
	{
		/// @dev Decode the live data.
		($settleUnsettled, $bidder, $bid) = abi.decode(
			$live,
			(bool, address, uint256)
		);
	}

	function encode(
		bool $settleUnsettled,
		address $bidder,
		uint256 $bid
	) public pure returns (bytes memory $live) {
		/// @dev Encode the live data.
		$live = abi.encode($settleUnsettled, $bidder, $bid);
	}
}
