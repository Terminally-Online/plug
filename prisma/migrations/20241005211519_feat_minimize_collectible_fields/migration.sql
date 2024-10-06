/*
  Warnings:

  - You are about to drop the column `imageUrl` on the `Collectible` table. All the data in the column will be lost.
  - You are about to drop the column `interface` on the `Collectible` table. All the data in the column will be lost.
  - You are about to drop the column `isSpam` on the `Collectible` table. All the data in the column will be lost.
  - You are about to drop the column `previewUrl` on the `Collectible` table. All the data in the column will be lost.
  - You are about to drop the column `videoUrl` on the `Collectible` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Collectible" DROP COLUMN "imageUrl",
DROP COLUMN "interface",
DROP COLUMN "isSpam",
DROP COLUMN "previewUrl",
DROP COLUMN "videoUrl";
