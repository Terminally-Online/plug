export const EIP712_TYPES = {
	EIP712Domain: [
		{ name: "name", type: "string" },
		{ name: "version", type: "string" },
		{ name: "chainId", type: "uint256" },
		{ name: "verifyingContract", type: "address" }
	]
} as const

export const LIVE_PLUGS_TYPES = {
	Plug: [
		{ name: "target", type: "address" },
		{ name: "value", type: "uint256" },
		{ name: "data", type: "bytes" }
	],
	Plugs: [
		{ name: "socket", type: "address" },
		{ name: "plugs", type: "Plug[]" },
		{ name: "solver", type: "bytes" },
		{ name: "salt", type: "bytes32" }
	],
	LivePlugs: [
		{ name: "plugs", type: "Plugs" },
		{ name: "signature", type: "bytes" }
	]
} as const

export const DEFAULT_SCHEMA = {
	config: {
		contract: {
			name: "PlugTypes",
			filename: "Plug.Types",
			license: "MIT",
			solidity: "0.8.23"
		},
		types: EIP712_TYPES,
		dangerous: {
			excludeCoreTypes: false,
			useOverloads: false,
			useDocs: false,
			packetHashName: (typeName: string) =>
				typeName.slice(0, 1).toUpperCase() + typeName.slice(1)
		},
		out: {
			schema: "./src/contracts/abstracts/",
			documentation: "./dist/docs/",
			zod: "./src/lib/types/schema/"
		}
	},
	types: LIVE_PLUGS_TYPES
} as const
