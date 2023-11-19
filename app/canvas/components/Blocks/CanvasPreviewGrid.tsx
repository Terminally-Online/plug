import type { FC } from "react";
import { memo } from "react";

import { Canvas } from "@prisma/client";

import CanvasPreview from "./CanvasPreview";

export type CanvasPreviewGridProps = {
  canvases: Array<Canvas>;
};

export const CanvasPreviewGrid: FC<CanvasPreviewGridProps> = ({ canvases }) => (
  <div className="w-full h-full bg-stone-900 grid grid-cols-3 grid-rows-4">
    {canvases.map((canvas) => (
      <CanvasPreview key={canvas.id} canvas={canvas} />
    ))}
  </div>
);

export default memo(CanvasPreviewGrid);
