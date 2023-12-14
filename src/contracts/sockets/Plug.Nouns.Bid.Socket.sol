// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { INounsSeeder, INounsToken } from "../interfaces/nouns/INounsToken.sol";

import { Ownable } from "solady/src/auth/Ownable.sol";

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { NounsBidLib } from "../libraries/nouns/Nouns.Bid.Lib.sol";
import { NounsBidFuse } from "../fuses/nouns/Nouns.Bid.Fuse.sol";

import { BytesLib } from "../libraries/BytesLib.sol";

/**
 * @title Plug Nouns Bid Socket
 * @notice This contract is responsible for coordinating the placing of bids in the
 *         Nouns Auction house and the following management of wins and losses without
 *         holding funds or collecting a TVL. ETH is automatically distributed when used,
 *         and Nouns are automatically distributed by the following bidder (in protocol)
 *         although the automated distribution can be front-run manually.
 * @author @nftchance (chance@utc24.io)
 * @author @masonchain
 */
contract PlugNounsBidSocket is PlugSocket, NounsBidFuse, Ownable {
    using BytesLib for bytes;

    /// @dev Accessible interface to the active Nouns Auction House.
    INounsToken public immutable NOUN;

    /// @dev The hippest way to reference ETH with a token address.
    address private constant DOLPHIN_ETH = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @dev On bid execution, 0.0002% of the bid is sent to the protocol.
    /// @notice In effect, that means it costs < $5 to prevent thousands
    ///         in lost economic value and/or collection appraisal (RFV).
    uint256 private constant PROTOCOL_FEE = 2;

    /// @dev Details of which Noun is active.
    uint256 currentNoun;

    /// @dev Account for the last bid made on each Noun.
    mapping(uint256 => address) public bids;

    constructor(address $auctionHouse, address $nouns, address $owner) NounsBidFuse($auctionHouse) {
        /// @dev Initialize the Auction House and Nouns interfaces.
        NOUN = INounsToken($nouns);

        /// @dev Initialize the owner.
        _initializeOwner($owner);

        /// @dev Initialize the Plug Socket.
        _initializeSocket("NounsBidSocket", "0.0.1");
    }

    receive() external payable virtual {
        _receive();
    }

    /**
     * @dev Handle when ETH is received.
     */
    function _receive() internal {
        /// @dev Prevent receiving funds from anyone other than the Auction House.
        if (msg.sender == address(AUCTION_HOUSE)) {
            /// @dev Determine who the last bidder was.
            address bidder = bids[currentNoun];

            /// @dev The winner in our hearts is not the winner on the field.
            delete bids[currentNoun];

            /// @dev Return the money to the hopeful bidder.
            give(bidder);
        }
        /// @dev Money has been deposited.
        else {
            give(msg.sender);
        }
    }

    /**
     * @dev Handle when money is received on behalf of a user.
     * @param $onBehalf The address of the sender.
     */
    function give(address $onBehalf) public payable virtual {
        /// @dev Account for the money deposited by the sender.
        balances[$onBehalf] += msg.value;

        emit NounsBidLib.Given($onBehalf, msg.value);
    }

    /**
     * @dev Use the money provided by the sender to bid on the current auction.
     * @param $value The amount of money to use.
     */
    function use(uint256 $value) internal {
        /// @dev Get the current state of the auction.
        (uint256 $nounId,,, uint256 $endTime,, bool $settled) = AUCTION_HOUSE.auction();

        /// @dev If the auction has concluded and a new auction has not
        ///      been scheduled, then settle the active auction and create one.
        if ($endTime <= block.timestamp && !$settled) {
            AUCTION_HOUSE.settleCurrentAndCreateNewAuction();

            /// @dev Get the current state of the active auction.
            ($nounId,,,,,) = AUCTION_HOUSE.auction();
        }

        /// @dev If a Noun was won, but has not yet been transferred to
        ///      to the winner, then transfer it to the winner.
        if ($nounId > currentNoun && bids[currentNoun] != address(0)) {
            take(address(NOUN), currentNoun);
        }

        /// @dev Update the balance of the sender.
        balances[_msgSender()] -= $value;

        /// @dev Update the value that is being bid.
        uint256 bidValue = $value - ($value / PROTOCOL_FEE);

        /// @dev Account for the protocol fee.
        balances[owner()] += $value - bidValue;

        /// @dev Bid on the auction.
        AUCTION_HOUSE.createBid{ value: bidValue }($nounId);

        /// @dev Track the current noun.
        currentNoun = $nounId;

        /// @dev Set the bidder as the current winner of the auction.
        bids[$nounId] = _msgSender();

        emit NounsBidLib.Used(msg.sender, _msgSender(), $value, $nounId);
    }

    /**
     * @dev Withdraw the money that the sender has deposited.
     * @param $asset The address of the asset being withdrawn.
     * @param $value The token id of the Noun being withdrawn.
     */
    function take(address $asset, uint256 $value) public {
        /// @dev If the asset is ETH, then transfer the ETH to the sender.
        if ($asset == DOLPHIN_ETH) {
            /// @dev Make sure the sender has sufficient balance to cover the call.
            if (balances[_msgSender()] < $value) {
                revert NounsBidLib.InsufficientBalance();
            }

            balances[_msgSender()] -= $value;

            /// @dev Transfer the ETH to the sender.
            (bool success,) = _msgSender().call{ value: $value }("");

            if (!success) revert NounsBidLib.InsufficientOwnership();

            emit NounsBidLib.Taken(msg.sender, _msgSender(), $asset, $value);
        }
        /// @dev If the asset is Nouns, then transfer the Nouns to the sender.
        else if ($asset == address(NOUN)) {
            /// @dev Confirm that the auction has been settled to this contract.
            if (NOUN.ownerOf($value) != address(this)) {
                revert NounsBidLib.InsufficientOwnership();
            }

            address winner = bids[$value];

            /// @dev Remove the winning bid from circulation.
            delete bids[$value];

            NOUN.transferFrom(address(this), winner, $value);

            emit NounsBidLib.Taken(_msgSender(), winner, $asset, $value);
        }
        /// @dev If the asset is not ETH or Nouns, then revert.
        else {
            revert NounsBidLib.InsufficientReason();
        }
    }
}
