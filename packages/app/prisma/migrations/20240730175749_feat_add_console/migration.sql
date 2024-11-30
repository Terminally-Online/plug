-- AlterTable
ALTER TABLE "OpenseaCollection" ALTER COLUMN "createdDate" SET DEFAULT CURRENT_TIMESTAMP;

-- CreateTable
CREATE TABLE "ConsoleColumn" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "key" TEXT NOT NULL,
    "index" INTEGER NOT NULL,
    "width" INTEGER,
    "consoleId" TEXT,

    CONSTRAINT "ConsoleColumn_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Console" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "Console_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Console_id_key" ON "Console"("id");

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_consoleId_fkey" FOREIGN KEY ("consoleId") REFERENCES "Console"("id") ON DELETE SET NULL ON UPDATE CASCADE;
