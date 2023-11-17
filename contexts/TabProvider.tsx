"use client"

import { usePathname, useRouter } from "next/navigation"
import { FC, PropsWithChildren, createContext, useContext, useEffect, useState } from "react"

type Tab = { 
  label: string
  color: `#${string}`
  href: string,
  active?: boolean
}

export const TabsContext = createContext<{
  tabs: Tab[]
  handleAdd: () => void
  handleRemove: (index: number) => void
  handleMove: (index: number, newIndex: number) => void
}>({ tabs: [], handleAdd: () => {}, handleRemove: () => {}, handleMove: () => {} })

export const TabsProvider: FC<PropsWithChildren> = ({ children }) => {
  const router = useRouter()
  const path = usePathname()

  const [tabs, setTabs] = useState<Tab[]>(() => {
    const savedTabs = localStorage.getItem('tabs');

    return savedTabs ? JSON.parse(savedTabs) : [];
  });

  const handleAdd = () => { 
    setTabs(tabs => {
      const tab: Tab = {
        label: `Untitled Canvas`,
        color: `#${Math.floor(Math.random()*16777215).toString(16)}`,
        href: `/canvas/${tabs.length + 1}`,
        active: false
      }

      // * Go to the newly created page.
      router.push(tab.href)

      return [...tabs, tab]
    })
  }

  const handleRemove = (index: number) => {
    setTabs(tabs => {
      const newTabs = [...tabs];

      newTabs.splice(index, 1);

      // * When removing, we may remove the active tab. In that case, we need to
      //   redirect to the last tab in the list.
      if (newTabs.length > 0) {
        const lastTab = newTabs[newTabs.length - 1];

        router.push(lastTab.href);
      }
      // * If there are no more tabs, we need to redirect to the home page.
      else router.push('/canvas');

      return newTabs;
    })
  }

  const handleMove = (index: number, newIndex: number) => { 
    setTabs(tabs => {
      const newTabs = [...tabs];

      const [removed] = newTabs.splice(index, 1);

      newTabs.splice(newIndex, 0, removed);

      return newTabs;
    })
  }

  useEffect(() => {
    localStorage.setItem('tabs', JSON.stringify(tabs));
  }, [tabs]);

  useEffect(() => { 
    // * Mark the selected tab as active when fit.
    setTabs(tabs => tabs.map(tab => ({ ...tab, active: tab.href === path })))
  }, [path])

  return <TabsContext.Provider value={{ tabs, handleAdd, handleRemove, handleMove }}>
    {children}
  </TabsContext.Provider>
}

export const useTabs = () => useContext(TabsContext)
