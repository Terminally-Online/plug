"use client";

import { FC, useEffect, useMemo } from "react";

import { useDrop } from "react-dnd";

import type { DragItem } from "@/lib/types";
import { ItemTypes } from "@/lib/constants";
import CanvasStore from "@/lib/store";
import { snapToGrid } from "@/lib/functions/snap-to-grid";
import { useTabs } from "@/contexts/TabsProvider";
import { t } from "@/app/api/trpc/client";

import { Position } from "./Position";
import { Drag } from "./Drag";
import { Box } from "./blocks/Box";
import { Markdown } from "./blocks/Markdown";
import Toolbar from "./Toolbar";
import Plug from "./blocks/Plug";

export type CanvasProps = {
  frame: string;
  id: string;
};

export const Canvas: FC<CanvasProps> = ({ id }) => {
  const { handleAdd } = useTabs();

  const [canvas, initialCanvasQuery] = t.canvas.get.useSuspenseQuery(id);

  // * The refetch isn't super ideal, but need to get it working.
  //   I am not sure how to replace it until I can figure out how to generate the id from the frontend, otherwise I am
  //      still not sure how you resolve the duplicates.
  const addComponent = t.canvas.component.add.useMutation({
    onSettled: () => {
      initialCanvasQuery.refetch();
    },
  });

  const moveComponent = t.canvas.component.move.useMutation({
    onSettled: () => {
      initialCanvasQuery.refetch();
    },
  });

  const components = useMemo(() => {
    if (!canvas) return null;

    return canvas.components.reduce(
      (acc, component) => {
        acc[component.id] = component;
        return acc;
      },
      {} as Record<string, (typeof canvas.components)[0]>
    );
  }, [canvas]);

  // TODO: Use this once we store the response in state instead of refetching.
  // t.canvas.onUpdate.useSubscription(undefined, {
  //   onData(updatedCanvas) {
  //     setCanvas(updatedCanvas);
  //   },
  // });

  useEffect(() => {
    if (!canvas) return;

    handleAdd({
      label: canvas.name,
      color: canvas.color,
      href: `/canvas/${canvas.id}`,
      active: true,
    });
  }, [canvas, handleAdd]);

  const [, drop] = useDrop(
    () => ({
      accept: [ItemTypes.Box, ItemTypes.Markdown],
      drop(item: DragItem, monitor) {
        const delta = monitor.getDifferenceFromInitialOffset();

        if (!delta) return;

        let { id, top, left } = item;

        left = Math.round(left + delta.x);
        top = Math.round(top + delta.y);

        if (snapToGrid) [left, top] = snapToGrid(left, top);

        moveComponent.mutate({ id: canvas.id, component: { id, left, top } });
      },
    }),
    [moveComponent]
  );

  const handleClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const isShiftClick = e.shiftKey;
    const isCtrlClick = e.ctrlKey;
    const isCommandClick = e.metaKey;
    const isAltClick = e.altKey;

    const ready = isShiftClick || isCommandClick || isCtrlClick || isAltClick;

    if (!ready) return;

    const width = 400;
    const left = CanvasStore.pointer.x - width / 2;
    const top = CanvasStore.pointer.y - width / 4;
    const type = ItemTypes.Plug;

    addComponent.mutate({
      id,
      component: {
        left,
        top,
        type,
        width,
        height: width,
        content: "",
      },
    });
  };

  if (!canvas) return null;

  return (
    <>
      <div
        ref={drop}
        className="relative w-screen h-screen overscroll-none"
        style={{
          transform: `scale(${(CanvasStore.scale.x, CanvasStore.scale.y)})`,
          transformOrigin: "top left",
        }}
        onClick={handleClick}
      >
        {components &&
          Object.keys(components).map((key) => {
            const componentTypes = {
              [ItemTypes.Box]: Box,
              [ItemTypes.Markdown]: Markdown,
              [ItemTypes.Plug]: Plug,
            };

            const component = components[key];
            const Component = componentTypes[component.type];

            console.log("component", component.content);

            return (
              <Position key={component.id} {...component}>
                <Component id={component.id}>
                  {JSON.stringify(component.content)}
                </Component>
              </Position>
            );
          })}

        <Drag />
      </div>

      <Toolbar />
    </>
  );
};

export default Canvas;
