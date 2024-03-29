import {
	type FC,
	type PropsWithChildren,
	useEffect,
	useMemo,
	useState
} from "react"

import Image from "next/image"

import { useAccount, useBalance, useDisconnect, useSwitchChain } from "wagmi"

import {
	ExclamationTriangleIcon,
	ExitIcon,
	GlobeIcon
} from "@radix-ui/react-icons"

import { useDomain } from "@/contexts/DomainProvider"
import {
	blockExplorerAddress,
	chainImage,
	formatName,
	truncateBalance
} from "@/lib/blockchain"

export const Wallet: FC<PropsWithChildren> = () => {
	const { address } = useAccount()

	const { accessible, chainId, domain, handleDomain } = useDomain()

	const { switchChain } = useSwitchChain()
	const { disconnect } = useDisconnect()

	const { data } = useBalance({ address, chainId: 1 })
	const { decimals, symbol, value } = data ?? {}

	const [isDisconnecting, setIsDisconnecting] = useState(false)

	const balance = useMemo(
		() => truncateBalance(value, decimals),
		[value, decimals]
	)

	const blockExplorer = useMemo(
		() => blockExplorerAddress(chainId, address),
		[chainId, address]
	)

	useEffect(() => {
		if (isDisconnecting) {
			const timeout = setTimeout(() => {
				setIsDisconnecting(false)
			}, 5000)

			return () => clearTimeout(timeout)
		}
	}, [isDisconnecting])

	if (!address) return null

	return (
		<div className="mt-[-46px] h-screen w-[360px]">
			<div className="m-auto mt-[46px] flex w-full py-16 text-center text-white">
				<h1 className="mx-auto flex flex-row text-4xl">
					{balance} <span className="ml-2 opacity-60">{symbol}</span>
				</h1>
			</div>

			<div className="flex w-full flex-row border-y-[1px] border-stone-950 text-center">
				<a
					target="_blank"
					rel="noreferrer"
					href={blockExplorer}
					className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950`}
				>
					<GlobeIcon
						className="mr-2 opacity-60"
						width={16}
						height={16}
					/>
					Explorer
				</a>

				<button
					onClick={() =>
						isDisconnecting
							? disconnect()
							: setIsDisconnecting(true)
					}
					className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-red-500 active:text-stone-950 ${
						isDisconnecting ? "active" : ""
					}`}
				>
					<ExitIcon
						className="mr-2 opacity-60"
						width={16}
						height={16}
					/>
					{isDisconnecting ? "Confirm" : "Log out"}
				</button>
			</div>

			<div className="flex w-full flex-row border-b-[1px] border-stone-950 text-center">
				<button
					onClick={() =>
						handleDomain({ ...domain, testnets: !domain.testnets })
					}
					className="text-md group pointer-events-auto flex h-full w-full items-center justify-center justify-center border-r-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
				>
					<ExclamationTriangleIcon
						className="mr-2 opacity-60"
						width={16}
						height={16}
					/>
					{domain.testnets ? "Hide" : "Show"} Testnets
				</button>
			</div>

			<div className="flex w-full flex-col">
				<div className="mb-auto">
					{accessible.map(({ id, name }) => {
						const active = chainId === Number(id)

						return (
							<button
								key={id}
								onClick={() =>
									switchChain({ chainId: Number(id) })
								}
								className={`text-md group pointer-events-auto mt-auto flex h-full h-min w-full items-center border-b-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950 ${
									active ? "active" : ""
								}`}
							>
								<Image
									src={chainImage(id)}
									alt="Ethereum"
									className="mr-2 h-4 w-4 rounded-full"
									width={16}
									height={16}
								/>
								<span className="opacity-60">
									{formatName(name)}
								</span>
							</button>
						)
					})}
				</div>
			</div>
		</div>
	)
}

export default Wallet
