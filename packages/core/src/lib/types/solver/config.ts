import { DEFAULT_NETWORKS, Network, Retries } from "@/src/lib"

export type BaseSolverConfig = Partial<Retries> &
	(
		| Record<"networks", Record<keyof typeof DEFAULT_NETWORKS, Network>>
		| undefined
	)

export type SolverConfig = Retries & {
	networks: Record<keyof typeof DEFAULT_NETWORKS, Network>
}
