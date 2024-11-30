/*
  Warnings:

  - You are about to drop the column `endDate` on the `QueuedWorkflow` table. All the data in the column will be lost.
  - You are about to drop the column `startDate` on the `QueuedWorkflow` table. All the data in the column will be lost.
  - You are about to drop the column `nextSimulationAt` on the `Workflow` table. All the data in the column will be lost.
  - Added the required column `startAt` to the `QueuedWorkflow` table without a default value. This is not possible if the table is not empty.

*/
-- DropIndex
DROP INDEX "QueuedWorkflow_startDate_idx";

-- DropIndex
DROP INDEX "Workflow_nextSimulationAt_idx";

-- AlterTable
ALTER TABLE "QueuedWorkflow" DROP COLUMN "endDate",
DROP COLUMN "startDate",
ADD COLUMN     "endAt" TIMESTAMP(3),
ADD COLUMN     "nextSimulationAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "startAt" TIMESTAMP(3) NOT NULL;

-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "nextSimulationAt";

-- CreateIndex
CREATE INDEX "QueuedWorkflow_startAt_idx" ON "QueuedWorkflow"("startAt");

-- CreateIndex
CREATE INDEX "QueuedWorkflow_nextSimulationAt_idx" ON "QueuedWorkflow"("nextSimulationAt");
