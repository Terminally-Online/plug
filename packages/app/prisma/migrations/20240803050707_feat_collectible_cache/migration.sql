/*
  Warnings:

  - Added the required column `cacheId` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" ADD COLUMN     "cacheId" TEXT NOT NULL;

-- CreateTable
CREATE TABLE "OpenseaCollectibleCache" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "chain" TEXT NOT NULL,
    "owner" TEXT NOT NULL,

    CONSTRAINT "OpenseaCollectibleCache_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "OpenseaCollectibleCache"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
