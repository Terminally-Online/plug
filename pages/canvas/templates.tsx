import Link from "next/link"

import { HandIcon, PlusIcon } from "lucide-react"

import { Input } from "@/components/ui/input"
import { TabsProvider } from "@/contexts/TabsProvider"
import { NextPageWithLayout } from "@/lib/types"

// TODO: Update the search bar to just work with any url.
// TODO: Implement a select that allows you to choose from a dropdown.
//	- This select should be used for:
//		- Sorting
// TODO: Sort by createdAt, updatedAt, and name.

// TODO: After alpha we will want to add selection of tags.

export const Templates: NextPageWithLayout = () => {
	return (
		<div className="grid grid-cols-12 grid-rows-1">
			<div className="col-span-3 flex flex-col">
				<div className="group col-span-3 flex flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white">
					<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
						<HandIcon
							width={18}
							height={18}
							className="opacity-60"
						/>
					</div>

					<h1 className="max-w-[280px] text-2xl">
						Welcome to the Plug Discovery Hub.
					</h1>
					<p className="max-w-[240px] text-sm opacity-60">
						Here you can find templates to get you started that have
						been made by the Plug community and team.
					</p>
				</div>

				<div className="top-18 sticky col-span-3 h-full border-r-[1px] border-stone-950">
					<Input
						placeholder="SEARCH TEMPLATES"
						className="relative w-full border-b-[1px] border-stone-950 bg-stone-900 py-8 text-white"
					/>

					<div>Sort By</div>
				</div>
			</div>

			<div className="col-span-9">
				<Link
					href="/canvas/create"
					className="group flex w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
				>
					<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
						<PlusIcon
							width={18}
							height={18}
							className="opacity-60"
						/>
					</div>

					<h1 className="max-w-[380px] text-2xl">
						Build your canvas on top of an existing foundation.
					</h1>
					<p className="max-w-[240px] text-sm opacity-60">
						Start from scratch and build out your own approach.
					</p>
				</Link>
				<Link
					href="/canvas/create"
					className="group flex w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
				>
					<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
						<PlusIcon
							width={18}
							height={18}
							className="opacity-60"
						/>
					</div>

					<h1 className="max-w-[380px] text-2xl">
						Build your canvas on top of an existing foundation.
					</h1>
					<p className="max-w-[240px] text-sm opacity-60">
						Start from scratch and build out your own approach.
					</p>
				</Link>
				<Link
					href="/canvas/create"
					className="group flex w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
				>
					<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
						<PlusIcon
							width={18}
							height={18}
							className="opacity-60"
						/>
					</div>

					<h1 className="max-w-[380px] text-2xl">
						Build your canvas on top of an existing foundation.
					</h1>
					<p className="max-w-[240px] text-sm opacity-60">
						Start from scratch and build out your own approach.
					</p>
				</Link>
			</div>
		</div>
	)
}

Templates.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Templates
