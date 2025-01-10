# Lifecycle

Driving Plug exists a global lifecycle that is used to drive the simulation & execution of Plugs. Without this lifecycle, everything will stand still so this is the most important piece of the system to understand.

## The Lifecycle

At the highest level, the Solver operates with an internal clock that ticks every second. This clock drives a "cron-like" function that checks if any Plugs needs to be simulated (and executed). The core function is as follows:

1. Execution cron-job ticks.
2. Grabs Plugs that need to be simulated from app endpoint.
3. Build the transaction(s) that will be simulated.
4. Early return the simulations (transactions) that were not successfully built.
5. Simulate each of the transactions that were successfully built.
6. Run any transactions that were successfully simulated.
7. Return the results of the life cycle.

With this lifecycle running, we now have the ability to create Plugs, define their execution, the Solver cherry-picking and simulating the transaction it can build, running the transaction, and returning the results.

## Frequencies

The most important piece of the lifecycle is execution frequency defined on the "Queued Plug". With frequencies a user has the ability to set an "execution frequency" which defines the schedule on which an execution will be built, simulated, and run onchain.

When a simulation response is returned from the Solver, the next simulation is calculated based on the frequency that is defined as the "number of days" between runs.

- 0 [once]: Only allow running once and then stop future processing.
- 1 [daily]: Only allow running up to once every day.
- 7 [weekly]: Only allow running up to once every week.

When frequency is zero, the next simulation is set to null preventing any future executions. If the value is non-zero then the next simulation is calculated based on:

- `hour = 60 * 60 * 1000`
- `day = hour * 24`
- `next = now() + frequency * day`

---

frequency = 7
schedule = none
period start (January 1st)
period end (January 8th = January 1st + 7 days)
next simulation = period start

period expectations:

- 1st of January: 1st - 8th of January
- 8th of January: 8th - 15th of January
- 15th of January: 15th - 22nd of January
- 22nd of January: 22nd - 29th of January
- 29th of January: 29th - 5th of February

end = period start + frequency

when it is successful:
move period start to period end (start + frequency) when we have a successful transaction in that period

if we are moving the period start, the next simulation is at the same period start

if we are not moving it, the next simulation is now + workflow simulation frequency (this is a different frequency value which defines the rate of simulations)

---

Notably, this implementation will result in schedule drift. Right now we do not have a solution for this.

Let's say right now we schedule something on the 1st of January with a weekly frequency. A simple assumption that if it executes on the 5nd of January, the period would reset on the 8th of January as it's reached the new start of the period. However, in the current implementation, on the 5th the period would adjust to be the 5th - 12th of January.

---

The execution has two windows: global - represents the entire time the transaction is allowed to be run, local - frequency defined windows wherein the tx can occur (ie: this is allowed to occur once a week)

One solution is to track this local window on the execution.

When the execution is created on Jan 1 with a weekly frequency, the value nextWindow is set to Jan 7. On Jan 7, nextWindow is updated to nextWindow+frequency

When an execution is being evalatued for run, we must make sure that the local window is active by checking if now() < nextWindow

The nextSimulation datetime is simply updated to be every 10 minutes because we want to keep testing it until it works.

If a simulation is successful, we update the nextWindow and nextSimulation by adding the frequency to both. If it fails, we only update nextSimulation by adding 10 minutes.

## Building

## Simulations

The single most important piece of the simulation definition is the frequency.

## Executions (Runs)

### Frequencies

## Conditions

### Pause

### Stop (Completed)

if frequency > 0 && now() + frequency < endDate
then nextSimulation = now() + frequency
else nextSimulation is null ???
