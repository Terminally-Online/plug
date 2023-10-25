import { TypedDataToPrimitiveTypes } from 'abitype'

import { TypedData } from 'viem'

import { constants } from '@nftchance/emporium-types'

type ConstantTypes = typeof constants.types

export type Types<TTypes extends ConstantTypes = ConstantTypes> = TTypes

// Our types constant is declared like:
// const types = {
//     Transaction: [
// 		{ name: 'to', type: 'address' },
// 		{ name: 'gasLimit', type: 'uint256' },
// 		{ name: 'data', type: 'bytes' }
// 	],
// 	SignedDelegation: [
// 		{ name: 'delegation', type: 'Delegation' },
// 		{ name: 'signature', type: 'bytes' }
// 	],
// 	Invocation: [
// 		{ name: 'transaction', type: 'Transaction' },
// 		{ name: 'authority', type: 'SignedDelegation[]' }
// 	]
// } as const
// And we want to narrow the type only to the keys of the types constant that have
// an object in the array with the name 'signature'. This means, SignedTypes should
// SignedDelegation and not Invocation or Transaction.

export type FilterKeysWithSigned<T> = {
	[K in keyof T]: `Signed${Capitalize<string & K>}` extends keyof T
		? K
		: never
}[keyof T]

export type TypedIntent<TTypes extends TypedData> =
	TypedDataToPrimitiveTypes<TTypes>[keyof TypedDataToPrimitiveTypes<TTypes>]

export type SignedTypeIntents<TTypes extends TypedData> = {
	[K in keyof TypedDataToPrimitiveTypes<TTypes> as K extends `Signed${string}`
		? K
		: never]: TypedDataToPrimitiveTypes<TTypes>[K]
}

export type SignedTypeIntent<TTypes extends TypedData> =
	SignedTypeIntents<TTypes>[keyof SignedTypeIntents<TTypes>]
