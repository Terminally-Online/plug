/*
  Warnings:

  - You are about to drop the `TokenBalance` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `TokenBalanceCache` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "TokenBalance" DROP CONSTRAINT "TokenBalance_cacheSocketId_cacheChain_fkey";

-- DropForeignKey
ALTER TABLE "TokenBalanceCache" DROP CONSTRAINT "TokenBalanceCache_socketId_fkey";

-- DropTable
DROP TABLE "TokenBalance";

-- DropTable
DROP TABLE "TokenBalanceCache";

-- CreateTable
CREATE TABLE "ImplementationBalance" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "balance" DOUBLE PRECISION NOT NULL,
    "implementationChain" TEXT NOT NULL,
    "implementationContract" TEXT NOT NULL,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "ImplementationBalance_pkey" PRIMARY KEY ("socketId","implementationChain","implementationContract")
);

-- AddForeignKey
ALTER TABLE "ImplementationBalance" ADD CONSTRAINT "ImplementationBalance_implementationChain_implementationCo_fkey" FOREIGN KEY ("implementationChain", "implementationContract") REFERENCES "Implementation"("chain", "contract") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ImplementationBalance" ADD CONSTRAINT "ImplementationBalance_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
