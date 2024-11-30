/*
  Warnings:

  - You are about to drop the `Tag` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `_TagToWorkflow` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "_TagToWorkflow" DROP CONSTRAINT "_TagToWorkflow_A_fkey";

-- DropForeignKey
ALTER TABLE "_TagToWorkflow" DROP CONSTRAINT "_TagToWorkflow_B_fkey";

-- DropTable
DROP TABLE "Tag";

-- DropTable
DROP TABLE "_TagToWorkflow";
