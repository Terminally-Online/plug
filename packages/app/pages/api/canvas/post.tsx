/* eslint-disable @next/next/no-img-element */
import { NextRequest } from "next/server"
import { ImageResponse } from "@vercel/og"

export const config = {
  runtime: "edge",
}

export default async function handler(req: NextRequest) {
  const regular = await fetch(
    new URL("../../../assets/Satoshi-Regular.ttf", import.meta.url)
  ).then((res) => res.arrayBuffer())
  const bold = await fetch(
    new URL("../../../assets/Satoshi-Bold.ttf", import.meta.url)
  ).then((res) => res.arrayBuffer())

  try {
    const { searchParams } = req.nextUrl

    const name = searchParams.get("name") ?? ""
    const description = searchParams.get("description") ?? ""
    const author = searchParams.get("author") ?? "drakedanner"
    const at = searchParams.get("at") ? new Date(searchParams.get("at") ?? "") : new Date()

    const formatTitle = (title: string) => {
      const lowercaseWords = ['a', 'an', 'the', 'and', 'but', 'or', 'for', 'nor', 'on', 'at', 'to', 'from', 'by', 'in', 'of'];
      const specialTerms: Record<string, string> = {
        'defi': 'DeFi',
        'nft': 'NFT',
        'dao': 'DAO',
        'web3': 'Web3',
        'plug': 'Plug',
        'and': '&'
      };

      const words = title.split(' ');

      return words.map((word, index) => {
        const lowerWord = word.toLowerCase();
        if (lowerWord in specialTerms) {
          return specialTerms[lowerWord];
        }

        if (index === 0 || index === words.length - 1 || !lowercaseWords.includes(lowerWord)) {
          return word.charAt(0).toUpperCase() + word.slice(1);
        }

        return lowerWord;
      }).join(' ');
    }

    return new ImageResponse(
      <div
        tw="flex flex-col w-full h-full relative justify-end relative"
        style={{
          fontFamily: "Satoshi",
          background: "#FEFFF7"
        }}
      >
        <div tw="flex flex-col p-12">
          <img tw="h-16 my-12" src="https://onplug.io/plug-logo-lime.svg" alt="plug logo" />

          <h1 tw="relative text-[110px] font-black text-left max-w-[90%] text-[#385842]">
            {formatTitle(name)}
          </h1>
          <p tw="text-[24px] opacity-40 font-bold w-[60%]">
            {description.length > 80 ? `${description}...` : description}
          </p>
        </div>

        <div tw="flex flex-row justify-between bg-[#D2F38A] px-12 py-4 font-bold">
          <div tw="flex flex-row items-center">
            <img tw="w-8 h-8 rounded-full mr-4" src={`https://cdn.onplug.io/users/${author}.png`} alt={author} />
            <p>{author}</p>
          </div>
          <p>{at.toLocaleDateString()}</p>
        </div>
      </div>,
      {
        width: 1200,
        height: 630,
        fonts: [
          {
            name: "Satoshi",
            data: regular,
            style: "normal",
            weight: 400,
          },
          {
            name: "Satoshi",
            data: bold,
            style: "normal",
            weight: 700,
          },
        ],
      }
    )
  } catch (e: any) {
    return new Response("Failed to generate image, " + e.message, {
      status: 500,
    })
  }
}
