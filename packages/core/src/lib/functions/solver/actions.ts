import { ActionProvider, actions } from "@/src/lib"

export const getProvider = (key: `${string}.${string}.${string}`) => {
	const [serviceKey, providerKey, actionKey] = key.split(".")

	if (!serviceKey || !providerKey || !actionKey) return

	return actions[serviceKey]?.[providerKey]
}

export const getAction = (actionProvider: ActionProvider, key: string) => {
	return actionProvider.actions[key]
}
