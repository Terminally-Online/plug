import { SingleBar } from "cli-progress"
import { z } from "zod"

// ! These fields will always be present on the exported object regardless of
//   whether they are present in the markdown file.
type MarkdownAttributes = Partial<
	Record<"title" | "description" | "image" | "slug" | "created" | "updated" | "tags" | "related" | "inbound", string>
>

// * Base markdown file extended with inferred or manually added fields.
type TypedAttributes = MarkdownAttributes &
	Record<"filename" | "content" | "className", string> &
	Record<"created", Date> &
	Partial<Record<"updated", Date>> &
	Partial<Record<"inbound" | "related" | "tags", string[]>> &
	Partial<Record<"unlisted", boolean>> &
	Record<string, string | string[] | Date>

const smallWords = [
	"a",
	"an",
	"the",
	"and",
	"but",
	"or",
	"for",
	"nor",
	"on",
	"at",
	"to",
	"from",
	"by",
	"in",
	"of",
	"up",
	"as",
	"it",
	"is",
	"has",
	"am",
	"are",
	"be",
	"can",
	"do",
	"did",
	"his",
	"her",
	"if",
	"its",
	"my",
	"not",
	"our",
	"out",
	"off",
	"so",
	"us",
	"was",
	"we",
	"you",
	"yet",
	"with",
	"vs",
	"vs."
]

const uppercaseWords = ["erc", "eip", "eth", "dao", "http"]

const schema = z.object({
	title: z.string({
		required_error: "Required field in attributes missing: .title"
	}),
	image: z.string({
		required_error: "Required field in attributes missing: .image"
	}),
	created: z.date({
		required_error: "Required field in attributes missing: .created"
	}),
	className: z.string().default("")
})

// * Parse the string attribute value into the correct type.
export const getAttribute = (key: keyof MarkdownAttributes, value: string) => {
	// ! We convert it to a Date here although it will be converted back to a string
	//   so that we can go from %Y-%m-%d to %Y-%m-%dT%H:%M:%S.%LZ for the front-end.
	if (["created", "updated"].includes(key)) {
		return new Date(value)
	}

	// ! Split the string-value tags into an array of tags.
	// * We do not include the inbound attribute here because it is
	//   automatically generated and should not be included in the front-end.
	if (["related", "tags"].includes(key)) {
		return value.split(",").map(item => item.trim())
	}

	return value.trim()
}

export const getAttributes = (attributesStr: string): TypedAttributes | undefined => {
	const attributes = attributesStr
		.split("\n")
		.filter(item => item !== "")
		.reduce((acc, item) => {
			// * Assemble the key-value pairs.
			// ! Only split on the first colon to allow for colons in the value.
			const [key, ...value] = item.split(":")

			const trimmedKey = key.trim() as keyof MarkdownAttributes

			// * Add the key-value pair to the accumulator.
			acc[key.trim()] = getAttribute(trimmedKey, value.join(":").trim())

			return acc
		}, {} as TypedAttributes)

	if (Object.keys(attributes).length === 0) return

	return attributes
}

// * Get the contents of the file paths provided.
export const getMarkdown = (files: string[], contents: string[], bar?: SingleBar) => {
	const base = files
		.filter(file => file.endsWith(".md"))
		.map((file, index) => {
			// * Get all the content within the header which is in the format of:
			//   ---
			//   key: value
			//   ---
			const [, attributesStr, content] = contents[index].split("---")
			const attributes = getAttributes(attributesStr)

			const filename = file.replace(".md", "")

			if (bar) {
				bar.update({ filename: filename.slice(0, 50) })
				bar.increment()
			}

			let { title, description, image, slug, created, className, ...attributesNested } = attributes ?? {}

			const validated = schema.safeParse({
				title,
				image,
				created: created,
				className: className
			})

			if (validated.success === false) {
				for (const issue in validated.error.issues)
					console.error("\x1b[31m%s\x1b[0m", `! ${filename}: ${validated.error.issues[issue].message}`)

				return
			}

			const words = (title ?? "").split(" ")
			const titleCasedWords = words.map((word, index) => {
				if (uppercaseWords.filter(item => word.toLowerCase().startsWith(item)).length > 0) {
					// ! If it ends in an 's', lowercase the 's'
					if (word.charAt(word.length - 1) === "s") {
						return word.slice(0, -1).toUpperCase() + word.slice(-1)
					}

					return word.toUpperCase()
				}

				if (index === 0 || index === words.length - 1 || !smallWords.includes(word)) {
					return word.charAt(0).toUpperCase() + word.slice(1)
				}

				return word
			})

			title = titleCasedWords.join(" ").replace(/([^.!?])$/, "$1.")

			description = `${(description ?? content)
				.slice(0, 180)
				// * Remove all new lines.
				.replace(/\n/g, " ")
				// * Remove all characters that are non-alphanumeric and not punctuation.
				.replace(/[^a-zA-Z0-9.,!\- ]/g, "")
				.trim()}...`

			return {
				filename,
				slug: slug ?? filename,
				title,
				description,
				image,
				content,
				attributes: {
					created: validated.data.created,
					className: validated.data.className,
					...attributesNested
				}
			}
		})

	const inbound = base.map(file => {
		// * Remove files that could not be compiled
		if (!file) return

		const inbound = base
			.filter(item => !item?.attributes?.unlisted && item?.attributes?.inbound)
			.reduce((acc, item) => {
				if (!file.filename || !item?.attributes?.related) return acc

				const isRelated =
					item?.attributes?.related?.includes(file.filename) || item?.attributes?.related?.includes(file.slug)

				if (isRelated) {
					acc.push(item.slug)
				}

				return acc
			}, [] as string[])

		if (inbound.length > 0)
			return {
				...file,
				attributes: {
					...file.attributes,
					inbound
				}
			}

		return file
	})

	return inbound
		.sort((a, b) => {
			// * Order chronologically from newest to oldest.
			if (!a?.attributes?.created || !b?.attributes?.created) return 0

			return b.attributes.created.getTime() - a.attributes.created.getTime()
		})
		.reduce(
			(acc, item) => {
				if (!item?.slug) return acc

				acc[item.slug.replaceAll("-", "")] = item

				return acc
			},
			{} as Record<string, (typeof inbound)[0]>
		)
}
