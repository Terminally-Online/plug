import { Button } from "@/components/shared/buttons/button"
import { useState, useEffect } from "react"
import { Search } from "../../inputs/search"
import { SearchIcon } from "lucide-react"
import { cn, formatTitle } from "@/lib";
import { Counter } from "@/components/shared/utils/counter";
import { api } from "@/server/client";

interface Message {
	text: string;
	isSent: boolean;
}

const TypingIndicator = () => {
	return (
		<div className="flex justify-start">
			<div className="bg-plug-green/10 rounded-lg p-3 flex gap-1 items-center">
				<div className="w-2 h-2 rounded-full bg-black/60 animate-bounce" style={{ animationDelay: '0ms' }} />
				<div className="w-2 h-2 rounded-full bg-black/60 animate-bounce" style={{ animationDelay: '150ms' }} />
				<div className="w-2 h-2 rounded-full bg-black/60 animate-bounce" style={{ animationDelay: '300ms' }} />
			</div>
		</div>
	);
};

const DEFAULT_MESSAGES = [
	"Help me open an Aave V3 position",
	"What can I do with my tokens?",
	"What can I do with Plug?",
]

export const ColumnChat = ({ index }: { index: number }) => {
	
	const [message, setMessage] = useState("")
	const [messages, setMessages] = useState<Message[]>(() => [{
		text: "👋 Hey, I'm Piggy. I can answer nearly any question about anything you see here in Plug.\n\nMy knowledge may be limited.",
		isSent: false
	}])
	const [isTyping, setIsTyping] = useState(false)
	const [activeTools, setActiveTools] = useState<string[]>([])

	const chat = api.biblo.chat.message.useMutation()

	const handleSubmit = async (sent: string) => {
		if (!sent.trim()) return;

		setMessages([...messages, { text: sent, isSent: true }]);
		setIsTyping(true);
		setMessage("");

		try {
			const response = await chat.mutateAsync({
				message: sent,
				history: messages.map(msg => ({
					content: msg.text,
					role: msg.isSent ? 'user' : 'assistant'
				}))
			});

			if (response.tools?.length) {
				setActiveTools(prev => [...new Set([...prev, ...response.tools])]);
			}
			setMessages(prev => [...prev, { text: response.reply, isSent: false }, ...response.additionalMessages.map(text => ({ text, isSent: false }))]);
		} catch (error: any) {
			console.error('Detailed Error:', {
				error,
				cause: error.cause,
				data: error.data,
				shape: error.shape
			});
			
			const errorMessage = error.shape?.message || error.message || 'Unknown error occurred';
			const errorCode = error.shape?.code || error.code;
			
			setMessages(prev => [...prev, { 
				text: `Error Code ${errorCode}: ${errorMessage}`,
				isSent: false 
			}]);
		} finally {
			setIsTyping(false);
		}
	}

	return <>
		<div className="flex flex-col gap-2 p-4 h-full relative">
			<div className="absolute top-0 left-0 right-0 h-24 bg-gradient-to-b from-plug-white to-transparent" />


			<div className="messages flex flex-col-reverse gap-2 overflow-y-auto h-full pb-2 pt-24">
				{isTyping && <TypingIndicator />}
				{[...messages].reverse().map((msg, index) => (
					<div
						key={index}
						className={`flex ${msg.isSent ? 'justify-end' : 'justify-start'} px-2`}
					>
						<div
							className={cn(
								"font-bold max-w-[85%] rounded-lg p-3 whitespace-pre-wrap break-words",
								"shadow-sm",
								msg.isSent
									? 'bg-plug-yellow text-plug-green'
									: 'bg-plug-green/10 text-black/60'
							)}
						>
							{msg.text}
						</div>
					</div>
				))}
			</div>

			<div className="flex flex-col gap-2 h-max bg-plug-white pt-2">
				{activeTools.length > 0 && (
					<div className="flex flex-row gap-1 overflow-x-scroll w-full group">
						<Button className="ml-auto group-hover:hidden flex flex-row gap-2" variant="secondary" onClick={() => { }} sizing="sm">
							<Counter count={activeTools.length} /> Tools Used
						</Button>
						<div className="hidden group-hover:flex flex-row gap-2">
							{activeTools.map((tool, index) => (
								<Button
									key={index}
									className="w-max"
									variant="secondary"
									sizing="sm"
									onClick={() => { }}
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
							<span className="opacity-60">&quot;</span>{msg}<span className="opacity-60">&quot;</span>
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
						if (e.key === 'Enter' && !e.shiftKey) {
							e.preventDefault();
							handleSubmit(message);
						}
					}}
				/>
				<Button className="w-full py-4" type="submit" onClick={() => handleSubmit(message)}>Send</Button>
			</div>
		</div>
	</>
}
