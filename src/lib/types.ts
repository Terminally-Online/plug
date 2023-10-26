import { TypedDataToPrimitiveTypes } from 'abitype'

import { TypedData } from 'viem'

export type IntentType<T> = Exclude<
	{
		[K in keyof T]: `Signed${Capitalize<string & K>}` extends keyof T
			? K
			: never
	}[keyof T],
	'EIP712Domain'
>

export type TypedIntents<TTypes extends TypedData> =
	TypedDataToPrimitiveTypes<TTypes>

export type TypedIntent<TTypes extends TypedData> =
	TypedIntents<TTypes>[IntentType<TTypes>]
