import { Suspense } from "react";

import { redirect } from "next/navigation";

import { getSession } from "next-auth/react";
import type { GetServerSideProps, InferGetServerSidePropsType } from "next";

import { api } from "@/lib/api";

import Search from "@/components/canvas/Search";
import Block from "@/components/canvas/Block";

export default async function Page({
  search,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const [canvases] = api.canvas.all.useSuspenseQuery();
  const createCanvas = api.canvas.create.useMutation();

  if (!search && canvases.length === 0) {
    const canvas = await createCanvas.mutateAsync({
      name: "My First Canvas",
      public: false,
    });

    redirect(`/canvas/${canvas.id}`);
  }

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

export const getServerSideProps = (async (context) => {
  const session = await getSession(context);

  if (!session) {
    return {
      redirect: {
        destination: `/connect`,
        permanent: false,
      },
    };
  }

  return {
    props: {
      search: context.query.search || "",
    },
  };
}) satisfies GetServerSideProps<{
  search: string | string[] | undefined;
}>;
