"use client"

import { useCallback, useMemo } from "react"

import { formatAddress, getBlockExplorerAddress } from "@/lib"

export const useBlockchain = ({ address }: { address: string }) => {
	const displayAddress = useMemo(() => formatAddress(address), [address])

	const blockExplorer = useCallback(
		(chainId: number) => {
			getBlockExplorerAddress(chainId, address)
		},
		[address]
	)

	return { address, displayAddress, blockExplorer }
}

export default useBlockchain
