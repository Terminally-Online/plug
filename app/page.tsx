import Link from "next/link";

export default function Page() {
  return <div className="w-screen h-screen flex flex-col items-center justify-center gap-4">
    <h1 className="text-2xl">Create a new Board</h1>
    <Link href={`/create`}>
      <button className="bg-black/40 text-white font-bold rounded-sm p-1 px-2">
        Create Board
      </button>
    </Link>
  </div>
}
