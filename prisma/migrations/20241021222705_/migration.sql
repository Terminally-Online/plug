-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_tokenId_collectionAddress_collectionChain_fkey" FOREIGN KEY ("tokenId", "collectionAddress", "collectionChain") REFERENCES "CollectibleMetadata"("tokenId", "collectionAddress", "collectionChain") ON DELETE RESTRICT ON UPDATE CASCADE;
