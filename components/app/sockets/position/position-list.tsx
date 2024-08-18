import {
	FC,
	HTMLAttributes,
	useCallback,
	useEffect,
	useMemo,
	useState
} from "react"

import Image from "next/image"

import axios from "axios"

import { Accordion, Counter } from "@/components/shared"
import { useBalances, useSockets } from "@/contexts"
import { cn } from "@/lib"

import { SocketPositionItem } from "./position-item"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id, className, ...props }) => {
	const { positions } = useBalances()

	const { defi } = positions || {}

	if (positions === undefined) return null

	return (
		<div
			className={cn("flex min-h-[calc(100vh-200px)]", className)}
			{...props}
		>
			<div className="mx-auto flex w-full flex-col gap-2">
				{Object.keys(positions.defi).map((protocol: string) => (
					<SocketPositionItem
						key={protocol}
						id={id}
						position={positions.defi[protocol]}
					/>
				))}
			</div>
		</div>
	)
}
