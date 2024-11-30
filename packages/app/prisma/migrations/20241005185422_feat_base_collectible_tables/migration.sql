/*
  Warnings:

  - Added the required column `tokenId` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Made the column `cacheSocketId` on table `Collectible` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheSocketId_fkey";

-- AlterTable
ALTER TABLE "Collectible" ADD COLUMN     "collectionAddress" TEXT,
ADD COLUMN     "collectionChain" TEXT,
ADD COLUMN     "tokenId" TEXT NOT NULL,
ALTER COLUMN "cacheSocketId" SET NOT NULL;

-- CreateTable
CREATE TABLE "Collection" (
    "address" TEXT NOT NULL,
    "chain" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "slug" TEXT NOT NULL,
    "collection" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "imageUrl" TEXT NOT NULL,
    "bannerImageUrl" TEXT NOT NULL,
    "owner" TEXT NOT NULL,
    "category" TEXT NOT NULL,
    "isDisabled" BOOLEAN NOT NULL,
    "isNsfw" BOOLEAN NOT NULL,
    "traitOffersEnabled" BOOLEAN NOT NULL,
    "collectionOffersEnabled" BOOLEAN NOT NULL,
    "openseaUrl" TEXT NOT NULL,
    "projectUrl" TEXT NOT NULL,
    "wikiUrl" TEXT NOT NULL,
    "discordUrl" TEXT NOT NULL,
    "telegramUrl" TEXT NOT NULL,
    "twitterUsername" TEXT NOT NULL,
    "instagramUsername" TEXT NOT NULL,
    "totalSupply" INTEGER NOT NULL,

    CONSTRAINT "Collection_pkey" PRIMARY KEY ("address","chain")
);

-- CreateTable
CREATE TABLE "CollectibleMetadata" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "traits" JSONB[],
    "color" TEXT,
    "collectibleId" TEXT NOT NULL,

    CONSTRAINT "CollectibleMetadata_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "Collection_slug_chain_idx" ON "Collection"("slug", "chain");

-- CreateIndex
CREATE UNIQUE INDEX "CollectibleMetadata_collectibleId_key" ON "CollectibleMetadata"("collectibleId");

-- AddForeignKey
ALTER TABLE "CollectibleMetadata" ADD CONSTRAINT "CollectibleMetadata_collectibleId_fkey" FOREIGN KEY ("collectibleId") REFERENCES "Collectible"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey" FOREIGN KEY ("collectionAddress", "collectionChain") REFERENCES "Collection"("address", "chain") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheSocketId_fkey" FOREIGN KEY ("cacheSocketId") REFERENCES "CollectibleCache"("socketId") ON DELETE RESTRICT ON UPDATE CASCADE;
