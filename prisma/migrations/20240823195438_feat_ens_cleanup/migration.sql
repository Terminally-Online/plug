/*
  Warnings:

  - The primary key for the `ENS` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `ensAvatar` on the `ENS` table. All the data in the column will be lost.
  - You are about to drop the column `ensName` on the `ENS` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[name]` on the table `ENS` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `name` to the `ENS` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_ensName_fkey";

-- DropIndex
DROP INDEX "ENS_ensName_key";

-- AlterTable
ALTER TABLE "ENS" DROP CONSTRAINT "ENS_pkey",
DROP COLUMN "ensAvatar",
DROP COLUMN "ensName",
ADD COLUMN     "avatar" TEXT,
ADD COLUMN     "name" TEXT NOT NULL,
ADD CONSTRAINT "ENS_pkey" PRIMARY KEY ("name");

-- CreateIndex
CREATE UNIQUE INDEX "ENS_name_key" ON "ENS"("name");

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_ensName_fkey" FOREIGN KEY ("ensName") REFERENCES "ENS"("name") ON DELETE SET NULL ON UPDATE CASCADE;
