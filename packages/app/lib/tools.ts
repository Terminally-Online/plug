
export type ToolFunction = {
  name: string;
  description: string;
  execute: (...args: any[]) => Promise<any>;
}

export const TOOLS: Record<string, ToolFunction> = {
  "schemas": {
    name: "schemas",
    description: "List all available schemas in the system",
    execute: async () => {
      return ["User", "Plug", "Holdings"];
    }
  },
  "holdings": {
    name: "holdings",
    description: "Get current holdings information",
    execute: async () => {
      return { /* holdings data */ };
    }
  },
  "price": {
    name: "price",
    description: "Get current price information for a specific item",
    execute: async (itemId: string) => {
      return { /* price data */ };
    }
  },
}
