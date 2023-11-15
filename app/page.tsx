"use client";

import { PointerEvent, useEffect, useRef, WheelEvent } from "react";

import useSize from "@react-hook/size";

import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import Canvas from "./components/Canvas";
import CanvasStore from "./lib/store";
import useRenderLoop from "./lib/hooks/useRenderLoop";

export default function Page() {
  const canvas = useRef<HTMLDivElement>(null);

  const frame = useRenderLoop(60);
  const [width, height] = useSize(canvas);

  const wheelListener = (e: WheelEvent) => {
    e.stopPropagation();

    const friction = 1;
    const event = e as WheelEvent;
    const deltaX = event.deltaX * friction;
    const deltaY = event.deltaY * friction;

    if (!event.ctrlKey) {
      CanvasStore.moveCamera(deltaX, deltaY);
    } else {
      CanvasStore.zoomCamera(deltaX, deltaY);
    }
  };

  const pointerListener = (event: PointerEvent) => {
    CanvasStore.movePointer(event.clientX, event.clientY);
  };

  useEffect(() => {
    if (width === 0 || height === 0) return;

    CanvasStore.initialize(width, height);
  }, [width, height]);

  return (
    <div className="w-full h-full text-black">
      <div
        className="w-full h-full relative overflow-hidden overscroll-none"
        ref={canvas}
        onWheel={wheelListener}
        onPointerMove={pointerListener}
      >
        <DndProvider backend={HTML5Backend}>
          <Canvas frame={frame}></Canvas>
        </DndProvider>
      </div>
    </div>
  );
}
