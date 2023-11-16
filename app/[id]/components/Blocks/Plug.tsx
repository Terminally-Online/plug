import type { FC, PropsWithChildren } from "react";
import { memo, useState } from "react";

import { z } from "zod";

import PlugSwitcher from "./PlugSwitcher";

export interface PlugProps {
  preview?: boolean;
}

const nounsSchema = z.object({
  price: z.number(),
})

const thresholdSchema = z.object({
  threshold: z.number().default(Date.now()),
})

const pins = [
  {
    label: "Nouns",
    pins: [
      {
        label: "Can Bid on Noun",
        value: "can-bid",
        schema: nounsSchema,
      },
      {
        label: "Place Bid on Noun",
        value: "place-bid",
        schema: nounsSchema,
      }
    ],
  },
  {
    label: "Schedule",
    pins: [
      {
        label: "Within Window",
        value: "within-window",
        schema: thresholdSchema,
      },
      {
        label: 'Before Block Number',
        value: 'before-block-number',
        schema: thresholdSchema,
      },
      {
        label: 'After Block Number',
        value: 'after-block-number',
        schema: thresholdSchema,
      },
      {
        label: 'Before Timestamp',
        value: 'before-timestamp',
        schema: thresholdSchema,
      },
      {
        label: 'After Timestamp',
        value: 'after-timestamp',
        schema: thresholdSchema,
      }
    ],
  }
]

export const Plug: FC<PropsWithChildren<PlugProps>> = memo(function Plug({ children, preview }) {
  const [pins, setPins] = useState<string[]>([1]);

  return (
    <div
      className="bg-stone-900 text-white cursor-move flex flex-col items-center justify-center"
      role={preview ? "PlugPreview" : "Plug"}
    >
      {pins.map(() => <>
        <PlugSwitcher />

        <div className="relative flex flex-col items-center justify-center">
          <div className="h-10 w-[1px] bg-stone-950 rounded-md" />

          <button 
            type="button" 
            className="text-xs absolute border-[1px] border-stone-950 bg-stone-800 text-white/40 hover:bg-stone-900 hover:text-white rounded-full w-[20px] h-[20px] flex items-center justify-center transition-all duration-2oo ease-in-out"
            onClick={() => setPins((pins) => [...pins, pins.length + 1])}
          >
            +
          </button>
        </div>
      </>)}
    </div>
  );
});
