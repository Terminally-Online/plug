/*
  Warnings:

  - You are about to drop the column `name` on the `Socket` table. All the data in the column will be lost.
  - Added the required column `name` to the `UserSocket` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Socket" DROP COLUMN "name";

-- AlterTable
ALTER TABLE "UserSocket" ADD COLUMN     "name" TEXT NOT NULL;
