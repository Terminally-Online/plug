/*
  Warnings:

  - The primary key for the `ImplementationBalance` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - The required column `id` was added to the `ImplementationBalance` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- AlterTable
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "ImplementationBalance_pkey" PRIMARY KEY ("id");
