"use client"

import type { FC, PropsWithChildren } from "react";
import { memo } from "react";

import { WagmiConfig, createConfig, mainnet } from 'wagmi'
import { createPublicClient, http } from 'viem'

const config = createConfig({ 
  autoConnect: true, 
  publicClient: createPublicClient({
    chain: mainnet,
    transport: http()
  }) 
})

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
  return <WagmiConfig config={config}>
    {children}
  </WagmiConfig>;
}

export default memo(WalletProvider);
