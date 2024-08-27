-- DropForeignKey
ALTER TABLE "OpenseaCollectible" DROP CONSTRAINT "OpenseaCollectible_cacheSocketId_cacheChain_fkey";

-- DropForeignKey
ALTER TABLE "OpenseaCollectibleCache" DROP CONSTRAINT "OpenseaCollectibleCache_socketId_fkey";

-- DropForeignKey
ALTER TABLE "PositionCache" DROP CONSTRAINT "PositionCache_socketId_fkey";

-- AddForeignKey
ALTER TABLE "PositionCache" ADD CONSTRAINT "PositionCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OpenseaCollectible" ADD CONSTRAINT "OpenseaCollectible_cacheSocketId_cacheChain_fkey" FOREIGN KEY ("cacheSocketId", "cacheChain") REFERENCES "OpenseaCollectibleCache"("socketId", "chain") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OpenseaCollectibleCache" ADD CONSTRAINT "OpenseaCollectibleCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
