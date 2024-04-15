// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { ERC721 } from "solady/tokens/ERC721.sol";
import { LibString } from "solady/utils/LibString.sol";

contract PlugMockERC721 is ERC721 {
    using LibString for uint256;

    function mint(address $to, uint256 $tokenId) public {
        _mint($to, $tokenId);
    }

    function name() public pure override returns (string memory) {
        return "MockERC721";
    }

    function symbol() public pure override returns (string memory) {
        return "MERC721";
    }

    function tokenURI(uint256 $tokenId) public pure override returns (string memory) {
        return string(abi.encodePacked("https://mock.com/", $tokenId.toString()));
    }
}
