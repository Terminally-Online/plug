import { GetStaticProps, InferGetStaticPropsType } from "next"
import { useRouter } from "next/router"

import { motion } from "framer-motion"

import { postAnimations } from "@/components/blog/animations"
import { InfoCard } from "@/components/landing/cards/info"
import { LandingContainer } from "@/components/landing/layout/container"
import { StaticLayout } from "@/components/landing/layout/static"
import { Button } from "@/components/shared/buttons/button"
import { getPosts, Post, routes } from "@/lib"

export const getStaticProps = (async () => {
	const { posts } = getPosts(1, 100)
	return { props: { posts } }
}) satisfies GetStaticProps<{ posts: Post[] }>

const Page = ({ posts }: InferGetStaticPropsType<typeof getStaticProps>) => {
	const router = useRouter()

	return (
		<StaticLayout title="Posts">
			<LandingContainer>
				<div className="flex flex-col pb-32 pt-8">
					<div className="flex flex-col gap-4 md:my-auto xl:gap-8">
						<motion.h1
							className="text-[48px] font-black leading-tight md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[1080px] xl:text-[96px]"
							initial={{ transform: "translateY(-20px)", opacity: 0 }}
							whileInView={{
								transform: ["translateY(-20px)", "translateY(0px)"],
								opacity: [0, 1]
							}}
							transition={{ duration: 0.3 }}
						>
							All The Latest Posts of the Plug Ecosystem
						</motion.h1>

						<motion.p
							className="max-w-[480px] text-xl font-bold text-plug-green/40 md:max-w-[520px] lg:max-w-[620px] lg:text-[24px] xl:max-w-[620px] 2xl:max-w-[780px]"
							initial={{ transform: "translateY(20px)", opacity: 0 }}
							whileInView={{
								transform: ["translateY(20px)", "translateY(0px)"],
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.3,
								delay: 0.15
							}}
						>
							Learn about automation strategies, protocol integrations, and stay updated with the latest
							developments in the Plug ecosystem.
						</motion.p>

						<Button
							variant="none"
							className="mt-8 w-max rounded-md border-[1px] border-plug-yellow/30 bg-plug-yellow px-8 py-3 text-center font-black text-plug-green filter backdrop-blur-xl transition-all duration-200 ease-in-out hover:bg-plug-yellow/50"
							onClick={() => router.push(`/posts/${posts[0].slug}`)}
						>
							Read Latest
						</Button>
					</div>

					<div className="grid gap-2 pt-20 lg:grid-cols-2 2xl:grid-cols-3">
						{posts.map(post => (
							<InfoCard
								key={post.slug}
								text={post.title}
								author={post.attributes.author}
								description={`${post.description.slice(0, 140)}...`}
								href={`/posts/${post.slug}`}
								className="h-[480px] cursor-pointer"
							>
								{post.slug in postAnimations && (
									<div className="h-full w-full">
										{postAnimations[post.slug as keyof typeof postAnimations]}
									</div>
								)}

								<div className="absolute bottom-[45%] left-0 right-0 top-[25%] bg-gradient-to-b from-plug-white/0 to-plug-white" />
								<div className="absolute bottom-0 left-0 right-0 top-[55%] bg-plug-white" />
							</InfoCard>
						))}

						<motion.div
							className="flex h-[480px] cursor-pointer flex-col items-center justify-center gap-2 rounded-lg border-[1px] border-plug-green/10 bg-plug-green p-16 text-plug-yellow transition-all duration-200 ease-in-out hover:bg-plug-green hover:text-plug-yellow"
							initial={{ transform: "translateY(20px)", opacity: 0 }}
							whileInView={{
								transform: ["translateY(20px)", "translateY(0px)"],
								opacity: [0, 1]
							}}
							transition={{ duration: 0.3 }}
							onClick={() => router.push(routes.app)}
						>
							<p className="text-2xl font-bold">Get Started Now</p>
							<p className="mx-auto text-center font-bold opacity-60">
								It only takes a few seconds to create your first Plug and have it running so that you
								never miss another opportunity.
							</p>
							<Button
								variant="white"
								onClick={() => router.push(routes.app)}
								className="mt-8 border-plug-yellow bg-plug-yellow text-plug-green"
							>
								Enter App
							</Button>
						</motion.div>
					</div>
				</div>
			</LandingContainer>
		</StaticLayout>
	)
}

export default Page
