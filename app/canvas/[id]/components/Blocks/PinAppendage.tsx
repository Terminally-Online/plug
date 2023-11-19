import { FC, memo } from "react";

import type { Pin } from "../../lib/types";

export type PinAppendageProps = {
  pin: Pin;
  onClick: () => void;
  isAvailable?: boolean;
};

export const PinAppendage: FC<PinAppendageProps> = ({
  pin,
  onClick,
  isAvailable = true,
}) => {
  return (
    <div className="relative flex flex-col items-center justify-center">
      <div className="h-10 w-[1px] bg-stone-950 rounded-md" />

      {isAvailable && pin.type === "if" && (
        <button
          type="button"
          className="text-xs absolute border-[1px] border-stone-950 bg-stone-800 text-white/40 hover:bg-stone-900 hover:text-white rounded-full w-[20px] h-[20px] flex items-center justify-center transition-all duration-2oo ease-in-out"
          onClick={onClick}
        >
          +
        </button>
      )}
    </div>
  );
};

export default memo(PinAppendage);
