import type { FC, PropsWithChildren } from "react";
import { memo, useState } from "react";

import type { Pin as PinType } from "../../lib/types";
import { pins } from "../../lib/constants";

import PinAppendage from "./PinAppendage";
import Pin from "./Pin";

export type PlugProps = {
  preview?: boolean;
}

export const Plug: FC<PropsWithChildren<PlugProps>> = ({ children, preview }) => {
  const [selectedPins, setSelectedPins] = useState([pins[0].pins[0]]);

  const handleChange = (index: number, pin: PinType) => {
    setSelectedPins((previousSelectedPins) => {
      const newSelectedPins = [...previousSelectedPins];
      newSelectedPins[index] = pin;

      // * If we have terminated the chain remove trailing pins.
      if(pin.type === "then") newSelectedPins.splice(index + 1, newSelectedPins.length - index);

      return newSelectedPins;
    });
  }

  const handleAddition = (index: number) => {
    setSelectedPins((previousSelectedPins) => {
      const newSelectedPins = [...previousSelectedPins];
      newSelectedPins.splice(index + 1, 0, pins[0].pins[0]);
      return newSelectedPins;
    })
  }

  return (
    <div
      className="bg-stone-900 text-white cursor-move flex flex-col items-center justify-center"
      role={preview ? "PlugPreview" : "Plug"}
    >
      {selectedPins.map((pin, index) => <div key={index} className="w-full h-full">
        <Pin selectedPin={pin} onPinChange={(newPin) => handleChange(index, newPin)} />
        <PinAppendage pin={pin} onClick={() => handleAddition(index)}/>
      </div>)}
    </div>
  );
};

export default memo(Plug);
