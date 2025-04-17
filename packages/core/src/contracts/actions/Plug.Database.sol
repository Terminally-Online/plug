// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.26;

/**
 * @title Plug Database
 * @notice The chain is a database and a Plug Socket is empowered with a global
 *         scratch space for usable identity co-located variables and state.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugDatabase {
    mapping(address => mapping(bytes32 => bytes32)) public bytes32Storage;
    mapping(address => mapping(bytes32 => uint256)) public uintStorage;
    mapping(address => mapping(bytes32 => int256)) public intStorage;
    mapping(address => mapping(bytes32 => address)) public addressStorage;
    mapping(address => mapping(bytes32 => bool)) public boolStorage;
    mapping(address => mapping(bytes32 => string)) public stringStorage;
    mapping(address => mapping(bytes32 => bytes)) public bytesStorage;

    mapping(address => mapping(bytes32 => bool)) private intExists;
    mapping(address => mapping(bytes32 => bool)) private boolExists;

    uint8 public constant TYPE_BYTES32 = 1;
    uint8 public constant TYPE_UINT256 = 2;
    uint8 public constant TYPE_INT256 = 3;
    uint8 public constant TYPE_ADDRESS = 4;
    uint8 public constant TYPE_BOOL = 5;
    uint8 public constant TYPE_STRING = 6;
    uint8 public constant TYPE_BYTES = 7;

    // Helper functions for type-based key derivation

    /**
     * @notice Set a bytes32 value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function set(bytes32 key, bytes32 value) external returns (bytes32 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BYTES32);
        bytes32Storage[msg.sender][typedKey] = value;
        return value;
    }

    /**
     * @notice Set a uint256 value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setUint(bytes32 key, uint256 value) external returns (uint256 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_UINT256);
        uintStorage[msg.sender][typedKey] = value;
        return value;
    }

    /**
     * @notice Set an int256 value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setInt(bytes32 key, int256 value) external returns (int256 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_INT256);
        intStorage[msg.sender][typedKey] = value;
        intExists[msg.sender][typedKey] = true;
        return value;
    }

    /**
     * @notice Set an address value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setAddress(bytes32 key, address value) external returns (address result) {
        bytes32 typedKey = _deriveKey(key, TYPE_ADDRESS);
        addressStorage[msg.sender][typedKey] = value;
        return value;
    }

    /**
     * @notice Set a boolean value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setBool(bytes32 key, bool value) external returns (bool result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BOOL);
        boolStorage[msg.sender][typedKey] = value;
        boolExists[msg.sender][typedKey] = true;
        return value;
    }

    /**
     * @notice Set a string value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setString(
        bytes32 key,
        string calldata value
    )
        external
        returns (string memory result)
    {
        bytes32 typedKey = _deriveKey(key, TYPE_STRING);
        stringStorage[msg.sender][typedKey] = value;
        return value;
    }

    /**
     * @notice Set a bytes value for a given key in the caller's storage
     * @param key The key to set
     * @param value The value to store
     * @return result The stored value
     */
    function setBytes(bytes32 key, bytes calldata value) external returns (bytes memory result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BYTES);
        bytesStorage[msg.sender][typedKey] = value;
        return value;
    }

    /**
     * @notice Get a bytes32 value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function get(address sender, bytes32 key) external view returns (bytes32 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BYTES32);
        return bytes32Storage[sender][typedKey];
    }

    /**
     * @notice Get a uint256 value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getUint(address sender, bytes32 key) external view returns (uint256 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_UINT256);
        return uintStorage[sender][typedKey];
    }

    /**
     * @notice Get an int256 value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getInt(address sender, bytes32 key) external view returns (int256 result) {
        bytes32 typedKey = _deriveKey(key, TYPE_INT256);
        return intStorage[sender][typedKey];
    }

    /**
     * @notice Get an address value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getAddress(address sender, bytes32 key) external view returns (address result) {
        bytes32 typedKey = _deriveKey(key, TYPE_ADDRESS);
        return addressStorage[sender][typedKey];
    }

    /**
     * @notice Get a boolean value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getBool(address sender, bytes32 key) external view returns (bool result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BOOL);
        return boolStorage[sender][typedKey];
    }

    /**
     * @notice Get a string value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getString(address sender, bytes32 key) external view returns (string memory result) {
        bytes32 typedKey = _deriveKey(key, TYPE_STRING);
        return stringStorage[sender][typedKey];
    }

    /**
     * @notice Get a bytes value for a given key from the specified address's storage
     * @param sender The address whose storage to read
     * @param key The key to retrieve
     * @return result The stored value
     */
    function getBytes(address sender, bytes32 key) external view returns (bytes memory result) {
        bytes32 typedKey = _deriveKey(key, TYPE_BYTES);
        return bytesStorage[sender][typedKey];
    }

    /**
     * @notice Check if a value exists for a given type and key
     * @param sender The address whose storage to check
     * @param key The key to check
     * @param typeId The type to check for
     * @return True if a value exists for the given type and key
     */
    function exists(address sender, bytes32 key, uint8 typeId) external view returns (bool) {
        bytes32 typedKey = _deriveKey(key, typeId);

        if (typeId == TYPE_BYTES32) {
            return bytes32Storage[sender][typedKey] != bytes32(0);
        } else if (typeId == TYPE_UINT256) {
            return uintStorage[sender][typedKey] != 0;
        } else if (typeId == TYPE_INT256) {
            return intExists[sender][typedKey];
        } else if (typeId == TYPE_ADDRESS) {
            return addressStorage[sender][typedKey] != address(0);
        } else if (typeId == TYPE_BOOL) {
            return boolExists[sender][typedKey];
        } else if (typeId == TYPE_STRING) {
            return bytes(stringStorage[sender][typedKey]).length > 0;
        } else if (typeId == TYPE_BYTES) {
            return bytesStorage[sender][typedKey].length > 0;
        }

        return false;
    }

    /**
     * @notice Get the type of a value stored at a given key
     * @param sender The address whose storage to check
     * @param key The key to check
     * @return The type identifier (0 if not found)
     */
    function getType(address sender, bytes32 key) external view returns (uint8) {
        // For performance, check types in order of likely frequency
        for (uint8 i = TYPE_BYTES32; i <= TYPE_BYTES; i++) {
            if (this.exists(sender, key, i)) {
                return i;
            }
        }

        return 0; // Not found
    }

    /**
     * @notice Delete a value for a given key and type from the caller's storage
     * @param key The key to delete
     * @param typeId The type of the value to delete
     */
    function removeWithType(bytes32 key, uint8 typeId) public {
        require(typeId >= TYPE_BYTES32 && typeId <= TYPE_BYTES, "Invalid type ID");

        bytes32 typedKey = _deriveKey(key, typeId);

        if (typeId == TYPE_BYTES32) {
            delete bytes32Storage[msg.sender][typedKey];
        } else if (typeId == TYPE_UINT256) {
            delete uintStorage[msg.sender][typedKey];
        } else if (typeId == TYPE_INT256) {
            delete intStorage[msg.sender][typedKey];
            delete intExists[msg.sender][typedKey];
        } else if (typeId == TYPE_ADDRESS) {
            delete addressStorage[msg.sender][typedKey];
        } else if (typeId == TYPE_BOOL) {
            delete boolStorage[msg.sender][typedKey];
            delete boolExists[msg.sender][typedKey];
        } else if (typeId == TYPE_STRING) {
            delete stringStorage[msg.sender][typedKey];
        } else if (typeId == TYPE_BYTES) {
            delete bytesStorage[msg.sender][typedKey];
        }
    }

    /**
     * @notice Delete all types of values for a given key
     * @param key The key to delete
     */
    function remove(bytes32 key) external {
        for (uint8 i = TYPE_BYTES32; i <= TYPE_BYTES; i++) {
            if (this.exists(msg.sender, key, i)) {
                removeWithType(key, i);
            }
        }
    }

    /**
     * @notice Batch set multiple key-value pairs of the same type for gas efficiency
     * @param keys Array of keys to set
     * @param values Array of values to store
     * @param typeId The type of values being set
     */
    function batchSet(bytes32[] calldata keys, bytes32[] calldata values, uint8 typeId) external {
        require(keys.length == values.length, "Array length mismatch");
        require(typeId == TYPE_BYTES32, "Only bytes32 batch operations supported");

        for (uint256 i = 0; i < keys.length; i++) {
            bytes32 typedKey = _deriveKey(keys[i], typeId);
            bytes32Storage[msg.sender][typedKey] = values[i];
        }
    }

    /**
     * @notice Get the storage key for a given value, prefixed with its type
     * @dev Applies a type prefix to the highest byte of the key for type-based storage
     * @param baseKey The original key to derive from
     * @param typePrefix The type identifier to prefix the key with
     * @return A bytes32 key with the type prefix in the highest byte
     */
    function _deriveKey(bytes32 baseKey, uint8 typePrefix) internal pure returns (bytes32) {
        return bytes32(uint256(baseKey) | (uint256(typePrefix) << 248));
    }

    /**
     * @notice Extract the type identifier from a typed key
     * @dev Retrieves the highest byte from the key which represents the type
     * @param key The typed key to extract the type from
     * @return The type identifier as a uint8
     */
    function _extractType(bytes32 key) internal pure returns (uint8) {
        return uint8(uint256(key) >> 248);
    }

    /**
     * @notice Extract the original key without the type prefix
     * @dev Masks out the highest byte (type prefix) from the key
     * @param key The typed key to extract the base key from
     * @return The original key without the type prefix
     */
    function _extractBaseKey(bytes32 key) internal pure returns (bytes32) {
        return bytes32(
            uint256(key) & 0x00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
        );
    }
}
