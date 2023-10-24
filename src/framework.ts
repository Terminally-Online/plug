import { ContractTransactionResponse, Signer, TypedDataDomain } from 'ethers'

import { TypedDataToPrimitiveTypes } from 'abitype'

import { constants } from '@nftchance/emporium-types'

import { Intent } from './intent'
import { Types } from './lib/types'

type Contract = {
	deploymentTransaction(): ContractTransactionResponse
	getAddress(): Promise<string>
}

export class Framework {
	info: {
		domain: TypedDataDomain
		types: Types
	} | null = null
	signedIntents: Array<unknown> = []

	constructor(public readonly contract: Contract) {}

	async init(
		name: string,
		version: string,
		chainId?: bigint,
		types = constants.types
	) {
		chainId = this.contract.deploymentTransaction()?.chainId

		if (!chainId) throw new Error('Chain ID not found')

		this.contract.getAddress().then(verifyingContract => {
			this.info = {
				domain: {
					chainId,
					verifyingContract,
					name,
					version
				},
				types
			}
		})

		return this
	}

	build(intent: TypedDataToPrimitiveTypes<Types>[keyof Types]) {
		if (!this.info) throw new Error('Contract info not initialized')

		return new Intent(this.info.domain, this.info.types, intent)
	}

	async sign<TType extends TypedDataToPrimitiveTypes<Types>[keyof Types]>(
		signer: Signer,
		intent: TType
	) {
		return this.build(intent).init(signer, (signedIntent: {}) => {
			if (!signedIntent) throw new Error('Signed intent not initialized')

			this.signedIntents.push(signedIntent)

			return signedIntent
		})
	}
}
