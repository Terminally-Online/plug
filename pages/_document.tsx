import { Head, Html, Main, NextScript } from "next/document"

export default function Document() {
	const title = "Plug"
	const description =
		"Plug brings trustless automation and scheduling to the entirety of the Ethereum ecosystem with an intent powered framework. The outcomes you want can finally be achieved without being glued to your device."

	return (
		<Html lang="en">
			<Head>
				{/* Favicon */}
				<link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png?v=2" />
				<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png?v=2" />
				<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png?v=2" />
				<link rel="manifest" href="/site.webmanifest?v=2" />
				<link rel="mask-icon" href="/safari-pinned-tab.svg?v=2" color="#00ef35" />
				<link rel="shortcut icon" href="/favicon.ico?v=2" />
				<meta name="msapplication-TileColor" content="#00ef35" />
				<meta name="theme-color" content="#ffffff" />

				{/* Basic Meta */}
				<meta charSet="utf-8" />
				<meta name="description" content={description} />

				{/* Open Graph */}
				<meta property="og:title" content={title} />
				<meta property="og:type" content="website" />
				<meta property="og:url" content="https://onplug.io" />
				<meta property="og:image" content="https://onplug.io/opengraph.png" />
				<meta property="og:description" content={description} />

				{/* Twitter */}
				<meta name="twitter:title" content={title} />
				<meta name="twitter:card" content="summary_large_image" />
				<meta name="twitter:site" content="@onplug_io" />
				<meta name="twitter:creator" content="@onplug_io" />
				<meta name="twitter:description" content={description} />
				<meta name="twitter:image" content="https://onplug.io/opengraph.png" />
			</Head>

			<body>
				<Main />
				<NextScript />
			</body>
		</Html>
	)
}
