// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

interface NounsArtInterface {
    function backgrounds(uint256) external view returns (string memory);
    function accessories(uint256 index) external view returns (bytes memory);
    function bodies(uint256 index) external view returns (bytes memory);
    function glasses(uint256 index) external view returns (bytes memory);
    function heads(uint256 index) external view returns (bytes memory);
}
