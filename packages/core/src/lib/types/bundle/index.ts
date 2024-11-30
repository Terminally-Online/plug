import { BaseSchemaConfig, SchemaConfig } from "@/src/lib/types/schema"
import { BaseSolverConfig, SolverConfig } from "@/src/lib"

export * from "./evm"

export type BaseConfig = BaseSolverConfig & BaseSchemaConfig
export type Config = SolverConfig & SchemaConfig
