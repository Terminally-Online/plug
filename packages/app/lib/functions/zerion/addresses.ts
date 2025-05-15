export const zerionChains = {
	arbitrum: {
		type: "chains",
		id: "arbitrum",
		attributes: {
			external_id: "0xa4b1",
			name: "Arbitrum",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/arbitrum.png"
			},
			explorer: {
				name: "Arbiscan",
				token_url_format: "https://arbiscan.io/token/{ADDRESS}",
				tx_url_format: "https://arbiscan.io/tx/{HASH}",
				home_url: "https://arbiscan.io"
			},
			rpc: {
				public_servers_url: ["https://arb1.arbitrum.io/rpc"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/arbitrum"
		}
	},
	"binance-smart-chain": {
		type: "chains",
		id: "binance-smart-chain",
		attributes: {
			external_id: "0x38",
			name: "BNB Chain",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/bsc.png"
			},
			explorer: {
				name: "BscScan",
				token_url_format: "https://bscscan.com/token/{ADDRESS}",
				tx_url_format: "https://bscscan.com/tx/{HASH}",
				home_url: "https://bscscan.com"
			},
			rpc: {
				public_servers_url: [
					"https://bsc-dataseed1.defibit.io",
					"https://bsc-dataseed2.defibit.io",
					"https://bsc-dataseed3.defibit.io",
					"https://bsc-dataseed4.defibit.io",
					"https://bsc-dataseed1.ninicoin.io",
					"https://bsc-dataseed2.ninicoin.io",
					"https://bsc-dataseed3.ninicoin.io",
					"https://bsc-dataseed4.ninicoin.io",
					"https://bsc-dataseed1.binance.org",
					"https://bsc-dataseed2.binance.org",
					"https://bsc-dataseed3.binance.org",
					"https://bsc-dataseed4.binance.org"
				]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xb8c77482e45f1f44de1745f52c74426c631bdd52"
				},
				data: {
					type: "fungibles",
					id: "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
				},
				data: {
					type: "fungibles",
					id: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/binance-smart-chain"
		}
	},
	ethereum: {
		type: "chains",
		id: "ethereum",
		attributes: {
			external_id: "0x1",
			name: "Ethereum",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/ethereum.png"
			},
			explorer: {
				name: "Etherscan",
				token_url_format: "https://etherscan.io/token/{ADDRESS}",
				tx_url_format: "https://etherscan.io/tx/{HASH}",
				home_url: "https://etherscan.io"
			},
			rpc: {
				public_servers_url: ["https://eth.llamarpc.com", "https://mainnet.gateway.tenderly.co"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/ethereum"
		}
	},
	blast: {
		type: "chains",
		id: "blast",
		attributes: {
			external_id: "0x13e31",
			name: "Blast",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/81457"
			},
			explorer: {
				name: "Blastscan",
				token_url_format: "https://blastscan.io/token/{ADDRESS}",
				tx_url_format: "https://blastscan.io/tx/{HASH}",
				home_url: "https://blastscan.io"
			},
			rpc: {
				public_servers_url: ["https://rpc.blast.io/"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/blast"
		}
	},
	abstract: {
		type: "chains",
		id: "abstract",
		attributes: {
			external_id: "0xab5",
			name: "Abstract",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/abstract.png"
			},
			explorer: {
				name: "Abstract Explorer",
				token_url_format: "https://abscan.org/token/{ADDRESS}",
				tx_url_format: "https://abscan.org/tx/{HASH}",
				home_url: "https://abscan.org"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/abstract"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/abstract"
		}
	},
	avalanche: {
		type: "chains",
		id: "avalanche",
		attributes: {
			external_id: "0xa86a",
			name: "Avalanche",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/avalanche.png"
			},
			explorer: {
				name: "SnowScan",
				token_url_format: "https://snowscan.xyz/token/{ADDRESS}",
				tx_url_format: "https://snowscan.xyz/tx/{HASH}",
				home_url: "https://snowscan.xyz"
			},
			rpc: {
				public_servers_url: ["https://api.avax.network/ext/bc/C/rpc"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/43e05303-bf43-48df-be45-352d7567ff39"
				},
				data: {
					type: "fungibles",
					id: "43e05303-bf43-48df-be45-352d7567ff39"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"
				},
				data: {
					type: "fungibles",
					id: "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/avalanche"
		}
	},
	opbnb: {
		type: "chains",
		id: "opbnb",
		attributes: {
			external_id: "0xcc",
			name: "opBNB",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/opBNB.png"
			},
			explorer: {
				name: "opBNB Block Explorer",
				token_url_format: "https://mainnet.opbnbscan.com/token/{ADDRESS}",
				tx_url_format: "https://mainnet.opbnbscan.com/tx/{HASH}",
				home_url: "https://mainnet.opbnbscan.com"
			},
			rpc: {
				public_servers_url: [
					"https://opbnb-mainnet-rpc.bnbchain.org",
					"https://opbnb-mainnet.nodereal.io/v1/64a9df0874fb4a93b9d0a3849de012d3",
					"wss://opbnb-mainnet.nodereal.io/ws/v1/64a9df0874fb4a93b9d0a3849de012d3",
					"https://opbnb-mainnet.nodereal.io/v1/e9a36765eb8a40b9bd12e680a1fd2bc5",
					"wss://opbnb-mainnet.nodereal.io/ws/v1/e9a36765eb8a40b9bd12e680a1fd2bc5",
					"https://opbnb-rpc.publicnode.com",
					"wss://opbnb-rpc.publicnode.com",
					"https://opbnb.drpc.org",
					"wss://opbnb.drpc.org"
				]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xb8c77482e45f1f44de1745f52c74426c631bdd52"
				},
				data: {
					type: "fungibles",
					id: "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/opbnb"
		}
	},
	"astar-zkevm": {
		type: "chains",
		id: "astar-zkevm",
		attributes: {
			external_id: "0xec0",
			name: "Astar zkEVM",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/3776"
			},
			explorer: {
				name: "Blockscout Astar zkEVM explorer",
				token_url_format: "https://astar-zkevm.explorer.startale.com/token/{ADDRESS}",
				tx_url_format: "https://astar-zkevm.explorer.startale.com/tx/{HASH}",
				home_url: "https://astar-zkevm.explorer.startale.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.startale.com/astar-zkevm"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/astar-zkevm"
		}
	},
	"zksync-era": {
		type: "chains",
		id: "zksync-era",
		attributes: {
			external_id: "0x144",
			name: "ZKsync Era",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/324"
			},
			explorer: {
				name: "ZKsync Era Block Explorer",
				token_url_format: "https://era.zksync.network/address/{ADDRESS}",
				tx_url_format: "https://era.zksync.network/tx/{HASH}",
				home_url: "https://era.zksync.network"
			},
			rpc: {
				public_servers_url: ["https://mainnet.era.zksync.io"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/zksync-era"
		}
	},
	base: {
		type: "chains",
		id: "base",
		attributes: {
			external_id: "0x2105",
			name: "Base",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/8453"
			},
			explorer: {
				name: "Base Explorer",
				token_url_format: "https://basescan.org/token/{ADDRESS}",
				tx_url_format: "https://basescan.org/tx/{HASH}",
				home_url: "https://basescan.org"
			},
			rpc: {
				public_servers_url: ["https://mainnet.base.org"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/base"
		}
	},
	berachain: {
		type: "chains",
		id: "berachain",
		attributes: {
			external_id: "0x138de",
			name: "Berachain",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/berra.png"
			},
			explorer: {
				name: "Berachain Explorer",
				token_url_format: "https://berascan.com/token/{ADDRESS}",
				tx_url_format: "https://berascan.com/tx/{HASH}",
				home_url: "https://berascan.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/berachain"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/7795d362-16cd-4418-b715-03e99e360823"
				},
				data: {
					type: "fungibles",
					id: "7795d362-16cd-4418-b715-03e99e360823"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/6a42d07c-fe9a-42bc-813e-2db0c5a747a1"
				},
				data: {
					type: "fungibles",
					id: "6a42d07c-fe9a-42bc-813e-2db0c5a747a1"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/berachain"
		}
	},
	bob: {
		type: "chains",
		id: "bob",
		attributes: {
			external_id: "0xed88",
			name: "BOB",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/bob.png"
			},
			explorer: {
				name: "bobscout",
				token_url_format: "https://explorer.gobob.xyz/token/{ADDRESS}",
				tx_url_format: "https://explorer.gobob.xyz/tx/{HASH}",
				home_url: "https://explorer.gobob.xyz"
			},
			rpc: {
				public_servers_url: ["https://rpc.gobob.xyz", "wss://rpc.gobob.xyz"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/bob"
		}
	},
	cyber: {
		type: "chains",
		id: "cyber",
		attributes: {
			external_id: "0x1d88",
			name: "Cyber",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/cyber.png"
			},
			explorer: {
				name: "Cyber Explorer",
				token_url_format: "https://cyberscan.co/token/{ADDRESS}",
				tx_url_format: "https://cyberscan.co/tx/{HASH}",
				home_url: "https://cyberscan.co"
			},
			rpc: {
				public_servers_url: [
					"https://cyber.alt.technology/",
					"wss://cyber-ws.alt.technology/",
					"https://rpc.cyber.co/",
					"wss://rpc.cyber.co/"
				]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/cyber"
		}
	},
	degen: {
		type: "chains",
		id: "degen",
		attributes: {
			external_id: "0x27bc86aa",
			name: "Degen Chain",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/666666666"
			},
			explorer: {
				name: "Degen Explorer",
				token_url_format: "https://explorer.degen.tips/token/{ADDRESS}",
				tx_url_format: "https://explorer.degen.tips/tx/{HASH}",
				home_url: "https://explorer.degen.tips"
			},
			rpc: {
				public_servers_url: ["https://rpc.degen.tips"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/d590ac9c-6971-42db-b900-0bd057033ae0"
				},
				data: {
					type: "fungibles",
					id: "d590ac9c-6971-42db-b900-0bd057033ae0"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/06d135ee-95f3-489d-a0a3-70129d9f952c"
				},
				data: {
					type: "fungibles",
					id: "06d135ee-95f3-489d-a0a3-70129d9f952c"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/degen"
		}
	},
	fraxtal: {
		type: "chains",
		id: "fraxtal",
		attributes: {
			external_id: "0xfc",
			name: "Fraxtal",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/fraxtal.png"
			},
			explorer: {
				name: "Fraxscan",
				token_url_format: "https://fraxscan.com/token/{ADDRESS}",
				tx_url_format: "https://fraxscan.com/tx/{HASH}",
				home_url: "https://fraxscan.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.frax.com"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/50e9e0d9-d591-4083-977a-48921057b02d"
				},
				data: {
					type: "fungibles",
					id: "50e9e0d9-d591-4083-977a-48921057b02d"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/fraxtal"
		}
	},
	"gravity-alpha": {
		type: "chains",
		id: "gravity-alpha",
		attributes: {
			external_id: "0x659",
			name: "Gravity Alpha",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/gravity.png"
			},
			explorer: {
				name: "Gravity Alpha Mainnet Explorer",
				token_url_format: "https://explorer.gravity.xyz/token/{ADDRESS}",
				tx_url_format: "https://explorer.gravity.xyz/tx/{HASH}",
				home_url: "https://explorer.gravity.xyz"
			},
			rpc: {
				public_servers_url: ["https://rpc.gravity.xyz"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/21d6737a-ebfe-44f2-99b7-78e8282dc9ef"
				},
				data: {
					type: "fungibles",
					id: "21d6737a-ebfe-44f2-99b7-78e8282dc9ef"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/071102b7-cab5-4e48-836d-9b2375b91794"
				},
				data: {
					type: "fungibles",
					id: "071102b7-cab5-4e48-836d-9b2375b91794"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/gravity-alpha"
		}
	},
	ink: {
		type: "chains",
		id: "ink",
		attributes: {
			external_id: "0xdef1",
			name: "Ink",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/ink.png"
			},
			explorer: {
				name: "Ink Explorer",
				token_url_format: "https://explorer.inkonchain.com/token/{ADDRESS}",
				tx_url_format: "https://explorer.inkonchain.com/tx/{HASH}",
				home_url: "https://explorer.inkonchain.com"
			},
			rpc: {
				public_servers_url: ["https://rpc-qnd.inkonchain.com"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/ink"
		}
	},
	lens: {
		type: "chains",
		id: "lens",
		attributes: {
			external_id: "0xe8",
			name: "Lens",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/lens.png"
			},
			explorer: {
				name: "Lens Explorer",
				token_url_format: "https://explorer.lens.xyz/token/{ADDRESS}",
				tx_url_format: "https://explorer.lens.xyz/tx/{HASH}",
				home_url: "https://explorer.lens.xyz"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/lens"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x40d16fc0246ad3160ccc09b8d0d3a2cd28ae6c2f"
				},
				data: {
					type: "fungibles",
					id: "0x40d16fc0246ad3160ccc09b8d0d3a2cd28ae6c2f"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/1e6d7e24-70a8-4b3c-8150-8baf83b8cbad"
				},
				data: {
					type: "fungibles",
					id: "1e6d7e24-70a8-4b3c-8150-8baf83b8cbad"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/lens"
		}
	},
	lisk: {
		type: "chains",
		id: "lisk",
		attributes: {
			external_id: "0x46f",
			name: "Lisk",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/lisk.png"
			},
			explorer: {
				name: "Lisk",
				token_url_format: "https://blockscout.lisk.com//token/{ADDRESS}",
				tx_url_format: "https://blockscout.lisk.com//tx/{HASH}",
				home_url: "https://blockscout.lisk.com/"
			},
			rpc: {
				public_servers_url: ["https://rpc.api.lisk.com"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/lisk"
		}
	},
	mode: {
		type: "chains",
		id: "mode",
		attributes: {
			external_id: "0x868b",
			name: "Mode",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/mode.png"
			},
			explorer: {
				name: "Blockscout",
				token_url_format: "https://explorer.mode.network/token/{ADDRESS}",
				tx_url_format: "https://explorer.mode.network/tx/{HASH}",
				home_url: "https://explorer.mode.network"
			},
			rpc: {
				public_servers_url: ["https://mainnet.mode.network", "https://mode.drpc.org", "wss://mode.drpc.org"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/mode"
		}
	},
	mantle: {
		type: "chains",
		id: "mantle",
		attributes: {
			external_id: "0x1388",
			name: "Mantle",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/mantle.png"
			},
			explorer: {
				name: "Mantle Explorer",
				token_url_format: "https://explorer.mantle.xyz/token/{ADDRESS}",
				tx_url_format: "https://explorer.mantle.xyz/tx/{HASH}",
				home_url: "https://explorer.mantle.xyz"
			},
			rpc: {
				public_servers_url: [
					"https://rpc.mantle.xyz",
					"https://mantle-rpc.publicnode.com",
					"wss://mantle-rpc.publicnode.com"
				]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/f8e50e85-dc0b-4820-a1d8-1f98db6e60f8"
				},
				data: {
					type: "fungibles",
					id: "f8e50e85-dc0b-4820-a1d8-1f98db6e60f8"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/mantle"
		}
	},
	optimism: {
		type: "chains",
		id: "optimism",
		attributes: {
			external_id: "0xa",
			name: "Optimism",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/optimism.png"
			},
			explorer: {
				name: "Etherscan",
				token_url_format: "https://optimistic.etherscan.io/token/{ADDRESS}",
				tx_url_format: "https://optimistic.etherscan.io/tx/{HASH}",
				home_url: "https://optimistic.etherscan.io"
			},
			rpc: {
				public_servers_url: ["https://mainnet.optimism.io"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/optimism"
		}
	},
	polygon: {
		type: "chains",
		id: "polygon",
		attributes: {
			external_id: "0x89",
			name: "Polygon",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/polygon.png"
			},
			explorer: {
				name: "PolygonScan",
				token_url_format: "https://polygonscan.com/token/{ADDRESS}",
				tx_url_format: "https://polygonscan.com/tx/{HASH}",
				home_url: "https://polygonscan.com"
			},
			rpc: {
				public_servers_url: [
					"https://polygon-rpc.com",
					"https://rpc-mainnet.matic.network",
					"https://matic-mainnet.chainstacklabs.com",
					"https://rpc-mainnet.maticvigil.com",
					"https://rpc-mainnet.matic.quiknode.pro",
					"https://matic-mainnet-full-rpc.bwarelabs.com"
				]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/7560001f-9b6d-4115-b14a-6c44c4334ef2"
				},
				data: {
					type: "fungibles",
					id: "7560001f-9b6d-4115-b14a-6c44c4334ef2"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/ef4dfcc9-4a7e-4a92-a538-df3d6f53e517"
				},
				data: {
					type: "fungibles",
					id: "ef4dfcc9-4a7e-4a92-a538-df3d6f53e517"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/polygon"
		}
	},
	celo: {
		type: "chains",
		id: "celo",
		attributes: {
			external_id: "0xa4ec",
			name: "Celo",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/42220"
			},
			explorer: {
				name: "Celoscan",
				token_url_format: "https://celoscan.io/token/{ADDRESS}",
				tx_url_format: "https://celoscan.io/tx/{HASH}",
				home_url: "https://celoscan.io"
			},
			rpc: {
				public_servers_url: ["https://forno.celo.org", "wss://forno.celo.org/ws"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x471ece3750da237f93b8e339c536989b8978a438"
				},
				data: {
					type: "fungibles",
					id: "0x471ece3750da237f93b8e339c536989b8978a438"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/celo"
		}
	},
	"manta-pacific": {
		type: "chains",
		id: "manta-pacific",
		attributes: {
			external_id: "0xa9",
			name: "Manta Pacific",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/manta.png"
			},
			explorer: {
				name: "manta-pacific Explorer",
				token_url_format: "https://pacific-explorer.manta.network/token/{ADDRESS}",
				tx_url_format: "https://pacific-explorer.manta.network/tx/{HASH}",
				home_url: "https://pacific-explorer.manta.network"
			},
			rpc: {
				public_servers_url: [
					"https://pacific-rpc.manta.network/http",
					"https://manta-pacific.drpc.org",
					"wss://manta-pacific.drpc.org"
				]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/manta-pacific"
		}
	},
	xdai: {
		type: "chains",
		id: "xdai",
		attributes: {
			external_id: "0x64",
			name: "Gnosis Chain",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/xdai.png"
			},
			explorer: {
				name: "GnosisScan",
				token_url_format: "https://gnosisscan.io/token/{ADDRESS}",
				tx_url_format: "https://gnosisscan.io/tx/{HASH}",
				home_url: "https://gnosisscan.io"
			},
			rpc: {
				public_servers_url: [
					"https://gnosis-mainnet.public.blastapi.io",
					"https://gnosischain-rpc.gateway.pokt.network",
					"https://rpc.ankr.com/gnosis",
					"https://rpc.gnosischain.com"
				]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/b99ea659-0ab1-4832-bf44-3bf1cc1acac7"
				},
				data: {
					type: "fungibles",
					id: "b99ea659-0ab1-4832-bf44-3bf1cc1acac7"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xe91d153e0b41518a2ce8dd3d7944fa863463a97d"
				},
				data: {
					type: "fungibles",
					id: "0xe91d153e0b41518a2ce8dd3d7944fa863463a97d"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/xdai"
		}
	},
	fantom: {
		type: "chains",
		id: "fantom",
		attributes: {
			external_id: "0xfa",
			name: "Fantom",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/fantom.png"
			},
			explorer: {
				name: "FtmScan",
				token_url_format: "https://ftmscan.com/token/{ADDRESS}",
				tx_url_format: "https://ftmscan.com/tx/{HASH}",
				home_url: "https://ftmscan.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.ftm.tools"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x4e15361fd6b4bb609fa63c81a2be19d873717870"
				},
				data: {
					type: "fungibles",
					id: "0x4e15361fd6b4bb609fa63c81a2be19d873717870"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83"
				},
				data: {
					type: "fungibles",
					id: "0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/fantom"
		}
	},
	ronin: {
		type: "chains",
		id: "ronin",
		attributes: {
			external_id: "0x7e4",
			name: "Ronin",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/2020"
			},
			explorer: {
				name: "Ronin Block Explorer",
				token_url_format: "https://app.roninchain.com/token/{ADDRESS}",
				tx_url_format: "https://app.roninchain.com/tx/{HASH}",
				home_url: "https://app.roninchain.com"
			},
			rpc: {
				public_servers_url: ["https://api.roninchain.com/rpc"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/3e1f750c-aff1-4918-8a18-4e71f28ffa47"
				},
				data: {
					type: "fungibles",
					id: "3e1f750c-aff1-4918-8a18-4e71f28ffa47"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/6dd2b90b-c236-498d-b2f4-f1098c56ddeb"
				},
				data: {
					type: "fungibles",
					id: "6dd2b90b-c236-498d-b2f4-f1098c56ddeb"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/ronin"
		}
	},
	aurora: {
		type: "chains",
		id: "aurora",
		attributes: {
			external_id: "0x4e454152",
			name: "Aurora",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/aurora.png"
			},
			explorer: {
				name: "Aurora Explorer",
				token_url_format: "https://explorer.aurora.dev/token/{ADDRESS}",
				tx_url_format: "https://explorer.aurora.dev/tx/{HASH}",
				home_url: "https://explorer.aurora.dev"
			},
			rpc: {
				public_servers_url: ["https://mainnet.aurora.dev"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/aurora"
		}
	},
	linea: {
		type: "chains",
		id: "linea",
		attributes: {
			external_id: "0xe708",
			name: "Linea",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/59144"
			},
			explorer: {
				name: "Etherscan",
				token_url_format: "https://lineascan.build/token/{ADDRESS}",
				tx_url_format: "https://lineascan.build/tx/{HASH}",
				home_url: "https://lineascan.build"
			},
			rpc: {
				public_servers_url: ["https://rpc.linea.build", "wss://rpc.linea.build"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/linea"
		}
	},
	"metis-andromeda": {
		type: "chains",
		id: "metis-andromeda",
		attributes: {
			external_id: "0x440",
			name: "Metis Andromeda",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/metis.png"
			},
			explorer: {
				name: "blockscout",
				token_url_format: "https://andromeda-explorer.metis.io/token/{ADDRESS}",
				tx_url_format: "https://andromeda-explorer.metis.io/tx/{HASH}",
				home_url: "https://andromeda-explorer.metis.io"
			},
			rpc: {
				public_servers_url: [
					"https://andromeda.metis.io/?owner=1088",
					"https://metis.drpc.org",
					"wss://metis.drpc.org"
				]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x9e32b13ce7f2e80a01932b42553652e053d6ed8e"
				},
				data: {
					type: "fungibles",
					id: "0x9e32b13ce7f2e80a01932b42553652e053d6ed8e"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/metis-andromeda"
		}
	},
	ape: {
		type: "chains",
		id: "ape",
		attributes: {
			external_id: "0x8173",
			name: "Ape Chain",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/apechain.png"
			},
			explorer: {
				name: "Apechain Explorer",
				token_url_format: "https://apechain.calderaexplorer.xyz/token/{ADDRESS}",
				tx_url_format: "https://apechain.calderaexplorer.xyz/tx/{HASH}",
				home_url: "https://apechain.calderaexplorer.xyz/"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/ape"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x4d224452801aced8b2f0aebe155379bb5d594381"
				},
				data: {
					type: "fungibles",
					id: "0x4d224452801aced8b2f0aebe155379bb5d594381"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/ape"
		}
	},
	"cronos-zkevm": {
		type: "chains",
		id: "cronos-zkevm",
		attributes: {
			external_id: "0x184",
			name: "Cronos zkEVM",
			icon: {
				url: "https://protocol-icons.s3.amazonaws.com/icons/cronos-zkevm.png"
			},
			explorer: {
				name: "Cronos zkEVM",
				token_url_format: "https://explorer.zkevm.cronos.org//token/{ADDRESS}",
				tx_url_format: "https://explorer.zkevm.cronos.org//tx/{HASH}",
				home_url: "https://explorer.zkevm.cronos.org/"
			},
			rpc: {
				public_servers_url: ["https://mainnet.zkevm.cronos.org"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/56f2a7d9-87e1-47da-90ab-ea0b5ae004c3"
				},
				data: {
					type: "fungibles",
					id: "56f2a7d9-87e1-47da-90ab-ea0b5ae004c3"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/cronos-zkevm"
		}
	},
	"polygon-zkevm": {
		type: "chains",
		id: "polygon-zkevm",
		attributes: {
			external_id: "0x44d",
			name: "Polygon zkEVM",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/1101"
			},
			explorer: {
				name: "blockscout",
				token_url_format: "https://zkevm.polygonscan.com/token/{ADDRESS}",
				tx_url_format: "https://zkevm.polygonscan.com/tx/{HASH}",
				home_url: "https://zkevm.polygonscan.com"
			},
			rpc: {
				public_servers_url: ["https://zkevm-rpc.com"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/polygon-zkevm"
		}
	},
	polynomial: {
		type: "chains",
		id: "polynomial",
		attributes: {
			external_id: "0x1f48",
			name: "Polynomial",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/polynomial.png"
			},
			explorer: {
				name: "PolynomialExplorer",
				token_url_format: "https://polynomialscan.io//token/{ADDRESS}",
				tx_url_format: "https://polynomialscan.io//tx/{HASH}",
				home_url: "https://polynomialscan.io/"
			},
			rpc: {
				public_servers_url: ["https://rpc.polynomial.fi"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/polynomial"
		}
	},
	rari: {
		type: "chains",
		id: "rari",
		attributes: {
			external_id: "0x52415249",
			name: "Rari",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/1380012617"
			},
			explorer: {
				name: "rarichain-explorer",
				token_url_format: "https://mainnet.explorer.rarichain.org/token/{ADDRESS}",
				tx_url_format: "https://mainnet.explorer.rarichain.org/tx/{HASH}",
				home_url: "https://mainnet.explorer.rarichain.org"
			},
			rpc: {
				public_servers_url: ["https://rari.calderachain.xyz/http"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/rari"
		}
	},
	redstone: {
		type: "chains",
		id: "redstone",
		attributes: {
			external_id: "0x2b2",
			name: "Redstone",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/redstone.png"
			},
			explorer: {
				name: "blockscout",
				token_url_format: "https://explorer.redstone.xyz/token/{ADDRESS}",
				tx_url_format: "https://explorer.redstone.xyz/tx/{HASH}",
				home_url: "https://explorer.redstone.xyz"
			},
			rpc: {
				public_servers_url: ["https://rpc.redstonechain.com", "wss://rpc.redstonechain.com"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/redstone"
		}
	},
	scroll: {
		type: "chains",
		id: "scroll",
		attributes: {
			external_id: "0x82750",
			name: "Scroll",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/scroll.png"
			},
			explorer: {
				name: "Scrollscan",
				token_url_format: "https://scrollscan.com/token/{ADDRESS}",
				tx_url_format: "https://scrollscan.com/tx/{HASH}",
				home_url: "https://scrollscan.com"
			},
			rpc: {
				public_servers_url: [
					"https://rpc.scroll.io",
					"https://rpc-scroll.icecreamswap.com",
					"https://rpc.ankr.com/scroll",
					"https://scroll-mainnet.chainstacklabs.com"
				]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/scroll"
		}
	},
	sei: {
		type: "chains",
		id: "sei",
		attributes: {
			external_id: "0x531",
			name: "Sei",
			icon: {
				url: "https://protocol-icons.s3.amazonaws.com/icons/sei.png"
			},
			explorer: {
				name: "Sei Explorer",
				token_url_format: "https://seistream.app/token/{ADDRESS}",
				tx_url_format: "https://seistream.app/tx/{HASH}",
				home_url: "https://seistream.app"
			},
			rpc: {
				public_servers_url: ["https://evm-rpc.sei-apis.com"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/d1706cbf-d14c-4065-8f8e-cf428dd5bd98"
				},
				data: {
					type: "fungibles",
					id: "d1706cbf-d14c-4065-8f8e-cf428dd5bd98"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/sei"
		}
	},
	soneium: {
		type: "chains",
		id: "soneium",
		attributes: {
			external_id: "0x74c",
			name: "Soneium",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/soneium.png"
			},
			explorer: {
				name: "Soneium Explorer",
				token_url_format: "https://soneium.blockscout.com/token/{ADDRESS}",
				tx_url_format: "https://soneium.blockscout.com/tx/{HASH}",
				home_url: "https://soneium.blockscout.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/soneium"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/soneium"
		}
	},
	sonic: {
		type: "chains",
		id: "sonic",
		attributes: {
			external_id: "0x92",
			name: "Sonic",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/sonic_s.png"
			},
			explorer: {
				name: "Sonic Explorer",
				token_url_format: "https://sonicscan.org/token/{ADDRESS}",
				tx_url_format: "https://sonicscan.org/tx/{HASH}",
				home_url: "https://sonicscan.org"
			},
			rpc: {
				public_servers_url: ["https://rpc.soniclabs.com"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/ee334614-0375-469e-b699-e2df1bf4185d"
				},
				data: {
					type: "fungibles",
					id: "ee334614-0375-469e-b699-e2df1bf4185d"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/d22aba16-16b9-4f1d-943b-955e7a8f92cd"
				},
				data: {
					type: "fungibles",
					id: "d22aba16-16b9-4f1d-943b-955e7a8f92cd"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/sonic"
		}
	},
	swellchain: {
		type: "chains",
		id: "swellchain",
		attributes: {
			external_id: "0x783",
			name: "Swellchain",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/swellchain.png"
			},
			explorer: {
				name: "Swell Explorer",
				token_url_format: "https://explorer.swellnetwork.io//token/{ADDRESS}",
				tx_url_format: "https://explorer.swellnetwork.io//tx/{HASH}",
				home_url: "https://explorer.swellnetwork.io/"
			},
			rpc: {
				public_servers_url: ["https://swell-mainnet.alt.technology", "https://rpc.ankr.com/swell"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/swellchain"
		}
	},
	taiko: {
		type: "chains",
		id: "taiko",
		attributes: {
			external_id: "0x28c58",
			name: "Taiko",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/taiko.png"
			},
			explorer: {
				name: "Taiko Explorer",
				token_url_format: "https://taikoscan.io/token/{ADDRESS}",
				tx_url_format: "https://taikoscan.io/tx/{HASH}",
				home_url: "https://taikoscan.io/"
			},
			rpc: {
				public_servers_url: ["https://rpc.taiko.xyz", "wss://ws.taiko.xyz"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/taiko"
		}
	},
	tomochain: {
		type: "chains",
		id: "tomochain",
		attributes: {
			external_id: "0x58",
			name: "Viction",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/viction.png"
			},
			explorer: {
				name: "Vicscan",
				token_url_format: "https://www.vicscan.xyz//token/{ADDRESS}",
				tx_url_format: "https://www.vicscan.xyz//tx/{HASH}",
				home_url: "https://www.vicscan.xyz/"
			},
			rpc: {
				public_servers_url: ["https://rpc.viction.xyz"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/2912694c-c696-4549-acfb-f43a87938623"
				},
				data: {
					type: "fungibles",
					id: "2912694c-c696-4549-acfb-f43a87938623"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/tomochain"
		}
	},
	unichain: {
		type: "chains",
		id: "unichain",
		attributes: {
			external_id: "0x82",
			name: "Unichain",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/unichain.png"
			},
			explorer: {
				name: "Unichain Explorer",
				token_url_format: "https://unichain.blockscout.com/token/{ADDRESS}",
				tx_url_format: "https://unichain.blockscout.com/tx/{HASH}",
				home_url: "https://unichain.blockscout.com"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/unichain"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/unichain"
		}
	},
	okbchain: {
		type: "chains",
		id: "okbchain",
		attributes: {
			external_id: "0xc4",
			name: "X Layer",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/okx.png"
			},
			explorer: {
				name: "OKLink",
				token_url_format: "https://www.oklink.com/xlayer/token/{ADDRESS}",
				tx_url_format: "https://www.oklink.com/xlayer/tx/{HASH}",
				home_url: "https://www.oklink.com/xlayer"
			},
			rpc: {
				public_servers_url: ["https://rpc.xlayer.tech", "https://xlayerrpc.okx.com"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0x75231f58b43240c9718dd58b4967c5114342a86c"
				},
				data: {
					type: "fungibles",
					id: "0x75231f58b43240c9718dd58b4967c5114342a86c"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/8ecdd365-3776-4ec2-803b-194862b161ba"
				},
				data: {
					type: "fungibles",
					id: "8ecdd365-3776-4ec2-803b-194862b161ba"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/okbchain"
		}
	},
	world: {
		type: "chains",
		id: "world",
		attributes: {
			external_id: "0x1e0",
			name: "World Chain",
			icon: {
				url: "https://protocol-icons.s3.amazonaws.com/worldchain.png"
			},
			explorer: {
				name: "WorldScan",
				token_url_format: "https://worldscan.org/token/{ADDRESS}",
				tx_url_format: "https://worldscan.org/tx/{HASH}",
				home_url: "https://worldscan.org"
			},
			rpc: {
				public_servers_url: ["https://worldchain-mainnet.g.alchemy.com/public"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/world"
		}
	},
	zero: {
		type: "chains",
		id: "zero",
		attributes: {
			external_id: "0x849ea",
			name: "ZERϴ",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/chainlist/543210"
			},
			explorer: {
				name: "ZERϴ Explorer",
				token_url_format: "https://explorer.zero.network/token/{ADDRESS}",
				tx_url_format: "https://explorer.zero.network/tx/{HASH}",
				home_url: "https://explorer.zero.network"
			},
			rpc: {
				public_servers_url: ["https://rpc.zerion.io/v1/zero"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: true
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/zero"
		}
	},
	"zklink-nova": {
		type: "chains",
		id: "zklink-nova",
		attributes: {
			external_id: "0xc5cc4",
			name: "ZkLink Nova",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/zklink.png"
			},
			explorer: {
				name: "zkLink Nova Block Explorer",
				token_url_format: "https://explorer.zklink.io/token/{ADDRESS}",
				tx_url_format: "https://explorer.zklink.io/tx/{HASH}",
				home_url: "https://explorer.zklink.io"
			},
			rpc: {
				public_servers_url: ["https://rpc.zklink.io", "wss://rpc.zklink.io"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/zklink-nova"
		}
	},
	"re-al": {
		type: "chains",
		id: "re-al",
		attributes: {
			external_id: "0x1b254",
			name: "re.al",
			icon: {
				url: "https://chain-icons.s3.us-east-1.amazonaws.com/real.png"
			},
			explorer: {
				name: "blockscout",
				token_url_format: "https://explorer.re.al/token/{ADDRESS}",
				tx_url_format: "https://explorer.re.al/tx/{HASH}",
				home_url: "https://explorer.re.al"
			},
			rpc: {
				public_servers_url: ["https://real.drpc.org", "wss://real.drpc.org"]
			},
			flags: {
				supports_trading: false,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/3b5814a4-7fa3-48e3-b389-638925def2cf"
				},
				data: {
					type: "fungibles",
					id: "3b5814a4-7fa3-48e3-b389-638925def2cf"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/re-al"
		}
	},
	zora: {
		type: "chains",
		id: "zora",
		attributes: {
			external_id: "0x76adf1",
			name: "Zora",
			icon: {
				url: "https://chain-icons.s3.amazonaws.com/zora"
			},
			explorer: {
				name: "Zora Network Explorer",
				token_url_format: "https://explorer.zora.energy/token/{ADDRESS}",
				tx_url_format: "https://explorer.zora.energy/tx/{HASH}",
				home_url: "https://explorer.zora.energy"
			},
			rpc: {
				public_servers_url: ["https://rpc.zora.energy/"]
			},
			flags: {
				supports_trading: true,
				supports_sending: true,
				supports_bridge: false
			}
		},
		relationships: {
			native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/eth"
				},
				data: {
					type: "fungibles",
					id: "eth"
				}
			},
			wrapped_native_fungible: {
				links: {
					related: "https://api.zerion.io/v1/fungibles/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				},
				data: {
					type: "fungibles",
					id: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
				}
			}
		},
		links: {
			self: "https://api.zerion.io/v1/chains/zora"
		}
	}
}
