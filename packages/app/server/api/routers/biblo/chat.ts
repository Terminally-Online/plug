import { z } from "zod"
import { Anthropic } from '@anthropic-ai/sdk'
import { createTRPCRouter, protectedProcedure } from "../../trpc"
import { env } from "@/env"
import { TOOLS } from "@/lib/tools"
import { schemas } from "@/lib"
import { TRPCError } from "@trpc/server"

const CLAUDE_MODEL = "claude-3-haiku-20240307"

const anthropic = new Anthropic({
	apiKey: env.ANTHROPIC_KEY
})

const messageSchema = z.object({
	message: z.string(),
	history: z.array(z.object({
		content: z.string(),
		role: z.enum(['user', 'assistant'])
	}))
})

interface MessageResponse {
    reply: string;
    additionalMessages?: string[];
    tools: string[];
}

export const chat = createTRPCRouter({
	message: protectedProcedure
		.input(messageSchema)
		.mutation(async ({ ctx, input }) => {
			const socket = await ctx.db.userSocket.findUnique({ where: { id: ctx.session.user.id } })
			if (!socket) throw new TRPCError({ code: "NOT_FOUND" })

			const messages = input.history.map(msg => ({
				role: msg.role,
				content: msg.content
			}))

			messages.push({
				role: 'user',
				content: input.message
			})

			const initialResponse = await anthropic.messages.create({
				model: CLAUDE_MODEL,
				max_tokens: 1024,
				messages,
				system: `You are Biblo, a founder of Plug that is the user-facing member helping them be as successful as possible. You have pride in your work and want them to use and succeed. 
				    Plug is an if-this-then-that like system that uses actions defined in schemas to build workflows. 
                                    They build a workflow in Plug and have everything automatically executed. Your purpose is to get them to build effective and profitable Plugs.
				    This is very important, users do not use different apps and we do not send them anywhere else. We recommend they build a Plug when we have that protocol schema supported. 
                  		    Available tools: ${Object.keys(TOOLS).map(toolKey => `${TOOLS[toolKey].name}: ${TOOLS[toolKey].description}`)}.
                  		    When appropriate, suggest relevant tools from this list.
                  		    You can provide multiple messages by using Claude's natural break points in your response.
	          		    You do not address or acknowledge the messages providing you data in <SYSTEM_RESPONSE> tags.
                  		    Current user: ${ctx.session.user.id}
				`,
			})
			const initialReply = initialResponse.content[0].type === 'text' ? initialResponse.content[0].text : ''
			const neededTools = parseToolSuggestions(initialReply)
			if (neededTools.length > 0) {
				const toolResults = await executeTools(socket.socketAddress, neededTools)
				const finalResponse = await anthropic.messages.create({
					model: CLAUDE_MODEL,
					max_tokens: 1024,
					messages: [
						...messages,
						{
							role: 'assistant',
							content: initialReply
						},
						{
							role: 'user',
							content: `<SYSTEM_RESPONSE>The results from the tools you requested are: ${JSON.stringify(toolResults, null, 2)}\nPlease provide a concise follow up if needed.</SYSTEM_RESPONSE>`
						}
					],
					system: `You are Biblo, a helpful assistant for the Plug platform.
						Keep responses clear and concise.
						Current user: ${ctx.session.user.id}`,
				})

				const finalMessages = finalResponse.content
					.filter(content => content.type === 'text')
					.map(content => content.text)

				return {
					reply: finalMessages[0],
					additionalMessages: finalMessages.slice(1),
					tools: neededTools
				}
			}

			const finalMessages = initialResponse.content
				.filter(content => content.type === 'text')
				.map(content => content.text)

			return {
				reply: finalMessages[0],
				additionalMessages: finalMessages.slice(1),
				tools: neededTools
			}
		})
})

async function executeTools(socketAddress: string, tools: string[]) {
	const results: Record<string, any> = {}
	for (const tool of tools) {
		switch (tool) {
			case 'holdings':
				results.holdings = [{ "name": "USDC", "value": 123456789 }]
				break
			case 'schemas':
				results.schemas = await schemas(undefined, undefined, 8453, socketAddress)
				break
			case 'schema':
				results.schemas = await schemas(undefined, undefined, 8453, socketAddress)
				break
		}
	}
	return results
}

function parseToolSuggestions(text: string | undefined): string[] {
	if (!text) return []
	const tools = ['holdings', 'schemas', 'price']
	return tools.filter(tool =>
		text.toLowerCase().includes(tool.toLowerCase())
	)
}
