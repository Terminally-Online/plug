// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../interfaces/Plug.Connector.Interface.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";

// Scratchpad:
// if (
// 	($minute == type(uint8).max || $minute == currentMinute) &&
// 	($hour == type(uint8).max || $hour == currentHour) &&
// 	($day == type(uint8).max || $day == currentDay) &&
// 	($month == type(uint8).max || $month == currentMonth) &&
// 	($dayOfWeek == type(uint8).max || $dayOfWeek == currentDayOfWeek) ==
// 	false
// ) {
// 	revert PlugLib.ThresholdInvalid();
// }

/**
 * @title Plug Cron
 * @notice This Plug enables the ability to run actions on a recurring basis
 *         following the typical cron job pattern.
 * @notice Use cases for having an onchain cron job schedule:
 *     - Running a function on a declared schedule with:
 *      ┌───────────── minute (0–59)
 *      │ ┌───────────── hour (0–23)
 *      │ │ ┌───────────── day of the month (1–31)
 *      │ │ │ ┌───────────── month (1–12)
 *      │ │ │ │ ┌───────────── day of the week (0–6) (Sunday to Saturday;
 *      │ │ │ │ │                                   7 is also Sunday on some systems)
 *      │ │ │ │ │
 *      │ │ │ │ │
 *      * * * * * <command to execute>*
 * @dev To denote a state of general acceptance (*), the asterisk is used in common
 *      cron syntax however in this implementation, the asterisk is replaced with
 *      type(uint8).max to represent the maximum value of a uint8 which will never
 *      be reached in the context of a cron job.
 * @author nftchance (chance@onplug.io)
 */
