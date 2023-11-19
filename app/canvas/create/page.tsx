import { Suspense } from "react"

import { getServerSession } from "next-auth"
import { redirect } from "next/navigation"

import { authOptions } from "@/app/api/auth/[...nextauth]/route"
import { p } from "@/server/prisma"
import Search from "./components/Search"
import Block from "./components/Block"

async function getCanvases({ userId, count = 10, page = 1, search = '' }: {
  userId: string
  count?: number
  page?: number
  search?: string
}) {
  // TODO: Replace this with `t.infinite()` once it is implemented.
  return await p.canvas.findMany({ 
    where: { 
      userId,
      name: { contains: search }
    },
    take: count,
    skip: count * page
  })
}

export default async function Page({ searchParams }: { 
  searchParams: { search?: string }
}) { 
  const session = await getServerSession(authOptions)
  const search = searchParams.search

  const address = session.address

  if(!address) redirect('/connect')

  const canvases = await getCanvases({
    userId: address,
    search
  })

  // * If they are creating one, and do not have any canvases yet, bump them through this flow and take them to the tutorial board.
  if(!search && canvases.length === 0) console.log('TODO: Should have skipped through this flow')

  return <div className="bg-stone-900 w-screen h-screen flex flex-col gap-2">
    {/* TODO: Implement a loading indicator */}
    <Suspense fallback={<div>Loading...</div>}>
      <Block vertical={canvases.length === 0}/>

      {canvases.length > 0 ? <Search /> : <></>}
    </Suspense>
  </div>
}
