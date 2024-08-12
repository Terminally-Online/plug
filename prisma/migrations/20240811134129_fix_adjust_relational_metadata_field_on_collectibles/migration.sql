/*
  Warnings:

  - You are about to drop the column `metadataId` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[collectibleId]` on the table `OpenseaCollectibleMetadata` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_metadataId_fkey";

-- DropIndex
DROP INDEX "OpenseaCollectible_metadataId_key";

-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "metadataId";

-- AlterTable
ALTER TABLE "OpenseaCollectibleMetadata" ADD COLUMN     "collectibleId" TEXT;

-- CreateIndex
CREATE UNIQUE INDEX "OpenseaCollectibleMetadata_collectibleId_key" ON "OpenseaCollectibleMetadata"("collectibleId");

-- AddForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" ADD CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey" FOREIGN KEY ("collectibleId") REFERENCES "OpenseaCollectible"("id") ON DELETE SET NULL ON UPDATE CASCADE;
