import { default as fse } from "fs-extra";
import { resolve } from "pathe";
import { defineConfig, HeadConfig } from "vitepress";

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
  appearance: false,
  themeConfig: {
    logo: "/logo-black.svg",
    siteTitle: false,

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
            ],
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
          }
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

  transformHead({ assets }) {
    const head: HeadConfig[] = [];

    head.push([
      "meta",
      {
        property: "og:image",
        content: "https://docs.onplug.io/opengraph.png",
      },
    ]);

    // make the image big
    head.push([
      "meta",
      {
        property: "og:image:width",
        content: "1920",
      },
    ]);

    head.push([
      "meta",
      {
        property: "og:image:height",
        content: "1080",
      },
    ]);

    head.push([
      "meta",
      {
        property: "twitter:image",
        content: "https://docs.onplug.io/opengraph.png",
      },
    ]);

    // make the image large
    head.push([
      "meta",
      {
        property: "twitter:image:width",
        content: "1920",
      },
    ]);

    head.push([
      "meta",
      {
        property: "twitter:image:height",
        content: "1080",
      },
    ]);

    head.push([
      "meta",
      {
        property: "twitter:card",
        content: "summary_large_image",
      },
    ]);

    const font = assets.find(() => /Satoshi-Variable\.\w+\.woff2/);
    if (font)
      head.push([
        "link",
        {
          rel: "preload",
          href: font,
          as: "font",
          type: "font/woff2",
          crossorigin: "",
        },
      ]);

    return head;
  },

  markdown: {
    lineNumbers: true,
  },
});
