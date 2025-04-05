import { Contract } from '@/src/lib/types'

export * from './bundle'
export * from './schema'

export const contractsPath = 'src/contracts'

const base = (name: string): Contract => ({
	name,
	relativePath: '../base/'
})

export const router = base('Plug.sol')
export const factory = base('Plug.Factory.sol')
export const socket = base('Plug.Socket.sol')
export const ticket = base('Plug.Ticket.sol')
export const token = base('Plug.Token.sol')
export const rewards = base('Plug.Rewards.sol')

export const assert = base('Plug.Assert.sol')
export const boolean = base('Plug.Boolean.sol')
export const coercion = base('Plug.Coercion.sol')
export const database = base('Plug.Database.sol')
export const evm = base('Plug.EVM.sol')
export const math = base('Plug.Math.sol')

export const constantContracts: Readonly<Array<Contract>> = [
	factory,
	ticket,
	assert,
	boolean,
	coercion,
	database,
	evm,
	math,
	token,
	rewards
] as const
export const etchContracts: Readonly<Array<Contract>> = [
	router,
	factory,
	socket
] as const
