import Image from "next/image"

import { signOut } from "next-auth/react"

import BlockiesSvg from "blockies-react-svg"
import { useDisconnect, useEnsAvatar, useEnsName } from "wagmi"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/buttons"
import { useFrame, useSockets } from "@/contexts"
import { formatAddress } from "@/lib/functions"

import { normalize } from "viem/ens"

export const AccountFrame = () => {
	const { frameVisible } = useFrame()
	const { address } = useSockets()

	const { data: ensName } = useEnsName({ address: address as `0x${string}` })
	const { data: ensAvatar } = useEnsAvatar({
		name: normalize(ensName ?? "") || undefined
	})

	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess(data) {
				signOut({ callbackUrl: "/" })
			}
		}
	})

	if (!address) return null

	return (
		<Frame
			className="z-[2]"
			icon={
				ensAvatar ? (
					<Image
						src={ensAvatar}
						alt="ENS Avatar"
						width={24}
						height={24}
						className="rounded-md"
					/>
				) : (
					<BlockiesSvg
						className="h-6 w-6 rounded-md"
						address={address}
					/>
				)
			}
			label="Account"
			visible={frameVisible === "account"}
		>
			<div className="flex flex-col gap-2">
				<p className="flex font-bold">
					Address
					<span className="ml-auto flex flex-row items-center gap-2 text-opacity-40">
						{formatAddress(address)}
					</span>
				</p>

				<p className="flex font-bold">
					ENS
					<span className="ml-auto flex flex-row items-center gap-2 text-opacity-40">
						{ensName}
					</span>
				</p>
			</div>

			<Button
				variant="destructive"
				className="mt-4 w-full"
				onClick={() => disconnect()}
			>
				Log Out
			</Button>
		</Frame>
	)
}
