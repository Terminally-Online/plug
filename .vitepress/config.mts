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
  title: "Plug Documentation",
  description: "Documentation for the Plug protocol.",
  appearance: "dark",
  themeConfig: {
    logo: { light: "/logo-dark.svg", dark: "/logo-white.svg" },

    nav: [{ text: "Home", link: "https://www.onplug.io" }],

    sidebar: [
      {
        text: "Introduction",
        items: [
          {
            text: "Why Plug",
            link: "/introduction/why-plug",
          },
          {
            text: "If This, Then That",
            link: "/introduction/if-this-then-that",
          },
          {
            text: "FAQ",
            link: "/introduction/frequently-asked-questions",
          },
        ],
      },
      {
        text: "Plugs",
        collapsed: false,
        items: [
          {
            text: "Introduction",
            link: "/plugs/introduction",
            items: [
              {
                text: "Imperative Transactions",
                link: "/plugs/imperative-transactions",
              },
              {
                text: "Declarative Messages",
                link: "/plugs/declarative-messages",
              },
            ],
          },
          {
            text: "Execution Paths",
            link: "/plugs/execution-paths",
            items: [
              {
                text: "Single Lane",
                link: "/plugs/execution-paths/single-lane",
              },
              {
                text: "Multi-Dimensional",
                link: "/plugs/execution-paths/multi-dimensional",
              },
              {
                text: "Native Transactions",
                link: "/plugs/execution-paths/native-transactions",
              },
              {
                text: "Meta-Transactions",
                link: "/plugs/execution-paths/meta-transactions",
              },
              {
                text: "Channels",
                link: "/plugs/execution-paths/channels",
              },
            ],
          },
        ],
      },
      {
        text: "Types and Decoders",
        collapsed: true,
        items: [
          {
            text: "EIP-712",
            link: "/decoders/eip-712",
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
            items: getItems("./generated/base-types"),
          },
          {
            text: "Hash Getters",
            link: "/decoders/hash-getters",
            items: getItems("./generated/hash-getters"),
          },
          {
            text: "Digest Getters",
            link: "/decoders/digest-getters",
            items: getItems("./generated/digest-getters"),
          },
          {
            text: "Signer Getters",
            link: "/decoders/signer-getters",
            items: getItems("./generated/signer-getters"),
          },
        ],
      },
      {
        text: "Core Abstracts",
        collapsed: true,
        items: [
          {
            text: "FuseEnforcer",
            link: "/core/fuse",
            items: [
              {
                text: "enforceFuse",
                link: "/core/fuse/enforce-fuse",
              },
            ],
          },
          {
            text: "Plug",
            link: "/core/framework",
            items: [
              {
                text: "contractInvoke",
                link: "/core/framework/contract-invoke",
              },
              {
                text: "invoke",
                link: "/core/framework/invoke",
              },
            ],
          },
        ],
      },
      {
        text: "Deployable Instances",
        collapsed: true,
        items: [
          {
            text: "Enforcers",
            link: "/instances/enforcers",
            items: [
              {
                text: "Allowed Methods",
                link: "/instances/enforcers/allowed-methods",
              },
              {
                text: "Block Number",
                link: "/instances/enforcers/block-number",
              },
              {
                text: "Timestamp",
                link: "/instances/enforcers/timestamp",
              },
              {
                text: "Schedule Windows",
                link: "/instances/enforcers/schedule-windows",
              },
              {
                text: "Limited Calls",
                link: "/instances/enforcers/limited-calls",
              },
              {
                text: "ERC20Allowance",
                link: "/instances/enforcers/erc20-allowance",
              },
              {
                text: "Revocation",
                link: "/instances/enforcers/revocation",
              },
            ],
          },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/nftchance/plug" },
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
