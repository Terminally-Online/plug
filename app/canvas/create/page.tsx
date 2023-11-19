import { Suspense } from "react";

import { getServerSession } from "next-auth";
import { redirect } from "next/navigation";

import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import Search from "./components/Search";
import Block from "./components/Block";
import { getServerClient } from "@/app/api/trpc/client.server";

export default async function Page({
  searchParams,
}: {
  searchParams: { search?: string };
}) {
  const session = await getServerSession(authOptions);
  const t = getServerClient(session);
  const search = searchParams.search;
  const address = session.address;

  if (!address) redirect("/connect");

  // TODO: Re-implement the search functionality by adding input to the procedure.
  const canvases = await t.all();

  // * If they are creating one, and do not have any canvases yet, bump
  //   them through this flow and take them to the tutorial board.
  if (!search && canvases.length === 0)
    redirect(
      `/canvas/${
        (await t.create({ name: "My First Canvas", public: false })).id
      }`
    );

  return (
    <div className="bg-stone-900 w-screen h-screen flex flex-col gap-2">
      {/* TODO: Implement a loading indicator */}
      <Suspense fallback={<div>Loading...</div>}>
        <Block vertical={canvases.length === 0} />

        {canvases.length > 0 ? <Search /> : <></>}
      </Suspense>
    </div>
  );
}
