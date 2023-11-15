import type { FC, PropsWithChildren } from "react";
import { memo } from "react";

export interface PlugProps {
  preview?: boolean;
}

export const Plug: FC<PropsWithChildren<PlugProps>> = memo(function Plug({ children, preview }) {
  return (
    <div
      className="cursor-move border-dashed border-[1px] border-gray-400 p-2"
      role={preview ? "PlugPreview" : "Plug"}
    >
      {children}
    </div>
  );
});
