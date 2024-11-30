/*
  Warnings:

  - You are about to drop the column `consoleId` on the `ConsoleColumn` table. All the data in the column will be lost.
  - You are about to drop the `Console` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_consoleId_fkey";

-- AlterTable
ALTER TABLE "ConsoleColumn" DROP COLUMN "consoleId",
ADD COLUMN     "socketAddress" TEXT,
ADD COLUMN     "userAddress" TEXT;

-- DropTable
DROP TABLE "Console";

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_userAddress_socketAddress_fkey" FOREIGN KEY ("userAddress", "socketAddress") REFERENCES "UserSocket"("userAddress", "socketAddress") ON DELETE SET NULL ON UPDATE CASCADE;
