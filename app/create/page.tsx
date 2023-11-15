import { redirect } from "next/navigation";

export default function Page() { 
  redirect(`/${Math.random().toString(36).substring(7)}`);
}
