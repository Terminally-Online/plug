export const VIEW_KEYS = {
	HOME: "HOME",
	ADD: "ADD",
	PLUGS: "PLUGS",
	DISCOVER: "DISCOVER",
	PLUG: "PLUG",
	MY_PLUGS: "MY_PLUGS",
	ACTIVITY: "ACTIVITY",
	ASSETS: "ASSETS",
	TOKENS: "TOKENS",
	COLLECTIBLES: "COLLECTIBLES",
	POSITIONS: "POSITIONS",
	EARNINGS: "EARNINGS",
	SETTINGS: "SETTINGS"
}

export const DEFAULT_VIEWS = [
	// Prepare the single page experience for the mobile view.
	{ key: VIEW_KEYS.PLUGS, index: -1 },
	// Prepare the console view for the multi-column view.
	{ key: VIEW_KEYS.PLUGS, index: 0 },
	{ key: VIEW_KEYS.ACTIVITY, index: 1 }
]
