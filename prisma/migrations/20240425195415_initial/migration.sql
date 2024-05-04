/*
  Warnings:

  - The primary key for the `LivePlugs` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `address` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `chainId` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `name` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `version` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `executor` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `maxFeePerGas` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `maxPriorityFeePerGas` on the `Plugs` table. All the data in the column will be lost.
  - You are about to drop the column `socket` on the `Plugs` table. All the data in the column will be lost.
  - The primary key for the `User` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `email` on the `User` table. All the data in the column will be lost.
  - You are about to drop the column `emailVerified` on the `User` table. All the data in the column will be lost.
  - You are about to drop the column `id` on the `User` table. All the data in the column will be lost.
  - You are about to drop the column `image` on the `User` table. All the data in the column will be lost.
  - You are about to drop the column `lastBlockIndexed` on the `Vault` table. All the data in the column will be lost.
  - You are about to drop the column `ownerAddress` on the `Vault` table. All the data in the column will be lost.
  - You are about to drop the `Canvas` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Chain` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Component` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Current` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Domain` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Fuse` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Plug` table. If the table is not empty, all the data it contains will be lost.
  - The required column `id` was added to the `LivePlugs` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.
  - Added the required column `updatedAt` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultAddress` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultChainId` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultUserAddress` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultAddress` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultChainId` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `vaultUserAddress` to the `Plugs` table without a default value. This is not possible if the table is not empty.
  - The required column `address` was added to the `User` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.
  - Added the required column `updatedAt` to the `User` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Canvas" DROP CONSTRAINT "Canvas_userId_fkey";

-- DropForeignKey
ALTER TABLE "Component" DROP CONSTRAINT "Component_canvasId_fkey";

-- DropForeignKey
ALTER TABLE "Component" DROP CONSTRAINT "Component_selectingId_fkey";

-- DropForeignKey
ALTER TABLE "Fuse" DROP CONSTRAINT "Fuse_plugId_fkey";

-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_address_chainId_fkey";

-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_address_name_version_chainId_fkey";

-- DropForeignKey
ALTER TABLE "Plug" DROP CONSTRAINT "Plug_currentName_currentType_currentData_fkey";

-- DropForeignKey
ALTER TABLE "Plug" DROP CONSTRAINT "Plug_plugsId_fkey";

-- DropForeignKey
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_chainId_fkey";

-- DropForeignKey
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_ownerAddress_fkey";

-- DropIndex
DROP INDEX "User_email_key";

-- DropIndex
DROP INDEX "Vault_address_chainId_key";

-- AlterTable
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_pkey",
DROP COLUMN "address",
DROP COLUMN "chainId",
DROP COLUMN "name",
DROP COLUMN "version",
ADD COLUMN     "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "id" TEXT NOT NULL,
ADD COLUMN     "updatedAt" TIMESTAMP(3) NOT NULL,
ADD COLUMN     "vaultAddress" TEXT NOT NULL,
ADD COLUMN     "vaultChainId" INTEGER NOT NULL,
ADD COLUMN     "vaultUserAddress" TEXT NOT NULL,
ADD CONSTRAINT "LivePlugs_pkey" PRIMARY KEY ("id");

-- AlterTable
ALTER TABLE "Plugs" DROP COLUMN "executor",
DROP COLUMN "maxFeePerGas",
DROP COLUMN "maxPriorityFeePerGas",
DROP COLUMN "socket",
ADD COLUMN     "chainId" INTEGER[],
ADD COLUMN     "vaultAddress" TEXT NOT NULL,
ADD COLUMN     "vaultChainId" INTEGER NOT NULL,
ADD COLUMN     "vaultUserAddress" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "User" DROP CONSTRAINT "User_pkey",
DROP COLUMN "email",
DROP COLUMN "emailVerified",
DROP COLUMN "id",
DROP COLUMN "image",
ADD COLUMN     "address" TEXT NOT NULL,
ADD COLUMN     "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "updatedAt" TIMESTAMP(3) NOT NULL,
ADD CONSTRAINT "User_pkey" PRIMARY KEY ("address");

-- AlterTable
ALTER TABLE "Vault" DROP COLUMN "lastBlockIndexed",
DROP COLUMN "ownerAddress",
ADD CONSTRAINT "Vault_pkey" PRIMARY KEY ("address", "chainId");

-- DropTable
DROP TABLE "Canvas";

-- DropTable
DROP TABLE "Chain";

-- DropTable
DROP TABLE "Component";

-- DropTable
DROP TABLE "Current";

-- DropTable
DROP TABLE "Domain";

-- DropTable
DROP TABLE "Fuse";

-- DropTable
DROP TABLE "Plug";

-- CreateTable
CREATE TABLE "Workflow" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL DEFAULT 'Untitled Workflow',
    "workflowForkedId" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "userAddress" TEXT,

    CONSTRAINT "Workflow_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "UserVault" (
    "userAddress" TEXT NOT NULL,
    "vaultAddress" TEXT NOT NULL,
    "vaultChainId" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "UserVault_pkey" PRIMARY KEY ("userAddress","vaultAddress","vaultChainId")
);

-- CreateTable
CREATE TABLE "Action" (
    "id" TEXT NOT NULL,
    "index" INTEGER NOT NULL,
    "target" TEXT NOT NULL,
    "value" TEXT NOT NULL,
    "data" TEXT NOT NULL,
    "workflowId" TEXT,
    "plugsId" TEXT,

    CONSTRAINT "Action_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Execution" (
    "id" TEXT NOT NULL,
    "simulation" BOOLEAN NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "plugsId" TEXT,

    CONSTRAINT "Execution_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Workflow" ADD CONSTRAINT "Workflow_userAddress_fkey" FOREIGN KEY ("userAddress") REFERENCES "User"("address") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserVault" ADD CONSTRAINT "UserVault_userAddress_fkey" FOREIGN KEY ("userAddress") REFERENCES "User"("address") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserVault" ADD CONSTRAINT "UserVault_vaultAddress_vaultChainId_fkey" FOREIGN KEY ("vaultAddress", "vaultChainId") REFERENCES "Vault"("address", "chainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES "Workflow"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Action" ADD CONSTRAINT "Action_plugsId_fkey" FOREIGN KEY ("plugsId") REFERENCES "Plugs"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plugs" ADD CONSTRAINT "Plugs_vaultUserAddress_vaultAddress_vaultChainId_fkey" FOREIGN KEY ("vaultUserAddress", "vaultAddress", "vaultChainId") REFERENCES "UserVault"("userAddress", "vaultAddress", "vaultChainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_vaultUserAddress_vaultAddress_vaultChainId_fkey" FOREIGN KEY ("vaultUserAddress", "vaultAddress", "vaultChainId") REFERENCES "UserVault"("userAddress", "vaultAddress", "vaultChainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Execution" ADD CONSTRAINT "Execution_plugsId_fkey" FOREIGN KEY ("plugsId") REFERENCES "Plugs"("id") ON DELETE SET NULL ON UPDATE CASCADE;
