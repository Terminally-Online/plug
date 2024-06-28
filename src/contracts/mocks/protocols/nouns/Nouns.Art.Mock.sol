// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { NounsArtInterface } from "../../../interfaces/protocols/nouns/Nouns.Art.Interface.sol";

contract NounsArtMock is NounsArtInterface {
    function backgrounds(uint256 index) public pure returns (string memory) {
        index;
        return string(bytes(abi.encode(keccak256("background"))));
    }

    function heads(uint256 index) public pure returns (bytes memory) {
        index;
        return bytes(abi.encode(keccak256("head")));
    }

    function bodies(uint256 index) public pure returns (bytes memory) {
        index;
        return bytes(abi.encode(keccak256("body")));
    }

    function accessories(uint256 index) public pure returns (bytes memory) {
        index;
        return bytes(abi.encode(keccak256("accessory")));
    }

    function glasses(uint256 index) public pure returns (bytes memory) {
        index;
        return bytes(abi.encode(keccak256("glasses")));
    }
}
