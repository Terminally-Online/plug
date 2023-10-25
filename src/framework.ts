import {
	GetContractReturnType,
	GetTypedDataDomain,
	TypedData,
	WalletClient
} from 'viem'

import { Intent } from './intent'
import { IntentType, TypedIntent } from './lib/types'

export class Framework<
	TTypes extends TypedData,
	TIntentType extends IntentType<TTypes> = IntentType<TTypes>,
	TIntent extends TypedIntent<TTypes> = TypedIntent<TTypes>
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

		const types = Object.fromEntries(
			Object.entries(this.info.types).filter(
				([key]) => key !== 'EIP712Domain'
			)
		) as TTypes

		return new Intent<TTypes, TIntentType, TIntent>(
			this.info.domain,
			types,
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
