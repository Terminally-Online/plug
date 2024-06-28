// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { NounsTokenInterface } from "../../../libraries/protocols/nouns/Plug.Nouns.Lib.sol";

contract NounsTokenMock is NounsTokenInterface {
    function seeds(uint256)
        public
        pure
        override
        returns (uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
    {
        return (0, 0, 0, 0, 0);
    }
}
