import { FC, HTMLAttributes, useState } from "react"
import { Button } from "@/components/shared/buttons/button"
import { api } from "@/server/client"
import { Square, Loader, PlayIcon } from "lucide-react"
import { useSetAtom } from "jotai"
import { columnAtomFamily, columnsStorageAtom } from "@/state/column-atoms"
import { DEFAULT_COLUMNS } from "@/state/columns"
import { useResponse } from "@/lib/hooks/useResponse"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const setColumns = useSetAtom(columnsStorageAtom)

	const [killed, setKilled] = useState(false)

	const { isLoading } = useResponse(() => api.solver.killer.killed.useQuery(undefined), {
		onSuccess: data => setKilled(data.killed)
	})

	const toggleSolverMutation = api.solver.killer.kill.useMutation({
		onMutate: () => setKilled(!killed),
		onSuccess: data => setKilled(data.killed)
	})

	return (
		<div {...props} className="p-4 flex items-center flex-col gap-2">
			<Button
				variant={isLoading ? "primaryDisabled" : "primary"}
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

			<Button
				variant={isLoading ? "primaryDisabled" : "primary"}
				onClick={() => setColumns(DEFAULT_COLUMNS)}
				className="w-full py-4 justify-center flex flex-row items-center gap-2"
			>
				<Square className="h-4 w-4 fill-current opacity-60" />
				Reset Columns
			</Button>
		</div>
	)
}
