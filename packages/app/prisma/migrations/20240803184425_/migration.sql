/*
  Warnings:

  - You are about to drop the `Price` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropTable
DROP TABLE "Price";

-- CreateTable
CREATE TABLE "TokenPrice" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "chain" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "decimals" INTEGER,
    "symbol" TEXT,
    "price" DOUBLE PRECISION,
    "timestamp" INTEGER,
    "confidence" DOUBLE PRECISION,
    "change" DOUBLE PRECISION,

    CONSTRAINT "TokenPrice_pkey" PRIMARY KEY ("id")
);
