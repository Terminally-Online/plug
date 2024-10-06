/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Collectible` table. All the data in the column will be lost.
  - Added the required column `interface` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isSpam` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Made the column `cacheSocketId` on table `Collectible` required. This step will fail if there are existing NULL values in that column.
  - Made the column `collectionAddress` on table `Collectible` required. This step will fail if there are existing NULL values in that column.
  - Made the column `collectionChain` on table `Collectible` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheSocketId_fkey";

-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey";

-- DropIndex
DROP INDEX "Collectible_id_key";

-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
DROP COLUMN "id",
ADD COLUMN     "imageUrl" TEXT,
ADD COLUMN     "interface" TEXT NOT NULL,
ADD COLUMN     "isSpam" BOOLEAN NOT NULL,
ADD COLUMN     "previewUrl" TEXT,
ADD COLUMN     "videoUrl" TEXT,
ALTER COLUMN "cacheSocketId" SET NOT NULL,
ALTER COLUMN "collectionAddress" SET NOT NULL,
ALTER COLUMN "collectionChain" SET NOT NULL,
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("cacheSocketId", "tokenId", "collectionAddress", "collectionChain");

-- AlterTable
ALTER TABLE "Collection" ALTER COLUMN "iconUrl" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey" FOREIGN KEY ("collectionAddress", "collectionChain") REFERENCES "Collection"("address", "chain") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheSocketId_fkey" FOREIGN KEY ("cacheSocketId") REFERENCES "CollectibleCache"("socketId") ON DELETE RESTRICT ON UPDATE CASCADE;
