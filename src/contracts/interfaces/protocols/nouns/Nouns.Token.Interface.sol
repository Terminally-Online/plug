// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

interface NounsTokenInterface {
    function seeds(uint256)
        external
        view
        returns (uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses);
}

interface INounsSeeder {
    struct Seed {
        uint48 background;
        uint48 body;
        uint48 accessory;
        uint48 head;
        uint48 glasses;
    }
}
