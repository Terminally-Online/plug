-- AlterTable
ALTER TABLE "ConsoleColumn" ADD COLUMN     "viewAsId" TEXT;

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_viewAsId_fkey" FOREIGN KEY ("viewAsId") REFERENCES "UserSocket"("id") ON DELETE SET NULL ON UPDATE CASCADE;
