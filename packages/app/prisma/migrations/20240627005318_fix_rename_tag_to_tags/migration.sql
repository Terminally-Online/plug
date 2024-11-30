/*
  Warnings:

  - You are about to drop the column `tag` on the `Workflow` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "tag",
ADD COLUMN     "tags" JSONB DEFAULT '[]';
