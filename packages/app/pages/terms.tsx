import { StaticLayout } from "@/components/landing/layout/static"
import { LandingContainer } from "@/components/landing/layout/container"
import ReactMarkdown from "react-markdown"
import fs from "fs"
import path from "path"
import Image from "next/image"
import { ExternalLink } from "lucide-react"

export const getStaticProps = async () => {
	const termsPath = path.join(process.cwd(), "content/terms.md")
	const termsContent = fs.readFileSync(termsPath, "utf8")

	return {
		props: {
			termsContent
		}
	}
}

const TermsOfService = ({ termsContent }: { termsContent: string }) => {
	return (
		<StaticLayout title="Terms of Service" description="Terms of Service for Plug">
			<LandingContainer>
				<div className="mx-auto flex w-full max-w-3xl flex-col gap-4 overflow-x-hidden pb-24 pt-8">
					<h1 className="text-[48px] font-black leading-tight md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[1240px] xl:text-[96px]">
						Terms of Service
					</h1>
					<div className="max-w-none">
						<ReactMarkdown
							components={{
								h1: ({ children }) => (
									<h1 className="text-[48px] font-black md:text-[72px] lg:text-[96px] xl:text-[120px]">{children}</h1>
								),
								h2: ({ children }) => <h2 className="mt-8 text-[2.5rem] font-black">{children}</h2>,
								h3: ({ children }) => <h3 className="mt-8 text-[2.25rem] font-black">{children}</h3>,
								h4: ({ children }) => <h4 className="mt-8 text-[2rem] font-black">{children}</h4>,
								h5: ({ children }) => <h5 className="mt-8 text-[1.5rem] font-black">{children}</h5>,
								h6: ({ children }) => <h6 className="mt-8 text-[1.25rem] font-black">{children}</h6>,
								p: ({ children }) => <p className="break-words mb-6">{children}</p>,
								ul: ({ children }) => <ul className="ml-6 list-disc break-words opacity-80 mb-6">{children}</ul>,
								ol: ({ children }) => <ol className="ml-6 list-decimal break-words opacity-80 mb-6">{children}</ol>,
								li: ({ children }) => <li className="mb-2 ml-2 list-item">{children}</li>,
								a: ({ children, ...props }) => {
									return (
										<span
											{...props}
											className="group inline cursor-pointer break-words underline transition-opacity duration-200 hover:opacity-80"
											onClick={e => {
												e.preventDefault()
												e.stopPropagation()
												if (props.href) {
													window.open(props.href, "_blank", "noopener,noreferrer")
												}
											}}
										>
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
							{termsContent}
						</ReactMarkdown>
					</div>
				</div>
			</LandingContainer>
		</StaticLayout>
	)
}

export default TermsOfService 