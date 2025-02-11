-- CreateTable
CREATE TABLE "Message" (
    "id" TEXT NOT NULL,
    "socketId" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "isUser" BOOLEAN NOT NULL,
    "timeSent" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "Message_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Message" ADD CONSTRAINT "Message_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE CASCADE ON UPDATE CASCADE;
