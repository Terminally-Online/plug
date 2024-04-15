import { Contract } from '@/src/lib/types'

export const contractsPath = 'src/contracts'

const base = (name: string): Contract => ({
    name,
    relativePath: '../base/'
})

const plug = (name: string): Contract => ({
    name,
    relativePath: '../plugs/'
})

const protocolPlug = (name: string): Contract => ({
    name,
    relativePath: '../plugs/protocols/'
})

const socket = (name: string): Contract => ({
    name,
    relativePath: '../sockets/'
})

export const router = base('Plug.sol')
export const factory = base('Plug.Factory.sol')
export const treasury = base('Plug.Treasury.sol')

export const balance = plug('Plug.Balance.sol')
export const balanceSemiFungible = plug('Plug.Balance.SemiFungible.sol')
export const baseFee = plug('Plug.BaseFee.sol')
export const blockNumber = plug('Plug.BlockNumber.sol')
export const calendar = plug('Plug.Calendar.sol')
export const limitedCalls = plug('Plug.LimitedCalls.sol')
export const revocation = plug('Plug.Revocation.sol')
export const timestamp = plug('Plug.Timestamp.sol')

export const fraxlend = protocolPlug('Plug.Fraxlend.APY.sol')
export const nounsBid = protocolPlug('Plug.Nouns.Bid.sol')
export const nounsId = protocolPlug('Plug.Nouns.Id.sol')
export const nounsTrait = protocolPlug('Plug.Nouns.Trait.sol')

export const vault = socket('Plug.Vault.Socket.sol')

export const constantContracts: Readonly<Array<Contract>> = [factory] as const
export const etchContracts: Readonly<Array<Contract>> = [
    // ! Bases
    router,
    factory,
    treasury,
    // ! Sockets
    vault,
    // ! Fuses
    balance,
    balanceSemiFungible,
    baseFee,
    blockNumber,
    calendar,
    limitedCalls,
    revocation,
    timestamp,
    // ! Protocols
    fraxlend,
    nounsBid,
    nounsId,
    nounsTrait,
] as const
