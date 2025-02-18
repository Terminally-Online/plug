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
	return (
		<StaticLayout title={post.title}>
			<LandingContainer>
				<motion.article
					initial={{ opacity: 0, y: 20 }}
					animate={{ opacity: 1, y: 0 }}
					className="mx-auto flex max-w-3xl flex-col gap-4 pb-24 font-bold"
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
								<h1 className="text-[60px] font-black md:text-[80px] 2xl:text-[180px]">{children}</h1>
							),
							h2: ({ children }) => (
								<h2 className="mt-8 text-[48px] font-black md:text-[60px] 2xl:text-[60px]">
									{children}
								</h2>
							),
							h3: ({ children }) => (
								<h3 className="mt-8 text-[36px] font-black md:text-[48px] 2xl:text-[48px]">
									{children}
								</h3>
							),
							h4: ({ children }) => (
								<h4 className="mt-8 text-[24px] font-black md:text-[36px] 2xl:text-[36px]">
									{children}
								</h4>
							),
							h5: ({ children }) => (
								<h5 className="mt-8 text-[20px] font-black md:text-[24px] 2xl:text-[24px]">
									{children}
								</h5>
							),
							h6: ({ children }) => (
								<h6 className="mt-8 text-[16px] font-black md:text-[20px] 2xl:text-[18px]">
									{children}
								</h6>
							),
							p: ({ children }) => <p className="text-justify opacity-80">{children}</p>,
							ul: ({ children }) => <ul className="list-disc text-justify opacity-60">{children}</ul>,
							ol: ({ children }) => <ol className="list-decimal text-justify opacity-60">{children}</ol>,
							li: ({ children }) => <li className="mb-2 list-item">{children}</li>,
							a: ({ children, ...props }) => {
								const faviconUrl = getFavicon(props.href ?? "")

								return (
									<span
										{...props}
										className="group inline cursor-pointer underline transition-opacity duration-200 hover:opacity-80"
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
								<span className="group relative transition-all duration-200">
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
