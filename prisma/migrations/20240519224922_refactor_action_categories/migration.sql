/*
  Warnings:

  - You are about to drop the column `name` on the `ActionCategory` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_actionCategoryId_fkey";

-- AlterTable
ALTER TABLE "Action" ALTER COLUMN "actionCategoryId" DROP NOT NULL;

-- AlterTable
ALTER TABLE "ActionCategory" DROP COLUMN "name";

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_actionCategoryId_fkey" FOREIGN KEY ("actionCategoryId") REFERENCES "ActionCategory"("id") ON DELETE SET NULL ON UPDATE CASCADE;
