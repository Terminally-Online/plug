import * as z from "zod"
import { CompleteConsoleColumn, ConsoleColumnModel } from "./index"

export const _ConsoleModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
})

export interface CompleteConsole extends z.infer<typeof _ConsoleModel> {
  columns: CompleteConsoleColumn[]
}

/**
 * ConsoleModel contains all relations on your model in addition to the scalars
 *
 * NOTE: Lazy required in case of potential circular dependencies within schema
 */
export const ConsoleModel: z.ZodSchema<CompleteConsole> = z.lazy(() => _ConsoleModel.extend({
  columns: ConsoleColumnModel.array(),
}))
