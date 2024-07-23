/*
  Warnings:

  - You are about to drop the column `banner_image_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `collection_offers_enabled` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `created_date` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `discord_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `image_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `instagram_username` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `is_disabled` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `is_nsfw` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `opensea_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `project_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `telegram_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `total_supply` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `trait_offers_enabled` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `twitter_username` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `wiki_url` on the `OpenseaCollection` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[slug]` on the table `OpenseaCollection` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `bannerImageUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `collectionOffersEnabled` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `createdDate` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `discordUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `imageUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `instagramUsername` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isDisabled` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isNsfw` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `openseaUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `projectUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `telegramUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `totalSupply` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `traitOffersEnabled` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `twitterUsername` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `wikiUrl` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollection" DROP COLUMN "banner_image_url",
DROP COLUMN "collection_offers_enabled",
DROP COLUMN "created_date",
DROP COLUMN "discord_url",
DROP COLUMN "image_url",
DROP COLUMN "instagram_username",
DROP COLUMN "is_disabled",
DROP COLUMN "is_nsfw",
DROP COLUMN "opensea_url",
DROP COLUMN "project_url",
DROP COLUMN "telegram_url",
DROP COLUMN "total_supply",
DROP COLUMN "trait_offers_enabled",
DROP COLUMN "twitter_username",
DROP COLUMN "wiki_url",
ADD COLUMN     "bannerImageUrl" TEXT NOT NULL,
ADD COLUMN     "collectionOffersEnabled" BOOLEAN NOT NULL,
ADD COLUMN     "createdDate" TIMESTAMP(3) NOT NULL,
ADD COLUMN     "discordUrl" TEXT NOT NULL,
ADD COLUMN     "imageUrl" TEXT NOT NULL,
ADD COLUMN     "instagramUsername" TEXT NOT NULL,
ADD COLUMN     "isDisabled" BOOLEAN NOT NULL,
ADD COLUMN     "isNsfw" BOOLEAN NOT NULL,
ADD COLUMN     "openseaUrl" TEXT NOT NULL,
ADD COLUMN     "projectUrl" TEXT NOT NULL,
ADD COLUMN     "telegramUrl" TEXT NOT NULL,
ADD COLUMN     "totalSupply" INTEGER NOT NULL,
ADD COLUMN     "traitOffersEnabled" BOOLEAN NOT NULL,
ADD COLUMN     "twitterUsername" TEXT NOT NULL,
ADD COLUMN     "wikiUrl" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "OpenseaCollection_slug_key" ON "OpenseaCollection"("slug");
