"use client"

import { getCsrfToken, signIn, useSession } from "next-auth/react";
import { SiweMessage } from "siwe";
import { useAccount, useConnect, useNetwork, useSignMessage } from "wagmi"
import { InjectedConnector } from "wagmi/connectors/injected";
import { useEffect } from "react";

export const Button = () => { 
  const { signMessageAsync } = useSignMessage()
  const { chain } = useNetwork()
  const { address, isConnected } = useAccount()
  const { connect } = useConnect({ connector: new InjectedConnector() })
  const { data: session } = useSession()

  const handleLogin = async () => {
    if(!isConnected) connect()

    try { 
      const callbackUrl = '/protected'
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
        redirect: false,
        signature,
        callbackUrl
      })
    } catch (e) {
      window.alert(e)
    }
  }

  useEffect(() => { 
    if(isConnected && !session) handleLogin()
  }, [isConnected])

  return <button onClick={(e) => { 
    e.preventDefault()
    handleLogin()
  }}>Sign In</button>
}
