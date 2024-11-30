/*
  Warnings:

  - You are about to drop the column `socketId` on the `FarcasterUser` table. All the data in the column will be lost.
  - You are about to drop the column `ensAvatar` on the `UserSocket` table. All the data in the column will be lost.
  - You are about to drop the column `ensName` on the `UserSocket` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "FarcasterUser" DROP CONSTRAINT "FarcasterUser_socketId_fkey";

-- DropIndex
DROP INDEX "FarcasterUser_socketId_key";

-- AlterTable
ALTER TABLE "FarcasterUser" DROP COLUMN "socketId";

-- AlterTable
ALTER TABLE "UserSocket" DROP COLUMN "ensAvatar",
DROP COLUMN "ensName";

-- CreateTable
CREATE TABLE "ENS" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "ensName" TEXT NOT NULL,
    "ensAvatar" TEXT,

    CONSTRAINT "ENS_pkey" PRIMARY KEY ("ensName")
);

-- CreateTable
CREATE TABLE "SocketIdentity" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "socketId" TEXT NOT NULL,
    "farcasterUserId" TEXT,
    "ensName" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "ENS_ensName_key" ON "ENS"("ensName");

-- CreateIndex
CREATE UNIQUE INDEX "SocketIdentity_socketId_key" ON "SocketIdentity"("socketId");

-- CreateIndex
CREATE UNIQUE INDEX "SocketIdentity_farcasterUserId_key" ON "SocketIdentity"("farcasterUserId");

-- CreateIndex
CREATE UNIQUE INDEX "SocketIdentity_ensName_key" ON "SocketIdentity"("ensName");

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_farcasterUserId_fkey" FOREIGN KEY ("farcasterUserId") REFERENCES "FarcasterUser"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_ensName_fkey" FOREIGN KEY ("ensName") REFERENCES "ENS"("ensName") ON DELETE RESTRICT ON UPDATE CASCADE;
