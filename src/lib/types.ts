import { TypedDataParameter } from 'abitype'

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

export type SignedTypes<TTypes extends ConstantTypes = ConstantTypes> = {
	[TKey in keyof TTypes]: TTypes[TKey] extends Array<infer TType>
		? TType extends TypedDataParameter
			? TType['name'] extends 'signature'
				? TKey
				: never
			: never
		: never
}[keyof TTypes]
