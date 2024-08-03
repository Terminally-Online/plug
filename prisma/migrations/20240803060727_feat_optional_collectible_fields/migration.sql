-- AlterTable
ALTER TABLE "OpenseaCollectible" ALTER COLUMN "displayImageUrl" DROP NOT NULL,
ALTER COLUMN "imageUrl" DROP NOT NULL,
ALTER COLUMN "metadataUrl" DROP NOT NULL;
