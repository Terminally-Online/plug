import { TypedData, TypedDataDomain, TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetContractReturnType,
	GetTypedDataDomain,
	GetTypedDataPrimaryType,
	WalletClient
} from 'viem'

import { Plug } from '@/src/core/plug'

import { Domain, TypedDataToKeysWithLivePair } from '../../lib/types'

export class PlugSDK<
	C extends WalletClient,
	T extends TypedData,
	K extends TypedDataToKeysWithLivePair<T>
> {
	public readonly info: {
		domain: GetTypedDataDomain['domain']
		types: T
	} | null = null

	constructor(
		domain: Domain,
		public readonly contract: GetContractReturnType | `0x${string}`,
		types: T
	) {
		this.info = {
			// Initialize the domain that is used in the domain hash onchain.
			domain: {
				...domain,
				verifyingContract:
					typeof contract === 'string' ? contract : contract.address
			},
			// Declare the types that are used in the typed data.
			types
		}
	}

	// Build a Plug intent without signing it.
	build<TK extends GetTypedDataPrimaryType<T>>(
		intentType: TK extends string ? GetTypedDataPrimaryType<T, TK> : never,
		intent: TypedDataToPrimitiveTypes<T>[TK],
		domain?: TypedDataDomain
	): Plug<C, T, TK> {
		domain = domain || this.info?.domain

		if (!this.info) throw new Error('Contract info not initialized')

		return new Plug<C, T, TK>(
			domain,
			this.info.types,
			intentType,
			intent as TypedDataToPrimitiveTypes<T>[TK]
		)
	}

	// Initialize and sign a pre-built Plug intent.
	async sign<TK extends K & GetTypedDataPrimaryType<T>>(
		client: C,
		plug: Plug<C, T, TK>
	): Promise<Plug<C, T, TK>> {
		return await plug.init({ client })
	}
}
