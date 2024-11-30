-- CreateTable
CREATE TABLE "Category" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "Category_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_CategoryToWorkflow" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_CategoryToWorkflow_AB_unique" ON "_CategoryToWorkflow"("A", "B");

-- CreateIndex
CREATE INDEX "_CategoryToWorkflow_B_index" ON "_CategoryToWorkflow"("B");

-- AddForeignKey
ALTER TABLE "_CategoryToWorkflow" ADD CONSTRAINT "_CategoryToWorkflow_A_fkey" FOREIGN KEY ("A") REFERENCES "Category"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_CategoryToWorkflow" ADD CONSTRAINT "_CategoryToWorkflow_B_fkey" FOREIGN KEY ("B") REFERENCES "Workflow"("id") ON DELETE CASCADE ON UPDATE CASCADE;
