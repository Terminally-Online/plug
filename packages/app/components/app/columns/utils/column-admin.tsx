import { FC, HTMLAttributes, useState } from "react"
import { Button } from "@/components/shared/buttons/button"
import { api } from "@/server/client"
import { Square, Loader, PlayIcon } from "lucide-react"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const [killed, setKilled] = useState(false)

	const { isLoading } = api.solver.killer.killed.useQuery(undefined, {
		onSuccess: data => setKilled(data.killed)
	})

	const toggleSolverMutation = api.solver.killer.kill.useMutation({
		onMutate: () => setKilled(!killed),
		onSuccess: data => setKilled(data.killed)
	})

	return (
		<div {...props} className="p-4 flex items-center justify-between">
			<Button
				variant={isLoading ? "primaryDisabled" : killed ? "primary" : "destructive"}
				onClick={() => toggleSolverMutation.mutate()}
				disabled={isLoading}
				className="w-full py-4 justify-center flex flex-row items-center gap-2"
			>
				{isLoading ? (
					<Loader className="h-4 w-4 animate-spin opacity-60" />
				) : killed ? (
					<PlayIcon className="h-4 w-4 fill-current opacity-60" />
				) : (
					<Square className="h-4 w-4 fill-current opacity-60" />
				)}

				{killed ? "Revive" : "Kill"} Solver
			</Button>
		</div>
	)
}
