import Link from "next/link"

import { routes } from "@/lib"

import { Blob } from "./blob"
import { LandingContainer } from "./layout"

export const Demo = () => {
	return (
		<div className="relative z-[1] my-[80px]">
			<LandingContainer className="relative mb-[40px] flex flex-col gap-4">
				<div className="flex flex-row items-center gap-12">
					<h1 className="min-w-[640px] text-[64px] font-bold">The Hub For All Your Onchain Activity.</h1>
					<div className="h-[2px] w-full bg-grayscale-100" />
					<Link
						className="whitespace-nowrap font-bold opacity-40 transition-opacity duration-200 ease-in-out hover:opacity-100"
						href={`${routes.documentation}/introduction/integrations`}
						target="_blank"
						rel="noreferrer"
					>
						Experience the Difference
					</Link>
					<div className="h-[2px] w-24 bg-grayscale-100" />
				</div>
				<p className="max-w-[540px] text-[18px] font-bold opacity-40">
					Manage your assets, build strategies with protocol composability, view your activity, and more all
					in one place with the simplest tool to 10x the amount of profit you can make.
				</p>
			</LandingContainer>

			<Blob left={"55%"} top={"300"} width={"1000"} height={"500"} />

			<div className="z-[120] h-[800px] w-full bg-grayscale-0" />
		</div>
	)
}
