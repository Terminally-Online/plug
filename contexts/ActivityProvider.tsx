import { ContextType, createContext, FC, PropsWithChildren, useContext, useState } from "react"

import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state"

export const ActivityContext = createContext<{
	activities: RouterOutputs["plugs"]["activity"]["get"]
	isLoading: boolean
}>({ activities: [], isLoading: true })

export const ActivityProvider: FC<PropsWithChildren> = ({ children }) => {
	const { isAnonymous } = useSocket()

	const { isLoading } = api.plugs.activity.get.useQuery(undefined, {
		enabled: isAnonymous === false,
		onSuccess: data => setActivities(data)
	})

	const [activities, setActivities] = useState<ContextType<typeof ActivityContext>["activities"]>([])

	api.plugs.activity.onActivity.useSubscription(undefined, {
		onData: data => {
			if (activities.find(activity => activity.id === data.id)) return
			setActivities(prev => (prev ? [data, ...prev] : [data]))
		}
	})

	return <ActivityContext.Provider value={{ activities, isLoading }}>{children}</ActivityContext.Provider>
}

export const useActivities = () => {
	return useContext(ActivityContext)
}
