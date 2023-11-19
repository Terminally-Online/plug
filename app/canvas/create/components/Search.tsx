"use client"

import type { FC } from "react";
import { memo, useEffect } from "react";

import { useRouter, useSearchParams } from "next/navigation";

import { useDebounce } from "@/lib/hooks/useDebounce";
import { Input } from "@/components/ui/input";

export type SearchProps = { 
  baseUrl?: string
}

const Search: FC<SearchProps> = ({ baseUrl }) => { 
  const router = useRouter()
  const searchParams = useSearchParams()

  const search = searchParams.get('search') ?? ''

  const { debounce, value, debounced } = useDebounce({ initial: search })

  useEffect(() => { 
    // * This is designed to fire once debounced is ahead of search.
    if(search !== debounced) router.push(`${baseUrl ?? '/canvas/create'}?search=${debounced}`)
  }, [search, debounced])

  return <>
    <Input 
      placeholder="Search all Canvases..." 
      className="w-full text-white"
      value={value}
      onChange={(e) => { 
        debounce(e.target.value)
      }}
    />
  </>
}

export default memo(Search)
