import { Socket } from "@prisma/client"
import { colors } from "../colors"
import { RouterOutputs } from "@/server/client"

type CreatePlaceholderPlugProps = { name: string, color: keyof typeof colors, views: number, forks: number }
export const createPlaceholderPlug = ({ name, color, views, forks }: CreatePlaceholderPlugProps): NonNullable<RouterOutputs['plugs']['all']>[number] => {
	return {
		id: '',
		name,
		color,
		views: [{ views }],
		forkCount: forks,
		socketId: "",
		isCurated: false,
		isPrivate: false,
		actions: "",
		tags: [],
		frequency: 0,
		plugForkedId: null,
		intentIds: [],
		namedAt: new Date(),
		renamedAt: new Date(),
		createdAt: new Date(),
		updatedAt: new Date(),
		socket: {
			id: "",
			admin: false,
			socketAddress: "",
			deploymentDelegate: "",
			deploymentFactory: "",
			deploymentNonce: 0,
			deploymentImplementation: "",
			deploymentSalt: "",
			identity: {
				socketId: "",
				referralCode: "",
				requestedAt: new Date(),
				approvedAt: new Date(),
				onboardingColor: "",
				onboardingCount: 2,
				ens: null,
				farcasterId: null,
				referrerId: null,
				onboardingAt: new Date(),
				onboardedAt: new Date(),
				createdAt: new Date(),
				updatedAt: new Date()
			},
			createdAt: new Date(),
			updatedAt: new Date()
		}
	}
}

const PLACEHOLDER_MINT_BASEPAINT = createPlaceholderPlug({ name: "Mint Daily Basepaint", color: 'basepaint', views: 16293, forks: 3014 })
const PLACEHOLDER_TREASURY_GAS = createPlaceholderPlug({ name: "Refill Treasury Gas", color: "plug", views: 24781, forks: 9 })
const PLACEHOLDER_LIQUIDATE_THE_LEVERAGE = createPlaceholderPlug({ name: "Liquidiate the Leverage", color: "morpho", views: 71920, forks: 152 })
const PLACEHOLDER_COVER_HEALTH_FACTOR = createPlaceholderPlug({ name: "Protect Health Factor", color: "aave", views: 37014, forks: 591 })
const PLACEHOLDER_LOOP_BORROWS = createPlaceholderPlug({ name: "Loop USDC Borrows", color: "euler", views: 17930, forks: 81 })
const PLACEHOLDER_STEAL_THE_MOON = createPlaceholderPlug({ name: "Steal the Moon", color: "plug", views: 930, forks: 801 })

export const PLACEHOLDER_PLUGS = [
	PLACEHOLDER_MINT_BASEPAINT,
	PLACEHOLDER_TREASURY_GAS,
	PLACEHOLDER_LIQUIDATE_THE_LEVERAGE,
	PLACEHOLDER_COVER_HEALTH_FACTOR,
	PLACEHOLDER_LOOP_BORROWS,
	PLACEHOLDER_STEAL_THE_MOON
]
