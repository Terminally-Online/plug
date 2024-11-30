/*
  Warnings:

  - Made the column `referralCode` on table `SocketIdentity` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "SocketIdentity" ALTER COLUMN "referralCode" SET NOT NULL;
