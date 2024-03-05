import { TypedDataToPrimitiveTypes } from 'abitype'

import { GetContractReturnType, TypedDataDefinition, WalletClient } from 'viem'

import { Plug } from '@/src/core/plug'

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

export class PlugSDK<
	TClient extends WalletClient = WalletClient,
	TDomain extends
		TypedDataDefinition['domain'] = TypedDataDefinition['domain'],
	TMessage extends TypedDataToPrimitiveTypes<
		typeof PLUGS_TYPES
	>['Plugs'] = TypedDataToPrimitiveTypes<typeof PLUGS_TYPES>['Plugs']
> {
	public plugs: Plug[] = []

	build(
		domain: TDomain,
		contract: GetContractReturnType | `0x${string}`,
		message: TMessage
	): Plug<TClient> {
		const plug = new Plug<TClient>(
			{
				...domain,
				verifyingContract:
					typeof contract === 'string' ? contract : contract.address
			},
			message
		)

		this.plugs.push(plug)

		return plug
	}

	latest() {
		if (this.plugs.length == 0) return undefined

		return this.plugs[this.plugs.length - 1]
	}
}
