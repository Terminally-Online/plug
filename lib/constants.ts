import { Contract } from "./types";

export const contracts: Array<Contract> = [
    {
        name: 'Plug.Router.Socket.sol',
        relativePath: '../sockets/',
        salt: '0x0',
        address: 'address(0)'
    }, {
        name: 'Plug.Vault.Socket.sol',
        relativePath: '../sockets/',
        salt: '0x0',
        address: 'address(0)'
    }
] as const
