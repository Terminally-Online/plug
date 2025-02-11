import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Anthropic } from "@anthropic-ai/sdk"
import { Message } from "@prisma/client"

import { env } from "@/env"
import { schemas } from "@/lib"
import { getPositions } from "@/lib/functions/zerion"
import { TOOLS } from "@/lib/tools"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

const CLAUDE_MODEL = "claude-3-5-haiku-20241022"

const anthropic = new Anthropic({
	apiKey: env.ANTHROPIC_KEY
})

// Add this before the parseDepositIntent function
interface Intent {
	chainId: number
	from: string
	inputs: Array<{
		protocol: string
		action: string
		token: string
		amount: string
	}>
}

export const chat = createTRPCRouter({
	message: protectedProcedure
		.input(
			z.object({
				message: z.string(),
				history: z.array(
					z.object({
						content: z.string(),
						role: z.enum(["user", "assistant"])
					})
				)
			})
		)
		.mutation(async ({ ctx, input }) => {
			const socket = await ctx.db.userSocket.findUnique({ where: { id: ctx.session.user.id } })
			if (!socket) throw new TRPCError({ code: "NOT_FOUND" })

			await ctx.db.message.create({
				data: {
					socketId: socket.id,
					content: input.message,
					isUser: true,
					timeSent: new Date(),
					tools: null
				}
			})

			// Parse deposit intent from user message
			const depositIntent = parseDepositIntent(input.message)
			if (depositIntent) {
				// Create a Workflow based on the deposit intent
				const workflow = await ctx.db.workflow.create({
					data: {
						socketId: socket.id,
						intentData: JSON.stringify(depositIntent),
						createdAt: new Date()
					}
				})
				console.log("Workflow Created:", workflow)

				// Now create the Execution based on the Workflow
				await ctx.db.execution.create({
					data: {
						workflowId: workflow.id,
						intentData: JSON.stringify(depositIntent),
						createdAt: new Date()
					}
				})
				console.log("Execution Created for Workflow:", workflow.id)
			}

			const messages = input.history.map(msg => ({
				role: msg.role,
				content: msg.content
			}))

			messages.push({
				role: "user",
				content: input.message
			})

			const personality = `
			    You are Piggy, a founder of Plug that is the user-facing member helping them be as successful as possible.

				- When making recommendations of what can be done, you only recommend data you have confirmed in your direct context.
				- When you are referencing a protocol retrieve the schema for that protocol and confirm the support we have for it.
				- When we do not support a protocol let the user know that we have notified the team and to check back soon.
				    
				Plug is an if-this-then-that like system that uses actions defined in schemas to build workflows.
				Users of Plug build workflows in Plug to automate their onchain activity.
				Your purpose is to get them to build effective and profitable Plugs.
				This is very important, users do not use different apps and we do not send them anywhere else.
				We recommend they build a Plug when we have that protocol schema supported.

				You do not use emojis in your responses.
				You understand that the user has a wallet with holdings that is seperate from their Plug account (Socket).
				You encourage the user to deposit their holdings into their Socket to use in Plugs.
			`

			const tools = `
				Available tools: ${Object.keys(TOOLS).map(toolKey => `${TOOLS[toolKey].name}: ${TOOLS[toolKey].description}`)}.
				When appropriate, suggest relevant tools from this list.
			`
			const user = `Current user: ${ctx.session.user.id}`

			const initialResponse = await anthropic.messages.create({
				model: CLAUDE_MODEL,
				max_tokens: 512,
				messages,
				system: `
					${personality}
					${tools}
					${user}
				`
			})
			const initialReply = initialResponse.content[0].type === "text" ? initialResponse.content[0].text : ""
			const neededTools = parseToolSuggestions(initialReply)

			await ctx.db.message.create({
				data: {
					socketId: socket.id,
					content: initialReply,
					isUser: false,
					timeSent: new Date(),
					tools: neededTools.length > 0 ? neededTools.join(", ") : null
				}
			})

			if (neededTools.length > 0) {
				const toolResults = await executeTools(
					ctx.session.user.id,
					socket.socketAddress,
					neededTools,
					input.message
				)
				console.log("Tool Results:", JSON.stringify(toolResults, null, 2))

				const finalResponse = await anthropic.messages.create({
					model: CLAUDE_MODEL,
					max_tokens: 1024,
					messages: [
						...messages,
						{
							role: "assistant",
							content: initialReply
						},
						{
							role: "user",
							content: `<SYSTEM_RESPONSE>The results from the tools you requested are: ${JSON.stringify(toolResults, null, 2)}\nPlease provide a concise follow up if needed.</SYSTEM_RESPONSE>`
						}
					],
					system: `
						${personality}
						${user}
					`
				})

				const finalMessages = finalResponse.content
					.filter(content => content.type === "text")
					.map(content => content.text)

				return {
					reply: finalMessages[0],
					additionalMessages: finalMessages.slice(1),
					tools: neededTools
				}
			}

			return {
				reply: initialReply,
				additionalMessages: [],
				tools: neededTools
			}
		}),

	getMessages: protectedProcedure.input(z.object({ socketId: z.string() })).query(async ({ ctx, input }) => {
		const messages = await ctx.db.message.findMany({
			where: { socketId: input.socketId },
			orderBy: { timeSent: "asc" }
		})
		return messages
	})
})

async function executeTools(socketId: string, socketAddress: string, tools: string[], message: string) {
	const results: Record<string, any> = {}
	for (const tool of tools) {
		if (tool in TOOLS) {
			try {
				results[tool] = await TOOLS[tool].execute(socketId, socketAddress)
				// If tool returned an error, log it but continue
				if (results[tool]?.error) {
					console.error(`Tool ${tool} failed:`, results[tool].error)
				}
			} catch (error) {
				console.error(`Tool ${tool} execution failed:`, error)
				results[tool] = { error: `Failed to execute ${tool}` }
			}
		}
	}
	return results
}

function parseToolSuggestions(text: string | undefined): string[] {
	if (!text) return []
	return Object.keys(TOOLS).filter(tool => text.toLowerCase().includes(tool.toLowerCase()))
}

async function handleError(ctx, errorMessage, socketId) {
	await ctx.db.message.create({
		data: {
			socketId: socketId,
			content: errorMessage,
			isUser: false,
			timeSent: new Date()
		}
	})
}

// Function to parse deposit intents from user messages
function parseDepositIntent(message: string): Intent | null {
	const regex = /deposit (\d+(\.\d+)?) (\w+) to Aave/i // Regex to match "deposit 0.005 ETH to Aave"
	const match = message.match(regex)

	if (match) {
		const amount = match[1]
		const token = match[3]
		return {
			chainId: 8453,
			from: "", // Placeholder for user's address
			inputs: [
				{
					protocol: "aave",
					action: "supply",
					token,
					amount
				}
			]
		}
	}
	return null // Return null if no match found
}
