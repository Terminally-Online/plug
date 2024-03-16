// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { ERC721 } from "solady/src/tokens/ERC721.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

import { LibString } from "solady/src/utils/LibString.sol";

contract PlugMockERC721 is ERC721, Ownable {
    using LibString for uint256;

    constructor() { }

    function initialize(address $owner) public {
        _initializeOwner($owner);
    }

    function mint(address $to, uint256 $tokenId) public onlyOwner {
        _mint($to, $tokenId);
    }

    function name() public pure override returns (string memory) {
        return "MockERC721";
    }

    function symbol() public pure override returns (string memory) {
        return "MERC721";
    }

    function tokenURI(uint256 $tokenId)
        public
        pure
        override
        returns (string memory)
    {
        return
            string(abi.encodePacked("https://mock.com/", $tokenId.toString()));
    }
}
