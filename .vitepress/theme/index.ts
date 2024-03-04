import type { Theme } from "vitepress";
import { h } from "vue";

import DefaultTheme from "vitepress/theme-without-fonts";
import "./style.css";

export default {
	extends: DefaultTheme,
	Layout: () => {
		return h(DefaultTheme.Layout, null, {});
	},
	enhanceApp() {},
} satisfies Theme;
