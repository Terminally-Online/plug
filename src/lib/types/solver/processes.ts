import { Collector, Executor } from "@/src/core/solver"

export type Collection<TCollector extends Collector<string, any>> = Parameters<
	TCollector["emit"]
>[1]

export type Execution<TExecutor extends Executor<`${string}Execution`, any>> =
	Parameters<TExecutor["execute"]>[0]
