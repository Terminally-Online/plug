/** @type {import('next').NextConfig} */

const TsconfigPathsPlugin = require("tsconfig-paths-webpack-plugin")

const nextConfig = {
	serverRuntimeConfig: {
		// Will only be available on the server side
	},
	publicRuntimeConfig: {
		// Will be available on both server and client
		APP_URL: process.env.NEXTAPP_URL,
		WS_URL: process.env.NEXTWS_URL
	},
	poweredByHeader: false,
	trailingSlash: true,
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
