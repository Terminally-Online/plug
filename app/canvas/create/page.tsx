export default function Page() { 
  return <div className="bg-stone-900 w-screen h-screen flex flex-col gap-2 items-center justify-center">
    <div className="border-[1px] border-stone-950 p-4">
      <h1 className="text-2xl text-white">Create a Canvas</h1>
    </div>

    <hr />

    <input placeholder="Search all Canvases..." />
  </div>
}
