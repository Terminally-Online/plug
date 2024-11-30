/*
  Warnings:

  - The primary key for the `Tag` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `name` on the `Tag` table. All the data in the column will be lost.
  - Added the required column `id` to the `Tag` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "_TagToWorkflow" DROP CONSTRAINT "_TagToWorkflow_A_fkey";

-- AlterTable
ALTER TABLE "Tag" DROP CONSTRAINT "Tag_pkey",
DROP COLUMN "name",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "Tag_pkey" PRIMARY KEY ("id");

-- AddForeignKey
ALTER TABLE "_TagToWorkflow" ADD CONSTRAINT "_TagToWorkflow_A_fkey" FOREIGN KEY ("A") REFERENCES "Tag"("id") ON DELETE CASCADE ON UPDATE CASCADE;
