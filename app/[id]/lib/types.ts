import { pins, ItemTypes } from "./constants";

export type DragItem = {
  id: string;
  type: string;
  left: number;
  top: number;
};

export type ComponentMap = {
  [key: string]: {
    type: (typeof ItemTypes)[keyof typeof ItemTypes];
    children: React.ReactNode;
    left: number;
    top: number;
    width?: number;
    height?: number;
  };
};

export type Pin = (typeof pins)[number]["pins"][number]
