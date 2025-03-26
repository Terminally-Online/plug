import { z } from 'zod'
import { AddressSchema,BytesSchema } from '@/src/lib'

            export const EIP712DomainSchema= z.object({
	name: z.string(),
	version: z.string(),
	chainId: z.number(),
	verifyingContract: AddressSchema
})
            export type EIP712Domain = z.infer<typeof EIP712DomainSchema>
        


            export const SliceSchema= z.object({
	index: z.bigint(),
	start: z.bigint(),
	length: z.bigint(),
	typeId: z.bigint()
})
            export type Slice = z.infer<typeof SliceSchema>
        


            export const UpdateSchema= z.object({
	start: z.bigint(),
	slice: SliceSchema
})
            export type Update = z.infer<typeof UpdateSchema>
        


            export const PlugSchema= z.object({
	selector: z.bigint(),
	to: AddressSchema,
	data: BytesSchema,
	value: z.bigint(),
	updates: z.array(UpdateSchema)
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
        