type CreatePlaceholderTokenProps = { 
	name: string
	icon: string; 
	values: Array<number>; 
} 

const createPlaceholderPosition = ({ name, icon, values }: CreatePlaceholderTokenProps) => {
	const balance = 1;
	const price = 1;

	const positions = values.map(value => ({
		value: value, 
		change: Math.random() * 2 - 1 
	}));

	return {
		icon: icon.startsWith('http') ? icon : `https://cdn.zerion.io/${icon}.png`,
		name,
		chain: "base",
		positions,
		fungible: {
			icon: icon.startsWith('http') ? icon : `https://cdn.zerion.io/${icon}.png`,
			name,
			symbol: "placeholder",
			balance,
			change: Math.random() * Math.max(10, Math.random() * 10),
			value: balance * price,
			price: price,
			verified: false,
			implementations: [{
				contract: icon,
				chain: "base",
				balance: 1234,
				percentage: 100,
				createdAt: new Date(),
				updatedAt: new Date(),
				decimals: 6,
				fungibleName: name,
				fungibleSymbol: name,
				balances: [{ balance: 12345 }]
			}]
		}
	}
}

const PLACEHOLDER_YEARN = createPlaceholderPosition({
	name: "Yearn",
	icon: "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
	values: [9300, 1521, 24.34] 
});
const PLACEHOLDER_AAVE = createPlaceholderPosition({
	name: "Aave",
	icon: "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9",
	values: [150, 21050, 350.21] 
});
const PLACEHOLDER_MORPHO = createPlaceholderPosition({
	name: "Morpho",
	icon: "38f1c334-bbe0-4d99-aef4-e6d0a3a4207b",
	values: [1650.40, 3,300.82]
});
const PLACEHOLDER_EULER = createPlaceholderPosition({
	name: "Euler",
	icon: "0xd9fcd98c322942075a5c3860693e9f4f03aae07b",
	values: [56321.96] 
});

export const PLACEHOLDER_POSITIONS = [
	PLACEHOLDER_EULER,
	PLACEHOLDER_AAVE,
	PLACEHOLDER_YEARN,
	PLACEHOLDER_MORPHO,
];
