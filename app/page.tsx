'use client'

import Link from "next/link";

export default function Page() {
  return <div className="w-screen h-screen flex flex-col items-center justify-center gap-4">
    <h1>We will move the landing page here now that this is live.</h1>

    <Link href="/canvas">Enter app</Link>
  </div>
}
