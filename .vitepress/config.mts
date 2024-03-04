import { default as fse } from "fs-extra";
import { resolve } from "pathe";
import { defineConfig } from "vitepress";

const rootDir = resolve(process.cwd());

const ordering = ["domain", "fuse", "Pin", "Transaction", "Breaker", "Plug"];

// * Get the generated files in a directory and create the array of items.
function getItems(directory: string) {
  const directoryPath = resolve(rootDir, directory);
  const files = fse.readdirSync(directoryPath);

  return files.map((file) => {
    const name = file.replace(".md", "");
    const link = `${directory}/${name}`;

    console.log(link);

    return {
      text: name,
      link: link.replace("./", ""),
    };
  });
}

export default defineConfig({
  title: "Plug",
  description: "Documentation for the Plug protocol and application.",
  appearance: "dark",
  themeConfig: {
    logo: { light: "/logo-dark.svg", dark: "/logo-white.svg" },

    nav: [{ text: "Home", link: "https://www.onplug.io" }],

    sidebar: [
      {
        text: "Introduction",
        items: [
          {
            text: "Getting Started",
            link: "/",
          },
          {
            text: "Why Plug",
            collapsed: false,
            items: [
              {
                text: "The Problem and Solution",
                link: "/introduction/why-plug",
              },
              {
                text: "Transaction Types",
                link: "/introduction/why/transactions",
              },
              {
                text: "Passive Management",
                link: "/introduction/why/passive-management",
              },
            ]
          },
          {
            text: "FAQ",
            link: "/introduction/frequently-asked-questions",
          },
        ],
      },
      {
        text: "Core Mechanisms",
        items: [
          {
            text: "EIP-712",
            link: "/decoders/eip-712",
            collapsed: true,
            items: [
              {
                text: "Live Pairs",
                link: "/decoders/eip-712/signed-pairs",
              },
              {
                text: "Automated Generation",
                link: "/decoders/eip-712/automated-generation",
              },
            ],
          },
          {
            text: "Base Types",
            link: "/decoders/base-types",
            collapsed: true,
            items: getItems("./generated/base-types"),
          },
          {
            text: "Hash Getters",
            link: "/decoders/hash-getters",
            collapsed: true,
            items: getItems("./generated/hash-getters"),
          },
          {
            text: "Digest Getters",
            link: "/decoders/digest-getters",
            collapsed: true,
            items: getItems("./generated/digest-getters"),
          },
          {
            text: "Signer Getters",
            link: "/decoders/signer-getters",
            collapsed: true,
            items: getItems("./generated/signer-getters"),
          },
          {
            text: "Fuses",
            link: "/core/fuses",
            collapsed: true,
            items: [
              {
                text: "encode",
                link: "/core/fuse/encode",
              },
              {
                text: "decode",
                link: "/core/fuse/decode",
              },
              {
                text: "enforceFuse",
                link: "/core/fuse/enforce-fuse",
              },
            ],
          },
          {
            text: "Sockets",
            link: "/core/sockets",
            collapsed: true,
            items: [
              {
                text: "signer",
                link: "/core/sockets/signer",
              },
              {
                text: "plug",
                link: "/core/sockets/plug",
              },
            ],
          },
          {
            text: "Routers",
            link: "/core/routers",
            collapsed: true,
            items: [
              {
                text: "plug",
                link: "/core/routers/plug",
              },
            ],
          },
          {
            text: "Solvers",
            link: "/core/solvers",
            collapsed: true,
            items: [
              {
                text: "solve",
                link: "/core/solvers/solve",
              },
            ],
          },
          {
            text: "Deployable Instances",
            collapsed: true,
            items: [
              {
                text: "Deterministic",
                link: "/instances/deployable/deterministic"
              },
              {
                text: "Fuses",
                collapsed: true,
                items: [
                  {
                    text: "Threshold",
                    link: "/instances/fuses/threshold",
                  },
                  {
                    text: "Schedule Windows",
                    link: "/instances/fuses/schedule-windows",
                  },
                  {
                    text: "Limited Calls",
                    link: "/instances/fuses/limited-calls",
                  },
                  {
                    text: "Revocation",
                    link: "/instances/fuses/revocation",
                  },
                ],
              },
              {
                text: "Sockets",
                collapsed: true,
                items: [
                  {
                    text: "Vaults",
                    link: "/instances/vaults"
                  },
                ]
              },
            ],
          },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/nftchance/plug" },
      { icon: "twitter", link: "https://twitter.com/onplug_io" },
    ],

    editLink: {
      pattern: "https://github.com/nftchance/plug-docs/edit/master/:path",
    },

    search: {
      provider: "local",
    },
  },

  lastUpdated: true,

  // * Load the font files.
  transformHead({ assets }) {
    const myFontFile = assets.find(() => /Satoshi-Variable\.\w+\.woff2/);

    if (myFontFile) {
      return [
        [
          "link",
          {
            rel: "preload",
            href: myFontFile,
            as: "font",
            type: "font/woff2",
            crossorigin: "",
          },
        ],
      ];
    }
  },

  markdown: {
    lineNumbers: true,
  },
});
