/*
  Warnings:

  - You are about to drop the column `onboardingColor` on the `UserSocket` table. All the data in the column will be lost.
  - You are about to drop the column `onboardingCount` on the `UserSocket` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "SocketIdentity" ADD COLUMN     "onboardingColor" TEXT,
ADD COLUMN     "onboardingCount" INTEGER NOT NULL DEFAULT 0;

-- AlterTable
ALTER TABLE "UserSocket" DROP COLUMN "onboardingColor",
DROP COLUMN "onboardingCount";
