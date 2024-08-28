export const VIEW_KEYS = {
	// Anonymous views
	HOME: "HOME",
	AUTHENTICATE: "AUTHENTICATE",
	ADD: "ADD",
	PLUGS: "PLUGS",
	DISCOVER: "DISCOVER",
	MY_PLUGS: "MY_PLUGS",
	PLUG: "PLUG",
	SEARCH: "SEARCH",
	ALERTS: "ALERTS",
	VIEW_AS: "VIEW_AS",

	// Authenticated views
	ACTIVITY: "ACTIVITY",
	ASSETS: "ASSETS",
	TOKENS: "TOKENS",
	COLLECTIBLES: "COLLECTIBLES",
	POSITIONS: "POSITIONS",
	EARNINGS: "EARNINGS",
	SETTINGS: "SETTINGS",

	// Admin views
	ADMIN: "ADMIN",
	PROFILE: "PROFILE"
}

export const DEFAULT_ANONYMOUS_VIEWS = [
	{ key: VIEW_KEYS.HOME, index: -1 },
	{ key: VIEW_KEYS.DISCOVER, index: 0 },
	{ key: VIEW_KEYS.MY_PLUGS, index: 1 }
]

export const DEFAULT_VIEWS = [
	{ key: VIEW_KEYS.HOME, index: -1 },
	{ key: VIEW_KEYS.PLUGS, index: 0 },
	{ key: VIEW_KEYS.ACTIVITY, index: 1 }
]
