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
				items: [
					{ text: 'Introduction', link: '/intents/introduction' },
					{
						text: 'Imperative Transactions',
						link: '/intents/imperative-transactions'
					},
					{
						text: 'Declarative Messages',
						link: '/intents/declarative-messages'
					},
					{
						text: 'Permissions',
						link: '/intents/permissions',
						items: [
							{
								text: 'Counterfactual',
								link: '/intents/counterfactual'
							},
							{ text: 'Revocation', link: '/intents/revocation' },
							{ text: 'Caveats', link: '/intents/caveats' }
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
								text: 'Channels',
								link: '/intents/execution-paths/channels'
							},
							{
								text: 'Native Transactions',
								link: '/intents/execution-paths/native-transactions'
							},
							{
								text: 'Meta-Transactions',
								link: '/intents/execution-paths/meta-transactions'
							}
						]
					}
				]
			},
			{
				text: 'Types & Decoders',
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
					},
					{
						text: 'Automated Generation',
						link: '/decoders/automated-generation',
						items: [
							{
								text: 'Overloads',
								link: '/decoders/automated-generation/overloads'
							}
						]
					}
				]
			},
			{
				text: 'Core',
				items: [
					{
						text: 'Framework',
						link: '/core/framework',
						items: [
							{ text: 'invoke', link: '/core/framework/invoke' },
							{
								text: 'contractInvoke',
								link: '/core/framework/contract-invoke'
							}
						]
					},
					{ text: 'Abstracts', link: '/core/abstracts' },
					{
						text: 'Enforcers',
						link: '/core/enforcers',
						items: [
							{
								text: 'enforceCaveat',
								link: '/core/enforcers/enforce-caveat'
							}
						]
					},
					{ text: 'Executors', link: '/core/executors' }
				]
			}
		],

		socialLinks: [
			{ icon: 'github', link: 'https://github.com/nftchance/emporium' }
		],

		editLink: {
			pattern:
				'https://github.com/nftchance/emporium/edit/main/packages/docs/:path'
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
