export const SOCKET_BASE_INCLUDE = {
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
			},
		},
		omit: {
			createdAt: true,
			updatedAt: true,
			socketId: true,
			farcasterId: true
		}
	}
}

export const SOCKET_BASE_QUERY = {
	include: SOCKET_BASE_INCLUDE
}
