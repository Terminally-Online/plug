import { chains } from "@/lib/blockchain"

export type ChainId = (typeof chains)[number]["id"]
