-- AlterTable
ALTER TABLE "Simulation" ADD COLUMN     "errors" TEXT[] DEFAULT ARRAY[]::TEXT[];
