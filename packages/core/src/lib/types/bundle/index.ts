import { DEFAULT_NETWORKS, BaseSchemaConfig, Network, Retries, SchemaConfig } from "@/src/lib"

export * from "./evm"
export * from "./network"

export type BaseConfig = BaseSchemaConfig & Partial<Retries> & (
    | Record<"networks", Record<keyof typeof DEFAULT_NETWORKS, Network>>
    | undefined
)
export type Config = SchemaConfig & Retries & {
    networks: Record<keyof typeof DEFAULT_NETWORKS, Network>
}
