/** @type {import('tailwindcss').Config} */
const plugin = require("tailwindcss/plugin")

module.exports = {
	darkMode: ["class"],
	content: [
		"./pages/**/*.{ts,tsx}",
		"./components/**/*.{ts,tsx}",
		"./app/**/*.{ts,tsx}",
		"./src/**/*.{ts,tsx}"
	],
	theme: {
		container: {
			center: true,
			padding: "2rem",
			screens: {
				"2xl": "1400px"
			}
		},
		extend: {
			fontFamily: { sans: ["var(--font-satoshi)", "sans-serif"] },

			colors: {
				background: "#FFFFFF",
				search: "#F8F8F8",

				'grayscale-0': "#f8f8f8",
				'grayscale-100': "#EBECEC",
				'grayscale-200': "#D6D8D8",
				'grayscale-300': "#C2C5C5",
				'grayscale-400': "#aeb2b1",
				'grayscale-500': "#999e9e",
				'grayscale-600': "#858b8b",
				'grayscale-700': "#717877",
				'grayscale-800': "#5d6564",
				'grayscale-900': "#485150",
				'grayscale-1000': "#343e3d",
				'grayscale-1100': "#1f2525",

				'plug-green': "#00E100",
				'plug-yellow': "#A3F700",
				'plug-red': "#FF0000",

				border: "hsl(var(--border))",
				input: "hsl(var(--input))",
				ring: "hsl(var(--ring))",
				background: "hsl(var(--background))",
				foreground: "hsl(var(--foreground))",

				primary: {
					DEFAULT: "hsl(var(--primary))",
					foreground: "hsl(var(--primary-foreground))"
				},
				secondary: {
					DEFAULT: "hsl(var(--secondary))",
					foreground: "hsl(var(--secondary-foreground))"
				},
				destructive: {
					DEFAULT: "hsl(var(--destructive))",
					foreground: "hsl(var(--destructive-foreground))"
				},
				muted: {
					DEFAULT: "hsl(var(--muted))",
					foreground: "hsl(var(--muted-foreground))"
				},
				accent: {
					DEFAULT: "hsl(var(--accent))",
					foreground: "hsl(var(--accent-foreground))"
				},
				popover: {
					DEFAULT: "hsl(var(--popover))",
					foreground: "hsl(var(--popover-foreground))"
				},
				card: {
					DEFAULT: "hsl(var(--card))",
					foreground: "hsl(var(--card-foreground))"
				}
			},

			backgroundImage: {
				'gradient-animated': 'linear-gradient(90deg, rgba(255,255,255,0.1), rgba(0,0,0,0.1), rgba(255,255,255,0.1))'
			},

			borderRadius: {
				lg: "16px",
				md: "10px",
				sm: "8px",
				xs: "4px"
			},

			keyframes: {
				"accordion-down": {
					from: { height: 0 },
					to: { height: "var(--radix-accordion-content-height)" }
				},
				"accordion-up": {
					from: { height: "var(--radix-accordion-content-height)" },
					to: { height: 0 }
				},
				"fade-in": {
					from: { opacity: 0 },
					to: { opacity: 1 },
				},
				"fade-out": {
					from: { opacity: 1 },
					to: { opacity: 0 },
				},
				"loading": {
					"0%": {
						"background-position": "0% 50%"
					},
					"50%": {
						"background-position": "100% 50%"
					},
					"100%": {
						"background-position": "0% 50%"
					}
				}
			},

			animation: {
				"accordion-down": "accordion-down 0.2s ease-out",
				"accordion-up": "accordion-up 0.2s ease-out",
				"fade-in": "fade-in 0.2s ease-out",
				"fade-out": "fade-out 0.2s ease-out",
				"loading": "loading 2s ease infinite"
			}
		}
	},
	plugins: [
		require("tailwindcss-animate"),
		plugin(function ({ addVariant }) {
			addVariant("active", ["&.active"])
			addVariant("group-active", ({ modifySelectors, separator }) => {
				modifySelectors(({ className }) => {
					return `.group-active${separator}${className}`
				})
			})
		})
	]
}
