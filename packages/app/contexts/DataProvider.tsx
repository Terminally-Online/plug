import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useSetAtom } from "jotai"

import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { useResponse } from "@/lib/hooks/useResponse"
import { plugsAtom, usePlugSubscriptions } from "@/state/plugs"
import { useConnect } from "@/lib"

type DataContextType = {}
export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ children }) => {
	const { account: { session } } = useConnect()

	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	useResponse(
		() => api.socket.get.useQuery(undefined, {
			enabled: session !== null,
		}),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() => api.solver.actions.schemas.useQuery(
			{ chainId: 8453 },
			{
				enabled: session !== null
			}
		),
		{ onSuccess: actions => setActions(actions) }
	)

	useResponse(() => api.plugs.all.useQuery(
		{ target: "mine" },
		{ enabled: session !== null }
	), {
		onSuccess: data =>
			setPlugs(prev => [...prev, ...data.filter(d => !prev.some(p => p.id === d.id))])
	})

	// TODO: This is going to cause a huge performance degredation if it is kept here.
	// useResponse(() => api.plugs.get.useQuery(
	// 	{ ids, viewed: Array.from(viewedPlugs) },
	// 	{ enabled: Boolean(session.data) && ids.length > 0 }
	// ), {
	// 	onSuccess: data => {
	// 		setPlugs(prev => {
	// 			const uniqueData = data.filter(d => !prev.some(p => p.id === d.id))
	// 			return [...prev, ...uniqueData]
	// 		})
	// 		setViewedPlugs(prev => {
	// 			const newSet = new Set([...Array.from(prev)].slice(-49))
	// 			data.forEach(plug => newSet.add(plug.id))
	// 			if (newSet.size > 50) {
	// 				const entries = Array.from(newSet)
	// 				newSet.clear()
	// 				entries.slice(-50).forEach(id => newSet.add(id))
	// 			}
	// 			return newSet
	// 		})
	// 	}
	// })

	usePlugSubscriptions({ enabled: session !== null })

	return (
		<DataContext.Provider value={{}}>
			{children}
		</DataContext.Provider>
	)
}

export const useData = () => useContext(DataContext)
