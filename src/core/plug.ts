import { GetTypedDataLivePlugs } from '@/lib/types/typedData'

import {
	hashTypedData,
	recoverTypedDataAddress,
	TypedDataDefinition,
	WalletClient
} from 'viem'
import { API } from './api'

export const PLUGS_TYPES = {
	Current: [
		{ name: 'ground', type: 'address' },
		{ name: 'voltage', type: 'uint256' },
		{ name: 'data', type: 'bytes' }
	],
	Fuse: [
		{ name: 'neutral', type: 'address' },
		{ name: 'live', type: 'bytes' }
	],
	Plug: [
		{ name: 'current', type: 'Current' },
		{ name: 'fuses', type: 'Fuse[]' }
	],
	Plugs: [
		{ name: 'plugs', type: 'Plug[]' },
		{ name: 'salt', type: 'bytes32' }
	]
} as const

export class Plug<
	TClient extends WalletClient = WalletClient,
	TDomain extends
		TypedDataDefinition['domain'] = TypedDataDefinition['domain'],
	TMessage extends TypedDataDefinition<
		typeof PLUGS_TYPES,
		'Plugs'
	>['message'] = TypedDataDefinition<typeof PLUGS_TYPES, 'Plugs'>['message'],
	TIntent extends GetTypedDataLivePlugs<'Plugs', TMessage> | undefined =
		| GetTypedDataLivePlugs<'Plugs', TMessage>
		| undefined
> {
	public readonly types: typeof PLUGS_TYPES
	public readonly primaryType: keyof typeof PLUGS_TYPES
	public client?: TClient
	public intent?: TIntent
	public apiClient: API

	constructor(
		public readonly domain: NonNullable<TDomain>,
		public readonly message: NonNullable<TMessage>,
		private readonly api = 'https://api.onplug.io/pool',
		private readonly apiKey = 'AAAAAAAAAAAAAAAAAAAA'
	) {
		this.types = PLUGS_TYPES
		this.primaryType = 'Plugs'
		this.apiClient = new API(this.api, this.apiKey)
	}

	lowercasePrimaryType() {
		return this.primaryType.toLowerCase()
	}

	async sign(client: TClient) {
		if (this.intent) return this.intent

		this.client = client

		if (!this.client?.account)
			throw new Error('Client has no initialized account.')

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
		} as NonNullable<typeof this.intent>

		return this.intent
	}

	async address(
		domain = this.domain,
		signature: `0x${string}` | undefined = this.intent?.signature
	) {
		if (!signature) throw new Error('Intent has no initialized signature')

		return await recoverTypedDataAddress({
			domain,
			types: this.types,
			primaryType: this.primaryType,
			message: this.message,
			signature
		})
	}

	async verify(signature: `0x${string}`, address: `0x${string}`) {
		return (await this.address(this.domain, signature)) === address
	}

	hash() {
		return hashTypedData({
			domain: this.domain,
			types: this.types,
			primaryType: this.primaryType,
			message: this.message
		})
	}

	async submit() {
		if (!this.client?.account)
			throw new Error('Client has no initialized account.')
		if (!this.intent) throw new Error('Plug has no initialized intent.')

		const body = {
			account: this.client.account.address,
			intent: this.intent
		}

		return await this.apiClient.post(body);
	}
}
