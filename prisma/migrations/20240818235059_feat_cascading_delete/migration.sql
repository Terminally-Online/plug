-- DropForeignKey
ALTER TABLE "Implementation" DROP CONSTRAINT "Implementation_fungibleName_fungibleSymbol_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_cacheId_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_fungibleName_fungibleSymbol_fkey";

-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_protocolName_fkey";

-- AddForeignKey
ALTER TABLE "Implementation" ADD CONSTRAINT "Implementation_fungibleName_fungibleSymbol_fkey" FOREIGN KEY ("fungibleName", "fungibleSymbol") REFERENCES "Fungible"("name", "symbol") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_fungibleName_fungibleSymbol_fkey" FOREIGN KEY ("fungibleName", "fungibleSymbol") REFERENCES "Fungible"("name", "symbol") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_protocolName_fkey" FOREIGN KEY ("protocolName") REFERENCES "Protocol"("name") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "PositionCache"("socketId") ON DELETE CASCADE ON UPDATE CASCADE;
