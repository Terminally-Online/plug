/*
  Warnings:

  - You are about to drop the column `metdataUrl` on the `OpenseaCollectible` table. All the data in the column will be lost.
  - Added the required column `metadataUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" DROP COLUMN "metdataUrl",
ADD COLUMN     "metadataUrl" TEXT NOT NULL;
