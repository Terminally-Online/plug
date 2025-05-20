/*
  Warnings:

  - You are about to drop the `Collectible` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `CollectibleCache` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `CollectibleMetadata` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Collection` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Fungible` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Implementation` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `ImplementationBalance` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Position` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `PositionCache` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Price` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Protocol` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheId_fkey";

-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_collectionAddress_collectionChain_fkey";

-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_metadataTokenId_metadataAddress_metadataChain_fkey";

-- DropForeignKey
ALTER TABLE "CollectibleCache" DROP CONSTRAINT "CollectibleCache_socketId_fkey";

-- DropForeignKey
ALTER TABLE "Implementation" DROP CONSTRAINT "Implementation_fungibleName_fungibleSymbol_fkey";

-- DropForeignKey
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_cacheId_fkey";

-- DropForeignKey
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_implementationChain_implementationCo_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_cacheId_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_fungibleName_fungibleSymbol_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_protocolName_fkey";

-- DropForeignKey
ALTER TABLE "PositionCache" DROP CONSTRAINT "PositionCache_socketId_fkey";

-- DropTable
DROP TABLE "Collectible";

-- DropTable
DROP TABLE "CollectibleCache";

-- DropTable
DROP TABLE "CollectibleMetadata";

-- DropTable
DROP TABLE "Collection";

-- DropTable
DROP TABLE "Fungible";

-- DropTable
DROP TABLE "Implementation";

-- DropTable
DROP TABLE "ImplementationBalance";

-- DropTable
DROP TABLE "Position";

-- DropTable
DROP TABLE "PositionCache";

-- DropTable
DROP TABLE "Price";

-- DropTable
DROP TABLE "Protocol";
