"use client";

import { ReactNode } from "react";

import { SessionProvider } from "next-auth/react";

import WalletProvider from "@/contexts/WalletProvider";

export default function NextAuthProvider({
  children,
}: {
  children: ReactNode;
}) {
  return <SessionProvider session={null}>
    <WalletProvider>
      {children}
    </WalletProvider>
  </SessionProvider>;
}
