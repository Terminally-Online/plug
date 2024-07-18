import { FC } from "react"

import Image from "next/image"
import Link from "next/link"

import { Activity, Book } from "lucide-react"

import { LandingContainer } from "@/components"
import { routes } from "@/lib"

// Use the base64 encoded SVG for the Twitter icon as lucide-react has deprecated it.
const twitter =
	"data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMjIgNHMtLjcgMi4xLTIgMy40YzEuNiAxMC05LjQgMTcuMy0xOCAxMS42IDIuMi4xIDQuNC0uNiA2LTJDMyAxNS41LjUgOS42IDMgNWMyLjIgMi42IDUuNiA0LjEgOSA0LS45LTQuMiA0LTYuNiA3LTMuOCAxLjEgMCAzLTEuMiAzLTEuMnoiIC8+Cjwvc3ZnPgo="

export const Navbar: FC = () => (
	<LandingContainer className="items-center gap-8 py-8">
		<Link href={routes.index}>
			<Image src="/black-logo.svg" alt="Logo" width={64} height={32} />
		</Link>
		<a href={routes.documentation} target="_blank" rel="noreferrer">
			<Book
				size={18}
				className="opacity-60 transition-opacity duration-200 hover:opacity-100"
			/>
		</a>
		<a href={routes.status} target="_blank" rel="noreferrer">
			<Activity
				size={18}
				className="opacity-60 transition-opacity duration-200 hover:opacity-100"
			/>
		</a>
		<a
			href={routes.twitter}
			target="_blank"
			rel="noreferrer"
			className="ml-auto"
		>
			<Image
				src={twitter}
				alt="Twitter"
				width={18}
				height={18}
				className="opacity-60 transition-opacity duration-200 hover:opacity-100"
			/>
		</a>
	</LandingContainer>
)

export default Navbar
