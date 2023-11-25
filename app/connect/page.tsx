import { redirect } from "next/navigation";

import { getServerSession } from "next-auth";

import { authOptions } from "@/server/auth";

// TODO: I may have disabled the error that was being thrown due to hydration. If it comes back, we will have to
//       use a dynamic import for Button. Ideally we will not have to do this because I do not want to have to
//       do this messed up import anytime I want to use the button.
import Button from "@/components/auth/button";

export const revalidate = 0;

export default async function Page() {
  const session = await getServerSession(authOptions);

  const username = session?.user?.name;

  if (username) redirect("/canvas");

  return (
    <>
      <h1>Log in to proceed</h1>

      <Button />
    </>
  );
}
