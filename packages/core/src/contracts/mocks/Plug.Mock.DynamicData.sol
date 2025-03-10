// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

/**
 * @title Plug Mock Dynamic Data
 * @notice A mock contract for testing dynamic data handling in the Plug framework.
 * @dev This contract is for testing purposes only.
 */
contract PlugMockDynamicData {
    // Structure to test nested dynamic data
    struct ComplexData {
        uint256 id;
        string name;
        uint256[] values;
    }

    // Events to verify correct data was received
    event ArrayReceived(uint256[] values);
    event StringReceived(string value);
    event BytesReceived(bytes value);
    event StructReceived(uint256 id, string name, uint256[] values);
    event NestedArrayReceived(uint256[][] values);

    /**
     * @notice Returns a dynamic uint256 array
     * @param length The length of the array to return
     * @return values The array with 1 to length values
     */
    function returnUintArray(uint256 length) external pure returns (uint256[] memory values) {
        values = new uint256[](length);
        for (uint256 i = 0; i < length; i++) {
            values[i] = i + 1;
        }
    }

    /**
     * @notice Returns a dynamic address array
     * @param length The length of the array to return
     * @return addresses The array with addresses
     */
    function returnAddressArray(uint256 length)
        external
        view
        returns (address[] memory addresses)
    {
        addresses = new address[](length);
        for (uint256 i = 0; i < length; i++) {
            // Create some deterministic addresses
            addresses[i] = address(uint160(i + 1 + uint160(address(this))));
        }
    }

    /**
     * @notice Returns a string of specified length
     * @param text The string to return
     * @return The input string
     */
    function returnString(string calldata text) external pure returns (string memory) {
        return text;
    }

    /**
     * @notice Returns raw bytes data
     * @param data The bytes to return
     * @return The input bytes
     */
    function returnBytes(bytes calldata data) external pure returns (bytes memory) {
        return data;
    }

    /**
     * @notice Returns a complex struct with dynamic fields
     * @param id The id to use
     * @param name The name to use
     * @param arrayLength The length of the values array
     * @return result A ComplexData struct
     */
    function returnStruct(
        uint256 id,
        string calldata name,
        uint256 arrayLength
    )
        external
        pure
        returns (ComplexData memory result)
    {
        uint256[] memory values = new uint256[](arrayLength);
        for (uint256 i = 0; i < arrayLength; i++) {
            values[i] = i + 10;
        }

        result = ComplexData({ id: id, name: name, values: values });
    }

    /**
     * @notice Returns a nested array (array of arrays)
     * @param outerLength The length of the outer array
     * @param innerLength The length of each inner array
     * @return result A nested array
     */
    function returnNestedArray(
        uint256 outerLength,
        uint256 innerLength
    )
        external
        pure
        returns (uint256[][] memory result)
    {
        result = new uint256[][](outerLength);

        for (uint256 i = 0; i < outerLength; i++) {
            result[i] = new uint256[](innerLength);

            for (uint256 j = 0; j < innerLength; j++) {
                result[i][j] = i * 100 + j;
            }
        }
    }

    /**
     * @notice Takes a uint256 array and emits an event with it
     * @param values The array to process
     * @return The sum of all values in the array
     */
    function processUintArray(uint256[] calldata values) external returns (uint256) {
        emit ArrayReceived(values);

        uint256 sum = 0;
        for (uint256 i = 0; i < values.length; i++) {
            sum += values[i];
        }
        return sum;
    }

    /**
     * @notice Takes a string and emits an event with it
     * @param text The string to process
     * @return The length of the string
     */
    function processString(string calldata text) external returns (uint256) {
        emit StringReceived(text);
        return bytes(text).length;
    }

    /**
     * @notice Takes bytes data and emits an event with it
     * @param data The bytes to process
     * @return The length of the data
     */
    function processBytes(bytes calldata data) external returns (uint256) {
        emit BytesReceived(data);
        return data.length;
    }

    /**
     * @notice Takes a struct and emits an event with its components
     * @param data The struct to process
     * @return The id from the struct
     */
    function processStruct(ComplexData calldata data) external returns (uint256) {
        emit StructReceived(data.id, data.name, data.values);
        return data.id;
    }

    /**
     * @notice Takes a nested array and emits an event with it
     * @param values The nested array to process
     * @return The count of all elements
     */
    function processNestedArray(uint256[][] calldata values) external returns (uint256) {
        emit NestedArrayReceived(values);

        uint256 elementCount = 0;
        for (uint256 i = 0; i < values.length; i++) {
            elementCount += values[i].length;
        }
        return elementCount;
    }

    /**
     * @notice Concatenates two dynamic arrays
     * @param array1 The first array
     * @param array2 The second array
     * @return result The concatenated array
     */
    function concatenateArrays(
        uint256[] calldata array1,
        uint256[] calldata array2
    )
        external
        pure
        returns (uint256[] memory result)
    {
        result = new uint256[](array1.length + array2.length);

        for (uint256 i = 0; i < array1.length; i++) {
            result[i] = array1[i];
        }

        for (uint256 i = 0; i < array2.length; i++) {
            result[array1.length + i] = array2[i];
        }
    }

    /**
     * @notice Concatenates two strings
     * @param str1 The first string
     * @param str2 The second string
     * @return The concatenated string
     */
    function concatenateStrings(
        string calldata str1,
        string calldata str2
    )
        external
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(str1, str2));
    }
}
