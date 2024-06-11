import { type FC, type PropsWithChildren, useEffect } from "react"

import { ChevronLeft, X } from "lucide-react"

import { Button } from "@/components/buttons"
import { cn } from "@/lib/utils"

import { Header } from "../header"

type Props = PropsWithChildren & {
	label: string
	visible: boolean
	icon?: JSX.Element
	handleBack?: () => void
	handleVisibleToggle: () => void
	hasOverlay?: boolean
} & React.HTMLAttributes<HTMLDivElement>

export const Frame: FC<Props> = ({
	children,
	className,
	label,
	visible,
	icon,
	handleBack,
	handleVisibleToggle,
	hasOverlay = false
}) => {
	useEffect(() => {
		const handleKeyDown = (event: KeyboardEvent) => {
			if (event.key === "Escape") handleVisibleToggle()
		}

		if (visible) {
			document.addEventListener("keydown", handleKeyDown)
		}

		return () => {
			document.removeEventListener("keydown", handleKeyDown)
		}
	}, [visible, handleVisibleToggle])

	return (
		<>
			{visible ? (
				<>
					<div
						className={cn(
							"fixed bottom-0 left-0 right-0 top-0 z-[1] cursor-pointer",
							(handleBack === undefined || hasOverlay === true) &&
								"bg-gradient-to-b from-black/10 to-black/30"
						)}
						onClick={handleVisibleToggle}
					/>

					<div
						className={cn(
							"fixed bottom-0 left-0 w-full rounded-t-[20px] bg-white px-6 py-8",
							className
						)}
					>
						<div className="flex flex-row items-center gap-2">
							{handleBack && (
								<Button
									variant="secondary"
									onClick={handleBack}
									className="mb-[20px] mr-2 h-min p-1"
								>
									<ChevronLeft
										size={14}
										className="opacity-60"
									/>
								</Button>
							)}

							<Header
								variant="frame"
								size="md"
								icon={icon}
								label={label}
								nextPadded={false}
								nextOnClick={handleVisibleToggle}
								nextLabel={
									<X size={14} className="opacity-60" />
								}
							/>
						</div>

						{children}
					</div>
				</>
			) : null}
		</>
	)
}
