import { ItemTypes } from "./lib/constants";
import { RECT_H, RECT_W } from "./lib/constants";

import { p } from "../lib/prisma";

import Viewport from "./components/Viewport";

const getCanvas = async (id: string) => {
  // * Get or create the test user -- Everyone is on a global user with no auth for now.
  const user = await p.user.upsert({
    where: { id: "tester" },
    update: {},
    create: {
      id: "tester",
    },
  });

  const HELLO_WORLD = {
    box: {
      type: ItemTypes.Box,
      children: `# Hello world`,
      left: RECT_W * 1.5,
      top: RECT_H * 1.5,
      width: 400,
      height: 400,
    },
  };

  const dbHelloWorld = {
      id,
      userId: user.id,
      components: {
        create: { content: HELLO_WORLD },
      },
    }

  // * Get or create the hello world canvas.
  const canvas = await p.canvas.upsert({
    where: { id },
    update: {},
    create: dbHelloWorld, 
    include: {
      components: true,
    },
  });

  return canvas;
};

export default async function Page({ params }: { params: { id: string } }) {
  const canvas = await getCanvas(params.id);

  const components = canvas.components.reduce((acc, curr) => {
    return { ...acc, ...(curr.content as Record<string, any>) };
  }, {});

  return <Viewport components={components} />;
}
