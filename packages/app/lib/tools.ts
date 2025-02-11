import { getPositions } from "@/lib/functions/zerion/positions"

import { schemas } from "./functions/plug"

export const TOOLS: Record<
	string,
	{
		name: string
		description: string
		execute: (socketId: string, socketAddress: string, protocol?: string, action?: string) => Promise<any>
	}
> = {
	introspection: {
		name: "introspection",
		description: "See what tools Morgan currently supports and has the ability to utilize in her chats with users.",
		execute: async () => {
			return `
					Available tools: ${Object.keys(TOOLS).map(toolKey => `${TOOLS[toolKey].name}: ${TOOLS[toolKey].description}`)}.
					When appropriate, suggest relevant tools from this list.
				`
		}
	},
	schema: {
		name: "schema",
		description:
			"Get detailed information about a specific protocol's actions. Use this when a user asks about a specific protocol (protocol: X) or action (action: Y). This tool helps users understand what they can do with a particular protocol.",
		execute: async (socketId, socketAddress, protocol, action) => {
			try {
				console.log("Schema tool called with:", protocol, action)
				return await schemas(protocol, action, 8453, "")
			} catch (error) {
				console.error("Schema tool failed:", error)
				return { error: "Failed to fetch schemas" }
			}
		}
	},
	schemas: {
		name: "schemas",
		description:
			"Get an overview of all available protocols and their actions. Use this when a user is new or unsure what they want to do. The response will help guide them to specific protocols they can then explore with the schema tool.",
		execute: async (socketId, socketAddress) => {
			try {
				return await schemas(undefined, undefined, 8453, "")
			} catch (error) {
				console.error("Schema tool failed:", error)
				return { error: "Failed to fetch schemas" }
			}
		}
	},
	user_holdings: {
		name: "user_holdings",
		description:
			"Get the current fungible holding data for the specific connected wallet and let the user know that they should deposit these assets to their Socket so they can be used in Plugs.",
		execute: async socketId => {
			try {
				return await getPositions(socketId)
			} catch (error) {
				console.error("User holdings tool failed:", error)
				return { error: "Failed to fetch user holdings" }
			}
		}
	},
	socket_holdings: {
		name: "socket_holdings",
		description:
			"Get the current fungible holding data for a specific socket (the Plug account of a connected wallet).",
		execute: async (socketId, socketAddress) => await getPositions(socketId, socketAddress)
	}
	// price: {
	// 	name: "price",
	// 	description: "Get current price information for specific tokens",
	// 	execute: async () => {
	// 		return await getPrices(tokens)
	// 	}
	// },
}
