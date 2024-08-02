import * as z from "zod"

export const FeatureRequestModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
  userAddress: z.string(),
  context: z.string(),
  message: z.string().nullish(),
})
