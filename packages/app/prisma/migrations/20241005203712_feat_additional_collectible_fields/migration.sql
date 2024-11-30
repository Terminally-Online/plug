/*
  Warnings:

  - Added the required column `interface` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `isSpam` to the `Collectible` table without a default value. This is not possible if the table is not empty.
  - Added the required column `name` to the `Collectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Collectible" ADD COLUMN     "imageUrl" TEXT,
ADD COLUMN     "interface" TEXT NOT NULL,
ADD COLUMN     "isSpam" BOOLEAN NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL,
ADD COLUMN     "previewUrl" TEXT,
ADD COLUMN     "videoUrl" TEXT;
