## 03/18/2025:

- [x] Get coils rendering in the app properly
- [x] Get chain filtering functioning again
- [x] Make sure the app supports having multiple types defined for user
   inputs and then values that will come from coils
   NOTES: For this, there are certain sentences where we need the original union 
          functionality that we already implemented. In this case though what we 
          really want is a pipe denoting the definition of a coil type instead of
          a union that allows the user to input multiple types
   NOTES: Need to make sure that there are no sentences that break this convention
          before implementing it because this will require an update of some kind
          to cord due to the way we are handling validation right now
- [x] Confirm the data is actually in the intent being saved
- [x] Wrap the handlers with consumption of coils
- [x] Remove the value validation from inside plug handlers
   NOTES: This was not actuallly implemented in many of the integrations except for the original
      ones because after a couple I realized that validation was a bad pattern anyways since it
      is all going to go to simulation anyways meaning the only other case is to hit a 400 error.
- [x] Balance function call for native assets

## 03/19/2025

- [x] Figure out why the readme of core does not include Plug.EVM.sol
- [x] Fix the release action so that I do not have to keep doing manual releases
- [ ] Update plugs to be defined with the type of call to be made
   - [ ] When one is not explicitly set we should default solve for
      call or call with value.
- [?] Build the proper coils for the onchain transaction
- [?] Append the coil definitions to a submit intent
   NOTES: This may already be done because I skipped to focusing on getting plugs e2e instead of 
          focusing on the other things I wanted to take care of right now.
- [ ] Refactor the actions directory into protocols
- [ ] Clear any linked inputs if the parent output is removed -- Automatic reconnection here would be nice

## 03/20/2025
- [ ] Fix the signatures to support the new shape
