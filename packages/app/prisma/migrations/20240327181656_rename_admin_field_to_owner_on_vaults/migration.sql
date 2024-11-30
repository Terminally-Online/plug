/*
  Warnings:

  - You are about to drop the column `adminAddressId` on the `Vault` table. All the data in the column will be lost.
  - Added the required column `ownerAddressId` to the `Vault` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_adminAddressId_fkey";

-- AlterTable
ALTER TABLE "Vault" DROP COLUMN "adminAddressId",
ADD COLUMN     "ownerAddressId" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "Vault" ADD CONSTRAINT "Vault_ownerAddressId_fkey" FOREIGN KEY ("ownerAddressId") REFERENCES "Address"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
