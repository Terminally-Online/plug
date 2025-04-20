// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

/**
 * @title Plug Boolean
 * @notice Plug Boolean enables the use of factual and counter-factual state and
 *         value checks on the defined data of an arbitrary Ethereum transaction
 *         to be used in subsequent transactions or validity assertion checks.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugBoolean {
    /**
     * @notice Perform a logical NOT operation
     * @param a Boolean operand to negate
     * @return result True if a is false, false if a is true
     */
    function isNot(bool a) public pure returns (bool result) {
        result = !a;
    }

    /**
     * @notice Perform a logical AND operation
     * @param a First boolean operand
     * @param b Second boolean operand
     * @return result True if both a AND b are true
     */
    function isAnd(bool a, bool b) public pure returns (bool result) {
        result = a && b;
    }

    /**
     * @notice Perform a logical OR operation
     * @param a First boolean operand
     * @param b Second boolean operand
     * @return result True if either a OR b is true
     */
    function isOr(bool a, bool b) public pure returns (bool result) {
        result = a || b;
    }

    /**
     * @notice Perform a logical XOR (exclusive OR) operation
     * @param a First boolean operand
     * @param b Second boolean operand
     * @return result True if a and b have different values (one is true, one is false)
     */
    function isXor(bool a, bool b) public pure returns (bool result) {
        result = a != b;
    }

    /**
     * @notice Perform a logical NAND (NOT AND) operation
     * @param a First boolean operand
     * @param b Second boolean operand
     * @return result True if it's not the case that both a AND b are true
     */
    function isNand(bool a, bool b) public pure returns (bool result) {
        result = !(a && b);
    }

    /**
     * @notice Perform a logical NOR (NOT OR) operation
     * @param a First boolean operand
     * @param b Second boolean operand
     * @return result True if both a and b are false
     */
    function isNor(bool a, bool b) public pure returns (bool result) {
        result = !(a || b);
    }

    /**
     * @notice Perform a logical implication operation (a implies b)
     * @param a First boolean operand (the antecedent)
     * @param b Second boolean operand (the consequent)
     * @return result True if a implies b (either a is false or b is true)
     */
    function isImplies(bool a, bool b) public pure returns (bool result) {
        result = !a || b;
    }

    /**
     * @notice Check if two values are equal
     * @param value First value to compare
     * @param threshold Second value to compare
     * @return result True if the values are equal
     */
    function isEqual(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value == threshold;
    }

    /**
     * @notice Check if two values are not equal
     * @param value First value to compare
     * @param threshold Second value to compare
     * @return result True if the values are not equal
     */
    function isNotEqual(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value != threshold;
    }

    /**
     * @notice Check if a value is greater than a threshold
     * @param value The value to check
     * @param threshold The threshold to compare against
     * @return result True if value > threshold
     */
    function isGreaterThan(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value > threshold;
    }

    /**
     * @notice Check if a value is greater than or equal to a threshold
     * @param value The value to check
     * @param threshold The threshold to compare against
     * @return result True if value >= threshold
     */
    function isGreaterThanOrEqual(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value >= threshold;
    }

    /**
     * @notice Check if a value is less than a threshold
     * @param value The value to check
     * @param threshold The threshold to compare against
     * @return result True if value < threshold
     */
    function isLessThan(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value < threshold;
    }

    /**
     * @notice Check if a value is less than or equal to a threshold
     * @param value The value to check
     * @param threshold The threshold to compare against
     * @return result True if value <= threshold
     */
    function isLessThanOrEqual(
        uint256 value,
        uint256 threshold
    )
        public
        pure
        returns (bool result)
    {
        result = value <= threshold;
    }

    /**
     * @notice Check if a value is between a minimum and maximum (inclusive)
     * @param value The value to check
     * @param min The minimum threshold (inclusive)
     * @param max The maximum threshold (inclusive)
     * @return result True if min <= value <= max
     */
    function isBetween(
        uint256 value,
        uint256 min,
        uint256 max
    )
        public
        pure
        returns (bool result)
    {
        result = value >= min && value <= max;
    }
}
