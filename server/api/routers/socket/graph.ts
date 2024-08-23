import { createTRPCRouter, protectedProcedure } from "../../trpc";
import { getFarcasterFollowing } from "@/lib";

export const graph = createTRPCRouter({ 
    // farcasterFollowing: protectedProcedure.query(async ({ ctx }) => { 
    //     const following = await getFarcasterFollowing(ctx.session.address);
    //     console.log(following)
    // })
})