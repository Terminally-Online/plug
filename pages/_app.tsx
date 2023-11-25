import { Session } from "next-auth";
import { SessionProvider, getSession } from "next-auth/react";
import type { AppType } from "next/app";

import WalletProvider from "@/contexts/WalletProvider";

import { api } from "@/utils/api";

import "./styles.css";

// Use of the <SessionProvider> is mandatory to allow components that call
// `useSession()` anywhere in your application to access the `session` object.
const MyApp: AppType<{
  session: Session | null;
}> = ({ Component, pageProps }) => {
  return (
    <SessionProvider session={pageProps.session}>
      <WalletProvider>
        <Component {...pageProps} />
      </WalletProvider>
    </SessionProvider>
  );
};

MyApp.getInitialProps = async ({ ctx }) => {
  return {
    session: await getSession(ctx),
  };
};

export default api.withTRPC(MyApp);
