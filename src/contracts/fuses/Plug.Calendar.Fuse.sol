// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    PlugFuseInterface,
    PlugTypesLib
} from "../interfaces/Plug.Fuse.Interface.sol";

library CalendarFuseLib {
    error CalendarLackingDuration();
    error CalendarLackingDays();
    error CalendarLackingRepeatsEvery();
    error CalendarLackingSufficientRepeatsEvery();
    error CalendarLackingStartTime();
    error CalendarLackingN();
    error CalendarLackingHorizon();
    error CalendarCaveatViolation();

    struct Period {
        uint32 startTime;
        uint32 endTime;
    }

    struct Calendar {
        Period[] periods;
    }
}

/**
 * @title Calendar Caveat
 * @notice This Fuse is responsible for providing a function to check
 *         whether or not the current time is within a given calendar of time
 *         that repeats every X seconds and lasts for Y seconds as well as
 *         helper functions needed to determine the next N calendar openings.
 * @dev When working with this Fuse, you will generate timestamps with:
 *    - Javascript / Typescript:
 *      - `const dateInSecs = Math.floor(new Date().getTime() / 1000);`
 *    - Rust:
 *      - `let date_in_secs = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs();`
 *    - Python:
 *      - `date_in_secs = int(time.time())`
 *    - Solidity:
 *      - `uint256 dateInSecs = block.timestamp;`
 * @author @nftchance (chance@onplug.io)
 */
