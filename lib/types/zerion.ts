export type ZerionPositions = {
	links: {
		self: string
	}
	data: Array<{
		type: string
		id: string
		attributes: {
			parent: null
			protocol: string | null
			name: string
			position_type: string
			quantity: {
				int: string
				decimals: number
				float: number
				numeric: string
			}
			value: number
			price: number
			changes: {
				absolute_1d: number
				percent_1d: number
			}
			fungible_info: {
				name: string
				symbol: string
				icon: {
					url: string
				}
				flags: {
					verified: boolean
				}
				implementations: Array<{
					chain_id: string
					address: string
					decimals: number
				}>
			}
			flags: {
				displayable: boolean
				is_trash: boolean
			}
			updated_at: string
			updated_at_block: number | null
			application_metadata?: {
				name: string
				icon: {
					url: string
				}
				url: string
			}
		}
		relationships: {
			chain: {
				links: {
					related: string
				}
				data: {
					type: string
					id: string
				}
			}
			fungible: {
				links: {
					related: string
				}
				data: {
					type: string
					id: string
				}
			}
			dapp?: {
				data: {
					type: string
					id: string
				}
			}
		}
	}>
}
