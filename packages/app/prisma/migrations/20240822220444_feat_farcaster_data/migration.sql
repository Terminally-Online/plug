-- CreateTable
CREATE TABLE "FarcasterUser" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fid" INTEGER NOT NULL,
    "username" TEXT NOT NULL,
    "displayName" TEXT,
    "avatar" TEXT,
    "socketId" TEXT,

    CONSTRAINT "FarcasterUser_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_FarcasterFollowing" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "FarcasterUser_fid_key" ON "FarcasterUser"("fid");

-- CreateIndex
CREATE UNIQUE INDEX "FarcasterUser_username_key" ON "FarcasterUser"("username");

-- CreateIndex
CREATE UNIQUE INDEX "FarcasterUser_socketId_key" ON "FarcasterUser"("socketId");

-- CreateIndex
CREATE UNIQUE INDEX "_FarcasterFollowing_AB_unique" ON "_FarcasterFollowing"("A", "B");

-- CreateIndex
CREATE INDEX "_FarcasterFollowing_B_index" ON "_FarcasterFollowing"("B");

-- AddForeignKey
ALTER TABLE "FarcasterUser" ADD CONSTRAINT "FarcasterUser_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_FarcasterFollowing" ADD CONSTRAINT "_FarcasterFollowing_A_fkey" FOREIGN KEY ("A") REFERENCES "FarcasterUser"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_FarcasterFollowing" ADD CONSTRAINT "_FarcasterFollowing_B_fkey" FOREIGN KEY ("B") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