contract PlugCalendarFuse is PlugFuseInterface {
    /// @dev The number of seconds in a day.
    uint32 private constant SECONDS_PER_DAY = 1 days;

    /// @dev Bits to shift over the start time in a schedule.
    uint256 private constant START_TIME_SHIFT = 72;
    /// @dev Bits to shift over the repeats every in a schedule.
    uint256 private constant REPEATS_EVERY_SHIFT = 40;
    /// @dev Bits to shift over the duration in a schedule.
    uint256 private constant DURATION_SHIFT = 8;
    /// @dev Bits to shift over the days of the week in a schedule.
    uint256 private constant DAYS_OF_WEEK_SHIFT = 0;

    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32
    )
        public
        view
        override
        returns (bytes memory $through)
    {
        uint256 schedule = abi.decode($live, (uint256));

        if (!isWithinCalendar(schedule)) {
            revert CalendarFuseLib.CalendarCaveatViolation();
        }

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @dev Unpack the schedule details from the schedule.
     * @param $schedule The schedule to unpack.
     * @return $startTime The time the schedule was declared.
     * @return $repeatsEvery The number of seconds between each calendar.
     * @return $duration The number of seconds each calendar lasts.
     * @return $daysOfWeek The days of the week the calendar is active.
     */
    function decode(uint256 $schedule)
        public
        pure
        returns (
            uint32 $startTime,
            uint32 $repeatsEvery,
            uint32 $duration,
            uint8 $daysOfWeek
        )
    {
        /// @dev Unpack the schedule details.
        $daysOfWeek = uint8($schedule);
        $duration = uint32($schedule >> DURATION_SHIFT);
        $repeatsEvery = uint32($schedule >> REPEATS_EVERY_SHIFT);
        $startTime = uint32($schedule >> START_TIME_SHIFT);
    }

    /**
     * @dev Pack the schedule details into a schedule.
     * @param $startTime The time the schedule was declared.
     * @param $repeatsEvery The number of seconds between each calendar.
     * @param $duration The number of seconds each calendar lasts.
     * @return $schedule The packed schedule.
     */
    function encode(
        uint32 $startTime,
        uint32 $repeatsEvery,
        uint32 $duration,
        uint8 $daysOfWeek
    )
        external
        pure
        returns (uint256 $schedule)
    {
        /// @dev Ensure the duration is greater than 0.
        if ($duration == 0) {
            revert CalendarFuseLib.CalendarLackingDuration();
        }

        /// @dev Must support at least one day.
        if ($daysOfWeek == 0) {
            revert CalendarFuseLib.CalendarLackingDays();
        }

        /// @dev Prevent weird overlapping calendars.
        if ($duration > $repeatsEvery) {
            revert
                CalendarFuseLib
                .CalendarLackingSufficientRepeatsEvery();
        }

        /// @dev Pack the schedule details.
        $schedule = (uint256($startTime) << START_TIME_SHIFT)
            | (uint256($repeatsEvery) << REPEATS_EVERY_SHIFT)
            | (uint256($duration) << DURATION_SHIFT)
            | uint256($daysOfWeek);
    }

    /**
     * @notice Check whether or not the current time is within the
     *         the declared calendar of availability by the schedule.
     * @param $schedule The schedule to check.
     * @return $isWithinCalendar Whether or not the current time is within the
     */
    function isWithinCalendar(uint256 $schedule)
        public
        view
        returns (bool)
    {
        /// @dev Get the schedule details.
        (
            uint32 startTime,
            uint32 repeatsEvery,
            uint32 duration,
            uint8 daysOfWeek
        ) = decode($schedule);

        /// @dev Ensure the current time is within the calendar.
        return _isWithinCalendar(
            startTime, repeatsEvery, duration, daysOfWeek
        );
    }

    /**
     * @dev Overloads the {isWithinCalendar} function to allow for a verbose call.
     * @param $startTime The time the schedule was declared.
     * @param $repeatsEvery The number of seconds between each calendar.
     * @param $duration The number of seconds each calendar lasts.
     */
    function isWithinCalendar(
        uint32 $startTime,
        uint32 $repeatsEvery,
        uint32 $duration,
        uint8 $daysOfWeek
    )
        external
        view
        returns (bool)
    {
        /// @dev Ensure the current time is within the calendar.
        return _isWithinCalendar(
            $startTime, $repeatsEvery, $duration, $daysOfWeek
        );
    }

    /**
     * @dev Determine the next N calendar openings for a given schedule.
     * @notice If you call this onchain you have sinned and you will not be forgiven.
     *         This is simply a utility function to help you determine and/or
     *         visualize the Openings of your schedule.
     * @param $schedule The schedule to check.
     * @param $n The number of calendar openings to return.
     * @return $calendars The next N calendar openings.
     */
    function toCalendars(
        uint256 $schedule,
        uint32 $n
    )
        external
        view
        returns (
            CalendarFuseLib.Calendar[] memory $calendars,
            uint32 $cursor
        )
    {
        /// @dev Load the stack.
        $calendars = new CalendarFuseLib.Calendar[]($n);

        /// @dev Get the schedule details.
        (
            uint32 startTime,
            uint32 repeatsEvery,
            uint32 duration,
            uint8 daysOfWeek
        ) = decode($schedule);

        /// @dev Calculate the cursor used to get the next batch of results
        ///      after the return of the requested batch.
        $cursor = uint32(block.timestamp) + $n * repeatsEvery;

        for (startTime; startTime < $cursor;) {
            /// @dev Add the next calendar to the list of calendars.
            $calendars[(startTime / repeatsEvery) % $n] =
                _toCalendar(startTime, duration, daysOfWeek);

            /// @dev Time travel into the future.
            unchecked {
                startTime += repeatsEvery;
            }
        }
    }

    /**
     * @dev Determine the active periods for a schedule calendar given a horizon
     *      to filter to the points at which the `daysOfWeek` condition
     *      is satisfied as a Calendar may contain multiple periods in which
     *      it can be settled.
     * @param $schedule The schedule to check.
     */
    function toCalendar(uint256 $schedule)
        external
        pure
        returns (CalendarFuseLib.Calendar memory $calendar)
    {
        /// @dev Get the schedule details.
        (uint32 startTime, uint32 repeatsEvery,, uint8 daysOfWeek) =
            decode($schedule);

        /// @dev Ensure the current time is within the calendar.
        return _toCalendar(startTime, repeatsEvery, daysOfWeek);
    }

    /**
     * @dev Overloaded version of {toCalendar} to allow for a verbose call.
     * @param $startTime The time the schedule was declared.
     * @param $duration The number of seconds each calendar lasts.
     * @param $daysOfWeek The days of the week the calendar is active.
     */
    function toCalendar(
        uint32 $startTime,
        uint32 $duration,
        uint8 $daysOfWeek
    )
        external
        pure
        returns (CalendarFuseLib.Calendar memory $calendar)
    {
        /// @dev Ensure the current time is within the calendar.
        return _toCalendar($startTime, $duration, $daysOfWeek);
    }

    /**
     * @dev Check whether or not the current time is on a given day of the week.
     * @param $daysOfWeek The days of the week to check.
     * @param $timestamp The timestamp to check.
     */
    function _isOnDayOfWeek(
        uint8 $daysOfWeek,
        uint32 $timestamp
    )
        internal
        pure
        returns (bool)
    {
        /// @dev Get the day of the week.
        uint8 dayOfWeek =
            uint8((($timestamp / SECONDS_PER_DAY) + 4) % 7);

        /// @dev Check if the day of the week is supported.
        return ($daysOfWeek >> dayOfWeek) & 1 == 1;
    }

    /**
     * @dev Determine the active periods for a schedule calendar given a horizon
     *      to filter to the points at which the `daysOfWeek` condition
     *      is satisfied as a Calendar may contain multiple periods in which
     *      it can be settled.
     * @param $startTime The time the schedule was declared.
     * @param $duration The number of seconds each calendar lasts.
     * @param $daysOfWeek The days of the week the calendar is active.
     * @return $calendar The active periods for the schedule calendar.
     */
    function _toCalendar(
        uint32 $startTime,
        uint32 $duration,
        uint8 $daysOfWeek
    )
        internal
        pure
        returns (CalendarFuseLib.Calendar memory $calendar)
    {
        /// @dev Calculate when this Calendar ends.
        uint32 calendarEndTime = $startTime + $duration;

        /// @dev Calculate the maximum number of days this intent may extend.
        uint32 daysInCalendar = $duration / SECONDS_PER_DAY;

        /// @dev Load the stack.
        $calendar.periods =
            new CalendarFuseLib.Period[](daysInCalendar);

        /// @dev Loop through every day in the calendar backwards.
        for (daysInCalendar; daysInCalendar >= 0; daysInCalendar--) {
            /// @dev Get the time `daysInCalendar` days after the start time.
            uint32 dayTime =
                $startTime + daysInCalendar * SECONDS_PER_DAY;

            /// @dev Day time will be the 24 hour increment of the start time,
            ///      however the day calculations roll by the start of the day.
            ///      Thus, we want to place the top of the period with no surplus.
            uint32 topDayTime = dayTime - (dayTime % SECONDS_PER_DAY);

            /// @dev Some Calendars may be longer than 1 day so we must check if
            ///      the period is active today.
            bool isOnDayOfWeek =
                _isOnDayOfWeek($daysOfWeek, topDayTime);

            /// @dev If it is not active there is no active period.
            /// @dev Do realize if the start time is not on a day of the week
            ///      that is allowed, then the period of the opening day will
            ///      be 0 seconds long.
            if (!isOnDayOfWeek) continue;

            /// @dev Calculate the start of this period by determining if we have
            ///      gone past the declared start time otherwise it started
            ///      at the top of the day.
            $calendar.periods[daysInCalendar].startTime =
                $startTime > topDayTime ? $startTime : topDayTime;

            /// @dev Calculate the last second of the day.
            uint32 bottomDayTime = topDayTime + SECONDS_PER_DAY - 1;

            /// @dev Calculate the end of this period by determining if we have
            ///      gone past the declared end time otherwise it ended
            ///      at the bottom of the day.
            $calendar.periods[daysInCalendar].endTime =
            calendarEndTime < bottomDayTime
                ? calendarEndTime
                : bottomDayTime;
        }
    }

    /**
     * @dev Check whether or not the current time is within a given calendar of time
     *      that repeats every X seconds and lasts for Y seconds.
     * @param $startTime The time the schedule was declared.
     * @param $repeatsEvery The number of seconds between each calendar.
     * @param $duration The number of seconds each calendar lasts.
     */
    function _isWithinCalendar(
        uint32 $startTime,
        uint32 $repeatsEvery,
        uint32 $duration,
        uint8 $daysOfWeek
    )
        internal
        view
        returns (bool)
    {
        /// @dev Get the current time.
        uint32 currentTime = uint32(block.timestamp);

        /// @dev Ensure the current time is after the start time.
        if (currentTime < $startTime) {
            revert CalendarFuseLib.CalendarLackingStartTime();
        }

        /// @dev Ensure the current time is on a supported day of the week.
        if (!_isOnDayOfWeek($daysOfWeek, currentTime)) {
            revert CalendarFuseLib.CalendarLackingDays();
        }

        /// @dev Get the time since the declaration of the schedule.
        uint32 timeElapsed = currentTime - $startTime;

        /// @notice $repeatsEvery may be zero for a one-time calendar
        ///         that does not repeat so we must check for this.
        if ($repeatsEvery == 0) {
            return currentTime < $startTime + $duration;
        }

        /// @dev Get the time since the start of the current calendar.
        uint32 currentCalendarOpen =
            $startTime + (timeElapsed / $repeatsEvery) * $repeatsEvery;

        /// @dev Ensure the current time is within the current calendar.
        return currentTime >= currentCalendarOpen
            && currentTime < currentCalendarOpen + $duration;
    }
}
