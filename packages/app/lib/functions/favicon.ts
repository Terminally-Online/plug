import axios from "axios"
import { load } from "cheerio"
import { SingleBar } from "cli-progress"

const DEBUG = process.env.NODE_ENV === "development"

// * Get the favicon URL and map it to each content URL.
const faviconUrls: Record<string, string> = {}

export const encodeFavicon = (mimeType: string, data: any) => {
	// * Confirm that the response is an image and not a rendered page.
	if (!mimeType.includes("image")) return

	return `data:${mimeType};base64,${Buffer.from(data, "binary").toString("base64")}`
}

export const getFaviconImage = async (url?: string) => {
	if (!url) return

	return await axios
		.get(`${url}`, {
			responseType: "arraybuffer"
		})
		.then(response => {
			return encodeFavicon(response.headers["content-type"], response.data)
		})
		// * If this fails, we will check the HTML. (Silently fail)
		.catch(() => undefined)
}

export const getFavicon = async (url: string) => {
	// * Attempt to recover the favicon URL from the HTML.
	const htmlFaviconResponse = await axios
		.get(url)
		.then(async htmlResponse => {
			const html = htmlResponse.data
			const $ = load(html)

			// * Recover the favicon from one of the many possible tags.
			const faviconTag =
				$('link[rel="icon"]').first() ||
				$('link[rel="shortcut icon"]').first() ||
				$('link[rel="apple-touch-icon"]').first()

			let faviconHref = faviconTag.attr("href")

			// ! The website used a relative path for the favicon URL.
			if (faviconHref?.startsWith("/")) {
				faviconHref = `${new URL(url).origin}${faviconHref}`

				if (DEBUG)
					console.log("\x1b[33m%s\x1b[0m", `! Relative to absolute path converstion for ${faviconHref}.`)
			}

			return await getFaviconImage(faviconHref)
		})
		.catch(e => {
			if (DEBUG && e.response?.status === 403) {
				const error = `! 403 Forbidden: Install a static reference for the favicon at ${url}.`

				console.log("\x1b[31m%s\x1b[0m", error)
			}

			return undefined
		})

	if (DEBUG) {
		if (htmlFaviconResponse !== undefined) console.log("\x1b[32m%s\x1b[0m", `✔︎ Found favicon in HTML for ${url}.`)
		else console.log("\x1b[31m%s\x1b[0m", `✘ No favicon in HTML for ${url}.`)
	}

	if (htmlFaviconResponse !== undefined) return htmlFaviconResponse

	let response = await getFaviconImage(`${new URL(url).origin}/favicon.ico`)

	if (response !== undefined) {
		if (DEBUG) console.log("\x1b[32m%s\x1b[0m", `✔︎ Found favicon in root for ${url}.`)

		return response
	}

	return response
}

export async function getFavicons(urls: string[], bar?: SingleBar) {
	await Promise.all(
		urls
			.map(url => url.replace("https://", "").replace("http://", ""))
			.sort()
			.map(async url => {
				// ! Don't return anything because we've already retrieved this url.
				if (faviconUrls[url] !== undefined) return

				const favicon = await getFavicon(`http://${url}`)

				if (bar) {
					bar.increment()
					bar.update({ filename: url.slice(0, 8) })
				}

				if (!favicon) return

				faviconUrls[url] = favicon
			})
	)

	return faviconUrls
}
