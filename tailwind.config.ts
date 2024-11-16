import type { Config } from "tailwindcss"

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["var(--font-satoshi)", "sans-serif"]
      },
      colors: {
        black: "#0E160E",
        background: "#FFFFFF",
        "grayscale-100": "#EBECEC",
        "plug-green": "#385842",
        "plug-yellow": "#D2F38A",
      }
    },
  },
  plugins: [],
}

export default config