import { z } from "zod"

export const AddressSchema = z.string().transform((val, ctx) => {
	if (!/^0x[a-fA-F0-9]{40}$/.test(val))
		ctx.addIssue({
			code: z.ZodIssueCode.custom,
			message: `Invalid address ${val}`
		})

	return val as `0x${string}`
})
export const AddressesSchema = z.union([AddressSchema, z.array(AddressSchema)])

export type Address = z.infer<typeof AddressSchema>
export type Addresses = z.infer<typeof AddressesSchema>

export const SignatureSchema = z.string().transform((val, ctx) => {
	if (!/^0x[a-fA-F0-9]{130}$/.test(val))
		ctx.addIssue({
			code: z.ZodIssueCode.custom,
			message: `Invalid signature ${val}`
		})

	return val as `0x${string}`
})
export type Signature = z.infer<typeof SignatureSchema>

export const Bytes32Schema = z.string().transform((val, ctx) => {
	if (!/^0x[a-fA-F0-9]{64}$/.test(val))
		ctx.addIssue({
			code: z.ZodIssueCode.custom,
			message: `Invalid bytes32 ${val}`
		})

	return val as `0x${string}`
})
export type Bytes32 = z.infer<typeof Bytes32Schema>

export const BytesSchema = z.string().transform((val, ctx) => {
	if (!/^0x[a-fA-F0-9]*$/.test(val))
		ctx.addIssue({
			code: z.ZodIssueCode.custom,
			message: `Invalid bytes ${val}`
		})

	return val as `0x${string}`
})
export type Bytes = z.infer<typeof BytesSchema>

export const EVM_TYPES = [
	// * Basic types
	["bool", "z.boolean()"],
	["string", "z.string()"],

	// * Format dependent types
	["address", "AddressSchema"],
	["bytes", "BytesSchema"],
	["bytes32", "Bytes32Schema"],

	// * Regex dependent types such as uint and int
	// `(u)int<M>`: (un)signed integer type of `M` bits, `0 < M <= 256`, `M % 8 == 0`
	// https://regexr.com/6v8hp
	[
		/^u?int(8|16|24|32|40|48|56|64|72|80|88|96|104|112|120|128|136|144|152|160|168|176|184|192|200|208|216|224|232|240|248|256)?$/,

		// allow zod number and bigint
		"z.bigint()"
	]
] as const
