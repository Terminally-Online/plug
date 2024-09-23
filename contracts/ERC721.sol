// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract Erc721 is ERC721 {
    constructor() ERC721("BINDING_ERC721", "BINDING_ERC721") {}
}
