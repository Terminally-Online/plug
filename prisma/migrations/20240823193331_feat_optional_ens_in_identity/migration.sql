-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_ensName_fkey";

-- AlterTable
ALTER TABLE "SocketIdentity" ALTER COLUMN "ensName" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_ensName_fkey" FOREIGN KEY ("ensName") REFERENCES "ENS"("ensName") ON DELETE SET NULL ON UPDATE CASCADE;
