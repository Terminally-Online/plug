import { Contract } from '@/src/lib/types'

export const contractsPath = 'src/contracts'

const base = (name: string): Contract => ({
    name,
    relativePath: '../base/'
})

const fuse = (name: string): Contract => ({
    name,
    relativePath: '../fuses/'
})

const protocolFuse = (name: string): Contract => ({
    name,
    relativePath: '../fuses/protocols/'
})

const socket = (name: string): Contract => ({
    name,
    relativePath: '../sockets/'
})

export const router = base('Plug.sol')
export const factory = base('Plug.Factory.sol')
export const treasury = base('Plug.Treasury.sol')

export const balance = fuse('Plug.Balance.Fuse.sol')
export const balanceSemiFungible = fuse('Plug.Balance.SemiFungible.Fuse.sol')
export const baseFee = fuse('Plug.BaseFee.Fuse.sol')
export const blockNumber = fuse('Plug.BlockNumber.Fuse.sol')
export const limitedCalls = fuse('Plug.LimitedCalls.Fuse.sol')
export const revocation = fuse('Plug.Revocation.Fuse.sol')
export const timestamp = fuse('Plug.Timestamp.Fuse.sol')
export const window = fuse('Plug.Window.Fuse.sol')

export const fraxlend = protocolFuse('Plug.Fraxlend.APY.Fuse.sol')
export const nounsBid = protocolFuse('Plug.Nouns.Bid.Fuse.sol')
export const nounsId = protocolFuse('Plug.Nouns.Id.Fuse.sol')
export const nounsTrait = protocolFuse('Plug.Nouns.Trait.Fuse.sol')

export const vault = socket('Plug.Vault.Socket.sol')

export const constantContracts = [factory, treasury]
export const etchContracts: Array<Contract> = [
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
    limitedCalls,
    revocation,
    timestamp,
    window,
    // ! Protocols
    nounsBid,
    nounsId,
    nounsTrait,
] as const
