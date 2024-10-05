/*
  Warnings:

  - The primary key for the `Collectible` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - A unique constraint covering the columns `[id]` on the table `Collectible` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `id` to the `Collectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "Collectible_pkey" PRIMARY KEY ("id");

-- CreateIndex
CREATE UNIQUE INDEX "Collectible_id_key" ON "Collectible"("id");
