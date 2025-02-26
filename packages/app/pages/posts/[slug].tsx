import { GetStaticProps, InferGetStaticPropsType } from "next"
import Image from "next/image"
import Link from "next/link"
import ReactMarkdown from "react-markdown"

import { motion } from "framer-motion"
import { ExternalLink } from "lucide-react"

import { postAnimations } from "@/components/blog/animations"
import { LandingContainer } from "@/components/landing/layout/container"
import { StaticLayout } from "@/components/landing/layout/static"
import { Counter } from "@/components/shared/utils/counter"
import { getFavicon, getPost, Post, PostLookup, posts } from "@/lib"

export const dynamicParams = false

export const getStaticProps = (async context => {
	const post = getPost(context.params?.slug as PostLookup)

	return {
		props: { post }
	}
}) satisfies GetStaticProps<{ post: Post }>

export async function getStaticPaths() {
	return {
		paths: Object.values(posts).map(post => ({ params: { slug: post.slug } })),
		fallback: false
	}
}

export default function Page({ post }: InferGetStaticPropsType<typeof getStaticProps>) {
	// Create a summary from content if no description
	const summaryText = post.description || post.content.substring(0, 160).replace(/[#*`]/g, '').trim();
	
	return (
		<StaticLayout 
			title={post.title}
			ogTitle={`Plug Posts | ${post.title}`}
			ogDescription={summaryText}
			ogImage={`${process.env.NEXT_PUBLIC_URL || ''}/api/canvas/post?name=${encodeURIComponent(post.title)}`}
		>
			<LandingContainer>
				<motion.article
					initial={{ opacity: 0, y: 20 }}
					animate={{ opacity: 1, y: 0 }}
					className="mx-auto flex max-w-3xl w-full overflow-x-hidden flex-col gap-4 pb-24 font-bold"
				>
					<motion.h1
						className="text-[48px] font-black leading-tight md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[1240px] xl:text-[96px]"
						initial={{ transform: "translateY(-20px)", opacity: 0 }}
						whileInView={{
							transform: ["translateY(-20px)", "translateY(0px)"],
							opacity: [0, 1]
						}}
						transition={{ duration: 0.3 }}
					>
						{post.title}
					</motion.h1>
					<div className="mb-4 flex flex-row items-center justify-between">
						<Link
							href={`https://twitter.com/${post.attributes.author}`}
							className="flex flex-row items-center gap-2"
						>
							<Image
								src={`/users/${post.attributes.author}.png`}
								alt={post.attributes.author ?? ""}
								width={24}
								height={24}
								className="rounded-full"
							/>
							<p>{post.attributes.author}</p>
						</Link>
						<p className="opacity-40">
							<Counter count={new Date(post.attributes.created).toLocaleDateString()} />
						</p>
					</div>

					{post.slug in postAnimations && <div className="min-h-[480px] overflow-hidden">
						{postAnimations[post.slug as keyof typeof postAnimations]}
					</div>}

					<ReactMarkdown
						components={{
							h1: ({ children }) => (
								<h1 className="text-[48px] md:text-[72px] lg:text-[96px] xl:text-[120px] font-black">{children}</h1>
							),
							h2: ({ children }) => (
								<h2 className="mt-8 text-[36px] md:text-[48px] lg:text-[60px] xl:text-[72px] font-black">
									{children}
								</h2>
							),
							h3: ({ children }) => (
								<h3 className="mt-8 text-[28px] md:text-[36px] lg:text-[48px] xl:text-[56px] font-black">
									{children}
								</h3>
							),
							h4: ({ children }) => (
								<h4 className="mt-8 text-[24px] md:text-[28px] lg:text-[36px] xl:text-[42px] font-black">
									{children}
								</h4>
							),
							h5: ({ children }) => (
								<h5 className="mt-8 text-[20px] md:text-[24px] lg:text-[28px] xl:text-[32px] font-black">
									{children}
								</h5>
							),
							h6: ({ children }) => (
								<h6 className="mt-8 text-[16px] md:text-[18px] lg:text-[20px] xl:text-[24px] font-black">
									{children}
								</h6>
							),
							p: ({ children }) => <p className="break-words opacity-80">{children}</p>,
							ul: ({ children }) => <ul className="list-disc break-words opacity-60">{children}</ul>,
							ol: ({ children }) => <ol className="list-decimal break-words opacity-60">{children}</ol>,
							li: ({ children }) => <li className="mb-2 list-item">{children}</li>,
							a: ({ children, ...props }) => {
								const faviconUrl = getFavicon(props.href ?? "")

								return (
									<span
										{...props}
										className="group inline cursor-pointer underline break-words transition-opacity duration-200 hover:opacity-80"
										onClick={e => {
											e.preventDefault()
											e.stopPropagation()
											if (props.href) {
												window.open(props.href, "_blank", "noopener,noreferrer")
											}
										}}
									>
										{faviconUrl && (
											<span className="relative mr-2 inline-block h-3 w-3 align-middle">
												<Image
													src={faviconUrl}
													alt="favicon"
													width={16}
													height={16}
													className="absolute bottom-0 left-0 right-0 top-0 transition-opacity duration-200 group-hover:opacity-0"
												/>
												<span className="absolute bottom-0 left-0 right-0 top-0 flex items-center justify-center opacity-0 transition-opacity duration-200 group-hover:opacity-100">
													<ExternalLink className="h-4 w-4 text-plug-green" />
												</span>
											</span>
										)}
										{children}
									</span>
								)
							},
							img: ({ src, alt }) => (
								<span className="group relative max-w-full overflow-hidden transition-all duration-200">
									<Image
										src={src ?? ""}
										alt={alt ?? ""}
										width={1920}
										height={1080}
										className="mx-auto max-h-[420px] w-full object-contain transition-all duration-200"
									/>
								</span>
							)
						}}
					>
						{post.content}
					</ReactMarkdown>
				</motion.article>
			</LandingContainer>
		</StaticLayout>
	)
}
