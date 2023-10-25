import { TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetContractReturnType,
	GetTypedDataDomain,
	TypedData,
	WalletClient
} from 'viem'

import { Intent } from './intent'
import { FilterKeysWithSigned } from './lib/types'

export class Framework<
	// * All of the EIP-712 types.
	TTypes extends TypedData,
	// * The key of of the EIP-712 types that have a `Signed` pair.
	TIntentType extends
		FilterKeysWithSigned<TTypes> = FilterKeysWithSigned<TTypes>,
	// * The EIP-712 type that is being signed.
	TIntent extends
		TypedDataToPrimitiveTypes<TTypes>[TIntentType] = TypedDataToPrimitiveTypes<TTypes>[TIntentType]
> {
	public info: {
		domain: GetTypedDataDomain['domain']
		types: TTypes
	} | null = null

	public signedIntents: Array<Intent<TTypes, TIntentType, TIntent>> = []

	constructor(public readonly contract: GetContractReturnType) {}

	init(name: string, version: string, chainId: number, types: TTypes) {
		this.info = {
			domain: {
				chainId,
				verifyingContract: this.contract.address,
				name,
				version
			},
			types
		}

		return this
	}

	build(
		intentType: TIntentType extends string ? TIntentType : never,
		intent: TIntent
	) {
		if (!this.info) throw new Error('Contract info not initialized')

		return new Intent<TTypes, TIntentType, TIntent>(
			this.info.domain,
			this.info.types,
			intentType,
			intent
		)
	}

	async sign(
		client: WalletClient,
		...[intentType, intent]: Parameters<this['build']>
	) {
		// * Build the intent and initialize it.
		await this.build(intentType, intent)
			.init(client, signedIntent => this.signedIntents.push(signedIntent))
			.catch(error => {
				throw new Error(`Signed intent not initialized: ${error}`)
			})

		// * Return the latest.
		return this.signedIntents[this.signedIntents.length - 1]
	}
}
