import { getEVMSchema, SchemaConfig } from '@/src/lib'
import { EIP712_TYPES } from '@/src/lib/constants/schema'

export const generateZodSchema = async (config: SchemaConfig) => {
	const types = {
		source: { ...EIP712_TYPES, ...config.types },
		schemas: [] as Array<string>,
		used: new Set<string>(),
		generated: new Set<string>()
	}

	for (const type of Object.keys(types.source)) {
		const fields = types.source[type as keyof (typeof types)['source']]

		if (!fields) continue

		const schema = fields
			.map(field => {
				const typeUsed =
					field.name === 'chainId'
						? 'z.number()'
						: getEVMSchema(field.type)

				if (
					typeUsed.startsWith('z.') === false &&
					types.generated.has(typeUsed) === false
				)
					types.used.add(typeUsed)

				return `\n\t${field.name}: ${typeUsed}`
			})
			.join(',')

		const typeName = `${type}Schema`

		types.generated.add(typeName)
		types.schemas.push(`
            export const ${typeName}= z.object({${schema}\n})
            export type ${type} = z.infer<typeof ${typeName}>
        `)
	}

	const imports = ["import { z } from 'zod'"]

	if (types.used.size > 0)
		imports.push(
			`import { ${Array.from(types.used.values()).join(',')} } from '@/src/lib'`
		)

	const schemas = [imports.join('\n'), types.schemas.join('\n\n')].join('\n')

	return { schemas }
}
