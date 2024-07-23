-- CreateTable
CREATE TABLE "OpenseaCollection" (
    "slug" TEXT NOT NULL,
    "collection" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "image_url" TEXT NOT NULL,
    "banner_image_url" TEXT NOT NULL,
    "owner" TEXT NOT NULL,
    "category" TEXT NOT NULL,
    "is_disabled" BOOLEAN NOT NULL,
    "is_nsfw" BOOLEAN NOT NULL,
    "trait_offers_enabled" BOOLEAN NOT NULL,
    "collection_offers_enabled" BOOLEAN NOT NULL,
    "opensea_url" TEXT NOT NULL,
    "project_url" TEXT NOT NULL,
    "wiki_url" TEXT NOT NULL,
    "discord_url" TEXT NOT NULL,
    "telegram_url" TEXT NOT NULL,
    "twitter_username" TEXT NOT NULL,
    "instagram_username" TEXT NOT NULL,
    "total_supply" INTEGER NOT NULL,
    "created_date" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "OpenseaCollection_pkey" PRIMARY KEY ("slug")
);
