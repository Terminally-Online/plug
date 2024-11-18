import { useRouter } from 'next/router'
import { useEffect, useRef } from 'react'
import { useColumnStore, usePlugStore } from '@/state'
import { COLUMNS } from '@/state'

export const useUrlParams = () => {
  const router = useRouter()
  const { columns, handle } = useColumnStore()
  const { plugs } = usePlugStore()
  const hasHandledInitialUrl = useRef(false)

  useEffect(() => {
    const plugId = router.query.plug as string
    
    if (!plugId || !plugs.length || hasHandledInitialUrl.current) return
    
    // Check if plug exists
    const plugExists = plugs.some(p => p.id === plugId)
    
    if (plugExists) {
      handle.add({
        index: columns[columns.length - 1]?.index + 1 || 0,
        key: COLUMNS.KEYS.PLUG,
        item: plugId,
        from: COLUMNS.KEYS.MY_PLUGS
      })

      // Clear the plug param from URL while preserving other params
      const { plug, ...restQuery } = router.query
      router.replace(
        {
          pathname: router.pathname,
          query: restQuery
        },
        undefined,
        { shallow: true }
      )
      
      hasHandledInitialUrl.current = true
    }
  }, [router, router.query, columns, plugs, handle])
}