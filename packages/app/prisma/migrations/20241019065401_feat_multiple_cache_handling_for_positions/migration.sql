/*
  Warnings:

  - The primary key for the `PositionCache` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - The required column `id` was added to the `PositionCache` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_cacheId_fkey";

-- AlterTable
ALTER TABLE "PositionCache" DROP CONSTRAINT "PositionCache_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "PositionCache_pkey" PRIMARY KEY ("id");

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "PositionCache"("id") ON DELETE CASCADE ON UPDATE CASCADE;
