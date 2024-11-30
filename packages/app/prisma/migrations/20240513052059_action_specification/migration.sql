/*
  Warnings:

  - You are about to drop the column `data` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `index` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `plugsId` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `target` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `value` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `workflowId` on the `Action` table. All the data in the column will be lost.
  - Added the required column `actionCategoryId` to the `Action` table without a default value. This is not possible if the table is not empty.
  - Added the required column `name` to the `Action` table without a default value. This is not possible if the table is not empty.
  - Added the required column `sentence` to the `Action` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_plugsId_fkey";

-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_workflowId_fkey";

-- AlterTable
ALTER TABLE "Action" DROP COLUMN "data",
DROP COLUMN "index",
DROP COLUMN "plugsId",
DROP COLUMN "target",
DROP COLUMN "value",
DROP COLUMN "workflowId",
ADD COLUMN     "actionCategoryId" TEXT NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL,
ADD COLUMN     "sentence" TEXT NOT NULL;

-- CreateTable
CREATE TABLE "ActionCategory" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "image" TEXT NOT NULL,
    "gradientFrom" TEXT NOT NULL,
    "gradientTo" TEXT NOT NULL,

    CONSTRAINT "ActionCategory_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "WorkflowAction" (
    "id" TEXT NOT NULL,
    "index" INTEGER NOT NULL,
    "target" TEXT NOT NULL,
    "value" TEXT NOT NULL,
    "data" TEXT NOT NULL,
    "workflowId" TEXT,
    "plugsId" TEXT,

    CONSTRAINT "WorkflowAction_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_actionCategoryId_fkey" FOREIGN KEY ("actionCategoryId") REFERENCES "ActionCategory"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "WorkflowAction" ADD CONSTRAINT "WorkflowAction_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "WorkflowAction" ADD CONSTRAINT "WorkflowAction_plugsId_fkey" FOREIGN KEY ("plugsId") REFERENCES "Plugs"("id") ON DELETE SET NULL ON UPDATE CASCADE;
