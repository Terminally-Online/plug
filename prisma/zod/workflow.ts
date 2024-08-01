import * as z from "zod"

export const _WorkflowModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
  userAddress: z.string(),
  name: z.string(),
  isCurated: z.boolean(),
  isPrivate: z.boolean(),
  actions: z.string(),
  color: z.string(),
  tags: z.string().array(),
  workflowForkedId: z.string().nullish(),
})
