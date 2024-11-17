import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["var(--font-satoshi)", "sans-serif"],
      },
      colors: {
        black: "#0E160E",
        white: "#FEFFF7",
        background: "#FEFFF7",

        "plug-green": "#385842",
        "plug-yellow": "#D2F38A",
        "plug-red": "#F38A8A",
      },
    },
  },
  plugins: [],
};

export default config;
