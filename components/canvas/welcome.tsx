import { useEffect, useState } from "react"

import { HandIcon, XIcon } from "lucide-react"

export const Welcome = () => {
	const [closed, setClosed] = useState(false)

	const handleClose = () => {
		localStorage.setItem("templates-welcome", "true")

		setClosed(true)
	}

	useEffect(() => {
		const hasClosedTemplateWelcome =
			localStorage.getItem("templates-welcome") === "true"

		if (hasClosedTemplateWelcome === false) setClosed(false)
	}, [])

	return (
		<>
			{!closed && (
				<div className="group relative col-span-3 flex flex-col items-center justify-center gap-2 border-[1px] border-l-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white">
					<button
						className="absolute right-4 top-4 cursor-pointer text-white"
						onClick={handleClose}
					>
						<XIcon width={10} height={10} className="opacity-60" />
					</button>

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
			)}
		</>
	)
}

export default Welcome
