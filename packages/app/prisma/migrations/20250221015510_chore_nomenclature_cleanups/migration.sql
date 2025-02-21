/*
  Warnings:

  - You are about to drop the column `workflowForkedId` on the `Plug` table. All the data in the column will be lost.
  - The primary key for the `View` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `workflowId` on the `View` table. All the data in the column will be lost.
  - Added the required column `plugId` to the `View` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "View" DROP CONSTRAINT "View_workflowId_fkey";

-- AlterTable
ALTER TABLE "Plug" DROP COLUMN "workflowForkedId",
ADD COLUMN     "plugForkedId" TEXT;

-- AlterTable
ALTER TABLE "View" DROP CONSTRAINT "View_pkey",
DROP COLUMN "workflowId",
ADD COLUMN     "plugId" TEXT NOT NULL,
ADD CONSTRAINT "View_pkey" PRIMARY KEY ("plugId", "date");

-- AddForeignKey
ALTER TABLE "View" ADD CONSTRAINT "View_plugId_fkey" FOREIGN KEY ("plugId") REFERENCES "Plug"("id") ON DELETE CASCADE ON UPDATE CASCADE;
