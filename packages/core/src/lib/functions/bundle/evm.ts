import { EVM_TYPES } from "@/src/lib"

export const getEVMSchema = (type: string): string => {
	if (type.includes("[]")) {
		return `z.array(${getEVMSchema(type.replace("[]", ""))})`
	}

	for (const [regex, schema] of EVM_TYPES) {
		if (typeof regex === "string") {
			if (regex === type) return schema
		} else if (regex.test(type)) return schema
	}

	return `${type}Schema`
}
