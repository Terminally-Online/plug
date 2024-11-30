---
head:
  - - meta
    - property: og:title
      content: Integrations
  - - meta
    - name: description
      content: Integrations into the Plug ecosystem.
  - - meta
    - property: og:description
      content: Integrations into the Plug ecosystem.
---

<script setup>
   import IntegrationList from './IntegrationList.vue'

    const protocolList = ['aave','aerodrome','alchemix','balancer','chainlink','compound','convex','curve','eigen-layer','ens','ethena','frax-lend','gearbox','hop','lido','maker','nouns','paraswap','rocket-pool','sushiswap','synthetix','uniswap','wasabi','yearn','zora']

    const chainList = ['arbitrum','avalanche','base','bera','blast','ethereum','optimism','polygon','scroll','zksync', 'zora']
</script>

# Integrations

<span style="color: rgba(0,0,0,0.6)">With the large catalog of protocol integrations you have a wide range of options and control. By surfacing all the primary actions of each protocol you can have all your onchain activity in one place. If there's an integration or chain you'd like to see that we do not have yet, please [reach out](https://twitter.com/onplug_io)!</span>

## Chains

With support for all the major Ethereum based chains you can use your money wherever. Plug can easily be deployed on any Ethereum based blockchain.

<IntegrationList :list="chainList" type="chain" />

## Protocols

Plug is a generalized framework and protocol that builds on top of existing protocols: it essentially integrates with everything in the Ethereum ecosystem. If you'd like Plug to integrate a protocol, please [reach out](https://twitter.com/onplug_io).

<IntegrationList :list="protocolList" type="protocol" />
