import { FC, PropsWithChildren, RefObject, useEffect, useRef, type JSX } from "react";

import { AnimatePresence, motion } from "framer-motion"
import { X } from "lucide-react"

import { cn } from "@/lib/utils"

type Props = {
	ref?: RefObject<HTMLInputElement | HTMLTextAreaElement | null>
	icon: JSX.Element
	placeholder: string
	search?: string
	clear?: boolean
	focus?: boolean
	textArea?: boolean
	isNumber?: boolean
	handleSearch?: (search: string) => void
	handleOnClick?: () => void
} & PropsWithChildren &
	React.HTMLAttributes<HTMLDivElement>

export const Search: FC<Props> = ({
	ref,
	icon,
	placeholder,
	search,
	clear = false,
	focus = false,
	textArea = false,
	isNumber = false,
	handleSearch,
	handleOnClick,
	className,
	children
}) => {
	const searchRef = useRef<HTMLInputElement | HTMLTextAreaElement>(null)
	ref = ref ? ref : searchRef

	const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
		if (handleSearch) handleSearch(event.target.value)
	}

	useEffect(() => {
		const handleResize = () => {
			if (!ref.current) return

			ref.current.style.height = "auto"
			ref.current.style.height = `${ref.current.scrollHeight}px`
		}

		if (textArea === false || !ref.current) return

		const currentRef = ref.current

		currentRef?.addEventListener("input", handleResize, false)
		currentRef?.addEventListener("change", handleResize)
		window.addEventListener("resize", handleResize)

		return () => {
			currentRef?.removeEventListener("input", handleResize, false)
			currentRef?.removeEventListener("change", handleResize)
			window.removeEventListener("resize", handleResize)
		}
	}, [textArea, ref])

	useEffect(() => {
		if (!focus || !ref.current) return

		ref.current.focus({ preventScroll: true })
	}, [focus, ref])

	return (
		<div className={cn("group flex flex-col gap-2", className)}>
			<div
				className={cn(
					"flex w-full cursor-pointer items-center gap-4 border-[1px] border-plug-green/10 p-4 px-6 transition-colors duration-200 ease-in-out",
					textArea ? "rounded-lg" : "rounded-[16px]",
					search && "border-plug-green/10 bg-white"
				)}
				onClick={handleOnClick ? () => handleOnClick() : () => ref.current?.focus()}
			>
				<div className={cn("w-max opacity-40", textArea && "mb-auto mt-1")}>{icon}</div>
				{textArea === false ? (
					<input
						ref={ref as RefObject<HTMLInputElement | null>}
						type={isNumber ? "number" : "text"}
						placeholder={placeholder}
						className="w-full cursor-pointer bg-transparent font-bold outline-none"
						value={search}
						onChange={e => (handleSearch ? handleSearch(e.target.value) : null)}
						autoCorrect="off"
					/>
				) : (
					<textarea
						ref={ref as RefObject<HTMLTextAreaElement | null>}
						placeholder={placeholder}
						className="max-h-[40vh] w-full cursor-pointer bg-transparent font-bold opacity-40 outline-none group-hover:opacity-100"
						value={search}
						onChange={e => (handleSearch ? handleChange(e) : null)}
					/>
				)}

				{clear && (
					<AnimatePresence>
						{handleSearch && search && (
							<motion.button
								onClick={() => handleSearch("")}
								className="group m-0 cursor-pointer p-1 transition duration-200 ease-in-out"
								initial={{ opacity: 0 }}
								animate={{ opacity: 1 }}
								exit={{ opacity: 0 }}
								transition={{
									duration: 0.2,
									ease: "easeInOut"
								}}
							>
								<X size={14} className="opacity-40" />
							</motion.button>
						)}
					</AnimatePresence>
				)}

				{children && <div className="ml-auto">{children}</div>}
			</div>
		</div>
	);
}
