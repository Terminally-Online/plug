"use client"

import { FC, HTMLAttributes, useMemo } from "react"

import { motion } from "framer-motion"
import { FileCog } from "lucide-react"

import { ActivityList, Counter, Header, StatCard } from "@/components"
import { cn } from "@/lib/utils"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement>> = ({
	...props
}) => {
	return (
		<div {...props}>
			<Header
				size="md"
				icon={<FileCog size={14} className="opacity-40" />}
				label="Runs"
			/>

			<ActivityList />
		</div>
	)
}
