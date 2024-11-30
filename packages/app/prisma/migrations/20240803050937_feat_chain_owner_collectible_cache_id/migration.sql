/*
  Warnings:

  - You are about to drop the column `cacheId` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - The primary key for the `OpenseaCollectibleCache` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `OpenseaCollectibleCache` table. All the data in the column will be lost.
  - Added the required column `cacheChain` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `cacheOwner` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_cacheId_fkey";

-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "cacheId",
ADD COLUMN     "cacheChain" TEXT NOT NULL,
ADD COLUMN     "cacheOwner" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "OpenseaCollectibleCache" DROP CONSTRAINT "OpenseaCollectibleCache_pkey",
DROP COLUMN "id",
ADD CONSTRAINT "OpenseaCollectibleCache_pkey" PRIMARY KEY ("chain", "owner");

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_cacheChain_cacheOwner_fkey" FOREIGN KEY ("cacheChain", "cacheOwner") REFERENCES "OpenseaCollectibleCache"("chain", "owner") ON DELETE RESTRICT ON UPDATE CASCADE;
