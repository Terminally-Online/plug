/*
  Warnings:

  - Made the column `price` on table `TokenPrice` required. This step will fail if there are existing NULL values in that column.
  - Made the column `change` on table `TokenPrice` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "TokenPrice" ALTER COLUMN "price" SET NOT NULL,
ALTER COLUMN "price" SET DEFAULT 0,
ALTER COLUMN "change" SET NOT NULL,
ALTER COLUMN "change" SET DEFAULT 0;
