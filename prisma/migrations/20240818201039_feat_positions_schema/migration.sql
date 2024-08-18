-- CreateTable
CREATE TABLE "Protocol" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "name" TEXT NOT NULL,
    "icon" TEXT NOT NULL,
    "url" TEXT NOT NULL,

    CONSTRAINT "Protocol_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Implementation" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "chain" TEXT NOT NULL,
    "contract" TEXT NOT NULL,
    "decimals" INTEGER NOT NULL,
    "fungibleName" TEXT NOT NULL,
    "fungibleSymbol" TEXT NOT NULL,

    CONSTRAINT "Implementation_pkey" PRIMARY KEY ("chain","contract")
);

-- CreateTable
CREATE TABLE "Fungible" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "name" TEXT NOT NULL,
    "symbol" TEXT NOT NULL,
    "icon" TEXT NOT NULL,
    "verified" BOOLEAN NOT NULL,

    CONSTRAINT "Fungible_pkey" PRIMARY KEY ("name","symbol")
);

-- CreateTable
CREATE TABLE "Position" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "chain" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "balance" DOUBLE PRECISION NOT NULL,
    "fungibleName" TEXT NOT NULL,
    "fungibleSymbol" TEXT NOT NULL,
    "protocolId" TEXT,
    "cacheId" TEXT NOT NULL,

    CONSTRAINT "Position_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PositionCache" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "socketId" TEXT NOT NULL,

    CONSTRAINT "PositionCache_pkey" PRIMARY KEY ("socketId")
);

-- AddForeignKey
ALTER TABLE "Implementation" ADD CONSTRAINT "Implementation_fungibleName_fungibleSymbol_fkey" FOREIGN KEY ("fungibleName", "fungibleSymbol") REFERENCES "Fungible"("name", "symbol") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_fungibleName_fungibleSymbol_fkey" FOREIGN KEY ("fungibleName", "fungibleSymbol") REFERENCES "Fungible"("name", "symbol") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_protocolId_fkey" FOREIGN KEY ("protocolId") REFERENCES "Protocol"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "PositionCache"("socketId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PositionCache" ADD CONSTRAINT "PositionCache_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
