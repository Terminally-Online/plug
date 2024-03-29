/** @type {import('next').NextConfig} */

const TsconfigPathsPlugin = require("tsconfig-paths-webpack-plugin")

const nextConfig = {
	poweredByHeader: false,
	trailingSlash: true,
	images: {
		remotePatterns: [
			{
				protocol: "https",
				hostname: "assets.smold.app"
			}
		]
	},
	/** We run eslint as a separate task in CI */
	eslint: { ignoreDuringBuilds: !!process.env.CI },
	webpack: config => {
		config.externals.push("pino-pretty", "lokijs", "encoding")
		config.externals.push({
			"utf-8-validate": "commonjs utf-8-validate",
			bufferutil: "commonjs bufferutil"
		})

		config.resolve.plugins.push(new TsconfigPathsPlugin({}))

		return config
	}
}

module.exports = nextConfig
