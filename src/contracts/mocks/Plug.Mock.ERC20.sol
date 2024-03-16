// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract PlugMockERC20 is ERC20, Ownable {
    constructor() ERC20() { }

    function initialize(address $owner) public {
        _initializeOwner($owner);
    }

    function mint(address $to, uint256 $amount) public onlyOwner {
        _mint($to, $amount);
    }

    function name() public pure override returns (string memory) {
        return "MockERC20";
    }

    function symbol() public pure override returns (string memory) {
        return "MERC20";
    }
}
