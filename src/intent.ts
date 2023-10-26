import {
    GetTypedDataDomain,
    TypedData,
    WalletClient,
    hashTypedData,
    recoverTypedDataAddress
} from 'viem'

import { IntentType, TypedIntent } from '@/lib/types'

export class Intent<
	TTypes extends TypedData,
	TIntentType extends IntentType<TTypes>,
	TIntent extends TypedIntent<TTypes>,
	TIntentTypeLower extends Lowercase<
		TIntentType extends string ? TIntentType : never
	> = Lowercase<TIntentType extends string ? TIntentType : never>
> {
	private client: WalletClient | undefined

	public intent:
		| (Record<'signature', `0x${string}`> &
				Record<TIntentTypeLower, TIntent>)
		| undefined

	constructor(
		public readonly domain: GetTypedDataDomain['domain'],
		public readonly types: TTypes,
		public readonly primaryType: TIntentType extends string
			? TIntentType
			: never,
		private readonly message: TIntent
	) {}

	lowercasePrimaryType() {
		return this.primaryType.toLowerCase() as TIntentTypeLower
	}

	async init(client: WalletClient, callback: (signedIntent: this) => void) {
		if (this.intent) return this

		this.client = client

		if (!this.client.account)
			throw new Error('Client account not initialized')

		this.intent = {
			[this.lowercasePrimaryType()]: this.message,
			signature: await this.client
				.signTypedData({
					account: this.client.account.address,
					domain: this.domain,
					types: this.types,
					primaryType: this.primaryType,
					message: this.message
				})
				.catch(error => {
					throw new Error(error)
				})
		} as this['intent']

		callback(this)
	}

	async address(signature = this.intent?.signature) {
		if (!signature) throw new Error('Signature not initialized')

		return await recoverTypedDataAddress({
			domain: this.domain,
			types: this.types,
			primaryType: this.primaryType,
			message: this.message,
			signature
		})
	}

	async verify(address: `0x${string}`) {
		return (await this.address()) === address
	}

	hash(message = this.message) {
		if (!message) throw new Error('Message not initialized')

		return hashTypedData({
			domain: this.domain,
			types: this.types,
			primaryType: this.primaryType,
			message
		})
	}
}
