import { Contract } from '@/src/lib/types'

export const contractsPath = 'src/contracts'

export const router: Contract = {
	name: 'Plug.sol',
	relativePath: '../base/'
} as const

export const factory: Contract = {
	name: 'Plug.Factory.sol',
	relativePath: '../base/'
} as const

export const baseFee: Contract = { 
    name: 'Plug.BaseFee.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const blockNumber: Contract = { 
    name: 'Plug.BlockNumber.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const clamp: Contract = {
    name: 'Plug.Clamp.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const limitedCalls: Contract = {
    name: 'Plug.LimitedCalls.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const nounsId: Contract = {
    name: 'Plug.NounsId.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const nounsTrait: Contract = {
    name: 'Plug.NounsTrait.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const revocation: Contract = {
    name: 'Plug.Revocation.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const timestamp: Contract = {
    name: 'Plug.Timestamp.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const window: Contract = {
    name: 'Plug.Window.Fuse.sol',
    relativePath: '../fuses/'
} as const

export const vaultContract: Contract = {
	name: 'Plug.Vault.Socket.sol',
	relativePath: '../sockets/'
} as const

export const etchContracts: Array<Contract> = [
	router,
	factory,
    baseFee,
    blockNumber,
    clamp,
    limitedCalls,
    nounsId,
    nounsTrait,
    revocation,
    timestamp,
    window,
] as const
