import { redirect } from 'next/navigation';

import { getServerSession } from 'next-auth';
import { authOptions } from '../api/auth/[...nextauth]/route';
import CanvasPreviewGrid from './components/Blocks/CanvasPreviewGrid';

export default async function Page() {
  const session = await getServerSession(authOptions)

  const username = session?.user?.name

  if(!username) redirect(`/connect`)

  // TODO: Retrieve these from trpc.
  const canvases = [{ 
    id: 1,
    name: "Untitled Canvas",
    updatedAt: new Date(),
  }, { 
    id: 1,
    name: "Untitled Canvas",
    updatedAt: new Date(),
  }, { 
    id: 1,
    name: "Untitled Canvas",
    updatedAt: new Date(),
  }, { 
    id: 1,
    name: "Untitled Canvas",
    updatedAt: new Date(),
  }, { 
    id: 1,
    name: "Untitled Canvas",
    updatedAt: new Date(),
  }]

  return <CanvasPreviewGrid canvases={canvases} />
}

