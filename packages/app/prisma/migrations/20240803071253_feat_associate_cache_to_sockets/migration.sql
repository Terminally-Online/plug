/*
  Warnings:

  - You are about to drop the column `cacheOwner` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - The primary key for the `OpenseaCollectibleCache` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `owner` on the `OpenseaCollectibleCache` table. All the data in the column will be lost.
  - Added the required column `cacheSocketId` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketId` to the `OpenseaCollectibleCache` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_cacheChain_cacheOwner_fkey";

-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "cacheOwner",
ADD COLUMN     "cacheSocketId" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "OpenseaCollectibleCache" DROP CONSTRAINT "OpenseaCollectibleCache_pkey",
DROP COLUMN "owner",
ADD COLUMN     "socketId" TEXT NOT NULL,
ADD CONSTRAINT "OpenseaCollectibleCache_pkey" PRIMARY KEY ("socketId", "chain");

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_cacheSocketId_cacheChain_fkey" FOREIGN KEY ("cacheSocketId", "cacheChain") REFERENCES "OpenseaCollectibleCache"("socketId", "chain") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OpenseaCollectibleCache" ADD CONSTRAINT "OpenseaCollectibleCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
