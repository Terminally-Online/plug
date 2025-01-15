-- AlterTable
ALTER TABLE "Execution" ADD COLUMN     "periodEndAt" TIMESTAMP(3),
ADD COLUMN     "periodStartAt" TIMESTAMP(3),
ALTER COLUMN "nextSimulationAt" DROP NOT NULL,
ALTER COLUMN "nextSimulationAt" DROP DEFAULT;
