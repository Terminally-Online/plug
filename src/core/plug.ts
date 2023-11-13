import { TypedData, TypedDataDomain, TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetTypedDataDomain,
	GetTypedDataPrimaryType,
	hashTypedData,
	recoverTypedDataAddress,
	WalletClient
} from 'viem'

import { TypedDataToLivePlug } from '@/lib/types'

export class Plug<
	C extends WalletClient,
	T extends TypedData,
	K extends GetTypedDataPrimaryType<T> = GetTypedDataPrimaryType<T>,
	U extends TypedDataToPrimitiveTypes<T>[K] = TypedDataToPrimitiveTypes<T>[K],
	S extends TypedDataToLivePlug<K, U> = TypedDataToLivePlug<K, U>
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

	// * Sign and initialize an intent in-framework.
	// ! This is only used when you have access to a `client` that can sign
	//   the messages such as a frontend. For an API, you would not
	//   consume this method in the Plug consumer or anywhere else.
	async init<
		P extends {
			client: C
		}
	>({ client }: P) {
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

	// * Recover the address from the message and signature of the intent.
	async address<
		P extends Partial<{
			domain: TypedDataDomain
			signature: `0x${string}`
		}>
	>({ domain, signature }: P = {} as P) {
		domain = domain ?? this.domain
		signature = signature ?? this.intent?.signature

		if (!signature) throw new Error('Signature not initialized')

		return await recoverTypedDataAddress({
			domain,
			types: this.types,
			primaryType: this.primaryType,
			message: this.message,
			signature
		})
	}

	// * Confirm the address of the message of intent.
	async verify<
		P extends Partial<{
			domain: TypedDataDomain
			signature: `0x${string}`
		}> & {
			address: `0x${string}`
		}
	>({ domain, address, signature }: P) {
		return (await this.address({ domain, signature })) === address
	}

	// * Hash the message of intent.
	hash<
		P extends Partial<{
			domain: TypedDataDomain
			message: U
		}>
	>({ domain, message }: P = {} as P) {
		domain = domain ?? this.domain
		message = message ?? this.message

		return hashTypedData({
			domain,
			types: this.types,
			primaryType: this.primaryType,
			message
		})
	}
}
