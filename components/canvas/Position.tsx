"use client";

import { PropsWithChildren, useRef } from "react";

import { DEBUG } from "@/lib/constants";
import CanvasStore from "@/lib/store";
import { inBounds } from "@/lib/functions/math-utils";
import { Draggable } from "./Draggable";

export type CanvasPosition = {
  id: string;
  top: number;
  left: number;
  width?: number;
  height?: number;
};

export const Position = ({
  id,
  left,
  top,
  width,
  height,
  children,
}: PropsWithChildren<CanvasPosition>) => {
  const ref = useRef<HTMLDivElement>(null);

  const screen = CanvasStore.screen;

  width = width ?? 400;
  height = height ?? 400;

  if (
    !inBounds(
      { left, top, width, height },
      {
        left: screen.x,
        top: screen.y,
        width: screen.width,
        height: screen.height,
      }
    )
  )
    return null;

  return (
    <div
      ref={ref}
      className="absolute inline-block"
      style={{
        left: `${left - screen.x}px`,
        top: `${top - screen.y}px`,
        width: `${width}px`,
        height: `${height}px`,
      }}
    >
      <Draggable id={id} role={id} left={left} top={top}>
        {children}
      </Draggable>

      {DEBUG && (
        <div
          className="absolute bg-red-400 p-2 rounded-sm text-xs text-red-700 font-bold tabular-nums"
          style={{
            top: "-60px",
            width: "max-content",
          }}
        >
          <p>
            {Math.round(left - screen.x)} x {Math.round(top - screen.y)} @{" "}
            {width ?? 0} x {height ?? 0}
          </p>
        </div>
      )}
    </div>
  );
};
