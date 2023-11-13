'use client'

import { memo } from "react";

import CanvasStore from "./CanvasStore";

import { RECT_H, RECT_W } from "../constants";
import { CanvasPosition, Position } from "./Foundation";

interface TextBlockProps extends CanvasPosition {
  text: string;
  color: string;
  left: number;
  top: number;
  width: number;
  height: number;
}

const TextBlock = ({
  text,
  color,
  left,
  top,
  width,
  height
}: TextBlockProps) => {
  return (
    <Position left={left} top={top} width={width} height={height}>
      <div
        className="flex items-center justify-center"
        style={{
          width: `${width}px`,
          height: `${height}px`,
          background: color
        }}
      >
        {text}
      </div>
    </Position>
  );
};

export const InfiniteCanvas = ({}: { frame: string }) => {
  const texts = [
    "Infinite",
    "Canvases",
    "Are",
    "Easy",
    "When",
    "You",
    "Know",
    "The",
    "Fundamentals"
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
    "#c6b9cd"
  ];

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
      {texts.map((text, index) => (
        <TextBlock
          key={index}
          text={text}
          color={colors[index]}
          left={(index % 3) * rectW}
          top={Math.floor(index / 3) * rectH}
          width={rectW}
          height={rectH}
        />
      ))}
    </div>
  );
};

export default memo(InfiniteCanvas);

