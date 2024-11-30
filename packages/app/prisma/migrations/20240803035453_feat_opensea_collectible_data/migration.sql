/*
  Warnings:

  - Added the required column `collection` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectionId` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `contract` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `description` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `displayAnimationUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `displayImageUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `identifier` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `imageUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isDisabled` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isNsfw` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `metdataUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `name` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `openseaUrl` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `tokenStandard` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" ADD COLUMN     "collection" TEXT NOT NULL,
ADD COLUMN     "collectionId" INTEGER NOT NULL,
ADD COLUMN     "contract" TEXT NOT NULL,
ADD COLUMN     "description" TEXT NOT NULL,
ADD COLUMN     "displayAnimationUrl" TEXT NOT NULL,
ADD COLUMN     "displayImageUrl" TEXT NOT NULL,
ADD COLUMN     "identifier" TEXT NOT NULL,
ADD COLUMN     "imageUrl" TEXT NOT NULL,
ADD COLUMN     "isDisabled" BOOLEAN NOT NULL,
ADD COLUMN     "isNsfw" BOOLEAN NOT NULL,
ADD COLUMN     "metdataUrl" TEXT NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL,
ADD COLUMN     "openseaUrl" TEXT NOT NULL,
ADD COLUMN     "tokenStandard" TEXT NOT NULL;
