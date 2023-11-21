"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  createWSClient,
  httpBatchLink,
  loggerLink,
  splitLink,
  wsLink,
} from "@trpc/client";
import React, { useState } from "react";

import WalletProvider from "@/contexts/WalletProvider";
import { t } from "@/app/api/trpc/client";
import { SessionProvider } from "next-auth/react";

export default function ClientProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [queryClient] = useState(() => new QueryClient({}));
  const [trpcClient] = useState(() =>
    t.createClient({
      links: [
        splitLink({
          condition: (op) => {
            return op.type === "subscription";
          },
          true: wsLink({
            client: createWSClient({
              url: `ws://localhost:3001/api/trpc`,
            }),
          }),
          false: httpBatchLink({
            url: `http://localhost:3000/api/trpc`,
          }),
        }),
        loggerLink({
          enabled: (opts) =>
            (process.env.NODE_ENV === "development" &&
              typeof window !== "undefined") ||
            (opts.direction === "down" && opts.result instanceof Error),
        }),
      ],
    })
  );

  return (
    <SessionProvider>
      <WalletProvider>
        <t.Provider client={trpcClient} queryClient={queryClient}>
          <QueryClientProvider client={queryClient}>
            {children}
          </QueryClientProvider>
        </t.Provider>
      </WalletProvider>
    </SessionProvider>
  );
}
