export const VIEW_KEYS = {
	// Public access panels
	AUTHENTICATE: "AUTHENTICATE",
	HOME: "HOME",
	ADD: "ADD",
	PLUGS: "PLUGS",
	DISCOVER: "DISCOVER",
	PLUG: "PLUG",
	MY_PLUGS: "MY_PLUGS",
	ACTIVITY: "ACTIVITY",
	ALERTS: "ALERTS",
	ASSETS: "ASSETS",
	TOKENS: "TOKENS",
	COLLECTIBLES: "COLLECTIBLES",
	POSITIONS: "POSITIONS",
	EARNINGS: "EARNINGS",
	SETTINGS: "SETTINGS",
	SEARCH: "SEARCH",
	// Admin panels
	ADMIN: "ADMIN",
	PROFILE: "PROFILE"
}

export const DEFAULT_VIEWS = [
	// Prepare the single page experience for the mobile view.
	{ key: VIEW_KEYS.HOME, index: -1 },
	// Prepare the console view for the multi-column view.
	{ key: VIEW_KEYS.PLUGS, index: 0 },
	{ key: VIEW_KEYS.ACTIVITY, index: 1 }
]
