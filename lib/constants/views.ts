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
	PROFILE: "PROFILE",

	// Admin views
	ADMIN: "ADMIN",

	// Temporal views
	APPLICATION: "APPLICATION"
}

const DEMO_VIEW_AS = "0x62180042606624f02d8a130da8a3171e9b33894d"
export const DEFAULT_VIEWS = [
	{ key: VIEW_KEYS.HOME, index: -1 },
	{ key: VIEW_KEYS.DISCOVER, index: 0 },
	{ key: VIEW_KEYS.MY_PLUGS, index: 1 },
	{ key: VIEW_KEYS.ACTIVITY, index: 2 },
	{ key: VIEW_KEYS.TOKENS, index: 3 }
]
export const DEFAULT_DEMO_VIEWS = DEFAULT_VIEWS.map(view => ({ ...view, viewAsId: DEMO_VIEW_AS }))

export const MOBILE_INDEX = -1
