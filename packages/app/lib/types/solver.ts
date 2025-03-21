import { z } from "zod"

const SchemasResponseIconSchema = z
	.object({
		default: z.string(),
		secondary: z.string()
	})
	.and(z.record(z.string(), z.string()))
export type SchemasResponseIcon = z.infer<typeof SchemasResponseIconSchema>

const SchemasResponseInfoSchema = z.object({
	label: z.string(),
	value: z.string()
})
export type SchemasResponseInfo = z.infer<typeof SchemasResponseInfoSchema>

const SchemasRequestValueManualSchema = z.object({
	key: z.string(),
	value: z.string(),
	name: z.string()
})
export type SchemasRequestValueManual = z.infer<typeof SchemasRequestValueManualSchema>

const SchemasRequestValueOptionExtensionSchema = z
	.object({
		label: z.string().optional(),
		icon: SchemasResponseIconSchema,
		info: SchemasResponseInfoSchema
	})
	.optional()
	.and(SchemasRequestValueManualSchema)
export type SchemasRequestValueOptionExtension = z.infer<typeof SchemasRequestValueOptionExtensionSchema>

const SchemasRequestValueSchema = SchemasRequestValueOptionExtensionSchema
export type SchemasRequestValue = z.infer<typeof SchemasRequestValueSchema>

const SchemasRequestValuesSchema = z.array(SchemasRequestValueSchema)
export type SchemasRequestValues = z.infer<typeof SchemasRequestValuesSchema>

const SchemasRequestValuesSchemaSet = z.record(z.number(), SchemasRequestValuesSchema)
export type SchemasRequestValuesSet = z.infer<typeof SchemasRequestValuesSchemaSet>

const SchemasRequestActionSchema = z.object({
	id: z.number(),
	protocol: z.string(),
	action: z.string(),
	values: SchemasRequestValuesSchema.optional()
})
export type SchemasRequestAction = z.infer<typeof SchemasRequestActionSchema>

const SchemasRequestActionsSchema = z.array(SchemasRequestActionSchema)
export type SchemasRequestActions = z.infer<typeof SchemasRequestActionsSchema>

const SchemasResponseChainSchema = z.object({
	name: z.string(),
	chainIds: z.array(z.number()),
	explorer: z.string(),
	icon: SchemasResponseIconSchema
})
export type SchemasResponseChain = z.infer<typeof SchemasResponseChainSchema>

const SchemasResponseChainsSchema = z.array(SchemasResponseChainSchema)
export type SchemasResponseChains = z.infer<typeof SchemasResponseChainsSchema>

const SchemasResponseOptionSchema = SchemasRequestValueSchema
export type SchemasResponseOption = z.infer<typeof SchemasResponseOptionSchema>

const SchemasResponseOptionsSchema = z.array(SchemasResponseOptionSchema)
export type SchemasResponseOptions = z.infer<typeof SchemasResponseOptionsSchema>

const SchemasResponseOptionsSetSchema: z.ZodType<
	Record<
		string,
		z.infer<typeof SchemasResponseOptionsSchema> | Record<string, z.infer<typeof SchemasResponseOptionsSchema>>
	>
> = z.record(z.string(), z.union([SchemasResponseOptionsSchema, z.record(z.string(), SchemasResponseOptionsSchema)]))
export type SchemasResponseOptionsSet = z.infer<typeof SchemasResponseOptionsSetSchema>

const SchemasResponseCoilsSchema = z.record(z.string(), z.string())
export type SchemasResponseCoils = z.infer<typeof SchemasResponseCoilsSchema>

const SchemasResponseSchemaSchema = z.object({
	metadata: z.object({
		icon: z.string(),
		tags: z.array(z.string()),
		chains: SchemasResponseChainsSchema
	}),
	schema: z.record(
		z.string(),
		z.object({
			type: z.string(),
			sentence: z.string(),
			options: SchemasResponseOptionsSetSchema.optional(),
			coils: SchemasResponseCoilsSchema.optional()
		})
	)
})
export type SchemasResponseSchema = z.infer<typeof SchemasResponseSchemaSchema>

const SchemasResponseSchema = z.record(z.string(), SchemasResponseSchemaSchema)
export type SchemasResponse = z.infer<typeof SchemasResponseSchema>

const SolverResponseRunSchema = <TDecoded extends z.ZodTypeAny = z.ZodRecord<z.ZodString, z.ZodUnknown>>(
	decodedSchema: TDecoded = z.record(z.string(), z.unknown()) as unknown as TDecoded
) =>
	z.object({
		id: z.string(),
		status: z.string(),
		result: z.string().optional(),
		error: z.string().optional(),
		errors: z.array(z.string()).optional(),
		gasEstimate: z.number().optional(),
		data: z
			.object({
				raw: z.instanceof(Uint8Array).optional(),
				decoded: decodedSchema.optional()
			})
			.optional(),
		intentId: z.string().optional(),
		createdAt: z.string()
	})
