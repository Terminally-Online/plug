import { authOptions } from "@/app/api/auth/[...nextauth]/route"
import { getServerSession } from "next-auth"
import { useSession } from "next-auth/react"
import { redirect } from "next/navigation"

export default async function Page() { 
  const session = await getServerSession(authOptions)

  const username = session?.user?.name

  if(!username) redirect('/connect')

  return <div className="bg-stone-900 w-screen h-screen flex flex-col gap-2 items-center justify-center">
    <div className="border-[1px] border-stone-950 p-4">
      <h1 className="text-2xl text-white">Create a Canvas</h1>
    </div>

    <hr />

    <input placeholder="Search all Canvases..." />
  </div>
}
