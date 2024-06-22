import { FC } from "react"

import { ActivityItem } from "./activity-item"

export const ActivityList: FC = () => {
	const activities = [
		{
			text: "Bid on Noun with Yellow Glasses",
			color: "green",
			status: "success",
			time: "2s ago"
		},
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
			text: "Bid on Noun with Yellow Glasses",
			color: "yellow",
			status: "success",
			time: "10m ago"
		},
		{
			text: "Bid on Noun with Yellow Glasses",

			color: "yellow",
			status: "warning",
			time: "15m ago"
		},
		{
			text: "Exit Aave When Below 10% APY",

			color: "blue",
			status: "success",
			time: "1h ago"
		},
		{
			text: "Borrow on Aave",
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
			text: "Top-Up Gearbox Loan Health Factor",
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

		// Add more activities as needed
	]

	if (activities.length === 0) {
		return (
			<p className="">
				Pending and completed Plug runs in your Socket will appear
				here...
			</p>
		)
	}

	return (
		<>
			{activities.map((activity, index) => (
				<ActivityItem
					key={index}
					text={activity.text}
					// @ts-ignore
					color={activity.color}
					status={activity.status}
					time={activity.time}
				/>
			))}
		</>
	)
}
