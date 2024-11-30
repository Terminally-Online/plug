-- AlterTable
ALTER TABLE "Collectible" ADD COLUMN     "metadataAddress" TEXT,
ADD COLUMN     "metadataChain" TEXT,
ADD COLUMN     "metadataTokenId" TEXT;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_metadataTokenId_metadataAddress_metadataChain_fkey" FOREIGN KEY ("metadataTokenId", "metadataAddress", "metadataChain") REFERENCES "CollectibleMetadata"("tokenId", "collectionAddress", "collectionChain") ON DELETE SET NULL ON UPDATE CASCADE;
