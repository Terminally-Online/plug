"use client"

import { useMemo, useRef, useState } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { ChevronRight, PartyPopper, TestTube2 } from "lucide-react"

import Fireworks from "./fireworks"

export default function Form() {
	const ref = useRef<HTMLInputElement>(null)

	const [form, setForm] = useState("")
	const [submit, setSubmit] = useState<boolean | number>(false)

	const isValid = useMemo(() => {
		if (form.length === 0) return false

		const isEmail = form.includes("@") && form.includes(".")
		const isAddress = form.length === 42 && form.startsWith("0x")
		const isENS = form.length > 0 && form.includes(".eth")

		return isEmail || isAddress || isENS
	}, [form])

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault()
		e.stopPropagation()

		if (form === "" || !isValid) ref.current?.focus()

		if (!isValid || submit) return

		setSubmit(true)

		fetch("/api", {
			method: "POST",
			body: JSON.stringify({ identifier: form })
		})
			.then(res => res.json())
			.then(res => {
				setSubmit(res.id)
			})
	}

	return (
		<>
			{submit !== false && (
				<Fireworks enabled={typeof submit === "number"} />
			)}

			<form
				onSubmit={handleSubmit}
				className="flex h-min flex-col items-center justify-center border-t-[1px] border-stone-950 lg:flex-row"
			>
				{!submit && (
					<AnimatePresence>
						<motion.div
							className="flex h-min min-w-full flex-row items-center lg:min-w-[75%]"
							initial={{ opacity: 0 }}
							animate={{ opacity: 1 }}
							exit={{ opacity: 0 }}
							transition={{ duration: 0.4, ease: "easeInOut" }}
						>
							<motion.div
								className="ml-4 lg:ml-8"
								initial={{ opacity: 0 }}
								animate={{ opacity: [0.4, 0.6, 0.4] }}
								transition={{
									duration: 1.2,
									ease: "easeInOut",
									repeat: Infinity
								}}
							>
								<ChevronRight
									className="text-white"
									size={24}
								/>
							</motion.div>

							<motion.input
								ref={ref}
								className="w-[100%] bg-transparent p-4 pl-0 text-white/60 placeholder-white/60 outline-none backdrop-blur-lg backdrop-filter"
								type="text"
								value={form}
								onChange={e => setForm(e.target.value.trim())}
								placeholder="YOUR ETHEREUM OR EMAIL ADDRESS"
								animate={{
									width: submit ? "0%" : "",
									padding: submit ? "0px" : "1rem 2rem",
									opacity: submit ? 0 : 1,
									color:
										form !== "" && !isValid ? "#ef4444" : ""
								}}
								autoComplete="off"
								spellCheck="false"
								transition={{
									duration: submit ? 0.4 : 0,
									ease: "easeInOut"
								}}
							/>
						</motion.div>
					</AnimatePresence>
				)}

				<motion.div
					className="min-w-full lg:min-w-[25%]"
					animate={{ width: submit ? "100%" : "" }}
					transition={{ duration: 0.4, ease: "easeInOut" }}
				>
					<motion.button
						type="submit"
						className="flex w-full flex-row items-center justify-center gap-4 bg-white p-4 px-8 text-black transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white disabled:cursor-not-allowed disabled:opacity-50"
						disabled={form !== "" && !isValid}
					>
						{submit ? (
							<PartyPopper className="opacity-40" size={18} />
						) : (
							<TestTube2 className="opacity-40" size={18} />
						)}

						{typeof submit === "number"
							? `YOUR SPOT HAS BEEN SAVED AS #${submit}.`
							: submit === true
							  ? "SUBMITTING..."
							  : form === "" || (form !== "" && isValid)
							    ? "REQUEST EARLY ACCESS"
							    : "INVALID ENTRY"}
					</motion.button>
				</motion.div>
			</form>
		</>
	)
}
