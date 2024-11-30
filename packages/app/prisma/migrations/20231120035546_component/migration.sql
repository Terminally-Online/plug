/*
  Warnings:

  - Added the required column `height` to the `Component` table without a default value. This is not possible if the table is not empty.
  - Added the required column `left` to the `Component` table without a default value. This is not possible if the table is not empty.
  - Added the required column `top` to the `Component` table without a default value. This is not possible if the table is not empty.
  - Added the required column `width` to the `Component` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Component" ADD COLUMN     "height" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "left" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "top" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "width" DOUBLE PRECISION NOT NULL;
