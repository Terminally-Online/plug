/*
  Warnings:

  - A unique constraint covering the columns `[metadataId]` on the table `OpenseaCollectible` will be added. If there are existing duplicate values, this will fail.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" ADD COLUMN     "metadataId" TEXT;

-- CreateTable
CREATE TABLE "OpenseaCollectibleMetadata" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "OpenseaCollectibleMetadata_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "OpenseaCollectible_metadataId_key" ON "OpenseaCollectible"("metadataId");

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_metadataId_fkey" FOREIGN KEY ("metadataId") REFERENCES "OpenseaCollectibleMetadata"("id") ON DELETE SET NULL ON UPDATE CASCADE;
