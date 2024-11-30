/*
  Warnings:

  - You are about to drop the column `socketChainId` on the `LivePlugs` table. All the data in the column will be lost.
  - You are about to drop the column `socketChainId` on the `Plugs` table. All the data in the column will be lost.
  - The primary key for the `Socket` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `chainId` on the `Socket` table. All the data in the column will be lost.
  - The primary key for the `UserSocket` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `socketChainId` on the `UserSocket` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "LivePlugs" DROP CONSTRAINT "LivePlugs_socketUserAddress_socketAddress_socketChainId_fkey";

-- DropForeignKey
ALTER TABLE "Plugs" DROP CONSTRAINT "Plugs_socketUserAddress_socketAddress_socketChainId_fkey";

-- DropForeignKey
ALTER TABLE "UserSocket" DROP CONSTRAINT "UserSocket_socketAddress_socketChainId_fkey";

-- AlterTable
ALTER TABLE "LivePlugs" DROP COLUMN "socketChainId";

-- AlterTable
ALTER TABLE "Plugs" DROP COLUMN "socketChainId";

-- AlterTable
ALTER TABLE "Socket" DROP CONSTRAINT "Socket_pkey",
DROP COLUMN "chainId",
ADD CONSTRAINT "Socket_pkey" PRIMARY KEY ("address");

-- AlterTable
ALTER TABLE "UserSocket" DROP CONSTRAINT "UserSocket_pkey",
DROP COLUMN "socketChainId",
ADD CONSTRAINT "UserSocket_pkey" PRIMARY KEY ("userAddress", "socketAddress");

-- AddForeignKey
ALTER TABLE "UserSocket" ADD CONSTRAINT "UserSocket_socketAddress_fkey" FOREIGN KEY ("socketAddress") REFERENCES "Socket"("address") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plugs" ADD CONSTRAINT "Plugs_socketUserAddress_socketAddress_fkey" FOREIGN KEY ("socketUserAddress", "socketAddress") REFERENCES "UserSocket"("userAddress", "socketAddress") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_socketUserAddress_socketAddress_fkey" FOREIGN KEY ("socketUserAddress", "socketAddress") REFERENCES "UserSocket"("userAddress", "socketAddress") ON DELETE RESTRICT ON UPDATE CASCADE;
