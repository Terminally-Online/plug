/*
  Warnings:

  - You are about to drop the column `implementation` on the `Socket` table. All the data in the column will be lost.
  - You are about to drop the column `salt` on the `Socket` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Socket" DROP COLUMN "implementation",
DROP COLUMN "salt",
ADD COLUMN     "deploymentAdmin" TEXT,
ADD COLUMN     "deploymentDelegate" TEXT,
ADD COLUMN     "deploymentImplementation" TEXT,
ADD COLUMN     "deploymentNonce" INTEGER;
