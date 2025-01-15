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

export type ZerionCollectibles = {
	links: {
		self: string
		next: string
	}
	data: Array<{
		type: string
		id: string
		attributes: {
			changed_at: string
			amount: string
			price: number
			value: number
			nft_info: {
				contract_address: string
				token_id: string
				name: string
				interface: string
				content: {
					preview: {
						url: string
					}
					detail: {
						url: string
					}
					video?: {
						url: string
					}
				}
				flags: {
					is_spam: boolean
				}
			}
			collection_info: {
				name: string
				description: string
				content: {
					icon: {
						url: string
					}
					banner: {
						url: string
					}
				}
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
			nft: {
				data: {
					type: string
					id: string
				}
			}
			nft_collection: {
				data: {
					type: string
					id: string
				}
			}
			wallet_nft_collection: {
				data: {
					type: string
					id: string
				}
			}
		}
	}>
}

export type ZerionFungibles = {
	links: {
		self: string
		next: string
	}
	data: Array<{
		type: string
		id: string
		attributes: {
			name: string
			symbol: string
			description: string
			icon: {
				url: string
			}
			flags: {
				verified: boolean
			}
			external_links: Array<{
				type: string
				name: string
				url: string
			}>
			implementations: Array<{
				chain_id: string
				address: string
				decimals: number
			}>
			market_data: {
				total_supply: number
				circulating_supply: number
				market_cap: number
				fully_diluted_valuation: number
				price: number
				changes: {
					percent_1d: number
					percent_30d: number
					percent_90d: number
					percent_365d: number
				}
			}
		}
		relationships: {
			chart_day: {
				links: {
					related: string
				}
				data: {
					type: string
					id: string
				}
			}
		}
		links: {
			self: string
		}
	}>
}
