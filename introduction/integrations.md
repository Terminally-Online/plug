---
head:
  - - meta
    - property: og:title
      content: Integrations
  - - meta
    - name: description
      content:
  - - meta
    - property: og:description
      content:
---

# Integrations

<span style="color: rgba(0,0,0,0.6)">With the large catalog of protocol integrations you have a wide range of options and control. By surfacing all the primary actions of each protocol you can have all your onchain activity in one place. If there's an integration or chain you'd like to see that we do not have yet, please [reach out](https://twitter.com/onplug_io)!</span>

<script setup>
    // TODO: Add Enso Finance integrations to this list.

    const protocolList = ['aave','aerodrome','alchemix','balancer','chainlink','compound','convex','curve','eigen-layer','ens','ethena','frax-lend','gearbox','hop','lido','maker','nouns','paraswap','rocket-pool','sushiswap','synthetix','uniswap','wasabi','yearn','zora']

    const bigList = ["ens"]

    const chainList = ['arbitrum','avalanche','base','bera','blast','ethereum','optimism','polygon','scroll','zksync', 'zora']

    const toChainImagePath = (str) => `/blockchain/${str}.png`
    const toProtocolImagePath = (str) => `/protocols/${str}.png`
    const toTitleCase = (str) => bigList.includes(str) 
        ? str.toUpperCase() 
        : str
            .replace(/-/g, ' ')
            .replace(/([a-z])([A-Z])|([A-Z])([A-Z][a-z])/g, "$1$3 $2$4")
            .split(" ")
            .map(word =>
                word.toUpperCase() === word && word.length > 1 
                    ? word 
                    : word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
            )
            .join(" ")
</script>

## Chains

With support for all the major Ethereum based chains you can use your money wherever. Plug can easily be deployed on any Ethereum based blockchain.

<div className="integrations">
    <div v-for="item in chainList" :key="item" className="integration">
        <img :src="toChainImagePath(item)" :alt="toTitleCase(item)">
        <p>{{ toTitleCase(item) }}</p>
        <p>Coming Soon</p>
    </div>
</div>

## Protocols

Plug is a generalized framework and protocol that builds on top of existing protocols: it essentially integrates with everything in the Ethereum ecosystem. If you'd like Plug to integrate a protocol, please [reach out](https://twitter.com/onplug_io).

<div className="integrations">
    <div v-for="item in protocolList" :key="item" className="integration">
        <img :src="toProtocolImagePath(item)" :alt="toTitleCase(item)">
        <p>{{ toTitleCase(item) }}</p>
        <p>Coming Soon</p>
    </div>
</div>
