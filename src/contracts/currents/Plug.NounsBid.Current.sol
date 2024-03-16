// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { ERC20WETHInterface } from "../interfaces/ERC.20.WETH.Interface.sol";
import { PlugNounsLib } from "../libraries/Plug.Nouns.Lib.sol";

/**
 * @title Plug Nouns Bid Current
 * @notice This contract is responsible for coordinating the placing of bids in the
 *         Nouns Auction house and the following management of wins and losses.
 *         Nouns are automatically distributed by the following bidder (in protocol)
 *         although the automated distribution can be front-run manually.
 * @notice The purpose of this contract is to automatically solve for the active id
 *         of the Noun that is currently being auctioned as well as handling
 *         auto-settlement of expired auctions so that users do not need to chain the
 *         conditionals or even be aware that is what is happening.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugNounsBidCurrent {
    /// @dev The WETH token address used in case the ETH transfer fails.
    ERC20WETHInterface private constant WETH =
        ERC20WETHInterface(payable(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2));

    /// @dev The hippest way to reference ETH with a token address.
    address private constant DOLPHIN_ETH =
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @dev Details of which Noun is active.
    uint256 currentNoun;

    /// @dev Account for the last bid made on each Noun.
    mapping(uint256 => address) public bids;

    /// @dev Receive ETH from the Auction House with no calldata.
    receive() external payable virtual {
        _receive();
    }

    /// @dev Fallback function to receive ETH from the Auction House
    ///      with calldata.
    /// @notice This isn't really needed, but it is nice to have just
    ///         in case something changes.
    fallback() external payable virtual {
        _receive();
    }

    /**
     * @dev Handle the receipt of ETH from the Auction House in the
     *      case that our last bidder was outbid.
     */
    function _receive() internal {
        /// @dev Prevent the receipt of funds from anyone but the Auction House.
        if (msg.sender != address(PlugNounsLib.AUCTION_HOUSE)) {
            revert PlugNounsLib.InvalidSender();
        }

        /// @dev Determine who the last bidder was.
        address bidder = bids[currentNoun];

        /// @dev The winner in our hearts is not the winner on the field.
        delete bids[currentNoun];

        /// @dev Get the balance of this contract.
        uint256 balance = address(this).balance;

        /// @dev Return the money to the hopeful bidder.
        /// @notice Gas usage is limited to 30,000 so that we do not
        ///         run out of gas in the case that the hopeful bidder
        ///         has a fallback function that consumes a lot of gas.
        (bool success,) = bidder.call{ value: balance, gas: 30_000 }("");

        /// @dev Emit an event to let the world know that we have returned
        ///      the money to the hopeful bidder.
        if (success) {
            emit PlugNounsLib.Give(bidder, DOLPHIN_ETH, msg.value);
        }
        /// @dev If a raw ETH transfer fails, then we wrap the ETH in WETH
        ///      and transfer it to the hopeful bidder.
        else {
            WETH.deposit{ value: balance }();
            WETH.transfer(bidder, balance);

            /// @dev Emit an event to let the world know that we have returned
            ///      the money to the hopeful bidder.
            emit PlugNounsLib.Give(bidder, address(WETH), msg.value);
        }
    }

    /**
     * @dev Use the money provided by the sender to bid on the current auction.
     * @param $value The amount of money to use.
     */
    function bid(uint256 $value) public payable {
        /// @dev Get the current state of the auction.
        (uint256 $nounId,,, uint256 $endTime,, bool $settled) =
            PlugNounsLib.AUCTION_HOUSE.auction();

        /// @dev If the auction has concluded and a new auction has not
        ///      been scheduled, then settle the active auction and create one.
        if ($endTime <= block.timestamp && !$settled) {
            PlugNounsLib.AUCTION_HOUSE.settleCurrentAndCreateNewAuction();

            /// @dev Get the current state of the active auction.
            ($nounId,,,,,) = PlugNounsLib.AUCTION_HOUSE.auction();
        }

        /// @dev If a Noun was won, but has not yet been transferred to
        ///      to the winner, then transfer it to the winner.
        /// @notice We check that it does not equal the zero address because
        ///         when a user is outbid, we delete the bid from the mapping
        ///         as well as we set it to the zero address after claiming
        ///         the Noun in the case that they won. Therefore,
        ///         if the ids are different and the address is not the zero
        ///         address, then we know that the Noun was won by our last
        ///         bidder but has not been transferred.
        if ($nounId > currentNoun && bids[currentNoun] != address(0)) {
            take(currentNoun);
        }

        /// @dev Track the current noun.
        currentNoun = $nounId;

        /// @dev Set the bidder as the current winner of the auction.
        bids[$nounId] = msg.sender;

        /// @dev Bid on the auction.
        PlugNounsLib.AUCTION_HOUSE.createBid{ value: $value }($nounId);

        /// @dev Emit an event to let the world know that we have bid on
        ///      the auction.
        emit PlugNounsLib.Bid(msg.sender, msg.sender, $value, $nounId);
    }

    /**
     * @dev Withdraw the money that the sender has deposited.
     * @notice Anyone can call this function though the token will always
     *         go to the winner and not the caller.
     * @param $nounId The id of the Noun being withdrawn.
     */
    function take(uint256 $nounId) public {
        /// @dev Confirm that the auction has been settled to this contract.
        if (PlugNounsLib.TOKEN.ownerOf($nounId) != address(this)) {
            revert PlugNounsLib.InsufficientOwnership();
        }

        /// @dev Warm up the slot of our winner.
        address winner = bids[$nounId];

        /// @dev Remove the winning bid from circulation.
        delete bids[$nounId];

        /// @dev Transfer the Noun to the winner.
        PlugNounsLib.TOKEN.transferFrom(address(this), winner, $nounId);

        /// @dev Emit an event to let the world know that we have
        ///      transferred the Noun to the winner.
        emit PlugNounsLib.Take(
            msg.sender, winner, address(PlugNounsLib.TOKEN), $nounId
        );
    }
}
