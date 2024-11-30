import { z } from "zod"

export const literalUnion = <T extends string | number>(
	constants: readonly T[]
) => {
	const literals = constants.map(x => z.literal(x)) as unknown as readonly [
		z.ZodLiteral<T>,
		z.ZodLiteral<T>,
		...z.ZodLiteral<T>[]
	]
	return z.union(literals)
}
