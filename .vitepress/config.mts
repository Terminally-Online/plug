import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
	title: 'Emporium',
	description: 'Documentation for the Emporium protocol.',
	themeConfig: {
		// https://vitepress.dev/reference/default-theme-config
		nav: [
			{ text: 'Home', link: '/' },
			{ text: 'Examples', link: '/markdown-examples' }
		],

		sidebar: [
			{
				text: 'Introduction',
				items: [
					{
						text: 'Why Emporium',
						link: '/introduction/why-emporium'
					},
					{
						text: 'If This, Then That',
						link: '/introduction/if-this-then-that'
					},
					{
						text: 'Getting Started',
						link: '/introduction/getting-started'
					},
					{
						text: 'FAQ',
						link: '/introduction/frequently-asked-questions'
					}
				]
			},
			{
				text: 'Intents',
				collapsed: false,
				items: [
					{
						text: 'Introduction',
						link: '/intents/introduction',
						items: [
							{
								text: 'Imperative Transactions',
								link: '/intents/imperative-transactions'
							},
							{
								text: 'Declarative Messages',
								link: '/intents/declarative-messages'
							}
						]
					},
					{
						text: 'Execution Paths',
						link: '/intents/execution-paths',
						items: [
							{
								text: 'Single Lane',
								link: '/intents/execution-paths/single-lane'
							},
							{
								text: 'Multi-Dimensional',
								link: '/intents/execution-paths/multi-dimensional'
							},
							{
								text: 'Native Transactions',
								link: '/intents/execution-paths/native-transactions'
							},
							{
								text: 'Meta-Transactions',
								link: '/intents/execution-paths/meta-transactions'
							},
							{
								text: 'Channels',
								link: '/intents/execution-paths/channels'
							}
						]
					}
				]
			},
			{
				text: 'Types & Decoders',
				collapsed: true,
				items: [
					{
						text: 'EIP-712',
						link: '/decoders/eip-712',
						items: [
							{
								text: 'Signed Pairs',
								link: '/decoders/eip-712/signed-pairs'
							}
						]
					},
					{
						text: 'Automated Generation',
						link: '/decoders/automated-generation'
					},
					{
						text: 'Base Types',
						link: '/decoders/base-types',
						items: [
							{
								text: 'EIP712Domain',
								link: '/decoders/base-types/eip-712-domain'
							},
							{
								text: 'Caveat',
								link: '/decoders/base-types/caveat'
							},
							{
								text: 'Permissions',
								link: '/decoders/base-types/permissions'
							},
							{
								text: 'SignedPermissions',
								link: '/decoders/base-types/signed-permissions'
							},
							{
								text: 'Transaction',
								link: '/decoders/base-types/transaction'
							},
							{
								text: 'Intent',
								link: '/decoders/base-types/intent'
							},
							{
								text: 'SignedIntent',
								link: '/decoders/base-types/signed-intent'
							}
						]
					},
					{
						text: 'Packet Hash Getters',
						link: '/decoders/packet-hash-getters',
						items: [
							{
								text: 'getEIP712DomainHash',
								link: '/decoders/packet-hash-getters/get-eip-712-domain-packet-hash'
							},
							{
								text: 'getCaveatHash',
								link: '/decoders/packet-hash-getters/get-caveat-packet-hash'
							},
							{
								text: 'getPermissionsHash',
								link: '/decoders/packet-hash-getters/get-permissions-packet-hash'
							},
							{
								text: 'getSignedPermissionsHash',
								link: '/decoders/packet-hash-getters/get-signed-permissions-packet-hash'
							},
							{
								text: 'getTransactionHash',
								link: '/decoders/packet-hash-getters/get-transaction-packet-hash'
							},
							{
								text: 'getIntentHash',
								link: '/decoders/packet-hash-getters/get-intent-packet-hash'
							},
							{
								text: 'getSignedIntentHash',
								link: '/decoders/packet-hash-getters/get-signed-intent-packet-hash'
							}
						]
					},
					{
						text: 'Digest Getters',
						link: '/decoders/digest-getters',
						items: [
							{
								text: 'getPermissionsDigest',
								link: '/decoders/digest-getters/get-permissions-digest'
							},
							{
								text: 'getIntentDigest',
								link: '/decoders/digest-getters/get-intent-digest'
							}
						]
					},
					{
						text: 'Signer Getters',
						link: '/decoders/signer-getters',
						items: [
							{
								text: 'getSignedPermissionsSigner',
								link: '/decoders/signer-getters/get-signed-permissions-signer'
							},
							{
								text: 'getSignedIntentSigner',
								link: '/decoders/signer-getters/get-signed-intent-signer'
							}
						]
					}
				]
			},
			{
				text: 'Core Abstracts',
				collapsed: true,
				items: [
					{
						text: 'Enforcers',
						link: '/core/enforcers',
						items: [
							{
								text: 'ThresholdEnforcer',
								link: '/core/enforcers/threshold-enforcer'
							}
						]
					},
					{
						text: 'CaveatEnforcer',
						link: '/core/framework',
						items: [
							{
								text: 'enforceCaveat',
								link: '/core/enforcers/enforce-caveat'
							}
						]
					},
					{
						text: 'Framework',
						link: '/core/framework',
						items: [
							{
								text: 'contractInvoke',
								link: '/core/framework/contract-invoke'
							},
							{
								text: 'invoke',
								link: '/core/framework/invoke'
							}
						]
					}
				]
			},
			{
				text: 'Instances',
				collapsed: true,
				items: [
					{
						text: 'Enforcers',
						link: '/instances/enforcers',
						items: [
							{
								text: 'Allowed Methods',
								link: '/instances/enforcers/allowed-methods'
							},
							{
								text: 'Block Number',
								link: '/instances/enforcers/block-number'
							},
							{
								text: 'Timestamp',
								link: '/instances/enforcers/timestamp'
							},
							{
								text: 'Schedule Windows',
								link: '/instances/enforcers/schedule-windows'
							},
							{
								text: 'Limited Calls',
								link: '/instances/enforcers/limited-calls'
							},
							{
								text: 'ERC20Allowance',
								link: '/instances/enforcers/erc20-allowance'
							},
							{
								text: 'Revocation',
								link: '/instances/enforcers/revocation'
							}
						]
					}
				]
			}
		],

		socialLinks: [
			{ icon: 'github', link: 'https://github.com/nftchance/emporium' }
		],

		editLink: {
			pattern:
				'https://github.com/nftchance/emporium-docs/edit/master/:path'
		},

		search: {
			provider: 'local'
		}
	},

	lastUpdated: true,

	// * Load the font files.
	transformHead({ assets }) {
		// adjust the regex accordingly to match your font
		const myFontFile = assets.find(file => /GeistMonoVF\.\w+\.woff2/)
		if (myFontFile) {
			return [
				[
					'link',
					{
						rel: 'preload',
						href: myFontFile,
						as: 'font',
						type: 'font/woff2',
						crossorigin: ''
					}
				]
			]
		}
	},

	markdown: {
		lineNumbers: true
	}
})
