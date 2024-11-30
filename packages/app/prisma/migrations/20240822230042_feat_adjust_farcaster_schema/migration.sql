/*
  Warnings:

  - You are about to drop the column `avatar` on the `FarcasterUser` table. All the data in the column will be lost.
  - You are about to drop the column `displayName` on the `FarcasterUser` table. All the data in the column will be lost.
  - You are about to drop the column `fid` on the `FarcasterUser` table. All the data in the column will be lost.
  - You are about to drop the column `username` on the `FarcasterUser` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "FarcasterUser_fid_key";

-- DropIndex
DROP INDEX "FarcasterUser_username_key";

-- AlterTable
ALTER TABLE "FarcasterUser" DROP COLUMN "avatar",
DROP COLUMN "displayName",
DROP COLUMN "fid",
DROP COLUMN "username";

-- CreateTable
CREATE TABLE "FarcasterUserAddress" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "FarcasterUserAddress_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_FarcasterAddresses" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_FarcasterAddresses_AB_unique" ON "_FarcasterAddresses"("A", "B");

-- CreateIndex
CREATE INDEX "_FarcasterAddresses_B_index" ON "_FarcasterAddresses"("B");

-- AddForeignKey
ALTER TABLE "_FarcasterAddresses" ADD CONSTRAINT "_FarcasterAddresses_A_fkey" FOREIGN KEY ("A") REFERENCES "FarcasterUser"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_FarcasterAddresses" ADD CONSTRAINT "_FarcasterAddresses_B_fkey" FOREIGN KEY ("B") REFERENCES "FarcasterUserAddress"("id") ON DELETE CASCADE ON UPDATE CASCADE;
