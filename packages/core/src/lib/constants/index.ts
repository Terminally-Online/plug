import { Contract } from '@/src/lib/types'

export * from './bundle'
export * from './schema'

export const contractsPath = 'src/contracts'

const base = (name: string): Contract => ({
	name,
	relativePath: '../base/'
})

const actions = (name: string): Contract => ({
	name,
	relativePath: '../actions/'
})

export const router = base('Plug.sol')
export const factory = base('Plug.Factory.sol')
export const socket = base('Plug.Socket.sol')
export const ticket = base('Plug.Ticket.sol')
export const token = base('Plug.Token.sol')
export const rewards = base('Plug.Rewards.sol')

export const assert = actions('Plug.Assert.sol')
export const boolean = actions('Plug.Boolean.sol')
export const coercion = actions('Plug.Coercion.sol')
export const database = actions('Plug.Database.sol')
export const evm = actions('Plug.EVM.sol')
export const math = actions('Plug.Math.sol')

export const constantContracts: Readonly<Array<Contract>> = [
	router,
	factory,
	socket,
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
	socket,
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
