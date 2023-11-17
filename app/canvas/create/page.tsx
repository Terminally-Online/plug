import { redirect } from "next/navigation";

export default function Page() { 
  redirect(`/canvas/${Math.random().toString(36).substring(7)}`);
}
