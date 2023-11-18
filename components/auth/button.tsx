"use client"

import { getCsrfToken, signIn, useSession } from "next-auth/react";
import { SiweMessage } from "siwe";
import { useAccount, useNetwork, useSignMessage } from "wagmi"
import { FC, PropsWithChildren, memo, useEffect } from "react";
import { useWeb3Modal } from "@web3modal/wagmi/react";

export type ButtonProps = { 
  callbackUrl?: string
}

export const Button: FC<PropsWithChildren<ButtonProps>> = ({ callbackUrl = '/canvas/' }) => { 
  const { data: session } = useSession()
  const { signMessageAsync } = useSignMessage()
  const { chain } = useNetwork()
  const { address, isConnected } = useAccount()
  const { open } = useWeb3Modal()

  const username = session?.user?.name

  const handleLogin = async () => {
    if(!isConnected) { 
      open(); 
      return 
    }

    try { 
      const message = new SiweMessage({
        domain: window.location.host,
        address,
        statement: "Sign in with Ethereum to the app.",
        uri: window.location.origin,
        version: "1",
        chainId: chain?.id,
        nonce: await getCsrfToken()
      })
      const signature = await signMessageAsync({ message: message.prepareMessage() })

      signIn("credentials", {
        message: JSON.stringify(message),
        redirect: true,
        signature,
        callbackUrl
      })
    } catch (e) {
      console.error(e)
    }
  }

  useEffect(() => { 
    if(isConnected && !session) handleLogin()
  }, [isConnected])

  return <button type="button" onClick={handleLogin}>
    {username ? username : isConnected ? 'Sign In' : 'Connect Wallet'}
  </button>
}

export default memo(Button);
