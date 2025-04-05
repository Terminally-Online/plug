// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

/**
 * @title Plug Assert
 * @notice Plug Assert provides core assertion functions for validating boolean results
 *         from Plug.Boolean.sol or other logic operations.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugAssert {
    /**
     * @notice Assert that a condition is true
     * @param condition The boolean condition to assert
     */
    function assertTrue(bool condition) public pure {
        require(condition, "PlugAssert:condition-false");
    }

    /**
     * @notice Assert that a condition is false
     * @param condition The boolean condition to assert
     */
    function assertFalse(bool condition) public pure {
        require(!condition, "PlugAssert:condition-true");
    }

    /**
     * @notice Assert that a condition is true with a custom error message
     * @param condition The boolean condition to assert
     * @param message Custom error message to display if condition is false
     */
    function assertTrue(bool condition, string memory message) public pure {
        require(condition, string.concat("PlugAssert:", message));
    }

    /**
     * @notice Assert that a condition is false with a custom error message
     * @param condition The boolean condition to assert
     * @param message Custom error message to display if condition is true
     */
    function assertFalse(bool condition, string memory message) public pure {
        require(!condition, string.concat("PlugAssert:", message));
    }

    /**
     * @notice Unconditionally fail with a custom error message
     * @param message Custom error message to display
     */
    function fail(string memory message) public pure {
        revert(string.concat("PlugAssert:", message));
    }
}
