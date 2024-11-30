/*
  Warnings:

  - You are about to drop the `Account` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Address` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Session` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `VerificationToken` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Account" DROP CONSTRAINT "Account_userId_fkey";

-- DropForeignKey
ALTER TABLE "Session" DROP CONSTRAINT "Session_userId_fkey";

-- DropForeignKey
ALTER TABLE "Vault" DROP CONSTRAINT "Vault_ownerAddress_fkey";

-- AlterTable
ALTER TABLE "User" ADD COLUMN     "nextVaultAddress" TEXT,
ADD COLUMN     "nextVaultSalt" TEXT;

-- DropTable
DROP TABLE "Account";

-- DropTable
DROP TABLE "Address";

-- DropTable
DROP TABLE "Session";

-- DropTable
DROP TABLE "VerificationToken";

-- AddForeignKey
ALTER TABLE "Vault" ADD CONSTRAINT "Vault_ownerAddress_fkey" FOREIGN KEY ("ownerAddress") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
