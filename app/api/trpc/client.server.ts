import { appRouter } from "@/server/routers/app";

// TODO: The null type here is not right, but we are pretending for now because it works since Typescript is not real type safety.
export const getServerClient = (session: null) =>
    appRouter.createCaller({
        session,
    });
