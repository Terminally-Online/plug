import localFont from "next/font/local";
import "./globals.css";

const satoshi = localFont({
  src: [
    { path: "../assets/Satoshi-Light.ttf", weight: "300" },
    { path: "../assets/Satoshi-Regular.ttf", weight: "400" },
    { path: "../assets/Satoshi-Bold.ttf", weight: "700" },
    { path: "../assets/Satoshi-Black.ttf", weight: "900" },
  ],
  variable: "--font-satoshi",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${satoshi.variable} bg-white`}>
      <body className="antialiased">{children}</body>
    </html>
  );
}
