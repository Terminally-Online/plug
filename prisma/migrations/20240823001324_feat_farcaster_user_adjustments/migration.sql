/*
  Warnings:

  - You are about to drop the `_FarcasterFollowing` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "FarcasterUser" DROP CONSTRAINT "FarcasterUser_socketId_fkey";

-- DropForeignKey
ALTER TABLE "_FarcasterFollowing" DROP CONSTRAINT "_FarcasterFollowing_A_fkey";

-- DropForeignKey
ALTER TABLE "_FarcasterFollowing" DROP CONSTRAINT "_FarcasterFollowing_B_fkey";

-- DropTable
DROP TABLE "_FarcasterFollowing";

-- CreateTable
CREATE TABLE "_FarcasterFollows" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_FarcasterFollows_AB_unique" ON "_FarcasterFollows"("A", "B");

-- CreateIndex
CREATE INDEX "_FarcasterFollows_B_index" ON "_FarcasterFollows"("B");

-- AddForeignKey
ALTER TABLE "FarcasterUser" ADD CONSTRAINT "FarcasterUser_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_FarcasterFollows" ADD CONSTRAINT "_FarcasterFollows_A_fkey" FOREIGN KEY ("A") REFERENCES "FarcasterUser"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_FarcasterFollows" ADD CONSTRAINT "_FarcasterFollows_B_fkey" FOREIGN KEY ("B") REFERENCES "FarcasterUser"("id") ON DELETE CASCADE ON UPDATE CASCADE;
