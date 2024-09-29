import { z } from 'zod'
import { AddressSchema, BytesSchema, Bytes32Schema } from '@/src/lib'

export const EIP712DomainSchema = z.object({
    name: z.string(),
    version: z.string(),
    chainId: z.number(),
    verifyingContract: AddressSchema
})
export type EIP712Domain = z.infer<typeof EIP712DomainSchema>



export const PlugSchema = z.object({
    target: AddressSchema,
    value: z.bigint(),
    data: BytesSchema
})
export type Plug = z.infer<typeof PlugSchema>



export const PlugsSchema = z.object({
    socket: AddressSchema,
    plugs: z.array(PlugSchema),
    solver: BytesSchema,
    salt: Bytes32Schema
})
export type Plugs = z.infer<typeof PlugsSchema>



export const LivePlugsSchema = z.object({
    plugs: PlugsSchema,
    signature: BytesSchema
})
export type LivePlugs = z.infer<typeof LivePlugsSchema>
