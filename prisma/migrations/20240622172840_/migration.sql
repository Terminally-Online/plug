/*
  Warnings:

  - You are about to drop the `Action` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Version` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_versionId_fkey";

-- DropForeignKey
ALTER TABLE "Version" DROP CONSTRAINT "Version_workflowId_fkey";

-- AlterTable
ALTER TABLE "Workflow" ADD COLUMN     "actions" TEXT NOT NULL DEFAULT '[]';

-- DropTable
DROP TABLE "Action";

-- DropTable
DROP TABLE "Version";
