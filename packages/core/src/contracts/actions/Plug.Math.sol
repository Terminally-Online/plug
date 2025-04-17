// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.26;

/**
 * @title Plug Math
 * @notice A collection of pure mathematical functions for use throughout the runtime
 *         execution of Plugs created by consumers and end-users of Plug.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugMath {
    /**
     * @notice Adds two numbers together
     * @param a First number
     * @param b Second number
     * @return result The sum of a and b
     */
    function add(int256 a, int256 b) public pure returns (int256 result) {
        return a + b;
    }

    /**
     * @notice Subtracts one number from another
     * @param a First number
     * @param b Number to subtract
     * @return result The difference (a - b)
     */
    function subtract(int256 a, int256 b) public pure returns (int256 result) {
        return a - b;
    }

    /**
     * @notice Multiplies two numbers
     * @param a First number
     * @param b Second number
     * @return result The product of a and b
     */
    function multiply(int256 a, int256 b) public pure returns (int256 result) {
        return a * b;
    }

    /**
     * @notice Divides one number by another
     * @param a Numerator
     * @param b Denominator
     * @return result The quotient (a / b)
     */
    function divide(int256 a, int256 b) public pure returns (int256 result) {
        require(b != 0, "PlugMath:divide-by-zero");
        return a / b;
    }

    /**
     * @notice Calculates the modulo of a number
     * @param a Dividend
     * @param b Divisor
     * @return result The remainder of a divided by b
     */
    function modulo(int256 a, int256 b) public pure returns (int256 result) {
        require(b != 0, "PlugMath:modulo-by-zero");
        return a % b;
    }

    /**
     * @notice Returns the minimum of two numbers
     * @param a First number
     * @param b Second number
     * @return result The smaller of a and b
     */
    function min(int256 a, int256 b) public pure returns (int256 result) {
        return a < b ? a : b;
    }

    /**
     * @notice Returns the maximum of two numbers
     * @param a First number
     * @param b Second number
     * @return result The larger of a and b
     */
    function max(int256 a, int256 b) public pure returns (int256 result) {
        return a > b ? a : b;
    }

    /**
     * @notice Calculates a number raised to a power
     * @param base The base number
     * @param exponent The exponent (must be non-negative)
     * @return result The base raised to the exponent
     */
    function power(
        int256 base,
        uint256 exponent
    )
        public
        pure
        returns (int256 result)
    {
        // Handle edge cases
        if (exponent == 0) return 1;
        if (base == 0) return 0;

        // For negative bases, handle even/odd exponents differently
        bool isNegative = base < 0;
        int256 absBase = isNegative ? -base : base;

        int256 temp = 1;
        for (uint256 i = 0; i < exponent; i++) {
            temp = temp * absBase;
        }

        // If base is negative and exponent is odd, result is negative
        return (isNegative && exponent % 2 == 1) ? -temp : temp;
    }

    /**
     * @notice Clamps a number between a minimum and maximum value
     * @param value The input value
     * @param minValue The minimum allowed value
     * @param maxValue The maximum allowed value
     * @return result The clamped value (between min and max)
     */
    function clamp(
        int256 value,
        int256 minValue,
        int256 maxValue
    )
        public
        pure
        returns (int256 result)
    {
        require(minValue <= maxValue, "PlugMath:min-exceeds-max");
        return min(max(value, minValue), maxValue);
    }
}
