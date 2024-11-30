/*
  Warnings:

  - You are about to drop the column `workflowId` on the `Action` table. All the data in the column will be lost.
  - Added the required column `versionId` to the `Action` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_workflowId_fkey";

-- AlterTable
ALTER TABLE "Action" DROP COLUMN "workflowId",
ADD COLUMN     "versionId" TEXT NOT NULL;

-- CreateTable
CREATE TABLE "Version" (
    "id" TEXT NOT NULL,
    "version" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "workflowId" TEXT NOT NULL,

    CONSTRAINT "Version_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Version_workflowId_version_key" ON "Version"("workflowId", "version");

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_versionId_fkey" FOREIGN KEY ("versionId") REFERENCES "Version"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Version" ADD CONSTRAINT "Version_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
