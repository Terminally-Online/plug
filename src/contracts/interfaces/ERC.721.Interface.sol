// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

interface ERC721Interface {
    function ownerOf(uint256 tokenId) external view returns (address);
}

