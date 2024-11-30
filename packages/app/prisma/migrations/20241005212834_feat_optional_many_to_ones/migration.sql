-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheSocketId_fkey";

-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey";

-- AlterTable
ALTER TABLE "Collectible" ALTER COLUMN "cacheSocketId" DROP NOT NULL,
ALTER COLUMN "collectionAddress" DROP NOT NULL,
ALTER COLUMN "collectionChain" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey" FOREIGN KEY ("collectionAddress", "collectionChain") REFERENCES "Collection"("address", "chain") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheSocketId_fkey" FOREIGN KEY ("cacheSocketId") REFERENCES "CollectibleCache"("socketId") ON DELETE SET NULL ON UPDATE CASCADE;
