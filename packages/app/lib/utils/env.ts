// import dotenv from "dotenv"
// import { z } from "zod"

// dotenv.config()

// const DEFAULT_HTTP = "http://localhost:3000"
// const DEFAULT_WS = "ws://localhost:3001"

// const envSchema = z.object({
// 	// Key infrastructure values
// 	DATABASE_URL: z
// 		.string()
// 		.optional()
// 		.default("postgresql://postgres:postgres@localhost:5434/postgres"),
// 	ALCHEMY_API_KEY: z.string(),
// 	NEXT_PUBLIC_WALLETCONNECT_ID: z.string(),

// 	// Public Next values
// 	NEXT_PUBLIC_APP_URL: z.string().optional().default(DEFAULT_HTTP),
// 	NEXT_PUBLIC_WS_URL: z.string().optional().default(DEFAULT_WS),
// 	NEXT_PUBLIC_EARLY_ACCESS: z.boolean().optional().default(false),

// 	// Private Next values
// 	NEXT_AUTH_URL: z.string().optional().default(DEFAULT_HTTP),
// 	NEXT_AUTH_SECRET: z.string()
// })

// export type Env = z.infer<typeof envSchema>

// const parsed = envSchema.safeParse(process.env)

// if (!parsed.success) {
// 	throw new Error(`Invalid environment variables: ${parsed.error}`)
// }

// export const env = parsed.data

// declare global {
// 	namespace NodeJS {
// 		interface ProcessEnv extends Env {}
// 	}
// }
