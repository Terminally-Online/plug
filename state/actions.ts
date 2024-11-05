import { useAtom } from "jotai"

import { ActionSchemas } from "@/lib"

import { atomWithStorage } from "jotai/utils"

export const actionsAtom = atomWithStorage<ActionSchemas>("actions", {})

export const useActions = () => useAtom(actionsAtom)
