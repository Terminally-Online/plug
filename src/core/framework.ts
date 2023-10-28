import { TypedData, TypedDataToPrimitiveTypes } from 'abitype'

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
		public readonly contract: GetContractReturnType
	) {
		this.info = {
			domain: {
				chainId,
				verifyingContract: this.contract.address,
				name,
				version
			},
			types
		}
	}

	build<TK extends K>(
		intentType: TK extends string
			? GetTypedDataPrimaryType<T, TK> & K
			: never,
		intent: TypedDataToPrimitiveTypes<T>[K]
	) {
		if (!this.info) throw new Error('Contract info not initialized')

		return new Intent<C, T, K>(
			this.info.domain,
			this.info.types,
			intentType,
			intent
		)
	}

	async sign<TK extends K>(
		client: C,
		// * Without the explicit declaration here, one can include fields that do
		//   not belong to the intent type as it would use union. We can only break
		//   out of the union after that type has been declared such as `intentType`
		//   setting a reference for the `TK` type.
		intentType: TK extends string
			? GetTypedDataPrimaryType<T, TK> & K
			: never,
		intent: TypedDataToPrimitiveTypes<T>[TK]
	) {
		// * Build the intent and initialize it.
		return (await this.build(
			intentType,
			intent as TypedDataToPrimitiveTypes<T>[K]
		).init(client)) as Intent<C, T, TK>
	}
}
