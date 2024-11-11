/*
  Warnings:

  - You are about to drop the column `views` on the `Workflow` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "views",
ADD COLUMN     "totalViews" INTEGER NOT NULL DEFAULT 0;

-- CreateTable
CREATE TABLE "View" (
    "date" TIMESTAMP(3) NOT NULL,
    "views" INTEGER NOT NULL DEFAULT 0,
    "workflowId" TEXT NOT NULL,

    CONSTRAINT "View_pkey" PRIMARY KEY ("workflowId","date")
);

-- CreateIndex
CREATE INDEX "View_date_idx" ON "View"("date");

-- AddForeignKey
ALTER TABLE "View" ADD CONSTRAINT "View_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
