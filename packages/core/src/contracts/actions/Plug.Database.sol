// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.23;

import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";

/**
 * @title Plug Database
 * @notice The chain is a database and a Plug Socket is empowered with a global
 *         scratch space for usable identity co-located variables and state.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugDatabase {
    mapping(address => mapping(bytes32 => bytes32)) public database;

    function set(bytes32 key, bytes32 value) external returns (bytes32 output) {
        return database[msg.sender][key] = value;
    }

    function get(bytes32 key) external view returns (bytes32 output) {
        return database[msg.sender][key];
    }
}
