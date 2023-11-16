
"use client"

import * as React from "react"
import {
  CaretSortIcon,
  CheckIcon,
} from "@radix-ui/react-icons"

import { cn } from "@/lib/utils"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"

import { Button } from "@/components/ui/button"
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command"

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"

import CanvasStore from "../../lib/store"

import { z } from "zod"
import { Input } from "@/components/ui/input"

const nounsSchema = z.object({
  price: z.number(),
})

const thresholdSchema = z.object({
  threshold: z.number().default(Date.now()),
})

const groups = [
  {
    label: "Nouns",
    teams: [
      {
        label: "Can Bid on Noun",
        value: "can-bid",
        schema: nounsSchema,
      },
      {
        label: "Place Bid on Noun",
        value: "place-bid",
        schema: nounsSchema,
      }
    ],
  },
  {
    label: "Schedule",
    teams: [
      {
        label: "Within Window",
        value: "within-window",
        schema: thresholdSchema,
      },
      {
        label: 'Before Block Number',
        value: 'before-block-number',
        schema: thresholdSchema,
      },
      {
        label: 'After Block Number',
        value: 'after-block-number',
        schema: thresholdSchema,
      },
      {
        label: 'Before Timestamp',
        value: 'before-timestamp',
        schema: thresholdSchema,
      },
      {
        label: 'After Timestamp',
        value: 'after-timestamp',
        schema: thresholdSchema,
      }
    ],
  }
]

type Team = (typeof groups)[number]["teams"][number]

type PopoverTriggerProps = React.ComponentPropsWithoutRef<typeof PopoverTrigger>

interface TeamSwitcherProps extends PopoverTriggerProps {}

export default function PlugSwitcher({ className }: TeamSwitcherProps) {
  const [open, setOpen] = React.useState(false)
  const [selectedTeam, setSelectedTeam] = React.useState<Team>(
    groups[0].teams[0]
  )

  CanvasStore.lockCamera(open)

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <div className="border border-stone-950 rounded-md w-full">
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a team"
            className={cn("w-full justify-between border-b-[1px] border-transparent border-b-stone-950 rounded-b-none", className)}
          >
            <Avatar className="mr-2 h-4 w-4">
              <AvatarImage
                src={`https://avatar.vercel.sh/${selectedTeam.value}.png`}
                alt={selectedTeam.label}
              />
              <AvatarFallback>SC</AvatarFallback>
            </Avatar>

            {selectedTeam.label}

            <CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>

        {Object.keys(selectedTeam.schema.shape).map((key) => (
          <div key={key} className="w-full flex flex-col">
            <Input
              type="text"
              id={key}
              placeholder={key
                .replace(/-/g, " ")
                .replace(/\w\S*/g, (w) => w.replace(/^\w/, (c) => c.toUpperCase()))}
              autoComplete="off"
              className="border-transparent"
            />
          </div>
        ))}
      </div>

      <PopoverContent className="w-[400px] p-0">
        <Command>
          <CommandList>
            <CommandInput placeholder="Search team..." />
            <CommandEmpty>No team found.</CommandEmpty>
            {groups.map((group) => (
              <CommandGroup key={group.label} heading={group.label}>
                {group.teams.map((team) => (
                  <CommandItem
                    key={team.value}
                    onSelect={() => {
                      setSelectedTeam(team)
                      setOpen(false)
                    }}
                    className="text-sm"
                  >
                    <Avatar className="mr-2 h-5 w-5">
                      <AvatarImage
                        src={`https://avatar.vercel.sh/${team.value}.png`}
                        alt={team.label}
                        className="grayscale"
                      />
                      <AvatarFallback>SC</AvatarFallback>
                    </Avatar>
                    {team.label}
                    <CheckIcon
                      className={cn(
                        "ml-auto h-4 w-4",
                        selectedTeam.value === team.value
                          ? "opacity-100"
                          : "opacity-0"
                      )}
                    />
                  </CommandItem>
                ))}
              </CommandGroup>
            ))}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}

