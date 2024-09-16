/*
  Warnings:

  - You are about to drop the `ConsoleColumn` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_socketId_fkey";

-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_viewAsId_fkey";

-- DropTable
DROP TABLE "ConsoleColumn";
