import { Contract } from "@/lib/types";

export const contractsPath = 'src/contracts'

export const routerContract: Contract = {
    name: 'Plug.Router.Socket.sol',
    relativePath: '../sockets/',
} as const

export const vaultContract: Contract = {
    name: 'Plug.Vault.Socket.sol',
    relativePath: '../sockets/',
} as const

export const receiverContract: Contract = { 
    name: 'Plug.Receiver.sol',
    relativePath: '../sockets'
}

export const etchContracts: Array<Contract> = [
    routerContract
] as const

export const mineContracts: Array<string> = [
    'Plug.Factory.sol',
    'Plug.Router.Socket.sol',
    'Plug.AllowedMethods.Fuse.sol',
    'Plug.BlockNumber.Fuse.sol',
    'Plug.Clamp.Fuse.sol',
    'Plug.LimitedCalls.Fuse.sol',
    'Plug.NounsId.Fuse.sol',
    'Plug.NounsTrait.Fuse.sol',
    'Plug.Revocation.Fuse.sol',
    'Plug.Timestamp.Fuse.sol',
    'Plug.Window.Fuse.sol',
    'Plug.NounsBid.Current.sol'
] as const
