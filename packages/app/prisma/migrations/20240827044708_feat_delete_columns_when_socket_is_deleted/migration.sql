-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_socketId_fkey";

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
