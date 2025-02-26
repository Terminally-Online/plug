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
    // We're no longer using these parameters in the display
    // const description = searchParams.get("description") ?? ""
    // const at = searchParams.get("at") ? new Date(searchParams.get("at") ?? "") : new Date()

    // Function to properly capitalize title words
    const formatTitle = (title: string) => {
      // Words to keep lowercase unless they're the first or last word
      const lowercaseWords = ['a', 'an', 'the', 'and', 'but', 'or', 'for', 'nor', 'on', 'at', 'to', 'from', 'by', 'in', 'of'];
      
      // Words to always uppercase (acronyms and special terms)
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
        // Check if it's a special term
        const lowerWord = word.toLowerCase();
        if (lowerWord in specialTerms) {
          return specialTerms[lowerWord];
        }
        
        // First word, last word, or not a lowercase word
        if (index === 0 || index === words.length - 1 || !lowercaseWords.includes(lowerWord)) {
          return word.charAt(0).toUpperCase() + word.slice(1);
        }
        
        return lowerWord;
      }).join(' ');
    }

    return new ImageResponse(
      <div
        tw="flex flex-col w-full h-full bg-[#FDFFF7] relative"
        style={{
          fontFamily: "Satoshi",
          background: "linear-gradient(30deg, #FDFFF7, #FDFFF7, rgba(210,243,138,0.2))"
        }}
      >
        {/* Background curved path inspired by landing page */}
        <svg
          viewBox="0 0 1200 630"
          fill="none"
          style={{
            position: "absolute",
            width: "100%",
            height: "100%",
            zIndex: 1
          }}
        >
          <path
            d="M1200 0C900 150 950 50 700 100C450 150 400 350 200 300C0 250 -100 450 100 550C300 650 500 500 800 550C1100 600 1050 450 1200 630"
            stroke="url(#gradient)"
            strokeWidth="40"
            fill="none"
            style={{
              opacity: 0.6
            }}
          />
          <defs>
            <linearGradient id="gradient" x1="0" y1="0" x2="100%" y2="100%" gradientUnits="userSpaceOnUse">
              <stop stopColor="#D2F38A" />
              <stop offset="1" stopColor="#385842" />
            </linearGradient>
          </defs>
        </svg>

        {/* Title container with higher z-index and adjusted positioning */}
        <div 
          tw="flex flex-col p-12 z-20 h-full justify-start pt-12"
          style={{
            position: "relative"
          }}
        >
          <h1 
            tw="text-[110px] font-black leading-[1.1] text-left max-w-[95%]"
            style={{
              backgroundImage: "linear-gradient(to right, #385842, #385842)",
              backgroundClip: "text",
              color: "transparent",
              textShadow: "0px 0px 10px rgba(210,243,138,0.5)",
              position: "relative",
              zIndex: 20,
              paddingBottom: "20px" // Add padding to prevent text cutoff
            }}
          >
            {formatTitle(name)}
          </h1>
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