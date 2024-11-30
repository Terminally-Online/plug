/*
  Warnings:

  - You are about to drop the `QueuedWorkflow` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "QueuedWorkflow" DROP CONSTRAINT "QueuedWorkflow_socketId_fkey";

-- DropForeignKey
ALTER TABLE "QueuedWorkflow" DROP CONSTRAINT "QueuedWorkflow_workflowId_fkey";

-- DropTable
DROP TABLE "QueuedWorkflow";

-- CreateTable
CREATE TABLE "Execution" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "startAt" TIMESTAMP(3) NOT NULL,
    "endAt" TIMESTAMP(3),
    "frequency" INTEGER NOT NULL DEFAULT 0,
    "nextSimulationAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "workflowId" TEXT NOT NULL,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "Execution_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "Execution_workflowId_idx" ON "Execution"("workflowId");

-- CreateIndex
CREATE INDEX "Execution_startAt_idx" ON "Execution"("startAt");

-- CreateIndex
CREATE INDEX "Execution_nextSimulationAt_idx" ON "Execution"("nextSimulationAt");

-- AddForeignKey
ALTER TABLE "Execution" ADD CONSTRAINT "Execution_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Execution" ADD CONSTRAINT "Execution_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
