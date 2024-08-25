import { getFarcasterFollowing } from "@/lib"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const graph = createTRPCRouter({
	// farcasterFollowing: protectedProcedure.query(async ({ ctx }) => {
	//     const following = await getFarcasterFollowing(ctx.session.address);
	//     console.log(following)
	// })
})
