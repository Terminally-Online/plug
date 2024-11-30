export const localRoutes = {
	brandKit: "/brand-kit/",
	comingSoon: "/coming-soon/",
	index: "/",
	app: "/app/"
}

export const staticRoutes = {
	vision: "/vision.pdf",
	memo: "/memo.pdf"
}

export const externalRoutes = {
	documentation: "https://docs.onplug.io",
	github: "https://github.com/nftchance/plug",
	status: "https://status.onplug.io",
	twitter: "https://twitter.com/onplug_io",
	earlyAccess: "https://docs.google.com/forms/d/e/1FAIpQLSf4ttqF5PizhP_F2jHBGTuaH-q6YunG4PkUcaK8JRhljXg5oQ/viewform"
}

export const routes = {
	...localRoutes,
	...staticRoutes,
	...externalRoutes
}
