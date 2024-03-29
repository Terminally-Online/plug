import { ReactElement, ReactNode } from "react"

import { NextPage } from "next"

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
	getLayout?: (page: ReactElement) => ReactNode
}

export type { DragItem, ComponentMap, Pins, Pin } from "@/lib/types/canvas"
