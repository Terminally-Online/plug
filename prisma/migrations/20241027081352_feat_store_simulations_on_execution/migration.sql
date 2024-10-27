/*
  Warnings:

  - You are about to drop the column `workflowId` on the `Simulation` table. All the data in the column will be lost.
  - Added the required column `executionId` to the `Simulation` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Simulation" DROP CONSTRAINT "Simulation_workflowId_fkey";

-- DropIndex
DROP INDEX "Simulation_status_idx";

-- DropIndex
DROP INDEX "Simulation_workflowId_idx";

-- AlterTable
ALTER TABLE "Simulation" DROP COLUMN "workflowId",
ADD COLUMN     "executionId" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "Simulation" ADD CONSTRAINT "Simulation_executionId_fkey" FOREIGN KEY ("executionId") REFERENCES "Execution"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
