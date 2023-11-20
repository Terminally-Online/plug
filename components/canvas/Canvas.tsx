"use client";

import { FC, useCallback, useEffect, useState } from "react";

import { useDrop } from "react-dnd";
import update from "immutability-helper";

import Link from "next/link";

import type { DragItem } from "@/lib/types";
import { DEBUG, ItemTypes } from "@/lib/constants";
import CanvasStore from "@/lib/store";
import { snapToGrid } from "@/lib/functions/snap-to-grid";

import { Position } from "./Position";
import { Drag } from "./Drag";
import Toolbar from "./Toolbar";
import { Box } from "./blocks/Box";
import { Markdown } from "./blocks/Markdown";
import Plug from "./blocks/Plug";

import { getServerClient } from "@/app/api/trpc/client.server";

export type CanvasProps = {
  frame: string;
  canvas: Awaited<
    ReturnType<ReturnType<typeof getServerClient>["canvas"]["get"]>
  >;
};

export type Components = Awaited<
  ReturnType<ReturnType<typeof getServerClient>["canvas"]["get"]>
>["components"];
export type Component = Components[0];

export const Canvas: FC<CanvasProps> = ({ frame, canvas }) => {
  const componentsMap = canvas.components.reduce(
    (acc, component) => {
      acc[component.id] = component;

      return acc;
    },
    {} as Record<
      string,
      Pick<
        Component,
        "id" | "type" | "content" | "left" | "top" | "width" | "height"
      >
    >
  );

  const [components, setComponents] = useState(componentsMap);

  const addComponent = useCallback(
    (
      id: string,
      left: number,
      top: number,
      type: string,
      content: React.ReactNode
    ) => {
      setComponents(
        update(components, {
          $merge: {
            [id]: {
              id,
              left,
              top,
              type,
              content,
            },
          },
        })
      );
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

  console.log(components);

  // * When a user double clicks, add a new component to the canvas.
  useEffect(() => {
    const handleDoubleClick = () => {
      const id = `box-${Object.keys(components).length + 1}`;
      const left = CanvasStore.pointer.x;
      const top = CanvasStore.pointer.y;
      const type = ItemTypes.Plug;

      addComponent(id, left, top, type, JSON.stringify({ name: "Test" }));
    };

    document.addEventListener("dblclick", handleDoubleClick);

    return () => document.removeEventListener("dblclick", handleDoubleClick);
  }, []);

  return (
    <>
      {DEBUG && (
        <div className="fixed top-8 right-0 text-red-700 bg-red-400 text-red-700 font-bold p-2 m-2 z-10 rounded-sm">
          <div>Components: {Object.keys(components).length}</div>

          <p>User id: tester</p>

          <p>
            Camera: {Math.round(CanvasStore.camera.x)},{" "}
            {Math.round(CanvasStore.camera.y)},{" "}
            {Math.round(CanvasStore.camera.z)}
          </p>
          <p>Locked: {CanvasStore.camera.locked.toString()}</p>
          <p>
            Scale: {Math.round(CanvasStore.scale.x)},{" "}
            {Math.round(CanvasStore.scale.y)}
          </p>
          <p>
            Screen: {Math.round(CanvasStore.screen.x)},{" "}
            {Math.round(CanvasStore.screen.y)}
          </p>
          <p>
            Pointer: {Math.round(CanvasStore.pointer.x)},{" "}
            {Math.round(CanvasStore.pointer.y)}
          </p>

          <div className="flex flex-row space-x-2 mt-4">
            <Link href={`/create`}>
              <button type="button" className="bg-red-700 text-white p-1 px-2">
                New Canvas
              </button>
            </Link>

            <button
              type="button"
              className="bg-red-700 text-white p-1 px-2"
              onClick={() => {
                const id = `box-${Object.keys(components).length + 1}`;
                const left = CanvasStore.pointer.x;
                const top = CanvasStore.pointer.y;
                const type = ItemTypes.Plug;

                // TODO: Right now we are stringifying the data (children) however we are just discarding the reference.
                // NOTES: The database is accepting raw JSON and does not have a declared shape because we do not want to
                //        enforce a specific structure through the schema because if we did, then we would have to
                //        make a database migration / handler update any time we want to add or remove a new pin.
                //        By doing it this, pins can be put behind feature flags when desired as well as we can just smash
                //        shit into the database and won't have to worry about refactoring for the update that enables
                //        users to create their own Pins.
                addComponent(
                  id,
                  left,
                  top,
                  type,
                  JSON.stringify({ name: "Test" })
                );
              }}
            >
              New Plug
            </button>
          </div>
        </div>
      )}

      <div
        ref={drop}
        className="relative w-screen h-screen overscroll-none"
        style={{
          transform: `scale(${(CanvasStore.scale.x, CanvasStore.scale.y)})`,
          transformOrigin: "top left",
        }}
      >
        {Object.keys(components).map((key) => {
          const componentTypes = {
            [ItemTypes.Box]: Box,
            [ItemTypes.Markdown]: Markdown,
            [ItemTypes.Plug]: Plug,
          };

          const component = components[key as keyof typeof components];
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
