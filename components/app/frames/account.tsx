import Image from "next/image"

import { signOut } from "next-auth/react"

import BlockiesSvg from "blockies-react-svg"
import { useDisconnect, useEnsAvatar, useEnsName } from "wagmi"

import { Button, Counter, Frame, StatCard } from "@/components"
import { useFrame, useSockets } from "@/contexts"

import { normalize } from "viem/ens"

export const AccountFrame = () => {
	const { id, isFrame } = useFrame({ id: "global", key: "account" })
	const { address } = useSockets()

	const { data: ensName } = useEnsName({ address: address as `0x${string}` })
	const { data: ensAvatar } = useEnsAvatar({
		name: normalize(ensName ?? "") || undefined
	})

	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut({ callbackUrl: "/" })
		}
	})

	if (!address) return null

	return (
		<Frame
			id={id}
			className="z-[2]"
			icon={
				ensAvatar ? (
					<Image
						src={ensAvatar}
						alt="ENS Avatar"
						width={24}
						height={24}
						className="h-6 w-6 rounded-sm"
					/>
				) : (
					<BlockiesSvg
						className="h-6 w-6 rounded-sm"
						address={address}
					/>
				)
			}
			label="Wallet"
			visible={isFrame}
		>
			<div className="flex flex-col gap-2">
				<div className="mb-4 flex flex-row items-center gap-2">
					<StatCard>
						<span className="mr-auto flex flex-row gap-2 text-2xl font-bold">
							<Counter className="w-max" count={201} />
							<span className="opacity-40">ETH</span>
						</span>
						<p className="font-bold opacity-40">Fees Earned</p>
					</StatCard>

					<StatCard>
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={321}
						/>
						<p className="font-bold opacity-40">Points Earned</p>
					</StatCard>
				</div>

				<p className="flex font-bold">
					<span className="w-full">Plugs Used</span>
					<Counter className="ml-auto opacity-40" count={31} />
				</p>
				<p className="flex font-bold">
					<span className="w-full">Runs</span>
					<Counter className="ml-auto opacity-40" count={412} />
				</p>
				<p className="flex font-bold">
					<span className="w-full">Users Onboarded</span>
					<Counter className="ml-auto opacity-40" count={51} />
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
