/*
  Warnings:

  - Made the column `collectibleId` on table `OpenseaCollectibleMetadata` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" DROP CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey";

-- AlterTable
ALTER TABLE "OpenseaCollectibleMetadata" ALTER COLUMN "collectibleId" SET NOT NULL;

-- AddForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" ADD CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey" FOREIGN KEY ("collectibleId") REFERENCES "OpenseaCollectible"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
