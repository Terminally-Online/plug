// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { NounsAuctionHouseInterface } from
    "../../../interfaces/protocols/nouns/Nouns.AuctionHouse.Interface.sol";

/**
 * @notice This mock is designed so that the first auction only ever runs and `settleCurrentAndCreateNewAuction`
 *         is called, but doesn't really do anything. If it needs to progress to the next auction, this contract
 *         needs to be updated to enable progression.
 */
contract NounsAuctionHouseMock is NounsAuctionHouseInterface {
    struct Bid {
        address bidder;
        uint256 value;
    }

    mapping(uint256 => Bid) public bids;

    function settleCurrentAndCreateNewAuction() external { }

    function auction()
        external
        view
        returns (
            uint256 nounId,
            uint256 amount,
            uint256 startTime,
            uint256 endTime,
            address bidder,
            bool settled
        )
    {
        return
            (uint256(0), uint256(0), uint256(0), block.timestamp + 14 days, bids[0].bidder, false);
    }

    function createBid(uint256 $nounId) external payable {
        require(msg.value > bids[$nounId].value, "Invalid bid amount");

        bids[$nounId] = Bid({ bidder: msg.sender, value: msg.value });
    }
}
