// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { LibString } from "solady/utils/LibString.sol";
import { FixedPointMathLib } from "solady/utils/FixedPointMathLib.sol";
import { DateTimeLib } from "solady/utils/DateTimeLib.sol";
import { LibSort } from "solady/utils/LibSort.sol";
import { SafeCastLib } from "solady/utils/SafeCastLib.sol";

/**
 * @title Plug Logic
 * @notice Plug Logic enables the use of runtime logic inside a Plug with
 *         coil derived transaction data, state and expectations.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugLogic {
    using LibString for string;
    using FixedPointMathLib for uint256;
    using SafeCastLib for uint256;
    using LibSort for bytes32[];

    function isTrue(bool value) public pure returns (bool) {
        return value;
    }

    function isFalse(bool value) public pure returns (bool) {
        return !value;
    }

    function isAnd(bool a, bool b) public pure returns (bool) {
        return a && b;
    }

    function isOr(bool a, bool b) public pure returns (bool) {
        return a || b;
    }

    function isNot(bool a) public pure returns (bool) {
        return !a;
    }

    function isXor(bool a, bool b) public pure returns (bool) {
        return a != b;
    }

    function isNand(bool a, bool b) public pure returns (bool) {
        return !(a && b);
    }

    function isNor(bool a, bool b) public pure returns (bool) {
        return !(a || b);
    }

    function isImplies(bool a, bool b) public pure returns (bool) {
        return !a || b;
    }

    function isGreaterThan(uint256 value, uint256 threshold) public pure returns (bool) {
        return value > threshold;
    }

    function isGreaterThanOrEqual(uint256 value, uint256 threshold) public pure returns (bool) {
        return value >= threshold;
    }

    function isLessThan(uint256 value, uint256 threshold) public pure returns (bool) {
        return value < threshold;
    }

    function isLessThanOrEqual(uint256 value, uint256 threshold) public pure returns (bool) {
        return value <= threshold;
    }

    function isEqual(uint256 value, uint256 threshold) public pure returns (bool) {
        return value == threshold;
    }

    function isNotEqual(uint256 value, uint256 threshold) public pure returns (bool) {
        return value != threshold;
    }

    function isBetween(uint256 value, uint256 min, uint256 max) public pure returns (bool) {
        return value >= min && value <= max;
    }

    function isMultipleOf(uint256 value, uint256 divisor) public pure returns (bool) {
        if (divisor == 0) return false;
        return value % divisor == 0;
    }

    function isPercentageOf(uint256 value, uint256 total, uint256 percentage) public pure returns (bool) {
        if (total == 0) return false;
        uint256 expectedPercentage = value.mulDiv(100, total);
        return expectedPercentage == percentage;
    }

    function stringsEqual(string calldata a, string calldata b) public pure returns (bool) {
        return LibString.eq(a, b);
    }

    function stringsNotEqual(string calldata a, string calldata b) public pure returns (bool) {
        return !LibString.eq(a, b);
    }

    function stringContains(string calldata str, string calldata substr) public pure returns (bool) {
        if (bytes(substr).length == 0) return false;
        return LibString.indexOf(str, substr) != type(uint256).max;
    }

    function stringNotContains(string calldata str, string calldata substr) public pure returns (bool) {
        if (bytes(substr).length == 0) return false;
        return LibString.indexOf(str, substr) == type(uint256).max;
    }

    function stringStartsWith(string calldata str, string calldata prefix) public pure returns (bool) {
        if (bytes(prefix).length == 0) return false;
        return bytes(str).length >= bytes(prefix).length && LibString.indexOf(str, prefix) == 0;
    }

    function stringEndsWith(string calldata str, string calldata suffix) public pure returns (bool) {
        if (bytes(suffix).length == 0) return false;
        return bytes(str).length >= bytes(suffix).length && LibString.endsWith(str, suffix);
    }

    function stringIsEmpty(string calldata str) public pure returns (bool) {
        return bytes(str).length == 0;
    }

    function stringHasValue(string calldata str) public pure returns (bool) {
        return bytes(str).length > 0;
    }

    function stringIsNumeric(string calldata str) public pure returns (bool) {
        bytes memory b = bytes(str);
        if (b.length == 0) return false;

        for (uint i = 0; i < b.length; i++) {
            if (!((b[i] >= bytes1('0') && b[i] <= bytes1('9')) || (i == 0 && b[i] == bytes1('-')))) {
                return false;
            }
        }
        return true;
    }

    function stringLength(string calldata str, uint256 expected) public pure returns (bool) {
        return bytes(str).length == expected;
    }

    function stringIsUppercase(string calldata str) public pure returns (bool) {
        return LibString.eq(str, LibString.toCase(str, true));
    }

    function stringIsLowercase(string calldata str) public pure returns (bool) {
        return LibString.eq(str, LibString.toCase(str, false));
    }

    function arrayContains(bytes32[] calldata arr, bytes32 value) public pure returns (bool) {
        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] == value) {
                return true;
            }
        }
        return false;
    }

    function arrayNotContains(bytes32[] calldata arr, bytes32 value) public pure returns (bool) {
        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] == value) {
                return false;
            }
        }
        return true;
    }

    function arrayIsEmpty(bytes32[] calldata arr) public pure returns (bool) {
        return arr.length == 0;
    }

    function arrayHasValue(bytes32[] calldata arr) public pure returns (bool) {
        return arr.length > 0;
    }

    function arrayLength(bytes32[] calldata arr, uint256 expected) public pure returns (bool) {
        return arr.length == expected;
    }

    function arrayContainsOnly(bytes32[] calldata arr, bytes32 value) public pure returns (bool) {
        if (arr.length == 0) return false;

        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] != value) {
                return false;
            }
        }
        return true;
    }

    function arrayIsSorted(bytes32[] calldata arr) public pure returns (bool) {
        if (arr.length <= 1) return true;

        for (uint i = 1; i < arr.length; i++) {
            if (arr[i] < arr[i-1]) {
                return false;
            }
        }
        return true;
    }

    function arrayHasUniqueElements(bytes32[] calldata arr) public pure returns (bool) {
        if (arr.length == 0) return true;

        for (uint i = 0; i < arr.length; i++) {
            for (uint j = i + 1; j < arr.length; j++) {
                if (arr[i] == arr[j]) {
                    return false;
                }
            }
        }
        return true;
    }

    function isBeforeTime(uint256 time, uint256 threshold) public pure returns (bool) {
        return time < threshold;
    }

    function isAfterTime(uint256 time, uint256 threshold) public pure returns (bool) {
        return time > threshold;
    }

    function isBetweenTimes(uint256 time, uint256 start, uint256 end) public pure returns (bool) {
        return time >= start && time <= end;
    }

    function isSameDay(uint256 timestamp1, uint256 timestamp2) public pure returns (bool) {
        (uint256 year1, uint256 month1, uint256 day1) = DateTimeLib.timestampToDate(timestamp1);
        (uint256 year2, uint256 month2, uint256 day2) = DateTimeLib.timestampToDate(timestamp2);

        return year1 == year2 && month1 == month2 && day1 == day2;
    }

    function isWeekday(uint256 timestamp) public pure returns (bool) {
        return !DateTimeLib.isWeekEnd(timestamp);
    }

    function isWeekend(uint256 timestamp) public pure returns (bool) {
        return DateTimeLib.isWeekEnd(timestamp);
    }

    function isLeapYear(uint256 year) public pure returns (bool) {
        if (year > type(uint16).max) return false;
        return DateTimeLib.isLeapYear(SafeCastLib.toUint16(year));
    }

    function any(bool[] calldata conditions) public pure returns (bool) {
        for (uint i = 0; i < conditions.length; i++) {
            if (conditions[i]) {
                return true;
            }
        }
        return false;
    }

    function all(bool[] calldata conditions) public pure returns (bool) {
        for (uint i = 0; i < conditions.length; i++) {
            if (!conditions[i]) {
                return false;
            }
        }
        return true;
    }

    function noneOf(bool[] calldata conditions) public pure returns (bool) {
        for (uint i = 0; i < conditions.length; i++) {
            if (conditions[i]) {
                return false;
            }
        }
        return true;
    }

    function one(bool[] calldata conditions) public pure returns (bool) {
        uint count = 0;
        for (uint i = 0; i < conditions.length; i++) {
            if (conditions[i]) count++;
            if (count > 1) return false;
        }
        return count == 1;
    }

    function minimumThreshold(bool[] calldata conditions, uint256 n) public pure returns (bool) {
        if (n == 0) return true;
        if (n > conditions.length) return false;

        uint count = 0;
        for (uint i = 0; i < conditions.length; i++) {
            if (conditions[i]) count++;
            if (count >= n) return true;
        }
        return false;
    }

    function maximumThreshold(bool[] calldata conditions, uint256 n) public pure returns (bool) {
        if (n >= conditions.length) return true;

        uint count = 0;
        for (uint i = 0; i < conditions.length; i++) {
            if (conditions[i]) count++;
            if (count > n) return false;
        }
        return true;
    }
}
