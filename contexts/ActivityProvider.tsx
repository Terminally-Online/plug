import { ContextType, createContext, FC, PropsWithChildren, useContext, useState } from "react"

import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state"

export const ActivityContext = createContext<{
	activities: RouterOutputs["plugs"]["activity"]["get"]
	isLoading: boolean
	handle: { toggle: (data: { id: string }) => void }
}>({ activities: [], isLoading: true, handle: { toggle: () => {} } })

export const ActivityProvider: FC<PropsWithChildren> = ({ children }) => {
	const { isAnonymous } = useSocket()

	const { isLoading } = api.plugs.activity.get.useQuery(undefined, {
		enabled: isAnonymous === false,
		onSuccess: data => setActivities(data)
	})

	const [activities, setActivities] = useState<ContextType<typeof ActivityContext>["activities"]>([])

	api.plugs.activity.onActivity.useSubscription(undefined, {
		onData: data => {
			// NOTE: If the activity item is already in the list, update its state.
			if (activities.find(activity => activity.id === data.id)) {
				setActivities(prev => prev.map(activity => (activity.id === data.id ? data : activity)))
			}

			setActivities(prev => (prev ? [data, ...prev] : [data]))
		}
	})

	const handle = {
		toggle: api.plugs.activity.toggle.useMutation({
			onMutate: () => {
				console.log("toggle state")
			}
		})
	}

	return (
		<ActivityContext.Provider
			value={{ activities, isLoading, handle: { toggle: data => handle.toggle.mutate(data) } }}
		>
			{children}
		</ActivityContext.Provider>
	)
}

export const useActivities = () => {
	return useContext(ActivityContext)
}
