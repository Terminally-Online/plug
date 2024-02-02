import { Contract, MinedContract } from "./types";

export const contractsPath = 'src/contracts'

export const routerContract: MinedContract = {
    name: 'Plug.Router.Socket.sol',
    relativePath: '../sockets/',
    salt: '0x0',
    address: 'address(0)'
} as const

export const vaultContract: MinedContract= {
    name: 'Plug.Vault.Socket.sol',
    relativePath: '../sockets/',
    salt: '0x0',
    address: 'address(0)'
} as const

export const receiverContract: Contract = { 
    name: 'Plug.Receiver.sol',
    relativePath: '../sockets'
}

export const contracts: Array<MinedContract> = [
    routerContract, 
    vaultContract
] as const
