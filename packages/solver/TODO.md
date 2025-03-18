## 03/18/2025:

- [x] Get coils rendering in the app properly
- [x] Get chain filtering functioning again
- [ ] Make sure the app supports having multiple types defined for user
   inputs and then values that will come from coils
   NOTES: For this, there are certain sentences where we need the original union 
          functionality that we already implemented. In this case though what we 
          really want is a pipe denoting the definition of a coil type instead of
          a union that allows the user to input multiple types.
   NOTES: Need to make sure that there are no sentences that break this convention
          before implementing it because this will require an update of some kind
          to cord due to the way we are handling validation right now.
- [ ] Remove the value validation from inside plug handlers
- [ ] Wrap the handlers with consumption of coils
- [ ] Refactor the actions directory into protocols
- [ ] Balance function call for native assets
- [ ] Append the coil definitions to a submit intent
- [ ] Build the proper coils for the onchain transaction
- [ ] Update plugs to be defined with the type of call to be made
   - [ ] When one is not explicitly set we should default solve for
      call or call with value.
- [ ] Fix the signatures to support the new shape
- [ ] Handle any other regressions we have created in the process of this update
- [ ] Clear any linked inputs if the parent output is removed -- Automatic reconnection here would be nice
