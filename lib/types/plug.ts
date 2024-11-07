export type ActionSchema = {
	metadata: {
		icon: string
	}
	schema: {
		[action: string]: {
			sentence: string
			fields: {
				name: string
				type: string
				options?: {
					value: string
					name: string
					label: string
					info?: string
					icon: string
				}[]
			}[]
		}
	}
}

export type ActionSchemas = {
	[protocol: string]: ActionSchema
}
