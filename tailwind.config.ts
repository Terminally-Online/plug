import type { Config } from 'tailwindcss'

import plugin from 'tailwindcss/plugin'

const config: Config = {
	darkMode: 'class',
	content: [
		'./src/pages/**/*.{js,ts,jsx,tsx,mdx}',
		'./src/components/**/*.{js,ts,jsx,tsx,mdx}',
		'./src/app/**/*.{js,ts,jsx,tsx,mdx}'
	],
	theme: {
		extend: {
			colors: {
				black: '#121111',
				white: '#ffffff',
				red: '#FF4700',
				orange: '#FF8C00',
				yellow: '#FFFC00'
			},
			backgroundImage: {
				'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
				'gradient-conic':
					'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))'
			},
			textShadow: {
				xs: '0 1px 2px var(--tw-text-shadow-color)',
				sm: '0 1px 2px var(--tw-shadow-color)',
				DEFAULT: '0 2px 4px var(--tw-shadow-color)',
				md: '0 4px 8px var(--tw-shadow-color)',
				lg: '0 8px 16px var(--tw-shadow-color)',
				xl: '0 12px 24px var(--tw-shadow-color)',
				'2xl': '0 16px 32px var(--tw-shadow-color)',
				'3xl': '0 20px 40px var(--tw-shadow-color)',
				blur: '0 0 60px var(--tw-shadow-color)'
			}
		}
	},
	plugins: [
		plugin(function ({ matchUtilities, theme }) {
			matchUtilities(
				{
					'text-shadow': value => ({
						textShadow: value
					})
				},
				{ values: theme('textShadow') }
			)
		}),
		plugin(function ({ addVariant }) {
			addVariant('active', ['&.active']),
				addVariant('group-active', ['group.active &'])
		})
	]
}
export default config
