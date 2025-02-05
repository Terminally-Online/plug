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

export const constantContracts: Readonly<Array<Contract>> = [factory, ticket] as const
export const etchContracts: Readonly<Array<Contract>> = [
	// ! Bases
	router,
	factory,
	socket
] as const
