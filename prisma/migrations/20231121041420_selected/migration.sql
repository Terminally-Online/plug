/*
  Warnings:

  - Added the required column `selectedById` to the `Component` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Component" ADD COLUMN     "selectedById" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "Component" ADD CONSTRAINT "Component_selectedById_fkey" FOREIGN KEY ("selectedById") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
