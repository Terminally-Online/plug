"use client"

import { useEffect, useState } from "react"

import { Moon, Sun } from "lucide-react"

export default function Theme() {
	const [theme, setTheme] = useState("dark")

	useEffect(() => {
		if (typeof window !== "undefined" && localStorage.theme)
			setTheme(localStorage.theme)
	}, [])

	useEffect(() => {
		if (theme === "dark") document.documentElement.classList.add("dark")
		else document.documentElement.classList.remove("dark")

		localStorage.theme = theme

		setTheme(theme)
	}, [theme])

	return (
		<div className="ml-auto flex flex-row items-center gap-4">
			{theme === "dark" ? (
				<button onClick={() => setTheme("light")}>
					<Sun
						className="text-stone-950/60 dark:text-white/40"
						size={18}
					/>
				</button>
			) : (
				<button onClick={() => setTheme("dark")}>
					<Moon
						className="text-stone-950/60 dark:text-white/40"
						size={18}
					/>
				</button>
			)}
		</div>
	)
}
