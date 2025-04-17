import { FC, HTMLAttributes, useState } from "react"

import { Loader, PlayIcon, Square } from "lucide-react"

import { useSetAtom } from "jotai"

import { Button } from "@/components/shared/buttons/button"
import { useResponse } from "@/lib/hooks/useResponse"
import { api } from "@/server/client"
import { DEFAULT_COLUMNS } from "@/state/columns"
import { columnsStorageAtom } from "@/state/columns"

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
		<div {...props} className="flex flex-col items-center gap-2 p-4">
			<Button
				variant={isLoading ? "primaryDisabled" : "primary"}
				onClick={() => toggleSolverMutation.mutate()}
				disabled={isLoading}
				className="flex w-full flex-row items-center justify-center gap-2 py-4"
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
				className="flex w-full flex-row items-center justify-center gap-2 py-4"
			>
				<Square className="h-4 w-4 fill-current opacity-60" />
				Reset Columns
			</Button>
		</div>
	)
}
