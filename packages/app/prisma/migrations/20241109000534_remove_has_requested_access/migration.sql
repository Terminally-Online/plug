/*
  Warnings:

  - You are about to drop the column `hasRequestedAccess` on the `SocketIdentity` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "SocketIdentity" DROP COLUMN "hasRequestedAccess";
