import { hashTypedData, recoverTypedDataAddress, WalletClient } from 'viem'

import { PLUGS_TYPES } from '@nftchance/plug-types'

import { PlugAPI } from '@/src/core/api'
import {
	Domain,
	PlugTypedIntent,
	PlugTypedMessage
} from '@/src/lib/types/typedData'

export class Plug<
	TClient extends WalletClient = WalletClient,
	TDomain extends Domain = Domain,
	TMessage extends PlugTypedMessage = PlugTypedMessage,
	TIntent extends PlugTypedIntent<TMessage> = PlugTypedIntent<TMessage>
> {
	public readonly types: typeof PLUGS_TYPES
	public readonly primaryType: keyof typeof PLUGS_TYPES
	public client?: TClient
	public intent?: TIntent
	public apiClient: PlugAPI

	constructor(
		public readonly domain: NonNullable<TDomain>,
		public readonly message: NonNullable<TMessage>,
		private readonly api = 'https://api.onplug.io/pool',
		private readonly apiKey = 'AAAAAAAAAAAAAAAAAAAA'
	) {
		this.types = PLUGS_TYPES
		this.primaryType = 'Plugs'
		this.apiClient = new PlugAPI(this.api, this.apiKey)
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

		return await this.apiClient.post(body)
	}
}
