/*
  Warnings:

  - You are about to drop the column `protocolId` on the `Position` table. All the data in the column will be lost.
  - The primary key for the `Protocol` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Protocol` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_protocolId_fkey";

-- AlterTable
ALTER TABLE "Position" DROP COLUMN "protocolId",
ADD COLUMN     "protocolName" TEXT;

-- AlterTable
ALTER TABLE "Protocol" DROP CONSTRAINT "Protocol_pkey",
DROP COLUMN "id",
ADD CONSTRAINT "Protocol_pkey" PRIMARY KEY ("name");

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_protocolName_fkey" FOREIGN KEY ("protocolName") REFERENCES "Protocol"("name") ON DELETE SET NULL ON UPDATE CASCADE;
