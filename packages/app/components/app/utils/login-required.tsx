import { FC } from "react"

export const LoginRequired: FC = () => {
	return (
		<div className="flex h-full w-full items-center justify-center bg-white">
			<div className="flex w-full max-w-md flex-col items-center gap-2 px-4">
				<h1 className="text-2xl font-bold">Welcome to Plug</h1>
				<p className="max-w-[360px] text-center text-sm font-bold opacity-40">
					Plug is currently in beta so you must log in by connecting your wallet to get started. Anonymous
					accounts are disabled until the beta is over.
				</p>
			</div>
		</div>
	)
}
