'use client'

import { PointerEvent, useEffect, useRef, WheelEvent } from "react";

import useSize from "@react-hook/size";

import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'

import InfiniteCanvas from "./components/InfiniteCanvas";
import CanvasStore from "./components/CanvasStore";
import useRenderLoop from "./components/RenderLoop";

export default function Canvas() {
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
    <DndProvider backend={HTML5Backend}>
      <div className="w-screen h-screen overscroll-none">
        <div
          className="w-full h-full relative overflow-hidden"
          ref={canvas}
          onWheel={wheelListener}
          onPointerMove={pointerListener}
        >
          <InfiniteCanvas frame={frame}></InfiniteCanvas>
        </div>
      </div>
    </DndProvider>
  );
};
