/*
  Warnings:

  - The primary key for the `LivePlugs` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `addressId` on the `LivePlugs` table. All the data in the column will be lost.
  - Added the required column `address` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `chainId` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `name` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `verifyingContract` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.
  - Added the required column `version` to the `LivePlugs` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_addressId_fkey";

-- AlterTable
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_pkey",
DROP COLUMN "addressId",
ADD COLUMN     "address" TEXT NOT NULL,
ADD COLUMN     "chainId" INTEGER NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL,
ADD COLUMN     "verifyingContract" TEXT NOT NULL,
ADD COLUMN     "version" TEXT NOT NULL,
ADD CONSTRAINT "LivePlugs_pkey" PRIMARY KEY ("plugsId", "address");

-- CreateTable
CREATE TABLE "Chain" (
    "id" TEXT NOT NULL,

    CONSTRAINT "Chain_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Vault" (
    "address" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "adminAddressId" TEXT NOT NULL,

    CONSTRAINT "Vault_pkey" PRIMARY KEY ("address")
);

-- CreateTable
CREATE TABLE "_ChainToVault" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_ChainToVault_AB_unique" ON "_ChainToVault"("A", "B");

-- CreateIndex
CREATE INDEX "_ChainToVault_B_index" ON "_ChainToVault"("B");

-- AddForeignKey
ALTER TABLE "Vault" ADD CONSTRAINT "Vault_adminAddressId_fkey" FOREIGN KEY ("adminAddressId") REFERENCES "Address"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_verifyingContract_name_version_chainId_fkey" FOREIGN KEY ("verifyingContract", "name", "version", "chainId") REFERENCES "Domain"("verifyingContract", "name", "version", "chainId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_address_fkey" FOREIGN KEY ("address") REFERENCES "Vault"("address") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_ChainToVault" ADD CONSTRAINT "_ChainToVault_A_fkey" FOREIGN KEY ("A") REFERENCES "Chain"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_ChainToVault" ADD CONSTRAINT "_ChainToVault_B_fkey" FOREIGN KEY ("B") REFERENCES "Vault"("address") ON DELETE CASCADE ON UPDATE CASCADE;
