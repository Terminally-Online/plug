/** @type {import('next').NextConfig} */

const TsconfigPathsPlugin = require("tsconfig-paths-webpack-plugin")

const remotePatterns = [
	{
		protocol: "https",
		hostname: "assets.smold.app"
	},
	{
		protocol: "https",
		hostname: "ipfs.io"
	},
	{
		protocol: "https",
		hostname: "**"
	}
]

const headers = [
	{
		key: "X-Frame-Options",
		value: "SAMEORIGIN"
	},
	{
		key: "X-XSS-Protection",
		value: "1; mode=block"
	},
	{
		key: "X-Content-Type-Options",
		value: "nosniff"
	},
	{
		key: "Referrer-Policy",
		value: "strict-origin-when-cross-origin"
	},
	{
		key: "Content-Security-Policy",
		value: "frame-ancestors 'none'"
	}
]

const nextConfig = {
	poweredByHeader: false,
	trailingSlash: true,
	images: {
		remotePatterns,
		dangerouslyAllowSVG: true,
		contentDispositionType: "attachment",
		contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
		formats: ["image/avif", "image/webp"]
	},
	/** We run eslint as a separate task in CI */
	eslint: { ignoreDuringBuilds: !!process.env.CI },
	headers: () => {
		return [
			{
				source: "/(.*)",
				headers
			}
		]
	},
	webpack: config => {
		config.externals.push("pino-pretty", "lokijs", "encoding", {
			"utf-8-validate": "commonjs utf-8-validate",
			bufferutil: "commonjs bufferutil"
		})
		config.resolve.plugins.push(new TsconfigPathsPlugin({}))

		return config
	},
	experimental: {
		optimizePackageImports: [
			// "./components",
			// "./lib",
			// "./state",
			// "./server",
			// "framer-motion",
			// "viem",
			// "viem/chains",
			// "viem/ens",
			// "viem/siwe",
			// "wagmi",
			// "lucide-react"
		]
	}
}

const withPWA = require('next-pwa')({
	dest: 'public',
	// disable: process.env.NODE_ENV === 'development'
})

module.exports = withPWA(nextConfig)