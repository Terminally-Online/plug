// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugEtcherLib,
    PlugTypesLib,
    Plug,
    PlugFactory,
    PlugSocket,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";
import { LibBytes } from "solady/utils/LibBytes.sol";
import { console } from "forge-std/console.sol";

/**
 * @title PlugSocketLibBytesTest
 * @notice Tests for LibBytes optimizations in PlugSocket
 */
contract PlugSocketLibBytesTest is Test {
    bytes public testOriginal;
    bytes public testSlice;
    uint256 public testStartPos;

    function setUp() public virtual {
        setUpPlug();

        // Setup test data for byte operations
        testOriginal = bytes("Hello, World!");
        testSlice = bytes(" LibBytes");
        testStartPos = 5; // Insert after "Hello"
    }

    /**
     * @notice Test the byte slicing functionality using LibBytes
     */
    function test_LibBytes_Slice() public {
        // Test slicing the original data
        bytes memory sliced = LibBytes.slice(testOriginal, 0, 5); // "Hello"
        assertEq(string(sliced), "Hello");

        sliced = LibBytes.slice(testOriginal, 7, 12); // "World"
        assertEq(string(sliced), "World");
    }

    /**
     * @notice Test the byte concatenation functionality using LibBytes
     */
    function test_LibBytes_Concat() public {
        // Test concatenating two byte arrays
        bytes memory firstPart = LibBytes.slice(testOriginal, 0, 5); // "Hello"
        bytes memory secondPart = bytes(", LibBytes World!");

        bytes memory result = LibBytes.concat(firstPart, secondPart);
        assertEq(string(result), "Hello, LibBytes World!");
    }

    /**
     * @notice Test the insertion of a slice into original data using LibBytes
     * This simulates the optimized _insertSlice functionality
     */
    function test_LibBytes_InsertSlice() public {
        // Simulate the optimized _insertSlice operation
        bytes memory firstPart = LibBytes.slice(testOriginal, 0, testStartPos);
        bytes memory lastPart = LibBytes.slice(testOriginal, testStartPos, testOriginal.length);

        bytes memory result = LibBytes.concat(firstPart, testSlice);
        result = LibBytes.concat(result, lastPart);

        // Verify the result
        assertEq(string(result), "Hello LibBytes, World!");

        // Compare with expected size
        assertEq(result.length, testOriginal.length + testSlice.length);
    }

    /**
     * @notice Gas comparison test between the optimized and unoptimized versions
     * Note: This is more of a benchmark than an actual test
     */
    function test_LibBytes_GasComparison() public {
        bytes memory original = bytes("ThisIsALongerStringToTestGasEfficiency");
        bytes memory slice = bytes("_OPTIMIZED_");
        uint256 startPos = 10;

        // Optimized version using LibBytes
        uint256 gasBefore = gasleft();
        bytes memory optimizedResult = _insertSliceOptimized(original, startPos, slice);
        uint256 gasAfterOptimized = gasBefore - gasleft();

        // Legacy version using loops
        gasBefore = gasleft();
        bytes memory legacyResult = _insertSliceLegacy(original, startPos, slice);
        uint256 gasAfterLegacy = gasBefore - gasleft();

        // Verify both methods produce the same result
        assertEq(string(optimizedResult), string(legacyResult));

        // Log gas usage (will be visible in verbose test output)
        console.log("Gas used (optimized):", gasAfterOptimized);
        console.log("Gas used (legacy):", gasAfterLegacy);
        console.log(
            "Gas saved:",
            gasAfterLegacy > gasAfterOptimized ? gasAfterLegacy - gasAfterOptimized : 0
        );
    }

    /**
     * @notice Legacy implementation of _insertSlice using loops
     */
    function _insertSliceLegacy(
        bytes memory original,
        uint256 startPos,
        bytes memory slice
    )
        internal
        pure
        returns (bytes memory)
    {
        require(startPos <= original.length, "Would overflow");

        bytes memory result = new bytes(original.length + slice.length);

        // Copy first part
        for (uint256 i = 0; i < startPos; i++) {
            result[i] = original[i];
        }

        // Insert slice
        for (uint256 i = 0; i < slice.length; i++) {
            result[startPos + i] = slice[i];
        }

        // Copy last part
        for (uint256 i = startPos; i < original.length; i++) {
            result[i + slice.length] = original[i];
        }

        return result;
    }

    /**
     * @notice Optimized implementation of _insertSlice using LibBytes
     */
    function _insertSliceOptimized(
        bytes memory original,
        uint256 startPos,
        bytes memory slice
    )
        internal
        pure
        returns (bytes memory)
    {
        require(startPos <= original.length, "Would overflow");

        bytes memory firstPart = LibBytes.slice(original, 0, startPos);
        bytes memory lastPart = LibBytes.slice(original, startPos, original.length);

        bytes memory result = LibBytes.concat(firstPart, slice);
        result = LibBytes.concat(result, lastPart);

        return result;
    }
}
