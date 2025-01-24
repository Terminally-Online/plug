import { FC, HTMLAttributes } from "react"
import { Button } from "@/components/shared/buttons/button"
import { api } from "@/server/client"
import { Square, Loader, PlayIcon } from "lucide-react"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const solverStatus = api.jobs.maintenance.getIsSolverKilled.useQuery(undefined, {
		refetchInterval: 30000,
	})

	const toggleSolverMutation = api.jobs.maintenance.toggleSolverKill.useMutation({
		onSuccess: () => {
			solverStatus.refetch()
		},
		onError: (error) => {
			console.error("Failed to update solver status:", error.message)
		}
	})

	const toggleSolver = () => {
		toggleSolverMutation.mutate()
	}

	const isKilled = solverStatus.data?.killed

	return (
		<div {...props} className="p-4 flex items-center justify-between">
			<span className="text-sm">
				Solver Status: {solverStatus.isLoading ? "Loading..." : isKilled ? "Killed" : "Running"}
			</span>
			<Button 
				variant={isKilled ? "primary" : "destructive"}
				onClick={toggleSolver}
				disabled={toggleSolverMutation.isLoading || solverStatus.isLoading}
				className="h-8 w-8 p-0 flex items-center justify-center"
			>
				{toggleSolverMutation.isLoading || solverStatus.isLoading ? (
					<Loader className="h-4 w-4 animate-spin" />
				) : isKilled ? (
					<PlayIcon className="h-4 w-4 fill-current" />
				) : (
					<Square className="h-4 w-4 fill-current" />
				)}
			</Button>
		</div>
	)
}
