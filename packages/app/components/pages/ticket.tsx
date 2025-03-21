import Image from "next/image"
import { createContext, FC, PropsWithChildren, useContext, useRef, useState } from "react"

import { cn } from "@/lib/utils"

type TicketProps = { color: string }

const MouseEnterContext = createContext<[boolean, React.Dispatch<React.SetStateAction<boolean>>] | undefined>(undefined)
export const useMouseEnter = () => {
	const context = useContext(MouseEnterContext)
	if (context === undefined) {
		throw new Error("useMouseEnter must be used within a MouseEnterProvider")
	}
	return context
}

const TicketContainer: FC<PropsWithChildren> = ({ children }) => {
	const containerRef = useRef<HTMLDivElement>(null)

	const [isMouseEntered, setIsMouseEntered] = useState(false)

	const handleMouseMove = (e: React.MouseEvent<HTMLDivElement>) => {
		if (!containerRef.current) return
		const { left, top, width, height } = containerRef.current.getBoundingClientRect()
		const x = (e.clientX - left - width / 2) / 25
		const y = (e.clientY - top - height / 2) / 25
		containerRef.current.style.transform = `rotateY(${x}deg) rotateX(${y}deg)`
	}

	const handleMouseEnter = (e: React.MouseEvent<HTMLDivElement>) => {
		setIsMouseEntered(true)
		if (!containerRef.current) return
	}

	const handleMouseLeave = (e: React.MouseEvent<HTMLDivElement>) => {
		if (!containerRef.current) return
		setIsMouseEntered(false)
		containerRef.current.style.transform = `rotateY(0deg) rotateX(0deg)`
	}

	return (
		<MouseEnterContext.Provider value={[isMouseEntered, setIsMouseEntered]}>
			<div className="group absolute -bottom-1/2 left-0 right-0 w-full rounded-lg p-8 px-12 transition-all duration-200 hover:bottom-0">
				<div
					className={cn("flex items-center justify-center")}
					style={{
						perspective: "2000px"
					}}
				>
					<div
						ref={containerRef}
						onMouseEnter={handleMouseEnter}
						onMouseMove={handleMouseMove}
						onMouseLeave={handleMouseLeave}
						style={{
							transformStyle: "preserve-3d"
						}}
					>
						{children}
					</div>
				</div>
			</div>
		</MouseEnterContext.Provider>
	)
}

export const Ticket: FC<TicketProps> = ({ color }) => {
	return (
		<div className="-z-1 absolute inset-0 select-none overflow-hidden rounded-b-lg [transform-style:preserve-3d] [&>*]:[transform-style:preserve-3d]">
			<TicketContainer>
				<Image
					className="h-full w-full rounded-lg border-[1px] blur-[80px] filter transition-all duration-200 group-hover:blur-none"
					src={`${process.env.NEXT_PUBLIC_APP_URL}/api/nft/image?color=${color.replace("#", "") || "FDFFF7"}`}
					alt="Plug Founding Ticket"
					width={1000}
					height={1600}
				/>
			</TicketContainer>
		</div>
	)
}
