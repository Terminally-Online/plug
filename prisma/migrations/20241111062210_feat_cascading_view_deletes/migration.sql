-- DropForeignKey
ALTER TABLE "View" DROP CONSTRAINT "View_workflowId_fkey";

-- AddForeignKey
ALTER TABLE "View" ADD CONSTRAINT "View_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE CASCADE ON UPDATE CASCADE;
