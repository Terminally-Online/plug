import { useEffect, useMemo, useState } from 'react'
import { createClient } from '@/lib'
import { mainnet } from 'viem/chains'
import { GetCodeReturnType } from 'viem'

export const useBytecode = (address: string) => {
	const [bytecode, setBytecode] = useState<GetCodeReturnType | null>(null)
	const [isLoading, setIsLoading] = useState(false)
	const [error, setError] = useState<Error | null>(null)

	const client = useMemo(() => createClient(mainnet.id), [])

	useEffect(() => {
		if (!address) return

		const getCode = async () => {
			try {
				setIsLoading(true)
				setError(null)
				const code = await client.getCode({
					address: address as `0x${string}`,
				})
				setBytecode(code)
			} catch (err) {
				console.error("Error fetching contract bytecode", err)
				setError(err instanceof Error ? err : new Error(String(err)))
			} finally {
				setIsLoading(false)
			}
		}

		getCode()
	}, [address, client])

	return { bytecode, error, isLoading, isError: error !== null }

}
