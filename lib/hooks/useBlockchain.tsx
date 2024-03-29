"use client"

import { useCallback, useMemo } from "react"

import { blockExplorerAddress, truncateAddress } from "@/lib/blockchain"

export const useBlockchain = ({ address }: { address: string }) => {
	const displayAddress = useMemo(() => truncateAddress(address), [address])

	const blockExplorer = useCallback(
		(chainId: number) => {
			blockExplorerAddress(chainId, address)
		},
		[address]
	)

	return { address, displayAddress, blockExplorer }
}

export default useBlockchain
