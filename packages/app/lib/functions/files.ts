import { readdirSync, statSync } from "fs"
import { join } from "path"

// * Recursively get all files in a directory.
export const getFiles = (dir: string, filesArray: string[] = []) => {
	const files = readdirSync(dir)

	files.forEach(file => {
		const fullPath = join(dir, file)

		// * If directory, execute a recursive call and get all the files inside.
		if (statSync(fullPath).isDirectory()) {
			filesArray = getFiles(fullPath, filesArray)
		} else {
			// * Push file to array.
			filesArray.push(file)
		}
	})

	return filesArray
}

export const getUrls = (contents: string[] = []) => {
	return (
		contents
			.flatMap(content => {
				const links = content.match(/\[([^\]]+)\]\(([^)]+)\)/g) ?? []

				return links
					.sort()
					.map(link => link.match(/\(([^)]+)\)/)?.[1] ?? "")
					.filter(url => {
						return !(
							!url ||
							!url.startsWith("http") ||
							url.startsWith("/") ||
							url.endsWith(".png") ||
							url.endsWith(".jpg") ||
							url.endsWith(".jpeg") ||
							url.endsWith(".gif")
						)
					})
			})
			// ! Remove duplicates.
			.filter((item, index, self) => self.indexOf(item) === index)
			.sort()
	)
}
