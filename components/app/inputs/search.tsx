import { FC, PropsWithChildren, RefObject, useEffect, useRef } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { X } from "lucide-react"

import { cn } from "@/lib/utils"

type Props = {
	ref?: RefObject<HTMLInputElement | HTMLTextAreaElement>
	icon: JSX.Element
	placeholder: string
	search?: string
	clear?: boolean
	textArea?: boolean
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
	textArea = false,
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

		ref.current?.addEventListener("input", handleResize, false)
		ref.current?.addEventListener("change", handleResize)
		window.addEventListener("resize", handleResize)

		return () => {
			ref.current?.removeEventListener("input", handleResize, false)
			ref.current?.removeEventListener("change", handleResize)
			window.removeEventListener("resize", handleResize)
		}
	}, [textArea, ref])

	return (
		<div className={cn("flex flex-col gap-2", className)}>
			<div
				className={cn(
					"flex w-full cursor-pointer items-center gap-4 bg-grayscale-0 p-4 px-6",
					textArea ? "rounded-lg" : "rounded-[16px]"
				)}
				onClick={
					handleOnClick
						? () => handleOnClick()
						: () => ref.current?.focus()
				}
			>
				<div className={cn("w-max", textArea && "mb-auto mt-1")}>
					{icon}
				</div>
				{textArea === false ? (
					<input
						ref={ref as RefObject<HTMLInputElement>}
						type="text"
						placeholder={placeholder}
						className="w-full cursor-pointer bg-transparent outline-none"
						value={search}
						onChange={e =>
							handleSearch ? handleSearch(e.target.value) : null
						}
					/>
				) : (
					<textarea
						ref={ref as RefObject<HTMLTextAreaElement>}
						placeholder={placeholder}
						className="max-h-[40vh] w-full cursor-pointer bg-transparent outline-none"
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
								<X size={14} className="opacity-60" />
							</motion.button>
						)}
					</AnimatePresence>
				)}

				{children && <div className="ml-auto">{children}</div>}
			</div>
		</div>
	)
}
