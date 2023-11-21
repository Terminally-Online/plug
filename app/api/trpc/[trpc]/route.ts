import { fetchRequestHandler } from "@trpc/server/adapters/fetch";

import { AppRouter, appRouter } from "@/server/routers/app";
import { getSession } from "next-auth/react";

const handler = (req: Request) =>
  fetchRequestHandler<AppRouter>({
    endpoint: "/api/trpc",
    req,
    router: appRouter,
    // TODO: Need to fix this to automatically consume the context from the server.
    createContext: async (opts) => {
      const session = await getSession({ req: opts.req });
      console.log("route session", session);
      return { session };
    },
    onError({ error }) {
      if (error.code === "INTERNAL_SERVER_ERROR") {
        // send to bug reporting
        console.error("Something went wrong", error);
      }
    },
    batching: {
      enabled: true,
    },
  });

export { handler as GET, handler as POST };
