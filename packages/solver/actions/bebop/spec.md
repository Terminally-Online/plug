Bebop Integration

## Overview

Bebop enables Plug users to make swaps and to check prices. The PMM API is used to get access to quotes and fills from professional market makers. The JAM API enables users to swap tokens not supported in the PMM API by accessing solver auctions.
## Supporting Documentation

- [API Comparison](https://docs.bebop.xyz/bebop/bebop-api-pmm-rfq/bebop-trading-apis-comparison)
- [Getting Quotes](https://docs.bebop.xyz/bebop/bebop-api-pmm-rfq/rfq-api-endpoints/trade/get-quote )

---
## Contract Interfacing

Bebop can be interacted with directly via the API.
  
## Scope

  
| System  | Name                 | Type       | Implemented | Notes                             |
| :------ | :------------------- | :--------- | :---------- | :-------------------------------- |
| PMM API | Swap At Least Tokens | Action     |             | Use JAM API if tokens not covered |
| PMM API | Market Swap          | Action     |             | Use JAM API if tokens not covered |
| PMM API | Check Price USD      | Constraint |             | Use JAM API if tokens not covered |
| PMM API | Check Price Custom   | Constraint |             | Use JAM API if tokens not covered |

## Swaps

While Plug offers different types of Swaps to the user, under the hood these swaps share many of the same actions.

[Retrieve a quote](https://docs.bebop.xyz/bebop/bebop-api-pmm-rfq/rfq-api-endpoints/trade/get-quote) The quote will indicate a price given the trade parameters that you have specified. Use https://api.bebop.xyz/pmm/arbitrum/docs#/v3/quote_v3_quote_get

Before swapping, we must [make approvals](https://docs.bebop.xyz/bebop/bebop-api-pmm-rfq/rfq-api-endpoints/trade/manage-approvals) In order for Bebop to facilitate the trade, it needs to have access to your tokens. Approvals can be made with standard ERC20 Approvals or Permit2. 

To use Permit2, specify `approval_type=Permit2` when retrieving a quote. The quote endpoint will return a `requiredSignatures` field with token addresses that you must provide permit signatures for.

You will need to make `ERC20.approve` allowances to the Permit2 contract(`0x000000000022D473030F116dDEE9F6B43aC78BA3`) before submitting an order.

You will need to create a permit2 signature for the tokens indicated in `quote.requiredSignatures` and submit this signature when calling `POST /order`

Sign the order and retrieve signatures from the user wallet (gasless). [Submit your signature alongside approval parameters to the order endpoint](https://docs.bebop.xyz/bebop/bebop-api-pmm-rfq/rfq-api-endpoints/trade/submit-order) (gasless).

Swaps should check if the pair is included in the ~50 tokens supported by the PMM API, if not the JAM API must be used.
### Swap At Least Tokens

"Swap {0<amountIn:float>} {1<tokenIn:address>} for at least {2<amountOut:int>} {3<tokenOut:address>}."

Users can input any token and amount for token in and token out. We could offer users Swap Exact Tokens, but this is a better user experience as users may get more than they are expecting back.

Plug can check quotes and fill this request if the amount is greater than or equal to the user requested amount.
### Market Swap

"Swap {0<amountIn:float>} {1<tokenIn:address>} for {2<tokenOut:address>}."

Users can input any token and amount for token in and only the token for token out. This swap would occur at market rate based on the quote given by Bebop.
### Check Price USD

"{0<amountIn:float>} {1<tokenIn:address>} is worth at least {2<dollar:float>} USD."

Users can input any token and amount to see if it is worth a USD amount. This can be done with the request for quote endpoint offered by Bebop.
### Check Price Custom

"{0<amountIn:float>} {1<tokenIn:address>} is worth at least {2<amountOut:int>} {3<tokenOut:address>}."

Users can input any token and amount to see if it is worth a different token or amount. This can be done with the request for quote endpoint offered by Bebop.
