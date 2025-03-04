// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.23;

/**
 * @title Plug Database
 * @notice The chain is a database and a Plug Socket is empowered with a global
 *         scratch space for usable identity co-located variables and state.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugDatabase {
    mapping(address => mapping(bytes32 => bytes32)) public database;

    /**
     * @notice Set a value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function set(bytes32 key, bytes32 value) external returns (bytes32 result) {
        return database[msg.sender][key] = value;
    }

    /**
     * @notice Get a value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function get(address sender, bytes32 key) external view returns (bytes32 result) {
        return database[sender][key];
    }
    
    /**
     * @notice Delete a value for a given key from the caller's storage
     * @param key The key to delete
     */
    function remove(bytes32 key) external {
        delete database[msg.sender][key];
    }
}
