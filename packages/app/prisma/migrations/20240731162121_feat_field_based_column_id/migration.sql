/*
  Warnings:

  - The primary key for the `ConsoleColumn` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `ConsoleColumn` table. All the data in the column will be lost.
  - Made the column `consoleId` on table `ConsoleColumn` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_consoleId_fkey";

-- AlterTable
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_pkey",
DROP COLUMN "id",
ALTER COLUMN "consoleId" SET NOT NULL,
ADD CONSTRAINT "ConsoleColumn_pkey" PRIMARY KEY ("consoleId", "index");

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_consoleId_fkey" FOREIGN KEY ("consoleId") REFERENCES "Console"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
