import Link from "next/link";

import Siwe from "@/components/auth/siwe";
import Layout from "@/components/auth/layout";

export default function IndexPage() {
  return (
    <Layout>
      <h1>Realtime Web3</h1>
      <p>
        The framework that gives you the tools need to craft the UX your users
        want.
      </p>

      <Siwe />
      <Link href="/example">API Examples</Link>
      <Link href="/protected">Protected Route</Link>
    </Layout>
  );
}
