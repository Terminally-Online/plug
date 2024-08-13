import { FC } from "react"

import { usePathname } from "next/navigation"

import { motion } from "framer-motion"

import { Workflow } from "@prisma/client"

import { Button, PlugGridItem } from "@/components"
import { usePlugs, useSockets } from "@/contexts"

type Props = React.HTMLAttributes<HTMLDivElement> & {
	id: string
	from: string
	plugs: Array<Workflow> | undefined
	count?: number
	search?: string
	handleReset?: () => void
}

export const PlugGrid: FC<Props> = ({
	id,
	from,
	plugs,
	count,
	search,
	handleReset,
	...props
}) => {
	const pathname = usePathname()

	const { handle } = useSockets()
	const { handle: handlePlugs } = usePlugs(id)

	if (plugs === undefined) return null

	return (
		<div {...props}>
			{plugs && plugs.length > 0 ? (
				<motion.div
					className="grid gap-1"
					style={{
						gridTemplateColumns: `repeat(auto-fit, minmax(160px, 1fr))`
					}}
					initial="hidden"
					animate="visible"
					variants={{
						hidden: { opacity: 0 },
						visible: {
							opacity: 1,
							transition: {
								staggerChildren: 0.05
							}
						}
					}}
				>
					{plugs
						.slice(0, count || plugs.length)
						.map((plug, index) => (
							<motion.div
								key={`${plug.id}-${index}`}
								variants={{
									hidden: { opacity: 0, y: 10 },
									visible: {
										opacity: 1,
										y: 0,
										transition: {
											type: "spring",
											stiffness: 100,
											damping: 10
										}
									}
								}}
							>
								<PlugGridItem id={id} from={from} plug={plug} />
							</motion.div>
						))}
				</motion.div>
			) : search === "" ? (
				<div className="my-64 flex flex-col gap-2 text-center">
					<p className="text-lg font-bold">No Plugs found.</p>
					<p className="mx-auto max-w-[320px] opacity-60">
						Create your first Plug from scratch or discover one of
						the existing curated and community Plugs now.
					</p>

					<div className="mx-auto mt-8 flex flex-row gap-1">
						<Button
							variant="secondary"
							onClick={() =>
								handle.columns.navigate({ id, key: "discover" })
							}
							className="w-max"
						>
							See Templates
						</Button>
						<Button
							className="w-max"
							onClick={() => handlePlugs.plug.add(pathname)}
						>
							Create
						</Button>
					</div>
				</div>
			) : search !== "" && plugs.length === 0 ? (
				<div className="mx-auto my-44 flex h-full max-w-[80%] flex-col gap-2 text-center">
					<p className="text-lg font-bold">No Plugs found.</p>
					<p className="mx-auto max-w-[320px] opacity-60">
						We looked through all of the results but could not find
						any matches. Reset your filter or try a different
						search.
					</p>

					{handleReset && (
						<div className="mx-auto mt-4 flex flex-row gap-1">
							<Button
								className="w-max"
								onClick={() => handleReset()}
							>
								Reset Filters
							</Button>
						</div>
					)}
				</div>
			) : (
				<></>
			)}
		</div>
	)
}
