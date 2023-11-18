import { redirect } from 'next/navigation';

import { getServerSession } from 'next-auth';
import { authOptions } from '../api/auth/[...nextauth]/route';

export default async function Page() {
  const session = await getServerSession(authOptions)

  const username = session?.user?.name

  if(!username) redirect(`/connect`)

  const canvases = [{ 
    label: "Hello world", 
    color: "#ff00ff", 
    href: "/canvas/hello-world", 
    active: false 
  }]

  return <div className="bg-stone-900 w-screen h-screen flex flex-col items-center justify-center gap-4">
      <h1 className="text-2xl text-white">Create a Canvas</h1>
  </div>
}

