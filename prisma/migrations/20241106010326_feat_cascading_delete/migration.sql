-- DropForeignKey
ALTER TABLE "Simulation" DROP CONSTRAINT "Simulation_executionId_fkey";

-- AddForeignKey
ALTER TABLE "Simulation" ADD CONSTRAINT "Simulation_executionId_fkey" FOREIGN KEY ("executionId") REFERENCES "Execution"("id") ON DELETE CASCADE ON UPDATE CASCADE;
