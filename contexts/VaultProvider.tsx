"use client"

import type { FC, PropsWithChildren } from "react"
import {
	createContext,
	useCallback,
	useContext,
	useEffect,
	useMemo,
	useState
} from "react"

import { WalletProvider } from "@/contexts"
import { api } from "@/lib/api"
import { Vault } from "@/server/api/routers/vault"

export type FlattenedVault = Omit<Vault, "chainId"> & {
	chainIds: [number]
}

export const VaultContext = createContext<{
	vault: Vault | undefined
	vaults: Array<Vault> | undefined
	handleAdd: (chainIds: Array<number>) => void
	handleSelect: (address: string) => void
	handleDeploy: (chainIds: Array<number>, version?: number) => void
}>({
	vault: undefined,
	vaults: undefined,
	handleAdd: () => {},
	handleSelect: () => {},
	handleDeploy: () => {}
})

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const VaultProvider: FC<PropsWithChildren> = ({ children }) => {
	const { data: apiVaults } = api.account.vaults.all.useQuery()

	const [vaults, setVaults] = useState<Array<Vault> | undefined>(apiVaults)
	const [vaultAddress, handleSelect] = useState<string | undefined>(undefined)

	// const { data: blockNumber } = useBlockNumber()
	const blockNumber = 0

	// TODO: Right now we have multiple of the same address on different chains and
	//       we need to flatten it so that it has multiple chain ids.
	const vault = useMemo(() => {
		if (!vaults) return undefined

		return vaults.find(vault => vault.address === vaultAddress)
	}, [vaults, vaultAddress])

	const handleVaultAdd = api.account.vaults.add.useMutation()

	const handleAdd = useCallback(
		(chainIds: Array<number>) =>
			handleVaultAdd.mutate({
				address: TEMPORARY_ADDRESS,
				chainIds: chainIds,
				lastBlockIndexed: blockNumber
			}),
		[blockNumber]
	)

	// * Update the local state of vaults to not need a refetch.
	const onAdd = (data: Array<Vault>) =>
		setVaults(prev => (prev ? [...prev, ...data] : [...data]))

	// * Loop through all of the incoming vaults and update the local state
	//   for all of the vaults that were updated.
	const onRename = (data: Array<Vault>) => {
		setVaults(prev => {
			const vaults: Array<Vault> = []

			for (const vault of prev || []) {
				const updated = data.find(
					({ address, chainId }) =>
						address === vault.address && chainId === vault.chainId
				)

				vaults.push(updated || vault)
			}

			return vaults
		})
	}

	api.account.vaults.onAdd.useSubscription(undefined, {
		onData: onAdd
	})

	api.account.vaults.onRename.useSubscription(undefined, {
		onData: onRename
	})

	// TODO: Use the contract write to deploy a new vault and recover
	//		 the information from the Vault.
	const handleDeploy = (chainIds: Array<number>, version: number = 0) => {}

	useEffect(() => {
		if (vaultAddress !== undefined || apiVaults === undefined) return

		handleSelect(apiVaults[0].address)
	}, [apiVaults, vaultAddress])

	return (
		<WalletProvider>
			<VaultContext.Provider
				value={{
					vault,
					vaults,
					handleAdd,
					handleSelect,
					handleDeploy
				}}
			>
				{children}
			</VaultContext.Provider>
		</WalletProvider>
	)
}

export const useVaults = () => useContext(VaultContext)
