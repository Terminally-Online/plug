import { signOut } from "next-auth/react"

import { useDisconnect } from "wagmi"

import BlockiesSvg from "blockies-react-svg"

import { StatCard } from "@/components/app/cards/stat"
import { Frame } from "@/components/app/frames/base"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnStore } from "@/state/columns"

// NOTE: This is only accessible on the mobile view so the index will always be -1.
export const AccountFrame = () => {
	const { socket, avatar } = useSocket()
	const { isFrame } = useColumnStore(COLUMNS.MOBILE_INDEX, "account")

	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut({ callbackUrl: "/" })
		}
	})

	if (!socket) return null

	return (
		<Frame
			index={COLUMNS.MOBILE_INDEX}
			className="z-[2]"
			icon={
				avatar ? (
					<Image src={avatar} alt="ENS Avatar" width={24} height={24} className="h-6 w-6 rounded-sm" />
				) : (
					<BlockiesSvg className="h-6 w-6 rounded-sm" address={socket.id} />
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
						<Counter className="mr-auto w-max text-2xl font-bold" count={321} />
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
				<Button variant="destructive" className="w-max" onClick={() => disconnect()}>
					Logout
				</Button>
				<Button variant="disabled" className="w-full" onClick={() => {}} disabled>
					Claim Earnings
				</Button>
			</div>
		</Frame>
	)
}
