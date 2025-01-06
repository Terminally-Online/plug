import * as progress from "cli-progress"
import { readFileSync, writeFileSync } from "fs"
import { join } from "path"

import { getFavicons } from "@/lib/functions/favicon"
import { getFiles, getUrls } from "@/lib/functions/files"
import { getMarkdown } from "@/lib/functions/markdown"

// * Define the static directories of the markdown files.
const BASE_DIR = "../posts"
const PUBLISHED_DIR = `${BASE_DIR}/published`

const bar = new progress.MultiBar(
	{
		format: "{title} | {bar} | {filename} | {percentage}% || {value}/{total} Chunks"
	},
	progress.Presets.shades_classic
)

;(async () => {
	const postsPath = join(PUBLISHED_DIR)
	const files = getFiles(postsPath)
	const filesBar = bar.create(files.length, 0)

	filesBar.update({ title: "Raw File" })
	const raw = files.map(filename => {
		filesBar.update({ filename })
		filesBar.increment()

		return readFileSync(join(postsPath, filename), "utf8")
	})
	filesBar.update({ filename: "Complete" })
	filesBar.stop()

	const markdownBar = bar.create(files.length, 0)
	markdownBar.update({ title: "Markdown" })
	const markdown = getMarkdown(files, raw, markdownBar)

	markdownBar.update({ filename: "Complete" })
	markdownBar.stop()

	const contentUrls = getUrls(Object.values(markdown).map(item => item?.content ?? "")).sort()

	const faviconBar = bar.create(contentUrls.length, 0)
	faviconBar.update({ title: "Favicons" })

	const favicons = await getFavicons(contentUrls, faviconBar)

	faviconBar.update({ filename: "Complete" })
	faviconBar.stop()

	// *  Generate the code that is used to interface with the markdown files.
	const code = `
        export const PAGE_SIZE = 20;

        export interface Post {
            filename: string;
            slug: string;
            title: string;
            description: string;
            image: string;
            content: string;
            attributes: {
                created: string;
            } & Partial<{
                updated: string;
                tags: string[];
                related: string[];
                inbound: string[];
                author: string;
                // display settings
                imagePosition: 'top' | 'bottom';
                imagePadded: 'true' | 'false';
                className: string;
                variant: string;
                unlisted: 'true' | 'false';
                sidebar: 'show' | 'hide' | undefined;
            }>
        }

        export type Posts = Record<string, Post>;

        export const faviconUrls = ${JSON.stringify(favicons)} as const;

        export const posts: Posts = ${JSON.stringify(markdown).replace(/"([^"]+)":/g, "$1:")};

        // * Get all the Posts for a given page.
        export const getPosts = (
            page = 1,
            pageSize = PAGE_SIZE,
            filter?:
                | Partial<Record<'date' | 'tag' | 'search', Partial<string>>>
                | undefined
        ): {
            posts: Post[]
            count: number
            hasNext: boolean
            random: Post
        } => {
            // ! Filter the posts before paginating so that we can get a final count.
            const filteredPosts = Object.values(
                Object.values(posts ?? {}).reduce<Record<string, Post>>(
                    (acc, article) => {
                        if(article.attributes.unlisted) return acc

                        if (!filter) {
                            acc[article.slug] = article

                            return acc
                        }

                        const { tag, date, search } = filter

                        // * Check the lowercase tags against the lowercase tag parameter.
                        const matchesTag =
                            !tag ||
                            (article.attributes.tags?.some(
                                articleTag =>
                                    articleTag.toLowerCase() === tag.toLowerCase()
                            ) ??
                                false)

                        // ! Require the exact date match.
                        const matchesDate =
                            !date ||
                            ((article.attributes.created?.includes(date) ||
                                article.attributes.updated?.includes(date)) ??
                                false)

                        const matchesSearch =
                            !search ||
                            article.title
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.description
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.content
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            (article.attributes.tags?.some(articleTag =>
                                articleTag.toLowerCase().includes(search.toLowerCase())
                            ) ??
                                false)

                        if (matchesTag && matchesDate && matchesSearch) {
                            acc[article.slug] = article
                        }

                        return acc
                    },
                    {}
                )
            )

            const count = filteredPosts.length
            const keys = Object.keys(posts)

            return {
                posts: filteredPosts.slice(
                    (page - 1) * pageSize,
                    page * pageSize
                ),
                count,
                hasNext: count > page * pageSize,
                random: posts[keys[Math.floor(Math.random() * keys.length)]]
            }
        }

        // * Get a specific Post by the value of the slug parameter in each Post.
        // ! The parameter can be the dictionary key or the slug which is a slug on the Post.
        export type PostLookupKey = keyof typeof posts
        export type PostLookup = (typeof posts)[PostLookupKey]['slug']

        export const getPost = (lookup: PostLookup) => {
            const article = posts[lookup.replaceAll('-', '') as PostLookupKey]

            if (!article) throw new Error('Post not found')

            return article
        }

        // * Get the favicon for a given URL.
        export const getFavicon = (url?: string) => {
            if (!url) return null

            const faviconUrl =
                faviconUrls[
                    url
                        .replace('https://', '')
                        .replace('http://', '') as keyof typeof faviconUrls
                ]

            if (!faviconUrl) return null

            return faviconUrl
        }
        `

	// * Write the code to the file.
	writeFileSync(join("lib/functions/posts.ts"), code)

	bar.stop()

	console.log("\n✔︎ Markdown files compiled successfully.")
})()
