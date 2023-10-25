import { TypedDataToPrimitiveTypes } from 'abitype'

import {
	Account,
	GetTypedDataDomain,
	hashTypedData,
	recoverTypedDataAddress,
	TypedData,
	WalletClient
} from 'viem'

import { FilterKeysWithSigned, SignedTypeIntent } from './lib/types'

export class Intent<
	TTypes extends TypedData,
	TIntentType extends
		FilterKeysWithSigned<TTypes> = FilterKeysWithSigned<TTypes>,
	// * The EIP-712 type that is being signed.
	TIntent extends
		TypedDataToPrimitiveTypes<TTypes>[TIntentType] = TypedDataToPrimitiveTypes<TTypes>[TIntentType]
> {
	client: WalletClient | undefined
	account: Account | undefined
	struct: SignedTypeIntent<TTypes> | undefined
	signature: `0x${string}` | undefined

	constructor(
		public readonly domain: GetTypedDataDomain['domain'],
		public readonly types: TTypes,
		public readonly primaryType: TIntentType extends string
			? TIntentType
			: never,
		public readonly message: TIntent
	) {
		// * Remove the EIP712Domain type from the types to avoid unused
		//   elements within the user-defined DAG.
		// this.types = Object.fromEntries(
		// 	Object.entries(types).filter(([key]) => key !== 'EIP712Domain')
		// )
	}

	async init(client: WalletClient, callback: (signedIntent: this) => void) {
		client
		callback
		if (this.struct) return this

		this.client = client

		if (!this.client.account)
			throw new Error('Client account not initialized')

		const signature = await this.client
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
		if (!signature) throw new Error('Signature not initialized')

		this.struct = {
			// delegation: this.message,
			signature: '0x0'
		}

		callback(this)
	}

	async address(signature = this.signature) {
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
