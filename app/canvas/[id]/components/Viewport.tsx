"use client";

import { PointerEvent, useEffect, useRef, WheelEvent } from "react";

import useSize from "@react-hook/size";

import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import Canvas from "./Canvas/Canvas";
import CanvasStore from "../lib/store";
import useRenderLoop from "../lib/hooks/useRenderLoop";
import { ComponentMap } from "../lib/types";
import { useTabs } from "@/contexts/TabsProvider";

export default function Viewport({
  id,
  components,
}: {
  id: string;
  components?: ComponentMap;
}) {
  const { handleAdd } = useTabs();

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
    handleAdd({
      label: `Canvas ${id}`,
      color: `#${Math.floor(Math.random() * 16777215).toString(16)}`,
      href: `/canvas/${id}`,
      active: true,
    });
  }, []);

  useEffect(() => {
    if (width === 0 || height === 0) return;

    CanvasStore.initialize(width, height);
  }, [width, height]);

  return (
    <div className="bg-stone-900 w-full h-full text-black dark:text-white">
      <div
        className="w-full h-full relative overflow-hidden overscroll-none"
        ref={canvas}
        onWheel={wheelListener}
        onPointerMove={pointerListener}
      >
        <DndProvider backend={HTML5Backend}>
          <Canvas frame={frame} id={id} components={components}></Canvas>
        </DndProvider>
      </div>
    </div>
  );
}
