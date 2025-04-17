// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { SuperchainERC20 } from "op/SuperchainERC20.sol";
import { PredeployAddresses } from "op/libraries/PredeployAddresses.sol";
import { Initializable } from "solady/utils/Initializable.sol";
import { Ownable } from "solady/auth/Ownable.sol";

import { PlugAddressesLib } from "../libraries/Plug.Lib.sol";

/**
 * @title PlugToken
 * @notice The "token" contract for the Plug ecosystem.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugToken is Initializable, Ownable, SuperchainERC20 {
    uint256 public constant TOTAL_SUPPLY = 9_000_000 ether;
    uint32 public transferUnlock = type(uint32).max;
    uint32 public bridgeUnlock = type(uint32).max;

    // @dev Mapping tracking which addresses have been given permission to transfer
    //      tokens prior to unlock. This is to be used by things such as onchain
    //      Reward Center contracts and distribution fans at the application layer.
    mapping(address => bool) public senderToAllowed;

    // @dev Initialize to address(1) to prevent hostile takeover of implementation.
    constructor() {
        _initializeOwner(address(1));
    }

    /**
     * @notice Initializes the token with a given unlock time, owner, and total supply.
     * @dev The complete total supply is instantly minted to the treasury for safe
     *      holding, distribution, and start-to-finish audit trails to verify the
     *      the supply distribution every step along the way.
     */
    function initialize(
        uint32 $unlock,
        address $owner,
        uint256 $totalSupply
    )
        public
        initializer
    {
        transferUnlock = $unlock;
        bridgeUnlock = $unlock;

        _initializeOwner($owner);
        _mint($owner, $totalSupply);
    }

    /**
     * @notice Sets the unlock time for transfers.
     * @param $unlock The new unlock time for transfers.
     */
    function setTransferUnlock(uint32 $unlock) public onlyOwner {
        transferUnlock = $unlock;
    }

    /**
     * @notice Sets the unlock time for bridge operations.
     * @param $unlock The new unlock time for bridge operations.
     */
    function setBridgeUnlock(uint32 $unlock) public onlyOwner {
        bridgeUnlock = $unlock;
    }

    /**
     * @notice Sets the sender allowed status for the token.
     * @param $sender The sender to set the allowed status for.
     * @param $allowed The allowed status to set for the sender.
     */
    function setSenderAllowed(
        address $sender,
        bool $allowed
    )
        external
        onlyOwner
    {
        senderToAllowed[$sender] = $allowed;
    }

    /**
     * @notice See {ERC20-name}.
     */
    function name() public pure override returns (string memory) {
        return "Plug";
    }

    /**
     * @notice See {ERC20-symbol}.
     */
    function symbol() public pure override returns (string memory) {
        return "PLUG";
    }

    /**
     * @notice See {ERC20-_beforeTokenTransfer}.
     */
    function _beforeTokenTransfer(
        address $from,
        address $to,
        uint256
    )
        internal
        virtual
        override
    {
        /// @dev If the bridge is engaging in an operation, we must check the bridge unlock
        //      first the following branch is determined by the variable allowance here.
        if (msg.sender == PredeployAddresses.SUPERCHAIN_TOKEN_BRIDGE) {
            require(block.timestamp >= bridgeUnlock, "PlugToken:bridge-locked");
            return;
        }

        /// @dev Skip all checks for minting/burning as we have no direct influence over
        //       minting beyond the entire supply being minted to the treasury.
        if ($from == address(0) || $to == address(0)) return;

        /// @dev If the transfer is not yet unlocked and the sender is not allowed,
        //      or the sender is not the owner, revert.
        if (
            block.timestamp < transferUnlock && !senderToAllowed[$from]
                && $from != owner()
        ) {
            revert("PlugToken:transfer-locked");
        }
    }
}
