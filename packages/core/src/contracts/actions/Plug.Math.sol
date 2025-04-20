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
     * @param x First number
     * @param y Second number
     * @return result The sum of x and y
     */
    function add(uint256 x, uint256 y) public pure returns (uint256 result) {
        return x + y;
    }

    /**
     * @notice Subtracts one number from another
     * @param x First number
     * @param y Number to subtract
     * @return result The difference (x - y)
     */
    function subtract(
        uint256 x,
        uint256 y
    )
        public
        pure
        returns (uint256 result)
    {
        return x - y;
    }

    /**
     * @notice Multiplies two numbers
     * @param x First number
     * @param y Second number
     * @return result The product of x and y
     */
    function multiply(
        uint256 x,
        uint256 y
    )
        public
        pure
        returns (uint256 result)
    {
        return x * y;
    }

    /**
     * @notice Divides one number by another
     * @param x Numerator
     * @param y Denominator
     * @return result The quotient (x / y)
     */
    function divide(
        uint256 x,
        uint256 y
    )
        public
        pure
        returns (uint256 result)
    {
        require(x != 0, "PlugMath:divide-by-zero");
        return x / y;
    }

    /**
     * @notice Calculates the modulo of a number
     * @param a Dividend
     * @param b Divisor
     * @return result The remainder of a divided by b
     */
    function modulo(
        uint256 a,
        uint256 b
    )
        public
        pure
        returns (uint256 result)
    {
        require(b != 0, "PlugMath:modulo-by-zero");
        return a % b;
    }

    /**
     * @notice Returns the minimum of two numbers
     * @param x First number
     * @param y Second number
     * @return result The smaller of a and b
     */
    function min(uint256 x, uint256 y) public pure returns (uint256 result) {
        return x < y ? x : y;
    }

    /**
     * @notice Returns the maximum of two numbers
     * @param x First number
     * @param y Second number
     * @return result The larger of x and y
     */
    function max(uint256 x, uint256 y) public pure returns (uint256 result) {
        return x > y ? x : y;
    }

    /**
     * @notice Calculates a number raised to a power
     * @param x The base number
     * @param y The exponent (must be non-negative)
     * @return result The x raised to the y
     */
    function power(uint256 x, uint256 y) public pure returns (uint256 result) {
        if (y == 0) return 1;
        if (x == 0) return 0;

        result = 1;
        for (uint256 i = 0; i < y; i++) {
            result = result * x;
        }

        return result;
    }
}
