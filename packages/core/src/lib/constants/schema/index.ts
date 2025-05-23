export const EIP712_TYPES = {
	EIP712Domain: [
		{ name: 'name', type: 'string' },
		{ name: 'version', type: 'string' },
		{ name: 'chainId', type: 'uint256' },
		{ name: 'verifyingContract', type: 'address' }
	]
} as const

export const LIVE_PLUGS_TYPES = {
	Slice: [
		{ name: 'index', type: 'uint8' },
		{ name: 'start', type: 'uint256' },
		{ name: 'length', type: 'uint256' },
		{ name: 'typeId', type: 'uint8' }
	],
	Update: [
		{ name: 'start', type: 'uint256' },
		{ name: 'slice', type: 'Slice' }
	],
	Plug: [
		{ name: 'data', type: 'bytes' },
		{ name: 'updates', type: 'Update[]' }
	],
	Plugs: [
		{ name: 'socket', type: 'address' },
		{ name: 'plugs', type: 'Plug[]' },
		{ name: 'solver', type: 'bytes' },
		{ name: 'salt', type: 'bytes' }
	],
	LivePlugs: [
		{ name: 'plugs', type: 'Plugs' },
		{ name: 'signature', type: 'bytes' }
	]
} as const

export const DEFAULT_SCHEMA = {
	config: {
		contract: {
			name: 'PlugTypes',
			filename: 'Plug.Types',
			license: 'MIT',
			solidity: '0.8.26'
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
			schema: './src/contracts/abstracts/',
			documentation: './dist/docs/',
			zod: './src/lib/types/schema/'
		}
	},
	types: LIVE_PLUGS_TYPES
} as const
