import { TypedData, TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetTypedDataDomain,
	GetTypedDataPrimaryType,
	hashTypedData,
	recoverTypedDataAddress,
	WalletClient
} from 'viem'

import {
	TypedDataToKeysWithSignedPair,
	TypedDataToSignedIntent
} from '@/lib/types'

export class Intent<
	C extends WalletClient,
	T extends TypedData,
	K extends TypedDataToKeysWithSignedPair<T>,
	U extends TypedDataToPrimitiveTypes<T>[K] = TypedDataToPrimitiveTypes<T>[K],
	S extends TypedDataToSignedIntent<K, U> = TypedDataToSignedIntent<K, U>
> {
	private client?: WalletClient
	public intent: S | undefined

	constructor(
		public readonly domain: GetTypedDataDomain['domain'],
		public readonly types: T,
		public readonly primaryType: K extends string
			? GetTypedDataPrimaryType<T, K>
			: never,
		public readonly message: U
	) {}

	lowercasePrimaryType() {
		return this.primaryType.toLowerCase()
	}

	async init(client: C) {
		if (this.intent) return this

		this.client = client

		if (!this.client.account)
			throw new Error('Client account not initialized')

		const signature = await this.client.signTypedData({
			account: this.client.account.address,
			domain: this.domain,
			types: this.types,
			primaryType: this.primaryType,
			message: this.message
		})

		this.intent = {
			[this.lowercasePrimaryType()]: this.message,
			signature
		} as S

		return this
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
