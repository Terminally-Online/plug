---
head:
  - - meta
    - property: og:title
      content: Fuse | Schedule Windows
  - - meta
    - name: description
      content: Declare a schedule for which your intents can be executed on.
  - - meta
    - property: og:description
      content: Declare a schedule for which your intents can be executed on.
---

# Schedule Windows Fuse

With a schedule users unlock the ability to declare complex schedules for their transactions to be executed that are remarkably similar to Google Calendar.

## Logic

Any intent using the [Schedule Window Fuse](/instances/fuse/schedule-windows) has the ability to declare:

- Start Time: The starting time of the relative period.
- Repeats Every: How often the window of availability is repeated.
- Duration: How long the window remains open for.
- Days of Week: Which days of the week the schedule is active on.

With these four key pieces combined a user has granular control over the times in which an intent can be executed. During simulation and execution if the current time is not within the open window, the transaction will revert preventing the execution.
