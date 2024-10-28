import { actions, categories } from "../constants"

export type Option = {
	icon: JSX.Element | undefined
	label: string
	value: string | number
	imagePath?: string
}

export type Value = string | Option | undefined | null

export type Action = {
	categoryName: keyof typeof categories
	actionName: keyof (typeof actions)[keyof typeof categories]
	values: Array<Value>
}

export type Actions = Array<Action>
