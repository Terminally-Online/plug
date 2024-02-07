import{_ as e,o as t,c as n,R as o}from"./chunks/framework.OBNbGUgv.js";const m=JSON.parse('{"title":"","description":"","frontmatter":{},"headers":[],"relativePath":"todo.md","filePath":"todo.md","lastUpdated":1702927677000}'),a={name:"todo.md"},i=o(`<blockquote><p>Hello, you have stumbled upon my working todo for Plug. It lives in the documentation here because I bounce from problem to problem without a real scope defined for the territory it resides within. I acknowledge this project is not setup in a way that enables simple contributions. That is intentional for now. Focusing on what matters and that is not a nicer README. That said, below you can see my todo, active notes and things as they were completed.</p></blockquote><ul><li><p>x: Done</p></li><li><p>h: Holding</p></li><li><p>o: Started</p></li><li><p>TODO:</p><ul><li><p>ALPHA:</p><ul><li><p>Protocol:</p><ul><li>[o] Deprecate all permission delegation functionality.</li><li>[ ] Implement a mechanism that allows us to mine the addresses for contracts that are deployed.</li></ul></li><li><p>Canvas:</p><ul><li>[x] Add distance constraint so that dragging does not impact clicking on components.</li><li>[ ] Move Pins out from the Plug component container so that they themselves are draggable. NOTE: Plugs is just a string of linked pins, not really a single component. NOTE: We basically already have everything setup for this, we just need to move the dragging one component lower.</li><li>[ ] Pin appendages placed on the head of a pin. NOTE: The connector will always be on the right side for output, and left side for linked pins. NOTE: Critical to note here is that the input of the following pin is not consumed. <ul><li>[ ] Each pin will have a starter and ender connector.</li><li>[ ] If the item is a <code>then</code>, it will not have an ender connector as it is the &quot;end&quot;.</li><li>[ ] Walk backwards through the linkage of a Sign pin to build the linearized loop.</li></ul></li><li>[ ] Draw colored lines between the linked pins. NOTE: We will just use a random color.</li><li>[ ] Bezier curves that connect each Pin in a Plug. NOTE: This is just a basic SVG path.</li><li>[ ] Add offset to the connection of the paths so that they do not look weird when dragged and overlapping.</li><li>[ ] Store the configuration of a Pin in the database.</li><li>[ ] Store the linkage in the database.</li><li>[ ] Make the grid mostly transparent with a mouse effect that only shows the grid around where the mouse is.</li><li>[ ] Given the capabilities of this app it is now probably worth looking into determining if we can pass in the inputs of previous pins inside the protocol. NOTE: Right now we are facing stack too deep issues, but we can remove a good amount of that if we require the use of a linearized loop instead of using the nested for loop architecture that is in place right now.</li><li>Drag controls: <ul><li>[ ] Right now when dropping a plug with multiple pins it is no longer on the grid. NOTE: I am not actually going to worry about this because each pin is going to be moved indivdually.</li><li>[ ] Camera Controls <ul><li>[ ] Make CMD + 0 reset the scale to 100%.</li><li>[ ] Recenter the camera requires moving the mouse in the grid to trigger the next render because the camera controls are outside of react render loop.</li><li>[x] Make sure items outside of bounds disappear.</li><li>[ ] CMD + Space + Drag to move the camera</li><li>[ ] Right now we can zoom in on elements that we should not be able to. NOTE: I think this started happening when we updated to the Pages Router, but not completely sure. When debugging this you will want to comment out the Toolbar so that you can determine if it is that element causing the issue beforedoing anything else. NOTE: When we break the zooming, the grid gets larger as well even though it is a fixed element which leads me to believe that we are zooming in on the parent element and this should not be happening.</li><li>[ ] Zooming out messes up the inbounds calculations.</li></ul></li><li>[ ] Item selection functionality. <ul><li>[ ] Be able to select a component.</li><li>[ ] Be able to delete it and select a group.</li><li>[ ] Drag selection.</li></ul></li></ul></li></ul></li><li><p>Templates &amp; Tabs:</p><ul><li>[x] Templates page <ul><li>[x] Tab Manager</li></ul></li><li>[ ] Move the tab logic into the database and use websockets to update everything.</li><li>[ ] Right now you can add too many tabs for the tab manager to display.</li></ul></li><li><p>Fuse Integration:</p><ul><li>[ ] Fuse integration. <ul><li>[ ] Revocation on deletion after the first signature has been posted to the database.</li></ul></li><li>[ ] Update the hello world to be the base noun plug.</li></ul></li><li><p>Deployment:</p><ul><li>[x] Get the a compiling build and register all the needed pipelines to cache when possible. CONCLUSION: Next through a huge fight when it came time and I am still not sure how I fixed it.</li><li>[x] Hand over dependency management to dependabot as we reach a state of not touching certain pieces anymore. CONCLUSION: Also got to include the submodules.</li><li>[x] Get landing page and app live in a staging environment. NOTE: This will not be a real staging environment, but it will be until we kill the landing app and move everything over which will happen come time of the first release.</li><li>[ ] Get the websocket running in the staging environment.</li></ul></li><li><p>[ ] Deprecate <code>packages/landing</code> NOTE: Do not do this until you are ready to roll out alpha because it replaces the early access signup with an enter app button.</p></li><li><p>[ ] Deprecate <code>packages/client</code> CONCLUSION: This actually is not going to be done because it will be a reference point for those that would like to build their own interface / provide an alternative to the BUSL licensed app.</p></li><li><p>[ ] Deprecate <code>packages/server</code> NOTE: Not really sure what is going to happen to this actually.</p></li></ul></li><li><p>BETA:</p><ul><li>[ ] Improve authentication implementation to support the use of email based accounts. NOTE: Realistically, there is no reason to limit to hot-wallets meaning wallets that are yet to be cold-started should still be able to architecture setups as well as sign orders. NOTE: There is also privy, but pricing is way more expensive. Imagine paying hundreds just for logging in alone. Bad. <ul><li>[ ] Add the ability to be an anonymous user (until signed in) <ul><li>[x] Implement the page redirects as a middleware so that we do not ever even get routed there. The reason this has not just been done is because I do not want to make two calls to the database (1 in the middleware and then another on the page) to determine if it is public / the accessing user has access so just not worrying about this for now. In MVP just going to design for public, but still require that users be logged in. <ul><li>[ ] When we do this, we will want to categorize them as a &#39;Anonymous Cow/Cat/Dog&#39; and show them as active on the board when they are viewing it.</li><li>[ ] For anonymous users, everything will always be read-only.</li></ul></li></ul></li></ul></li><li>[ ] Add &#39;home&#39; button to center the screen on the center of the board.</li><li>[ ] Drag and drop tab re-ordering. NOTE: I do not really care about this but I am sure someone will.</li></ul></li></ul></li></ul><p>PENDING:</p><pre><code>- [o] Finish the documentation.
  - [x] Remove the index page.
  - [x] Make the branding confirm to the branding that has now been defined.
  - [ ] Edit all of the static documentation and write the content for pieces that are not auto-generated.
</code></pre><p>BUGS:</p><pre><code>- [ ] Right now tabs will run into issues when you reach the edge of the page.
      NOTE: I am pretending this isnt happening and will get to it when I get to it.
- [ ] If you do not have any canvases in the database yet, create one, and then go back to the home page, if you then try and create another new one it forces you back to the newly created one instead of allowing you to access the page as you should because you no longer have zero canvases.
- [ ] Right now in search, if you hot reload the page or tab out, the api query results are re-added to the list.
      NOTE: I messed with lots of things and could not get this to stop happening. Would like to avoid the use of a watching useEffect that clears the list because I do not think it is needed.
</code></pre><p>DONE:</p><pre><code>- [x] ALPHA:
  - [x] Protocol:
    - [x] Add a passthrough lane to the fuses so that data can be passed through them.
          NOTE: This will be a sidecar to the active data of the fuse so that it can maintain its conditions while maybe manipulating the pass through.
    - [x] Implement a test for this.
    - [x] Fuse development.
      - [x] Add noun trait bid.
      - [x] Expose the contracts as exports from @nftchance/plug-core
    - [x] Replace the use of constructors with an internal \\_initializeSocket()
    - [x] Vault implementation.
          CONCLUSION: It turns out there still isn&#39;t a commonly agreed on vault implementation because they are designed to do very different things. - [ ] While we are not going to use ERC4626, it will be worthwhile to look into and determine if there are any features that need to be yoinked.
    - [x] Factory implementation for vaults and other contracts.
          CONCLUSION: This architecture was settled on because it allows us to reuse the same factory for all of our deployments while providing a simple interface that can be used for vanity address mining. While it is highly unlikely that we will do any address mining for a user, it will be worthwhile to do so for our fuses. Especially if we can get a cool vanity tag.
    - [x] Finalize the implementation of changes to the verification implementation.
      - [x] .forced was in the process of being implemented before I had to pivot and migrate away from the use of constructors.
    - [x] Confirm the protocol is not vulnerable to the same kind of attack as: https://blog.openzeppelin.com/arbitrary-address-spoofing-vulnerability-erc2771context-multicall-public-disclosure
          NOTE: While the architecture of the protocol in general is quite similar to what lead to the vulernability being possible, I do not think it is possible because we have force-resolved the sender at all times and never assume the response from a Fuse or execution can be trusted.
          NOTE: Actually, I think we may also be vulnerable because you would just replace \`multicall\` with \`plug\`

  - [x] API:
    - [x] There is an issue with EventEmitter where procedures are refusing to communicate with one another. In both routers, when making a subscription through an event emitter the vent never fires. The weird thing is that when I am not using an emitter, everything works fine with just an observable but the second I try and make an onUpdate it doesnt ever fire the event.
          NOTE: I have tried so many things and I am tired of wasting time on this. Everything else is setup fine and subscriptions that are not going through emitters work fine so just cannot justify wasting any more time on this at this point. If people complain because I got users then this will be taken care of immediately!
    - [x] Update \`ws.prod.ts\` to actually be ready for production.
    - [x] Improved account page.
    - [x] Infinite query to retrieve the canvases of an account.
    - [x] Improved /create experience to introduce grace period before actual creation.
    - [x] Searching canvases on the account page.
    - [x] Make the websocket subscriptions relative to the context of the subscription rather than being a general catch-all subscription.
          NOTE: Creation events should only notify of ones own account.
          NOTE: Update events should only notify of ones own account.
    - [x] Searching does not cause the results to refresh from the database even though the search params changed.

  - [x] Canvas:
    - [x] onCreate
    - [x] refactor canvas.onCreate -&gt; canvas.onAdd
    - [x] onUpdate
    - [x] Components:
      - [x] onAdd
      - [x] onMove
    - [x] Replace the placeholder component with a plug.
    - [x] The positioning of new components from addComponent is a little wonky.
          NOTE: Spent a little time trying to fix it, got it good enough for now. Will want to figure out what is going on here and fix it though.
      - [x] Make sure a connector between each pin is on the grid.
    - [x] Control:
      - [x] Add zoom buttons so that you don&#39;t have to scroll.
        - [x] Fix the zoom having adverse effects on react-dnd.
          - [x] If we cannot fix this, just disable zooming until we move to a proprietary dragging implementation. Not ideal but not sure what other option there event is.
                NOTE: There have been several issues dating all the way back to 2015 that have never been resolved.
                NOTE: There has not been a commit in 10 months. Think I made the wrong choice here.
                https://github.com/react-dnd/react-dnd/issues?q=is%3Aissue+is%3Aopen+zoom
        - [x] Better drag controls
          - [x] Replace react-dnd with https://docs.dndkit.com/introduction/installation
            - [x] Get the items lining up with the grid.
                  NOTE: My base assumption that it is some padding causing this issue.
                  CONCLUSION: Just some extra margin that was being placed because of the weird \`module.css\` files that are still being used during the migration process away from react-dnd.
            - [x] Get back the infinite canvas movement.
                  NOTE: Should the grid move when we move the camera?
              - [x] Make sure the grid works inside the infinite canvas.
              - [x] Connect the two pieces.
            - [x] Move components individually.
            - [x] Save the position of a dragged item in the canvas.
            - [x] Make sure that it works even when are zoomed out.
            - [x] Make sure that we can only drag items within the bounds of the window.
          - [x] Reenable the ability to add a component.
            - [x] Add a debug identifier for the location of each component.
            - [x] Make sure the addition of components is ailgned to the grid.
                  NOTE: I am really not sure what is going on here. Rather confused about this.
                  NOTE: In the process of solving this and while the coordinates in the database appear to be right, the coordinates being shown on the elements are not accurate to the layout. I cannot tell if this is because we have collisions between the two elements, but I do not think so.
                  CONCLUSION: I randomly fixed this. I expect it to come back. Right now the pointer coordinates are off by a very weird non-standard amount.
            - [x] Trigger a render when a new component is added.
      - [x] Get infinite canvas functional.
            NOTE: It is really rather simple once you understand what an infinite canvas actually is. Especially because we do not even actually care about it being infinite for real.
      - [x] Drag components.
            NOTE: We should just use react-dnd so that we do not have to worry about building our own drag implementation.
        - [x] Align items to the grid (grid snapping).
          - [x] Infinite canvas and drag and drop functional at the same time.
          - [x] Prevent a user from selecting the same pin multiple times in a plug.
          - [x] Prevent a user from adding a new pin when all available pins have been exhausted.
          - [x] Page where a user can see their canvases
          - [x] Tab functionality
          - [x] Active canvas
          - [x] Close a tab
          - [x] Add a tab
          - [x] Make adding a new tab direct to /canvas/create/
          - [x] Do not allow having multiple of the same tab open.
          - [x] Figure out what &#39;closed connection&#39; is originating from when signing in.
          - [x] Fully functional authentication implementation
          - [x] Implement a basic auth stack so that we can the basic api flowing.
                NOTE: In development we should just use a mock auth provider and ignore this piece for now.
          - [x] Get the basic authentication ux designed and boiled.
          - [x] Implement everything that is needed to actually run the websocket server and make sure everything is running properly.
          - [x] When this is in place bundle it up into \`pnpm dev\` to run everything at once
          - [x] Run a websocket server to power trpc
          - [x] Get user canvases from the database.
          - [x] Pick back up by finishing the conversion to page router. \`layout.tsx\` is the next thing to figure out since it is handled differently.
          - [x] Sync state changes to the database.
          - [x] Easy framework to add and remove component types.
          - [x] Store the state of the canvas into the database.
          - [x] When we sync the state maybe we shouldn&#39;t even worry about saving when a component is updated and instead operate on just a set interval to check if changes have been made since the last broadcast and stream out the changes to the database. Realisticaly, my wanting to support websockets did not arise from the need of realtime collaboration. It was simply there to solve the issues that would arise when someone has two tabs open or someone is viewing a board that is not theirs.
          - [x] Retrieve the state of the canvas from the database.
            - [h] Need to get live updates in case we have multiple windows open.
              NOTE: While it may seem like this is premature optimization, after having experience the difficulties of adding live responses without having properly planned for it.
          - [x] Figure out how to store complex component state in the canvas.
                NOTE: We also have superjson as an option thought it will not completely solve the problem so it is likely not the right choice.
      - [x] Final nomenclature update
        - [x] Update \`authority\` nomenclature to just use \`pin\` for increased conistency.
        - [x] &quot;Naming of things doesn&#39;t matter!&quot; -- Anyone that says this has never actually launched something in this industry and if they have and didn&#39;t learn this lesson then that is their own problem.
      - [x] Figure out what to do about naming.
        - [x] Get the domain
        - [x] Deploy on the new domains
        - [x] My favorite idea so far has been Plug cause 1) it is plug and play 2) it makes contracts plug and play with one another. It is also really short and I can get a domain that&#39;s easy to navigate to. Also, it&#39;s so simple that I finally have logo ideas.
        - [x] Rename the Github repositories
        - [x] Rename the Discord
        - [x] Deprecate the old npm packages
        - [x] Deploy documentation at docs.onplug.io
        - [x] Week 1 report
      - [x] Setup Supabase to be used for the production database on the landing site.
        - [x] How are you supposed to integrate schema generation into your CI/CD?
      - [x] Deploy to Vercel
      - [x] Make sure our API endpoints are protected
      - [x] Make sure the landing page is responsive.
      - [x] Get the raw client implementation functional with raw trpc connections.
      - [x] Package up the trpc connector into the sdk to offer a more explicit integration path.
            NOTE: This is where the \`process.env.API_URL\` and managed functions would be exposed.
      - [x] Make sure that we can encode and decode each enforcer.
</code></pre>`,8),s=[i];function l(r,h,d,c,u,p){return t(),n("div",null,s)}const w=e(a,[["render",l]]);export{m as __pageData,w as default};
