/*
  Warnings:

  - You are about to drop the column `selectedById` on the `Component` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Component" DROP CONSTRAINT "Component_selectedById_fkey";

-- AlterTable
ALTER TABLE "Component" DROP COLUMN "selectedById",
ADD COLUMN     "selectingId" TEXT;

-- AddForeignKey
ALTER TABLE "Component" ADD CONSTRAINT "Component_selectingId_fkey" FOREIGN KEY ("selectingId") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;
