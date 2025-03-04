// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.23;

/**
 * @title Plug Coercion
 * @notice A collection of type conversion functions for the Plug ecosystem
 *         that allow for dynamic type coercion during runtime execution.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugCoercion {
    /**
     * @notice Cast a uint256 to a uint160, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint160
     */
    function toUint160(uint256 y) public pure returns (uint160 z) {
        require((z = uint160(y)) == y, "PlugCoercion:uint160-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint128, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint128
     */
    function toUint128(uint256 y) public pure returns (uint128 z) {
        require((z = uint128(y)) == y, "PlugCoercion:uint128-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint96, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint96
     */
    function toUint96(uint256 y) public pure returns (uint96 z) {
        require((z = uint96(y)) == y, "PlugCoercion:uint96-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint64, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint64
     */
    function toUint64(uint256 y) public pure returns (uint64 z) {
        require((z = uint64(y)) == y, "PlugCoercion:uint64-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint32, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint32
     */
    function toUint32(uint256 y) public pure returns (uint32 z) {
        require((z = uint32(y)) == y, "PlugCoercion:uint32-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint16, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint16
     */
    function toUint16(uint256 y) public pure returns (uint16 z) {
        require((z = uint16(y)) == y, "PlugCoercion:uint16-overflow");
    }

    /**
     * @notice Cast a uint256 to a uint8, revert on overflow
     * @param y The uint256 to be downcasted
     * @return z The downcasted integer, now type uint8
     */
    function toUint8(uint256 y) public pure returns (uint8 z) {
        require((z = uint8(y)) == y, "PlugCoercion:uint8-overflow");
    }

    /**
     * @notice Cast an int256 to an int128, revert on overflow or underflow
     * @param y The int256 to be downcasted
     * @return z The downcasted integer, now type int128
     */
    function toInt128(int256 y) public pure returns (int128 z) {
        require((z = int128(y)) == y, "PlugCoercion:int128-overflow");
    }

    /**
     * @notice Cast an int256 to an int64, revert on overflow or underflow
     * @param y The int256 to be downcasted
     * @return z The downcasted integer, now type int64
     */
    function toInt64(int256 y) public pure returns (int64 z) {
        require((z = int64(y)) == y, "PlugCoercion:int64-overflow");
    }

    /**
     * @notice Cast an int256 to an int32, revert on overflow or underflow
     * @param y The int256 to be downcasted
     * @return z The downcasted integer, now type int32
     */
    function toInt32(int256 y) public pure returns (int32 z) {
        require((z = int32(y)) == y, "PlugCoercion:int32-overflow");
    }

    /**
     * @notice Cast an int256 to an int16, revert on overflow or underflow
     * @param y The int256 to be downcasted
     * @return z The downcasted integer, now type int16
     */
    function toInt16(int256 y) public pure returns (int16 z) {
        require((z = int16(y)) == y, "PlugCoercion:int16-overflow");
    }

    /**
     * @notice Cast an int256 to an int8, revert on overflow or underflow
     * @param y The int256 to be downcasted
     * @return z The downcasted integer, now type int8
     */
    function toInt8(int256 y) public pure returns (int8 z) {
        require((z = int8(y)) == y, "PlugCoercion:int8-overflow");
    }

    /**
     * @notice Converts an unsigned integer to a signed integer
     * @param y The uint256 value to convert
     * @return z The converted int256 value
     */
    function toInt256(uint256 y) public pure returns (int256 z) {
        require(y <= uint256(type(int256).max), "PlugCoercion:int256-overflow");
        z = int256(y);
    }

    /**
     * @notice Converts a signed integer to an unsigned integer
     * @param y The int256 value to convert
     * @return z The converted uint256 value
     */
    function toUint256(int256 y) public pure returns (uint256 z) {
        require(y >= 0, "PlugCoercion:negative-uint");
        z = uint256(y);
    }

    /**
     * @notice Converts an integer to bytes32
     * @param y The uint256 value to convert
     * @return z The bytes32 representation
     */
    function uintToBytes32(uint256 y) public pure returns (bytes32 z) {
        z = bytes32(y);
    }

    /**
     * @notice Converts bytes32 to an integer
     * @param y The bytes32 value to convert
     * @return z The uint256 representation
     */
    function bytes32ToUint(bytes32 y) public pure returns (uint256 z) {
        z = uint256(y);
    }

    /**
     * @notice Converts address to uint160
     * @param y The address to convert
     * @return z The uint160 representation
     */
    function addressToUint160(address y) public pure returns (uint160 z) {
        z = uint160(y);
    }

    /**
     * @notice Converts uint160 to address
     * @param y The uint160 to convert
     * @return z The address representation
     */
    function uint160ToAddress(uint160 y) public pure returns (address z) {
        z = address(y);
    }

    /**
     * @notice Converts bytes to bytes32, padding with zeros if necessary
     * @param y The bytes to convert
     * @return z The bytes32 representation
     */
    function bytesToBytes32(bytes memory y) public pure returns (bytes32 z) {
        require(y.length <= 32, "PlugCoercion:bytes-too-long");

        assembly {
            z := mload(add(y, 32))
        }
    }

    /**
     * @notice Converts bytes32 to bytes
     * @param y The bytes32 to convert
     * @return z The bytes representation
     */
    function bytes32ToBytes(bytes32 y) public pure returns (bytes memory z) {
        z = new bytes(32);

        assembly {
            mstore(add(z, 32), y)
        }
    }

    /**
     * @notice Converts a boolean to uint256 (1 for true, 0 for false)
     * @param y The boolean value to convert
     * @return z The uint256 representation
     */
    function boolToUint(bool y) public pure returns (uint256 z) {
        return y ? 1 : 0;
    }

    /**
     * @notice Converts uint256 to boolean (true for non-zero, false for zero)
     * @param y The uint256 value to convert
     * @return z The boolean representation
     */
    function uintToBool(uint256 y) public pure returns (bool z) {
        return y != 0;
    }

    /**
     * @notice Converts a boolean to int256 (1 for true, 0 for false)
     * @param y The boolean value to convert
     * @return z The int256 representation
     */
    function boolToInt(bool y) public pure returns (int256 z) {
        return y ? int256(1) : int256(0);
    }

    /**
     * @notice Converts int256 to boolean (true for non-zero, false for zero)
     * @param y The int256 value to convert
     * @return z The boolean representation
     */
    function intToBool(int256 y) public pure returns (bool z) {
        return y != 0;
    }
}
