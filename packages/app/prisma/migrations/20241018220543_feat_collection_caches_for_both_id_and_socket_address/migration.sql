/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `cacheSocketId` on the `Collectible` table. All the data in the column will be lost.
  - The primary key for the `CollectibleCache` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - Added the required column `cacheId` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - The required column `id` was added to the `CollectibleCache` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheSocketId_fkey";

-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
DROP COLUMN "cacheSocketId",
ADD COLUMN     "cacheId" TEXT NOT NULL,
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("cacheId", "tokenId", "collectionAddress", "collectionChain");

-- AlterTable
ALTER TABLE "CollectibleCache" DROP CONSTRAINT "CollectibleCache_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "CollectibleCache_pkey" PRIMARY KEY ("id");

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "CollectibleCache"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
