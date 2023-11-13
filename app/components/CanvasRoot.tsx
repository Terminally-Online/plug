'use client'

import { PointerEvent, useEffect, useRef, WheelEvent } from "react";

import useSize from "@react-hook/size";

import InfiniteCanvas from "./InfiniteCanvas";
import CanvasStore from "./CanvasStore";
import useRenderLoop from "./RenderLoop";

export default function CanvasRoot() {
  const canvas = useRef<HTMLDivElement>(null);

  const frame = useRenderLoop(60);
  const [width, height] = useSize(canvas);

  const wheelListener = (e: WheelEvent) => {
    e.preventDefault();
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
    <div className="w-full h-full overscroll-none">
      <div
        className="w-full h-full relative overflow-hidden overscroll-none"
        ref={canvas}
        onWheel={wheelListener}
        onPointerMove={pointerListener}
      >
        <InfiniteCanvas frame={frame}></InfiniteCanvas>
      </div>
    </div>
  );
};
