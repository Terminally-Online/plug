/*
  Warnings:

  - You are about to drop the `Socket` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "UserSocket" DROP CONSTRAINT "UserSocket_socketAddress_fkey";

-- DropTable
DROP TABLE "Socket";
