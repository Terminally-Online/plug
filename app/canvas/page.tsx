'use client'

import { useSession } from 'next-auth/react'

import Link from "next/link";

const LogInButton = () => <Link href={`/auth/signin`}>
  <button className="bg-black/40 text-white font-bold rounded-sm p-1 px-2">
    Log In 
  </button>
</Link>

const CreateCanvasButton = () => <Link href={`/create`}>
  <button className="bg-black/40 text-white font-bold rounded-sm p-1 px-2">
    Create Canvas
  </button>
</Link>

export default function Page() {
  const { data: session } = useSession()

  const username = session?.user?.name

  return <div className="bg-stone-900 w-screen h-screen flex flex-col items-center justify-center gap-4">
    <h1 className="text-2xl">{username ? 'Create a new Canvas' : 'Log In to Proceed'}</h1>
    {username ? <CreateCanvasButton /> : <LogInButton />}
  </div>
}
