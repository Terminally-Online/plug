import { FC, useCallback, useMemo, useState, useRef } from "react"
import { Frame, CollectibleImage, TransferRecipient, Counter, Image } from "@/components"
import { cn, chains } from "@/lib"
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

  const maxAmount = parseInt(collectible.amount)
  const [amount, setAmount] = useState("0")
  const [dragPercentage, setDragPercentage] = useState(0)
  const [isPrecise, setIsPrecise] = useState(false)
  const containerRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLInputElement>(null)

  const handleDragStart = useCallback(
    (e: React.MouseEvent<HTMLDivElement>) => {
      e.preventDefault()

      const activeElement = document.activeElement as HTMLElement
      if (activeElement && activeElement.tagName === "INPUT") {
        activeElement.blur()
      }

      const handleDrag = (e: MouseEvent) => {
        if (containerRef.current) {
          const rect = containerRef.current.getBoundingClientRect()
          const x = e.clientX - rect.left
          const rawPercentage = (x / rect.width)
          
          // Calculate the whole number amount based on the drag position
          const newAmount = Math.max(0, Math.min(Math.round(maxAmount * rawPercentage), maxAmount))
          
          // Update amount first
          setAmount(newAmount.toString())
          
          // Then calculate and set the percentage based on the selected amount
          const newPercentage = (newAmount / maxAmount) * 100
          setDragPercentage(newPercentage)
        }
      }

      const handleDragEnd = () => {
        document.removeEventListener("mousemove", handleDrag)
        document.removeEventListener("mouseup", handleDragEnd)
        if (inputRef.current) {
          inputRef.current.focus()
        }
      }

      document.addEventListener("mousemove", handleDrag)
      document.addEventListener("mouseup", handleDragEnd)
    },
    [maxAmount]
  )

  const handleAmountChange = (value: string) => {
    const numericValue = value.replace(/[^0-9]/g, "")
    
    if (numericValue === "") {
      setAmount("0")
      setDragPercentage(0)
    } else {
      const parsedValue = parseInt(numericValue)
      if (!isNaN(parsedValue)) {
        const clampedValue = Math.min(Math.max(parsedValue, 0), maxAmount)
        setAmount(clampedValue.toString())
        // Calculate percentage based on the new amount
        setDragPercentage((clampedValue / maxAmount) * 100)
      }
    }
  }

  const handleMaxClick = useCallback(() => {
    if (isERC1155) {
      setAmount(maxAmount.toString())
      setDragPercentage(100)
    }
  }, [isERC1155, maxAmount])
  
  const isReady = useMemo(() => {
    if (!isERC1155) return true
    const numAmount = parseInt(amount)
    return !isNaN(numAmount) && numAmount > 0 && numAmount <= maxAmount
  }, [amount, isERC1155, maxAmount])

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
      label={isERC1155 ? "Transfer Amount" : "Transfer NFT"}
      visible={isFrame}
      handleBack={() => frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-recipient`)}
      hasChildrenPadding={false}
      hasOverlay
    >
      <div className="relative mb-4 flex flex-col gap-4">
        {/* Recipient Preview */}
        <div className="relative z-10 px-6">
          <TransferRecipient
            address={recipient}
            handleSelect={() => frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-recipient`)}
          />
        </div>

        {/* Amount Input for ERC1155 */}
        {isERC1155 && (
          <div className="relative z-[5] flex flex-col gap-4">
            <div className="px-6">
              <div className="flex flex-col gap-2">
                <CollectibleImage
                  video={collectible.videoUrl?.includes("mp4") ? collectible.videoUrl : undefined}
                  image={collectible.imageUrl ?? undefined}
                  fallbackImage={collection.iconUrl ?? undefined}
                  name={collectible.name || collection.name}
                />
              </div>
            </div>

            <div
              className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-grayscale-100 p-4"
              ref={containerRef}
              onMouseDown={handleDragStart}
              onMouseEnter={() => setIsPrecise(true)}
              onMouseLeave={() => setIsPrecise(false)}
            >
              <div className="flex w-full flex-row">
                <div className="flex flex-row items-center gap-4 px-2">
                  <div className="h-8 w-8 min-w-8 overflow-hidden rounded-md">
                    <CollectibleImage
                      video={collectible.videoUrl?.includes("mp4") ? collectible.videoUrl : undefined}
                      image={collectible.imageUrl ?? undefined}
                      fallbackImage={collection.iconUrl ?? undefined}
                      name={collectible.name || collection.name}
                      size="sm"
                    />
                  </div>

                  <div className="flex flex-col items-start">
                    <p className="mr-auto font-bold">{`${collection.name} #${collectible.tokenId}`}</p>
                    <p className="flex flex-row text-sm font-bold text-black/40">
                      <Counter count={dragPercentage ?? 0} decimals={0} />%
                    </p>
                  </div>
                </div>

                <div className="ml-auto flex-col items-end px-2">
                  <div className="pointer-events-none relative flex h-full w-max min-w-32 flex-col items-center justify-center text-right">
                    {isPrecise && (
                      <input
                        ref={inputRef}
                        value={amount}
                        onChange={e => handleAmountChange(e.target.value)}
                        className="sr-only pointer-events-none absolute inset-0"
                        autoFocus
                      />
                    )}

                    <p
                      className="my-auto ml-auto flex flex-row font-bold tabular-nums transition-all duration-200 ease-in-out"
                      style={{ color: isPrecise ? color : undefined }}
                    >
                      <Counter count={amount ?? "0"} />

                      {isPrecise && (
                        <div
                          className="absolute -right-2 bottom-3 top-3 w-[3px] animate-pulse rounded-full"
                          style={{ backgroundColor: color }}
                        />
                      )}
                    </p>
                  </div>
                </div>
              </div>

              <div
                className="absolute inset-0 z-[-2] min-w-4 rounded-r-lg opacity-20 blur-2xl filter"
                style={{ width: `${dragPercentage}%`, backgroundColor: color }}
              >
                <div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
                <div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
                <div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
              </div>
            </div>

            <div className="flex flex-row items-center justify-between gap-4 px-6">
              <p className="flex flex-row items-center gap-1 font-bold tabular-nums">
                <Image
                  src={chains[1].logo}
                  alt={"ethereum"}
                  className="mr-2 h-4 w-4 rounded-full"
                  width={24}
                  height={24}
                />
                $0.50
              </p>
              <p
                className="ml-auto cursor-pointer font-bold text-black/40 hover:brightness-105"
                onClick={handleMaxClick}
                style={{ color: amount !== collectible.amount ? color : undefined }}
              >
                Max
              </p>
            </div>
          </div>
        )}

        {/* NFT Preview */}
        {!isERC1155 && (
          <div className="relative z-[1] flex flex-col gap-4">
            <div className="px-6">
              <div className="flex flex-col gap-2">
                <CollectibleImage
                  video={collectible.videoUrl?.includes("mp4") ? collectible.videoUrl : undefined}
                  image={collectible.imageUrl ?? undefined}
                  fallbackImage={collection.iconUrl ?? undefined}
                />
              </div>
            </div>
            
            <div className="flex flex-row items-center justify-between gap-4 px-6">
              <p className="flex flex-row items-center gap-1 font-bold tabular-nums">
                <Image
                  src={chains[1].logo}
                  alt={"ethereum"}
                  className="mr-2 h-4 w-4 rounded-full"
                  width={24}
                  height={24}
                />
                $0.50
              </p>
            </div>
          </div>
        )}

        {/* Transfer Button */}
        <div className="relative z-20 mx-6">
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
            {isReady ? "Send" : "Enter Amount"}
          </button>
        </div>
      </div>
    </Frame>
  )
}