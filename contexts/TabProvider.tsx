"use client"

import { FC, PropsWithChildren, createContext, useContext, useEffect, useState } from "react"

type Tab = { 
  label: string
  color: `#${string}`
  href: string
}

export const TabsContext = createContext<{
  tabs: Tab[]
  handleAdd: () => void
  handleRemove: (index: number) => void
  handleMove: (index: number, newIndex: number) => void
}>({ tabs: [], handleAdd: () => {}, handleRemove: () => {}, handleMove: () => {} })

export const TabsProvider: FC<PropsWithChildren> = ({ children }) => {
  const [tabs, setTabs] = useState<Tab[]>(() => {
    const savedTabs = localStorage.getItem('tabs');

    return savedTabs ? JSON.parse(savedTabs) : [];
  });

  const handleAdd = () => { 
    setTabs(tabs => [...tabs, {
      label: `Canvas ${tabs.length + 1}`,
      color: `#${Math.floor(Math.random()*16777215).toString(16)}`,
      href: `/canvas/${tabs.length + 1}`
    }])
  }

  const handleRemove = (index: number) => {
    setTabs(tabs => {
      const newTabs = [...tabs];

      newTabs.splice(index, 1);

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

  return <TabsContext.Provider value={{ tabs, handleAdd, handleRemove, handleMove }}>
    {children}
  </TabsContext.Provider>
}

export const useTabs = () => useContext(TabsContext)
