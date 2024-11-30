/*
  Warnings:

  - You are about to drop the column `completedAt` on the `Simulation` table. All the data in the column will be lost.
  - You are about to drop the column `startedAt` on the `Simulation` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Execution" ALTER COLUMN "status" SET DEFAULT 'active';

-- AlterTable
ALTER TABLE "Simulation" DROP COLUMN "completedAt",
DROP COLUMN "startedAt";
