/*
  Warnings:

  - You are about to drop the `OpenseaCollectible` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `OpenseaCollectibleCache` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `OpenseaCollectibleMetadata` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `OpenseaCollection` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_cacheSocketId_cacheChain_fkey";

-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_contract_cacheChain_fkey";

-- DropForeignKey
ALTER TABLE "OpenseaCollectibleCache" DROP CONSTRAINT "OpenseaCollectibleCache_socketId_fkey";

-- DropForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" DROP CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey";

-- DropTable
DROP TABLE "OpenseaCollectible";

-- DropTable
DROP TABLE "OpenseaCollectibleCache";

-- DropTable
DROP TABLE "OpenseaCollectibleMetadata";

-- DropTable
DROP TABLE "OpenseaCollection";
