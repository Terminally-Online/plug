-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_socketId_fkey";

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
