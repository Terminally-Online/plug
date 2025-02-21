/*
  Warnings:

  - You are about to drop the `Simulation` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Simulation" DROP CONSTRAINT "Simulation_executionId_fkey";

-- DropTable
DROP TABLE "Simulation";

-- CreateTable
CREATE TABLE "Run" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" TEXT NOT NULL,
    "result" JSONB,
    "error" TEXT,
    "errors" TEXT[] DEFAULT ARRAY[]::TEXT[],
    "gasEstimate" DOUBLE PRECISION,
    "intentId" TEXT NOT NULL,

    CONSTRAINT "Run_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Run" ADD CONSTRAINT "Run_intentId_fkey" FOREIGN KEY ("intentId") REFERENCES "Intent"("id") ON DELETE CASCADE ON UPDATE CASCADE;
