import { colors } from "../colors"

type CreatePlaceholderTokenProps = { status: string, name: string, days: number, color: keyof typeof colors }
const createDate = (days: number) => {
    const daysInMilliseconds = days * 86400000;
    return new Date(Date.now() + daysInMilliseconds);
};
const createPlaceholderActivity = ({ status, name, days, color }: CreatePlaceholderTokenProps) => {
	const createdAt = createDate(days)
	const endAt = Math.random() > 0.5 ? undefined : createDate(90)

	return {
		status,
		plug: { name, color },
		createdAt,
		startAt: createdAt,
		endAt,
		inputs: [],
		runs: []
	}
}

const PLACEHOLDER_MINT_BASEPAINT = createPlaceholderActivity({ status: "active", name: "Mint Daily Basepaint", days: 0, color: 'basepaint' })
const PLACEHOLDER_DELTA_NEUTRAL = createPlaceholderActivity({ status: "paused", name: "Refill Treasury Gas", days: -6, color: "plug" })
const PLACEHOLDER_LIQUIDATE_THE_LEVERAGE = createPlaceholderActivity({ status: "active", name: "Liquidiate the Leverage", days: -30, color: "morpho" })
const PLACEHOLDER_COVER_HEALTH_FACTOR = createPlaceholderActivity({ status: "active", name: "Protect Health Factor", days: -60, color: "aave" })
const PLACEHOLDER_STEAL_THE_MOON = createPlaceholderActivity({ status: "paused", name: "Loop USDC Borrows", days: -90, color: "euler" })

export const PLACEHOLDER_ACTIVITIES = [
	PLACEHOLDER_MINT_BASEPAINT,
	PLACEHOLDER_DELTA_NEUTRAL,
	PLACEHOLDER_COVER_HEALTH_FACTOR,
	PLACEHOLDER_LIQUIDATE_THE_LEVERAGE,
	PLACEHOLDER_STEAL_THE_MOON,
]
