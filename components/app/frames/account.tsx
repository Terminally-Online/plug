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
				<div className="mb-4 flex flex-row items-center gap-2">
					<div className="flex w-full flex-col rounded-lg bg-grayscale-100 p-4 font-bold">
						<span className="text-2xl">0.61 ETH</span>
						<span className="opacity-40">Fees Earned</span>
					</div>

					<div className="flex w-full flex-col rounded-lg bg-grayscale-100 p-4 font-bold">
						<span className="text-2xl">312</span>
						<span className="opacity-40">Points Earned</span>
					</div>
				</div>

				<p className="flex font-bold">
					Plugs Used
					<span className="ml-auto opacity-40">31</span>
				</p>
				<p className="flex font-bold">
					Runs
					<span className="ml-auto opacity-40">412</span>
				</p>
				<p className="flex font-bold">
					Users Onboarded
					<span className="ml-auto opacity-40">51</span>
				</p>
				<p className="flex font-bold">
					Account Age
					<span className="ml-auto opacity-40">18 days</span>
				</p>
			</div>

			<div className="mt-4 flex flex-row gap-2">
				<Button
					variant="destructive"
					className="w-max"
					onClick={() => disconnect()}
				>
					Logout
				</Button>
				<Button
					variant="disabled"
					className="w-full"
					onClick={() => {}}
					disabled
				>
					Claim Earnings
				</Button>
			</div>
		</Frame>
	)
}
