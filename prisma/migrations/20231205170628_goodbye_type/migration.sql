/*
  Warnings:

  - You are about to drop the column `type` on the `Component` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Component" DROP COLUMN "type";

-- DropEnum
DROP TYPE "ComponentType";
