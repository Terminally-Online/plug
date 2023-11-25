import { api } from "@/lib/api";

import CanvasPreviewGrid from "@/components/canvas/blocks/CanvasPreviewGrid";

import { getSession } from "next-auth/react";
import type { GetServerSideProps } from "next";

export default async function Page() {
  const [canvases] = api.canvas.all.useSuspenseQuery();

  return <CanvasPreviewGrid canvases={canvases} />;
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
    props: {},
  };
}) satisfies GetServerSideProps;
