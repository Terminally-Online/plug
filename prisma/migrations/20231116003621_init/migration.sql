/*
  Warnings:

  - You are about to drop the `ComponentOnCanvas` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "ComponentOnCanvas" DROP CONSTRAINT "ComponentOnCanvas_canvasId_fkey";

-- DropForeignKey
ALTER TABLE "ComponentOnCanvas" DROP CONSTRAINT "ComponentOnCanvas_componentId_fkey";

-- AlterTable
ALTER TABLE "Component" ADD COLUMN     "canvasId" TEXT;

-- DropTable
DROP TABLE "ComponentOnCanvas";

-- AddForeignKey
ALTER TABLE "Component" ADD CONSTRAINT "Component_canvasId_fkey" FOREIGN KEY ("canvasId") REFERENCES "Canvas"("id") ON DELETE SET NULL ON UPDATE CASCADE;
