"use client";

import { FC, memo, PointerEvent, useEffect, useRef, WheelEvent } from "react";

import useSize from "@react-hook/size";

import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import CanvasStore from "@/lib/store";
import useRenderLoop from "@/lib/hooks/useRenderLoop";
import { useTabs } from "@/contexts/TabsProvider";
import { getServerClient } from "@/app/api/trpc/client.server";
import { t } from "@/app/api/trpc/client";
import Canvas from "./Canvas";

export type ViewportProps = {
  canvas: Awaited<
    ReturnType<ReturnType<typeof getServerClient>["canvas"]["get"]>
  >;
};

const Viewport: FC<ViewportProps> = ({ canvas }) => {
  t.randomNumber.useSubscription(undefined, {
    onData: (randomNumber) => {
      console.log(randomNumber);
    },
  });

  const { handleAdd } = useTabs();

  const canvasRef = useRef<HTMLDivElement>(null);

  const frame = useRenderLoop(60);
  const [width, height] = useSize(canvasRef);

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
      label: canvas.name,
      color: canvas.color,
      href: `/canvas/${canvas.id}`,
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
        ref={canvasRef}
        onWheel={wheelListener}
        onPointerMove={pointerListener}
      >
        <DndProvider backend={HTML5Backend}>
          <Canvas frame={frame} canvas={canvas}></Canvas>
        </DndProvider>
      </div>
    </div>
  );
};

export default memo(Viewport);
