-- AlterTable
ALTER TABLE "Companion" ALTER COLUMN "lastFeedAt" DROP NOT NULL,
ALTER COLUMN "lastFeedAt" DROP DEFAULT;
