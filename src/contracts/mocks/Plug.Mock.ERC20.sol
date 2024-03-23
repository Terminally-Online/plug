// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

contract PlugMockERC20 is ERC20 {
    function mint(address $to, uint256 $amount) public {
        _mint($to, $amount);
    }

    function name() public pure override returns (string memory) {
        return "MockERC20";
    }

    function symbol() public pure override returns (string memory) {
        return "MERC20";
    }
}
