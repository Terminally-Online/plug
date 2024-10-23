-- CreateTable
CREATE TABLE "QueuedWorkflow" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "startDate" TIMESTAMP(3) NOT NULL,
    "endDate" TIMESTAMP(3),
    "frequency" INTEGER NOT NULL DEFAULT 0,
    "workflowId" TEXT NOT NULL,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "QueuedWorkflow_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "QueuedWorkflow_workflowId_idx" ON "QueuedWorkflow"("workflowId");

-- CreateIndex
CREATE INDEX "QueuedWorkflow_startDate_idx" ON "QueuedWorkflow"("startDate");

-- AddForeignKey
ALTER TABLE "QueuedWorkflow" ADD CONSTRAINT "QueuedWorkflow_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "QueuedWorkflow" ADD CONSTRAINT "QueuedWorkflow_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
