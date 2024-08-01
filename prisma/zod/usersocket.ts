import * as z from "zod"

export const _UserSocketModel = z.object({
  createdAt: z.date(),
  updatedAt: z.date(),
  name: z.string(),
  userAddress: z.string(),
  socketAddress: z.string(),
})
