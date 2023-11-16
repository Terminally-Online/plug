import type { FC, PropsWithChildren } from "react";
import { memo, useState } from "react";

import { pins } from "../../lib/constants";
import type { Pin as PinType } from "../../lib/types";
import { PinAppendage } from "./PinAppendage";

import Pin from "./Pin";

export type PlugProps = {
  preview?: boolean;
}

export const Plug: FC<PropsWithChildren<PlugProps>> = memo(function Plug({ children, preview }) {
  const [selectedPins, setSelectedPins] = useState([pins[0].pins[0]]);

  const handleChange = (index: number, pin: PinType) => {
    setSelectedPins((previousSelectedPins) => {
      const newSelectedPins = [...previousSelectedPins];

      newSelectedPins[index] = pin;

      // * If this is a 'then', remove all the pins after it
      if(pin.type === "then") {
        newSelectedPins.splice(index + 1, newSelectedPins.length - index);
      }

      return newSelectedPins;
    });
  }

  return (
    <div
      className="bg-stone-900 text-white cursor-move flex flex-col items-center justify-center"
      role={preview ? "PlugPreview" : "Plug"}
    >
      {selectedPins.map((pin, index) => <div key={index} className="w-full h-full">
        <Pin selectedPin={pin} onPinChange={(newPin) => { 
          handleChange(index, newPin);
        }} />

        <PinAppendage pin={pin} onClick={() => { 
          console.log('adding new')
          setSelectedPins((previousSelectedPins) => {
            const newSelectedPins = [...previousSelectedPins];

            newSelectedPins.splice(index + 1, 0, pins[0].pins[0]);

            return newSelectedPins;
          })
        }}/>
      </div>)}
    </div>
  );
});
