import { FC, useCallback, useMemo, useState } from "react"
import { Frame, CollectibleImage, TransferRecipient } from "@/components"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumns } from "@/state"

type CollectibleType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
type CollectionType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]

interface TransferNFTFrameProps {
  index: number
  collectible: CollectibleType
  collection: CollectionType
  recipient: string
  color: string
  textColor: string
  isERC1155: boolean
}

export const TransferNFTFrame: FC<TransferNFTFrameProps> = ({
  index,
  collectible,
  collection,
  recipient,
  color,
  textColor,
  isERC1155
}) => {
  const { isFrame, frame } = useColumns(
    index,
    `${collection.address}-${collection.chain}-${collectible.tokenId}-transfer-amount`
  )

  // Initialize with "1" for both ERC721 and ERC1155
  const [amount, setAmount] = useState("1")
  
  // Parse collectible amount as number for comparisons
  const maxAmount = parseInt(collectible.amount)
  
  const isReady = useMemo(() => {
    if (!isERC1155) return true
    const numAmount = parseInt(amount)
    return !isNaN(numAmount) && numAmount > 0 && numAmount <= maxAmount
  }, [amount, isERC1155, maxAmount])

  const handleMaxClick = useCallback(() => {
    if (isERC1155) {
      setAmount(collectible.amount)
    }
  }, [isERC1155, collectible.amount])

  return (
    <Frame
      index={index}
      icon={
        <div className="relative h-8 w-10">
          <div
            className="h-8 w-8 rounded-full bg-cover bg-center bg-no-repeat"
            style={{
              backgroundImage: `url(${collection.iconUrl})`
            }}
          />
        </div>
      }
      label={`Transfer ${collection.name} #${collectible.tokenId}`}
      visible={isFrame}
      handleBack={() => frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-recipient`)}
      hasChildrenPadding={false}
      hasOverlay
    >
      <div className="mb-4 flex flex-col gap-4">
        {/* Recipient Preview */}
        <div className="px-6">
          <TransferRecipient
            address={recipient}
            handleSelect={() => frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-recipient`)}
          />
        </div>

        {/* NFT Preview */}
        <div className="px-6">
          <CollectibleImage
            video={collectible.videoUrl?.includes("mp4") ? collectible.videoUrl : undefined}
            image={collectible.imageUrl}
            fallbackImage={collection.iconUrl}
            name={collectible.name || collection.name}
          />
        </div>

        {/* Amount Input for ERC1155 */}
        {isERC1155 && (
          <div className="flex flex-col gap-2 px-6">
            <div className="flex flex-row items-center justify-between">
              <p className="font-bold opacity-40">Amount</p>
              <p
                className="ml-auto cursor-pointer font-bold text-black/40 hover:brightness-105"
                onClick={handleMaxClick}
                style={{ color: amount !== collectible.amount ? color : undefined }}
              >
                Max
              </p>
            </div>
            <input
              type="text"
              inputMode="numeric"
              pattern="[0-9]*"
              value={amount}
              onChange={(e) => {
                const value = e.target.value.replace(/[^0-9]/g, "")
                if (value === "") {
                  setAmount("")
                  return
                }
                const parsed = parseInt(value)
                if (!isNaN(parsed)) {
                  setAmount(Math.min(parsed, maxAmount).toString())
                }
              }}
              className="w-full rounded-lg border p-4 font-bold"
              placeholder="Enter amount"
            />
            <p className="text-sm font-bold opacity-40">
              Available: {collectible.amount}
            </p>
          </div>
        )}

        {/* Transfer Button */}
        <div className="mx-6">
          <button
            className={cn(
              "flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
              !isReady && "bg-white"
            )}
            style={{
              backgroundColor: isReady ? color : "#FFFFFF",
              color: isReady ? textColor : color,
              borderColor: isReady ? "#FFFFFF" : color
            }}
            disabled={!isReady}
          >
            {isReady ? "Transfer" : "Enter Amount"}
          </button>
        </div>
      </div>
    </Frame>
  )
}