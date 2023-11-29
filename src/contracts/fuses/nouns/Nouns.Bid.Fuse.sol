// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {INounsAuctionHouse} from '../../interfaces/nouns/INounsAuctionHouse.sol';

import {NounsBidSocketHelpers} from '../../libraries/nouns/NounsBidSocketHelpers.sol';

import {Fuse} from '../../abstracts/Fuse.sol';

contract NounsBidFuse is Fuse {
	INounsAuctionHouse public immutable auctionHouse;

	/// @dev Keep track of the balances of each user.
	mapping(address => uint256) public balances;

	constructor(address $auctionHouse) {
		/// @dev Initialize the Auction House and Nouns interfaces.
		auctionHouse = INounsAuctionHouse($auctionHouse);
	}

	function enforceFuse(
		bytes calldata $live,
		Current calldata,
		bytes32
	) public view override returns (bytes memory $callback) {
		(bool $settleUnsettled, address $bidder, uint256 $bid) = decode($live);

		/// @dev Get the current state of the auction.
		(, , , uint256 $endTime, address $winner, bool $settled) = auctionHouse
			.auction();

		/// @dev Prevent the user from bidding on an auction that they
		///      have already won / are winning.
		if ($winner == $bidder)
			revert NounsBidSocketHelpers.InsufficientReason();

		/// @dev Prevent the user from bidding on an auction that has
		///      not yet been settled.
		if ($settled == false && $endTime <= block.timestamp)
			if ($settleUnsettled == false)
				revert NounsBidSocketHelpers.InsufficientSettlement();

		/// @dev Make sure the user has enough money to bid.
		if (balances[$bidder] < $bid)
			revert NounsBidSocketHelpers.InsufficientBalance();

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
