import {
	createUseReadContract,
	createUseSimulateContract,
	createUseWatchContractEvent,
	createUseWriteContract
} from 'wagmi/codegen'

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Plug
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugAbi = [
	{
		type: 'function',
		inputs: [],
		name: 'name',
		outputs: [{ name: '$name', internalType: 'string', type: 'string' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{
				name: '$livePlugs',
				internalType: 'struct PlugTypesLib.LivePlugs',
				type: 'tuple',
				components: [
					{
						name: 'plugs',
						internalType: 'struct PlugTypesLib.Plugs',
						type: 'tuple',
						components: [
							{
								name: 'socket',
								internalType: 'address',
								type: 'address'
							},
							{
								name: 'plugs',
								internalType: 'struct PlugTypesLib.Plug[]',
								type: 'tuple[]',
								components: [
									{
										name: 'current',
										internalType:
											'struct PlugTypesLib.Current',
										type: 'tuple',
										components: [
											{
												name: 'target',
												internalType: 'address',
												type: 'address'
											},
											{
												name: 'value',
												internalType: 'uint256',
												type: 'uint256'
											},
											{
												name: 'data',
												internalType: 'bytes',
												type: 'bytes'
											}
										]
									},
									{
										name: 'fuses',
										internalType:
											'struct PlugTypesLib.Fuse[]',
										type: 'tuple[]',
										components: [
											{
												name: 'target',
												internalType: 'address',
												type: 'address'
											},
											{
												name: 'data',
												internalType: 'bytes',
												type: 'bytes'
											}
										]
									}
								]
							},
							{
								name: 'salt',
								internalType: 'bytes32',
								type: 'bytes32'
							},
							{
								name: 'fee',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'maxFeePerGas',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'maxPriorityFeePerGas',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'executor',
								internalType: 'address',
								type: 'address'
							}
						]
					},
					{ name: 'signature', internalType: 'bytes', type: 'bytes' }
				]
			}
		],
		name: 'plug',
		outputs: [
			{ name: '$results', internalType: 'bytes[]', type: 'bytes[]' }
		],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [
			{
				name: '$livePlugs',
				internalType: 'struct PlugTypesLib.LivePlugs[]',
				type: 'tuple[]',
				components: [
					{
						name: 'plugs',
						internalType: 'struct PlugTypesLib.Plugs',
						type: 'tuple',
						components: [
							{
								name: 'socket',
								internalType: 'address',
								type: 'address'
							},
							{
								name: 'plugs',
								internalType: 'struct PlugTypesLib.Plug[]',
								type: 'tuple[]',
								components: [
									{
										name: 'current',
										internalType:
											'struct PlugTypesLib.Current',
										type: 'tuple',
										components: [
											{
												name: 'target',
												internalType: 'address',
												type: 'address'
											},
											{
												name: 'value',
												internalType: 'uint256',
												type: 'uint256'
											},
											{
												name: 'data',
												internalType: 'bytes',
												type: 'bytes'
											}
										]
									},
									{
										name: 'fuses',
										internalType:
											'struct PlugTypesLib.Fuse[]',
										type: 'tuple[]',
										components: [
											{
												name: 'target',
												internalType: 'address',
												type: 'address'
											},
											{
												name: 'data',
												internalType: 'bytes',
												type: 'bytes'
											}
										]
									}
								]
							},
							{
								name: 'salt',
								internalType: 'bytes32',
								type: 'bytes32'
							},
							{
								name: 'fee',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'maxFeePerGas',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'maxPriorityFeePerGas',
								internalType: 'uint256',
								type: 'uint256'
							},
							{
								name: 'executor',
								internalType: 'address',
								type: 'address'
							}
						]
					},
					{ name: 'signature', internalType: 'bytes', type: 'bytes' }
				]
			}
		],
		name: 'plug',
		outputs: [
			{ name: '$results', internalType: 'bytes[][]', type: 'bytes[][]' }
		],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [],
		name: 'symbol',
		outputs: [{ name: '$version', internalType: 'string', type: 'string' }],
		stateMutability: 'pure'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugBaseFeeFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugBaseFeeFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		name: 'encode',
		outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugBlockNumberFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugBlockNumberFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		name: 'encode',
		outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugClampFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugClampFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$min', internalType: 'uint256', type: 'uint256' },
			{ name: '$max', internalType: 'uint256', type: 'uint256' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$min', internalType: 'uint256', type: 'uint256' },
			{ name: '$max', internalType: 'uint256', type: 'uint256' }
		],
		name: 'encode',
		outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugFactory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugFactoryAbi = [
	{
		type: 'function',
		inputs: [
			{
				name: '$implementation',
				internalType: 'address',
				type: 'address'
			},
			{ name: '$admin', internalType: 'address', type: 'address' },
			{ name: '$salt', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'deploy',
		outputs: [
			{ name: '$alreadyDeployed', internalType: 'bool', type: 'bool' },
			{ name: '$vault', internalType: 'address', type: 'address' }
		],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [
			{
				name: '$implementation',
				internalType: 'address',
				type: 'address'
			},
			{ name: '$salt', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'getAddress',
		outputs: [{ name: '$vault', internalType: 'address', type: 'address' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{
				name: '$implementation',
				internalType: 'address',
				type: 'address'
			}
		],
		name: 'initCodeHash',
		outputs: [
			{ name: '$initCodeHash', internalType: 'bytes32', type: 'bytes32' }
		],
		stateMutability: 'view'
	},
	{
		type: 'event',
		anonymous: false,
		inputs: [
			{
				name: 'implementation',
				internalType: 'address',
				type: 'address',
				indexed: true
			},
			{
				name: 'vault',
				internalType: 'address',
				type: 'address',
				indexed: true
			},
			{
				name: 'salt',
				internalType: 'bytes32',
				type: 'bytes32',
				indexed: false
			}
		],
		name: 'SocketDeployed'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugLimitedCallsFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugLimitedCallsFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$terms', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$callCount', internalType: 'uint256', type: 'uint256' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$callCount', internalType: 'uint256', type: 'uint256' }
		],
		name: 'encode',
		outputs: [{ name: '$terms', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'nonpayable'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugNounsIdFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugNounsIdFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$live', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [{ name: '$value', internalType: 'uint256', type: 'uint256' }],
		name: 'encode',
		outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugNounsTraitFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugNounsTraitFuseAbi = [
	{
		type: 'constructor',
		inputs: [{ name: '$art', internalType: 'address', type: 'address' }],
		stateMutability: 'nonpayable'
	},
	{
		type: 'function',
		inputs: [],
		name: 'ACCESSORY_SELECTOR',
		outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'BACKGROUND_SELECTOR',
		outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'BODY_SELECTOR',
		outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'GLASSES_SELECTOR',
		outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'HEAD_SELECTOR',
		outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'cancelOwnershipHandover',
		outputs: [],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [
			{ name: 'pendingOwner', internalType: 'address', type: 'address' }
		],
		name: 'completeOwnershipHandover',
		outputs: [],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [{ name: '$live', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$selector', internalType: 'bytes32', type: 'bytes32' },
			{ name: '$trait', internalType: 'bytes32', type: 'bytes32' }
		],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$selector', internalType: 'bytes32', type: 'bytes32' },
			{ name: '$trait', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'encode',
		outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$selector', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'nounTrait',
		outputs: [
			{ name: '$traitHash', internalType: 'bytes32', type: 'bytes32' }
		],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'owner',
		outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: 'pendingOwner', internalType: 'address', type: 'address' }
		],
		name: 'ownershipHandoverExpiresAt',
		outputs: [{ name: 'result', internalType: 'uint256', type: 'uint256' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [],
		name: 'renounceOwnership',
		outputs: [],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [],
		name: 'requestOwnershipHandover',
		outputs: [],
		stateMutability: 'payable'
	},
	{
		type: 'function',
		inputs: [{ name: '$art', internalType: 'address', type: 'address' }],
		name: 'setArt',
		outputs: [],
		stateMutability: 'nonpayable'
	},
	{
		type: 'function',
		inputs: [
			{ name: 'newOwner', internalType: 'address', type: 'address' }
		],
		name: 'transferOwnership',
		outputs: [],
		stateMutability: 'payable'
	},
	{
		type: 'event',
		anonymous: false,
		inputs: [
			{
				name: 'pendingOwner',
				internalType: 'address',
				type: 'address',
				indexed: true
			}
		],
		name: 'OwnershipHandoverCanceled'
	},
	{
		type: 'event',
		anonymous: false,
		inputs: [
			{
				name: 'pendingOwner',
				internalType: 'address',
				type: 'address',
				indexed: true
			}
		],
		name: 'OwnershipHandoverRequested'
	},
	{
		type: 'event',
		anonymous: false,
		inputs: [
			{
				name: 'oldOwner',
				internalType: 'address',
				type: 'address',
				indexed: true
			},
			{
				name: 'newOwner',
				internalType: 'address',
				type: 'address',
				indexed: true
			}
		],
		name: 'OwnershipTransferred'
	},
	{ type: 'error', inputs: [], name: 'AlreadyInitialized' },
	{ type: 'error', inputs: [], name: 'NewOwnerIsZeroAddress' },
	{ type: 'error', inputs: [], name: 'NoHandoverRequest' },
	{ type: 'error', inputs: [], name: 'Unauthorized' }
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugRevocationFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugRevocationFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$sender', internalType: 'address', type: 'address' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [{ name: '$sender', internalType: 'address', type: 'address' }],
		name: 'encode',
		outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '', internalType: 'address', type: 'address' },
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'isRevoked',
		outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' },
			{ name: '$revoked', internalType: 'bool', type: 'bool' }
		],
		name: 'revoke',
		outputs: [],
		stateMutability: 'nonpayable'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugTimestampFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugTimestampFuseAbi = [
	{
		type: 'function',
		inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		name: 'decode',
		outputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$operator', internalType: 'uint128', type: 'uint128' },
			{ name: '$threshold', internalType: 'uint128', type: 'uint128' }
		],
		name: 'encode',
		outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	}
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugWindowFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugWindowFuseAbi = [
	{
		type: 'function',
		inputs: [
			{ name: '$schedule', internalType: 'uint256', type: 'uint256' }
		],
		name: 'decode',
		outputs: [
			{ name: '$startTime', internalType: 'uint32', type: 'uint32' },
			{ name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
			{ name: '$duration', internalType: 'uint32', type: 'uint32' },
			{ name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$startTime', internalType: 'uint32', type: 'uint32' },
			{ name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
			{ name: '$duration', internalType: 'uint32', type: 'uint32' },
			{ name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' }
		],
		name: 'encode',
		outputs: [
			{ name: '$schedule', internalType: 'uint256', type: 'uint256' }
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$live', internalType: 'bytes', type: 'bytes' },
			{
				name: '$current',
				internalType: 'struct PlugTypesLib.Current',
				type: 'tuple',
				components: [
					{
						name: 'target',
						internalType: 'address',
						type: 'address'
					},
					{ name: 'value', internalType: 'uint256', type: 'uint256' },
					{ name: 'data', internalType: 'bytes', type: 'bytes' }
				]
			},
			{ name: '', internalType: 'bytes32', type: 'bytes32' }
		],
		name: 'enforceFuse',
		outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$startTime', internalType: 'uint32', type: 'uint32' },
			{ name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
			{ name: '$duration', internalType: 'uint32', type: 'uint32' },
			{ name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' }
		],
		name: 'isWithinWindow',
		outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$schedule', internalType: 'uint256', type: 'uint256' }
		],
		name: 'isWithinWindow',
		outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
		stateMutability: 'view'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$schedule', internalType: 'uint256', type: 'uint256' }
		],
		name: 'toWindow',
		outputs: [
			{
				name: '$window',
				internalType: 'struct WindowFuseLib.Window',
				type: 'tuple',
				components: [
					{
						name: 'periods',
						internalType: 'struct WindowFuseLib.Period[]',
						type: 'tuple[]',
						components: [
							{
								name: 'startTime',
								internalType: 'uint32',
								type: 'uint32'
							},
							{
								name: 'endTime',
								internalType: 'uint32',
								type: 'uint32'
							}
						]
					}
				]
			}
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$startTime', internalType: 'uint32', type: 'uint32' },
			{ name: '$duration', internalType: 'uint32', type: 'uint32' },
			{ name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' }
		],
		name: 'toWindow',
		outputs: [
			{
				name: '$window',
				internalType: 'struct WindowFuseLib.Window',
				type: 'tuple',
				components: [
					{
						name: 'periods',
						internalType: 'struct WindowFuseLib.Period[]',
						type: 'tuple[]',
						components: [
							{
								name: 'startTime',
								internalType: 'uint32',
								type: 'uint32'
							},
							{
								name: 'endTime',
								internalType: 'uint32',
								type: 'uint32'
							}
						]
					}
				]
			}
		],
		stateMutability: 'pure'
	},
	{
		type: 'function',
		inputs: [
			{ name: '$schedule', internalType: 'uint256', type: 'uint256' },
			{ name: '$n', internalType: 'uint32', type: 'uint32' }
		],
		name: 'toWindows',
		outputs: [
			{
				name: '$windows',
				internalType: 'struct WindowFuseLib.Window[]',
				type: 'tuple[]',
				components: [
					{
						name: 'periods',
						internalType: 'struct WindowFuseLib.Period[]',
						type: 'tuple[]',
						components: [
							{
								name: 'startTime',
								internalType: 'uint32',
								type: 'uint32'
							},
							{
								name: 'endTime',
								internalType: 'uint32',
								type: 'uint32'
							}
						]
					}
				]
			},
			{ name: '$cursor', internalType: 'uint32', type: 'uint32' }
		],
		stateMutability: 'view'
	},
	{ type: 'error', inputs: [], name: 'WindowCaveatViolation' },
	{ type: 'error', inputs: [], name: 'WindowLackingDays' },
	{ type: 'error', inputs: [], name: 'WindowLackingDuration' },
	{ type: 'error', inputs: [], name: 'WindowLackingStartTime' },
	{ type: 'error', inputs: [], name: 'WindowLackingSufficientRepeatsEvery' }
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// React
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useReadPlug = /*#__PURE__*/ createUseReadContract({ abi: plugAbi })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"name"`
 */
export const useReadPlugName = /*#__PURE__*/ createUseReadContract({
	abi: plugAbi,
	functionName: 'name'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"symbol"`
 */
export const useReadPlugSymbol = /*#__PURE__*/ createUseReadContract({
	abi: plugAbi,
	functionName: 'symbol'
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useWritePlug = /*#__PURE__*/ createUseWriteContract({
	abi: plugAbi
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"plug"`
 */
export const useWritePlugPlug = /*#__PURE__*/ createUseWriteContract({
	abi: plugAbi,
	functionName: 'plug'
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useSimulatePlug = /*#__PURE__*/ createUseSimulateContract({
	abi: plugAbi
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"plug"`
 */
export const useSimulatePlugPlug = /*#__PURE__*/ createUseSimulateContract({
	abi: plugAbi,
	functionName: 'plug'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__
 */
export const useReadPlugBaseFeeFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugBaseFeeFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugBaseFeeFuseDecode = /*#__PURE__*/ createUseReadContract(
	{ abi: plugBaseFeeFuseAbi, functionName: 'decode' }
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugBaseFeeFuseEncode = /*#__PURE__*/ createUseReadContract(
	{ abi: plugBaseFeeFuseAbi, functionName: 'encode' }
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugBaseFeeFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugBaseFeeFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__
 */
export const useReadPlugBlockNumberFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugBlockNumberFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugBlockNumberFuseDecode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugBlockNumberFuseAbi,
		functionName: 'decode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugBlockNumberFuseEncode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugBlockNumberFuseAbi,
		functionName: 'encode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugBlockNumberFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugBlockNumberFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__
 */
export const useReadPlugClampFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugClampFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugClampFuseDecode = /*#__PURE__*/ createUseReadContract({
	abi: plugClampFuseAbi,
	functionName: 'decode'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugClampFuseEncode = /*#__PURE__*/ createUseReadContract({
	abi: plugClampFuseAbi,
	functionName: 'encode'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugClampFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugClampFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useReadPlugFactory = /*#__PURE__*/ createUseReadContract({
	abi: plugFactoryAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"getAddress"`
 */
export const useReadPlugFactoryGetAddress = /*#__PURE__*/ createUseReadContract(
	{ abi: plugFactoryAbi, functionName: 'getAddress' }
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"initCodeHash"`
 */
export const useReadPlugFactoryInitCodeHash =
	/*#__PURE__*/ createUseReadContract({
		abi: plugFactoryAbi,
		functionName: 'initCodeHash'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWritePlugFactory = /*#__PURE__*/ createUseWriteContract({
	abi: plugFactoryAbi
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"deploy"`
 */
export const useWritePlugFactoryDeploy = /*#__PURE__*/ createUseWriteContract({
	abi: plugFactoryAbi,
	functionName: 'deploy'
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useSimulatePlugFactory = /*#__PURE__*/ createUseSimulateContract({
	abi: plugFactoryAbi
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"deploy"`
 */
export const useSimulatePlugFactoryDeploy =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugFactoryAbi,
		functionName: 'deploy'
	})

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWatchPlugFactoryEvent =
	/*#__PURE__*/ createUseWatchContractEvent({ abi: plugFactoryAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"SocketDeployed"`
 */
export const useWatchPlugFactorySocketDeployedEvent =
	/*#__PURE__*/ createUseWatchContractEvent({
		abi: plugFactoryAbi,
		eventName: 'SocketDeployed'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useReadPlugLimitedCallsFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugLimitedCallsFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugLimitedCallsFuseDecode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugLimitedCallsFuseAbi,
		functionName: 'decode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugLimitedCallsFuseEncode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugLimitedCallsFuseAbi,
		functionName: 'encode'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useWritePlugLimitedCallsFuse =
	/*#__PURE__*/ createUseWriteContract({ abi: plugLimitedCallsFuseAbi })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useWritePlugLimitedCallsFuseEnforceFuse =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugLimitedCallsFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useSimulatePlugLimitedCallsFuse =
	/*#__PURE__*/ createUseSimulateContract({ abi: plugLimitedCallsFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useSimulatePlugLimitedCallsFuseEnforceFuse =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugLimitedCallsFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__
 */
export const useReadPlugNounsIdFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugNounsIdFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugNounsIdFuseDecode = /*#__PURE__*/ createUseReadContract(
	{ abi: plugNounsIdFuseAbi, functionName: 'decode' }
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugNounsIdFuseEncode = /*#__PURE__*/ createUseReadContract(
	{ abi: plugNounsIdFuseAbi, functionName: 'encode' }
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugNounsIdFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsIdFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useReadPlugNounsTraitFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugNounsTraitFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"ACCESSORY_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseAccessorySelector =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'ACCESSORY_SELECTOR'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"BACKGROUND_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseBackgroundSelector =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'BACKGROUND_SELECTOR'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"BODY_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseBodySelector =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'BODY_SELECTOR'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"GLASSES_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseGlassesSelector =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'GLASSES_SELECTOR'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"HEAD_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseHeadSelector =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'HEAD_SELECTOR'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugNounsTraitFuseDecode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'decode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugNounsTraitFuseEncode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'encode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugNounsTraitFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"nounTrait"`
 */
export const useReadPlugNounsTraitFuseNounTrait =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'nounTrait'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugNounsTraitFuseOwner =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'owner'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"ownershipHandoverExpiresAt"`
 */
export const useReadPlugNounsTraitFuseOwnershipHandoverExpiresAt =
	/*#__PURE__*/ createUseReadContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'ownershipHandoverExpiresAt'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useWritePlugNounsTraitFuse = /*#__PURE__*/ createUseWriteContract({
	abi: plugNounsTraitFuseAbi
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseCancelOwnershipHandover =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'cancelOwnershipHandover'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseCompleteOwnershipHandover =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'completeOwnershipHandover'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePlugNounsTraitFuseRenounceOwnership =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'renounceOwnership'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseRequestOwnershipHandover =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'requestOwnershipHandover'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"setArt"`
 */
export const useWritePlugNounsTraitFuseSetArt =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'setArt'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugNounsTraitFuseTransferOwnership =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'transferOwnership'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useSimulatePlugNounsTraitFuse =
	/*#__PURE__*/ createUseSimulateContract({ abi: plugNounsTraitFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseCancelOwnershipHandover =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'cancelOwnershipHandover'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseCompleteOwnershipHandover =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'completeOwnershipHandover'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePlugNounsTraitFuseRenounceOwnership =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'renounceOwnership'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseRequestOwnershipHandover =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'requestOwnershipHandover'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"setArt"`
 */
export const useSimulatePlugNounsTraitFuseSetArt =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'setArt'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugNounsTraitFuseTransferOwnership =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugNounsTraitFuseAbi,
		functionName: 'transferOwnership'
	})

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useWatchPlugNounsTraitFuseEvent =
	/*#__PURE__*/ createUseWatchContractEvent({ abi: plugNounsTraitFuseAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipHandoverCanceled"`
 */
export const useWatchPlugNounsTraitFuseOwnershipHandoverCanceledEvent =
	/*#__PURE__*/ createUseWatchContractEvent({
		abi: plugNounsTraitFuseAbi,
		eventName: 'OwnershipHandoverCanceled'
	})

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipHandoverRequested"`
 */
export const useWatchPlugNounsTraitFuseOwnershipHandoverRequestedEvent =
	/*#__PURE__*/ createUseWatchContractEvent({
		abi: plugNounsTraitFuseAbi,
		eventName: 'OwnershipHandoverRequested'
	})

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPlugNounsTraitFuseOwnershipTransferredEvent =
	/*#__PURE__*/ createUseWatchContractEvent({
		abi: plugNounsTraitFuseAbi,
		eventName: 'OwnershipTransferred'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useReadPlugRevocationFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugRevocationFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugRevocationFuseDecode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugRevocationFuseAbi,
		functionName: 'decode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugRevocationFuseEncode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugRevocationFuseAbi,
		functionName: 'encode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugRevocationFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugRevocationFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"isRevoked"`
 */
export const useReadPlugRevocationFuseIsRevoked =
	/*#__PURE__*/ createUseReadContract({
		abi: plugRevocationFuseAbi,
		functionName: 'isRevoked'
	})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useWritePlugRevocationFuse = /*#__PURE__*/ createUseWriteContract({
	abi: plugRevocationFuseAbi
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"revoke"`
 */
export const useWritePlugRevocationFuseRevoke =
	/*#__PURE__*/ createUseWriteContract({
		abi: plugRevocationFuseAbi,
		functionName: 'revoke'
	})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useSimulatePlugRevocationFuse =
	/*#__PURE__*/ createUseSimulateContract({ abi: plugRevocationFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"revoke"`
 */
export const useSimulatePlugRevocationFuseRevoke =
	/*#__PURE__*/ createUseSimulateContract({
		abi: plugRevocationFuseAbi,
		functionName: 'revoke'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__
 */
export const useReadPlugTimestampFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugTimestampFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugTimestampFuseDecode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugTimestampFuseAbi,
		functionName: 'decode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugTimestampFuseEncode =
	/*#__PURE__*/ createUseReadContract({
		abi: plugTimestampFuseAbi,
		functionName: 'encode'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugTimestampFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugTimestampFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__
 */
export const useReadPlugWindowFuse = /*#__PURE__*/ createUseReadContract({
	abi: plugWindowFuseAbi
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugWindowFuseDecode = /*#__PURE__*/ createUseReadContract({
	abi: plugWindowFuseAbi,
	functionName: 'decode'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugWindowFuseEncode = /*#__PURE__*/ createUseReadContract({
	abi: plugWindowFuseAbi,
	functionName: 'encode'
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugWindowFuseEnforceFuse =
	/*#__PURE__*/ createUseReadContract({
		abi: plugWindowFuseAbi,
		functionName: 'enforceFuse'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"isWithinWindow"`
 */
export const useReadPlugWindowFuseIsWithinWindow =
	/*#__PURE__*/ createUseReadContract({
		abi: plugWindowFuseAbi,
		functionName: 'isWithinWindow'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"toWindow"`
 */
export const useReadPlugWindowFuseToWindow =
	/*#__PURE__*/ createUseReadContract({
		abi: plugWindowFuseAbi,
		functionName: 'toWindow'
	})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"toWindows"`
 */
export const useReadPlugWindowFuseToWindows =
	/*#__PURE__*/ createUseReadContract({
		abi: plugWindowFuseAbi,
		functionName: 'toWindows'
	})
