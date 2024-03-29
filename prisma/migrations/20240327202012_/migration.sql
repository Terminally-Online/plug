/*
  Warnings:

  - The primary key for the `Chain` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `verifyingContract` on the `LivePlugs` table. All the data in the column will be lost.
  - The primary key for the `Vault` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `ownerAddressId` on the `Vault` table. All the data in the column will be lost.
  - You are about to drop the `_ChainToVault` table. If the table is not empty, all the data it contains will be lost.
  - A unique constraint covering the columns `[address,chainId]` on the table `Vault` will be added. If there are existing duplicate values, this will fail.
  - Changed the type of `id` on the `Chain` table. No cast exists, the column would be dropped and recreated, which cannot be done if there is data, since the column is required.
  - Added the required column `chainId` to the `Vault` table without a default value. This is not possible if the table is not empty.
  - Added the required column `ownerAddress` to the `Vault` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_address_fkey";

-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_verifyingContract_name_version_chainId_fkey";

-- DropForeignKey
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_ownerAddressId_fkey";

-- DropForeignKey
ALTER TABLE "_ChainToVault" DROP CONSTRAINT "_ChainToVault_A_fkey";

-- DropForeignKey
ALTER TABLE "_ChainToVault" DROP CONSTRAINT "_ChainToVault_B_fkey";

-- AlterTable
ALTER TABLE "Chain" DROP CONSTRAINT "Chain_pkey",
DROP COLUMN "id",
ADD COLUMN     "id" INTEGER NOT NULL,
ADD CONSTRAINT "Chain_pkey" PRIMARY KEY ("id");

-- AlterTable
ALTER TABLE "LivePlugs" DROP COLUMN "verifyingContract";

-- AlterTable
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_pkey",
DROP COLUMN "ownerAddressId",
ADD COLUMN     "chainId" INTEGER NOT NULL,
ADD COLUMN     "ownerAddress" TEXT NOT NULL;

-- DropTable
DROP TABLE "_ChainToVault";

-- CreateIndex
CREATE UNIQUE INDEX "Vault_address_chainId_key" ON "Vault"("address", "chainId");

-- AddForeignKey
ALTER TABLE "Vault" ADD CONSTRAINT "Vault_chainId_fkey" FOREIGN KEY ("chainId") REFERENCES "Chain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Vault" ADD CONSTRAINT "Vault_ownerAddress_fkey" FOREIGN KEY ("ownerAddress") REFERENCES "Address"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_address_name_version_chainId_fkey" FOREIGN KEY ("address", "name", "version", "chainId") REFERENCES "Domain"("verifyingContract", "name", "version", "chainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_address_chainId_fkey" FOREIGN KEY ("address", "chainId") REFERENCES "Vault"("address", "chainId") ON DELETE RESTRICT ON UPDATE CASCADE;
