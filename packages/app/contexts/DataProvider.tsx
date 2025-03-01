import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useAtom, useSetAtom } from "jotai"

import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { useResponse } from "@/lib/hooks/useResponse"
import { plugsAtom, usePlugSubscriptions, viewedPlugsAtom } from "@/state/plugs"
import { useColumnStore } from "@/state/columns"
import { useConnect } from "@/lib"

type DataContextType = {}
export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const { account: { isAuthenticated }} = useConnect()
	const { columns } = useColumnStore()

	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	const [viewedPlugs, setViewedPlugs] = useAtom(viewedPlugsAtom)

	const ids = (columns?.map(column => column?.item).filter(Boolean) as string[]) || []

	useResponse(
		() => api.socket.get.useQuery(undefined, {
			enabled: isAuthenticated,
		}),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() => api.solver.actions.schemas.useQuery(
			{ chainId: 8453 },
			{
				enabled: isAuthenticated,
				refetchInterval: 5 * 60 * 1000
			}
		),
		{ onSuccess: actions => setActions(actions) }
	)

	// useResponse(() => api.plugs.all.useQuery(
	// 	{ target: "mine" },
	// 	{ enabled: isAuthenticated, queryKeyHashFn: () => `${session?.address}-mine` }
	// ), {
	// 	onSuccess: data =>
	// 		setPlugs(prev => [...prev, ...data.filter(d => !prev.some(p => p.id === d.id))])
	// })
	//
	// useResponse(() => api.plugs.get.useQuery(
	// 	{ ids, viewed: Array.from(viewedPlugs) },
	// 	{ enabled: isAuthenticated && ids.length > 0 }
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

	usePlugSubscriptions()

	return (
		<DataContext.Provider value={{}}>
			{children}
		</DataContext.Provider>
	)
}

export const useData = () => useContext(DataContext)
