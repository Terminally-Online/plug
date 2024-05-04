export const localRoutes = {
	brandKit: "/brand-kit/",
	comingSoon: "/coming-soon/",
	index: "/"
}

export const staticRoutes = {
	vision: "/vision.pdf",
	memo: "/memo.pdf"
}

export const externalRoutes = {
	documentation: "https://docs.onplug.io",
	github: "https://github.com/nftchance/plug",
	status: "https://status.onplug.io",
	twitter: "https://twitter.com/onplug_io"
}

export const routes = {
	...localRoutes,
	...staticRoutes,
	...externalRoutes
}
