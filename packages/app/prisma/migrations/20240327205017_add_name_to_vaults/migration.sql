/*
  Warnings:

  - Added the required column `name` to the `Vault` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Vault" ADD COLUMN     "name" TEXT NOT NULL;
