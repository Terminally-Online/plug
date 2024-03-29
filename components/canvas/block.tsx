import type { FC } from "react"

import Link from "next/link"

import { MagicWandIcon, PlusIcon } from "@radix-ui/react-icons"

export const Block: FC = () => (
	<div className="flex w-full border-r-[1px] border-stone-950">
		<Link
			href="/canvas/create"
			className="group flex h-full w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
		>
			<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
				<PlusIcon width={18} height={18} className="opacity-60" />
			</div>

			<h1 className="text-2xl">New Canvas</h1>
			<p className="max-w-[240px] text-sm opacity-60">
				Start from scratch and build out your own approach.
			</p>
		</Link>

		<Link
			href="canvas/templates"
			className="group flex h-full w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
		>
			<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
				<MagicWandIcon width={18} height={18} className="opacity-60" />
			</div>

			<h1 className="text-2xl">Use Template</h1>
			<p className="max-w-[240px] text-sm opacity-60">
				Build on top a foundation that we have already created.
			</p>
		</Link>
	</div>
)

export default Block
