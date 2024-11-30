/*
  Warnings:

  - You are about to drop the `Identifier` table. If the table is not empty, all the data it contains will be lost.
  - A unique constraint covering the columns `[id]` on the table `User` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "Identifier" DROP CONSTRAINT "Identifier_userId_fkey";

-- DropTable
DROP TABLE "Identifier";

-- CreateIndex
CREATE UNIQUE INDEX "User_id_key" ON "User"("id");
