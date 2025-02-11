import { useEffect, useState } from "react"
import Markdown from "react-markdown"

import { SearchIcon } from "lucide-react"

import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { cn, formatTitle } from "@/lib"
import { api } from "@/server/client"

import { Search } from "../../inputs/search"

interface Message {
	text: string
	isSent: boolean
}

const TypingIndicator = () => {
	return (
		<div className="flex justify-start">
			<div className="flex items-center gap-1 rounded-lg bg-plug-green/10 p-3">
				<div className="h-2 w-2 animate-bounce rounded-full bg-black/60" style={{ animationDelay: "0ms" }} />
				<div className="h-2 w-2 animate-bounce rounded-full bg-black/60" style={{ animationDelay: "150ms" }} />
				<div className="h-2 w-2 animate-bounce rounded-full bg-black/60" style={{ animationDelay: "300ms" }} />
			</div>
		</div>
	)
}

const DEFAULT_MESSAGES = [
	"Help me open an Aave V3 position",
	"What can I do with my tokens?",
	"What can I do with Plug?"
]

export const ColumnChat = ({ index }: { index: number }) => {
	const [message, setMessage] = useState("")
	const [messages, setMessages] = useState<Message[]>(() => [
		{
			text: "👋 Hey, I'm Morgan. I can answer nearly any question about anything you see here in Plug.\n\nMy knowledge may be limited.",
			isSent: false
		}
	])
	const [isTyping, setIsTyping] = useState(false)
	const [activeTools, setActiveTools] = useState<string[]>([])

	const chat = api.biblo.chat.message.useMutation()

	const handleSubmit = async (sent: string) => {
		if (!sent.trim()) return

		setMessages(prev => [...prev, { text: sent, isSent: true }])
		setIsTyping(true)
		setMessage("")

		try {
			const response = await chat.mutateAsync({
				message: sent,
				history: messages.slice(-20).map(msg => ({
					content: msg.text,
					role: msg.isSent ? "user" : "assistant"
				}))
			})

			if (response.tools?.length) {
				setActiveTools(prev => [...new Set([...prev, ...response.tools])])
			}

			const fullResponse = [response.reply, ...response.additionalMessages].join("\n\n")

			setMessages(prev => [...prev, { text: fullResponse, isSent: false }])
		} catch (error: any) {
			console.error("Detailed Error:", {
				error,
				cause: error.cause,
				data: error.data,
				shape: error.shape
			})

			const errorMessage = error.shape?.message || error.message || "Unknown error occurred"
			const errorCode = error.shape?.code || error.code

			setMessages(prev => [
				...prev,
				{
					text: `Error Code ${errorCode}: ${errorMessage}`,
					isSent: false
				}
			])
		} finally {
			setIsTyping(false)
		}
	}

	return (
		<>
			<div className="relative flex h-full flex-col gap-2 p-4">
				<div className="absolute left-0 right-0 top-0 h-24 bg-gradient-to-b from-plug-white to-transparent" />

				<div className="messages flex h-full flex-col-reverse gap-2 overflow-y-auto pb-2 pt-24">
					{isTyping && <TypingIndicator />}
					{[...messages].reverse().map((msg, index) => (
						<div key={index} className={`flex ${msg.isSent ? "justify-end" : "justify-start"} px-2`}>
							<div
								className={cn(
									"max-w-[100%] rounded-lg p-4 font-bold",
									msg.isSent ? "bg-plug-yellow text-plug-green" : "bg-plug-green/10 text-black/60"
								)}
							>
								<Markdown
									components={{
										ul: ({ children }) => (
											<ul className="ml-8 mr-4 list-disc text-justify">{children}</ul>
										),
										ol: ({ children }) => (
											<ol className="ml-8 mr-4 list-decimal text-justify">{children}</ol>
										),
										li: ({ children }) => <li className="list-item">{children}</li>,
										p: ({ children }) => <p className="">{children}</p>
									}}
								>
									{msg.text}
								</Markdown>
							</div>
						</div>
					))}
				</div>

				<div className="flex h-max flex-col gap-2 bg-plug-white pt-2">
					{activeTools.length > 0 && (
						<div className="group flex w-full flex-row gap-1 overflow-x-scroll">
							<Button
								className="ml-auto flex flex-row gap-2 group-hover:hidden"
								variant="secondary"
								onClick={() => {}}
								sizing="sm"
							>
								<Counter count={activeTools.length} /> Tools Used
							</Button>
							<div className="hidden flex-row gap-2 group-hover:flex">
								{activeTools.map((tool, index) => (
									<Button
										key={index}
										className="w-max"
										variant="secondary"
										sizing="sm"
										onClick={() => {}}
									>
										{formatTitle(tool)}
									</Button>
								))}
							</div>
						</div>
					)}

					<div className="flex flex-row gap-2 overflow-x-scroll">
						{DEFAULT_MESSAGES.map((msg, index) => (
							<Button
								key={index}
								className="w-max"
								variant="secondary"
								onClick={() => handleSubmit(msg)}
								sizing="sm"
							>
								<span className="opacity-60">&quot;</span>
								{msg}
								<span className="opacity-60">&quot;</span>
							</Button>
						))}
					</div>

					<Search
						icon={<SearchIcon size={16} />}
						search={message}
						handleSearch={message => setMessage(message)}
						placeholder="Type a message..."
						className="message-input"
						onKeyDown={(e: React.KeyboardEvent) => {
							if (e.key === "Enter" && !e.shiftKey) {
								e.preventDefault()
								handleSubmit(message)
							}
						}}
					/>
					<Button className="w-full py-4" type="submit" onClick={() => handleSubmit(message)}>
						Send
					</Button>
				</div>
			</div>
		</>
	)
}
