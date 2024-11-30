/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - The primary key for the `CollectibleMetadata` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `collectibleCollectionAddress` on the `CollectibleMetadata` table. All the data in the column will be lost.
  - You are about to drop the column `collectibleCollectionChain` on the `CollectibleMetadata` table. All the data in the column will be lost.
  - You are about to drop the column `collectibleTokenId` on the `CollectibleMetadata` table. All the data in the column will be lost.
  - Added the required column `amount` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectionAddress` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectionChain` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.
  - Added the required column `tokenId` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "CollectibleMetadata" DROP CONSTRAINT "CollectibleMetadata_collectibleTokenId_collectibleCollecti_fkey";

-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
ADD COLUMN     "amount" TEXT NOT NULL,
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("cacheSocketId", "tokenId", "collectionAddress", "collectionChain");

-- AlterTable
ALTER TABLE "CollectibleMetadata" DROP CONSTRAINT "CollectibleMetadata_pkey",
DROP COLUMN "collectibleCollectionAddress",
DROP COLUMN "collectibleCollectionChain",
DROP COLUMN "collectibleTokenId",
ADD COLUMN     "collectionAddress" TEXT NOT NULL,
ADD COLUMN     "collectionChain" TEXT NOT NULL,
ADD COLUMN     "tokenId" TEXT NOT NULL,
ADD CONSTRAINT "CollectibleMetadata_pkey" PRIMARY KEY ("tokenId", "collectionAddress", "collectionChain");

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_tokenId_collectionAddress_collectionChain_fkey" FOREIGN KEY ("tokenId", "collectionAddress", "collectionChain") REFERENCES "CollectibleMetadata"("tokenId", "collectionAddress", "collectionChain") ON DELETE RESTRICT ON UPDATE CASCADE;
