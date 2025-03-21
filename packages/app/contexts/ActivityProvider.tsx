import { ContextType, createContext, FC, PropsWithChildren, useContext, useState } from "react"

import { useWatchPendingTransactions } from "wagmi"

import { useConnect } from "@/lib"
import { useResponse } from "@/lib/hooks/useResponse"
import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"

export const ActivityContext = createContext<{
	activities: RouterOutputs["plugs"]["activity"]["get"]
	isLoading: boolean
	handle: { toggle: (data: { id: string }) => void; delete: (data: { id: string }) => void }
}>({ activities: [], isLoading: true, handle: { toggle: () => {}, delete: () => {} } })

export const ActivityProvider: FC<PropsWithChildren> = ({ children }) => {
	const {
		account: { session }
	} = useConnect()
	const { isAnonymous } = useSocket()

	const [activities, setActivities] = useState<ContextType<typeof ActivityContext>["activities"]>([])

	const { isLoading } = useResponse(
		() =>
			api.plugs.activity.get.useQuery(undefined, {
				enabled: session !== null && isAnonymous === false
			}),
		{ onSuccess: data => setActivities(data) }
	)

	api.plugs.activity.onActivity.useSubscription(undefined, {
		onData: data => {
			if (activities.find(activity => activity.id === data.id)) {
				setActivities(prev => prev.map(activity => (activity.id === data.id ? data : activity)))
			} else {
				setActivities(prev => (prev ? [data, ...prev] : [data]))
			}
		}
	})
	api.plugs.activity.onDelete.useSubscription(undefined, {
		onData: data => {
			setActivities(prev => prev.filter(activity => activity.id !== data.id))
		}
	})

	const handle = {
		toggle: api.plugs.activity.toggleStatus.useMutation({
			onMutate: data => {
				setActivities(prev =>
					prev.map(activity =>
						activity.id === data.id
							? { ...activity, status: activity.status.trim() !== "active" ? "active" : "paused" }
							: activity
					)
				)
			}
		}),
		delete: api.plugs.activity.delete.useMutation({
			onMutate: data => {
				setActivities(prev => prev.filter(activity => activity.id !== data.id))
			}
		})
	}

	useWatchPendingTransactions({
		onTransactions: data => {
			console.log(data)
		},
		poll: true,
		pollingInterval: 1_000,
		syncConnectedChain: true
	})

	return (
		<ActivityContext.Provider
			value={{
				activities,
				isLoading,
				handle: {
					toggle: data => handle.toggle.mutate(data),
					delete: data => handle.delete.mutate(data)
				}
			}}
		>
			{children}
		</ActivityContext.Provider>
	)
}

export const useActivities = () => {
	return useContext(ActivityContext)
}
