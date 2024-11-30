/*
  Warnings:

  - You are about to drop the column `bannerImageUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `category` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `collection` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `collectionOffersEnabled` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `description` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `discordUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `imageUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `instagramUsername` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `isDisabled` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `isNsfw` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `name` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `openseaUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `owner` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `projectUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `slug` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `telegramUrl` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `totalSupply` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `traitOffersEnabled` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `twitterUsername` on the `Collection` table. All the data in the column will be lost.
  - You are about to drop the column `wikiUrl` on the `Collection` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "Collection_slug_chain_idx";

-- AlterTable
ALTER TABLE "Collection" DROP COLUMN "bannerImageUrl",
DROP COLUMN "category",
DROP COLUMN "collection",
DROP COLUMN "collectionOffersEnabled",
DROP COLUMN "description",
DROP COLUMN "discordUrl",
DROP COLUMN "imageUrl",
DROP COLUMN "instagramUsername",
DROP COLUMN "isDisabled",
DROP COLUMN "isNsfw",
DROP COLUMN "name",
DROP COLUMN "openseaUrl",
DROP COLUMN "owner",
DROP COLUMN "projectUrl",
DROP COLUMN "slug",
DROP COLUMN "telegramUrl",
DROP COLUMN "totalSupply",
DROP COLUMN "traitOffersEnabled",
DROP COLUMN "twitterUsername",
DROP COLUMN "wikiUrl";
