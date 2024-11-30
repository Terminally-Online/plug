/*
  Warnings:

  - Added the required column `lastBlockIndexed` to the `Vault` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Vault" ADD COLUMN     "lastBlockIndexed" INTEGER NOT NULL;
