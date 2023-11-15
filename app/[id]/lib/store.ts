"use client";

import {
  cameraToScreenCoordinates,
  scaleWithAnchorPoint,
} from "./functions/camera-utils";

import { CAMERA_ANGLE, RECT_H, RECT_W } from "./constants";

const isWindow = typeof window !== "undefined";

export type CanvasState = {
  /// * An infinite canvas means not everything can be rendered at once.
  shouldRender: boolean;
  pixelRatio: number; /// * Resolution for dip calculations
  container: {
    /// * Container dimensions (2d)
    width: number;
    height: number;
  };
  pointer: {
    /// * Pointer coordinates (2d)
    x: number;
    y: number;
  };
  camera: {
    /// * Camera coordinates (3d)
    x: number;
    y: number;
    z: number;
  };
  components: {
    [key: string]: {
      /// * Components to render
      id: string | number;
      top: number;
      left: number;
      height: number;
      width: number;
    };
  };
};

export const DEFAULT_CANVAS_STATE: CanvasState = {
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
};

let canvasData = DEFAULT_CANVAS_STATE;

export default class CanvasStore {
  private static get data() {
    return canvasData;
  }

  static initialize(width: number, height: number) {
    canvasData = {
      ...DEFAULT_CANVAS_STATE,
      container: { width, height },
      camera: {
        ...DEFAULT_CANVAS_STATE.camera,
        z: width / (2 * Math.tan(CAMERA_ANGLE)),
      },
    };
  }

  public static get camera() {
    return this.data.camera;
  }

  public static get screen() {
    const { x, y, z } = this.camera;

    const aspect = this.aspect;
    const angle = CAMERA_ANGLE;

    return cameraToScreenCoordinates(x, y, z, angle, aspect);
  }

  public static get scale() {
    const { width: w, height: h } = CanvasStore.screen;
    const { width: cw, height: ch } = CanvasStore.container;

    return { x: cw / w, y: ch / h };
  }

  public static get shouldRender() {
    if (!canvasData) return false;

    return canvasData.shouldRender;
  }

  public static set shouldRender(value: boolean) {
    if (!canvasData) return;

    canvasData.shouldRender = value;
  }

  public static get container() {
    if (!canvasData) return { width: 0, height: 0 };

    return canvasData.container;
  }

  public static get pointer() {
    if (!canvasData) return { x: 0, y: 0 };

    return canvasData.pointer;
  }

  public static get aspect() {
    if (!canvasData) return 0;

    return canvasData.container.width / canvasData.container.height;
  }

  public static isCameraInBounds(
    cameraX: number,
    cameraY: number,
    cameraZ: number
  ) {
    cameraX;
    cameraY;
    cameraZ;

    return true;

    // const angle = radians(30)
    // const { x, y, width, height } = cameraToScreenCoordinates(
    //   cameraX,
    //   cameraY,
    //   cameraZ,
    //   angle,
    //   this.aspect
    // )
    // const isXInBounds = x >= 0 && x <= this.data.canvas.width
    // const isYInBounds = y >= 0 && y <= this.data.canvas.height
    // return isXInBounds && isYInBounds
  }

  public static moveCamera(mx: number, my: number) {
    const scrollFactor = 1.5;

    const deltaX = mx * scrollFactor;
    const deltaY = my * scrollFactor;

    const { x, y, z } = this.camera;

    if (this.isCameraInBounds(x + deltaX, y + deltaY, z)) {
      this.data.camera.x += deltaX;
      this.data.camera.y += deltaY;

      /// * move pointer by the same amount
      this.shouldRender = true;
      this.movePointer(deltaY, deltaY);
    }
  }

  public static zoomCamera(deltaX: number, deltaY: number) {
    deltaX;

    // Normal zoom is quite slow, we want to scale the amount quite a bit
    const zoomScaleFactor = 10;

    const deltaAmount = zoomScaleFactor * Math.max(deltaY);
    const { x: oldX, y: oldY, z: oldZ } = this.camera;
    const oldScale = { ...this.scale };

    const { width: containerWidth, height: containerHeight } = this.container;
    const { width, height } = cameraToScreenCoordinates(
      oldX,
      oldY,
      oldZ + deltaAmount,
      CAMERA_ANGLE,
      this.aspect
    );

    const newScaleX = containerWidth / width;
    const newScaleY = containerHeight / height;
    const { x: newX, y: newY } = scaleWithAnchorPoint(
      this.pointer.x,
      this.pointer.y,
      oldX,
      oldY,
      oldScale.x,
      oldScale.y,
      newScaleX,
      newScaleY
    );

    const newZ = oldZ + deltaAmount;

    this.shouldRender = true;

    if (this.isCameraInBounds(oldX, oldY, newZ)) {
      this.data.camera = {
        x: newX,
        y: newY,
        z: newZ,
      };
    }
  }

  // pointer position from top left of the screen
  public static movePointer(deltaX: number, deltaY: number) {
    const scale = this.scale;

    const { x: left, y: top } = this.screen;

    this.data.pointer.x = left + deltaX / scale.x;
    this.data.pointer.y = top + deltaY / scale.y;

    console.log(this.data.pointer, left, top);
  }
}
