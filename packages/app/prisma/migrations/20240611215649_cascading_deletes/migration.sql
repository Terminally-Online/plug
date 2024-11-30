-- DropForeignKey
ALTER TABLE "Action" DROP CONSTRAINT "Action_versionId_fkey";

-- DropForeignKey
ALTER TABLE "Version" DROP CONSTRAINT "Version_workflowId_fkey";

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_versionId_fkey" FOREIGN KEY ("versionId") REFERENCES "Version"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Version" ADD CONSTRAINT "Version_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE CASCADE ON UPDATE CASCADE;
