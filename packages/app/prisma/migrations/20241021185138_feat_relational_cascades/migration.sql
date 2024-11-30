-- DropForeignKey
ALTER TABLE "Collectible" DROP CONSTRAINT "Collectible_cacheId_fkey";

-- DropForeignKey
ALTER TABLE "Workflow" DROP CONSTRAINT "Workflow_socketId_fkey";

-- AddForeignKey
ALTER TABLE "Workflow" ADD CONSTRAINT "Workflow_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Collectible" ADD CONSTRAINT "Collectible_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "CollectibleCache"("id") ON DELETE CASCADE ON UPDATE CASCADE;
