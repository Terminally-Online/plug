"use client";

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
