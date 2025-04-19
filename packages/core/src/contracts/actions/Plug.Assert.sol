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
}
