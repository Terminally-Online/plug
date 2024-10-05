/*
  Warnings:

  - You are about to drop the column `collectionSlug` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - The primary key for the `OpenseaCollection` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - Added the required column `address` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_collectionSlug_fkey";

-- DropIndex
DROP INDEX "OpenseaCollection_slug_key";

-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "collectionSlug";

-- AlterTable
ALTER TABLE "OpenseaCollection" DROP CONSTRAINT "OpenseaCollection_pkey",
ADD COLUMN     "address" TEXT NOT NULL,
ADD CONSTRAINT "OpenseaCollection_pkey" PRIMARY KEY ("address", "chain");

-- CreateTable
CREATE TABLE "Collectible" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "cacheSocketId" TEXT,

    CONSTRAINT "Collectible_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "CollectibleCache" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "CollectibleCache_pkey" PRIMARY KEY ("socketId")
);

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_contract_cacheChain_fkey" FOREIGN KEY ("contract", "cacheChain") REFERENCES "OpenseaCollection"("address", "chain") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheSocketId_fkey" FOREIGN KEY ("cacheSocketId") REFERENCES "CollectibleCache"("socketId") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "CollectibleCache" ADD CONSTRAINT "CollectibleCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
