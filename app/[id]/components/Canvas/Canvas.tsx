"use client";

import { memo, useCallback, useState } from "react";

import { useDrop } from "react-dnd";
import update from "immutability-helper";

import type { DragItem } from "../../lib/types";

import { DEBUG, RECT_H, RECT_W, ItemTypes } from "../../lib/constants";
import CanvasStore from "../../lib/store";
import { snapToGrid } from "../../lib/functions/snap-to-grid";

import { Position } from "./Position";
import { Drag } from "./Drag";
import { Box } from "../Blocks/Box";
import { Markdown } from "../Blocks/Markdown";
import Link from "next/link";

export type ComponentMap = {
  [key: string]: {
    type: (typeof ItemTypes)[keyof typeof ItemTypes];
    children: React.ReactNode;
    left: number;
    top: number;
    width?: number;
    height?: number;
  };
};

export const Canvas = ({}: { frame: string }) => {
  const [components, setComponents] = useState<ComponentMap>({
    box: {
      type: ItemTypes.Box,
      children: `## ${new Date()} | 6`,
      left: RECT_W * 1.5,
      top: RECT_H * 1.5,
      width: 400,
      height: 400,
    },
  });

  const scale = CanvasStore.scale;

  const addComponent = useCallback(
    (
      id: string,
      left: number,
      top: number,
      type: string,
      children: React.ReactNode
    ) => {
      setComponents((components) => ({
        ...components,
        [id]: {
          type,
          children,
          left,
          top,
          width: 400,
          height: 400,
        },
      }));
    },
    []
  );

  const moveComponent = useCallback(
    (id: string, left: number, top: number) => {
      setComponents(
        update(components, {
          [id]: {
            $merge: { left, top },
          },
        })
      );
    },
    [components]
  );

  const [, drop] = useDrop(
    () => ({
      accept: [ItemTypes.Box, ItemTypes.Markdown],
      drop(item: DragItem, monitor) {
        const delta = monitor.getDifferenceFromInitialOffset();

        if (!delta) return;

        let left = Math.round(item.left + delta.x);
        let top = Math.round(item.top + delta.y);

        if (snapToGrid) [left, top] = snapToGrid(left, top);

        moveComponent(item.id, left, top);
      },
    }),
    [moveComponent]
  );

  return (
    <>
      {DEBUG && (
        <div className="fixed top-0 right-0 text-white bg-red-400 text-red-700 font-bold p-2 m-2 z-10 rounded-sm">
          <div>Components: {Object.keys(components).length}</div>

          <p>
            Camera: {CanvasStore.camera.x}, {CanvasStore.camera.y},{" "}
            {CanvasStore.camera.z}
          </p>
          <p>
            Scale: {CanvasStore.scale.x}, {CanvasStore.scale.y}
          </p>
          <p>
            Screen: {CanvasStore.screen.x}, {CanvasStore.screen.y}
          </p>
          <p>
            Pointer: {CanvasStore.pointer.x}, {CanvasStore.pointer.y}
          </p>

          <div className="flex flex-row space-x-2 mt-4">
            <Link href={`/create`}>
              <button
                type="button"
                className="bg-red-700 text-white p-1 px-2"
              >
                New Canvas
              </button>
            </Link>

            <button
              type="button"
              className="bg-red-700 text-white p-1 px-2"
              onClick={() => {
                setComponents({});
              }}
            >
              Clear
            </button>

            <button
              type="button"
              className="bg-red-700 text-white p-1 px-2"
              onClick={() => {
                const id = `box-${Object.keys(components).length + 1}`;
                const left = CanvasStore.pointer.x;
                const top = CanvasStore.pointer.y;
                const type = ItemTypes.Box;

                addComponent(id, left, top, type, `## ${new Date()} | 6`);
              }}
            >
              New Box
            </button>

            <button
              type="button"
              className="bg-red-700 text-white p-1 px-2"
              onClick={() => {
                const id = `box-${Object.keys(components).length + 1}`;
                const left = CanvasStore.pointer.x;
                const top = CanvasStore.pointer.y;
                const type = ItemTypes.Markdown;

                addComponent(id, left, top, type, `## ${new Date()} | 6`);
              }}
            >
              New Markdown
            </button>
          </div>
        </div>
      )}

      <div
        ref={drop}
        className="relative w-screen h-screen"
        style={{
          transform: `scale(${(scale.x, scale.y)})`,
          transformOrigin: "top left",
        }}
      >
        {Object.keys(components).map((key) => {
          const componentTypes = {
            [ItemTypes.Box]: Box,
            [ItemTypes.Markdown]: Markdown,
          };

          const Component = componentTypes[components[key].type];

          return (
            <Position key={key} id={key} {...components[key]}>
              <Component>{components[key].children}</Component>
            </Position>
          );
        })}

        <Drag />
      </div>
    </>
  );
};

export default memo(Canvas);
