-- CreateTable
CREATE TABLE "Address" (
    "id" TEXT NOT NULL,

    CONSTRAINT "Address_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Domain" (
    "verifyingContract" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "version" TEXT NOT NULL,
    "chainId" INTEGER NOT NULL,

    CONSTRAINT "Domain_pkey" PRIMARY KEY ("verifyingContract","name","version","chainId")
);

-- CreateTable
CREATE TABLE "Current" (
    "name" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "data" TEXT NOT NULL,

    CONSTRAINT "Current_pkey" PRIMARY KEY ("name","type","data")
);

-- CreateTable
CREATE TABLE "Fuse" (
    "target" TEXT NOT NULL,
    "data" TEXT NOT NULL,
    "plugId" TEXT,

    CONSTRAINT "Fuse_pkey" PRIMARY KEY ("target","data")
);

-- CreateTable
CREATE TABLE "Plug" (
    "id" TEXT NOT NULL,
    "currentName" TEXT NOT NULL,
    "currentType" TEXT NOT NULL,
    "currentData" TEXT NOT NULL,
    "plugsId" TEXT,

    CONSTRAINT "Plug_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Plugs" (
    "id" TEXT NOT NULL,
    "socket" TEXT NOT NULL,
    "salt" TEXT NOT NULL,
    "fee" INTEGER NOT NULL,
    "maxFeePerGas" INTEGER NOT NULL,
    "maxPriorityFeePerGas" INTEGER NOT NULL,
    "executor" TEXT NOT NULL,

    CONSTRAINT "Plugs_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "LivePlugs" (
    "plugsId" TEXT NOT NULL,
    "signature" TEXT NOT NULL,
    "addressId" TEXT NOT NULL,

    CONSTRAINT "LivePlugs_pkey" PRIMARY KEY ("plugsId","addressId")
);

-- AddForeignKey
ALTER TABLE "Fuse" ADD CONSTRAINT "Fuse_plugId_fkey" FOREIGN KEY ("plugId") REFERENCES "Plug"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plug" ADD CONSTRAINT "Plug_currentName_currentType_currentData_fkey" FOREIGN KEY ("currentName", "currentType", "currentData") REFERENCES "Current"("name", "type", "data") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plug" ADD CONSTRAINT "Plug_plugsId_fkey" FOREIGN KEY ("plugsId") REFERENCES "Plugs"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_plugsId_fkey" FOREIGN KEY ("plugsId") REFERENCES "Plugs"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LivePlugs" ADD CONSTRAINT "LivePlugs_addressId_fkey" FOREIGN KEY ("addressId") REFERENCES "Address"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
