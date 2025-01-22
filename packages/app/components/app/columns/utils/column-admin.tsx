import { FC, HTMLAttributes } from "react"
import { Button } from "@/components/shared/buttons/button"
import { api } from "@/server/client"
import { Square, Play } from "lucide-react"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const solverStatus = api.jobs.maintenance.getSolverStatus.useQuery(undefined, {
		refetchInterval: 30000,
	})

	const toggleSolverMutation = api.jobs.maintenance.toggleSolverPause.useMutation({
		onSuccess: () => {
			solverStatus.refetch()
		},
		onError: (error) => {
			console.error("Failed to update solver status:", error.message)
		}
	})

	const toggleSolver = () => {
		toggleSolverMutation.mutate({
			action: solverStatus.data?.status === "stopped" ? "start" : "stop"
		})
	}

	const isStopped = solverStatus.data?.status === "stopped"

	return (
		<div {...props} className="p-4 flex items-center justify-between">
			<span className="text-sm">
				Solver Status: {solverStatus.isLoading ? "Loading..." : isStopped ? "Stopped" : "Running"}
			</span>
			<Button 
				variant={isStopped ? "primary" : "destructive"}
				onClick={toggleSolver}
				disabled={toggleSolverMutation.isLoading || solverStatus.isLoading}
				className="h-8 w-8 p-0"
			>
				{toggleSolverMutation.isLoading ? (
					<span className="animate-spin">⌛</span>
				) : isStopped ? (
					<Play className="h-4 w-4" />
				) : (
					<Square className="h-4 w-4" />
				)}
			</Button>
		</div>
	)
}
