import type { FC } from "react";
import { memo, useMemo } from "react";

import Link from "next/link";

import { Canvas } from "@prisma/client";

export type CanvasPreviewProps = {
  canvas: Canvas;
};

export const CanvasPreview: FC<CanvasPreviewProps> = ({ canvas }) => {
  const { id, name, color, updatedAt } = canvas;

  const href = `/canvas/${id}`;

  const durationDisplay = useMemo(() => {
    const duration = new Date().getTime() - updatedAt.getTime();

    const seconds = Math.floor(duration / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);

    if (hours > 0) return `${hours} hour${hours > 1 ? "s" : ""}`;
    if (minutes > 0) return `${minutes} minute${minutes > 1 ? "s" : ""}`;
    if (seconds > 0) return `${seconds} second${seconds > 1 ? "s" : ""}`;

    return "just now";
  }, [updatedAt]);

  return (
    <Link
      href={href}
      className="p-4 flex flex-row items-end border-[1px] border-stone-950 text-white hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
    >
      <div className="flex flex-col gap-2">
        <h1 className="text-lg font-bold flex flex-row items-center gap-4">
          <div
            className="w-4 h-4 rounded-full border-[1px] border-stone-950"
            style={{ background: color }}
          />
          {name}
        </h1>
        <p className="text-sm opacity-60">Edited {durationDisplay} ago</p>
      </div>
    </Link>
  );
};

export default memo(CanvasPreview);
