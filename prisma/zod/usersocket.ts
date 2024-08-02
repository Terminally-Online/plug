import * as z from "zod"
import { CompleteConsoleColumn, RelatedConsoleColumnModel } from "./index"

export const UserSocketModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
  socketAddress: z.string(),
})

export interface CompleteUserSocket extends z.infer<typeof UserSocketModel> {
  columns: CompleteConsoleColumn[]
}

/**
 * RelatedUserSocketModel contains all relations on your model in addition to the scalars
 *
 * NOTE: Lazy required in case of potential circular dependencies within schema
 */
export const RelatedUserSocketModel: z.ZodSchema<CompleteUserSocket> = z.lazy(() => UserSocketModel.extend({
  columns: RelatedConsoleColumnModel.array(),
}))
