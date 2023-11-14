'use client'

import {
  cameraToScreenCoordinates,
  scaleWithAnchorPoint,
} from "../camera-utils";

import { CAMERA_ANGLE, RECT_H, RECT_W } from "../constants";

export interface CanvasState {
  shouldRender: boolean;
  pixelRatio: number; // our resolution for dip calculations
  container: {
    //holds information related to our screen container
    width: number;
    height: number;
  };
  pointer: {
    x: number;
    y: number;
  };
  camera: {
    //holds camera state
    x: number;
    y: number;
    z: number;
  };
}

const getInitialCanvasState = (): CanvasState => {
  return {
    shouldRender: true,
    pixelRatio: window.devicePixelRatio || 1,
    container: {
      width: 0,
      height: 0,
    },
    pointer: {
      x: 0,
      y: 0,
    },
    camera: {
      x: 0,
      y: 0,
      z: 0,
    },
  };
};

let canvasData = typeof window !== 'undefined' ? getInitialCanvasState() : null;

export default class CanvasStore {
  private static get data() {
    const pixelRatio = typeof window !== 'undefined' ? window.devicePixelRatio || 1 : 1; 

    if (!canvasData)
      canvasData = {
        shouldRender: true,
        pixelRatio,
        container: {
          width: 0,
          height: 0,
        },
        pointer: {
          x: 0,
          y: 0,
        },
        camera: {
          x: 0,
          y: 0,
          z: 0,
        },
      };

    return canvasData;
  }

  static initialize(width: number, height: number) {
    const containerWidth = width;
    const containerHeight = height;

    canvasData = getInitialCanvasState();
    canvasData.pixelRatio = window.devicePixelRatio || 1;
    canvasData.container.width = containerWidth;
    canvasData.container.height = containerHeight;
    canvasData.camera.x = 1.5 * RECT_W;
    canvasData.camera.y = 1.5 * RECT_H;
    canvasData.camera.z = containerWidth / (2 * Math.tan(CAMERA_ANGLE));
  }

  public static get screen() {
    const { x, y, z } = this.camera;
    const aspect = this.aspect;
    const angle = CAMERA_ANGLE;
    return cameraToScreenCoordinates(x, y, z, angle, aspect);
  }

  public static get camera() {
    return this.data.camera;
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

  private static get container() {
    if (!canvasData) return { width: 0, height: 0 };

    return canvasData.container;
  }

  private static get pointer() {
    if (!canvasData) return { x: 0, y: 0 };

    return canvasData.pointer;
  }

  private static get aspect() {
    if (!canvasData) return 0;

    return canvasData.container.width / canvasData.container.height;
  }

  private static isCameraInBounds(
    cameraX: number,
    cameraY: number,
    cameraZ: number
  ) {
    cameraX
    cameraY
    cameraZ

    return true;

    // const angle = radians(30);
    // const { x, y, width, height } = cameraToScreenCoordinates(
    //   cameraX,
    //   cameraY,
    //   cameraZ,
    //   angle,
    //   this.aspect
    // );
    // const isXInBounds = x >= 0 && x <= this.data.canvas.width;
    // const isYInBounds = y >= 0 && y <= this.data.canvas.height;
    // return isXInBounds && isYInBounds;
  }

  public static moveCamera(mx: number, my: number) {
    const scrollFactor = 1.5;

    const deltaX = mx * scrollFactor,
      deltaY = my * scrollFactor;

    const { x, y, z } = this.camera;

    if (this.isCameraInBounds(x + deltaX, y + deltaY, z)) {
      this.data.camera.x += deltaX;
      this.data.camera.y += deltaY;

      // move pointer by the same amount
      this.shouldRender = true;
      this.movePointer(deltaY, deltaY);
    }
  }

  public static zoomCamera(deltaX: number, deltaY: number) {
    deltaX

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
  }
}

