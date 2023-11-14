'use client'

import { memo } from "react";

import CanvasStore from "./CanvasStore";
import Markdown from "./Markdown";

import { RECT_H, RECT_W } from "../constants";

export const InfiniteCanvas = ({}: { frame: string }) => {
  // TODO: Retrieve the components from the database.
  
  const components = [{
    type: 'md',
    content: '## test heading',
    left: RECT_W,
    top: RECT_H,
    width: RECT_W,
    height: RECT_H,
  }]

  const rectW = RECT_W;
  const rectH = RECT_H;
  const scale = CanvasStore.scale;

  return (
    <div
      className="w-full h-full"
      style={{
        transform: `scale(${(scale.x, scale.y)})`,
        transformOrigin: "top left"
      }}
    >
      <div
        className="relative"
        style={{
          width: `${rectW}px`,
          height: `${rectH}px`,
        }}
      >
        {components.map((component, index) => {
          switch (component.type) {
            case "md":
              return (
                <Markdown
                  key={index}
                  text={component.content}
                  left={component.left}
                  top={component.top}
                  width={component.width}
                  height={component.height}
                />
              );
            default:
              return null;
          }
        })}
      </div>
    </div>
  );
};

export default memo(InfiniteCanvas);

