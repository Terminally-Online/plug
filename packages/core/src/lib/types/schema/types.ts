import { z } from 'zod'
import { AddressSchema,BytesSchema } from '@/src/lib'

            export const EIP712DomainSchema= z.object({
	name: z.string(),
	version: z.string(),
	chainId: z.number(),
	verifyingContract: AddressSchema
})
            export type EIP712Domain = z.infer<typeof EIP712DomainSchema>
        


            export const PlugSchema= z.object({
	to: AddressSchema,
	data: BytesSchema,
	value: z.bigint(),
	gas: z.bigint()
})
            export type Plug = z.infer<typeof PlugSchema>
        


            export const PlugsSchema= z.object({
	socket: AddressSchema,
	plugs: z.array(PlugSchema),
	solver: BytesSchema,
	salt: BytesSchema
})
            export type Plugs = z.infer<typeof PlugsSchema>
        


            export const LivePlugsSchema= z.object({
	plugs: PlugsSchema,
	signature: BytesSchema
})
            export type LivePlugs = z.infer<typeof LivePlugsSchema>
        