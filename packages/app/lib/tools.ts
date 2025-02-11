import { getPrices } from "@/lib/functions/llama"
import { getPositions } from "@/lib/functions/zerion/positions"

import { schemas } from "./functions/plug"
import { intent } from "./functions/plug/solver"

// Define the Intent type directly in this file
interface Intent {
	chainId: number // e.g., 8453 for Base chain
	from: string // User's address
	inputs: Array<{
		protocol: string // e.g., "aave"
		action: string // e.g., "supply"
		token: string // Token address
		amount: string // Amount to deposit
	}>
}

export const TOOLS: Record<
	string,
	{
		name: string
		description: string
		execute: (socketId: string, socketAddress: string) => Promise<any>
	}
> = {
	schemas: {
		name: "schemas",
		description:
			"List all available schemas in the system for integrated protocols that allow users to create workflows and bundle transactions. In order to use a specific schema the user needs to access the 'schema' tool",
		execute: async () => {
			try {
				return await schemas(undefined, undefined, 8453, "")
			} catch (error) {
				console.error("Schema tool failed:", error)
				return { error: "Failed to fetch schemas" }
			}
		}
	},
	aave_supply: {
		name: "aave_supply",
		description:
			"Get the Aave supply schema which allows users to supply assets as collateral to Aave. Shows available tokens and their current supply APY.",
		execute: async (_, socketAddress) => {
			const result = await schemas("aave", "supply", 8453, socketAddress)
			console.log("Aave Supply Schema Response:", JSON.stringify(result, null, 2))
			return result
		}
	},
	aave_borrow: {
		name: "aave_borrow",
		description:
			"Get the Aave borrow schema which allows users to borrow assets against their supplied collateral. Shows available tokens and their current borrow APY.",
		execute: async (_, socketAddress) => await schemas("aave", "borrow", 8453, socketAddress)
	},
	aave_repay: {
		name: "aave_repay",
		description:
			"Get the Aave repay schema which allows users to repay borrowed assets. Shows current borrowed positions that can be repaid.",
		execute: async (_, socketAddress) => await schemas("aave", "repay", 8453, socketAddress)
	},
	aave_withdraw: {
		name: "aave_withdraw",
		description:
			"Get the Aave withdraw schema which allows users to withdraw their supplied assets. Shows current supplied positions that can be withdrawn.",
		execute: async (_, socketAddress) => await schemas("aave", "withdraw", 8453, socketAddress)
	},
	aave_deposit: {
		name: "aave_deposit",
		description: "Deposit assets into Aave V3.",
		execute: async (socketId, socketAddress) => {
			// Implement the deposit logic here
			const depositAmount = "0.005" // Example amount
			const tokenAddress = "0x..." // Example token address

			const depositIntent: Intent = {
				chainId: 8453,
				from: socketAddress,
				inputs: [
					{
						protocol: "aave",
						action: "supply",
						token: tokenAddress,
						amount: depositAmount
					}
				]
			}

			try {
				const response = await intent(depositIntent)
				return response
			} catch (error) {
				console.error("Deposit tool failed:", error)
				return { error: "Failed to execute deposit" }
			}
		}
	},
	aave_health_factor: {
		name: "aave_health_factor",
		description: "Check the health factor for Aave V3.",
		execute: async (socketId, socketAddress) => {
			// Implement the health factor check logic here
		}
	},
	aave_apy: {
		name: "aave_apy",
		description: "Check the APY for Aave V3.",
		execute: async (socketId, socketAddress) => {
			try {
				const response = await intent({
					chainId: 1, // or whatever chain you're targeting
					from: socketAddress,
					inputs: [
						{
							protocol: "aave_v3",
							action: "apy",
							direction: 1, // 1 for deposit, -1 for borrow
							token: "0x...", // token address
							operator: 1, // 1 for greater than, -1 for less than
							threshold: "0.05" // 5%
						}
					]
				})

				return response
			} catch (error) {
				console.error("APY tool failed:", error)
				return { error: "Failed to fetch APY data" }
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
	},
	aave_monitor_borrow_apy: {
		name: "aave_monitor_borrow_apy",
		description: "Monitor Aave V3 borrow APY for a specific asset",
		execute: async (socketId: string, socketAddress: string) => {
			try {
				// 1. First get the schema for the APY constraint
				const schemaResponse = await intent({
					chainId: 8453,
					from: socketAddress,
					inputs: [
						{
							protocol: "aave_v3",
							action: "schemas"
						}
					]
				})

				// 2. Find the APY constraint schema
				const apySchema = schemaResponse.schemas.find((s: any) => s.action === "apy" && s.type === "constraint")

				if (!apySchema) {
					return { error: "APY monitoring schema not found" }
				}

				// 3. Return the schema info so the UI can collect inputs
				return {
					schema: apySchema,
					requiredInputs: [
						{
							name: "token",
							type: "address",
							description: "Asset to monitor (e.g. ETH, USDC)"
						},
						{
							name: "threshold",
							type: "float",
							description: "APY threshold percentage"
						}
					]
				}
			} catch (error) {
				console.error("Failed to get APY schema:", error)
				return { error: "Failed to setup APY monitoring" }
			}
		}
	}
	// price: {
	// 	name: "price",
	// 	description: "Get current price information for specific tokens",
	// 	execute: async () => {
	// 		return await getPrices(tokens)
	// 	}
	// },
	// TODO: Add specific tools for schema
}
