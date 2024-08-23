import { FC, HTMLAttributes } from "react"

export const ConsoleAlerts: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id }) => {
	return (
		<div className="flex h-full flex-col items-center gap-2">
			<div className="my-auto text-center">
				<p className="font-bold">You are all caught up.</p>
				<p className="max-w-[320px] text-center text-black/60">
					When you have something important to do it will show up
					here.
				</p>
			</div>
		</div>
	)
}
