// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Erc20 is ERC20 {
    constructor() ERC20("BINDING_ERC20", "BINDING_ERC20") {}
}
