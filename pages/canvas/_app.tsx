import { PropsWithChildren } from "react";

import type { Metadata } from "next";

import { TabsProvider } from "@/contexts/TabsProvider";
import Hud from "@/components/canvas/Hud";

export const metadata: Metadata = {
  title: "Home | Plug",
  description: "Your homebase for all things Plug.",
};

export default function Layout({ children }: PropsWithChildren) {
  return (
    <TabsProvider>
      <Hud>{children}</Hud>
    </TabsProvider>
  );
}
