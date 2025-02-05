-- AlterTable
ALTER TABLE "UserSocket" ADD COLUMN     "onboardingColor" TEXT,
ADD COLUMN     "onboardingCount" INTEGER NOT NULL DEFAULT 0;
