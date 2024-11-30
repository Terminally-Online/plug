/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.

*/
-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("tokenId", "collectionAddress", "collectionChain");
