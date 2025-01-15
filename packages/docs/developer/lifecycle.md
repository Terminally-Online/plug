# Execution Lifecycle

Driving Plug exists a global lifecycle that is used to drive the building, simulating, and onchain execution of Plugs. Without this lifecycle, everything will stand still so this is the most important piece of the system to understand.

## The Flow

The lifecycle pattern to execute a Plug is as follows:

1. The Solver ticks every second.
2. Grabs Plugs that need to be simulated from the app endpoint.
3. Build the transaction(s) that will be simulated.
4. Early return the simulations (transactions) that were not successfully built.
5. Simulate each of the transactions that were successfully built.
6. Run any transactions that were successfully simulated.
7. Return the results of the life cycle.

With this lifecycle running, we now have the ability to create Plugs, define their actions, queued a one-time or scheduled execution, the Solver building and simulating a transaction, and returning the results to the application that users use.

## One-Time Executions

In the application there exists the ability to both "run" and "schedule" an execution:

- A run execution has no [frequency](#frequencies) and no [schedule](#schedules) set by the user.
- A scheduled execution has a [frequency](#frequencies) and a [schedule](#schedules) set by the user.

For one-time executions, the Solver will simulate the transaction once and attempt to build the transaction. If building and simulating is successful, the transaction will be executed. If the transaction is not successfully built or fails somewhere along the way, the Solver will return an error and the execution will be marked as failed without any retries to follow.

## Frequencies

As a complete concept [frequencies](#frequencies) control a majority of the time-based logic in the Solver and supporting systems to ensure that simulations and executions are processed at the correct times.

### Workflow

A [workflow frequency](#workflow) defines the rate of simulations. During Beta, this is set to 10 minutes by default for every workflow regardless of the [execution frequency](#execution).

Notably, today Plug does not use a "listener" system to trigger the simulations. Instead, it uses a cron-like job that ticks every second and pulls in new simulation targets from the application every 1 minute. This means that while you have a queued execution and a transaction has not run, the Solver will simulate (and attempt to execute the transaction if fit) as often as every 10 minutes without the user doing anything.

This is helpful in the case that an execution has been queued and the user wants to run it as soon as the conditions are met. Now, the simulation can fail to resolve several times and still execute once the chain state has changed in a way that enables the transaction to be built, simulated, and executed.

### Execution

A second form of frequencies the system uses is [execution frequencies](#execution) that define how often a transaction _could_ be run. A user may want a transaction to only run once, daily, or even monthly.

As we are using a cron-like system it is important to mention that in a typical cron-like system there would be the case of time drift because the system uses a notional definition of when it is time to run the job.

Imagine the case where a user defines a weekly schedule and the transaction runs on the 6th day of the 7 day week (period).

Does the user expect the window to slide to now and start the next 7 days? Of course not, because we already had a period defined. The period should slide over to the end of the current period and simulations should resume at the start of the next period.

![Execution Frequency Logic](/public/assets/execution-frequency.png)

To avoid the mentioned time drift issue we use a sliding window system to manage our simulation periods.

With a sliding window a transaction may only execute once within each period. If it executes in a period and there is the opportunity to have future periods, we slide the period over and set the next simulate date to the current end of the period (the new start of the next period). If there is no future period, we set the next simulate date to null and stop any future processing of the given execution.

## Schedules

A third layer of time control exists in the form of [execution schedules](#schedules). These schedules define the start and end dates of an execution irrespective of the [execution frequency](#execution) or [workflow frequency](#workflow).

![Execution Schedules](/public/assets/execution-schedules.png)

This is useful in the case where a user wants to queue an execution that should only run between a certain date range. With a schedule in place, it functions as a global container of the [period system](#execution) that puts maximum bounds on the execution

Even if the periods "should" continue based on the frequency if there is a schedule in place that has been concluded, the execution will not continue.
