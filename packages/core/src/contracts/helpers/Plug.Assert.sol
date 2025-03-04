// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugLogic } from "./Plug.Logic.sol";

/**
 * @title Plug Assert
 * @notice Plug Assert provides assertion functions that use PlugLogic for condition
 *         evaluation and revert with appropriate error messages when conditions are not met.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugAssert is PlugLogic {
    // ======== Boolean Assertions ========

    function assertTrue(bool value) external pure {
        require(isTrue(value), "PlugAssert:true-failed");
    }

    function assertFalse(bool value) external pure {
        require(isFalse(value), "PlugAssert:false-failed");
    }

    function assertAnd(bool a, bool b) external pure {
        require(isAnd(a, b), "PlugAssert:and-failed");
    }

    function assertOr(bool a, bool b) external pure {
        require(isOr(a, b), "PlugAssert:or-failed");
    }

    function assertNot(bool a) external pure {
        require(isNot(a), "PlugAssert:not-failed");
    }

    function assertXor(bool a, bool b) external pure {
        require(isXor(a, b), "PlugAssert:xor-failed");
    }

    function assertNand(bool a, bool b) external pure {
        require(isNand(a, b), "PlugAssert:nand-failed");
    }

    function assertNor(bool a, bool b) external pure {
        require(isNor(a, b), "PlugAssert:nor-failed");
    }

    function assertImplies(bool a, bool b) external pure {
        require(isImplies(a, b), "PlugAssert:implies-failed");
    }

    // ======== Number Assertions ========

    function assertGreaterThan(uint256 value, uint256 threshold) external pure {
        require(isGreaterThan(value, threshold), "PlugAssert:not-greater-than");
    }

    function assertGreaterThanOrEqual(uint256 value, uint256 threshold) external pure {
        require(isGreaterThanOrEqual(value, threshold), "PlugAssert:not-greater-than-or-equal");
    }

    function assertLessThan(uint256 value, uint256 threshold) external pure {
        require(isLessThan(value, threshold), "PlugAssert:not-less-than");
    }

    function assertLessThanOrEqual(uint256 value, uint256 threshold) external pure {
        require(isLessThanOrEqual(value, threshold), "PlugAssert:not-less-than-or-equal");
    }

    function assertEqual(uint256 value, uint256 threshold) external pure {
        require(isEqual(value, threshold), "PlugAssert:not-equal");
    }

    function assertNotEqual(uint256 value, uint256 threshold) external pure {
        require(isNotEqual(value, threshold), "PlugAssert:equal");
    }

    function assertBetween(uint256 value, uint256 min, uint256 max) external pure {
        require(isBetween(value, min, max), "PlugAssert:not-between");
    }

    function assertMultipleOf(uint256 value, uint256 divisor) external pure {
        require(isMultipleOf(value, divisor), "PlugAssert:not-multiple-of");
    }

    function assertPercentageOf(uint256 value, uint256 total, uint256 percentage) external pure {
        require(isPercentageOf(value, total, percentage), "PlugAssert:not-percentage-of");
    }

    // ======== String Assertions ========

    function assertStringsEqual(string calldata a, string calldata b) external pure {
        require(stringsEqual(a, b), "PlugAssert:strings-not-equal");
    }

    function assertStringsNotEqual(string calldata a, string calldata b) external pure {
        require(stringsNotEqual(a, b), "PlugAssert:strings-equal");
    }

    function assertStringContains(string calldata str, string calldata substr) external pure {
        require(stringContains(str, substr), "PlugAssert:string-does-not-contain");
    }

    function assertStringNotContains(string calldata str, string calldata substr) external pure {
        require(stringNotContains(str, substr), "PlugAssert:string-contains");
    }

    function assertStringStartsWith(string calldata str, string calldata prefix) external pure {
        require(stringStartsWith(str, prefix), "PlugAssert:string-does-not-start-with");
    }

    function assertStringEndsWith(string calldata str, string calldata suffix) external pure {
        require(stringEndsWith(str, suffix), "PlugAssert:string-does-not-end-with");
    }

    function assertStringIsEmpty(string calldata str) external pure {
        require(stringIsEmpty(str), "PlugAssert:string-not-empty");
    }

    function assertStringHasValue(string calldata str) external pure {
        require(stringHasValue(str), "PlugAssert:string-empty");
    }

    function assertStringIsNumeric(string calldata str) external pure {
        require(stringIsNumeric(str), "PlugAssert:string-not-numeric");
    }

    function assertStringLength(string calldata str, uint256 expected) external pure {
        require(stringLength(str, expected), "PlugAssert:string-length-mismatch");
    }

    function assertStringIsUppercase(string calldata str) external pure {
        require(stringIsUppercase(str), "PlugAssert:string-not-uppercase");
    }

    function assertStringIsLowercase(string calldata str) external pure {
        require(stringIsLowercase(str), "PlugAssert:string-not-lowercase");
    }

    // ======== Array Assertions ========

    function assertArrayContains(bytes32[] calldata arr, bytes32 value) external pure {
        require(arrayContains(arr, value), "PlugAssert:array-does-not-contain");
    }

    function assertArrayNotContains(bytes32[] calldata arr, bytes32 value) external pure {
        require(arrayNotContains(arr, value), "PlugAssert:array-contains");
    }

    function assertArrayIsEmpty(bytes32[] calldata arr) external pure {
        require(arrayIsEmpty(arr), "PlugAssert:array-not-empty");
    }

    function assertArrayHasValue(bytes32[] calldata arr) external pure {
        require(arrayHasValue(arr), "PlugAssert:array-empty");
    }
    
    function assertArrayLength(bytes32[] calldata arr, uint256 expected) external pure {
        require(arrayLength(arr, expected), "PlugAssert:array-length-mismatch");
    }

    function assertArrayContainsOnly(bytes32[] calldata arr, bytes32 value) external pure {
        require(arrayContainsOnly(arr, value), "PlugAssert:array-contains-other-values");
    }

    function assertArrayIsSorted(bytes32[] calldata arr) external pure {
        require(arrayIsSorted(arr), "PlugAssert:array-not-sorted");
    }

    function assertArrayHasUniqueElements(bytes32[] calldata arr) external pure {
        require(arrayHasUniqueElements(arr), "PlugAssert:array-has-duplicates");
    }

    // ======== Date & Time Assertions ========

    function assertBeforeTime(uint256 time, uint256 threshold) external pure {
        require(isBeforeTime(time, threshold), "PlugAssert:not-before-time");
    }

    function assertAfterTime(uint256 time, uint256 threshold) external pure {
        require(isAfterTime(time, threshold), "PlugAssert:not-after-time");
    }

    function assertBetweenTimes(uint256 time, uint256 start, uint256 end) external pure {
        require(isBetweenTimes(time, start, end), "PlugAssert:not-between-times");
    }

    function assertSameDay(uint256 timestamp1, uint256 timestamp2) external pure {
        require(isSameDay(timestamp1, timestamp2), "PlugAssert:not-same-day");
    }

    function assertWeekday(uint256 timestamp) external pure {
        require(isWeekday(timestamp), "PlugAssert:not-weekday");
    }

    function assertWeekend(uint256 timestamp) external pure {
        require(isWeekend(timestamp), "PlugAssert:not-weekend");
    }

    function assertLeapYear(uint256 year) external pure {
        require(isLeapYear(year), "PlugAssert:not-leap-year");
    }

    // ======== Advanced Assertions ========
    
    function assertAnyOf(bool[] calldata conditions) external pure {
        require(anyOf(conditions), "PlugAssert:none-true");
    }
    
    function assertAllOf(bool[] calldata conditions) external pure {
        require(allOf(conditions), "PlugAssert:not-all-true");
    }
    
    function assertNoneOf(bool[] calldata conditions) external pure {
        require(noneOf(conditions), "PlugAssert:some-true");
    }
    
    function assertExactlyOneOf(bool[] calldata conditions) external pure {
        require(exactlyOneOf(conditions), "PlugAssert:not-exactly-one-true");
    }
    
    function assertAtLeastN(bool[] calldata conditions, uint256 n) external pure {
        require(atLeastN(conditions, n), "PlugAssert:not-enough-true");
    }
    
    function assertAtMostN(bool[] calldata conditions, uint256 n) external pure {
        require(atMostN(conditions, n), "PlugAssert:too-many-true");
    }
}