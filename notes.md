Right now we are getting the following from the app API:

```json
{
    "result": {
        "data": {
            "json": [
                {
                    "id": "29594d1b-c824-43c1-af05-1fb2728ba0ef",
                    "actions": [
                        {
                            "categoryName": "plug",
                            "actionName": "baseFee",
                            "values": [
                                {
                                    "label": "Less than",
                                    "value": "<"
                                },
                                null
                            ]
                        }
                    ],
                    "socketId": "0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E"
                }
            ]
        }
    }
}
```

However, to build the transaction data our solver api is currently accepting:

```json
{
    "chainId": 1,
    "from": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "solver": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "actions": [
        {
            "type": "redeem",
            "inputs": {
                "protocol": "yearn_v3",
                "tokenIn": "0xBe53A109B494E5c9f97b9Cd39Fe969BE68BF6204",
                "tokenOut": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "amountIn": 20
            }
        }
    ]
}
```

So, we need to marry the two schemas together so that we can get a simulation to run. It will have a lot of actions, so we will step through each of them.
In order to do this properly, we need to make sure that the shape we go to will also support constraints.

We fundamentally need the "values" model to properly handle everything on the frontend.

I think we are moving in the right direction if we rewrite the existing solver schema to be:

```json
{
    "chainId": 1,
    "from": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "solver": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "actions": [
        {
            "protocol": "yearn_v3",
            "type": "redeem",
            "values": [
                {
                    "label": "tokenIn",
                    "value": "0xBe53A109B494E5c9f97b9Cd39Fe969BE68BF6204",
                    "type": "address",
                    "options": [
                        {
                            "label": "USDC",
                            "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                            "icon": "https://.../"
                        },
                        {
                            "label": "USDT",
                            "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                            "icon": "https://.../"
                        },
                        {
                            "label": "DAI",
                            "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                            "icon": "https://.../"
                        }
                    ],
                    "connector": null
                },
                {
                    "label": "tokenOut",
                    "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                    "type": "address",
                    "options": null,
                    "connector": 0
                },
                {
                    "label": "amountIn",
                    "value": 20,
                    "type": "uint256",
                    "options": null,
                    "connector": null
                }
            ]
        }
    ]
}
```

Now, this is a fully inclusive schema that contains both the app request and response. In practice, the app will want a way to get the options, connectors, etc. while only having to submit the required information.

For the request to get the information we would do something like:

```json
GetIntentRequest

{
    "chainId": 1,
    "protocol": "yearn_v3",
    "type": "redeem",
}

GetIntentResponse

{
    "chainId": 1,
    "protocol": "yearn_v3",
    "type": "redeem",
    "values": [
        {
            "label": "tokenIn",
            "type": "address",
            "options": [
                {
                    "label": "USDC",
                    "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                    "icon": "https://.../"
                },
                {
                    "label": "USDT",
                    "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                    "icon": "https://.../"
                },
                {
                    "label": "DAI",
                    "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                    "icon": "https://.../"
                }
            ],
            "connector": null
        },
        {
            "label": "tokenOut",
            "type": "address",
            "options": null,
            "connector": 0
        },
        {
            "label": "amountIn",
            "type": "uint256",
            "options": null,
            "connector": null
        }
    ]
}
```

For the post back to the solver, we will store the actions like:

```json
PostIntent

{
    "chainId": 1,
    "from": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "solver": "0x62180042606624f02d8a130da8a3171e9b33894d",
    "actions": [
        {
            "protocol": "yearn_v3",
            "type": "redeem",
            "values": [
                {
                    "label": "tokenIn",
                    "value": "0xBe53A109B494E5c9f97b9Cd39Fe969BE68BF6204"
                },
                {
                    "label": "tokenOut",
                    "value": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
                },
                {
                    "label": "amountIn",
                    "value": 20
                }
            ]
        }
    ]
}
```

Unhandled we need to return the sentence from the solver api to the app.

Then, the response from the solver with the final intent transaction will be:

```json


```
