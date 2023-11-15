"use client";

import { memo, useCallback, useState } from "react";
import update from "immutability-helper";

import { useDrop } from "react-dnd";

import { Drag } from "./Drag";

import { DragItem } from "../types";

import { RECT_H, RECT_W, ItemTypes } from "../lib/constants";
import CanvasStore from "../lib/store";

import { snapToGrid } from "../lib/functions/snap-to-grid";

import { DraggableBox } from "./Box/DraggableBox";
import { MarkdownDraggable } from "./Markdown/MarkdownDraggable";
import { CanvasPosition, Position } from "./Canvas/Position";

interface TextBlockProps extends CanvasPosition {
  text: string;
  color: string;
  width: number;
  height: number;
}

const TextBlock = ({
  text,
  color,
  left,
  top,
  width,
  height,
}: TextBlockProps) => {
  return (
    <Position left={left} top={top} width={width} height={height}>
      <div
        className="flex items-center justify-center"
        style={{
          width: `${width}px`,
          height: `${height}px`,
          background: color,
        }}
      >
        {text}
      </div>
    </Position>
  );
};

export type ComponentMap = {
  [key: string]: {
    type: (typeof ItemTypes)[keyof typeof ItemTypes];
    title: string;
    left: number;
    top: number;
  };
};

const DEBUG = true;

export const Canvas = ({}: { frame: string }) => {
  const [components, setComponents] = useState<ComponentMap>({
    a: {
      type: ItemTypes.Box,
      title: `## ${new Date()} | 6`,
      left: RECT_W / 2,
      top: RECT_H / 2,
    },
  });

  const scale = CanvasStore.scale;

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

  const addComponent = useCallback(
    (id: string, left: number, top: number, type: string) => {
      setComponents((components) => ({
        ...components,
        [id]: {
          type,
          title: `## ${new Date()} | 6`,
          left,
          top,
        },
      }));
    },
    []
  );

  //   const [, drop] = useDrop(
  //     () => ({
  //       accept: [ItemTypes.Box, ItemTypes.Markdown],
  //       drop(item: DragItem, monitor) {
  //         const delta = monitor.getDifferenceFromInitialOffset();

  //         if (!delta) return;

  //         let left = Math.round(item.left + delta.x);
  //         let top = Math.round(item.top + delta.y);

  //         if (snapToGrid) [left, top] = snapToGrid(left, top);

  //         moveComponent(item.id, left, top);
  //       },
  //     }),
  //     [moveComponent]
  //   );

  const texts = [
    "Infinite",
    "Canvases",
    "Are",
    "Easy",
    "When",
    "You",
    "Know",
    "The",
    "Fundamentals",
  ];

  const colors = [
    "#f1f7ed",
    "#61c9a8",
    "#7ca982",
    "#e0eec6",
    "#c2a83e",
    "#ff99c8",
    "#fcf6bd",
    "#9c92a3",
    "#c6b9cd",
  ];

  return (
    <>
      {DEBUG && (
        <div className="fixed top-0 right-0 text-white bg-black/60 p-2 m-2 z-10 rounded-sm">
          <div>Components: {Object.keys(components).length}</div>

          <p>
            Camera: {CanvasStore.camera.x}, {CanvasStore.camera.y},{" "}
            {CanvasStore.camera.z}
          </p>
          <p>
            Scale: {CanvasStore.scale.x}, {CanvasStore.scale.y}
          </p>
          <p>
            Pointer: {CanvasStore.pointer.x}, {CanvasStore.pointer.y}
          </p>

          <button
            type="button"
            onClick={() => {
              console.log("add markdown");
              // todo: add component at pointer

              const id = `box-${Object.keys(components).length + 1}`;
              const left = CanvasStore.pointer.x;
              const top = CanvasStore.pointer.y;
              const type = ItemTypes.Box;

              addComponent(id, left, top, type);
            }}
          >
            Add Box
          </button>
        </div>
      )}

      <div
        // ref={drop}
        className="relative w-full h-full bg-yellow-200"
        style={{
          transform: `scale(${(scale.x, scale.y)})`,
          transformOrigin: "top left",
        }}
      >
        <div
          className="relative w-full h-full"
          style={{
            width: `${RECT_W}px`,
            height: `${RECT_H}px`,
          }}
        >
          {texts.map((text, index) => (
            <TextBlock
              key={index}
              text={text}
              color={colors[index]}
              left={(index % 3) * RECT_W}
              top={Math.floor(index / 3) * RECT_H}
              width={RECT_W}
              height={RECT_H}
            />
          ))}
          {/* {Object.keys(components).map((key) => (
            <DraggableBox key={key} id={key} {...components[key]} />
          ))}

          <Drag /> */}
        </div>
      </div>
    </>
  );
};

export default memo(Canvas);
