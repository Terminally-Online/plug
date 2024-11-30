/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Collectible` table. All the data in the column will be lost.
  - The primary key for the `CollectibleMetadata` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `collectibleId` on the `CollectibleMetadata` table. All the data in the column will be lost.
  - You are about to drop the column `id` on the `CollectibleMetadata` table. All the data in the column will be lost.
  - Made the column `collectionAddress` on table `Collectible` required. This step will fail if there are existing NULL values in that column.
  - Made the column `collectionChain` on table `Collectible` required. This step will fail if there are existing NULL values in that column.
  - Added the required column `collectibleCollectionAddress` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectibleCollectionChain` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectibleTokenId` to the `CollectibleMetadata` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey";

-- DropForeignKey
ALTER TABLE "CollectibleMetadata" DROP CONSTRAINT "CollectibleMetadata_collectibleId_fkey";

-- DropIndex
DROP INDEX "CollectibleMetadata_collectibleId_key";

-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
DROP COLUMN "id",
ADD COLUMN     "collectibleMetadataId" TEXT,
ALTER COLUMN "collectionAddress" SET NOT NULL,
ALTER COLUMN "collectionChain" SET NOT NULL,
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("tokenId", "collectionAddress", "collectionChain");

-- AlterTable
ALTER TABLE "CollectibleMetadata" DROP CONSTRAINT "CollectibleMetadata_pkey",
DROP COLUMN "collectibleId",
DROP COLUMN "id",
ADD COLUMN     "collectibleCollectionAddress" TEXT NOT NULL,
ADD COLUMN     "collectibleCollectionChain" TEXT NOT NULL,
ADD COLUMN     "collectibleTokenId" TEXT NOT NULL,
ADD CONSTRAINT "CollectibleMetadata_pkey" PRIMARY KEY ("collectibleTokenId", "collectibleCollectionAddress", "collectibleCollectionChain");

-- AddForeignKey
ALTER TABLE "CollectibleMetadata" ADD CONSTRAINT "CollectibleMetadata_collectibleTokenId_collectibleCollecti_fkey" FOREIGN KEY ("collectibleTokenId", "collectibleCollectionAddress", "collectibleCollectionChain") REFERENCES "Collectible"("tokenId", "collectionAddress", "collectionChain") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey" FOREIGN KEY ("collectionAddress", "collectionChain") REFERENCES "Collection"("address", "chain") ON DELETE RESTRICT ON UPDATE CASCADE;
