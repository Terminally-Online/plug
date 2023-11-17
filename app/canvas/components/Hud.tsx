'use client'

import { FC, PropsWithChildren, memo, useState } from "react"

import { Cross1Icon, HomeIcon, PlusIcon } from "@radix-ui/react-icons"

import Link from "next/link";
import { useSession } from "next-auth/react";

export const Hud: FC<PropsWithChildren> = ({ children }) => { 
  const { data: session } = useSession()

  const username = session?.user?.name

  const [tabs, setTabs] = useState([{ 
    label: 'Nouns: Hello World',
    color: '#ff0000',
    href: '/canvas/hello-world-nouns'
  }])

  const handleCreate = () => { 
    setTabs(tabs => [...tabs, {
      label: `Canvas ${tabs.length + 1}`,
      color: `#${Math.floor(Math.random()*16777215).toString(16)}`,
      href: `/canvas/${tabs.length + 1}`
    }])
  }

  return <>
    <div className="bg-stone-900 border-b-[1px] border-b-stone-950 fixed top-0 left-0 w-screen z-[99999]">
      <div className="flex flex-row items-center h-8">
        <Link 
          href="/canvas" 
          className="text-white/60 p-2 h-full flex items-center justify-center text-md font-bold pointer-events-auto hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
        >
          <HomeIcon width={16} height={16} />
        </Link>

        {tabs.map(({ label, color, href }) => 
          <div className="group px-4 border-l-[1px] border-l-stone-950 bg-stone-900 hover:bg-stone-950 h-full flex flex-row items-center gap-4">
            <Link 
              href={href} 
              key={href} 
              className="text-white/60 group-hover:text-white text-sm transition-all duration-200 ease-in-out flex flex-row items-center gap-2"
            >
              <div className="w-2 h-2 rounded-full" style={{ backgroundColor: color }} />

              {label}
            </Link>

            <button 
              type="button" 
              className="opacity-0 group-hover:opacity-100 text-white/60 hover:text-white transition-all duration-200 ease-in-out" 
              onClick={() => setTabs(tabs => tabs.filter(tab => tab.href !== href))}>
              <Cross1Icon width={12} height={12} />
            </button>
          </div>
        )}

        <button 
          type="button" 
          className="p-2 h-full flex items-center justify-center border-x-[1px] border-x-stone-950 bg-stone-800 text-white/60 hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out" 
          onClick={handleCreate}>
          <PlusIcon width={16} height={16} />
        </button>

        <div className="ml-auto">
          {!username ? <Link 
            href="/auth/signin" 
            className="text-white p-2 text-md font-bold"
          >Log In</Link> : <p>{username}</p>}
        </div>
      </div>
    </div>

    {children}
  </>
}

export default memo(Hud)
