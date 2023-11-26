import { PropsWithChildren } from "react";

import type { Metadata } from "next";

import { TabsProvider } from "@/contexts/TabsProvider";

export const metadata: Metadata = {
  title: "Home | Plug",
  description: "Your homebase for all things Plug.",
};

export default function Layout({ children }: PropsWithChildren) {
  return (
    <TabsProvider>
      <h1>Test</h1>
      {children}
    </TabsProvider>
  );
}
