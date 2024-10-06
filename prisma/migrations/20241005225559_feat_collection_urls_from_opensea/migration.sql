/*
  Warnings:

  - You are about to drop the column `discordUrl` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `instagramUsername` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `openseaUrl` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `projectUrl` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `telegramUrl` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `twitterUsername` on the `OpenseaCollection` table. All the data in the column will be lost.
  - You are about to drop the column `wikiUrl` on the `OpenseaCollection` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Collection" ADD COLUMN     "discordUrl" TEXT,
ADD COLUMN     "instagramUsername" TEXT,
ADD COLUMN     "openseaUrl" TEXT,
ADD COLUMN     "projectUrl" TEXT,
ADD COLUMN     "telegramUrl" TEXT,
ADD COLUMN     "twitterUsername" TEXT,
ADD COLUMN     "wikiUrl" TEXT;

-- AlterTable
ALTER TABLE "OpenseaCollection" DROP COLUMN "discordUrl",
DROP COLUMN "instagramUsername",
DROP COLUMN "openseaUrl",
DROP COLUMN "projectUrl",
DROP COLUMN "telegramUrl",
DROP COLUMN "twitterUsername",
DROP COLUMN "wikiUrl";
