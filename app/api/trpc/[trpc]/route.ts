import { fetchRequestHandler } from "@trpc/server/adapters/fetch";

import { AppRouter, appRouter } from "@/server/routers/app";

const handler = (req: Request) =>
  fetchRequestHandler<AppRouter>({
    endpoint: "/api/trpc",
    req,
    router: appRouter,
    // TODO: Need to fix this to automatically consume the context from the server.
    createContext: () => ({ session: null }),
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
