import { type FC } from 'react'

import { api } from '@/lib/api'
import { cn } from '@/lib/utils'

import { MagicWandIcon, PlusIcon } from '@radix-ui/react-icons'
import Link from 'next/link'
import { useRouter } from 'next/router'

export type BlockProps = {
	vertical: boolean
}

export const Block: FC<BlockProps> = ({ vertical }) => {
	const router = useRouter()

	const createCanvas = api.canvas.create.useMutation({
		onSuccess: data => {
			router.push(`/canvas/${data.id}`)
		}
	})

	const handleCreate = () => {
		createCanvas.mutate({
			name: 'Untitled Canvas',
			public: false,
			color: `#${Math.floor(Math.random() * 16777215).toString(16)}`
		})
	}

	return (
		<div
			className={cn(
				'w-full flex',
				vertical ? 'h-full flex-col' : 'flex-row'
			)}
		>
			<button
				className="bg-stone-900 h-full w-full group text-white border-[1px] border-r-[0px] border-stone-950 p-8 flex flex-col gap-4 text-center items-center justify-center hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
				onClick={handleCreate}
			>
				<div className="bg-stone-800 border-[1px] border-stone-950 rounded-full w-min p-2 group-hover:bg-white">
					<PlusIcon width={18} height={18} className="opacity-60" />
				</div>

				<h1 className="text-2xl">New Canvas</h1>
				<p className="opacity-60 max-w-[180px]">
					Start from scratch and build out your own approach.
				</p>
			</button>

			<Link
				href="canvas/templates"
				className="h-full w-full group bg-white text-stone-950 border-[1px] border-stone-950 p-8 flex flex-col gap-4 text-center items-center justify-center hover:bg-stone-900 hover:text-white transition-all duration-200 ease-in-out"
			>
				<div className="text-white bg-stone-800 border-[1px] border-stone-950 rounded-full w-min p-2 group-hover:bg-white group-hover:text-stone-950">
					<MagicWandIcon
						width={18}
						height={18}
						className="opacity-60"
					/>
				</div>

				<h1 className="text-2xl">Use Template</h1>
				<p className="opacity-60 max-w-[180px]">
					Build on top a foundation that we have already created.
				</p>
			</Link>
		</div>
	)
}

export default Block