export type SolverResponseRun<TDecoded extends Record<string, unknown> = Record<string, unknown>> = z.infer<
	ReturnType<typeof SolverResponseRunSchema<z.ZodType<TDecoded>>>
>

const IntentResponseIntentSchema = <TDecoded extends z.ZodTypeAny = z.ZodRecord<z.ZodString, z.ZodUnknown>>(
	decodedSchema: TDecoded = z.record(z.string(), z.unknown()) as unknown as TDecoded
) =>
	z.object({
		id: z.string(),
		status: z.string(),
		chainId: z.number(),
		from: z.string(),
		inputs: SchemasRequestActionsSchema,
		frequency: z.number(),
		startAt: z.string(),
		endAt: z.string().optional(),
		periodEndAt: z.string().nullable().optional(),
		nextSimulationAt: z.string().nullable().optional(),
		runs: z.array(SolverResponseRunSchema(decodedSchema)),
		createdAt: z.string()
	})
export type IntentResponseIntent<TDecoded extends Record<string, unknown> = Record<string, unknown>> = z.infer<
	ReturnType<typeof IntentResponseIntentSchema<z.ZodType<TDecoded>>>
>

const SolutionResponseTransactionSchema = z.object({
	to: z.string() as z.ZodType<`0x${string}`>,
	data: z.string() as z.ZodType<`0x${string}`>,
	value: z.bigint(),
	meta: z.any()
})
export type SolutionResponseTransaction = z.infer<typeof SolutionResponseTransactionSchema>

const SolutionResponseRunSchema = <TDecoded extends z.ZodTypeAny = z.ZodRecord<z.ZodString, z.ZodUnknown>>(
	decodedSchema: TDecoded = z.record(z.string(), z.unknown()) as unknown as TDecoded
) =>
	z.object({
		id: z.string(),
		status: z.string(),
		from: z.string() as z.ZodType<`0x${string}`>,
		to: z.string() as z.ZodType<`0x${string}`>,
		value: z.bigint(),
		gasEstimate: z.number(),
		resultData: z.object({
			raw: z.string().optional(),
			decoded: decodedSchema.optional()
		}),
		intentId: z.string(),
		livePlugsId: z.string()
	})
export type SolutionResponseRun<TDecoded extends Record<string, unknown> = Record<string, unknown>> = z.infer<
	ReturnType<typeof SolutionResponseRunSchema<z.ZodType<TDecoded>>>
>

const IntentResponseSchema = <TDecoded extends z.ZodTypeAny = z.ZodRecord<z.ZodString, z.ZodUnknown>>(
	decodedSchema: TDecoded = z.record(z.string(), z.unknown()) as unknown as TDecoded
) =>
	z.object({
		transactions: z.array(SolutionResponseTransactionSchema),
		run: SolutionResponseRunSchema(decodedSchema).optional()
	})
export type IntentResponse<TDecoded extends Record<string, unknown> = Record<string, unknown>> = z.infer<
	ReturnType<typeof IntentResponseSchema<z.ZodType<TDecoded>>>
>

const KillResponseSchema = z.object({ killed: z.boolean() })
export type KillResponse = z.infer<typeof KillResponseSchema>

export const Schemas = {
	Request: {
		ValueManual: SchemasRequestValueManualSchema,
		ValueOptionExtension: SchemasRequestValueOptionExtensionSchema,
		Value: SchemasRequestValueSchema,
		Values: SchemasRequestValuesSchema,
		Action: SchemasRequestActionSchema,
		Actions: SchemasRequestActionsSchema
	},
	Response: {
		Icon: SchemasResponseIconSchema,
		Info: SchemasResponseInfoSchema,
		Chain: SchemasResponseChainSchema,
		Chains: SchemasResponseChainsSchema,
		Option: SchemasResponseOptionSchema,
		Options: SchemasResponseOptionsSchema,
		OptionsSet: SchemasResponseOptionsSetSchema,
		Coils: SchemasResponseCoilsSchema,
		Schema: SchemasResponseSchemaSchema,
		Response: SchemasResponseSchema,
		Kill: KillResponseSchema
	},
	Solver: {
		ResponseRun: SolverResponseRunSchema
	},
	Intent: {
		ResponseIntent: IntentResponseIntentSchema
	},
	Solution: {
		ResponseTransaction: SolutionResponseTransactionSchema,
		ResponseRun: SolutionResponseRunSchema,
		Response: IntentResponseSchema
	}
}
