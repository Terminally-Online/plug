import { createContext, useMemo, useState } from "react";

import { CAMERA_ANGLE, RECT_H, RECT_W } from "../lib/constants";
import { cameraToScreenCoordinates } from "../lib/functions/camera-utils";

const isWindow = typeof window !== "undefined";

export type CanvasContextProps = {
  children: React.ReactNode;
  width: number;
  height: number;
};

export const DEFAULT_CANVAS_STATE = {
  state: {
    shouldRender: true,
    pixelRatio: isWindow ? window.devicePixelRatio || 1 : 1,
    container: {
      width: 0,
      height: 0,
    },
    pointer: {
      x: 0,
      y: 0,
    },
    camera: {
      x: 1.5 * RECT_W,
      y: 1.5 * RECT_H,
      z: 0,
    },
    components: {},
  },
};

export const CanvasContext = createContext(DEFAULT_CANVAS_STATE);

export const CanvasProvider = ({
  children,
  width,
  height,
}: CanvasContextProps) => {
  const [state, setState] = useState({
    ...DEFAULT_CANVAS_STATE.state,
    container: { width, height },
    camera: {
      ...DEFAULT_CANVAS_STATE.state.camera,
      z: width / (2 * Math.tan(CAMERA_ANGLE)),
    },
    aspect: width / height,
  });

  const screen = useMemo(() => {
    const { x, y, z } = state.camera;

    return cameraToScreenCoordinates(x, y, z, CAMERA_ANGLE, state.aspect);
  }, [state]);

  const scale = {
    x: state.container.width / screen.width,
    y: state.container.height / screen.height,
  };

  return (
    <CanvasContext.Provider
      value={{
        state,
      }}
    >
      {children}
    </CanvasContext.Provider>
  );
};
