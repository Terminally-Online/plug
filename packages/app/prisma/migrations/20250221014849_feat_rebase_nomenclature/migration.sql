/*
  Warnings:

  - You are about to drop the `Execution` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `FeatureRequest` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `TokenPrice` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `UserSocket` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Workflow` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "CollectibleCache" DROP CONSTRAINT "CollectibleCache_socketId_fkey";

-- DropForeignKey
ALTER TABLE "Execution" DROP CONSTRAINT "Execution_workflowId_fkey";

-- DropForeignKey
ALTER TABLE "PositionCache" DROP CONSTRAINT "PositionCache_socketId_fkey";

-- DropForeignKey
ALTER TABLE "Simulation" DROP CONSTRAINT "Simulation_executionId_fkey";

-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_referrerId_fkey";

-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_socketId_fkey";

-- DropForeignKey
ALTER TABLE "View" DROP CONSTRAINT "View_workflowId_fkey";

-- DropForeignKey
ALTER TABLE "Workflow" DROP CONSTRAINT "Workflow_socketId_fkey";

-- DropTable
DROP TABLE "Execution";

-- DropTable
DROP TABLE "FeatureRequest";

-- DropTable
DROP TABLE "TokenPrice";

-- DropTable
DROP TABLE "UserSocket";

-- DropTable
DROP TABLE "Workflow";

-- CreateTable
CREATE TABLE "Socket" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "admin" BOOLEAN NOT NULL DEFAULT false,
    "socketAddress" TEXT NOT NULL,
    "salt" TEXT,
    "implementation" TEXT,

    CONSTRAINT "Socket_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Plug" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "namedAt" TIMESTAMP(3),
    "renamedAt" TIMESTAMP(3),
    "name" TEXT NOT NULL DEFAULT 'Untitled',
    "isCurated" BOOLEAN NOT NULL DEFAULT false,
    "isPrivate" BOOLEAN NOT NULL DEFAULT false,
    "actions" TEXT NOT NULL DEFAULT '[]',
    "color" TEXT NOT NULL,
    "tags" TEXT[],
    "workflowForkedId" TEXT,
    "frequency" INTEGER NOT NULL DEFAULT 10,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "Plug_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Intent" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" TEXT NOT NULL DEFAULT 'active',
    "chainId" INTEGER NOT NULL DEFAULT 1,
    "actions" TEXT NOT NULL DEFAULT '[]',
    "startAt" TIMESTAMP(3) NOT NULL,
    "endAt" TIMESTAMP(3),
    "frequency" INTEGER NOT NULL DEFAULT 0,
    "periodEndAt" TIMESTAMP(3),
    "nextSimulationAt" TIMESTAMP(3),
    "plugId" TEXT NOT NULL,

    CONSTRAINT "Intent_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Price" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "chain" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "price" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "change" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "decimals" INTEGER,
    "symbol" TEXT,
    "timestamp" INTEGER,
    "confidence" DOUBLE PRECISION,

    CONSTRAINT "Price_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "Socket_socketAddress_idx" ON "Socket"("socketAddress");

-- CreateIndex
CREATE INDEX "Intent_plugId_idx" ON "Intent"("plugId");

-- CreateIndex
CREATE INDEX "Intent_startAt_idx" ON "Intent"("startAt");

-- CreateIndex
CREATE INDEX "Intent_nextSimulationAt_idx" ON "Intent"("nextSimulationAt");

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "Socket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_referrerId_fkey" FOREIGN KEY ("referrerId") REFERENCES "Socket"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "View" ADD CONSTRAINT "View_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Plug"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plug" ADD CONSTRAINT "Plug_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "Socket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Intent" ADD CONSTRAINT "Intent_plugId_fkey" FOREIGN KEY ("plugId") REFERENCES "Plug"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Simulation" ADD CONSTRAINT "Simulation_executionId_fkey" FOREIGN KEY ("executionId") REFERENCES "Intent"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PositionCache" ADD CONSTRAINT "PositionCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "Socket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "CollectibleCache" ADD CONSTRAINT "CollectibleCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "Socket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
