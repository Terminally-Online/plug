import { TypedData, TypedDataDomain, TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetContractReturnType,
	GetTypedDataDomain,
	GetTypedDataPrimaryType,
	WalletClient
} from 'viem'

import { Intent } from '@/core/intent'

import { TypedDataToKeysWithSignedPair } from '../lib/types'

export class Framework<
	C extends WalletClient,
	T extends TypedData,
	K extends TypedDataToKeysWithSignedPair<T>
> {
	public readonly info: {
		domain: GetTypedDataDomain['domain']
		types: T
	} | null = null

	constructor(
		name: string,
		version: string,
		chainId: number,
		types: T,
		public readonly contract: GetContractReturnType | `0x${string}`
	) {
		contract = typeof contract === 'string' ? contract : contract.address

		this.info = {
			domain: {
				chainId,
				verifyingContract: contract,
				name,
				version
			},
			types
		}
	}

	build<TK extends GetTypedDataPrimaryType<T>>(
		intentType: TK extends string ? GetTypedDataPrimaryType<T, TK> : never,
		intent: TypedDataToPrimitiveTypes<T>[TK],
		domain?: TypedDataDomain
	) {
		domain = domain || this.info?.domain

		if (!this.info) throw new Error('Contract info not initialized')

		return new Intent<C, T, TK>(
			domain,
			this.info.types,
			intentType,
			intent as TypedDataToPrimitiveTypes<T>[TK]
		)
	}

	async sign<TK extends K & GetTypedDataPrimaryType<T>>(
		client: C,
		// * Without the explicit declaration here, one can include fields that do
		//   not belong to the intent type as it would use union. We can only break
		//   out of the union after that type has been declared such as `intentType`
		//   setting a reference for the `TK` type.
		intentType: TK extends string
			? GetTypedDataPrimaryType<T, TK> & K
			: never,
		intent: TypedDataToPrimitiveTypes<T>[TK],
		domain?: TypedDataDomain
	) {
		// * Build the intent and initialize it.
		return (await this.build(
			intentType,
			intent as TypedDataToPrimitiveTypes<T>[TK] &
				TypedDataToPrimitiveTypes<T>[K],
			domain
		).init({ client })) as Intent<C, T, TK>
	}
}
