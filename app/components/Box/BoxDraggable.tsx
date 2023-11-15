import { memo, useEffect } from "react";
import type { CSSProperties, FC } from "react";

import { useDrag } from "react-dnd";
import type { DragSourceMonitor } from "react-dnd";

import { getEmptyImage } from "react-dnd-html5-backend";

import { Box } from "./Box";
import { ItemTypes } from "../../lib/constants";
import { Position } from "../Canvas/Position";

function getStyles(isDragging: boolean): CSSProperties {
  return {
    opacity: isDragging ? 0 : 1,
    height: isDragging ? 0 : "",
  };
}

export type DraggableBoxProps = {
  id: string;
  title: string;
  left: number;
  top: number;
};

export const BoxDraggable: FC<DraggableBoxProps> = memo(
  function DraggableBox({ id, title, left, top }) {
    const [{ isDragging }, drag, preview] = useDrag(
      () => ({
        type: ItemTypes.Box,
        item: { id, left, top, title },
        collect: (monitor: DragSourceMonitor) => ({
          isDragging: monitor.isDragging(),
        }),
      }),
      [id, left, top, title]
    );

    useEffect(() => {
      preview(getEmptyImage(), { captureDraggingState: true });
    }, [preview]);

    return (
      <div
        ref={drag}
        style={getStyles(isDragging)}
        role="DraggableBox"
      >
        <Box title={title} />
      </div>
    );
  }
);
