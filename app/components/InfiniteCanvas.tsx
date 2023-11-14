'use client'

import { memo, useCallback, useState } from "react";
import update from 'immutability-helper'

import { useDrop } from "react-dnd";

import { Drag } from "./Drag";
import { ItemTypes } from "./ItemTypes";

import { RECT_H, RECT_W } from "../constants";
import { DragItem } from "../types";
import { snapToGrid } from "../snap-to-grid";

import CanvasStore from "./CanvasStore";

import { MarkdownDraggable } from "./Markdown/MarkdownDraggable";

export type ComponentMap = { 
  [key: string]: { 
    type: typeof ItemTypes[keyof typeof ItemTypes],
    title: string, 
    left: number, 
    top: number, 
  } 
}

export const InfiniteCanvas = ({}: { frame: string }) => {
  const [components, setComponents] = useState<ComponentMap>({ 
    default: {
      type: ItemTypes.Markdown,
      title: '## test heading',
      left: RECT_W / 2,
      top: RECT_H / 2,
   }
  });

  const scale = CanvasStore.scale;

  const moveComponent = useCallback((id: string, left: number, top: number) => {
    setComponents(        
      update(components, {
        [id]: {
          $merge: { left, top },
        },
      })
    );
  }, [components]);

  const [, drop] = useDrop(
    () => ({
      accept: ItemTypes.Markdown,
      drop(item: DragItem, monitor) { 
        const delta = monitor.getDifferenceFromInitialOffset()

        if(!delta) return

        let left = Math.round(item.left + delta.x)
        let top = Math.round(item.top + delta.y)

        if(snapToGrid) [left, top] = snapToGrid(left, top)

        moveComponent(item.id, left, top)
      }
    }), [moveComponent]
  )

  return (
    <div
      ref={drop}
      className="w-full h-full"
      style={{
        background: 'red',
        transform: `scale(${(scale.x, scale.y)})`,
        transformOrigin: "top left"
      }}
    >
      {Object.keys(components).map((key) => (
        <MarkdownDraggable 
          key={key} 
          id={key}
          {...components[key]}
        />
      ))}

      <Drag />
    </div>
  );
};

export default memo(InfiniteCanvas);
