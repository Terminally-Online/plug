import Image from "next/image"
import { FC, HTMLAttributes } from "react"

import { zeroAddress } from "viem"
import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"
import { useEnsAddress, useEnsAvatar, useEnsName } from "wagmi"

import { Avatar } from "@/components/app/sockets/profile"
import { Accordion } from "@/components/shared/utils/accordion"
import { formatAddress, greenGradientStyle } from "@/lib"
import { useAccount } from "@/lib/hooks/account/useAccount"

export const TransferRecipient: FC<
	HTMLAttributes<HTMLDivElement> & {
		isRecent?: boolean
		address: string
		handleSelect: (address: string) => void
	}
> = ({ isRecent = false, address = "", handleSelect, className }) => {
	const account = useAccount()

	const { data: ensAddress } = useEnsAddress({
		name: normalize(address),
		query: {
			enabled: address?.endsWith("eth") || false
		},
		chainId: mainnet.id
	})
	const {
		data: ensName,
		isFetching,
		isFetched
	} = useEnsName({
		address: ensAddress ?? (address as `0x${string}`),
		query: {
			enabled: (ensAddress ?? address ?? "").startsWith("0x") === true
		},
		chainId: mainnet.id
	})
	const { data: ensAvatar } = useEnsAvatar({
		name: ensName ?? "",
		query: {
			enabled: ensName !== undefined || address?.endsWith(".eth") || false
		},
		chainId: mainnet.id
	})

	const Badge = () => {
		if (isRecent)
			return (
				<p className="rounded-md bg-plug-green/10 px-2 py-1 text-sm">
					<span style={{ ...greenGradientStyle }}>Recent</span>
				</p>
			)
		if (account.address === address)
			return (
				<p className="rounded-md bg-plug-green/5 px-2 py-1 text-sm">
					<span className="opacity-40">Connected</span>
				</p>
			)
		return null
	}

	return (
		<Accordion className={className} onExpand={() => handleSelect(address)}>
			<div className="flex flex-row items-center gap-4">
				<div className="relative h-10 w-10 min-w-10 rounded-sm">
					{ensAvatar ? (
						<>
							<Image
								className="absolute left-1/2 top-1/2 h-12 w-48 -translate-x-1/2 blur-2xl filter"
								src={ensAvatar}
								alt="ENS Avatar"
								width={240}
								height={240}
							/>
							<Image
								className="relative rounded-sm"
								src={ensAvatar}
								alt="ENS Avatar"
								width={240}
								height={240}
							/>
						</>
					) : (
						<>
							<div className="absolute left-0 top-1/2 blur-lg filter">
								<Avatar name={address} className="rounded-sm" />
							</div>
							<Avatar name={address} className="rounded-sm" />
						</>
					)}
				</div>
				<div className="flex w-max flex-col truncate overflow-ellipsis text-left">
					<p className="font-bold">
						{isFetching && isFetched === false
							? "Loading..."
							: isFetched
								? ((address.includes("eth") ? address : ensName) ?? formatAddress(address))
								: address === ""
									? "Enter your search"
									: "No match found"}
					</p>
					<p className="text-sm font-bold opacity-40">
						{isFetching && isFetched === false
							? formatAddress(zeroAddress)
							: address.endsWith(".eth") && !ensName && isFetched === false
								? "No ENS record found for this address."
								: address.endsWith(".eth") || ensName
									? formatAddress(ensAddress ?? address)
									: "External account"}
					</p>
				</div>
				<div className="my-auto ml-auto w-max font-bold">
					<Badge />
				</div>
			</div>
		</Accordion>
	)
}
