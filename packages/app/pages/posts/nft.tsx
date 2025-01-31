import Image from "next/image"

const Page = () => {
	return <div className="grid grid-cols-8 gap-2 p-4">
		{Array.from({ length: 300 }).map((_, index) => (
			<div key={index} className="relative aspect-square w-full">
				<Image
					src={`http://localhost:3000/api/canvas/nft?number=${index}`}
					alt={`NFT ${index}`}
					fill
				/>
			</div>
		))}
	</div>
}

export default Page
