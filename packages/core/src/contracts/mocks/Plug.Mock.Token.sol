// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { ERC20 } from "solady/tokens/ERC20.sol";
import { Initializable } from "solady/utils/Initializable.sol";
import { Ownable } from "solady/auth/Ownable.sol";

/**
 * @title MockPlugToken
 * @notice A simplified mock of the PlugToken for testing purposes
 */
contract MockPlugToken is Initializable, Ownable, ERC20 {
    uint32 public transferUnlock;
    mapping(address => bool) public senderToAllowed;
    
    constructor() {
        _initializeOwner(address(1));
    }
    
    function initialize(uint32 $unlock, address $owner, uint256 $totalSupply) public initializer {
        transferUnlock = $unlock;
        _initializeOwner($owner);
        _mint($owner, $totalSupply);
    }
    
    function setTransferUnlock(uint32 $unlock) public onlyOwner {
        transferUnlock = $unlock;
    }
    
    function setSenderAllowed(address $sender, bool $allowed) external onlyOwner {
        senderToAllowed[$sender] = $allowed;
    }
    
    function name() public pure override returns (string memory) {
        return "Plug";
    }
    
    function symbol() public pure override returns (string memory) {
        return "PLUG";
    }
    
    function mint(address $to, uint256 $amount) public {
        _mint($to, $amount);
    }
    
    function _beforeTokenTransfer(
        address $from,
        address $to,
        uint256
    ) internal virtual override {
        // Skip checks for minting/burning
        if ($from == address(0) || $to == address(0)) return;
        
        // If the transfer is not yet unlocked and the sender is not allowed,
        // or the sender is not the owner, revert.
        if (block.timestamp < transferUnlock && !senderToAllowed[$from] && $from != owner()) {
            revert("PlugToken:transfer-locked");
        }
    }
}