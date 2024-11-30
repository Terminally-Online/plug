/*
  Warnings:

  - You are about to drop the column `actionCategoryId` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `info` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `primary` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the column `sentence` on the `Action` table. All the data in the column will be lost.
  - You are about to drop the `ActionCategory` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Execution` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `LivePlugs` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Plugs` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `WorkflowAction` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `actionName` to the `Action` table without a default value. This is not possible if the table is not empty.
  - Added the required column `categoryName` to the `Action` table without a default value. This is not possible if the table is not empty.
  - Added the required column `data` to the `Action` table without a default value. This is not possible if the table is not empty.
  - Added the required column `index` to the `Action` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_actionCategoryId_fkey";

-- DropForeignKey
ALTER TABLE "Execution" DROP CONSTRAINT "Execution_plugsId_fkey";

-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_plugsId_fkey";

-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_socketUserAddress_socketAddress_fkey";

-- DropForeignKey
ALTER TABLE "Plugs" DROP CONSTRAINT "Plugs_socketUserAddress_socketAddress_fkey";

-- DropForeignKey
ALTER TABLE "WorkflowAction" DROP CONSTRAINT "WorkflowAction_plugsId_fkey";

-- DropForeignKey
ALTER TABLE "WorkflowAction" DROP CONSTRAINT "WorkflowAction_workflowId_fkey";

-- AlterTable
ALTER TABLE "Action" DROP COLUMN "actionCategoryId",
DROP COLUMN "info",
DROP COLUMN "primary",
DROP COLUMN "sentence",
ADD COLUMN     "actionName" TEXT NOT NULL,
ADD COLUMN     "categoryName" TEXT NOT NULL,
ADD COLUMN     "data" TEXT NOT NULL,
ADD COLUMN     "index" INTEGER NOT NULL,
ADD COLUMN     "workflowId" TEXT;

-- DropTable
DROP TABLE "ActionCategory";

-- DropTable
DROP TABLE "Execution";

-- DropTable
DROP TABLE "LivePlugs";

-- DropTable
DROP TABLE "Plugs";

-- DropTable
DROP TABLE "WorkflowAction";

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE SET NULL ON UPDATE CASCADE;
