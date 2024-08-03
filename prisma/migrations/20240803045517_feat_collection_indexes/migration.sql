/*
  Warnings:

  - You are about to drop the column `collectionId` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - Added the required column `collectionSlug` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_collectionId_fkey";

-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "collectionId",
ADD COLUMN     "collectionSlug" TEXT NOT NULL;

-- CreateIndex
CREATE INDEX "OpenseaCollection_slug_chain_idx" ON "OpenseaCollection"("slug", "chain");

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_collectionSlug_fkey" FOREIGN KEY ("collectionSlug") REFERENCES "OpenseaCollection"("slug") ON DELETE RESTRICT ON UPDATE CASCADE;
