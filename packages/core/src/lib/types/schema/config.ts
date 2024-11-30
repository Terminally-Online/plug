import { TypedData } from "viem"

export type SchemaContract = {
	authors: Array<string> | string
	name: string
	filename: string
	license: string
	solidity: string
}

export type SchemaDangerous = {
	excludeCoreTypes: boolean
	useOverloads: boolean
	useDocs: boolean
	packetHashName: (typeName: string) => string
}

export type SchemaOutput = {
	schema: string
	documentation: string
	zod: string
}

export type BaseSchemaConfig = Partial<{
	contract: Partial<SchemaContract>
	out: Partial<SchemaOutput>
	dangerous: Partial<SchemaDangerous>
	types?: TypedData | undefined
}>

export type SchemaConfig = {
	contract: SchemaContract
	out: SchemaOutput
	dangerous: SchemaDangerous
	types: TypedData
}
