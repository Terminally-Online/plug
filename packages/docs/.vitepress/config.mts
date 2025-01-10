import { defineConfig, HeadConfig } from "vitepress";

export default defineConfig({
  // base: "/plug/",

  title: "Plug Documentation",
  titleTemplate: ":title | Plug Documentation",
  description: "Documentation for the Plug protocol and application.",

  appearance: false,
  lastUpdated: true,

  markdown: {
    lineNumbers: true,
  },

  themeConfig: {
    logo: "/favicon.ico",
    externalLinkIcon: true,

    nav: [{ text: "Home", link: "https://www.onplug.io" }],
    outline: { level: "deep" },

    sidebar: [
      {
        text: "Getting Started",
        items: [
          {
            text: "Introduction",
            link: "/",
          },
          {
            text: "Integrations",
            link: "/introduction/integrations",
          },
          {
            text: "Frequently Asked Questions",
            link: "/introduction/frequently-asked-questions",
          },
        ],
      },

      {
        text: "Core Concepts",
        items: [
          {
            text: "Architecture",
            link: "/concepts/architecture",
          },
          {
            text: "Constraints",
            link: "/concepts/constraints",
          },
          {
            text: "Actions",
            link: "/concepts/actions",
          },
          {
            text: "Strategies",
            link: "/concepts/strategies",
          },
        ],
      },

      {
        text: "Developer",
        items: [
          {
            text: "Schedule Lifecycle",
            link: "/developer/scheduling",
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

  transformHead({ assets }) {
    const head: HeadConfig[] = [];

    for (const item of [
      {
        property: "og:image",
        content: "https://docs.onplug.io/opengraph.png",
      },
      {
        property: "og:image:width",
        content: "1920",
      },
      {
        property: "og:image:height",
        content: "1080",
      },
      {
        property: "twitter:image",
        content: "https://docs.onplug.io/opengraph.png",
      },
      {
        property: "twitter:image:width",
        content: "1920",
      },
      {
        property: "twitter:image:height",
        content: "1080",
      },
      {
        property: "twitter:card",
        content: "summary_large_image",
      },
    ])
      head.push(["meta", item]);

    head.push([
      "link",
      {
        rel: "preload",
        href: assets.find(() => /Satoshi-Variable\.\w+\.woff2/) ?? "",
        as: "font",
        type: "font/woff2",
        crossorigin: "",
      },
    ]);
    head.push(["link", { rel: "icon", href: "/favicon.ico" }]);

    head.push([
      "script",
      {
        async: "",
        src: "https://www.googletagmanager.com/gtag/js?id=G-CH66Y2E034",
      },
    ]);
    head.push([
      "script",
      {},
      `window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());
      gtag('config', 'G-CH66Y2E034');`,
    ]);

    return head;
  },
});
