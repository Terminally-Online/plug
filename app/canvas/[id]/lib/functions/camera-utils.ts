export const cameraToScreenCoordinates = (
  x: number,
  y: number,
  z: number,
  cameraAngle: number,
  screenAspect: number
) => {
  const width = 2 * z * Math.tan(cameraAngle);
  const height = width / screenAspect;
  const screenX = x - width / 2;
  const screenY = y - height / 2;
  return { x: screenX, y: screenY, width, height };
};

// After changing scale, we return an x y position so that the relative position from top left remains constant
// This way, after zooming, we make sure to set x and y so that users pointer position remains unchanged
export const scaleWithAnchorPoint = (
  anchorPointX: number,
  anchorPointY: number,
  cameraX1: number,
  cameraY1: number,
  scaleX1: number,
  scaleY1: number,
  scaleX2: number,
  scaleY2: number
) => {
  const cameraX2 =
    (anchorPointX * (scaleX2 - scaleX1) + scaleX1 * cameraX1) / scaleX2;
  const cameraY2 =
    (anchorPointY * (scaleY2 - scaleY1) + scaleY1 * cameraY1) / scaleY2;

  return { x: cameraX2, y: cameraY2 };
};
