import type { FC } from "react";
import { memo } from "react";

export interface PlugProps {
  title: string;
  preview?: boolean;
}

export const Plug: FC<PlugProps> = memo(function Plug({ title, preview }) {
  return (
    <div
      className="cursor-move border-dashed border-[1px] border-gray-400 p-2"
      role={preview ? "PlugPreview" : "Plug"}
    >
      {title}
    </div>
  );
});
