import { Suspense } from "react";

import type { GetServerSideProps } from "next";
import { Session } from "next-auth";
import { getSession } from "next-auth/react";

import { api } from "@/lib/api";
import CanvasPreviewGrid from "@/components/canvas/blocks/CanvasPreviewGrid";
import { TabsProvider } from "@/contexts/TabsProvider";

export default function Page() {
  const { data: canvases } = api.canvas.all.useQuery();

  return (
    <TabsProvider>
      <Suspense fallback={<div>Loading...</div>}>
        <CanvasPreviewGrid canvases={canvases} />
      </Suspense>
    </TabsProvider>
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
      session,
    },
  };
}) satisfies GetServerSideProps<{ session: Session | null }>;
