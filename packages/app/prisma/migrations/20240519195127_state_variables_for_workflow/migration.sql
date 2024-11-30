/*
  Warnings:

  - You are about to drop the column `curated` on the `Workflow` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "curated",
ADD COLUMN     "isCurated" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "isPrivate" BOOLEAN NOT NULL DEFAULT false;
