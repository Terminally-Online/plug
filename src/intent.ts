import { TypedDataToPrimitiveTypes } from 'abitype'

import {
	GetTypedDataDomain,
	GetTypedDataPrimaryType,
	hashTypedData,
	recoverTypedDataAddress,
	TypedData,
	WalletClient
} from 'viem'

import { TypedDataToKeysWithSignedPair } from './lib/types'

const DEBUG_TYPES = {
	Mail: [
		{ name: 'from', type: 'Person' },
		{ name: 'to', type: 'Person' },
		{ name: 'contents', type: 'string' }
	],
	Person: [
		{ name: 'name', type: 'string' },
		{ name: 'wallet', type: 'address' }
	],
	Email: [
		{ name: 'from', type: 'Person' },
		{ name: 'to', type: 'Person[]' },
		{ name: 'mail', type: 'SignedMail[]' }
	],
	SignedMail: [
		{ name: 'mail', type: 'Mail' },
		{ name: 'signature', type: 'bytes' }
	],
	SignedEmail: [
		{ name: 'email', type: 'Email' },
		{ name: 'signature', type: 'bytes' }
	]
} as const

export type IntentType<T> = Exclude<
	{
		[K in keyof T]: `Signed${Capitalize<string & K>}` extends keyof T
			? K
			: never
	}[keyof T],
	'EIP712Domain'
>

export class Intent<
	C extends WalletClient,
	T extends TypedData,
	K extends TypedDataToKeysWithSignedPair<T>,
	U extends TypedDataToPrimitiveTypes<T>[K] = TypedDataToPrimitiveTypes<T>[K],
	S extends Record<'signature', `0x${string}`> & {
		[TK in K as Lowercase<string & TK>]: U
	} = Record<'signature', `0x${string}`> & {
		[TK in K as Lowercase<string & TK>]: U
	}
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

		await this.client.signTypedData({
			account: this.client.account.address,
			domain: this.domain,
			types: DEBUG_TYPES,
			primaryType: 'Mail',
			message: {
				from: { name: 'Bob', wallet: '0x0' },
				to: { name: 'Alice', wallet: '0x0' },
				contents: 'Hello, world!'
			}
		})

		if (!this.message) return

		if (!this.domain) throw new Error('Domain not initialized')
		if (!this.types) throw new Error('Types not initialized')
		if (!this.message) throw new Error('Message not initialized')
		if (!this.primaryType) throw new Error('Primary type not initialized')
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
