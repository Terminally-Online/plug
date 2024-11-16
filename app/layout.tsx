import localFont from "next/font/local"

const satoshi = localFont({
  src: [
    { path: "../assets/Satoshi-Light.ttf", weight: "300" },
    { path: "../assets/Satoshi-Regular.ttf", weight: "400" },
    { path: "../assets/Satoshi-Bold.ttf", weight: "700" },
    { path: "../assets/Satoshi-Black.ttf", weight: "900" }
  ],
  variable: "--font-satoshi"
})

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className={`${satoshi.variable}`}>
      <body>{children}</body>
    </html>
  )
}