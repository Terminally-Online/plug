export const SOCKET_BASE_INCLUDE = {
	columns: {
		orderBy: { index: "asc" },
		include: {
			viewAs: { include: { identity: { include: { ens: true } } } }
		},
		omit: { createdAt: true, updatedAt: true, socketId: true, viewAsId: true }
	},
	identity: {
		include: {
			farcaster: {
				include: {
					following: {
						include: {
							addresses: {
								omit: { createdAt: true, updatedAt: true }
							}
						},
						omit: { createdAt: true, updatedAt: true }
					}
				},
				omit: { createdAt: true, updatedAt: true }
			},
			ens: {
				omit: { createdAt: true, updatedAt: true }
			}
		},
		omit: {
			createdAt: true,
			updatedAt: true,
			socketId: true,
			farcasterId: true,
			ensName: true
		}
	}
} as const

export const SOCKET_BASE_QUERY = {
	include: SOCKET_BASE_INCLUDE,
	omit: { createdAt: true, updatedAt: true }
}
