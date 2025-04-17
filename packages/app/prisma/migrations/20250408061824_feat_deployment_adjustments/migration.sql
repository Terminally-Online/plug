/*
  Warnings:

  - You are about to drop the column `deploymentAdmin` on the `Socket` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Socket" DROP COLUMN "deploymentAdmin",
ADD COLUMN     "deploymentSalt" TEXT;
