import { TRPCError } from "@trpc/server";
import axios from "axios";
import { getAPIKey } from "@/lib";
import { getDominantColor } from "@/server/color";
import { PrismaClient } from "@prisma/client";

export async function getMetadataForToken(
  db: PrismaClient,
  {
    type,
    address,
    chain,
    tokenId
  }: {
    type: "ERC20" | "ERC721" | "ERC1155";
    address: string;
    chain: string;
    tokenId: string;
  }
) {
  if (type === "ERC20") throw new TRPCError({ code: "NOT_IMPLEMENTED" });

  const metadataCache = await db.collectibleMetadata.findUnique({
    where: {
      tokenId_collectionAddress_collectionChain: {
        tokenId,
        collectionAddress: address,
        collectionChain: chain
      }
    }
  });

  if (metadataCache) return metadataCache;

  const url = `https://api.opensea.io/api/v2/chain/${chain}/contract/${address}/nfts/${tokenId}`;
  const response = await axios.get(url, {
    headers: {
      Accept: "application/json",
      "x-api-key": getAPIKey()
    }
  });

  if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" });

  const collectible = await db.collectible.findFirst({
    where: {
      tokenId,
      collectionAddress: address,
      collectionChain: chain
    },
    include: { collection: true }
  });

  if (collectible === null) throw new TRPCError({ code: "NOT_FOUND" });

  const traits = response.data.nft.traits === null ? [] : response.data.nft.traits;
  const colorUrl = collectible.previewUrl ?? collectible.collection.iconUrl ?? "";
  const color = await getDominantColor(colorUrl);

  return await db.collectibleMetadata.upsert({
    where: {
      tokenId_collectionAddress_collectionChain: {
        tokenId,
        collectionAddress: address,
        collectionChain: chain
      }
    },
    create: {
      tokenId,
      collectionAddress: address,
      collectionChain: chain,
      traits,
      color
    },
    update: {
      traits,
      color
    }
  });
}