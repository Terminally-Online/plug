/*
  Warnings:

  - You are about to drop the column `endAt` on the `QueuedWorkflow` table. All the data in the column will be lost.
  - You are about to drop the column `nextSimulationAt` on the `QueuedWorkflow` table. All the data in the column will be lost.
  - You are about to drop the column `startAt` on the `QueuedWorkflow` table. All the data in the column will be lost.
  - Added the required column `startDate` to the `QueuedWorkflow` table without a default value. This is not possible if the table is not empty.
  - Made the column `frequency` on table `QueuedWorkflow` required. This step will fail if there are existing NULL values in that column.

*/
-- DropIndex
DROP INDEX "QueuedWorkflow_nextSimulationAt_idx";

-- DropIndex
DROP INDEX "QueuedWorkflow_startAt_idx";

-- AlterTable
ALTER TABLE "QueuedWorkflow" DROP COLUMN "endAt",
DROP COLUMN "nextSimulationAt",
DROP COLUMN "startAt",
ADD COLUMN     "endDate" TIMESTAMP(3),
ADD COLUMN     "startDate" TIMESTAMP(3) NOT NULL,
ALTER COLUMN "frequency" SET NOT NULL,
ALTER COLUMN "frequency" SET DEFAULT 0;

-- AlterTable
ALTER TABLE "Workflow" ADD COLUMN     "nextSimulationAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- CreateIndex
CREATE INDEX "QueuedWorkflow_startDate_idx" ON "QueuedWorkflow"("startDate");

-- CreateIndex
CREATE INDEX "Workflow_nextSimulationAt_idx" ON "Workflow"("nextSimulationAt");
