import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useState
} from "react"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { api } from "@/server/client"

const consoleModel = Prisma.validator<Prisma.ConsoleDefaultArgs>()({
	include: { columns: true }
})
export type ConsoleModel = Prisma.ConsoleGetPayload<typeof consoleModel>

export const ConsoleContext = createContext<{
	console: ConsoleModel | undefined
	handle: {
		add: (data: { key: string; id?: string }) => void
		remove: (id: string) => void
		resize: (data: { id: string; width: number }) => void
		move: (data: { from: number; to: number }) => void
	}
}>({
	console: undefined,
	handle: {
		add: () => {},
		remove: () => {},
		resize: () => {},
		move: () => {}
	}
})

export const ConsoleProvider: FC<PropsWithChildren> = ({ children }) => {
	const { data, refetch } = api.console.get.useQuery()

	const [console, setConsole] = useState<ConsoleModel | undefined>(data)

	const handle = {
		add: api.console.add.useMutation({
			// TODO: Generate and hunt uuids and create it onMutate instead
			//       of waiting on the server.
			onSuccess: data => setConsole(data)
		}),
		remove: api.console.remove.useMutation({
			onMutate: data => {
				const previousConsole = console

				setConsole(
					previousConsole && {
						...previousConsole,
						columns: previousConsole.columns
							.filter(column => column.id !== data)
							.map((column, index) => ({
								...column,
								index
							}))
					}
				)

				return previousConsole
			},
			onError: (_, __, context) => setConsole(context)
		}),
		resize: api.console.resize.useMutation({
			onMutate: data => {
				const previousConsole = console

				setConsole(
					previousConsole && {
						...previousConsole,
						columns: previousConsole.columns.map(column =>
							column.id === data.id
								? { ...column, width: data.width }
								: column
						)
					}
				)

				return previousConsole
			},
			onError: (_, __, context) => setConsole(context)
		}),
		move: api.console.move.useMutation({
			onMutate: data => {
				if (!console) return

				const previousConsole = console

				const columns = console.columns
				const [removed] = columns.splice(data.from, 1)
				columns.splice(data.to, 0, removed)

				setConsole(
					previousConsole && {
						...previousConsole,
						columns: columns.map((column, index) => ({
							...column,
							index
						}))
					}
				)

				return previousConsole
			},
			onError: (_, __, context) => setConsole(context)
		})
	}

	return (
		<ConsoleContext.Provider
			value={{
				console,
				handle: {
					add: data => handle.add.mutate(data),
					remove: data => handle.remove.mutate(data),
					resize: data => handle.resize.mutate(data),
					move: data => handle.move.mutate(data)
				}
			}}
		>
			{children}
		</ConsoleContext.Provider>
	)
}

export const useConsole = () => useContext(ConsoleContext)
