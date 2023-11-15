import { p } from '../lib/prisma'

import Viewport from "./components/Viewport";

const getCanvas = async (id: string) => {
  // * Get or create the test user -- Everyone is on a global user with no auth for now.
  const user = await p.user.upsert({
    where: { id: 'tester' },
    update: {},
    create: {
      id: 'tester'
    }
  })


  // * Get or create the canvas.
  const canvas = await p.canvas.upsert({
    where: { id },
    update: {},
    create: {
      id,
      userId: user.id
    },
    include: {
      components: true
    }
  })
    
  return canvas;
}

export default async function Page({ params }: { params: { id: string } }) { 
  const canvas = await getCanvas(params.id);

  console.log(canvas)

  return <Viewport />
}
