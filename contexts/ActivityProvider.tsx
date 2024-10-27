import { ContextType, createContext, FC, PropsWithChildren, useContext, useState } from "react"

import { api, RouterOutputs } from "@/server/client"

import { useSocket } from "@/state"

export const ActivityContext = createContext<{
	activities: RouterOutputs["plug"]["action"]["getQueued"]
	isLoading: boolean
}>({ activities: [], isLoading: true })

export const ActivityProvider: FC<PropsWithChildren> = ({ children }) => {
	const { isAnonymous } = useSocket()

	const { isLoading } = api.plug.action.getQueued.useQuery(undefined, {
		enabled: isAnonymous === false,
		onSuccess: data => setActivities(data)
	})

	const [activities, setActivities] = useState<ContextType<typeof ActivityContext>["activities"]>([])

	return <ActivityContext.Provider value={{ activities, isLoading }}>{children}</ActivityContext.Provider>
}

export const useActivities = () => {
	return useContext(ActivityContext)
}
