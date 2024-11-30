-- AlterTable
ALTER TABLE "SocketIdentity" ADD COLUMN     "approvedAt" TIMESTAMP(3),
ADD COLUMN     "hasRequestedAccess" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "referredBy" TEXT,
ADD COLUMN     "requestedAt" TIMESTAMP(3);
