export * from "./analytics"
export * from "./chains"
export * from "./colors"
export * from "./routes"
export * from "./tokens"
export * from "./wallet"

export const tags = [
	"All",
	"DeFi",
	"NFT",
	"Consumer",
	"Degen",
	"MEV",
	"Social",
	"Trading",
	"Lending",
	"Borrowing"
] as const

export const greenGradientStyle = {
	background: "linear-gradient(30deg, #385842, #D2F38A)",
	WebkitBackgroundClip: "text",
	WebkitTextFillColor: "transparent"
}

export const sunGradientStyle = {
	background: "linear-gradient(30deg, #FFA800, #FAFF00)",
	WebkitBackgroundClip: "text",
	WebkitTextFillColor: "transparent"
}
