// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { ERC721 } from "solady/tokens/ERC721.sol";
import { Initializable } from "solady/utils/Initializable.sol";
import { Ownable } from "solady/auth/Ownable.sol";
import { LibString } from "solady/utils/LibString.sol";

import { PlugLib, PlugAddressesLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Ticket
 * @notice This contract implements a soulbound (non-transferable) ERC721 token
 *         that represents participation within the Plug ecosystem.
 * @dev Tokens are minted sequentially and can only be minted once per address.
 *      Once minted, tokens cannot be transferred between addresses (soulbound).
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugTicket is Initializable, Ownable, ERC721 {
    using LibString for uint256;

    uint256 public totalSupply;

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
    function initialize() public initializer {
        _initializeOwner(PlugAddressesLib.PLUG_OWNER_ADDRESS);
    }

    /**
     * @notice Mints a new soulbound ticket to the owner of the calling contract.
     * @dev Can only be called by a contract, and will mint to that contract's owner
     *      Each address can only mint one token, and tokens are minted sequentially.
     */
    function mint() public {
        if (msg.sender.code.length == 0) revert PlugLib.CallerMustBeContract();
        if (balanceOf(msg.sender) != 0) revert PlugLib.AlreadyMinted();

        _mint(msg.sender, totalSupply++);
    }

    /**
     * @dev See {ERC721-transferFrom}.
     * @notice Overridden to implement soulbound behavior. Only allows transfers
     *         to/from zero address (minting/burning).
     */
    function transferFrom(
        address $from,
        address $to,
        uint256 $id
    )
        public
        payable
        virtual
        override
    {
        if ($from == address(0) || $to == address(0)) {
            super.transferFrom($from, $to, $id);
            return;
        }
        revert PlugLib.NonTransferableToken();
    }

    /**
     * See {ERC721-name}
     */
    function name() public pure override returns (string memory) {
        return "Plug: Ticket";
    }

    /**
     * See {ERC721-symbol}
     */
    function symbol() public pure override returns (string memory) {
        return "TICKET";
    }

    /**
     * See {ERC721-tokenURI}
     */
    function tokenURI(uint256 $id) public pure override returns (string memory) {
        return string.concat("https://onplug.io/canvas/nft/", $id.toString());
    }
}
