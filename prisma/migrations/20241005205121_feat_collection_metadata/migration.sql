/*
  Warnings:

  - Added the required column `description` to the `Collection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `iconUrl` to the `Collection` table without a default value. This is not possible if the table is not empty.
  - Added the required column `name` to the `Collection` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Collection" ADD COLUMN     "description" TEXT NOT NULL,
ADD COLUMN     "iconUrl" TEXT NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL;