contract PlugCron is PlugConnectorInterface {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32 $plugsHash) public virtual {
        $plugsHash;

        /// @dev Decode the terms definition of the cron job.
        (
            uint8 usesPerPeriod,
            uint24 tolerance,
            uint8 minute,
            uint8 hour,
            uint8 day,
            uint8 month,
            uint8 dayOfWeek
        ) = decode($terms);

        // TODO: Implement consumption limit.
        usesPerPeriod;

        tolerance;
        minute;
        hour;
        day;
        month;
        dayOfWeek;
    }

    /**
     * See {PlugConnectorInterface-decode}.
     */
    function decode(bytes calldata $terms)
        public
        pure
        returns (
            uint8 $usesPerPeriod,
            uint24 $tolerance,
            uint8 $minute,
            uint8 $hour,
            uint8 $day,
            uint8 $month,
            uint8 $dayOfWeek
        )
    {
        ($usesPerPeriod, $tolerance, $minute, $hour, $day, $month, $dayOfWeek) =
            abi.decode($terms, (uint8, uint24, uint8, uint8, uint8, uint8, uint8));
    }

    /**
     * See {PlugConnectorInterface-encode}.
     */
    function encode(
        uint8 $usesPerPeriod,
        uint24 $tolerance,
        uint8 $minute,
        uint8 $hour,
        uint8 $day,
        uint8 $month,
        uint8 $dayOfWeek
    )
        public
        pure
        returns (bytes memory $terms)
    {
        $terms = abi.encode($usesPerPeriod, $tolerance, $minute, $hour, $day, $month, $dayOfWeek);
    }

    /**
     * @notice Calculate year/month/day from the number of days since 1970/01/01 using
     *         the date conversion algorithm and adding the offset 2440588 so that
     *         1970/01/01 is day 0
     * @dev Fliegel and van Flandern (1968) published compact computer algorithms
     *      for converting between Julian dates and Gregorian calendar dates.
     *      Their algorithms were presented in the Fortran programming language,
     *      and take advantage of the truncation feature of integer arithmetic.
     * @return $month is the month, a number from 1 to 12;
     * @return $day is the day of the month, a number in the range 1-31;
     */
    function _daysToDate(uint256 $days) internal pure returns (uint256 $month, uint256 $day) {
        unchecked {
            int256 L = int256($days) + 2_509_157;
            int256 N = (4 * L) / 146_097;
            L = L - (146_097 * N + 3) / 4;

            int256 Y = (4000 * (L + 1)) / 1_461_001;
            L = L - (1461 * Y) / 4 + 31;

            int256 M = (80 * L) / 2447;
            L = M / 11;

            $month = uint256(M + 2 - 12 * L);
            $day = uint256(L - (2447 * M) / 80);
        }
    }

    /**
     * @notice Calculate the hour(s) of the time held in the timestamp.
     * @param $timestamp The timestamp to calculate the hour for.
     * @return $hour The hour of the time held in the timestamp.
     */
    function _hour(uint256 $timestamp) internal pure returns (uint256 $hour) {
        $hour = ($timestamp % 1 days) / 1 hours;
    }

    /**
     * @notice Calculate the minute(s) of the time held in the timestamp.
     * @param $timestamp The timestamp to calculate the minute for.
     * @return $minute The minute of the time held in the timestamp.
     */
    function _minute(uint256 $timestamp) internal pure returns (uint256 $minute) {
        $minute = ($timestamp % 1 hours) / 1 minutes;
    }

    /**
     * @notice Calculate the second(s) of the time held in the timestamp.
     * @param $timestamp The timestamp to calculate the second for.
     * @return $second The second of the time held in the timestamp.
     */
    function _second(uint256 $timestamp) internal pure returns (uint256 $second) {
        $second = $timestamp % 1 minutes;
    }

    /**
     * @notice Calculate the time for a given timestamp.
     * @param $timestamp The timestamp to calculate the time for.
     * @return $hour The hour of the time held in the timestamp.
     * @return $minute The minute of the time held in the timestamp.
     */
    function _time(uint256 $timestamp) internal pure returns (uint256 $hour, uint256 $minute) {
        $hour = _hour($timestamp);
        $minute = _minute($timestamp);
    }

    /**
     * @notice Calculate the month of the timestamp.
     * @param $timestamp The timestamp to calculate the month for.
     * @return $month The month of the timestamp.
     */
    function _month(uint256 $timestamp) internal pure returns (uint256 $month) {
        ($month,) = _daysToDate($timestamp / 1 days);
    }

    /**
     * @notice Calculate the day of the timestamp.
     * @param $timestamp The timestamp to calculate the day for.
     * @return $day The day of the timestamp.
     */
    function _day(uint256 $timestamp) internal pure returns (uint256 $day) {
        (, $day) = _daysToDate($timestamp / 1 days);
    }

    /**
     * @notice Calculate the date for a given timestamp.
     * @param $timestamp The timestamp to calculate the date for.
     * @return $month The month of the timestamp.
     * @return $day The day of the timestamp.
     */
    function _date(uint256 $timestamp) internal pure returns (uint256 $month, uint256 $day) {
        ($month, $day) = _daysToDate($timestamp / 1 days);
    }

    /**
     * @notice Calculate the day of the week for a given timestamp.
     * @dev Returning a numerical representation of the day of the week binds
     *      days to the range of: 0 (Monday) → 6 (Sunday)
     * @param $timestamp The timestamp to calculate the day of the week for.
     * @return $dow The day of the week for the given timestamp.
     */
    function _dayOfWeek(uint256 $timestamp) internal pure returns (uint256 $dow) {
        $dow = (($timestamp / 1 days) + 3) % 7;
    }

    /**
     * @notice Calculate the date and time for a given timestamp.
     * @param $timestamp The timestamp to calculate the date and time for.
     * @return $month The month of the timestamp.
     * @return $day The day of the timestamp.
     * @return $hour The hour of the timestamp.
     * @return $minute The minute of the timestamp.
     */
    function _datetime(uint256 $timestamp)
        internal
        pure
        returns (uint256 $month, uint256 $day, uint256 $hour, uint256 $minute)
    {
        ($month, $day) = _date($timestamp);
        ($hour, $minute) = _time($timestamp);
    }

    /**
     * @notice Calculate the date, time, and day of the week for a given timestamp.
     * @param $timestamp The timestamp to calculate the date, time, and day of the week for.
     * @return $month The month of the timestamp.
     * @return $day The day of the timestamp.
     * @return $hour The hour of the timestamp.
     * @return $minute The minute of the timestamp.
     * @return $dow The day of the week for the given timestamp.
     */
    function _cron(uint256 $timestamp)
        internal
        pure
        returns (uint256 $month, uint256 $day, uint256 $hour, uint256 $minute, uint256 $dow)
    {
        ($month, $day, $hour, $minute) = _datetime($timestamp);
        $dow = _dayOfWeek($timestamp);
    }
}
