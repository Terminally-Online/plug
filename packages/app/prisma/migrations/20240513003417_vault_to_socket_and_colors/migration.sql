/*
  Warnings:

  - You are about to drop the column `vaultAddress` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `vaultChainId` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `vaultUserAddress` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `vaultAddress` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `vaultChainId` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `vaultUserAddress` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `nextVaultAddress` on the `User` table. All the data in the column will be lost.
  - You are about to drop the column `nextVaultSalt` on the `User` table. All the data in the column will be lost.
  - You are about to drop the `UserVault` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Vault` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `socketAddress` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketChainId` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketUserAddress` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketAddress` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketChainId` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `socketUserAddress` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `color` to the `Workflow` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_vaultUserAddress_vaultAddress_vaultChainId_fkey";

-- DropForeignKey
ALTER TABLE "Plugs" DROP CONSTRAINT "Plugs_vaultUserAddress_vaultAddress_vaultChainId_fkey";

-- DropForeignKey
ALTER TABLE "UserVault" DROP CONSTRAINT "UserVault_userAddress_fkey";

-- DropForeignKey
ALTER TABLE "UserVault" DROP CONSTRAINT "UserVault_vaultAddress_vaultChainId_fkey";

-- AlterTable
ALTER TABLE "LivePlugs" DROP COLUMN "vaultAddress",
DROP COLUMN "vaultChainId",
DROP COLUMN "vaultUserAddress",
ADD COLUMN     "socketAddress" TEXT NOT NULL,
ADD COLUMN     "socketChainId" INTEGER NOT NULL,
ADD COLUMN     "socketUserAddress" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "Plugs" DROP COLUMN "vaultAddress",
DROP COLUMN "vaultChainId",
DROP COLUMN "vaultUserAddress",
ADD COLUMN     "socketAddress" TEXT NOT NULL,
ADD COLUMN     "socketChainId" INTEGER NOT NULL,
ADD COLUMN     "socketUserAddress" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "User" DROP COLUMN "nextVaultAddress",
DROP COLUMN "nextVaultSalt",
ADD COLUMN     "nextSocketAddress" TEXT,
ADD COLUMN     "nextSocketSalt" TEXT;

-- AlterTable
ALTER TABLE "Workflow" ADD COLUMN     "color" TEXT NOT NULL,
ALTER COLUMN "name" SET DEFAULT 'Untitled Plug';

-- DropTable
DROP TABLE "UserVault";

-- DropTable
DROP TABLE "Vault";

-- CreateTable
CREATE TABLE "Socket" (
    "address" TEXT NOT NULL,
    "chainId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "Socket_pkey" PRIMARY KEY ("address","chainId")
);

-- CreateTable
CREATE TABLE "UserSocket" (
    "userAddress" TEXT NOT NULL,
    "socketAddress" TEXT NOT NULL,
    "socketChainId" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "UserSocket_pkey" PRIMARY KEY ("userAddress","socketAddress","socketChainId")
);

-- AddForeignKey
ALTER TABLE "UserSocket" ADD CONSTRAINT "UserSocket_userAddress_fkey" FOREIGN KEY ("userAddress") REFERENCES "User"("address") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserSocket" ADD CONSTRAINT "UserSocket_socketAddress_socketChainId_fkey" FOREIGN KEY ("socketAddress", "socketChainId") REFERENCES "Socket"("address", "chainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plugs" ADD CONSTRAINT "Plugs_socketUserAddress_socketAddress_socketChainId_fkey" FOREIGN KEY ("socketUserAddress", "socketAddress", "socketChainId") REFERENCES "UserSocket"("userAddress", "socketAddress", "socketChainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_socketUserAddress_socketAddress_socketChainId_fkey" FOREIGN KEY ("socketUserAddress", "socketAddress", "socketChainId") REFERENCES "UserSocket"("userAddress", "socketAddress", "socketChainId") ON DELETE RESTRICT ON UPDATE CASCADE;
