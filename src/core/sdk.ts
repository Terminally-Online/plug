import { GetContractReturnType, WalletClient } from 'viem'

import { Plug } from '@/src/core/plug'
import { Domain, PlugPrimitiveTypes } from '@/src/lib/types/typedData'

export class PlugSDK<
	TClient extends WalletClient = WalletClient,
	TDomain extends Domain = Domain,
	TMessage extends PlugPrimitiveTypes = PlugPrimitiveTypes
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
