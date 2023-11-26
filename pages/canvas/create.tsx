import { Suspense } from "react";

import { redirect } from "next/navigation";

import { getSession } from "next-auth/react";
import type { GetServerSideProps, InferGetServerSidePropsType } from "next";

import { type NextPageWithLayout } from "@/lib/types";
import { api } from "@/lib/api";
import Search from "@/components/canvas/Search";
import Block from "@/components/canvas/Block";
import { TabsProvider } from "@/contexts/TabsProvider";

const Page: NextPageWithLayout<
  InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
  const { data: canvases } = api.canvas.all.useQuery();

  return (
    <div className="bg-stone-900 h-full flex flex-col gap-2">
      <Suspense fallback={<div>Loading...</div>}>
        <Block vertical={(canvases && canvases.length === 0) || false} />

        {canvases && canvases.length > 0 ? <Search /> : <></>}
      </Suspense>
    </div>
  );
};

export const getServerSideProps = (async (context) => {
  if (!(await getSession(context))) {
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

Page.getLayout = (page) => <TabsProvider>{page}</TabsProvider>;

export default Page;
