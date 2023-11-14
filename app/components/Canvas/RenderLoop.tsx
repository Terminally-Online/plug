"use client";

import CanvasStore from "../../store";
import { useEffect, useRef, useState } from "react";

class RenderLoop {
  private lastFrameTime: number = 0;
  private lastRequestId: number | null = null;
  constructor(
    private fps: number = 0,
    private draw: () => void
  ) {}

  initialize(fps: number) {
    this.fps = fps;
  }

  start() {
    this.lastFrameTime = performance.now();
    this.loop();
  }

  stop() {
    if (this.lastRequestId) cancelAnimationFrame(this.lastRequestId);
    this.lastRequestId = null;
  }

  private get fpsInterval() {
    return 1000 / this.fps;
  }

  private loop() {
    this.lastRequestId = requestAnimationFrame(() => this.loop());
    const now = performance.now();
    const elapsed = now - this.lastFrameTime;
    if (elapsed > this.fpsInterval) {
      this.lastFrameTime = now - (elapsed % this.fpsInterval);
      this.draw();
    }
  }
}

let renderLoop: RenderLoop;
export function getRenderLoop(fps = 15, draw: () => void) {
  if (!renderLoop) return new RenderLoop(fps, draw);
  else return renderLoop;
}

export const useRenderLoop = (fps: number = 15) => {
  const [frame, setFrame] = useState("0");
  const loop = useRef<RenderLoop>(
    getRenderLoop(fps, () => {
      if (CanvasStore.shouldRender) {
        setFrame(`${performance.now()}`);
        CanvasStore.shouldRender = false;
      }
    })
  );

  useEffect(() => {
    CanvasStore.shouldRender = true;

    loop.current.start();

    return () => {
      loop.current.stop();
    };
  }, []);

  return frame;
};

export default useRenderLoop;
