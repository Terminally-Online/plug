-- DropForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" DROP CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey";

-- AddForeignKey
ALTER TABLE "OpenseaCollectibleMetadata" ADD CONSTRAINT "OpenseaCollectibleMetadata_collectibleId_fkey" FOREIGN KEY ("collectibleId") REFERENCES "OpenseaCollectible"("id") ON DELETE CASCADE ON UPDATE CASCADE;
