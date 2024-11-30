/*
  Warnings:

  - You are about to drop the column `collection` on the `OpenseaCollectible` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "collection",
ALTER COLUMN "collectionId" SET DATA TYPE TEXT;

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_collectionId_fkey" FOREIGN KEY ("collectionId") REFERENCES "OpenseaCollection"("slug") ON DELETE RESTRICT ON UPDATE CASCADE;
