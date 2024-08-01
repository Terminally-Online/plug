import * as z from "zod"
import { CompleteConsole, ConsoleModel } from "./index"

export const _ConsoleColumnModel = z.object({
  id: z.string(),
  createdAt: z.date(),
  updatedAt: z.date(),
  key: z.string(),
  index: z.number().int(),
  width: z.number().int().nullish(),
  consoleId: z.string(),
})

export interface CompleteConsoleColumn extends z.infer<typeof _ConsoleColumnModel> {
  console: CompleteConsole
}

/**
 * ConsoleColumnModel contains all relations on your model in addition to the scalars
 *
 * NOTE: Lazy required in case of potential circular dependencies within schema
 */
export const ConsoleColumnModel: z.ZodSchema<CompleteConsoleColumn> = z.lazy(() => _ConsoleColumnModel.extend({
  console: ConsoleModel,
}))
