import { useEffect, useState } from "react"

import { base } from 'viem/chains'
import { createClient, NATIVE_TOKEN_ADDRESS } from '../../constants'

const client = createClient(base.id)

const erc20Abi = [
	{
		inputs: [{ name: "owner", type: "address" }, { name: "spender", type: "address" }],
		name: "allowance",
		outputs: [{ name: "amount", type: "uint256" }],
		stateMutability: "view",
		type: "function",
	},
] as const;

export type UseAllowanceProps = { token: string, owner: string, spender?: string, decimals?: number }

export const useAllowance = ({ token, owner, spender }: UseAllowanceProps) => {
	const [allowance, setAllowance] = useState(BigInt(0))

	useEffect(() => {
		const getApproval = async () => {
			if (token === NATIVE_TOKEN_ADDRESS || !spender) return

			const amount = await client.readContract({
				address: token as `0x${string}`,
				abi: erc20Abi,
				functionName: 'allowance',
				args: [owner as `0x${string}`, spender as `0x${string}`]
			})
			setAllowance(amount)
		}

		getApproval()
	}, [token, owner, spender])

	return { approval: allowance }
}
