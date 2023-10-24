import {
	AddressLike,
	Signer,
	TypedDataDomain,
	TypedDataEncoder,
	verifyTypedData
} from 'ethers'

import { SignedTypes } from './lib/types'

export class Intent<TIntent, TSignedIntent extends SignedTypes> {
	encoder: TypedDataEncoder
	signer: Signer | null = null
	signedMessage: TSignedIntent | null = null

	constructor(
		public readonly domain: TypedDataDomain,
		public readonly types: Record<string, Array<TypedDataField>>,
		public readonly message: TIntent
	) {
		// * Remove the EIP712Domain type from the types to avoid unused
		//   elements within the user-defined DAG.
		this.types = Object.fromEntries(
			Object.entries(types).filter(([key]) => key !== 'EIP712Domain')
		)

		/// * Go ahead and instantiate the encoder.
		this.encoder = new TypedDataEncoder(
			this.types as Record<string, Array<TypedDataField>>
		)
	}

	private _primaryType(): keyof TSignedIntent {
		// TODO: Should solve for this instead of requiring the manual declaration.
		// If one is not found, throw instead of returning undefined.

		throw new Error('Not implemented.')
	}

	async init(
		signer: Signer | null,
		callback: (
			error: Error | null,
			signedIntent: TSignedIntent | null
		) => TSignedIntent
	) {
		if (this.signedMessage) return this

		if (!signer) throw new Error('Signer not initialized')

		this.signer = signer

		this.signer
			.signTypedData(
				this.domain,
				// ! Hacky way to get around using abitype with Ethers.
				this.types as unknown as Record<string, Array<TypedDataField>>,
				this.message as Record<string, unknown>
			)
			.then(signature => {
				const signedIntent: TSignedIntent = {
					[this._primaryType()]: this.message,
					signature
				}

				this.signedMessage = signedIntent

				callback(null, this.signedMessage)
			})
			.catch(error => {
				callback(error, null)
			})

		return this
	}

	address(signature?: string) {
		// * If a manual signature was not provided, go ahead and retrieve the built one.
		if (signature === undefined) signature = this.signedMessage?.signature

		if (signature === null) throw new Error('Signature not initialized')

		return verifyTypedData(
			this.domain,
			this.types as unknown as Record<string, Array<TypedDataField>>,
			this.message,
			signature
		)
	}

	verify(address: AddressLike | `0x${string}`) {
		return this.address() === address
	}

	hash(message?: TIntent) {
		if (message === undefined) message = this.message

		if (message === undefined) throw new Error('Message not initialized')

		return TypedDataEncoder.hash(
			this.domain,
			this.types as unknown as Record<string, Array<TypedDataField>>,
			message
		)
	}
}
