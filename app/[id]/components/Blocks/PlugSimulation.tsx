import type { FC } from "react";
import { memo } from "react"

import { LightningBoltIcon } from "@radix-ui/react-icons";
import { Pin } from "../../lib/types";

export type PlugSimulationProps = { 
    pins: Array<Pin>
}

export const PlugSimulation: FC<PlugSimulationProps> = ({ pins }) => { 
    const endsWithThen = pins[pins.length - 1].type === "then";

    if(!endsWithThen) return null;

    return <button 
        type="button" 
        className="flex flex-row items-center justify-center gap-2 bg-stone-800 w-full p-2 rounded-sm border-[1px] border-stone-950 text-sm font-bold hover:bg-white hover:text-black transition-all duration-200 ease-in-out"
    >
        <LightningBoltIcon className="w-3 h-3 text-white/60" /> 
        Simulate
    </button> 
}

export default memo(PlugSimulation)
