// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Ownable } from "solady/auth/Ownable.sol";
import { ERC721 } from "solady/tokens/ERC721.sol";
import { LibString } from "solady/utils/LibString.sol";

/**
 * @title Plug Ticket
 * @notice This contract implements a soulbound (non-transferable) ERC721 token
 *         that represents participation within the Plug ecosystem.
 * @dev Tokens are minted sequentially and can only be minted once per address.
 *      Once minted, tokens cannot be transferred between addresses (soulbound).
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugTicket is Ownable, ERC721 {
    using LibString for uint256;

    error AlreadyMinted();
    error NonTransferableToken();
    error CallerMustBeContract();

    string private base;
    uint256 private nextId;

    constructor(address $owner) {
        _initializeOwner($owner);
    }

    /**
     * @notice Mints a new soulbound ticket to the owner of the calling contract.
     * @dev Can only be called by a contract, and will mint to that contract's owner
     *      Each address can only mint one token, and tokens are minted sequentially.
     */
    function mint() public {
        if (msg.sender.code.length == 0) revert CallerMustBeContract();

        address user = Ownable(msg.sender).owner();

        if (balanceOf(user) != 0) revert AlreadyMinted();

        _mint(user, nextId++);
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
        revert NonTransferableToken();
    }

    /**
     * @notice Sets the base URI for computing {tokenURI}
     * @dev Only callable by contract owner
     * @param $uri The new base URI to be used for all tokens
     */
    function setBase(string memory $uri) public onlyOwner {
        base = $uri;
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
        return "PLUGT";
    }

    /**
     * See {ERC721-tokenURI}
     */
    function tokenURI(uint256 $id) public view override returns (string memory) {
        return string.concat(base, $id.toString());
    }
}
