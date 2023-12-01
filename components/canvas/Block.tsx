import { type FC } from 'react'

import Link from 'next/link'
import { useRouter } from 'next/router'

import { MagicWandIcon, PlusIcon } from '@radix-ui/react-icons'

import { api } from '@/lib/api'
import { cn } from '@/lib/utils'

export const Block: FC = () => {
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
		<div className={cn('flex w-full')}>
			<button
				className="group flex h-full w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-stone-950 bg-stone-900 p-8 py-12 text-center text-white transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
				onClick={handleCreate}
			>
				<div className="w-min rounded-full border-[1px] border-stone-950 bg-stone-900 p-2 group-hover:bg-white">
					<PlusIcon width={18} height={18} className="opacity-60" />
				</div>

				<h1 className="text-2xl">New Canvas</h1>
				<p className="max-w-[240px] text-sm opacity-60">
					Start from scratch and build out your own approach.
				</p>
			</button>

			<Link
				href="canvas/templates"
				className="group flex h-full w-full flex-col items-center justify-center gap-2 border-[1px] border-r-[0px] border-stone-950 bg-stone-900 p-8 py-12 text-center text-white transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
			>
				<div className="w-min rounded-full border-[1px] border-stone-950 bg-stone-900 p-2 group-hover:bg-white">
					<MagicWandIcon
						width={18}
						height={18}
						className="opacity-60"
					/>
				</div>

				<h1 className="text-2xl">Use Template</h1>
				<p className="max-w-[240px] text-sm opacity-60">
					Build on top a foundation that we have already created.
				</p>
			</Link>
		</div>
	)
}

export default Block
