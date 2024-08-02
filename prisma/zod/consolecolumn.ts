import * as z from "zod"
import { CompleteUserSocket, RelatedUserSocketModel } from "./index"

export const ConsoleColumnModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
  key: z.string(),
  index: z.number().int(),
  width: z.number().int().nullish(),
  socketId: z.string(),
})

export interface CompleteConsoleColumn extends z.infer<typeof ConsoleColumnModel> {
  socket: CompleteUserSocket
}

/**
 * RelatedConsoleColumnModel contains all relations on your model in addition to the scalars
 *
 * NOTE: Lazy required in case of potential circular dependencies within schema
 */
export const RelatedConsoleColumnModel: z.ZodSchema<CompleteConsoleColumn> = z.lazy(() => ConsoleColumnModel.extend({
  socket: RelatedUserSocketModel,
}))
