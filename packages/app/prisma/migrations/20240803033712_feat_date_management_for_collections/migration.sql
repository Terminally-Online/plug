/*
  Warnings:

  - You are about to drop the column `createdDate` on the `OpenseaCollection` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "OpenseaCollection" DROP COLUMN "createdDate",
ADD COLUMN     "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- CreateTable
CREATE TABLE "OpenseaCollectible" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "OpenseaCollectible_pkey" PRIMARY KEY ("id")
);
