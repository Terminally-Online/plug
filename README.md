# 🔌 Plug Pitch

I was losing my sanity dealing with updating all of the raise materials. So, this repository takes all the base images and generates everything for you in a second or two.

To utilize this repository there are a few dependencies you must have installed:

```ml
├─ ffmpeg - "A complete, cross-platform solution to record, convert and stream audio and video."
└─ pnpm — "Efficient package manager for Node modules."
```

To generate a suite of assets you will run:

```bash
pnpm generate
or
pnpm generate <email>
```

When doing this, you will be met with a:

```ml
├─ deck (pdf) - "Traditional version of a pitch deck."
├─ memo (pdf) - "Memo that gives more granular detail of Plug, the vision, team and market."
├─ one-pager - "Introduction that provides context for a conversation/call"
│  ├─ plug-with-chance-one-pager (pdf) - "Includes a call link to schedule with Chance."
│  ├─ plug-with-reka-one-pager (pdf) - "Includes a call link to schedule with Chance and Reka."
│  └─ plug-with-drake-one-pager (pdf) - "Includes a call link to schedule with Chance and Drake."
├─ vision-deck — "Sharable and watermarked version of the deck in pdf and gif format"
│  └─ vision (pdf & gif) - "Narrative based deck that provides vision excluding traditional pitch details."
```

As each asset is generated you can see the in terminal where it was generated as well as click the file path to go straight to it. Finally, due to the watermarking you will have a local reference of each of the files used to generate the pdfs. You do not need to do anything with these.

To update the assets that are used in `/base/` simply drop in the new files and everything will continue working. The following name conventions must be followed:

```ml
├─ deck (pdf) - "Images can be numbered or be exported from Figma as Seed - Boomer - N"
├─ memo (pdf) - "Images can be numbered or be exported from Figma as Investor Memo - N"
├─ one-pager - "Image must be named one-pager"
└─ vision-deck — "Images can be numbered or be exported from Figma as Seed - N"
```
