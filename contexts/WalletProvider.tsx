import type { FC, PropsWithChildren } from "react";
import { memo } from "react";

import { WagmiConfig } from "wagmi";
import { mainnet } from "wagmi/chains";
import { createWeb3Modal, defaultWagmiConfig } from "@web3modal/wagmi/react";

const projectId = process.env.NEXT_PUBLIC_PROJECT_ID || "";
const chains = [mainnet];
const metadata = {
  name: "Plug",
  description: '"IF This, Then That" for Ethereum.',
  url: "https://onplug.io",
  icons: ["https://onplug.io/favicon.ico"],
};
const config = defaultWagmiConfig({ chains, projectId, metadata });

createWeb3Modal({ wagmiConfig: config, projectId, chains });

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
  return <WagmiConfig config={config}>{children}</WagmiConfig>;
};

export default memo(WalletProvider);
