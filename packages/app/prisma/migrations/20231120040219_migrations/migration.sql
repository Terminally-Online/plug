-- CreateEnum
CREATE TYPE "ComponentType" AS ENUM ('MARKDOWN', 'BOX', 'PLUG');

-- AlterTable
ALTER TABLE "Component" ADD COLUMN     "type" "ComponentType" NOT NULL DEFAULT 'PLUG';
