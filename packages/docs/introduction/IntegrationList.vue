<template>
  <div :style="styles.integrations">
    <div
      v-for="item in list"
      :key="item"
      :style="styles.integration"
      @mouseenter="onMouseEnter"
      @mouseleave="onMouseLeave"
    >
      <img
        :src="getImagePath(item)"
        :alt="toTitleCase(item)"
        :style="styles.image"
      />
      <p :style="styles.title">{{ toTitleCase(item) }}</p>
      <p :style="styles.subtitle">Coming Soon</p>
    </div>
  </div>
</template>

<script setup>
import { defineProps, reactive } from "vue";

const props = defineProps({
  list: {
    type: Array,
    default: () => [],
  },
  type: String,
});

const styles = reactive({
  integrations: {
    display: "grid",
    gridTemplateColumns: "repeat(auto-fill, minmax(200px, 1fr))",
    gap: "20px",
  },
  integration: {
    border: "1px solid rgba(0, 0, 0, 0.1)",
    padding: "1.5rem",
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    borderRadius: "6px",
    transition: "background 0.2s",
    cursor: "pointer",
  },
  image: {
    width: "25%",
    borderRadius: "50%",
    marginBottom: "20px",
  },
  title: {
    margin: "0",
    fontWeight: "700",
  },
  subtitle: {
    margin: "0",
    opacity: "0.6",
    fontSize: "12px",
    fontWeight: "400",
  },
});

const onMouseEnter = (event) => {
  event.currentTarget.style.background = "rgba(0, 0, 0, 0.04)";
};

const onMouseLeave = (event) => {
  event.currentTarget.style.background = "";
};

const bigList = ["ens"];

const toChainImagePath = (str) => `/plug/blockchain/${str}.png`;
const toProtocolImagePath = (str) => `/plug/protocols/${str}.png`;
const getImagePath = (str) =>
  props.type === "chain" ? toChainImagePath(str) : toProtocolImagePath(str);

const toTitleCase = (str) =>
  bigList.includes(str)
    ? str.toUpperCase()
    : str
        .replace(/-/g, " ")
        .replace(/([a-z])([A-Z])|([A-Z])([A-Z][a-z])/g, "$1$3 $2$4")
        .split(" ")
        .map((word) =>
          word.toUpperCase() === word && word.length > 1
            ? word
            : word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
        )
        .join(" ");
</script>
