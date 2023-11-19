import { getServerSession } from "next-auth";
import { redirect } from "next/navigation";
import { authOptions } from "../api/auth/[...nextauth]/route";

// TODO: I may have disabled the error that was being thrown due to hydration. If it comes back, we will have to
//       use a dynamic import for Button. Ideally we will not have to do this because I do not want to have to
//       do this messed up import anytime I want to use the button.
import Button from "@/components/auth/button";

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
