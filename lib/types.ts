import { z } from "zod";

import { pins, ItemTypes } from "./constants";
import { ReactElement, ReactNode } from "react";
import { NextPage } from "next";

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

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

// infer the type and pass it down to schema
export type Pins = Array<{
  label: string;
  pins: Array<{
    label: string;
    value: string;
    type: "if" | "then";
    schema: z.ZodObject<any>;
  }>;
}>;

export type Pin = (typeof pins)[number]["pins"][number];
