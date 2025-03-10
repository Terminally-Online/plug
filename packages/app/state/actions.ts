import { useAtom } from "jotai"

import { SchemasResponse } from "@/lib"

import { atomWithStorage } from "jotai/utils"

export const actionsAtom = atomWithStorage<SchemasResponse>("plug.actions", {})

export const useActions = () => useAtom(actionsAtom)
