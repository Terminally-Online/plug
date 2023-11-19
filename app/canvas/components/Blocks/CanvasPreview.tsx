import type { FC } from "react";
import { memo } from "react";

import Link from "next/link";
import { Canvas } from "@prisma/client";

export type CanvasPreviewProps = {
  canvas: Canvas;
};

export const CanvasPreview: FC<CanvasPreviewProps> = ({ canvas }) => {
  const { id, name, updatedAt } = canvas;

  const href = `/canvas/${id}`;
  const duration = new Date().getTime() - updatedAt.getTime();
  const durationInMinutes = Math.floor(duration / 1000 / 60);

  return (
    <Link
      href={href}
      className="p-4 flex flex-row items-end border-[1px] border-stone-950 text-white hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
    >
      <div className="flex flex-col">
        <h1 className="text-lg font-bold">{name}</h1>
        <p className="text-sm opacity-60">
          Edited {durationInMinutes} minutes ago
        </p>
      </div>
    </Link>
  );
};

export default memo(CanvasPreview);
