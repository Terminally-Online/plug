import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useMemo } from "react"

import { ActivityItem, Animate, Callout } from "@/components"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"

import { ActivityFrame } from "../../frames/sockets/activity"

const activities = [
	{
		text: "Lend ETH with High APY",
		color: "blue",
		status: "success",
		time: "12s ago"
	},
	{
		text: "Buy Beta When Majors Move",
		color: "green",
		status: "success",
		time: "3m ago"
	},
	{
		text: "Exit Aave When Below 10% APY",

		color: "blue",
		status: "success",
		time: "1h ago"
	},
	{
		text: "Manage Aave Position",
		color: "blue",
		status: "success",
		time: "2h ago"
	},
	{
		text: "Lend ETH with High APY",
		color: "blue",
		status: "success",
		time: "3h ago"
	},
	{
		text: "Balance Gearbox",
		color: "orange",
		status: "success",
		time: "6h ago"
	},
	{
		text: "Get Tip Allowance & Auto-Compound ENJOY Rewards",
		color: "cyan",
		status: "success",
		time: "9h ago"
	},
	{
		text: "Borrow on Aave",
		color: "blue",
		status: "warning",
		time: "16h ago"
	},
	{
		text: "Lend ETH with High APY",
		color: "blue",
		status: "error",
		time: "1d ago"
	},
	{
		text: "Bridge to Optimism, Base, and Bera",
		color: "red",
		status: "success",
		time: "2d ago"
	},
	{
		text: "Bid on Noun with Yellow Glasses",
		color: "yellow",
		status: "warning",
		time: "2d ago"
	},
	{
		text: "Fill Ethena Liquidity Cap to Limit",
		color: "green",
		status: "success",
		time: "2d ago"
	},
	{
		text: "Renew ENS Annually at Low Gas",
		color: "green",
		status: "success",
		time: "2d ago"
	},
	{
		text: "Enter Yearn If Above 50%",
		color: "blue",
		status: "success",
		time: "2d ago"
	},
	{
		text: "Borrow on Gearbox",
		color: "orange",
		status: "success",
		time: "3d ago"
	},
	{
		text: "Top-Up Gearbox Loan Health Factor",
		color: "orange",
		status: "success",
		time: "3d ago"
	},
	{
		text: "Exit Aave When Below 10% APY",
		color: "blue",
		status: "success",
		time: "3d ago"
	},
	{
		text: "Lend ETH with High APY",
		color: "blue",
		status: "success",
		time: "3d ago"
	},
	{
		text: "Bid on Noun with Yellow Glasses",
		color: "yellow",
		status: "warning",
		time: "3d ago"
	},
	{
		text: "Bid on Noun with Yellow Glasses",
		color: "yellow",
		status: "warning",
		time: "4d ago"
	},
	{
		text: "Borrow on Gearbox",
		color: "orange",
		status: "success",
		time: "4d ago"
	},
	{
		text: "Top-Up Gearbox Loan Health Factor",
		color: "orange",
		status: "success",
		time: "4d ago"
	},
	{
		text: "Bid on Noun with Yellow Glasses",
		color: "yellow",
		status: "warning",
		time: "5d ago"
	},
	{
		text: "Lend ETH with High APY",
		color: "green",
		status: "success",
		time: "5d ago"
	},
	{
		text: "Bid on Noun with Yellow Glasses",
		color: "yellow",
		status: "warning",
		time: "6d ago"
	},
	{
		text: "Get Tip Allowance & Auto-Compound ENJOY Rewards",
		color: "blue",
		status: "success",
		time: "8d ago"
	},
	{
		text: "Top-Up Gearbox Loan Health Factor",
		color: "orange",
		status: "error",
		time: "11d ago"
	}
]

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { data: session } = useSession()
	const { isAnonymous } = useSocket()
	const { isExternal } = useColumns(index)

	const visibleActivities = useMemo(() => {
		if (!session || (isAnonymous && isExternal === false)) return Array(10).fill(undefined)

		return activities
	}, [session, isAnonymous, isExternal])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Animate.List>
				{visibleActivities.map((activity, activityIndex) => (
					<Animate.ListItem key={activityIndex}>
						<ActivityItem index={index} activityIndex={activityIndex} activity={activity} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="activity" isAbsolute={true} />

			{visibleActivities
				.filter(activity => Boolean(activity))
				.map((activity, activityIndex) => (
					<ActivityFrame
						key={activityIndex}
						index={index}
						activityIndex={activityIndex}
						activity={{ name: activity.text, status }}
					/>
				))}
		</div>
	)
}
