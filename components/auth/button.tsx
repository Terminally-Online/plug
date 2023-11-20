"use client";

import { getCsrfToken, signIn, useSession } from "next-auth/react";
import { SiweMessage } from "siwe";
import { useAccount, useNetwork, useSignMessage } from "wagmi";
import { FC, PropsWithChildren, memo } from "react";
import { useWeb3Modal } from "@web3modal/wagmi/react";

export type ButtonProps = {
  callbackUrl?: string;
  redirect?: boolean;
};

export const Button: FC<PropsWithChildren<ButtonProps>> = ({
  callbackUrl = "/canvas/",
  redirect = true,
}) => {
  const { data: session } = useSession();
  const { signMessageAsync } = useSignMessage();
  const { chain } = useNetwork();
  const { address, isConnected } = useAccount();
  const { open } = useWeb3Modal();

  const username = session?.user?.name;

  const handleLogin = async () => {
    if (!isConnected) {
      open();
      return;
    }

    try {
      const message = new SiweMessage({
        domain: window.location.host,
        address,
        statement:
          "Sign into with Plug by signing this message to prove that you are the owner of this address.",
        uri: window.location.origin,
        version: "1",
        chainId: chain?.id,
        nonce: await getCsrfToken(),
      });
      const signature = await signMessageAsync({
        message: message.prepareMessage(),
      });

      signIn("credentials", {
        message: JSON.stringify(message),
        redirect,
        signature,
        callbackUrl,
      });
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <button type="button" onClick={handleLogin}>
      {username ? username : isConnected ? "Sign In" : "Connect Wallet"}
    </button>
  );
};

export default memo(Button);
