-- CreateTable
CREATE TABLE "Companion" (
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "name" TEXT NOT NULL,
    "feedCount" INTEGER NOT NULL DEFAULT 0,
    "lastFeedAt" TIMESTAMP(3),
    "socketId" TEXT NOT NULL,

    CONSTRAINT "Companion_pkey" PRIMARY KEY ("socketId")
);

-- CreateIndex
CREATE INDEX "Companion_name_idx" ON "Companion"("name");

-- AddForeignKey
ALTER TABLE "Companion" ADD CONSTRAINT "Companion_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "SocketIdentity"("socketId") ON DELETE CASCADE ON UPDATE CASCADE;
